from fastapi import FastAPI
from pydantic import BaseModel
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch
import requests
import json
import re
import unicodedata
from typing import Optional, List
import spacy
import os

app = FastAPI()
user_context = {}

# Load Qwen model and tokenizer
qwen_tokenizer = AutoTokenizer.from_pretrained("Qwen/Qwen1.5-0.5B-Chat", trust_remote_code=True)
qwen_model = AutoModelForCausalLM.from_pretrained("Qwen/Qwen1.5-0.5B-Chat", trust_remote_code=True).eval()

# Load spaCy English model for food extraction
nlp = spacy.load("en_core_web_sm")

MISTRAL_API_KEY = os.getenv("MISTRAL_API_KEY") or "2fgXh3wnRm2MZlIvScsKGVf5NAupcXS4"

class Prompt(BaseModel):
    session_id: str
    prompt: str
    additional_prompt: Optional[str] = ""

class FoodInput(BaseModel):
    food_name: str

FOOD_KEYWORDS = {
    'food', 'meal', 'dish', 'fruit', 'vegetable', 'protein', 'cheese',
    'chicken', 'fish', 'bread', 'egg', 'salad', 'soup', 'burrito', 'pancake',
    'smoothie', 'salmon', 'tofu', 'tempeh', 'nuts', 'berries', 'oatmeal',
    'granola', 'yogurt', 'bacon', 'wrap', 'sandwich', 'toast', 'bowl',
    'pizza', 'rice', 'beef', 'pork', 'shrimp', 'seafood', 'lentils', 'beans'
}

GENERIC_FOOD_BLACKLIST = {
    "some breakfast food options", "foods", "a bowl", "bowl", "some food options",
    "food", "meal", "dishes", "options", "breakfast", "lunch", "dinner",
    "some dishes", "these dishes", "some chopped vegetables", "your meal",
    "your dish", "these foods"
}

def sanitize_text(text: str) -> str:
    # Remove non-ASCII chars, e.g. Chinese or emojis
    return unicodedata.normalize('NFKD', text).encode('ascii', 'ignore').decode('ascii')

def parse_prompt_with_mistral(user_input: str) -> str:
    url = "https://api.mistral.ai/v1/chat/completions"
    headers = {
        "Authorization": f"Bearer {MISTRAL_API_KEY}",
        "Content-Type": "application/json"
    }
    payload = {
        "model": "mistral-small",
        "messages": [
            {
                "role": "system",
                "content": "You are a helpful assistant that converts informal Bahasa Indonesia into clear, pure English prompts without foreign words or mixed language."
            },
            {
                "role": "user",
                "content": user_input
            }
        ]
    }
    response = requests.post(url, headers=headers, json=payload)
    if response.status_code == 200:
        return response.json()["choices"][0]["message"]["content"].strip()
    return f"[Mistral API Error {response.status_code}] {response.text}"

def generate_response_from_qwen(prompt: str, max_tokens: int = 512) -> str:
    input_text = f"<|user|>\n{prompt}\n<|assistant|>\n"
    input_ids = qwen_tokenizer.encode(input_text, return_tensors="pt")
    with torch.no_grad():
        output = qwen_model.generate(input_ids, max_new_tokens=max_tokens)
    decoded = qwen_tokenizer.decode(output[0], skip_special_tokens=True)
    return decoded.split("<|assistant|>")[-1].strip() if "<|assistant|>" in decoded else decoded.strip()

def extract_food_names_via_spacy(text: str) -> List[str]:
    doc = nlp(text.lower())
    noun_phrases = [chunk.text.strip() for chunk in doc.noun_chunks]
    filtered_phrases = [phrase for phrase in noun_phrases if any(word in phrase for word in FOOD_KEYWORDS)]
    return list(set(filtered_phrases))

def clean_food_names(food_names: List[str]) -> List[str]:
    cleaned = []
    for f in food_names:
        f_clean = f.strip().lower()
        f_clean = re.sub(r"^\d+\.\s*", "", f_clean)  # Remove numbered list prefix
        if re.search(r"\d", f_clean):
            continue
        if f_clean in GENERIC_FOOD_BLACKLIST:
            continue
        if len(f_clean) < 3:
            continue
        cleaned.append(f_clean)
    return list(set(cleaned))

@app.post("/generate")
async def generate(prompt: Prompt):
    session_id = prompt.session_id
    user_input = prompt.prompt
    extra = prompt.additional_prompt.strip() if prompt.additional_prompt else ""

    mistral_clean = parse_prompt_with_mistral(user_input)
    mistral_clean = re.sub(r"```.*?```", "", mistral_clean, flags=re.DOTALL).replace("Sure!", "").replace("Here's your prompt:", "").strip()

    final_input = f"{mistral_clean} {extra}".strip()
    previous = user_context.get(session_id)
    final_prompt = f"{previous} {final_input}".strip() if previous else final_input
    user_context[session_id] = final_prompt

    response_raw = generate_response_from_qwen(final_prompt)
    response_clean = sanitize_text(response_raw)

    food_names_raw = extract_food_names_via_spacy(response_clean)
    cleaned_food_names = clean_food_names(food_names_raw)

    return {
        "session_id": session_id,
        "original_input": user_input,
        "processed_prompt": mistral_clean,
        "additional_prompt": extra,
        "final_combined_prompt": final_prompt,
        "response": response_clean,
        "food_names": cleaned_food_names
    }

@app.post("/generate-nutrition")
async def generate_nutrition(data: FoodInput):
    prompt = f"""Format the result in this JSON structure:

{{
  "food": "Tempe Goreng",
  "nutrients": [
    {{
      "name": "calories",
      "size": "200kcal",
      "variants": ["Iron", "Magnesium"]
    }},
    {{
      "name": "protein",
      "size": "15g",
      "variants": []
    }}
  ]
}}

Now do the same for this food:

Food: {data.food_name}
"""
    input_text = f"<|user|>\n{prompt}\n<|assistant|>\n"
    input_ids = qwen_tokenizer.encode(input_text, return_tensors="pt")
    with torch.no_grad():
        output = qwen_model.generate(input_ids, max_new_tokens=512)
    decoded = qwen_tokenizer.decode(output[0], skip_special_tokens=True)
    response_text = decoded.split("<|assistant|>")[-1].strip() if "<|assistant|>" in decoded else decoded.strip()

    if "```" in response_text:
        response_text = response_text.replace("```json", "").replace("```", "").strip()

    try:
        first_pass = json.loads(response_text) if response_text.startswith('{') else response_text
        response_json = json.loads(first_pass) if isinstance(first_pass, str) else first_pass
    except Exception as e:
        response_json = {
            "raw_output": response_text,
            "warning": f"Output is not valid JSON after cleanup: {str(e)}"
        }

    return {
        "food_name": data.food_name,
        "response": response_json
    }

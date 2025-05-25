package requests

type Medical struct {
	Allergies     []string `json:"allergies"`
	MedicalRecord []string `json:"medicalRecord"`
}

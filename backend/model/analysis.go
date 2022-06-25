package model

type Analysis struct {
	Title           string `json:"title"` // Input from user or calculate from FunctionDeclare
	FunctionDeclare string `json:"function_declare"`
	GitUrl          string `json:"git_url"`
	Note            string `json:"note"`
	Code            string `json:"code"`
}

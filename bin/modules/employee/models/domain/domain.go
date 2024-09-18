package models

type EmployeeRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	JobTitle string `json:"job_title"`
	Company  string `json:"company"`
}

type EmployeeGetByIDRequest struct {
	ID string `json:"id"`
}

type EmployeeResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	JobTitle string `json:"job_title"`
	Company  string `json:"company"`
}

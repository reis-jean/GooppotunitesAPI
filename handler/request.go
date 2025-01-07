package handler

import "fmt"

// funções generica para retornar erro de campos obrigatorios
func errParamsIsRequired(name, typ string) error{
	return fmt.Errorf("param: %s (type: %s) is Required", name, typ)
}

type CreateOpeningRequest struct {
	Role    string `json: "role"`
	Company string `json: "company"`
	Location string `json: "location"`
	Remote  *bool  `json: "remote"`
	Link    string `json: "link"`
	Salary  int64  `json: "salary"`
}

func (r *CreateOpeningRequest) validate() error {

	
	if r.Role == "" {
		return errParamsIsRequired("role", "string")
	}
	if r.Company == "" { 
		return errParamsIsRequired("company", "string")
	}
	if r.Remote == nil {
		return errParamsIsRequired("remote", "string")
	}
	if r.Link == "" {
		return errParamsIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return errParamsIsRequired("salary", "int")
	}

	return nil 
}
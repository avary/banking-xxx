package dto

type NewCustomerRequest struct {
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
}

type CustomerUpdateRequest struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Status  string `json:"status"`
}

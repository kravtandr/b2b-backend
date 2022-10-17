package domain

type FastRegistrationForm struct {
	Name       string `json:"name"`
	LegalName  string `json:"legal_name"`
	Itn        string `json:"itn"`
	OwnerName  string `json:"owner_name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Post       string `json:"post"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Country    string `json:"country"`
}

type Search struct {
	Name string `json:"name"`
}

type ValidateErrors struct {
	Errors []string `json:"erros"`
}

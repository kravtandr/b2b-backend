package domain

type Company struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	LegalName    string   `json:"legal_name"`
	Itn          string   `json:"itn"`
	Psrn         string   `json:"psrn"`
	Address      string   `json:"address"`
	LegalAddress string   `json:"legal_address"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Link         string   `json:"link"`
	Activity     string   `json:"activity"`
	OwnerId      int      `json:"owner_id"`
	Rating       int      `json:"rating"`
	Docks        []string `json:"docs"`
}
type CompanySearch struct {
	Name string `json:"name"`
}

type Employee struct {
	Post       string `json:"post"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Group_id   int    `json:"group_id"`
}

type Employees []Employee

type Companies []Company

type CompanyStorage interface {
	Add(value Company) error
	AddBaseCompany(value Company, post string) error
	CompaniesUsersLink(value Company, post string) error
	GetByEmail(key string) (value Company, err error)
	GetCompanyById(key string) (value Company, err error)
	GetCompanyEmployees(key string) (value Employees, err error)
	//SearchCompanies(key string) (value Companies, err error)

	//GetCompaniesByCategoryId(key string) (value Companies, err error)

}

type CompanyUseCase interface {
	Add(value *Company) error
	AddBaseCompany(value *Company, post string) error
	GetByEmail(key string) (value Company, err error)
	GetCompanyById(key string) (value []byte, err error)
	GetCompanyEmployees(key string) (value []byte, err error)

	//SearchCompanies(key CompanySearch) (value []byte, err error)
	//Validate(company *Company) bool
}

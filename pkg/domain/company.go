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

	Password string `json:"password"`
}
type CompanySearch struct {
	Name string `json:"name"`
}

type Companies []Company

type CompanyStorage interface {
	Add(value Company) error
	GetByEmail(key string) (value Company, err error)
	//SearchCompanies(key string) (value Companies, err error)
	//GetCompanyById(key string) (value Company, err error)
	//GetCompaniesByCategoryId(key string) (value Companies, err error)

}

type CompanyUseCase interface {
	Add(value *Company) error
	GetByEmail(key string) (value Company, err error)
	//SearchCompanies(key CompanySearch) (value []byte, err error)
	//GetCompanyById(key string) (value []byte, err error)
	//GetCompaniesByCategoryId(key string) (value []byte, err error)
	//Validate(company *Company) bool
	//Login(company *Company) (int, error)
	//Registration(company *Company) (int, error)
}

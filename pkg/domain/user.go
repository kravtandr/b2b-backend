package domain

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Country    string `json:"country"`
	GroupId    int    `json:"group_id"`
}

type Users []User

type PublicUser struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
}

type UserEmail struct {
	Email string `json:"email"`
}

type UserStorage interface {
	Add(user User) error
	//AddCompany(company Company) error
	GetByEmail(key string) (value User, err error)
	GetPublicUserByEmail(key string) (value PublicUser, err error)
	GetPublicUserById(id string) (value PublicUser, err error)
	//SearchCompanies(key string) (value Companies, err error)
	//GetCompanyById(key string) (value Company, err error)
	//GetCompaniesByCategoryId(key string) (value Companies, err error)

}

type UserUseCase interface {
	Add(user User) error
	Login(user *User) (int, error)
	Validate(user *User) bool
	Registration(user *User) (int, error)
	//RegistrationCompany(company *Company) (int, error)
	GetByEmail(key string) (value User, err error)
	GetPublicUserByEmail(key string) (value []byte, err error)
	GetPublicUserById(id string) (value []byte, err error)
	//SearchCompanies(key CompanySearch) (value []byte, err error)

	//GetCompanyById(key string) (value []byte, err error)
	//GetCompaniesByCategoryId(key string) (value []byte, err error)

}

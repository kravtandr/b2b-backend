package domain

type OrderForm struct {
	Id               int    `json:"id"`
	Role             bool   `json:"role"`
	Product_category string `json:"product_category"`
	Product_name     string `json:"product_name"`
	Order_text       string `json:"order_text"`
	Order_comments   string `json:"order_comments"`
	Fio              string `json:"fio"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Company_name     string `json:"company_name"`
	Itn              string `json:"itn"`
}

type OrderForms []OrderForm

type PublicOrderForm struct {
	Id               int    `json:"id"`
	Role             string `json:"role"`
	Product_category string `json:"product_category"`
	Product_name     string `json:"product_name"`
	Order_text       string `json:"order_text"`
	Order_comments   string `json:"order_comments"`
	Fio              string `json:"fio"`
	Email            int    `json:"email"`
	Phone            string `json:"phone"`
	Company_name     string `json:"company_name"`
}

type OrderFormStorage interface {
	Add(orderForm OrderForm) error
	//AddCompany(company Company) error
	// GetByEmail(key string) (value User, err error)
	// GetPublicUserByEmail(key string) (value PublicUser, err error)
	// GetPublicUserById(id string) (value PublicUser, err error)
	//SearchCompanies(key string) (value Companies, err error)
	//GetCompanyById(key string) (value Company, err error)
	//GetCompaniesByCategoryId(key string) (value Companies, err error)

}

type OrderFormUseCase interface {
	Add(orderForm OrderForm) (int, error)
	// Login(user *User) (int, error)
	Validate(orderForm *OrderForm) bool
	// Registration(user *User) (int, error)
	//RegistrationCompany(company *Company) (int, error)
	// GetByEmail(key string) (value User, err error)
	// GetPublicUserByEmail(key string) (value []byte, err error)
	// GetPublicUserById(id string) (value []byte, err error)
	//SearchCompanies(key CompanySearch) (value []byte, err error)

	//GetCompanyById(key string) (value []byte, err error)
	//GetCompaniesByCategoryId(key string) (value []byte, err error)

}

package models

type Company struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	LegalName    string `json:"legal_name"`
	Itn          string `json:"itn"`
	Psrn         string `json:"psrn"`
	Address      string `json:"address"`
	LegalAddress string `json:"legal_address"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Link         string `json:"link"`
	Activity     string `json:"activity"`
	OwnerId      int64  `json:"owner_id"`
	Rating       int64  `json:"rating"`
	Verified     int64  `json:"verified"`
	Photo        string `json:"photo"`
	//Docks        []string `json:"docs"`
}
type CompanySearchByName struct {
	Name string `json:"name"`
}

type CompanySearchById struct {
	Id int64 `json:"id"`
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

type CompanyFullInfo struct {
	Company   Company   `json:"company"`
	Employees Employees `json:"employees"`
}

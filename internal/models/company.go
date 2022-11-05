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
	//Docks        []string `json:"docs"`
}

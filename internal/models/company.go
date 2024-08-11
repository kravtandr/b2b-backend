package models

import "b2b/m/internal/services/auth/models"

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

type CompanyUpdateProfileRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Address      string `json:"address"`
	LegalAddress string `json:"legal_address"`
	Itn          string `json:"itn"`
	Phone        string `json:"phone"`
	Link         string `json:"link"`
	Activity     string `json:"activity"`
}

type CompanyUpdateProfileResponse struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Address      string `json:"address"`
	LegalAddress string `json:"legal_address"`
	Phone        string `json:"phone"`
	Link         string `json:"link"`
	Activity     string `json:"activity"`
}

//тотальный кал, переделать

type CompanyWithCookie struct {
	Token        string `json:"token"`
	Cookie       string `json:"-"`
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

type CompanyUser struct {
	Post      string `json:"post"`
	CompanyId int64  `json:"company_id"`
	UserId    int64  `json:"user_id"`
	Itn       string `json:"itn"`
}

type CompanyAndOwner struct {
	Owner   models.User `json:"owner"`
	Company Company     `json:"company"`
	Post    string      `json:"post"`
}

type PublicCompanyAndOwnerRequest struct {
	Owner   UpdateUserRequest           `json:"owner"`
	Company CompanyUpdateProfileRequest `json:"company"`
	Post    string                      `json:"post"`
}

type PublicCompanyAndOwnerResponse struct {
	Owner   UpdateUserResponse           `json:"owner"`
	Company CompanyUpdateProfileResponse `json:"company"`
	Post    string                       `json:"post"`
}

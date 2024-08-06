package models

type Profile struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name" valid:"required"`
	Surname     string      `json:"surname" valid:"required"`
	Patronymic  string      `json:"patronymic"`
	Country     string      `json:"country"`
	Email       string      `json:"email"`
	Balance     int64       `json:"balance"`
	Avatar      string      `json:"avatar"`
	Description string      `json:"description"`
	Company     Company     `json:"company"`
	CompanyUser CompanyUser `json:"companyUser"`
}

type UpdateProfileRequest struct {
	Name        string `json:"name" valid:"required"`
	Surname     string `json:"surname" valid:"required"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type UpdateUserRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UpdateUserResponse struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
}

type PublicUser struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
}

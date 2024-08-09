package models

type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Country    string `json:"country"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	GroupId    int    `json:"group_id"`
	Balance    int64  `json:"balance"`
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

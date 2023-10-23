package models

type User struct {
	Id         int64  `json:"id"`
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

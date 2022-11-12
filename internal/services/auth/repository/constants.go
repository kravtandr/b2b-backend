package repository

const (
	getUserByEmailRequest  = "SELECT id, password FROM Users WHERE email = $1"
	getUserByIDRequest     = "SELECT id, name, surname, patronymic, email, password, group_id FROM Users WHERE id = $1"
	createUserRequest      = "INSERT INTO Users (name, surname, patronymic, email, password, country) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	createCreateUpdateUser = "UPDATE Users SET name = $2, surname = $3, patronymic = $4, email =$5 ,password = $6 WHERE id = $1 RETURNING id, name, surname, patronymic, email, password"
	//updateUserRequest     = "UPDATE Users SET "
	//updateUserName        = "name=$"
	//updateUserSurname     = "surname=$"
	//updateUserPass        = "password=$"
	//updateUserEmail       = "email=$"
	//updateUserDescription = "description=$"
	//updateUserAvatar      = "avatar=$"
	whereStatement = "WHERE id = $1"
	//updateUserReturning   = "RETURNING name, surname, email, avatar, description"
	createCompanyByUserId       = "SELECT id, name, description, legal_name, itn, psrn, address,legal_address,email, phone, link, activity, owner_id, rating , verified  FROM Companies WHERE owner_id = $1"
	createCreateCompany         = "INSERT INTO companies (name, legal_name, itn, email, owner_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	createCreateUserCompanyLink = "INSERT INTO companiesusers (post, company_id, user_id, itn) VALUES ($1, $2, $3, $4) RETURNING id"
	createUserSession           = "INSERT INTO Cookies (user_id, hash) VALUES ($1, $2)"
	validateUserSession         = "SELECT user_id FROM Cookies WHERE hash = $1"
	removeUserSession           = "DELETE FROM Cookies WHERE hash = $1"
	GetUserInfoQuery            = "SELECT id, name, surname FROM Users WHERE id = $1"
)

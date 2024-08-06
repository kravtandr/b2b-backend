package repository

const (
	createUserRequest      = "INSERT INTO Users (name, surname, patronymic, email, password, country) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	getUserByEmailRequest  = "SELECT id, name, surname,patronymic, email, password FROM Users WHERE email = $1"
	getUserByIDRequest     = "SELECT id, name, surname, patronymic, country, email, password, group_id, balance FROM Users WHERE id = $1"
	createCreateUpdateUser = "UPDATE Users SET name = $2, surname = $3, patronymic = $4, email =$5 ,password = $6 WHERE id = $1 RETURNING id, name, surname, patronymic, email, password"
	GetUserInfoQuery       = "SELECT id, name, surname FROM Users WHERE id = $1"

	createCreateUserCompanyLink = "INSERT INTO companiesusers (post, company_id, user_id, itn) VALUES ($1, $2, $3, $4) RETURNING id"
	createGetCompanyUserLink    = "SELECT  id, post, company_id, user_id, itn FROM CompaniesUsers WHERE user_id = $1 and company_id = $2"

	createCreateCompany   = "INSERT INTO companies (name, legal_name, itn, email, owner_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	createCompanyByUserId = "SELECT id, name, description, legal_name, itn, psrn, address,legal_address,email, phone, link, activity, owner_id, rating , verified  FROM Companies WHERE owner_id = $1"

	createUserSession   = "INSERT INTO Cookies (user_id, hash) VALUES ($1, $2)"
	validateUserSession = "SELECT user_id FROM Cookies WHERE hash = $1"
	removeUserSession   = "DELETE FROM Cookies WHERE hash = $1"

	createUpdateUserBalance = "UPDATE Users SET balance = $2 WHERE id = $1 RETURNING id, name, surname, patronymic, email, password"

	createAddPayment        = "INSERT INTO Payments (user_id, payment_id, amount, type) VALUES ($1, $2, $3, $4) RETURNING id, user_id, payment_id, amount, status, paid, type, credited, created_at"
	createGetPayment        = "SELECT id, user_id, payment_id, amount, status, paid, type, credited, created_at FROM Payments WHERE payment_id = $1"
	createUpdatePayment     = "UPDATE Payments SET user_id = $1, payment_id = $2, amount = $3, status = $4, paid = $5, type = $6, credited = $7 WHERE payment_id = $2 RETURNING id, user_id, payment_id, amount, status, paid, type, credited, created_at"
	createGetUserPayments   = "SELECT id, user_id, payment_id, amount, status, paid, type, credited, created_at FROM Payments WHERE user_id = $1"
	createCountUserPayments = "SELECT COUNT(id) FROM Payments WHERE user_id = $1"

	//updateUserRequest     = "UPDATE Users SET "
	//updateUserName        = "name=$"
	//updateUserSurname     = "surname=$"
	//updateUserPass        = "password=$"
	//updateUserEmail       = "email=$"
	//updateUserDescription = "description=$"
	//updateUserAvatar      = "avatar=$"
	// whereStatement = "WHERE id = $1"
	//updateUserReturning   = "RETURNING name, surname, email, avatar, description"
)

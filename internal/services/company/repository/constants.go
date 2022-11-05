package repository

const (
	getCompanyByIDRequest = "SELECT id, name, description, legal_name, itn, psrn, address,legal_address,email, phone, link, activity, owner_id, rating , verified FROM Companies WHERE id = $1"
)

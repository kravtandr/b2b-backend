package repository

const (
	getCompanyByIDRequest                   = "SELECT id, name, description, legal_name, itn, psrn, address,legal_address,email, phone, link, activity, owner_id, rating , verified, photo_obj_name FROM Companies WHERE id = $1"
	createGetCompanyByOwnerIdAndItn         = "SELECT id, name, description, legal_name, itn, psrn, address,legal_address,email, phone, link, activity, owner_id, rating , verified, photo_obj_name FROM Companies WHERE owner_id = $1 and itn = $2"
	createUpdateCompanyById                 = "UPDATE Companies SET name = $2, description = $3, address = $4, legal_address= $5, phone = $6, link = $7, activity = $8, photo_obj_name = $9 WHERE id = $1 RETURNING id, name, description, legal_name, itn, psrn, address,legal_address,email, phone, link, activity, owner_id, rating , verified, photo_obj_name "
	createUpdateCompanyUsersLink            = "UPDATE CompaniesUsers SET post = $3 WHERE company_id = $1 and user_id = $2 RETURNING id, post, company_id, user_id, itn"
	createGetCompanyUserLinkByOwnerIdAndItn = "SELECT  id, post, company_id, user_id, itn FROM CompaniesUsers WHERE user_id = $1 and itn = $2"
	createGetProductsCompaniesLink          = "SELECT  id, company_id, product_id, amount FROM CompaniesProducts WHERE product_id = $1"
	createGetCompanyPhotos                  = "SELECT photo_obj_name FROM Companies WHERE id = $1"
)

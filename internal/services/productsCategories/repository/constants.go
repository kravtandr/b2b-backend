package repository

const (
	createGetCategoryById           = "SELECT id, name, description From categories WHERE id = $1"
	createGetProductById            = "SELECT id, name, description,price FROM products WHERE id = $1"
	createGetAllCategories          = "SELECT id, name From categories"
	createSearchCategories          = "SELECT id, name, description FROM categories WHERE name ~* $1 OFFSET $2 LIMIT $3"
	createGetProductsList           = "SELECT id, name, description,price FROM products OFFSET $1 LIMIT $2"
	createSearchProducts            = "SELECT id, name, description,price FROM products WHERE name ~* $1 OFFSET $2 LIMIT $3"
	createAddProduct                = "INSERT INTO Products (name, description, price) VALUES ($1, $2, $3) RETURNING id, name, description, price"
	createAddProductsCategoriesLink = "INSERT INTO ProductsCategories (product_id, category_id) VALUES ($1, $2)"
	createAddCompaniesProductsLink  = "INSERT INTO CompaniesProducts (company_id, product_id, addedBy, amount, pay_way, delivery_way, adress) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	createAddProductDocuments       = "INSERT INTO ProductDocuments (product_id, document_obj_name) VALUES ($1, $2)"
	createAddProductPhotos          = "INSERT INTO ProductPhotos (product_id, photo_obj_name) VALUES ($1, $2)"
	createGetProductPhotos          = "SELECT photo_obj_name FROM ProductPhotos WHERE product_id = $1"
	createGetProductDocuments       = "SELECT document_obj_name FROM ProductDocuments WHERE product_id = $1"
)

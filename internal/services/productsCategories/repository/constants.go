package repository

const (
	createAddProduct          = "INSERT INTO Products (name, description, price) VALUES ($1, $2, $3) RETURNING id, name, description, price"
	createAddProductDocuments = "INSERT INTO ProductDocuments (product_id, document_obj_name) VALUES ($1, $2)"
	createAddProductPhotos    = "INSERT INTO ProductPhotos (product_id, photo_obj_name) VALUES ($1, $2)"
	createGetProductById      = "SELECT id, name, description,price FROM products WHERE id = $1"
	createGetProductPhotos    = "SELECT photo_obj_name FROM ProductPhotos WHERE product_id = $1"
	createGetProductDocuments = "SELECT document_obj_name FROM ProductDocuments WHERE product_id = $1"
	createUpdateProduct       = "UPDATE Products SET name = $2, description = $3, price = $4 WHERE id = $1"

	createGetProductsList          = "SELECT id, name, description,price FROM products OFFSET $1 LIMIT $2"
	createSearchProducts           = "SELECT id, name, description,price FROM products WHERE name ~* $1 OFFSET $2 LIMIT $3"
	createGetCompanyProducts       = "SELECT id, name, description,price FROM products as pr JOIN (SELECT company_id, product_id FROM companiesproducts) as cppr on cppr.product_id = pr.id WHERE cppr.company_id = $1 OFFSET $2 LIMIT $3"
	createGetProductsListByFilters = "SELECT pr.id, pr.name, description,price, cat.id, cat.name FROM products as pr JOIN (SELECT product_id, category_id FROM productscategories) as prcat ON pr.id = prcat.product_id JOIN (SELECT categories.id, categories.name From categories) as cat ON  prcat.category_id = cat.id WHERE pr.name ~* $1 AND (cat.name  ~* $2 AND cat.id = Any ($3)) AND price BETWEEN $4 AND $5 OFFSET $6 LIMIT $7"

	createGetCategoryById = "SELECT id, name, description From categories WHERE id = $1"

	createGetAllCategories = "SELECT id, name From categories"
	createSearchCategories = "SELECT id, name, description FROM categories WHERE name ~* $1 OFFSET $2 LIMIT $3"

	createAddProductsCategoriesLink    = "INSERT INTO ProductsCategories (product_id, category_id) VALUES ($1, $2)"
	createGetProductsCategoriesLink    = "SELECT product_id, category_id FROM ProductsCategories WHERE product_id = $1"
	createUpdateProductsCategoriesLink = "UPDATE ProductsCategories SET category_id = $2 WHERE product_id = $1"

	createAddCompaniesProductsLink = "INSERT INTO CompaniesProducts (company_id, product_id, addedBy, amount, pay_way, delivery_way, adress) VALUES ($1, $2, $3, $4, $5, $6, $7)"
)

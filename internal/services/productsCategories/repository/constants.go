package repository

const (
	createGetCategoryById           = "SELECT id, name, description From categories WHERE id = $1"
	createGetProductById            = "SELECT id, name, description,price, photo FROM products WHERE id = $1"
	createGetAllCategories          = "SELECT id, name From categories"
	createSearchCategories          = "SELECT id, name, description FROM categories WHERE name ~ $1 OFFSET $2 LIMIT $3"
	createGetProductsList           = "SELECT id, name, description,price, photo FROM products OFFSET $1 LIMIT $2"
	createSearchProducts            = "SELECT id, name, description,price, photo FROM products WHERE name ~ $1 OFFSET $2 LIMIT $3"
	createAddProduct                = "INSERT INTO Products (name, description, price, photo, docs) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, description, price, photo"
	createAddProductsCategoriesLink = "INSERT INTO ProductsCategories (product_id, category_id) VALUES ($1, $2)"
	createAddCompaniesProductsLink  = "INSERT INTO CompaniesProducts (company_id, product_id, addedBy, amount, pay_way, delivery_way, adress) VALUES ($1, $2, $3, $4, $5, $6, $7)"
)

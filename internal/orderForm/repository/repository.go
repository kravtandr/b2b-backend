package OrderFormRepository

import (
	"b2b/m/pkg/domain"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type OrderFormStorage struct {
	dataHolder *pgxpool.Pool
}

func NewOrderFormStorage(DB *pgxpool.Pool) domain.OrderFormStorage {
	return &OrderFormStorage{dataHolder: DB}
}

func (u *OrderFormStorage) Add(value domain.OrderForm) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		log.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(),
		`INSERT INTO OrderForm ("role", "product_category", "product_name", "order_text", "order_comments", "fio", "email", "phone", "company_name", "itn") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,

		value.Role,
		value.Product_category,
		value.Product_name,
		value.Order_text,
		value.Order_comments,
		value.Fio,
		value.Email,
		value.Phone,
		value.Company_name,
		value.Itn,
	)
	return err
}

// func (u *UserStorage) GetByEmail(key string) (value domain.User, err error) {
// 	var user domain.User

// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		log.Printf("Error while getting user")
// 		return user, err
// 	}
// 	defer conn.Release()

// 	err = conn.QueryRow(context.Background(),
// 		`SELECT id, "name", "surname", "patronymic", "email","password", "group_id"
// 		FROM Users WHERE email = $1`,
// 		key,
// 	).Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Email, &user.Password, &user.GroupId)
// 	return user, err
// }

// func (u *UserStorage) GetPublicUserByEmail(key string) (value domain.PublicUser, err error) {
// 	var user domain.PublicUser

// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		log.Printf("Error while getting user")
// 		return user, err
// 	}
// 	defer conn.Release()

// 	err = conn.QueryRow(context.Background(),
// 		`SELECT name, surname, patronymic,  email
// 		FROM Users WHERE email = $1`,
// 		key,
// 	).Scan(&user.Name, &user.Surname, &user.Patronymic, &user.Email)
// 	return user, err
// }

// func (u *UserStorage) GetPublicUserById(id string) (value domain.PublicUser, err error) {
// 	var user domain.PublicUser

// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		log.Printf("Error while getting user")
// 		return user, err
// 	}
// 	defer conn.Release()

// 	err = conn.QueryRow(context.Background(),
// 		`SELECT name, surname, patronymic,  email
// 		FROM Users WHERE id = $1`,
// 		id,
// 	).Scan(&user.Name, &user.Surname, &user.Patronymic, &user.Email)
// 	return user, err
// }

// func (u *UserStorage) SearchCompanies(key string) (value domain.Companies, err error) {
// 	companies := make([]domain.Company, 0)
// 	//param := "^" + key
// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		log.Printf("Connection error while adding trip ", err)
// 		return companies, err
// 	}
// 	defer conn.Release()
// 	// rows, err := conn.Query(context.Background(),
// 	// 	`SELECT cp.id, cp.email, cp.name, cp.legal_name, cp.itn, cp.psrn, cp.adress, cp.phone, cp.link, cp.category_id
// 	// FROM Company AS cp
// 	// WHERE cp.name ~ $1`,
// 	// 	param)
// 	// if err != nil {
// 	// 	log.Printf("Error while getting companies")
// 	// 	return companies, err
// 	// }
// 	// var cp domain.Company
// 	// for rows.Next() {
// 	// 	rows.Scan(&cp.Id, &cp.Email, &cp.Name, &cp.LegalName, &cp.Itn, &cp.Psrn, &cp.Adress, &cp.Phone, &cp.Link, &cp.CategoryId)
// 	// 	companies = append(companies, cp)
// 	// }
// 	// if rows.Err() != nil {
// 	// 	log.Printf("Error while scanning places", rows.Err())
// 	// 	return companies, err
// 	// }
// 	// if len(companies) == 0 {
// 	// 	return companies, err
// 	// }
// 	return companies, err
// }

// func (u *CompanyStorage) GetCompanyById(key string) (value domain.Company, err error) {
// 	var company domain.Company

// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		log.Printf("Error while getting user")
// 		return company, err
// 	}
// 	defer conn.Release()
// 	log.Printf("key = ", key)
// 	err = conn.QueryRow(context.Background(),
// 		`SELECT id, email, name, legal_name, itn, psrn, adress, phone, link, category_id
// 		FROM Company WHERE id = $1`,
// 		key,
// 	).Scan(&company.Id, &company.Email, &company.Name, &company.LegalName, &company.Itn, &company.Psrn, &company.Adress, &company.Phone, &company.Link, &company.CategoryId)

// 	return company, err
// }

// // func (u *CompanyStorage) GetCompaniesByCategoryTitle(key string) (value domain.Companies, err error) {
// // 	var companies domain.Companies

// // 	conn, err := u.dataHolder.Acquire(context.Background())
// // 	if err != nil {
// // 		log.Printf("Error while getting companies")
// // 		return companies, err
// // 	}
// // 	defer conn.Release()
// // 	err = conn.QueryRow(context.Background(),
// // 		`SELECT cp.id, cp.email, cp.name, cp.legal_name, cp.itn, cp.psrn, cp.adress, cp.phone, cp.link, cp.category_id
// // 		FROM Company AS cp
// // 		JOIN Category AS ct ON cp.category_id = ct.id
// // 		WHERE cp.category_id = $1`,
// // 		key,
// // 	).Scan(&category.Id, &category.Title, &category.IndustryId)

// // 	return category, err
// // }

// func (u *CompanyStorage) GetCompaniesByCategoryId(key string) (value domain.Companies, err error) {
// 	companies := make([]domain.Company, 0)
// 	var category domain.Category

// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		log.Printf("Connection error while adding trip ", err)
// 		return companies, err
// 	}
// 	defer conn.Release()

// 	err = conn.QueryRow(context.Background(),
// 		`SELECT id, title, industry_id
// 		FROM Category WHERE id = $1`,
// 		key,
// 	).Scan(&category.Id, &category.Title, &category.IndustryId)
// 	if err != nil {
// 		log.Printf("Error while getting category")
// 		return companies, err
// 	}

// 	rows, err := conn.Query(context.Background(),
// 		`SELECT cp.id, cp.email, cp.name, cp.legal_name, cp.itn, cp.psrn, cp.adress, cp.phone, cp.link, cp.category_id
// 	FROM Company AS cp
// 	JOIN Category AS ct ON cp.category_id = ct.id
// 	WHERE cp.category_id = $1`,
// 		category.Id)
// 	if err != nil {
// 		log.Printf("Error while getting companies")
// 		return companies, err
// 	}
// 	var cp domain.Company
// 	for rows.Next() {
// 		rows.Scan(&cp.Id, &cp.Email, &cp.Name, &cp.LegalName, &cp.Itn, &cp.Psrn, &cp.Adress, &cp.Phone, &cp.Link, &cp.CategoryId)
// 		companies = append(companies, cp)
// 	}
// 	if rows.Err() != nil {
// 		log.Printf("Error while scanning places", rows.Err())
// 		return companies, err
// 	}
// 	if len(companies) == 0 {
// 		return companies, err
// 	}
// 	return companies, err
// }

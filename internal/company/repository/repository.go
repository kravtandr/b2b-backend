package companyRepository

import (
	"b2b/m/pkg/domain"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/webdeskltd/dadata.v2"
)

type CompanyStorage struct {
	dataHolder *pgxpool.Pool
	daData     *dadata.DaData
}

func NewCompanyStorage(DB *pgxpool.Pool, daData *dadata.DaData) domain.CompanyStorage {
	return &CompanyStorage{dataHolder: DB, daData: daData}
}

func (u *CompanyStorage) Add(value domain.Company) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		log.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`INSERT INTO Companies ("name", "description", "legal_name", "itn", "psrn", "address","legal_address","email", "phone", "link", "activity", "owner_id", "rating"  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,

		value.Name,
		value.Description,
		value.LegalName,
		value.Itn,
		value.Psrn,
		value.Address,
		value.LegalAddress,
		value.Email,
		value.Phone,
		value.Link,
		value.Activity,
		value.OwnerId,
		value.Rating,
	)
	return err
}

func (u *CompanyStorage) AddBaseCompany(value domain.Company, post string) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		log.Printf("Connection error while AddBaseCompany ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`INSERT INTO Companies ("name", "legal_name", "itn", "email", "owner_id") VALUES ($1, $2, $3, $4, $5)`,

		value.Name,
		value.LegalName,
		value.Itn,
		value.Email,
		value.OwnerId,
	)
	if err != nil {
		log.Printf("Connection error while AddBaseCompany ", err)
		return err
	}
	err = u.CompaniesUsersLink(value, post)
	if err != nil {
		log.Printf("Connection error while CompaniesUsersLink ", err)
		return err
	}
	return err
}

func (u *CompanyStorage) CompaniesUsersLink(value domain.Company, post string) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		log.Printf("Connection error while AddBaseCompany ", err)
		return err
	}
	defer conn.Release()
	value, err = u.GetByEmail(value.Email)
	if err != nil {
		log.Printf("Connection error while GetByEmail ", err)
		return err
	}
	_, err = conn.Exec(context.Background(),
		`INSERT INTO companiesusers ("post", "company_id", "user_id") VALUES ($1, $2, $3)`,

		post,
		value.Id,
		value.OwnerId,
	)
	if err != nil {
		log.Printf("Connection error while adding user ", err)
		return err
	}
	return err

}

func (u *CompanyStorage) GetByEmail(key string) (value domain.Company, err error) {
	var company domain.Company

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		log.Printf("Error while getting user")
		return company, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT id, name, description, legal_name, itn, psrn, address, legal_address, email, phone, link, activity, owner_id, rating, verified
		FROM Companies WHERE email = $1`,
		key,
	).Scan(&company.Id, &company.Name, &company.Description, &company.LegalName, &company.Itn, &company.Psrn, &company.Address, &company.LegalAddress, &company.Email, &company.Phone, &company.Link, &company.Activity, &company.OwnerId, &company.Rating, &company.Verified)

	return company, err
}

func (u *CompanyStorage) GetCompanyById(key string) (value domain.Company, err error) {
	var company domain.Company

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		log.Printf("Error while getting user")
		return company, err
	}
	defer conn.Release()
	log.Printf("key = ", key)
	err = conn.QueryRow(context.Background(),
		`SELECT id, name, description, legal_name, itn, psrn, address,legal_address,email, phone, link, activity, owner_id, rating , verified 
	FROM Companies WHERE id = $1`,
		key,
	).Scan(&company.Id, &company.Name, &company.Description, &company.LegalName, &company.Itn, &company.Psrn, &company.Address, &company.LegalAddress, &company.Email, &company.Phone, &company.Link, &company.Activity, &company.OwnerId, &company.Rating, &company.Verified)

	return company, err
}

func (u *CompanyStorage) GetCompanyEmployees(key string) (value domain.Employees, err error) {
	var employees domain.Employees

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		log.Printf("Error while getting user")
		return employees, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(),
		`SELECT   cu.post, cu.user_id, u.name, u.surname, u.patronymic, u.email, u.country, u.group_id
	FROM companiesusers AS cu  JOIN Users u on u.id = cu.user_id WHERE cu.company_id = $1`,
		key)
	if err != nil {
		log.Printf("Error while getting employees")
		return employees, err
	}
	var employee domain.Employee
	for rows.Next() {
		rows.Scan(&employee.Post, &employee.Id, &employee.Name, &employee.Surname, &employee.Patronymic, &employee.Email, &employee.Country, &employee.Group_id)
		employees = append(employees, employee)
	}
	if rows.Err() != nil || len(employees) == 0 {
		log.Printf("Error while scanning employees", rows.Err())
		return employees, errors.New("no such company")
	}
	if len(employees) == 0 {
		log.Printf("Error no employees", rows.Err())
		return employees, errors.New("no employees")
	}

	return employees, err
}

func (u *CompanyStorage) GetCompanyFullInfo(key string) (value domain.CompanyFullInfo, err error) {
	var companyFullInfo domain.CompanyFullInfo
	companyFullInfo.Company, err = u.GetCompanyById(key)
	if err != nil {
		log.Printf("Error while getting GetCompanyFullInfo")
		return companyFullInfo, err
	}
	companyFullInfo.Employees, err = u.GetCompanyEmployees(key)
	if err != nil {
		log.Printf("Error while getting GetCompanyFullInfo")
		return companyFullInfo, err
	}
	return companyFullInfo, err
}

func (u *CompanyStorage) GetCompanyByItnDaData(key string) (value []dadata.ResponseParty, err error) {
	var companyFullInfo []dadata.ResponseParty
	companyFullInfo, err = u.daData.SuggestParties(dadata.SuggestRequestParams{Query: "\"" + fmt.Sprint(key) + "\""})
	if err != nil {
		log.Printf("Error while getting GetCompanyByItnDaData")
		return companyFullInfo, err
	}
	return companyFullInfo, err
}

// func (u *CompanyStorage) SearchCompanies(key string) (value domain.Companies, err error) {
// 	companies := make([]domain.Company, 0)
// 	param := "^" + key
// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		log.Printf("Connection error while adding trip ", err)
// 		return companies, err
// 	}
// 	defer conn.Release()
// 	rows, err := conn.Query(context.Background(),
// 		`SELECT cp.id, cp.email, cp.name, cp.legal_name, cp.itn, cp.psrn, cp.adress, cp.phone, cp.link, cp.category_id
// 	FROM Company AS cp
// 	WHERE cp.name ~ $1`,
// 		param)
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

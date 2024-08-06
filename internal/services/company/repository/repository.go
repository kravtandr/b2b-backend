package repository

import (
	"b2b/m/internal/services/company/models"
	"b2b/m/pkg/errors"
	"b2b/m/pkg/hasher"
	"b2b/m/pkg/helpers"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	minio "github.com/minio/minio-go/v7"
)

type CompanyRepository interface {
	GetCompanyById(ctx context.Context, ID int64) (*models.Company, error)
	GetCompanyByOwnerIdAndItn(ctx context.Context, company models.Company) (*models.Company, error)
	UpdateCompanyById(ctx context.Context, newCompany models.Company) (*models.Company, error)
	UpdateCompanyUsersLink(ctx context.Context, companyId int64, userId int64, post string) (string, error)
	GetCompanyUserLinkByOwnerIdAndItn(ctx context.Context, id int64, itn string) (*models.CompaniesUsersLink, error)
	GetCompanyByProductId(ctx context.Context, ID int64) (*models.Company, error)
	GetProductsCompaniesLink(ctx context.Context, productId int64) (*models.ProductsCompaniesLink, error)
}

type companyRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
	minioClient  *minio.Client
}

func (a *companyRepository) UpdateCompanyById(ctx context.Context, newCompany models.Company) (*models.Company, error) {
	query := a.queryFactory.CreateUpdateCompanyById(newCompany)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	repoCompany := &models.Company{}
	if err := row.Scan(
		&repoCompany.Id, &repoCompany.Name, &repoCompany.Description, &repoCompany.LegalName,
		&repoCompany.Itn, &repoCompany.Psrn, &repoCompany.Address, &repoCompany.LegalAddress,
		&repoCompany.Email, &repoCompany.Phone, &repoCompany.Link, &repoCompany.Activity,
		&repoCompany.OwnerId, &repoCompany.Rating, &repoCompany.Verified, &repoCompany.Photo,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return repoCompany, nil
}

func (a *companyRepository) GetCompanyByOwnerIdAndItn(ctx context.Context, company models.Company) (*models.Company, error) {
	query := a.queryFactory.GetCompanyByOwnerIdAndItn(company.OwnerId, company.Itn)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	repoCompany := &models.Company{}
	if err := row.Scan(
		&company.Id, &company.Name, &company.Description, &company.LegalName,
		&company.Itn, &company.Psrn, &company.Address, &company.LegalAddress,
		&company.Email, &company.Phone, &company.Link, &company.Activity,
		&company.OwnerId, &company.Rating, &company.Verified, &company.Photo,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return repoCompany, nil
}

func (a *companyRepository) GetCompanyUserLinkByOwnerIdAndItn(ctx context.Context, id int64, itn string) (*models.CompaniesUsersLink, error) {
	query := a.queryFactory.CreateGetCompanyUserLinkByOwnerIdAndItn(id, itn)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	companiesUsersLink := &models.CompaniesUsersLink{}
	if err := row.Scan(
		&companiesUsersLink.Id, &companiesUsersLink.Post, &companiesUsersLink.CompanyId,
		&companiesUsersLink.UserId, &companiesUsersLink.Itn,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.CompanyUsersLinkNotExist
		}
		return nil, err
	}
	return companiesUsersLink, nil
}
func (a *companyRepository) UpdateCompanyUsersLink(ctx context.Context, companyId int64, userId int64, post string) (string, error) {
	query := a.queryFactory.CreateUpdateCompanyUsersLink(companyId, userId, post)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	companiesUsersLink := &models.CompaniesUsersLink{}
	if err := row.Scan(
		&companiesUsersLink.Id, &companiesUsersLink.Post, &companiesUsersLink.CompanyId, &companiesUsersLink.UserId, &companiesUsersLink.Itn,
	); err != nil {
		if err == pgx.ErrNoRows {
			return "no rows", errors.CompanyUsersLinkNotExist
		}

		return fmt.Sprint(err), err
	}
	return companiesUsersLink.Post, nil
}

func (a *companyRepository) GetCompanyById(ctx context.Context, ID int64) (*models.Company, error) {
	query := a.queryFactory.CreateGetCompanyByID(ID)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	company := &models.Company{}
	if err := row.Scan(
		&company.Id, &company.Name, &company.Description, &company.LegalName,
		&company.Itn, &company.Psrn, &company.Address, &company.LegalAddress,
		&company.Email, &company.Phone, &company.Link, &company.Activity,
		&company.OwnerId, &company.Rating, &company.Verified, &company.Photo,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.CompanyDoesNotExist
		}

		return nil, err
	}

	return company, nil
}

func (a *companyRepository) GetProductsCompaniesLink(ctx context.Context, productId int64) (*models.ProductsCompaniesLink, error) {
	query := a.queryFactory.CreateGetProductsCompaniesLink(productId)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	productsCompaniesLink := &models.ProductsCompaniesLink{}
	if err := row.Scan(
		&productsCompaniesLink.Id, &productsCompaniesLink.CompanyId, &productsCompaniesLink.ProductId, &productsCompaniesLink.Amount,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.CompanyUsersLinkNotExist
		}

		return nil, err
	}
	return productsCompaniesLink, nil
}

func (a *companyRepository) GetCompanyByProductId(ctx context.Context, ID int64) (*models.Company, error) {
	productsCompaniesLink, err := a.GetProductsCompaniesLink(ctx, ID)
	if err != nil {
		return nil, err
	}
	company, err := a.GetCompanyById(ctx, productsCompaniesLink.CompanyId)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (a companyRepository) PutPhotos(ctx context.Context, Company *models.Company) (*models.Company, error) {

	dec, contentType, err := helpers.DecodeImgFromBase64(ctx, Company.Photo)
	if err != nil {
		log.Println("Error while DecodeImgFromBase64", err)
		return &models.Company{}, errors.CantDecodeImgFromBase64
	}
	opts := minio.PutObjectOptions{ContentType: contentType}

	// create tmp file to avoid creating file with eq filename
	f, err := os.CreateTemp("", hasher.SimpleHash(Company.Name))
	if err != nil {
		log.Println("Error while create tmp file", err)
		return &models.Company{}, errors.CantCreateTmpFile
	}
	if dec != nil {
		if _, err := f.Write(dec); err != nil {
			log.Println("Error while write tmp file", err)
			return &models.Company{}, errors.CantWriteTmpFile
		}
	} else {
		log.Println("No bytes to write file", err)
		return &models.Company{}, errors.NoBytesToWrite
	}

	// TODO в константы
	BucketName := "photo"

	_, err = a.minioClient.FPutObject(
		ctx,
		BucketName, // константа с именем бакета
		f.Name(),   // имя
		f.Name(),   // путь откуда сохранять в минио
		opts,       // тип контента
	)
	if err != nil {
		log.Println("Error in FPutObject ", err)
		return &models.Company{}, errors.ErrorMinioFPutObject
	}
	//fmt.Println(info)
	Company.Photo = f.Name()
	if err := f.Close(); err != nil {
		log.Println(err)
		return &models.Company{}, errors.ErrorMinioFPutObject
	}
	//delite tmp file
	if err := os.Remove(f.Name()); err != nil {
		log.Println(err)
		return &models.Company{}, errors.ErrorMinioFPutObject
	}

	return Company, nil
}

func (a companyRepository) GetCompanyPhotos(ctx context.Context, company *models.Company) (*models.Company, error) {
	query := a.queryFactory.CreateGetCompanyPhotos(company.Id)
	var objName string
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &models.Company{}, err
	}
	// Обнуляем company.Photo перед циклом
	company.Photo = ""
	defer rows.Close()
	for rows.Next() {
		func() {
			err = rows.Scan(&objName)
			var BucketName string = "photo"
			reader, err := a.minioClient.GetObject(
				ctx,
				BucketName,               // константа с именем бакета
				objName,                  // имя
				minio.GetObjectOptions{}, // тип контента
			)
			if err != nil {
				log.Println("Cant`t get image from image-storage:", err)
			}
			defer reader.Close()
			// log.Println("reader  = ", reader)
			// log.Println("objName  = ", objName)
			if err != nil {
				log.Println("Error in a.minioClient.GetObject ", err)
			}
			imageBytes := make([]byte, 1)
			imageBytes, err = io.ReadAll(reader)
			if err != nil && err != io.EOF {
				log.Println("Error in io.ReadAll(image) ", err)
			}
			data := imageBytes
			base64photo := helpers.EncodeImgToBase64(ctx, data)
			//log.Println("imageBytes  = ", data[10:40])
			//log.Println("ImgToBase64  = ", base64photo[10:40])
			// Product.Photo = append(Product.Photo, base64photo)
			company.Photo = fmt.Sprint(helpers.CheckSum(base64photo))

		}()
	}
	if rows.Err() != nil {
		return &models.Company{}, err
	}
	return company, err
}

func NewCompanyRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) CompanyRepository {
	return &companyRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}

package repository

import (
	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/errors"
	helpers "b2b/m/pkg/helpers"
	"context"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	minio "github.com/minio/minio-go/v7"
)

type ProductsCategoriesRepository interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error)
	SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error)
	GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error)
	SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.ProductsList, error)
	AddProduct(ctx context.Context, Product *models.Product) (*models.Product, error)
	AddProductsCategoriesLink(ctx context.Context, productId int64, categoryId int64) error
	AddCompaniesProductsLink(ctx context.Context, CompaniesProducts *models.CompaniesProducts) error
}

type productsCategoriesRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
	minioClient  *minio.Client
}

func (a productsCategoriesRepository) GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error) {
	query := a.queryFactory.CreateGetCategoryById(CategoryId.Id)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	category := &models.Category{}
	if err := row.Scan(
		&category.Id, &category.Name, &category.Description,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Category{}, errors.CategoryDoesNotExist
		}
		return &models.Category{}, err
	}
	return category, nil
}

func (a productsCategoriesRepository) GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error) {
	query := a.queryFactory.CreateGetProductById(ProductId.Id)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	product := &models.Product{}
	if err := row.Scan(
		&product.Id, &product.Name, &product.Description, &product.Price,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Product{}, errors.ProductDoesNotExist
		}
		return &models.Product{}, err
	}
	product, err := a.GetProductWithPhotosAndDocuments(ctx, product)
	if err != nil {
		return &models.Product{}, err
	}
	return product, nil
}

func (a productsCategoriesRepository) GetProductWithPhotosAndDocuments(ctx context.Context, product *models.Product) (*models.Product, error) {
	productWithPhotos, err := a.GetProductPhotos(ctx, product)
	if err != nil {
		return &models.Product{}, err
	}
	log.Println("productWithPhotos ", productWithPhotos)
	//TODO add documents
	// productWithPhotosAndDocuments, err := a.GetProductDocuments(ctx, productWithPhotos)
	// if err != nil {
	// 	return &models.Product{}, err
	// }
	// log.Println("productWithPhotosAndDocuments ", productWithPhotosAndDocuments)
	return productWithPhotos, nil
}

func (a productsCategoriesRepository) PutPhotos(ctx context.Context, Product *models.Product) (*models.Product, error) {

	for i, photo := range Product.Photo {

		dec, contentType, err := helpers.DecodeImgFromBase64(ctx, photo)
		if err != nil {
			log.Println("Error while DecodeImgFromBase64", err)
			return &models.Product{}, errors.CantDecodeImgFromBase64
		}
		opts := minio.PutObjectOptions{ContentType: contentType}

		// create tmp file to avoid creating file with eq filename
		f, err := os.CreateTemp("", Product.Name)
		if err != nil {
			log.Println("Error while create tmp file", err)
			return &models.Product{}, errors.CantCreateTmpFile
		}
		if dec != nil {
			if _, err := f.Write(dec); err != nil {
				log.Println("Error while write tmp file", err)
				return &models.Product{}, errors.CantWriteTmpFile
			}
		} else {
			log.Println("No bytes to write file", err)
			return &models.Product{}, errors.NoBytesToWrite
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
			return &models.Product{}, errors.ErrorMinioFPutObject
		}
		//fmt.Println(info)
		Product.Photo[i] = f.Name()
		defer f.Close()
		//delite tmp file
		os.Remove(f.Name())
		if err != nil {
			log.Println(err)
		}

	}
	return Product, nil
}

func (a productsCategoriesRepository) AddProductDocuments(ctx context.Context, Product *models.Product) error {
	for _, document := range Product.Docs {
		query := a.queryFactory.CreateAddProductDocuments(Product.Id, document)
		_ = a.conn.QueryRow(ctx, query.Request, query.Params...)
	}
	return nil
}

func (a productsCategoriesRepository) AddProductPhotos(ctx context.Context, Product *models.Product) error {
	for _, photo := range Product.Photo {
		query := a.queryFactory.CreateAddProductPhotos(Product.Id, photo)
		_ = a.conn.QueryRow(ctx, query.Request, query.Params...)
	}
	return nil
}

func (a productsCategoriesRepository) AddProduct(ctx context.Context, Product *models.Product) (*models.Product, error) {
	// store base64 undecoded or decode and store image?????????
	// Very roughly, the final size of Base64-encoded binary data is equal to 1.37 times the original data size
	// if decode and store as image had to convert to right image type png/jpeg and encode to response
	// how to store decoded image??? can it be stored in var???
	// decoding base64 easier way to store in minio and have previev
	// if store base64 how to put base64 to minio??? what size? on the other side no decoding
	// decode to img -> minio store -> encode to base64

	product := &models.Product{}
	product = Product
	product, err := a.PutPhotos(ctx, Product)
	if err != nil {
		log.Println("Error in PutPhotos ", err)
		return &models.Product{}, err
	}
	query := a.queryFactory.CreateAddProduct(product)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	if err := row.Scan(
		&product.Id, &product.Name, &product.Description, &product.Price,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Product{}, errors.ProductDoesNotExist
		}
		return &models.Product{}, err
	}
	err = a.AddProductPhotos(ctx, product)
	if err != nil {
		log.Println("Error in AddProductPhotos ", err)
		return &models.Product{}, err
	}
	err = a.AddProductDocuments(ctx, product)
	if err != nil {
		log.Println("Error in AddProductDocuments ", err)
		return &models.Product{}, err
	}
	//base64 in response
	result, err := a.GetProductById(ctx, &models.ProductId{Id: product.Id})
	if err != nil {
		log.Println("Error in GetProductById ", err)
		return &models.Product{}, err
	}
	return result, nil
}

func (a productsCategoriesRepository) AddProductsCategoriesLink(ctx context.Context, productId int64, categoryId int64) error {
	query := a.queryFactory.CreateAddProductsCategoriesLink(productId, categoryId)
	_ = a.conn.QueryRow(ctx, query.Request, query.Params...)
	return nil
}

func (a productsCategoriesRepository) AddCompaniesProductsLink(ctx context.Context, CompaniesProducts *models.CompaniesProducts) error {
	query := a.queryFactory.CreateAddCompaniesProductsLink(CompaniesProducts)
	_ = a.conn.QueryRow(ctx, query.Request, query.Params...)
	return nil
}

func (a productsCategoriesRepository) SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error) {
	query := a.queryFactory.CreateSearchCategories(SearchBody)
	var category models.Category
	var categories []models.Category
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &categories, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Name, &category.Description)
		categories = append(categories, category)
	}
	if rows.Err() != nil {
		return &categories, err
	}
	return &categories, err
}

func (a productsCategoriesRepository) GetProductPhotos(ctx context.Context, Product *models.Product) (*models.Product, error) {
	query := a.queryFactory.CreateGetProductPhotos(Product.Id)
	var objName string
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &models.Product{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&objName)
		//filename := strings.Split(objName, "_")
		//log.Println("GET FROM MINO objName", objName, "........")
		log.Println("GET FROM MINO objName", objName, "........")

		BucketName := "photo"
		// image, err := a.minioClient.GetObject(
		// 	ctx,
		// 	BucketName,               // константа с именем бакета
		// 	objName,                  // имя
		// 	minio.GetObjectOptions{}, // тип контента
		// )

		reader, err := a.minioClient.GetObject(
			ctx,
			BucketName,
			objName,
			minio.GetObjectOptions{},
		)
		if err != nil {
			log.Println("Cant`t get image from image-storage:", err)
			return &models.Product{}, err
		}
		defer reader.Close()
		if err != nil {
			log.Println("Error in a.minioClient.GetObject ", err)
			return &models.Product{}, err
		}
		log.Println("image from minio ", reader)
		//base64.StdEncoding.Encode(image)
		imageBytes, err := io.ReadAll(reader)
		if err != nil {
			log.Println("Error in io.ReadAll(image) ", err)
			return &models.Product{}, err
		}
		log.Println("imageBytes  = ", imageBytes)
		Product.Photo = append(Product.Photo, helpers.EncodeImgToBase64(ctx, imageBytes))
	}
	if rows.Err() != nil {
		return &models.Product{}, err
	}
	return Product, err
}

func (a productsCategoriesRepository) GetProductDocuments(ctx context.Context, Product *models.Product) (*models.Product, error) {
	query := a.queryFactory.CreateGetProductDocuments(Product.Id)
	var objName string
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &models.Product{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&objName)
		BucketName := "photo"
		docs, err := a.minioClient.GetObject(
			ctx,
			BucketName,               // константа с именем бакета
			objName,                  // имя
			minio.GetObjectOptions{}, // тип контента
		)
		if err != nil {
			log.Println("Error in a.minioClient.GetObject ", err)
			return &models.Product{}, err
		}
		docsBytes, err := io.ReadAll(docs)
		if err != nil {
			log.Println("Error in  io.ReadAll(docs) ", err)
			return &models.Product{}, err
		}
		Product.Docs = append(Product.Docs, helpers.EncodeImgToBase64(ctx, docsBytes))
	}
	if rows.Err() != nil {
		return &models.Product{}, err
	}
	return Product, err
}

func (a productsCategoriesRepository) GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error) {
	query := a.queryFactory.CreateGetProductsList(SkipLimit)
	var product models.Product
	var products models.ProductsList
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		log.Println("Error in Query GetProductsList", err)
		return &products, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
		productWithPhotos, err := a.GetProductWithPhotosAndDocuments(ctx, &product)
		if err != nil {
			log.Println("Error in  GetProductsList->GetProductWithPhotosAndDocuments", err)
			return &products, err
		}
		log.Println(productWithPhotos.Photo)
		product.Photo = productWithPhotos.Photo
		products = append(products, product)
	}

	if rows.Err() != nil {
		log.Println("Error in GetProductsList rows scan ", err)
		return &products, err
	}

	return &products, err
}

func (a productsCategoriesRepository) SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.ProductsList, error) {
	query := a.queryFactory.CreateSearchProducts(SearchBody)
	var product models.Product
	var products models.ProductsList
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &products, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
		product, err := a.GetProductWithPhotosAndDocuments(ctx, &product)
		if err != nil {
			return &products, err
		}
		products = append(products, *product)
	}
	if rows.Err() != nil {
		return &products, err
	}
	return &products, err
}

func NewProductsCategoriesRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
	minioClient *minio.Client,
) ProductsCategoriesRepository {
	return &productsCategoriesRepository{
		queryFactory: queryFactory,
		conn:         conn,
		minioClient:  minioClient,
	}
}

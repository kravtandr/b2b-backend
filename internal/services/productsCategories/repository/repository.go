package repository

import (
	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/errors"
	hasher "b2b/m/pkg/hasher"
	helpers "b2b/m/pkg/helpers"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	minio "github.com/minio/minio-go/v7"
)

type ProductsCategoriesRepository interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error)
	SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error)
	GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error)
	GetProductsListByFilters(ctx context.Context, filters *models.ProductsFilters) (*models.ProductsList, error)
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
	//log.Println("productWithPhotos ", productWithPhotos)
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
		f, err := os.CreateTemp("", hasher.SimpleHash(Product.Name))
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
		if err := f.Close(); err != nil {
			log.Println(err)
			return &models.Product{}, errors.ErrorMinioFPutObject
		}
		//delite tmp file
		if err := os.Remove(f.Name()); err != nil {
			log.Println(err)
			return &models.Product{}, errors.ErrorMinioFPutObject
		}

	}
	return Product, nil
}

func (a productsCategoriesRepository) AddProductDocuments(ctx context.Context, Product *models.Product) error {
	for _, document := range Product.Docs {
		query := a.queryFactory.CreateAddProductDocuments(Product.Id, document)
		if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
			log.Println("ERROR AddProductDocuments", err)
			return err
		}

	}
	return nil
}

func (a productsCategoriesRepository) AddProductPhotos(ctx context.Context, Product *models.Product) error {
	for _, photo := range Product.Photo {
		query := a.queryFactory.CreateAddProductPhotos(Product.Id, photo)
		if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
			log.Println("ERROR AddProductPhotos", err)
			return err
		}
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
	log.Println("Start AddProduct")
	log.Println("PutPhotos...")
	Product, err := a.PutPhotos(ctx, Product)
	if err != nil {
		log.Println("Error in PutPhotos ", err)
		return &models.Product{}, err
	}
	log.Println("PutPhotos - OK")
	log.Println("CreateAddProduct...")
	log.Println("a.queryFactory.CreateAddProduct...")
	query := a.queryFactory.CreateAddProduct(Product)
	log.Println("a.queryFactory.CreateAddProduct - OK")
	log.Println("a.conn.QueryRow...")
	//timeout 15 sek
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	log.Println("_______________________________")
	err = a.conn.Ping(ctx)
	if err != nil {
		log.Println("PING ERR", err)
	} else {
		log.Println("PING OK")
	}
	log.Println(ctx, query.Request, query.Params)
	log.Println("_______________________________")
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	log.Println("a.conn.QueryRow - OK")
	if err := row.Scan(
		&Product.Id, &Product.Name, &Product.Description, &Product.Price,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Product{}, errors.ProductDoesNotExist
		}
		return &models.Product{}, err
	}
	log.Println("CreateAddProduct - OK")
	log.Println("AddProductPhotos...")
	err = a.AddProductPhotos(ctx, Product)
	if err != nil {
		log.Println("Error in AddProductPhotos ", err)
		return &models.Product{}, err
	}
	log.Println("AddProductPhotos - OK")
	log.Println("AddProductDocuments...")
	err = a.AddProductDocuments(ctx, Product)
	if err != nil {
		log.Println("Error in AddProductDocuments ", err)
		return &models.Product{}, err
	}
	log.Println("AddProductDocuments - OK")
	//base64 in response
	log.Println("GetProductById...")
	result, err := a.GetProductById(ctx, &models.ProductId{Id: Product.Id})
	if err != nil {
		log.Println("Error in GetProductById ", err)
		return &models.Product{}, err
	}
	log.Println("GetProductById - OK")
	log.Println("END AddProduct")
	return result, nil
	//return &models.Product{}, nil

}

func (a productsCategoriesRepository) AddProductsCategoriesLink(ctx context.Context, productId int64, categoryId int64) error {
	query := a.queryFactory.CreateAddProductsCategoriesLink(productId, categoryId)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		log.Println("ERROR AddProductsCategoriesLink", err)
		return err
	}
	return nil
}

func (a productsCategoriesRepository) AddCompaniesProductsLink(ctx context.Context, CompaniesProducts *models.CompaniesProducts) error {
	query := a.queryFactory.CreateAddCompaniesProductsLink(CompaniesProducts)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		log.Println("ERROR AddCompaniesProductsLink", err)
		return err
	}
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
	// Обнуляем Product.Photo перед циклом
	Product.Photo = nil
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
			Product.Photo = append(Product.Photo, base64photo)
			Product.Photo = append(Product.Photo, fmt.Sprint(helpers.CheckSum(base64photo)))

		}()
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
	var productsWithPhoto models.ProductsList
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		log.Println("Error in Query GetProductsList", err)
		return &products, err
	}
	defer rows.Close()
	for rows.Next() {
		func() {
			product = models.Product{}
			//productWithPhotos := &models.Product{}
			err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
			if err != nil {
				log.Println("Error in  GetProductsList->GetProductWithPhotosAndDocuments", err)
				//return &products, err
			}
			products = append(products, product)
		}()
	}
	for _, item := range products {
		func() {
			productWithPhotos, err := a.GetProductWithPhotosAndDocuments(ctx, &item)
			if err != nil {
				log.Println("Error in  GetProductsList->GetProductWithPhotosAndDocuments", err)
				//return &products, err
			}
			productsWithPhoto = append(productsWithPhoto, *productWithPhotos)
		}()
	}

	if rows.Err() != nil {
		log.Println("Error in GetProductsList rows scan ", err)
		return &products, err
	}

	return &productsWithPhoto, err
}

func (a productsCategoriesRepository) GetProductsListByFilters(ctx context.Context, filters *models.ProductsFilters) (*models.ProductsList, error) {
	query := a.queryFactory.CreateGetProductsListByFilters(filters)
	var product models.Product
	var products models.ProductsList
	var productsWithPhoto models.ProductsList
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		log.Println("Error in Query GetProductsListByFilters", err)
		return &products, err
	}
	defer rows.Close()
	for rows.Next() {
		func() {
			product = models.Product{}
			err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
			if err != nil {
				log.Println("Error in  GetProductsListByFilters->GetProductWithPhotosAndDocuments", err)
			}
			products = append(products, product)
		}()
	}
	for _, item := range products {
		func() {
			productWithPhotos, err := a.GetProductWithPhotosAndDocuments(ctx, &item)
			if err != nil {
				log.Println("Error in  GetProductsListByFilters->GetProductWithPhotosAndDocuments", err)
			}
			productsWithPhoto = append(productsWithPhoto, *productWithPhotos)
		}()
	}

	if rows.Err() != nil {
		log.Println("Error in GetProductsListByFilters rows scan ", err)
		return &products, err
	}

	return &productsWithPhoto, err
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

package repository

import (
	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/errors"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

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
		&product.Id, &product.Name, &product.Description, &product.Price, &product.Photo,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Product{}, errors.ProductDoesNotExist
		}
		return &models.Product{}, err
	}
	return product, nil
}

func (a productsCategoriesRepository) PutPhotos(ctx context.Context, Product *models.Product) (*models.Product, error) {

	for i, photo := range Product.Photo {
		// //upload
		// imageName := "test_png.png"
		// filePath := "./test_png.png"
		var dec []byte
		var opts minio.PutObjectOptions
		coI := strings.Index(string(photo), ",")
		switch strings.TrimSuffix(photo[5:coI], ";base64") {
		case "image/png":
			opts = minio.PutObjectOptions{ContentType: "image/png"}
			dec, err := base64.StdEncoding.DecodeString(photo)
			if err != nil {
				log.Println("Error while decode base64 string", dec)
				return &models.Product{}, err
			}
			// img, err := png.Decode(bytes.NewReader(dec))
			// if err != nil {
			// 	log.Println("Error while png.Decode")
			// 	return &models.Product{}, err
			// }

		case "image/jpeg", "image/jpg":
			opts = minio.PutObjectOptions{ContentType: "image/jpeg"}
			dec, err := base64.StdEncoding.DecodeString(photo)
			if err != nil {
				log.Println("Error while decode base64 string", dec)
				return &models.Product{}, err
			}
			// img, err := jpeg.Decode(bytes.NewReader(dec))
			// if err != nil {
			// 	log.Println("Error while jpeg.Decode")
			// 	return &models.Product{}, err
			// }

		default:
			log.Println("Unsuppotred file type")
			return &models.Product{}, errors.DeniedAccess
		}

		// create tmp file to avoid creating file with eq filename
		f, err := os.CreateTemp("/tmp", Product.Name)
		if err != nil {
			log.Println("Error while create tmp file")
			return &models.Product{}, err
		}
		if dec != nil {
			if _, err := f.Write(dec); err != nil {
				log.Println("Error while write tmp file")
				return &models.Product{}, err
			}
		} else {
			log.Println("No bytes to write file")
			return &models.Product{}, err
		}
		defer f.Close()

		// TODO в константы
		BucketName := "photo"

		info, err := a.minioClient.FPutObject(
			context.Background(),
			BucketName,                     // константа с именем бакета
			Product.Name+f.Name()+"_photo", // имя
			"/tmp/"+f.Name(),               // путь откуда сохранять в минио
			opts,                           // тип контента
		)
		if err != nil {
			log.Println("Error in FPutObject ", err)
		}
		fmt.Println(info)
		Product.Photo[i] = Product.Name + "_" + f.Name() + "_photo"
		//delite tmp file
		os.Remove(f.Name())
		if err != nil {
			log.Println(err)
		}

	}
	return Product, nil
}

func (a productsCategoriesRepository) AddProduct(ctx context.Context, Product *models.Product) (*models.Product, error) {
	// store base64 undecoded or decode and store image?????????
	// Very roughly, the final size of Base64-encoded binary data is equal to 1.37 times the original data size
	// if decode and store as image had to convert to right image type png/jpeg and encode to response
	// how to store decoded image??? can it be stored in var???
	// decoding base64 easier way to store in minio and have previev
	// if store base64 how to put base64 to minio??? what size? on the other side no decoding
	// decode to img -> minio store -> encode to base64
	query := a.queryFactory.CreateAddProduct(Product)

	// log.Printf("Successfully uploaded %s of size %d\n", imageName, info.Size)
	// fmt.Println("MINIP INFO", info)
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
	product, err := a.PutPhotos(ctx, product)
	if err != nil {
		log.Println("Error in PutPhotos ", err)
	}

	return product, nil
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

func (a productsCategoriesRepository) GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error) {
	query := a.queryFactory.CreateGetProductsList(SkipLimit)
	var product models.Product
	var products models.ProductsList
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &products, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Photo)
		products = append(products, product)
	}
	if rows.Err() != nil {
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
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Photo)
		products = append(products, product)
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

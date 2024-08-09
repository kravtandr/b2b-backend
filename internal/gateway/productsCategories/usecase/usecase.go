package usecase

import (
	company_usecase "b2b/m/internal/gateway/company/usecase"
	auth_usecase "b2b/m/internal/gateway/user/usecase"
	"b2b/m/internal/models"
	chttp "b2b/m/pkg/customhttp"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"
	"database/sql"
	"log"
)

type ProductsCategoriesUseCase interface {
	AddProduct(ctx context.Context, request *models.UserInfoAndAddProductByFormRequest) (*models.GetProduct, error)
	GetProductById(ctx context.Context, request *models.GetProductByIdRequest) (*models.GetProductByIdResponse, error)
	UpdateProduct(ctx context.Context, request *models.UserInfoAndUpdateProductByFormRequest) (*models.GetProduct, error)

	GetProductsList(ctx context.Context, request *chttp.QueryParam) (*models.GetProductsList, error)
	GetProductsListByFilters(ctx context.Context, params *chttp.QueryParam, request *models.GetProductsByFilters) (*models.ProductsWithCategoryAndCompany, error)
	SearchProducts(ctx context.Context, request *chttp.SearchItemNameWithSkipLimit) (*models.GetProductsList, error)
	GetCompanyProducts(ctx context.Context, request *models.GetCompanyProductsRequest, params *chttp.QueryParam) (*models.GetProductsList, error)

	GetCategoryById(ctx context.Context, request *models.GetCategoryByIdRequest) (*models.GetCategoryByIdResponse, error)
	SearchCategories(ctx context.Context, request *chttp.SearchItemNameWithSkipLimit) (*[]models.GetCategoryByIdResponse, error)
}

type productsCategoriesUseCase struct {
	ProductsCategoriesGRPC ProductsCategoriesGRPC
	AuthUsecase            auth_usecase.UserUsecase
	CompanyUsecase         company_usecase.CompanyUseCase
}

func (u *productsCategoriesUseCase) GetCategoryById(ctx context.Context, request *models.GetCategoryByIdRequest) (*models.GetCategoryByIdResponse, error) {
	response, err := u.ProductsCategoriesGRPC.GetCategoryById(ctx, &productsCategories_service.GetCategoryByID{
		Id: request.Id,
	})
	if err != nil {
		return &models.GetCategoryByIdResponse{}, err
	}
	description := sql.NullString{
		String: response.Description.String_,
		Valid:  response.Description.Valid}
	return &models.GetCategoryByIdResponse{
		Id:          request.Id,
		Name:        response.Name,
		Description: description,
	}, nil
}

func (u *productsCategoriesUseCase) UpdateProduct(ctx context.Context, request *models.UserInfoAndUpdateProductByFormRequest) (*models.GetProduct, error) {
	userProfile, err := u.AuthUsecase.Profile(ctx, request.UserProfile.Id)
	if err != nil {
		log.Println("Gateway -> Usecase -> UpdateProduct -> u.AuthUsecase.Profile ERROR", err)
		return &models.GetProduct{}, err
	}
	request.UserProfile = *userProfile
	description := &productsCategories_service.SqlNullString{
		String_: request.Product.Description,
		Valid:   true}

	// var photo string
	// var photos []string
	// for _, result := range request.Product.Photo {

	// 	modelCategories = append(modelCategories, modelCategory)
	// }
	response, err := u.ProductsCategoriesGRPC.UpdateProduct(ctx, &productsCategories_service.UpdateProductRequest{
		Id:           request.Product.Id,
		Name:         request.Product.Name,
		CategoryId:   request.Product.CategoryId,
		Description:  description,
		Price:        request.Product.Price,
		Amount:       request.Product.Amount,
		PayWay:       request.Product.PayWay,
		Adress:       request.Product.Adress,
		DeliveryWay:  request.Product.DeliveryWay,
		ProductPhoto: request.Product.Photo,
		Docs:         request.Product.Docs,
		UserId:       request.UserProfile.Id,
		CompanyId:    request.UserProfile.Company.Id,
	})
	if err != nil {
		log.Println("Gateway -> Usecase -> UpdateProduct -> u.ProductsCategoriesGRPC.UpdateProduct ERROR", err)
		return &models.GetProduct{}, err
	}
	respDescription := sql.NullString{
		String: response.Description.String_,
		Valid:  response.Description.Valid}
	return &models.GetProduct{
		Id:          response.Id,
		Name:        response.Name,
		Description: respDescription,
		Price:       response.Price,
		Photo:       response.Photo,
		Docs:        response.Documents,
	}, nil

}

func (u *productsCategoriesUseCase) AddProduct(ctx context.Context, request *models.UserInfoAndAddProductByFormRequest) (*models.GetProduct, error) {
	userProfile, err := u.AuthUsecase.Profile(ctx, request.UserProfile.Id)
	if err != nil {
		log.Println("Gateway -> Usecase -> AddProduct -> u.AuthUsecase.Profile ERROR", err)
		return &models.GetProduct{}, err
	}
	request.UserProfile = *userProfile
	description := &productsCategories_service.SqlNullString{
		String_: request.Product.Description,
		Valid:   true}

	// var photo string
	// var photos []string
	// for _, result := range request.Product.Photo {

	// 	modelCategories = append(modelCategories, modelCategory)
	// }
	response, err := u.ProductsCategoriesGRPC.AddProduct(ctx, &productsCategories_service.AddProductRequest{
		Name:         request.Product.Name,
		CategoryId:   request.Product.CategoryId,
		Description:  description,
		Price:        request.Product.Price,
		Amount:       request.Product.Amount,
		PayWay:       request.Product.PayWay,
		Adress:       request.Product.Adress,
		DeliveryWay:  request.Product.DeliveryWay,
		ProductPhoto: request.Product.Photo,
		Docs:         request.Product.Docs,
		UserId:       request.UserProfile.Id,
		CompanyId:    request.UserProfile.Company.Id,
	})
	if err != nil {
		log.Println("Gateway -> Usecase -> AddProduct -> u.ProductsCategoriesGRPC.AddProduct ERROR", err)
		return &models.GetProduct{}, err
	}
	respDescription := sql.NullString{
		String: response.Description.String_,
		Valid:  response.Description.Valid}
	return &models.GetProduct{
		Id:          response.Id,
		Name:        response.Name,
		Description: respDescription,
		Price:       response.Price,
		Photo:       response.Photo,
		Docs:        response.Documents,
	}, nil
}
func (u *productsCategoriesUseCase) GetProductById(ctx context.Context, request *models.GetProductByIdRequest) (*models.GetProductByIdResponse, error) {
	response, err := u.ProductsCategoriesGRPC.GetProductById(ctx, &productsCategories_service.GetProductByID{
		Id: request.Id,
	})
	if err != nil {
		log.Println("Error GetProductById -> GetProductById", err)
		return &models.GetProductByIdResponse{}, err
	}
	productCompany, err := u.CompanyUsecase.GetCompanyByProductId(ctx, request.Id)
	if err != nil {
		log.Println("Error GetProductById -> GetCompanyByProductId", err)
		return &models.GetProductByIdResponse{}, err
	}
	description := sql.NullString{
		String: response.Description.String_,
		Valid:  response.Description.Valid}
	return &models.GetProductByIdResponse{
		Id:          request.Id,
		Name:        response.Name,
		Description: description,
		Price:       response.Price,
		Photo:       response.Photo,
		Docs:        response.Documents,
		Company: models.Company{
			Id:           productCompany.Id,
			Name:         productCompany.Name,
			Description:  productCompany.Description,
			LegalName:    productCompany.LegalName,
			Itn:          productCompany.Itn,
			Psrn:         productCompany.Psrn,
			Address:      productCompany.Address,
			LegalAddress: productCompany.LegalAddress,
			Email:        productCompany.Email,
			Phone:        productCompany.Phone,
			Link:         productCompany.Link,
			Activity:     productCompany.Activity,
			OwnerId:      productCompany.OwnerId,
			Rating:       productCompany.Rating,
			Verified:     productCompany.Verified,
			Photo:        productCompany.Photo,
		},
	}, nil
}

func (u *productsCategoriesUseCase) SearchCategories(ctx context.Context, request *chttp.SearchItemNameWithSkipLimit) (*[]models.GetCategoryByIdResponse, error) {
	SearchResults, err := u.ProductsCategoriesGRPC.SearchCategories(ctx, &productsCategories_service.SearchItemNameWithSkipLimitRequest{
		Name:  request.Name,
		Skip:  request.Skip,
		Limit: request.Limit,
	})
	if err != nil {
		return nil, err
	}

	var modelCategory models.GetCategoryByIdResponse
	var modelCategories []models.GetCategoryByIdResponse
	for _, result := range SearchResults.Categories {
		description := sql.NullString{
			String: result.Description.String_,
			Valid:  result.Description.Valid}
		modelCategory = models.GetCategoryByIdResponse{
			Id:          result.Id,
			Name:        result.Name,
			Description: description,
		}
		modelCategories = append(modelCategories, modelCategory)
	}
	return &modelCategories, nil
}

func (u *productsCategoriesUseCase) SearchProducts(ctx context.Context, request *chttp.SearchItemNameWithSkipLimit) (*models.GetProductsList, error) {
	SearchResults, err := u.ProductsCategoriesGRPC.SearchProducts(ctx, &productsCategories_service.SearchItemNameWithSkipLimitRequest{
		Name:  request.Name,
		Skip:  request.Skip,
		Limit: request.Limit,
	})
	if err != nil {
		return nil, err
	}

	var modelProduct models.GetProduct
	var modelProducts models.GetProductsList
	for _, result := range SearchResults.Products {
		description := sql.NullString{
			String: result.Description.String_,
			Valid:  result.Description.Valid}
		modelProduct = models.GetProduct{
			Id:          result.Id,
			Name:        result.Name,
			Description: description,
			Price:       result.Price,
			Photo:       result.Photo,
			Docs:        result.Documents,
		}
		modelProducts = append(modelProducts, modelProduct)
	}
	return &modelProducts, nil
}
func (u *productsCategoriesUseCase) GetProductsListByFilters(ctx context.Context, params *chttp.QueryParam, request *models.GetProductsByFilters) (*models.ProductsWithCategoryAndCompany, error) {
	response, err := u.ProductsCategoriesGRPC.GetProductsListByFilters(ctx, &productsCategories_service.GetProductsListByFiltersRequest{
		ProductName:      request.Product_name,
		CategoryName:     request.Category_name,
		CategoriesIds:    request.Categories_ids,
		PriceLowerLimit:  request.Price_lower_limit,
		PriceHigherLimit: request.Price_higher_limit,
		Skip:             params.Skip,
		Limit:            params.Limit,
	})
	if err != nil {
		return &models.ProductsWithCategoryAndCompany{}, err
	}
	var modelProduct models.ProductWithCategoryAndCompany
	var modelProducts models.ProductsWithCategoryAndCompany
	var company models.Company

	for _, product := range response.Products {

		description := sql.NullString{
			String: product.Description.String_,
			Valid:  product.Description.Valid}
		company_response, err := u.CompanyUsecase.GetCompanyByProductId(ctx, product.Id)

		if err != nil {
			return nil, err
		}
		company = models.Company{
			Id:           company_response.Id,
			Name:         company_response.Name,
			Description:  company_response.Description,
			LegalName:    company_response.LegalName,
			Itn:          company_response.Itn,
			Psrn:         company_response.Psrn,
			Address:      company_response.Address,
			LegalAddress: company_response.LegalAddress,
			Email:        company_response.Email,
			Phone:        company_response.Phone,
			Link:         company_response.Link,
			Activity:     company_response.Activity,
			OwnerId:      company_response.OwnerId,
			Rating:       company_response.Rating,
			Verified:     company_response.Verified,
			Photo:        company_response.Photo,
		}
		modelProduct = models.ProductWithCategoryAndCompany{
			Id:           product.Id,
			Name:         product.Name,
			Description:  description,
			Price:        product.Price,
			Photo:        product.Photo,
			Docs:         product.Documents,
			CategoryId:   product.CategoryId,
			CategoryName: product.CategoryName,
			Company:      company,
		}
		modelProducts = append(modelProducts, modelProduct)
	}
	return &modelProducts, nil
}

func (u *productsCategoriesUseCase) GetProductsList(ctx context.Context, request *chttp.QueryParam) (*models.GetProductsList, error) {
	response, err := u.ProductsCategoriesGRPC.GetProductsList(ctx, &productsCategories_service.GetProductsListRequest{
		Skip:  request.Skip,
		Limit: request.Limit,
	})
	if err != nil {
		return &models.GetProductsList{}, err
	}
	var modelProduct models.GetProduct
	var modelProducts models.GetProductsList
	for _, result := range response.Products {
		description := sql.NullString{
			String: result.Description.String_,
			Valid:  result.Description.Valid}
		modelProduct = models.GetProduct{
			Id:          result.Id,
			Name:        result.Name,
			Description: description,
			Price:       result.Price,
			Photo:       result.Photo,
			Docs:        result.Documents,
		}
		modelProducts = append(modelProducts, modelProduct)
	}
	return &modelProducts, nil
}

func (u *productsCategoriesUseCase) GetCompanyProducts(ctx context.Context, request *models.GetCompanyProductsRequest, params *chttp.QueryParam) (*models.GetProductsList, error) {
	response, err := u.ProductsCategoriesGRPC.GetCompanyProducts(ctx, &productsCategories_service.GetCompanyProductsRequest{
		CompanyId: request.CompanyId,
		Skip:      params.Skip,
		Limit:     params.Limit,
	})
	if err != nil {
		return &models.GetProductsList{}, err
	}
	var modelProduct models.GetProduct
	var modelProducts models.GetProductsList
	for _, result := range response.Products {
		description := sql.NullString{
			String: result.Description.String_,
			Valid:  result.Description.Valid}
		modelProduct = models.GetProduct{
			Id:          result.Id,
			Name:        result.Name,
			Description: description,
			Price:       result.Price,
			Photo:       result.Photo,
			Docs:        result.Documents,
		}
		modelProducts = append(modelProducts, modelProduct)
	}
	return &modelProducts, nil
}

func NewProductsCategoriesUseCase(PCgrpc ProductsCategoriesGRPC, authUsecase auth_usecase.UserUsecase, companyUsecase company_usecase.CompanyUseCase) ProductsCategoriesUseCase {
	return &productsCategoriesUseCase{ProductsCategoriesGRPC: PCgrpc, AuthUsecase: authUsecase, CompanyUsecase: companyUsecase}
}

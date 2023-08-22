package usecase

import (
	auth_usecase "b2b/m/internal/gateway/user/usecase"
	"b2b/m/internal/models"
	chttp "b2b/m/pkg/customhttp"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"
	"database/sql"
	"log"
)

type ProductsCategoriesUseCase interface {
	GetCategoryById(ctx context.Context, request *models.GetCategoryByIdRequest) (*models.GetCategoryByIdResponse, error)
	GetProductById(ctx context.Context, request *models.GetProductByIdRequest) (*models.GetProductByIdResponse, error)
	SearchCategories(ctx context.Context, request *chttp.SearchItemNameWithSkipLimit) (*[]models.GetCategoryByIdResponse, error)
	GetProductsList(ctx context.Context, request *chttp.QueryParam) (*models.GetProductsList, error)
	SearchProducts(ctx context.Context, request *chttp.SearchItemNameWithSkipLimit) (*models.GetProductsList, error)
	AddProduct(ctx context.Context, request *models.UserInfoAndAddProductByFormRequest) (*models.GetProduct, error)
}

type productsCategoriesUseCase struct {
	ProductsCategoriesGRPC ProductsCategoriesGRPC
	AuthUsecase            auth_usecase.UserUsecase
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
	}, nil
}
func (u *productsCategoriesUseCase) GetProductById(ctx context.Context, request *models.GetProductByIdRequest) (*models.GetProductByIdResponse, error) {
	response, err := u.ProductsCategoriesGRPC.GetProductById(ctx, &productsCategories_service.GetProductByID{
		Id: request.Id,
	})
	if err != nil {
		log.Println("Error GetProductById", err)
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
		}
		modelProducts = append(modelProducts, modelProduct)
	}
	return &modelProducts, nil
}

func NewProductsCategoriesUseCase(PCgrpc ProductsCategoriesGRPC, authUsecase auth_usecase.UserUsecase) ProductsCategoriesUseCase {
	return &productsCategoriesUseCase{ProductsCategoriesGRPC: PCgrpc, AuthUsecase: authUsecase}
}

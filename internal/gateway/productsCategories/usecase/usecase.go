package usecase

import (
	"b2b/m/internal/models"
	chttp "b2b/m/pkg/customhttp"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"
	"database/sql"
)

type ProductsCategoriesUseCase interface {
	GetCategoryById(ctx context.Context, request *models.GetCategoryByIdRequest) (*models.GetCategoryByIdResponse, error)
	SearchCategories(ctx context.Context, request *chttp.SearchItemName) (*[]models.GetCategoryByIdResponse, error)
	GetProductsList(ctx context.Context, request *chttp.QueryParam) (*models.GetProductsList, error)
	SearchProducts(ctx context.Context, request *chttp.SearchItemNameWithSkipLimit) (*models.GetProductsList, error)
}

type productsCategoriesUseCase struct {
	productsCategoriesGRPC productsCategoriesGRPC
}

func (u *productsCategoriesUseCase) GetCategoryById(ctx context.Context, request *models.GetCategoryByIdRequest) (*models.GetCategoryByIdResponse, error) {
	response, err := u.productsCategoriesGRPC.GetCategoryById(ctx, &productsCategories_service.GetCategoryByID{
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

func (u *productsCategoriesUseCase) SearchCategories(ctx context.Context, request *chttp.SearchItemName) (*[]models.GetCategoryByIdResponse, error) {
	SearchResults, err := u.productsCategoriesGRPC.SearchCategories(ctx, &productsCategories_service.SearchItemNameRequest{
		Name: request.Name,
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
	SearchResults, err := u.productsCategoriesGRPC.SearchProducts(ctx, &productsCategories_service.SearchItemNameWithSkipLimitRequest{
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
	response, err := u.productsCategoriesGRPC.GetProductsList(ctx, &productsCategories_service.GetProductsListRequest{
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

func NewProductsCategoriesUseCase(grpc productsCategoriesGRPC) ProductsCategoriesUseCase {
	return &productsCategoriesUseCase{productsCategoriesGRPC: grpc}
}

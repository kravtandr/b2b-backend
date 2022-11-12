package usecase

import (
	"b2b/m/internal/models"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"
	"database/sql"
)

type ProductsCategoriesUseCase interface {
	GetCategoryById(ctx context.Context, request *models.GetCategoryByIdRequest) (*models.GetCategoryByIdResponse, error)
	SearchCategories(ctx context.Context, request *models.SearchCategory) (*[]models.GetCategoryByIdResponse, error)
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

func (u *productsCategoriesUseCase) SearchCategories(ctx context.Context, request *models.SearchCategory) (*[]models.GetCategoryByIdResponse, error) {
	SearchResults, err := u.productsCategoriesGRPC.SearchCategories(ctx, &productsCategories_service.SearchCategoriesRequest{
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

func NewProductsCategoriesUseCase(grpc productsCategoriesGRPC) ProductsCategoriesUseCase {
	return &productsCategoriesUseCase{productsCategoriesGRPC: grpc}
}

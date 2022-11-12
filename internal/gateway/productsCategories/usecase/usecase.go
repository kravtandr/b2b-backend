package usecase

import (
	"b2b/m/internal/models"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"
)

type ProductsCategoriesUseCase interface {
	GetAllcategories(ctx context.Context, request *models.GetCategorieByIdRequest) (*models.GetCategorieByIdResponse, error)
}

type productsCategoriesUseCase struct {
	productsCategoriesGRPC productsCategoriesGRPC
}

func (u *productsCategoriesUseCase) GetAllcategories(ctx context.Context, request *models.GetCategorieByIdRequest) (*models.GetCategorieByIdResponse, error) {
	response, err := u.productsCategoriesGRPC.GetCategoryById(ctx, &productsCategories_service.GetCategoryByID{
		Id: request.Id,
	})
	if err != nil {
		return &models.GetCategorieByIdResponse{}, err
	}

	return &models.GetCategorieByIdResponse{
		Id:   request.Id,
		Name: response.Name,
	}, nil
}

func NewProductsCategoriesUseCase(grpc productsCategoriesGRPC) ProductsCategoriesUseCase {
	return &productsCategoriesUseCase{productsCategoriesGRPC: grpc}
}

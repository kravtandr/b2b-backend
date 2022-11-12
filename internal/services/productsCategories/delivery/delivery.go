package delivery

import (
	"b2b/m/internal/services/productsCategories/models"
	"b2b/m/internal/services/productsCategories/usecase"
	"b2b/m/pkg/error_adapter"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"
)

type productsCategoriesDelivery struct {
	productsCategoriesUseCase usecase.ProductsCategoriesUseCase
	errorAdapter              error_adapter.ErrorAdapter
	productsCategories_service.UnimplementedProductsCategoriesServiceServer
}

func (a *productsCategoriesDelivery) GetCategoryById(ctx context.Context, request *productsCategories_service.GetCategoryByID) (*productsCategories_service.GetCategory, error) {
	resp, err := a.productsCategoriesUseCase.GetCategoryById(ctx, &models.CategorieId{
		Id: request.Id,
	})
	if err != nil {
		return &productsCategories_service.GetCategory{}, a.errorAdapter.AdaptError(err)
	}

	return &productsCategories_service.GetCategory{
		Id:   resp.Id,
		Name: resp.Name,
	}, nil
}

func NewProductsCategoriesDelivery(
	productsCategoriesUseCase usecase.ProductsCategoriesUseCase,
	errorAdapter error_adapter.ErrorAdapter,
) productsCategories_service.ProductsCategoriesServiceServer {
	return &productsCategoriesDelivery{
		productsCategoriesUseCase: productsCategoriesUseCase,
		errorAdapter:              errorAdapter,
	}
}

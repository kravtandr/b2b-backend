package delivery

import (
	"b2b/m/internal/services/productsCategories/models"
	"b2b/m/internal/services/productsCategories/usecase"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/error_adapter"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"
)

type productsCategoriesDelivery struct {
	productsCategoriesUseCase usecase.ProductsCategoriesUseCase
	errorAdapter              error_adapter.ErrorAdapter
	productsCategories_service.UnimplementedProductsCategoriesServiceServer
}

func (a *productsCategoriesDelivery) GetProductsList(ctx context.Context, request *productsCategories_service.GetProductsListRequest) (*productsCategories_service.GetProductsListResponse, error) {
	resp, err := a.productsCategoriesUseCase.GetProductsList(ctx, &chttp.QueryParam{
		Skip:  request.Skip,
		Limit: request.Limit,
	})
	if err != nil {
		return &productsCategories_service.GetProductsListResponse{}, a.errorAdapter.AdaptError(err)
	}
	var res productsCategories_service.GetProductsListResponse
	var modelProduct *productsCategories_service.GetProduct

	for _, result := range *resp {
		description := productsCategories_service.SqlNullString{
			String_: result.Description.String,
			Valid:   result.Description.Valid}
		modelProduct = &productsCategories_service.GetProduct{
			Id:          result.Id,
			Name:        result.Name,
			Description: &description,
			Price:       result.Price,
			Photo:       result.Photo,
		}
		res.Products = append(res.Products, modelProduct)

	}
	return &res, nil
}

func (a *productsCategoriesDelivery) SearchProducts(ctx context.Context, request *productsCategories_service.SearchItemNameWithSkipLimitRequest) (*productsCategories_service.GetProductsListResponse, error) {
	resp, err := a.productsCategoriesUseCase.SearchProducts(ctx, &chttp.SearchItemNameWithSkipLimit{
		Name:  request.Name,
		Skip:  request.Skip,
		Limit: request.Limit,
	})
	if err != nil {
		return &productsCategories_service.GetProductsListResponse{}, a.errorAdapter.AdaptError(err)
	}
	var res productsCategories_service.GetProductsListResponse
	var modelProduct *productsCategories_service.GetProduct

	for _, result := range *resp {
		description := productsCategories_service.SqlNullString{
			String_: result.Description.String,
			Valid:   result.Description.Valid}
		modelProduct = &productsCategories_service.GetProduct{
			Id:          result.Id,
			Name:        result.Name,
			Description: &description,
			Price:       result.Price,
			Photo:       result.Photo,
		}
		res.Products = append(res.Products, modelProduct)

	}
	return &res, nil
}

func (a *productsCategoriesDelivery) GetCategoryById(ctx context.Context, request *productsCategories_service.GetCategoryByID) (*productsCategories_service.GetCategory, error) {
	resp, err := a.productsCategoriesUseCase.GetCategoryById(ctx, &models.CategoryId{
		Id: request.Id,
	})
	if err != nil {
		return &productsCategories_service.GetCategory{}, a.errorAdapter.AdaptError(err)
	}
	description := productsCategories_service.SqlNullString{
		String_: resp.Description.String,
		Valid:   resp.Description.Valid}
	return &productsCategories_service.GetCategory{
		Id:          resp.Id,
		Name:        resp.Name,
		Description: &description,
	}, nil
}

func (a *productsCategoriesDelivery) SearchCategories(ctx context.Context, request *productsCategories_service.SearchItemNameRequest) (*productsCategories_service.GetCategories, error) {
	resp, err := a.productsCategoriesUseCase.SearchCategories(ctx, request.Name)
	if err != nil {
		return &productsCategories_service.GetCategories{}, a.errorAdapter.AdaptError(err)
	}

	var res productsCategories_service.GetCategories
	var modelCategory *productsCategories_service.GetCategory

	for _, result := range *resp {
		description := productsCategories_service.SqlNullString{
			String_: result.Description.String,
			Valid:   result.Description.Valid}
		modelCategory = &productsCategories_service.GetCategory{
			Id:          result.Id,
			Name:        result.Name,
			Description: &description,
		}
		res.Categories = append(res.Categories, modelCategory)

	}
	return &res, nil
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

package usecase

import (
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"

	"google.golang.org/grpc"
)

type productsCategoriesGRPC interface {
	GetCategoryById(ctx context.Context, in *productsCategories_service.GetCategoryByID, opts ...grpc.CallOption) (*productsCategories_service.GetCategory, error)
	SearchCategories(ctx context.Context, in *productsCategories_service.SearchItemNameRequest, opts ...grpc.CallOption) (*productsCategories_service.GetCategories, error)
	GetProductsList(ctx context.Context, in *productsCategories_service.GetProductsListRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsListResponse, error)
	SearchProducts(ctx context.Context, in *productsCategories_service.SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsListResponse, error)
}

package usecase

import (
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"

	"google.golang.org/grpc"
)

type productsCategoriesGRPC interface {
	GetCategoryById(ctx context.Context, in *productsCategories_service.GetCategoryByID, opts ...grpc.CallOption) (*productsCategories_service.GetCategory, error)
	SearchCategories(ctx context.Context, in *productsCategories_service.SearchCategoriesRequest, opts ...grpc.CallOption) (*productsCategories_service.GetCategories, error)
	GetProductsList(ctx context.Context, in *productsCategories_service.SearchCategoriesRequest, opts ...grpc.CallOption) (*productsCategories_service.GetCategories, error)
	SearchProducts(ctx context.Context, in *productsCategories_service.SearchCategoriesRequest, opts ...grpc.CallOption) (*productsCategories_service.GetCategories, error)
}

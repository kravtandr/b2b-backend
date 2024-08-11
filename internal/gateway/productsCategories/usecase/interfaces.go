package usecase

import (
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"context"

	"google.golang.org/grpc"
)

type ProductsCategoriesGRPC interface {
	AddProduct(ctx context.Context, in *productsCategories_service.AddProductRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProduct, error)
	GetProductById(ctx context.Context, in *productsCategories_service.GetProductByID, opts ...grpc.CallOption) (*productsCategories_service.GetProduct, error)
	UpdateProduct(ctx context.Context, in *productsCategories_service.UpdateProductRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProduct, error)

	GetProductsList(ctx context.Context, in *productsCategories_service.GetProductsListRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsListResponse, error)
	GetProductsListByFilters(ctx context.Context, in *productsCategories_service.GetProductsListByFiltersRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsByFiltersResponse, error)
	SearchProducts(ctx context.Context, in *productsCategories_service.SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsListResponse, error)
	GetCompanyProducts(ctx context.Context, in *productsCategories_service.GetCompanyProductsRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsListResponse, error)

	GetCategoryById(ctx context.Context, in *productsCategories_service.GetCategoryByID, opts ...grpc.CallOption) (*productsCategories_service.GetCategory, error)
	SearchCategories(ctx context.Context, in *productsCategories_service.SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*productsCategories_service.GetCategories, error)
}

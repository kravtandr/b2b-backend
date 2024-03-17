package usecase

import (
	company_service "b2b/m/pkg/services/company"
	"context"

	"google.golang.org/grpc"
)

type CompanyGRPC interface {
	GetCompanyById(ctx context.Context, in *company_service.GetCompanyRequestById, opts ...grpc.CallOption) (*company_service.GetPrivateCompanyResponse, error)
	UpdateCompanyByOwnerId(ctx context.Context, in *company_service.UpdateCompanyRequest, opts ...grpc.CallOption) (*company_service.GetCompanyAndPostResponse, error)
	GetCompanyByProductId(ctx context.Context, in *company_service.IdRequest, opts ...grpc.CallOption) (*company_service.GetPrivateCompanyResponse, error)
}

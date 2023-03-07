package usecase

import (
	company_service "b2b/m/pkg/services/company"
	"context"
	"google.golang.org/grpc"
)

type CompanyGRPC interface {
	GetCompanyById(ctx context.Context, in *company_service.GetCompanyRequestById, opts ...grpc.CallOption) (*company_service.GetCompanyResponse, error)
	UpdateCompanyByOwnerId(ctx context.Context, in *company_service.UpdateCompanyRequest, opts ...grpc.CallOption) (*company_service.GetCompanyAndPostResponse, error)
}

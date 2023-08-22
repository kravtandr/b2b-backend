package setup

import (
	"context"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"

	"b2b/m/internal/services/productsCategories/config"
	"b2b/m/internal/services/productsCategories/delivery"
	"b2b/m/internal/services/productsCategories/repository"
	productsCategories_usecase "b2b/m/internal/services/productsCategories/usecase"
	"b2b/m/pkg/error_adapter"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/helpers"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}
	// Initialize minio client object.
	minioClient, err := minio.New(cfg.MinioUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioUser, cfg.MinioPass, ""),
		Secure: cfg.MinioSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	BucketName := "photo"
	location := "eu-central-1"
	err = minioClient.MakeBucket(context.Background(), BucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(context.Background(), BucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", BucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", BucketName)
	}

	productsCategoriesRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewProductsCategoriesRepository(
			repository.NewQueryFactory(), conn, minioClient,
		),
	)
	productsCategoriesUseCase := productsCategories_usecase.NewProductsCategoriesUseCase(productsCategoriesRepo)
	productsCategoriesDelivery := delivery.NewProductsCategoriesDelivery(productsCategoriesUseCase, error_adapter.NewErrorAdapter(grpc_errors.PreparedProductsServiceErrorMap))

	server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	productsCategories_service.RegisterProductsCategoriesServiceServer(server, productsCategoriesDelivery)
	grpc_prometheus.Register(server)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}

package config

import (
	logs "b2b/m/pkg/logger"
	"context"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	ENDPOINT                          string `envconfig:"ENDPOINT" required:"true"`
	AuthServiceEndpoint               string `envconfig:"AUTH_ENDPOINT" required:"true"`
	FastOrderServiceEndpoint          string `envconfig:"FASTORDER_ENDPOINT" required:"true"`
	CompanyServiceEndpoint            string `envconfig:"COMPANY_ENDPOINT" required:"true"`
	ProductsCategoriesServiceEndpoint string `envconfig:"PRODUCTSCATEGORIES_ENDPOINT" required:"true"`
	ChatServiceEndpoint               string `envconfig:"CHAT_ENDPOINT" required:"true"`
	HTTPPort                          string `envconfig:"HTTP_PORT" required:"true"`
	DBUrl                             string `envconfig:"DB_URL" required:"true"`
	DADATA_API_KEY                    string `envconfig:"DADATA_API_KEY" required:"true"`
	DADATA_SECRET_KEY                 string `envconfig:"DADATA_SECRET_KEY" required:"true"`
	UKASSA_SECRET_KEY                 string `envconfig:"UKASSA_SECRET_KEY" required:"true"`
	UKASSA_SHOP_ID                    string `envconfig:"UKASSA_SHOP_ID" required:"true"`
	// TODO add yookassa here
	// DefaultBucket    string `envconfig:"S3_BUCKET" required:"true"`
	// S3Endpoint       string `envconfig:"S3_ENDPOINT" required:"true"`
	// S3PublicEndpoint string `envconfig:"S3_PUBLIC_ENDPOINT" required:"true"`
	// SecretKey        string `envconfig:"S3_SECRET_KEY" required:"true"`
	// ID               string `envconfig:"S3_SECRET_ID" required:"true"`

	Ctx    context.Context
	Cancel func()
	Logger *zap.Logger
}

func (c *Config) Setup() error {

	if err := envconfig.Process("GATEWAY", c); err != nil {
		return err
	}
	lgr := logs.BuildLogger()
	c.Logger = lgr.Logger
	c.Ctx, c.Cancel = context.WithCancel(context.Background())
	fmt.Println("GATEWAY_CHAT_ENDPOINT", c.ChatServiceEndpoint)
	return nil
}

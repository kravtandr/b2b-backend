package config

import (
	"context"

	logs "b2b/m/pkg/logger"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	DBUrl     string `envconfig:"DB_URL" required:"true"`
	GRPCPort  string `envconfig:"GRPC_PORT" required:"true"`
	MinioUrl  string `envconfig:"MINIO_URL" required:"true"`
	MinioUser string `envconfig:"MINIO_USER" required:"true"`
	MinioPass string `envconfig:"MINIO_PASS" required:"true"`
	MinioSSL  bool   `envconfig:"MINIO_SSL" required:"true"`

	Ctx    context.Context
	Cancel func()
	Logger *zap.Logger
}

func (c *Config) Setup() error {
	if err := envconfig.Process("PRODUCTSCATEGORIES", c); err != nil {
		return err
	}

	lgr := logs.BuildLogger()
	c.Logger = lgr.Logger
	c.Ctx, c.Cancel = context.WithCancel(context.Background())

	return nil
}

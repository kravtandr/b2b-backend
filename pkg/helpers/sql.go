package helpers

import (
	"context"

	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func CreatePGXPool(ctx context.Context, connection string, logger *zap.Logger) (conn *pgxpool.Pool, err error) {
	config, err := pgxpool.ParseConfig(connection)
	if err != nil {
		return nil, err
	}
	config.ConnConfig.Logger = zapadapter.NewLogger(logger)

	return pgxpool.ConnectConfig(ctx, config)
}

func TimestapzToString(connection string, logger *zap.Logger) {

}

// const ctLayout = "2006-01-02T15:04:05Z"

// testTime, err := time.Parse(time.RFC3339, ctLayout)
// if err != nil {
// 	fmt.Println(err)

// }

// fmt.Println(testTime)

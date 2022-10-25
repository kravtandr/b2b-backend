package repository

import (
	"b2b/m/internal/services/fastOrder/models"
	"context"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type FastOrderRepository interface {
	FastOrder(ctx context.Context, user *models.OrderForm) error
}

type fastOrderRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (a *fastOrderRepository) FastOrder(ctx context.Context, order *models.OrderForm) error {
	query := a.queryFactory.CreateFastOrder(order)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		return err
	}
	return nil
}

func NewFastOrderRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) FastOrderRepository {
	return &fastOrderRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}

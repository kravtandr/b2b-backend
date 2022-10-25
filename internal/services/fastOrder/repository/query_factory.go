package repository

import (
	"b2b/m/internal/services/fastOrder/models"
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateFastOrder(user *models.OrderForm) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateFastOrder(order *models.OrderForm) *query.Query {
	return &query.Query{
		Request: CreateFastOrder,
		Params: []interface{}{
			order.Role, order.Product_category, order.Product_name, order.Order_text, order.Order_comments, order.Fio, order.Email, order.Phone, order.Company_name, order.Itn,
		},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}

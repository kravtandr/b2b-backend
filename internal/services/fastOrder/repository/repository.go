package repository

import (
	"b2b/m/internal/services/fastOrder/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

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
	sendMessageToBot(order)
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

func sendMessageToBot(order *models.OrderForm) {
	if order.Company_name == "test" {
		return
	}
	orderDetails := fmt.Sprintf("Id: %d\nНазвание компании: %s\nФИО: %s\nEmail: %s\nТелефон: %s\nИНН: %s\n"+
		"Категория товара: %s\nНазвание товара: %s\nОписание заказа: %s\nКомментарии к заказу: %s", order.Id,
		order.Company_name, order.Fio, order.Email, order.Phone, order.Itn, order.Product_category, order.Product_name,
		order.Order_text, order.Order_comments)
	chatId := "-1001625808882"
	botToken := "5653463229:AAGkQgSIkaSH-MaE9-PswOQWtMQZITJ2_Hk"
	query := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", botToken,
		chatId, url.QueryEscape(orderDetails))
	_, err := http.Get(query)
	if err != nil {
		log.Fatalln(err)
	}
}

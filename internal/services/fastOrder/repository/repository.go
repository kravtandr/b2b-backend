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
	LandingOrder(ctx context.Context, user *models.LandingForm) error
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
	sendMessageToBotNewFastOrder(order)
	return nil
}

func (a *fastOrderRepository) LandingOrder(ctx context.Context, order *models.LandingForm) error {
	query := a.queryFactory.CreateLandingOrder(order)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		return err
	}
	sendMessageToBotNewLandingOrder(order)
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

func sendMessageToBotNewLandingOrder(order *models.LandingForm) {
	if order.Order_text == "test" {
		return
	}
	orderDetails := fmt.Sprintf("Email: %s\nИНН: %s\n"+
		"Категория товара: %s\nАдрес доставки: %s\nДата доставки: %s\nОписание заявки: %s",
		order.Email, order.Itn, order.Product_category, order.Delivery_address, order.Delivery_date,
		order.Order_text)
	chatId := "-1001625808882"
	botToken := "5653463229:AAGkQgSIkaSH-MaE9-PswOQWtMQZITJ2_Hk"
	query := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", botToken,
		chatId, url.QueryEscape(orderDetails))
	_, err := http.Get(query)
	if err != nil {
		log.Fatalln(err)
	}
}

func sendMessageToBotNewFastOrder(order *models.OrderForm) {
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

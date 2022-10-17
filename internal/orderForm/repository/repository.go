package OrderFormRepository

import (
	"b2b/m/pkg/domain"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type OrderFormStorage struct {
	dataHolder *pgxpool.Pool
}

func NewOrderFormStorage(DB *pgxpool.Pool) domain.OrderFormStorage {
	return &OrderFormStorage{dataHolder: DB}
}

func (u *OrderFormStorage) Add(value domain.OrderForm) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		log.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(),
		`INSERT INTO OrderForm ("role", "product_category", "product_name", "order_text", "order_comments", "fio", "email", "phone", "company_name", "itn") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,

		value.Role,
		value.Product_category,
		value.Product_name,
		value.Order_text,
		value.Order_comments,
		value.Fio,
		value.Email,
		value.Phone,
		value.Company_name,
		value.Itn,
	)
	return err
}

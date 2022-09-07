package main

import (
	catd "b2b/m/internal/category/delivery"
	comd "b2b/m/internal/company/delivery"
	ordford "b2b/m/internal/orderForm/delivery"
	ud "b2b/m/internal/user/delivery"
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/fasthttp/router"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/valyala/fasthttp"
	"gopkg.in/webdeskltd/dadata.v2"
)

func SetUpRouter(db *pgxpool.Pool, daData *dadata.DaData) *router.Router {
	r := router.New()
	r = ud.SetUpUserRouter(db, daData, r)
	r = comd.SetUpCompanyRouter(db, daData, r)
	r = catd.SetUpCategoryRouter(db, r)
	r = ordford.SetUpOrderFormRouter(db, r)
	return r
}

func corsMiddleware(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", `http://bi-tu-bi.ru`) // prod
		//ctx.Response.Header.Set("Access-Control-Allow-Origin", `http://localhost:3000`) // local
		ctx.Response.Header.Set("Content-Type", "application/json; charset=utf8")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		ctx.Response.Header.Set("Access-Control-Expose-Headers", "Authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Max-Age", "3600")

		// ctx.Response.Header.Set(
		// 	"Access-Control-Allow-Headers",
		// 	"Origin, Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers",
		// )
		if bytes.Equal(ctx.Method(), []byte(fasthttp.MethodOptions)) {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		handler(ctx)
	}
}

func main() {
	fmt.Println("starting server at :8080")
	url := "postgres://b2b:b2b@localhost:5432/b2b"
	dbpool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	daData := dadata.NewDaData("42e877cc6e66e3cc70c47a2f42966120cfcea751", "984e0c50d52dd2611b98609eaa7c82268e46297e")

	r := SetUpRouter(dbpool, daData)
	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}

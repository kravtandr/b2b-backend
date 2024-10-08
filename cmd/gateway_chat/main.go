package main

import (
	"bytes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"b2b/m/internal/gateway_chat/config"
	"b2b/m/internal/gateway_chat/setup"

	"github.com/valyala/fasthttp"
)

func main() {
	var cfg config.Config
	if err := cfg.Setup(); err != nil {
		log.Fatal("failed to setup cfg: ", err)
		return
	}

	logger := cfg.Logger.Sugar()
	p, stop, err := setup.Setup(cfg)
	if err != nil {
		logger.Fatal("msg", "failed to setup server", "error", err)
		return
	}

	go func() {
		if err := fasthttp.ListenAndServe(cfg.HTTPPort, corsMiddleware(p.GetHandler())); err != nil {
			logger.Fatal("failed to start server")
			return
		}
	}()
	logger.Info("gateway_chat started ...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		logger.Info("msg", "received signal, exiting", "signal", sig)
		cfg.Cancel()
		stop()

		logger.Info("msg", " goodbye")
	}(<-c)
}

func corsMiddleware(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	var cfg config.Config
	if err := cfg.Setup(); err != nil {
		log.Fatal("failed to setup cfg: ", err)
		return nil
	}
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", cfg.ENDPOINT)
		// ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:3000") // local
		// ctx.Response.Header.Set("Access-Control-Allow-Origin", "https://bi-tu-bi.ru") // deploy
		ctx.Response.Header.Set("Content-Type", "application/json; charset=utf8")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		ctx.Response.Header.Set("Access-Control-Expose-Headers", "Authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Max-Age", "3600")

		if bytes.Equal(ctx.Method(), []byte(fasthttp.MethodOptions)) {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		handler(ctx)
	}
}

// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"b2b/m/internal/services/chat/config"
	"b2b/m/internal/services/chat/setup"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fasthttp/websocket"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
)

var addr = flag.String("addr", "127.0.0.1:5001", "http service address")

var upgrader = websocket.FastHTTPUpgrader{}

func main() {
	var cfg config.Config
	if err := cfg.Setup(); err != nil {
		log.Fatal("failed to setup cfg: ", err)
		os.Exit(3)
		return
	}

	logger := cfg.Logger.Sugar()
	server, cancel, err := setup.SetupServer(cfg)
	if err != nil {
		logger.Fatal("msg", "failed to setup server", "error", err)
		return
	}

	go func() {
		logger.Info("msg", "starting grpc server", "port", cfg.GRPCPort)
		lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
		if err != nil {
			logger.Error("msg", "grpc server listen", "err", err)
			os.Exit(1)
		}
		logger.Info("msg", "grpc server listener started")

		if err := server.Serve(lis); err != nil {
			logger.Error("msg", "grpc server run failuer", "err", err)
			os.Exit(1)
		}
	}()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	logger.Info("chat service started ...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		logger.Info("msg", "received signal, exiting", "signal", sig)
		server.GracefulStop()
		cfg.Cancel()
		cancel()

		logger.Info("msg", " goodbye")
	}(<-c)

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
	WSserver := fasthttp.Server{
		Name:    "EchoExample",
		Handler: requestHandler,
	}
	log.Fatal(WSserver.ListenAndServe(*addr))
}

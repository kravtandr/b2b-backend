// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/fasthttp/websocket"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

var addr = flag.String("addr", "localhost:5001", "http service address")

var upgrader = websocket.FastHTTPUpgrader{}

type Msg struct {
	Msg       string `json:"msg"`
	UserID    int64  `json:"userID"`
	ProductID int64  `json:"productID"`
}

func initDB(ctx *fasthttp.RequestCtx) {
	//jdbc:postgresql://localhost:5431/b2b
	//postgres://b2b:b2b@postgres:5432/b2b
	config, err := pgxpool.ParseConfig("postgres://b2b:b2b@localhost:5431/b2b")
	if err != nil {
		log.Println("IN CHAT ParseConfig ERROR:", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Println("IN CHAT ConnectConfig ERROR:", err)
	}
	pool.Ping(context.Background())
	if err != nil {
		log.Println("IN CHAT Ping DB ERROR:", err)
	}
}

func WriteNewMsg(msg Msg) error {
	// row := pool.QueryRow(context.Background(), "SELECT * from users;")
	// repoCompany := &models.PublicUser{}
	// if err := row.Scan(
	// 	&repoCompany.Email,
	// ); err != nil {
	// 	log.Println("IN CHAT DB ERROR:", err)
	// 	if err == pgx.ErrNoRows {
	// 		return
	// 	}

	// 	return
	// }
	return nil
}

func echoView(ctx *fasthttp.RequestCtx) {
	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		defer ws.Close()
		//первое сообщение приходит с фронта
		msg := Msg{Msg: "Сколько единиц в комлекте?", UserID: 1, ProductID: 1}
		bytes, _ := json.Marshal(msg)
		// 1 - binary, 2 - text
		err := ws.WriteMessage(1, bytes)
		initDB(ctx)
		if err != nil {
			log.Println("WS write:", err)
		}
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				//когда приходит сообщение записываю его в бд
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			//когда отправляю сообщение записываю его в бд
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})

	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}

// func homeView(ctx *fasthttp.RequestCtx) {
// 	ctx.SetContentType("text/html")
// 	homeTemplate.Execute(ctx, "ws://"+string(ctx.Host())+"/echo")
// }

func main() {
	fmt.Println("________________________________________________________")
	fmt.Println("________________________________________________________")
	flag.Parse()
	log.SetFlags(0)

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/ws":
			echoView(ctx)
		// case "/":
		// 	homeView(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	server := fasthttp.Server{
		Name:    "EchoExample",
		Handler: requestHandler,
	}

	log.Fatal(server.ListenAndServe(*addr))
}

// var homeTemplate = template.Must(template.New("").Parse(`
// <!DOCTYPE html>
// <html>
// <head>
// <meta charset="utf-8">
// <script>
// window.addEventListener("load", function(evt) {

//     var output = document.getElementById("output");
//     var input = document.getElementById("input");
//     var ws;

//     var print = function(message) {
//         var d = document.createElement("div");
//         d.textContent = message;
//         output.appendChild(d);
//         output.scroll(0, output.scrollHeight);
//     };

//     document.getElementById("open").onclick = function(evt) {
//         if (ws) {
//             return false;
//         }
//         ws = new WebSocket("{{.}}");
//         ws.onopen = function(evt) {
//             print("OPEN");
//         }
//         ws.onclose = function(evt) {
//             print("CLOSE");
//             ws = null;
//         }
//         ws.onmessage = function(evt) {
//             print("RESPONSE: " + evt.data);
//         }
//         ws.onerror = function(evt) {
//             print("ERROR: " + evt.data);
//         }
//         return false;
//     };

//     document.getElementById("send").onclick = function(evt) {
//         if (!ws) {
//             return false;
//         }
//         print("SEND: " + input.value);
//         ws.send(input.value);
//         return false;
//     };

//     document.getElementById("close").onclick = function(evt) {
//         if (!ws) {
//             return false;
//         }
//         ws.close();
//         return false;
//     };

// });
// </script>
// </head>
// <body>
// <table>
// <tr><td valign="top" width="50%">
// <p>Click "Open" to create a connection to the server,
// "Send" to send a message to the server and "Close" to close the connection.
// You can change the message and send multiple times.
// <p>
// <form>
// <button id="open">Open</button>
// <button id="close">Close</button>
// <p><input id="input" type="text" value="Hello world!">
// <button id="send">Send</button>
// </form>
// </td><td valign="top" width="50%">
// <div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
// </td></tr></table>
// </body>
// </html>
// `))

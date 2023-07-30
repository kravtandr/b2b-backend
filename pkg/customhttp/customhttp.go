package customhttp

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	cnst "b2b/m/pkg/constants"

	"github.com/valyala/fasthttp"
)

type Resp struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type QueryParam struct {
	Skip  int64
	Limit int64
}

type SearchItemName struct {
	Name string `json:"name"`
}

type SearchItemNameWithSkipLimit struct {
	Name  string
	Skip  int64
	Limit int64
}

func ApiResp(value interface{}, err error) ([]byte, error) {
	var r Resp
	if err != nil {
		r = Resp{Data: value, Msg: fmt.Sprintln(err)}
	} else {
		r = Resp{Data: value, Msg: "OK"}

	}
	bytes, err := json.Marshal(r)
	if err != nil {
		r.Msg = "error while marshalling JSON"
		log.Printf("error while marshalling JSON: %s", err)
	}
	return bytes, err
}

func GetQueryParams(ctx *fasthttp.RequestCtx) (*QueryParam, error) {
	skip_byte := ctx.QueryArgs().Peek("skip")
	skip, err := strconv.ParseInt(string(skip_byte), 10, 64)
	if err != nil {
		return &QueryParam{Skip: 0, Limit: 1}, err
	}
	limit_byte := ctx.QueryArgs().Peek("limit")
	limit, err := strconv.ParseInt(string(limit_byte), 10, 64)
	if err != nil {
		return &QueryParam{Skip: 0, Limit: 1}, err
	}
	var params = &QueryParam{
		Skip:  skip,
		Limit: limit,
	}
	fmt.Println("GetQueryParams || SKIP ====== ", skip, "LIMIT ====== ", limit)
	return params, nil
}

func SetCookieAndSession(ctx *fasthttp.RequestCtx, cookie string) {
	var c fasthttp.Cookie
	c.SetPath("/")
	c.SetKey(cnst.CookieName)
	c.SetValue(cookie)
	c.SetMaxAge(int(time.Hour))
	c.SetHTTPOnly(true)
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	ctx.Response.Header.SetCookie(&c)
}

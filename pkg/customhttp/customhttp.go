package customhttp

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

type Resp struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type QueryParam struct {
	Skip  int
	Limit int
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
	skip, err := strconv.Atoi(ctx.UserValue("skip").(string))
	if err != nil {
		return &QueryParam{}, err
	}
	limit, err := strconv.Atoi(ctx.UserValue("limit").(string))
	if err != nil {
		return &QueryParam{}, err
	}
	var params = &QueryParam{
		Skip:  skip,
		Limit: limit,
	}
	return params, nil
}

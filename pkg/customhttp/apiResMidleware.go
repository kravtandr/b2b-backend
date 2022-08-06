package customhttp

import (
	"encoding/json"
	"fmt"
	"log"
)

type Resp struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func ApiResp(value interface{}, err error) ([]byte, error) {
	r := Resp{Data: value, Msg: fmt.Sprintln(err)}
	bytes, err := json.Marshal(r)
	if err != nil {
		r.Msg = "error while marshalling JSON"
		log.Printf("error while marshalling JSON: %s", err)
	}
	return bytes, err
}

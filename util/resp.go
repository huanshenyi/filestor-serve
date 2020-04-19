package util

import (
	"encoding/json"
	"fmt"
	"log"
)

// httpリスポンスの構造
type RespMsg struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

// NewRespMsg: responseを生成する
func NewRespMsg(code int, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// JSONBytes:objectをjson形式のバイナリリスト
func (resp *RespMsg)JSONBytes()[]byte {
	r, err := json.Marshal(resp)
	if err != nil{
		log.Println(err)
	}
	return r
}

// JSONString:objectをjson形式のstringに変化
func (resp *RespMsg)JSONString()string {
	r, err := json.Marshal(resp)
	if err != nil{
		log.Println(err)
	}
	return string(r)
}

// GenSimpleRespStream : 只包含code和message的响应体([]byte)
func GenSimpleRespStream(code int, msg string) []byte {
	return []byte(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
}

// GenSimpleRespString : 只包含code和message的响应体(string)
func GenSimpleRespString(code int, msg string) string {
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg)
}


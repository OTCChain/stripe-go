package rpc

import (
	"encoding/json"
	"fmt"
)

const (
	vsn = "1.0"
)

type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JError(msg error) *JsonError {
	return &JsonError{
		Code:    -1,
		Message: msg.Error(),
	}
}

func JSError(fm string, msg ...interface{}) *JsonError {
	return &JsonError{
		Code:    -1,
		Message: fmt.Sprintf(fm, msg...),
	}
}

type JsonRpcMessageItem struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      uint32          `json:"id,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *JsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func (i *JsonRpcMessageItem) String() string {
	s := fmt.Sprintf("\n<-------------rpc request-----------")
	s += fmt.Sprintf("\n*version:			%s", i.Version)
	s += fmt.Sprintf("\n*ID:			%d", i.ID)
	s += fmt.Sprintf("\n*Params:			%s", string(i.Params))
	s += fmt.Sprintf("\n----------------------------------->\n")

	return s
}

type JsonRpcMessage []*JsonRpcMessageItem

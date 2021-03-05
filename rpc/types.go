package rpc

import "encoding/json"

const (
	vsn = "1.0"
)

type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type JsonRpcMessageItem struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      uint32          `json:"id,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *JsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type JsonRpcMessage []*JsonRpcMessageItem

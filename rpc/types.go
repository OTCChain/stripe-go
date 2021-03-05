package rpc

import "encoding/json"

type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type JsonRpcMessageItem struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *JsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type JsonRpcMessage []*JsonRpcMessageItem

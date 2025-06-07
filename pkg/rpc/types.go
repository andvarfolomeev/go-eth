package rpc

import "encoding/json"

type Request struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

type Response struct {
	Jsonrpc string           `json:"jsonrpc"`
	Id      int              `json:"id"`
	Result  *json.RawMessage `json:"result"`
	Error   *Error           `json:"error"`
}

package ethrpc

import (
	"fmt"
)

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("jsonrpc error: %d: %s", e.Code, e.Message)
}

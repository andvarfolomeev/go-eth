package batcher

import "go-eth/pkg/rpc"

type RpcBatcher struct {
	reqs []rpc.Request
}

func New() *RpcBatcher {
	return &RpcBatcher{}
}

func (b *RpcBatcher) Add(method string, params ...any) int {
	id := len(b.reqs) + 1
	b.reqs = append(b.reqs, rpc.Request{
		Jsonrpc: "2.0",
		Id:      id,
		Method:  method,
		Params:  params,
	})
	return id
}

func (b *RpcBatcher) Request() []rpc.Request {
	return b.reqs
}

func (b *RpcBatcher) Clear() {
	b.reqs = b.reqs[:0]
}

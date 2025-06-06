package batcher

import "go-eth/pkg/rpc"

type Batcher struct {
	reqs []rpc.Request
}

func New() *Batcher {
	return &Batcher{}
}

func (b *Batcher) Add(method string, params ...any) int {
	id := len(b.reqs) + 1
	b.reqs = append(b.reqs, rpc.Request{
		Jsonrpc: "2.0",
		Id:      id,
		Method:  method,
		Params:  params,
	})
	return id
}

func (b *Batcher) Request() []rpc.Request {
	return b.reqs
}

func (b *Batcher) Clear() {
	b.reqs = b.reqs[:0]
}

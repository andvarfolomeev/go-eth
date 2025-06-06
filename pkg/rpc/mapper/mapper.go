package mapper

import (
	"errors"
	"fmt"
	"go-eth/pkg/rpc"
)

type RpcMapper struct {
	data map[int]rpc.Response
}

func New(resps []rpc.Response) *RpcMapper {
	data := make(map[int]rpc.Response)

	for _, resp := range resps {
		id := resp.Id
		data[id] = resp
	}

	return &RpcMapper{
		data: data,
	}
}

func (mb RpcMapper) GetByID(id int) (rpc.Response, error) {
	data, ok := mb.data[id]
	if !ok {
		return rpc.Response{}, errors.New(fmt.Sprintf("Failed to find response by id = %d", id))
	}

	return data, nil
}

package mapper

import (
	"errors"
	"fmt"
	"go-eth/pkg/rpc"
)

type Mapper struct {
	data map[int]rpc.Response
}

func New(resps []rpc.Response) *Mapper {
	data := make(map[int]rpc.Response)

	for _, resp := range resps {
		id := resp.Id
		data[id] = resp
	}

	return &Mapper{
		data: data,
	}
}

func (mb Mapper) GetByID(id int) (rpc.Response, error) {
	data, ok := mb.data[id]
	if !ok {
		return rpc.Response{}, errors.New(fmt.Sprintf("Failed to find response by id = %d", id))
	}

	return data, nil
}

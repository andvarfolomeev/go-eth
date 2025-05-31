package ethindexer

import (
	"context"

	"github.com/andvarfolomeev/go-eth/pkg/ethrpc"
)

func getEarlistAndLastBlock(ctx context.Context, client ethrpc.Client) (string, string, error) {
	type result struct {
		earliest string
		latest   string
		err      error
	}

	resCh := make(chan result, 2)

	go func() {
		block, err := client.BlockByNumber(ctx, "earliest")
		if err != nil {
			resCh <- result{err: err}
		}
		resCh <- result{earliest: block.Number}
	}()

	go func() {
		block, err := client.BlockByNumber(ctx, "latest")
		if err != nil {
			resCh <- result{err: err}
		}
		resCh <- result{latest: block.Number}
	}()

	var earliest, latest string

	for range 2 {
		select {
		case <-ctx.Done():
			return "", "", ctx.Err()
		case res := <-resCh:
			switch {
			case res.err != nil:
				return "", "", res.err
			case res.earliest != "":
				earliest = res.earliest
			case res.latest != "":
				earliest = res.latest
			}
		}
	}

	return earliest, latest, nil
}

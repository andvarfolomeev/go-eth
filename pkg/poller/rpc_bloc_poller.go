package poller

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"go-eth/pkg/conv"
	"go-eth/pkg/rpc"
)

const (
	InitialBlock = "0"
)

type RPCBlockPoller struct {
	client   *rpc.Client
	interval time.Duration
}

func NewRPCBlockPoller(client *rpc.Client, interval time.Duration) BlockPoller {
	return &RPCBlockPoller{
		client:   client,
		interval: interval,
	}
}

func (ebp *RPCBlockPoller) Poll(ctx context.Context, outCh chan string, errCh chan error) {
	go ebp.poll(ctx, outCh, errCh)
}

func (ebp *RPCBlockPoller) poll(ctx context.Context, outCh chan string, errCh chan error) {
	defer close(outCh)
	defer close(errCh)

	ticker := time.NewTicker(ebp.interval)
	defer ticker.Stop()

	latestBlockHex := InitialBlock

	for {
		select {
		case <-ticker.C:
			newBlocks, err := ebp.getNewBlocks(ctx, latestBlockHex)

			if err != nil {
				select {
				case errCh <- err:
				case <-ctx.Done():
					return
				}
				continue
			}

			for _, blockHex := range newBlocks {
				select {
				case outCh <- blockHex:
				case <-ctx.Done():
					return
				}
			}

			if len(newBlocks) > 0 {
				latestBlockHex = newBlocks[len(newBlocks)-1]
			}

		case <-ctx.Done():
			return
		}
	}
}

func (ebp *RPCBlockPoller) getNewBlocks(ctx context.Context, latestBlockHex string) ([]string, error) {
	resp, err := ebp.client.Call(
		ctx,
		rpc.Request{
			Jsonrpc: "2.0",
			Id:      0,
			Method:  "eth_blockNumber",
			Params:  []interface{}{},
		},
	)

	if err != nil {
		return nil, err
	}

	var newLatestBlock string
	if err := json.Unmarshal(*resp.Result, &newLatestBlock); err != nil {
		return nil, nil
	}

	if latestBlockHex == InitialBlock {
		return []string{newLatestBlock}, nil
	}

	blocks, err := generateBlockRange(latestBlockHex, newLatestBlock)

	if err != nil {
		return nil, err
	}

	return blocks, nil
}

func generateBlockRange(start, end string) ([]string, error) {
	startUint64, err := conv.HexToUint64(start)

	if err != nil {
		return nil, fmt.Errorf("Invalid start block hex %q: %w", start, err)
	}

	endUint64, err := conv.HexToUint64(end)

	if err != nil {
		return nil, fmt.Errorf("Invalid end block hex %q: %w", start, err)
	}

	if startUint64 > endUint64 {
		startUint64, endUint64 = endUint64, startUint64
	}

	blocks := make([]string, 0, endUint64-startUint64)

	for i := startUint64 + 1; i < endUint64; i++ {
		blocks = append(blocks, fmt.Sprintf("0x%s", strconv.FormatUint(i, 16)))
	}

	return blocks, nil
}

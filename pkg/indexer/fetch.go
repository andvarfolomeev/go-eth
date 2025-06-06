package indexer

import (
	"context"
	"go-eth/pkg/eth"
	"go-eth/pkg/rpc/batcher"
	"go-eth/pkg/rpc/mapper"
)

func (ix *Indexer) fetchFullBlockWorker(ctx context.Context, inCh chan string, outCh chan eth.FullRawBlock, errCh chan error) {
	for {
		select {
		case blockHex, ok := <-inCh:
			if !ok {
				return
			}

			rawBlock, err := ix.fetchFullBlock(ctx, blockHex)

			if err != nil {
				select {
				case errCh <- err:
				case <-ctx.Done():
				}
			}

			select {
			case outCh <- rawBlock:
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
		}
	}
}

func (ix *Indexer) fetchFullBlock(ctx context.Context, blockHex string) (eth.FullRawBlock, error) {
	batchBuilder := batcher.New()

	blockResId := batchBuilder.Add("eth_getBlockByNumber", blockHex, true)
	reciptsResId := batchBuilder.Add("eth_getBlockReceipts", blockHex)

	resp, err := ix.client.BatchCall(ctx, batchBuilder.Request()...)

	if err != nil {
		return eth.FullRawBlock{}, err
	}

	mapper := mapper.New(resp)

	rawBlock, err := mapper.GetByID(blockResId)

	if err != nil {
		return eth.FullRawBlock{}, err
	}

	rawRecipts, err := mapper.GetByID(reciptsResId)

	if err != nil {
		return eth.FullRawBlock{}, err
	}

	if rawBlock.Error != nil {
		return eth.FullRawBlock{}, rawBlock.Error
	}

	if rawRecipts.Error != nil {
		return eth.FullRawBlock{}, rawRecipts.Error
	}

	return eth.FullRawBlock{
		Block:    rawBlock.Result,
		Receipts: rawRecipts.Result,
	}, nil
}

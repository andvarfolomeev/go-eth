package indexer

import (
	"context"
	"encoding/json"
	"go-eth/pkg/conv"
	"go-eth/pkg/eth"
	"go-eth/pkg/model"
)

func (ix *Indexer) deserializeFullRawBlockWorker(ctx context.Context, inCh chan eth.FullRawBlock, outCh chan model.FullBlock, errCh chan error) {
	for {
		select {
		case fullRawBlock := <-inCh:
			fullBlock, err := ix.deserializeFullRawBlock(fullRawBlock)

			if err != nil {
				select {
				case errCh <- err:
				case <-ctx.Done():
					return
				}
			}

			select {
			case outCh <- fullBlock:
			case <-ctx.Done():
				return
			}

		case <-ctx.Done():
			return
		}
	}
}

func (ix *Indexer) deserializeFullRawBlock(fullRawBlock eth.FullRawBlock) (model.FullBlock, error) {
	var (
		ethBlock   eth.BlockWithTxs
		ethRecipts []eth.Receipt
	)

	if err := json.Unmarshal(*fullRawBlock.Block, &ethBlock); err != nil {
		return model.FullBlock{}, err
	}

	block, err := conv.EthBlockToModel(&ethBlock)

	if err != nil {
		return model.FullBlock{}, err
	}

	if err := json.Unmarshal(*fullRawBlock.Receipts, &ethRecipts); err != nil {
		return model.FullBlock{}, err
	}

	recipts := make([]*model.Receipt, 0, len(ethRecipts))

	for _, ethReceipt := range ethRecipts {
		recipt, err := conv.EthReceiptToModel(&ethReceipt)

		if err != nil {
			return model.FullBlock{}, err
		}

		recipts = append(recipts, recipt)
	}

	txs := make([]*model.Tx, 0, len(ethBlock.Transactions))

	for _, ethTx := range ethBlock.Transactions {
		tx, err := conv.EthTxToModel(&ethTx)

		if err != nil {
			return model.FullBlock{}, err
		}

		txs = append(txs, tx)
	}

	return model.FullBlock{
		Block:   block,
		Txs:     txs,
		Recipts: recipts,
	}, nil
}

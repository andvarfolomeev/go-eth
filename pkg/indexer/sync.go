package indexer

import (
	"context"
	"go-eth/pkg/model"
	"log/slog"
)

func (ix *Indexer) Sync(ctx context.Context, errCh chan error) {
	earliestRawBlock, err := ix.fetchFullBlock(ctx, "earliest")

	if err != nil {
		errCh <- err
		return
	}

	latestRawBlock, err := ix.fetchFullBlock(ctx, "latest")

	if err != nil {
		errCh <- err
		return
	}

	earliestBlock, err := ix.deserializeFullRawBlock(earliestRawBlock)

	if err != nil {
		errCh <- err
		return
	}

	latestBlock, err := ix.deserializeFullRawBlock(latestRawBlock)

	if err != nil {
		errCh <- err
		return
	}

	err = ix.saveFullBlocks(ctx, []model.FullBlock{
		earliestBlock,
		latestBlock,
	})

	if err != nil {
		errCh <- err
		return
	}

	slog.Info("Successful sync")
}

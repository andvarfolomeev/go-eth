package indexer

import (
	"context"
	"database/sql"
	"go-eth/pkg/model"
	"go-eth/pkg/repositories"
	"log/slog"
	"time"
)

func (ix *Indexer) saveFullBlockWorker(ctx context.Context, fullBlockCh chan model.FullBlock, errCh chan error) {
	batchSize := 500
	batch := make([]model.FullBlock, 0, batchSize)

	tickerDuration := time.Second * 5
	ticker := time.NewTicker(tickerDuration)

	saveBatch := func() {
		if len(batch) == 0 {
			return
		}

		start := time.Now()
		err := ix.saveFullBlocks(ctx, batch)
		elapsed := time.Since(start).Milliseconds()

		if err != nil {
			errCh <- err
		}

		slog.Info("Saved batch", "len", len(batch), "elapsed (ms)", elapsed)

		batch = batch[:0]

	}

	for {
		select {
		case <-ticker.C:
			saveBatch()
		case fullBlock := <-fullBlockCh:
			batch = append(batch, fullBlock)

			if len(batch) <= batchSize {
				continue
			}

			saveBatch()
			ticker.Reset(tickerDuration)
		case <-ctx.Done():
			return
		}
	}
}

func (ix *Indexer) saveFullBlocks(ctx context.Context, fullBlocks []model.FullBlock) error {
	tx, err := ix.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})

	if err != nil {
		return err
	}

	blockRep := repositories.NewBlockRepository(tx)
	txRep := repositories.NewTxRepository(tx)
	receiptRep := repositories.NewReceiptRepository(tx)

	for _, fullBlock := range fullBlocks {
		err = blockRep.Save(ctx, fullBlock.Block)

		if err != nil {
			return err
		}

		for _, tx := range fullBlock.Txs {
			err = txRep.Save(ctx, tx)

			if err != nil {
				return err
			}
		}
		for _, receipt := range fullBlock.Recipts {
			err = receiptRep.Save(ctx, receipt)

			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

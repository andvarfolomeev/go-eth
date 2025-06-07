package indexer

import (
	"context"
	"go-eth/pkg/conv"
	"go-eth/pkg/repositories"
	"log/slog"
)

func (ix *Indexer) GapsGenerate(ctx context.Context, outCh chan string, errCh chan error) {
	blockRep := repositories.NewBlockRepository(ix.db)
	gaps, err := blockRep.GetGaps(ctx)

	if err != nil {
		errCh <- err
	}

	go func() {
		for _, gap := range gaps {
			for blockNumber := gap.Start; blockNumber < gap.End; blockNumber++ {
				blockHex := conv.Int64ToHex(blockNumber)
				select {
				case outCh <- blockHex:
				case <-ctx.Done():
					slog.Info("close")
					return
				}
			}
		}
	}()
}

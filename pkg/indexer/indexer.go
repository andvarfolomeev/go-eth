package indexer

import (
	"context"
	"database/sql"
	"sync"

	"go-eth/pkg/eth"
	"go-eth/pkg/model"
	"go-eth/pkg/rpc"
)

type Indexer struct {
	client *rpc.Client
	db     *sql.DB
}

func New(client *rpc.Client, db *sql.DB) *Indexer {
	return &Indexer{
		client: client,
		db:     db,
	}
}

func (ix *Indexer) Start(ctx context.Context, workerCount int, inCh chan string, errCh chan error) {
	ix.Sync(ctx, errCh)

	go ix.GapsGenerate(ctx, inCh, errCh)

	fullRawBlockCh := make(chan eth.FullRawBlock)
	domainFullBlockCh := make(chan model.FullBlock)

	ix.StartFetchFullBlockWorker(ctx, workerCount, inCh, fullRawBlockCh, errCh)
	ix.StartDeserializeFullRawBlockWorker(ctx, workerCount, fullRawBlockCh, domainFullBlockCh, errCh)

	go ix.saveFullBlockWorker(ctx, domainFullBlockCh, errCh)
}

func (ix *Indexer) StartFetchFullBlockWorker(ctx context.Context, workerCount int, inCh chan string, outCh chan eth.FullRawBlock, errCh chan error) {
	var wg sync.WaitGroup

	wg.Add(workerCount)

	for range workerCount {
		go func() {
			defer wg.Done()

			ix.fetchFullBlockWorker(ctx, inCh, outCh, errCh)
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()
}

func (ix *Indexer) StartDeserializeFullRawBlockWorker(ctx context.Context, workerCount int, inCh chan eth.FullRawBlock, outCh chan model.FullBlock, errCh chan error) {
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			ix.deserializeFullRawBlockWorker(ctx, inCh, outCh, errCh)
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()
}

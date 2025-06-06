package main

import (
	"context"

	"go-eth/internal/config"
	"go-eth/internal/jsonlogger"
	"go-eth/pkg/indexer"
	"go-eth/pkg/intrastructure"
	"go-eth/pkg/rpc"
)

func main() {
	options := config.ParseOptions()
	jsonlogger.SetupLoggerAsDefault(options.LogLevel)

	ctx := context.Background()
	defer ctx.Done()

	blockHexCh := make(chan string)
	defer close(blockHexCh)

	errCh := make(chan error)
	defer close(errCh)

	db := intrastructure.NewDB(options.DatabaseURL)
	client := rpc.New(options.RpcURL)

	// poller := poller.NewRPCBlockPoller(client, time.Second)
	// poller.Poll(ctx, blockHexCh, errCh)

	ix := indexer.NewIndexer(client, db)
	ix.Start(ctx, options.FetchWorkers, blockHexCh, errCh)

	for {
		select {
		case err := <-errCh:
			panic(err)
		case <-ctx.Done():
		}
	}
}

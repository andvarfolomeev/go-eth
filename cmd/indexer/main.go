package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"go-eth/internal/config"
	"go-eth/internal/jsonlogger"
	"go-eth/pkg/indexer"
	"go-eth/pkg/intrastructure"
	"go-eth/pkg/rpc"
)

func main() {
	options := config.ParseOptions()
	jsonlogger.SetupLoggerAsDefault(options.LogLevel)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	blockHexCh := make(chan string)
	defer close(blockHexCh)

	errCh := make(chan error)
	defer close(errCh)

	db := intrastructure.NewDB(options.DatabaseURL)
	client := rpc.New(options.RpcURL)

	// poller := poller.New(client, time.Second)
	// poller.Poll(ctx, blockHexCh, errCh)

	ix := indexer.New(client, db)
	ix.Start(ctx, options.FetchWorkers, blockHexCh, errCh)

	for {
		select {
		case err := <-errCh:
			panic(err)
		case <-ctx.Done():
			slog.Info("Stopping...")
			os.Exit(0)
		}
	}
}

package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/andvarfolomeev/go-eth/internal/jsonlogger"
	"github.com/andvarfolomeev/go-eth/pkg/ethrpc"
)

func main() {
	options := parseOptions()
	jsonlogger.SetupLoggerAsDefault(options.logLevel)

	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 1000*time.Millisecond)
	defer cancel()

	client := ethrpc.New("https://eth.llamarpc.com")
	defer client.Close()

	blockHex, err := client.BlockNumber(ctx)

	if err != nil {
		log.Panic(err)
	}

	block, err := client.BlockWithTxsByNumber(ctx, *blockHex)

	if err != nil {
		log.Panic(err)
	}

	slog.Info("Getted block", "block", block)

}

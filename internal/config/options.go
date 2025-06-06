package config

import (
	"flag"
	"fmt"
	"log/slog"
	"strings"
)

type ProgramOptions struct {
	LogLevel     slog.Level
	FetchWorkers int
	DatabaseURL  string
	RpcURL       string
}

type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

var allowedLevels = []LogLevel{LevelDebug, LevelInfo, LevelWarn, LevelError}

func (l *LogLevel) String() string {
	return string(*l)
}

func (l *LogLevel) Set(s string) error {
	s = strings.ToLower(s)
	for _, level := range allowedLevels {
		if s == string(level) {
			*l = level
			return nil
		}
	}
	return fmt.Errorf("invalid log level: %s", s)
}

func strToSlogLevel(s LogLevel) slog.Level {
	mapping := map[LogLevel]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"Error": slog.LevelError,
	}

	return mapping[s]
}

func ParseOptions() *ProgramOptions {
	var (
		databaseURL string
		rpcURL      string
		logLevelStr LogLevel = LevelInfo
	)
	flag.Var(&logLevelStr, "log-level", "log level(debug, info, warn, error)")
	fetchWorkers := flag.Int("fetch-workers", 5, "fetch worker counts")
	flag.StringVar(&databaseURL, "pg-url", "", "PostgreSQL URL")
	flag.StringVar(&rpcURL, "rpc-url", "https://ethereum-rpc.publicnode.com", "RPC URL")

	flag.Parse()

	return &ProgramOptions{
		LogLevel:     strToSlogLevel(logLevelStr),
		FetchWorkers: *fetchWorkers,
		DatabaseURL:  databaseURL,
		RpcURL:       rpcURL,
	}
}

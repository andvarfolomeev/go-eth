package main

import (
	"flag"
	"fmt"
	"log/slog"
	"strings"
)

type ProgramOptions struct {
	logLevel slog.Level
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

func parseOptions() *ProgramOptions {
	var logLevelStr LogLevel = LevelInfo
	flag.Var(&logLevelStr, "log-level", "log level(debug, info, warn, error)")

	flag.Parse()

	return &ProgramOptions{
		logLevel: strToSlogLevel(logLevelStr),
	}
}

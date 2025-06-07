package jsonlogger

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"time"
)

type JsonLogger struct {
	w     io.Writer
	level slog.Level
}

func New(w io.Writer, level slog.Level) *JsonLogger {
	return &JsonLogger{w: w, level: level}
}

func (h *JsonLogger) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

func (h *JsonLogger) Handle(ctx context.Context, r slog.Record) error {
	m := map[string]interface{}{
		"level": r.Level.String(),
		"time":  r.Time.Format(time.RFC3339),
		"msg":   r.Message,
	}

	r.Attrs(func(a slog.Attr) bool {
		m[a.Key] = a.Value.Any()
		return true
	})

	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	_, err = h.w.Write(append(b, '\n'))
	return err
}

func (h *JsonLogger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *JsonLogger) WithGroup(name string) slog.Handler {
	return h
}

func SetupLoggerAsDefault(logLevel slog.Level) {
	logger := slog.New(New(os.Stdout, logLevel))
	slog.SetDefault(logger)
}

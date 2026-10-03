package slogdiscard

import (
	"context"
	"log/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHnadler())
}

type DiscardHnadler struct{}

func NewDiscardHnadler() *DiscardHnadler {
	return &DiscardHnadler{}
}

func (h *DiscardHnadler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *DiscardHnadler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *DiscardHnadler) WithGroup(_ string) slog.Handler {
	return h
}

func (h *DiscardHnadler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

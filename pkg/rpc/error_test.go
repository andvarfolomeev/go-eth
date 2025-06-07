package rpc_test

import (
	"errors"
	"testing"

	"go-eth/pkg/rpc"
)

func TestRPCError_Error(t *testing.T) {
	err := &rpc.Error{
		Code:    -32601,
		Message: "method not found",
	}

	t.Run("implements error", func(t *testing.T) {
		var e error = err
		if e.Error() == "" {
			t.Fatal("expected non-empty error string")
		}
	})

	t.Run("string format", func(t *testing.T) {
		expected := "jsonrpc error: -32601: method not found"
		if err.Error() != expected {
			t.Errorf("unexpected error string: got %q, want %q", err.Error(), expected)
		}
	})

	t.Run("errors.As works", func(t *testing.T) {
		var rpcErr *rpc.Error
		if !errors.As(err, &rpcErr) {
			t.Errorf("errors.As failed to cast to *RPCError")
		}
	})
}

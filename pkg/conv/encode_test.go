package conv_test

import (
	"testing"

	"go-eth/pkg/conv"
)

func TestUint64ToHex(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected string
	}{
		{"zero", 0, "0x0"},
		{"small number", 26, "0x1a"},
		{"max uint64", ^uint64(0), "0xffffffffffffffff"},
		{"power of two", 1024, "0x400"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.Uint64ToHex(tt.input)
			if got != tt.expected {
				t.Errorf("Uint64ToHex(%d) = %s, want %s", tt.input, got, tt.expected)
			}
		})
	}
}

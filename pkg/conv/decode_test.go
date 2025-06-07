package conv_test

import (
	"testing"
	"time"

	"go-eth/pkg/conv"
)

func TestHexToUint64(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint64
		wantErr bool
	}{
		{"simple hex", "0xa", 10, false},
		{"without prefix", "a", 10, false},
		{"zero", "0x0", 0, false},
		{"invalid hex", "0xg", uint64(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.HexToUint64(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToUint64() error = %v, wantError = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HexToUint64() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestHexToBigInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{"simple hex", "0x1a", "26", false},
		{"zero", "0x0", "0", false},
		{"valid large hex", "0xffff", "65535", false},
		{"invalid hex", "zzz", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.HexToBigInt(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToBigInt() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.expected {
				t.Errorf("HexToBigInt() = %v, want = %v", got.String(), tt.expected)
			}
		})
	}
}

func TestHexToTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantErr  bool
	}{
		{"valid timestamp", "0x5f5e100", time.Unix(100000000, 0), false},
		{"zero timestamp", "0x0", time.Unix(0, 0), false},
		{"invalid hex", "not-a-hex", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := conv.HexToTime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToTime() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.expected) {
				t.Errorf("HexToTime() = %v, want = %v", got, tt.expected)
			}
		})
	}
}

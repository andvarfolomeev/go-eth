package ethconv

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func HexToUint64(s string) (uint64, error) {
	return strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 64)
}

func HexToBigInt(s string) (*big.Int, error) {
	n := new(big.Int)
	n, ok := n.SetString(strings.TrimPrefix(s, "0x"), 16)
	if !ok {
		return nil, fmt.Errorf("invalid big hex: %s", s)
	}
	return n, nil
}

func HexToTime(s string) (time.Time, error) {
	sec, err := HexToUint64(s)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(sec), 0), nil
}

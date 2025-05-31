package ethconv

import "fmt"

func Uint64ToHex(n uint64) string {
	return fmt.Sprintf("0x%x", n)
}

package conv

import "fmt"

func Uint64ToHex(n uint64) string {
	return fmt.Sprintf("0x%x", n)
}

func Int64ToHex(n int64) string {
	return fmt.Sprintf("0x%x", n)
}

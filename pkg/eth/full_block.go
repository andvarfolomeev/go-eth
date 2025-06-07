package eth

import "encoding/json"

type FullBlock struct {
	Block   *BlockWithTxs
	Recipts *[]Receipt
}

type FullRawBlock struct {
	Block    *json.RawMessage
	Receipts *json.RawMessage
}

type ByteCodeResult struct {
	Address string
	Code    []byte
}

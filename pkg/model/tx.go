package model

type Tx struct {
	Hash        []byte
	BlockHash   []byte
	BlockNumber int64

	To   []byte
	From []byte

	Nonce int64

	GasLimit             int64
	GasPrice             int64
	MaxFeePerGas         int64
	MaxPriorityFeePerGas int64

	Value string
	Input []byte

	Type int8
}

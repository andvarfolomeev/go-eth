package model

type Receipt struct {
	BlockHash       []byte
	BlockNumber     int64
	TransactionHash []byte

	GasUsed           int64
	CumulativeGasUsed int64
	EffectiveGasPrice int64

	ContractAddress *string // null → pointer

	Status int8
}

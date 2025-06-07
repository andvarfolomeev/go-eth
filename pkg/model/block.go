package model

type Block struct {
	Number     int64
	Hash       []byte
	ParentHash []byte

	StateRoot        []byte
	ReceiptsRoot     []byte
	TransactionsRoot []byte

	Difficulty int64

	GasLimit int64
	GasUsed  int64

	Timestamp int64

	Nonce   []byte
	MixHash []byte

	Sha3Uncles []byte

	ExtraData []byte
	LogsBloom []byte
	Size      int64

	Miner []byte
}

package eth

type Tx struct {
	Hash        string `json:"hash"`
	BlockHash   string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`

	To   string `json:"to"`
	From string `json:"from"`

	Nonce string `json:"nonce"`

	GasLimit             string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`

	Value string `json:"value"`
	Input string `json:"input"`

	Type string `json:"type"`

	R       string `json:"r"`
	S       string `json:"s"`
	V       string `json:"v"`
	YParity string `json:"yParity"`
}

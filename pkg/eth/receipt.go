package eth

type Receipt struct {
	BlockHash       string `json:"blockHash"`
	BlockNumber     string `json:"blockNumber"`
	TransactionHash string `json:"transactionHash"`

	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	EffectiveGasPrice string `json:"effectiveGasPrice"`

	ContractAddress *string `json:"contractAddress"` // null → pointer

	Logs   []Log  `json:"logs"`
	Status string `json:"status"`
}

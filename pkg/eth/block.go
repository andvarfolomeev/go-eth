package eth

type Block struct {
	Number     string   `json:"number"`
	Hash       string   `json:"hash"`
	ParentHash string   `json:"parentHash"`
	Uncles     []string `json:"uncles"`

	StateRoot        string `json:"stateRoot"`
	ReceiptsRoot     string `json:"receiptsRoot"`
	TransactionsRoot string `json:"transactionsRoot"`

	Difficulty string `json:"difficulty"`

	GasLimit string `json:"gasLimit"`
	GasUsed  string `json:"gasUsed"`

	Timestamp string `json:"timestamp"`

	Nonce   string `json:"nonce"`
	MixHash string `json:"mixHash"`

	Sha3Uncles string `json:"sha3Uncles"`

	ExtraData string `json:"extraData"`
	LogsBloom string `json:"logsBloom"`
	Size      string `json:"size"`

	Miner string `json:"miner"`

	Transactions []string `json:"transactions"`
}

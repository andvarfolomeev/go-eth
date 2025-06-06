package model

type FullBlock struct {
	Block   *Block
	Txs     []*Tx
	Recipts []*Receipt
}

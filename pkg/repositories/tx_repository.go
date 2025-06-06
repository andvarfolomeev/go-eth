package repositories

import (
	"context"
	"go-eth/pkg/model"
)

type TxRepository struct {
	execer Execer
}

func NewTxRepository(execer Execer) *TxRepository {
	return &TxRepository{
		execer: execer,
	}
}

func (r *TxRepository) Save(ctx context.Context, tx *model.Tx) error {
	query := `
		INSERT INTO txs (
			hash, block_hash, block_number, "to", "from", nonce,
			gas, gas_price, max_fee_per_gas, max_priority_fee_per_gas,
			value, input, type
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10,
			$11, $12, $13
		)
		ON CONFLICT (hash) DO NOTHING
	`

	_, err := r.execer.ExecContext(ctx, query,
		tx.Hash,
		tx.BlockHash,
		tx.BlockNumber,
		tx.To,
		tx.From,
		tx.Nonce,
		tx.GasLimit,
		tx.GasPrice,
		tx.MaxFeePerGas,
		tx.MaxPriorityFeePerGas,
		tx.Value,
		tx.Input,
		tx.Type,
	)

	return err
}

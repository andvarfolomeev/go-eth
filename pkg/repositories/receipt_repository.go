package repositories

import (
	"context"
	"go-eth/pkg/model"
)

type ReceiptRepository struct {
	execer Execer
}

func NewReceiptRepository(execer Execer) *ReceiptRepository {
	return &ReceiptRepository{
		execer: execer,
	}
}

func (r *ReceiptRepository) Save(ctx context.Context, receipt *model.Receipt) error {
	query := `
		INSERT INTO receipts (
			block_hash, block_number, transaction_hash,
			gas_used, cumulative_gas_used, effective_gas_price,
			contract_address, status
		) VALUES (
			$1, $2, $3,
			$4, $5, $6,
			$7, $8
		)
		ON CONFLICT (transaction_hash) DO NOTHING
	`

	_, err := r.execer.ExecContext(ctx, query,
		receipt.BlockHash,
		receipt.BlockNumber,
		receipt.TransactionHash,
		receipt.GasUsed,
		receipt.CumulativeGasUsed,
		receipt.EffectiveGasPrice,
		receipt.ContractAddress,
		receipt.Status,
	)

	return err
}

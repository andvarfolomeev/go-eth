package repositories

import (
	"context"
	"go-eth/pkg/model"
)

type BlockRepository struct {
	execer Execer
}

func NewBlockRepository(execer Execer) *BlockRepository {
	return &BlockRepository{
		execer: execer,
	}
}

func (r *BlockRepository) Save(ctx context.Context, blockEntry *model.Block) error {
	query := `
		INSERT INTO blocks (
			number, hash, parent_hash, state_root, receipts_root, transactions_root,
			difficulty, gas_limit, gas_used, timestamp, nonce, mix_hash,
			sha3_uncles, extra_data, logs_bloom, size, miner
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11, $12,
			$13, $14, $15, $16, $17
		)
		ON CONFLICT (number) DO NOTHING
	`

	_, err := r.execer.ExecContext(ctx, query,
		blockEntry.Number,
		blockEntry.Hash,
		blockEntry.ParentHash,
		blockEntry.StateRoot,
		blockEntry.ReceiptsRoot,
		blockEntry.TransactionsRoot,
		blockEntry.Difficulty,
		blockEntry.GasLimit,
		blockEntry.GasUsed,
		blockEntry.Timestamp,
		blockEntry.Nonce,
		blockEntry.MixHash,
		blockEntry.Sha3Uncles,
		blockEntry.ExtraData,
		blockEntry.LogsBloom,
		blockEntry.Size,
		blockEntry.Miner,
	)

	return err
}

func (r *BlockRepository) GetGaps(ctx context.Context) ([]model.Gap, error) {
	query := `
		SELECT number + 1 AS start, next_nr - 1 AS end
		FROM (
	  		SELECT number, LEAD(number) OVER (ORDER BY number) AS next_nr
			FROM blocks
		) AS heights_lead
		WHERE number + 1 <> next_nr
		ORDER BY number ASC;
	`

	rows, err := r.execer.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gaps := make([]model.Gap, 0)
	for rows.Next() {
		var gap model.Gap
		if err := rows.Scan(&gap.Start, &gap.End); err != nil {
			return nil, err
		}
		gaps = append(gaps, gap)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return gaps, nil

}

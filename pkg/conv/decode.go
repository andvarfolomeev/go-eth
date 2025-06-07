package conv

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"go-eth/pkg/eth"
	"go-eth/pkg/model"
)

func HexToUint64(s string) (uint64, error) {
	return strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 64)
}

func HexToInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseInt(strings.TrimPrefix(s, "0x"), 16, 64)
}

func HexToBigInt(s string) (*big.Int, error) {
	n := new(big.Int)
	n, ok := n.SetString(strings.TrimPrefix(s, "0x"), 16)
	if !ok {
		return nil, fmt.Errorf("invalid big hex: %s", s)
	}
	return n, nil
}

func HexToTime(s string) (time.Time, error) {
	sec, err := HexToUint64(s)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(sec), 0), nil
}

func HexToBytes(hexStr string) ([]byte, error) {
	cleaned := strings.TrimPrefix(hexStr, "0x")
	return hex.DecodeString(cleaned)
}

func EthBlockToModel(rawBlock *eth.BlockWithTxs) (*model.Block, error) {
	number, err := HexToInt64(rawBlock.Number)
	if err != nil {
		return nil, err
	}

	difficulty, err := HexToInt64(rawBlock.Difficulty)
	if err != nil {
		return nil, err
	}

	gasLimit, err := HexToInt64(rawBlock.GasLimit)
	if err != nil {
		return nil, err
	}

	gasUsed, err := HexToInt64(rawBlock.GasUsed)
	if err != nil {
		return nil, err
	}

	timestamp, err := HexToInt64(rawBlock.Timestamp)
	if err != nil {
		return nil, err
	}

	size, err := HexToInt64(rawBlock.Size)
	if err != nil {
		return nil, err
	}

	hash, err := HexToBytes(rawBlock.Hash)
	if err != nil {
		return nil, err
	}

	parentHash, err := HexToBytes(rawBlock.ParentHash)
	if err != nil {
		return nil, err
	}

	stateRoot, err := HexToBytes(rawBlock.StateRoot)
	if err != nil {
		return nil, err
	}

	receiptsRoot, err := HexToBytes(rawBlock.ReceiptsRoot)
	if err != nil {
		return nil, err
	}

	transactionsRoot, err := HexToBytes(rawBlock.TransactionsRoot)
	if err != nil {
		return nil, err
	}

	nonce, err := HexToBytes(rawBlock.Nonce)
	if err != nil {
		return nil, err
	}

	mixHash, err := HexToBytes(rawBlock.MixHash)
	if err != nil {
		return nil, err
	}

	sha3Uncles, err := HexToBytes(rawBlock.Sha3Uncles)
	if err != nil {
		return nil, err
	}

	extraData, err := HexToBytes(rawBlock.ExtraData)
	if err != nil {
		return nil, err
	}

	logsBloom, err := HexToBytes(rawBlock.LogsBloom)
	if err != nil {
		return nil, err
	}

	miner, err := HexToBytes(rawBlock.Miner)
	if err != nil {
		return nil, err
	}

	return &model.Block{
		Number:           number,
		Hash:             hash,
		ParentHash:       parentHash,
		StateRoot:        stateRoot,
		ReceiptsRoot:     receiptsRoot,
		TransactionsRoot: transactionsRoot,
		Difficulty:       difficulty,
		GasLimit:         gasLimit,
		GasUsed:          gasUsed,
		Timestamp:        timestamp,
		Nonce:            nonce,
		MixHash:          mixHash,
		Sha3Uncles:       sha3Uncles,
		ExtraData:        extraData,
		LogsBloom:        logsBloom,
		Size:             size,
		Miner:            miner,
	}, nil
}

func EthReceiptToModel(r *eth.Receipt) (*model.Receipt, error) {
	blockNumber, err := HexToInt64(r.BlockNumber)
	if err != nil {
		return nil, err
	}

	gasUsed, err := HexToInt64(r.GasUsed)
	if err != nil {
		return nil, err
	}

	cumulativeGasUsed, err := HexToInt64(r.CumulativeGasUsed)
	if err != nil {
		return nil, err
	}

	effectiveGasPrice, err := HexToInt64(r.EffectiveGasPrice)
	if err != nil {
		return nil, err
	}

	statusInt, err := HexToInt64(r.Status)
	if err != nil {
		return nil, err
	}
	status := int8(statusInt)

	blockHash, err := HexToBytes(r.BlockHash)
	if err != nil {
		return nil, err
	}

	txHash, err := HexToBytes(r.TransactionHash)
	if err != nil {
		return nil, err
	}

	return &model.Receipt{
		BlockHash:         blockHash,
		BlockNumber:       blockNumber,
		TransactionHash:   txHash,
		GasUsed:           gasUsed,
		CumulativeGasUsed: cumulativeGasUsed,
		EffectiveGasPrice: effectiveGasPrice,
		ContractAddress:   r.ContractAddress,
		Status:            status,
	}, nil
}

func EthTxToModel(t *eth.Tx) (*model.Tx, error) {
	blockNumber, err := HexToInt64(t.BlockNumber)
	if err != nil {
		return nil, err
	}

	gasLimit, err := HexToInt64(t.GasLimit)
	if err != nil {
		return nil, err
	}

	gasPrice, err := HexToInt64(t.GasPrice)
	if err != nil {
		return nil, err
	}

	maxFeePerGas, err := HexToInt64(t.MaxFeePerGas)
	if err != nil {
		return nil, err
	}

	maxPriorityFeePerGas, err := HexToInt64(t.MaxPriorityFeePerGas)
	if err != nil {
		return nil, err
	}

	txTypeInt, err := HexToInt64(t.Type)
	if err != nil {
		return nil, err
	}
	txType := int8(txTypeInt)

	hash, err := HexToBytes(t.Hash)
	if err != nil {
		return nil, err
	}

	blockHash, err := HexToBytes(t.BlockHash)
	if err != nil {
		return nil, err
	}

	to, err := HexToBytes(t.To)
	if err != nil {
		return nil, err
	}

	from, err := HexToBytes(t.From)
	if err != nil {
		return nil, err
	}

	nonce, err := HexToInt64(t.Nonce)
	if err != nil {
		return nil, err
	}

	return &model.Tx{
		Hash:                 hash,
		BlockHash:            blockHash,
		BlockNumber:          blockNumber,
		To:                   to,
		From:                 from,
		Nonce:                nonce,
		GasLimit:             gasLimit,
		GasPrice:             gasPrice,
		MaxFeePerGas:         maxFeePerGas,
		MaxPriorityFeePerGas: maxPriorityFeePerGas,
		Value:                t.Value,
		Input:                []byte(t.Input),
		Type:                 txType,
	}, nil
}

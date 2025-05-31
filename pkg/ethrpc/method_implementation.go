package ethrpc

import (
	"context"
)

func (c *Client) BlockNumber(ctx context.Context) (*string, error) {
	var result string
	err := c.Do(ctx, "eth_blockNumber", []interface{}{}, &result)
	return &result, err
}

func (c *Client) ChainId(ctx context.Context) (*string, error) {
	var result string
	err := c.Do(ctx, "eth_chainId", []interface{}{}, &result)
	return &result, err
}

func (c *Client) Balance(ctx context.Context, address, block string) (*string, error) {
	var result string
	err := c.Do(ctx, "eth_getBalance", []interface{}{address, block}, &result)
	return &result, err
}

func (c *Client) BlockByHash(ctx context.Context, blockHash string) (*Block, error) {
	result := Block{}
	err := c.Do(ctx, "eth_getBlockByHash", []interface{}{blockHash, false}, &result)
	return &result, err
}

func (c *Client) BlockWithTxsByHash(ctx context.Context, blockHash string) (*Block, error) {
	result := Block{}
	err := c.Do(ctx, "eth_getBlockByHash", []interface{}{blockHash, true}, &result)
	return &result, err
}

func (c *Client) BlockByNumber(ctx context.Context, blockHex string) (*Block, error) {
	var result Block
	err := c.Do(ctx, "eth_getBlockByNumber", []interface{}{blockHex, false}, &result)
	return &result, err
}

func (c *Client) BlockWithTxsByNumber(ctx context.Context, blockHex string) (*BlockWithTxs, error) {
	var result BlockWithTxs
	err := c.Do(ctx, "eth_getBlockByNumber", []interface{}{blockHex, true}, &result)
	return &result, err
}

func (c *Client) BlockReceipts(ctx context.Context, block string) (*[]TxReceipt, error) {
	var result []TxReceipt
	err := c.Do(ctx, "eth_getBlockByNumber", []interface{}{block}, &result)
	return &result, err
}

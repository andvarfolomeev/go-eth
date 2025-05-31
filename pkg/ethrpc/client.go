package ethrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Client struct {
	url    string
	client *http.Client
	id     int
}

func New(url string) *Client {
	return &Client{
		url:    url,
		client: &http.Client{},
		id:     1,
	}
}

func (c *Client) Close() {
	c.client.CloseIdleConnections()
}

func (c *Client) Do(ctx context.Context, method string, params []interface{}, result interface{}) error {
	rpcReqBody := RPCRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      c.id,
	}
	c.id++

	data, err := json.Marshal(rpcReqBody)
	if err != nil {
		return err
	}

	slog.Debug("JSON RPC Request", "body", &rpcReqBody)

	req, err := http.NewRequestWithContext(ctx, "POST", c.url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var rpcResp RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return err
	}

	slog.Debug("JSON RPC Response", "body", &rpcResp)

	if rpcResp.Error != nil {
		return rpcResp.Error
	}

	if result != nil {
		if err := json.Unmarshal(rpcResp.Result, result); err != nil {
			return err
		}
	}

	return nil
}

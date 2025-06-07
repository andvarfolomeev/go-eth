package rpc

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
}

func New(url string) *Client {
	return &Client{
		url:    url,
		client: &http.Client{},
	}
}

func (c *Client) Close() {
	c.client.CloseIdleConnections()
}

func (c *Client) Call(ctx context.Context, rpcReqBody Request) (Response, error) {
	var resBody Response
	err := c.doRequest(ctx, rpcReqBody, &resBody, "JSON RPC Request", "JSON RPC Response")
	if err != nil {
		return Response{}, err
	}
	if resBody.Error != nil {
		slog.Error("JSON RPC Error", "error", resBody.Error)
		return resBody, resBody.Error
	}
	return resBody, nil
}

func (c *Client) BatchCall(ctx context.Context, rpcReqBody ...Request) ([]Response, error) {
	var resBody []Response
	err := c.doRequest(ctx, rpcReqBody, &resBody, "Batched JSON RPC Request", "Batched JSON RPC Response")
	if err != nil {
		return nil, err
	}
	for _, resp := range resBody {
		if resp.Error != nil {
			slog.Error("JSON RPC Error in batch", "id", resp.Id, "error", resp.Error)
		}
	}
	return resBody, nil
}

func (c *Client) doRequest(ctx context.Context, reqBody any, resBody any, reqLogMsg, resLogMsg string) error {
	data, err := json.Marshal(reqBody)
	if err != nil {
		slog.Error("Failed to marshal request", "error", err)
		return err
	}

	slog.Debug(reqLogMsg, "body", reqBody)

	req, err := http.NewRequestWithContext(ctx, "POST", c.url, bytes.NewBuffer(data))
	if err != nil {
		slog.Error("Failed to create HTTP request", "error", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		slog.Error("HTTP request failed", "error", err)
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(resBody); err != nil {
		slog.Error("Failed to decode response", "error", err)
		return err
	}

	slog.Debug(resLogMsg, "body", resBody)
	return nil
}

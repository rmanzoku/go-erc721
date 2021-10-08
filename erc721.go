package erc721

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rmanzoku/go-erc721/bind"
)

type ERC721 struct {
	rpc     string
	address string
}

func NewERC721(rpc, address string) *ERC721 {
	return &ERC721{rpc, address}
}

type rpcRequest struct {
	JsonRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
	ID      uint              `json:"id"`
}

type rpcResponce struct {
	JsonRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      uint   `json:"id"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (e *ERC721) call(opts *bind.CallOpts, params []json.RawMessage) (string, error) {
	in := &rpcRequest{
		JsonRPC: "2.0",
		Method:  "eth_call",
		Params:  params,
		ID:      1010101,
	}
	input, err := json.Marshal(in)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(opts.Context, "POST", e.rpc, bytes.NewReader(input))
	if err != nil {
		return "", err
	}

	cli := new(http.Client)
	resp, err := cli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := new(rpcResponce)
	err = json.Unmarshal(b, res)
	if err != nil {
		return "", err
	}

	if res.Error != nil {
		return "", fmt.Errorf("code:%d message:%s", res.Error.Code, res.Error.Message)
	}

	return res.Result, nil
}

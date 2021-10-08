package erc721

import (
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/rmanzoku/go-erc721/bind"
)

func (e *ERC721) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	buf := make([]byte, 32)
	tokenId.FillBytes(buf)
	data := "0x" + "c87b56dd" + hex.EncodeToString(buf)

	in := struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}{
		To:   e.address,
		Data: data,
	}

	param, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	latest, _ := json.Marshal("latest")

	result, err := e.call(opts, []json.RawMessage{param, latest})
	if err != nil {
		return "", err
	}
	b, err := decodeString(result)
	if err != nil {
		return "", err
	}
	uri := string(b[64:]) // remove string header
	return uri, nil
}

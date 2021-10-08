package erc721_test

import (
	"os"
	"testing"

	"github.com/rmanzoku/go-erc721"
)

const (
	rpc     = "https://polygon-rpc.com/"
	address = "0xb862AeC93F0169249935f82FD98E6a494F53C287"
)

var e *erc721.ERC721

func TestMain(m *testing.M) {
	e = erc721.NewERC721(rpc, address)
	os.Exit(m.Run())
}

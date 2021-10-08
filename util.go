package erc721

import (
	"encoding/hex"
	"hash"
	"strings"

	"golang.org/x/crypto/sha3"
)

func decodeString(s string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(s, "0x"))
}

func encodeToString(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func keccak256Hex(data []byte) string {
	hash := keccak256(data)
	return encodeToString(hash)
}

func keccak256(data ...[]byte) []byte {
	type keccakState interface {
		hash.Hash
		Read([]byte) (int, error)
	}
	b := make([]byte, 32)
	d := sha3.NewLegacyKeccak256().(keccakState)
	for _, b := range data {
		d.Write(b)
	}
	_, _ = d.Read(b)
	return b
}

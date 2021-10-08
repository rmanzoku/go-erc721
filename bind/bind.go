package bind

import (
	"context"
	"math/big"
)

type CallOpts struct {
	Pending     bool            // Whether to operate on the pending state or the last known one
	From        string          // Optional the sender address, otherwise the first account is used
	BlockNumber *big.Int        // Optional the block number on which the call should be performed
	Context     context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

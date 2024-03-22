package arbitrum

import (
	"context"

	"github.com/Pond-International/arb-go-ethereum/arbitrum_types"
	"github.com/Pond-International/arb-go-ethereum/core"
	"github.com/Pond-International/arb-go-ethereum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}

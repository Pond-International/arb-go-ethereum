package arbitrum

import (
	"context"

	"github.com/Pond-International/arb-go-ethereum/common/hexutil"
	"github.com/Pond-International/arb-go-ethereum/core"
	"github.com/Pond-International/arb-go-ethereum/internal/ethapi"
	"github.com/Pond-International/arb-go-ethereum/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}

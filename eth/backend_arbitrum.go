package eth

import (
	"context"

	"github.com/Pond-International/arb-go-ethereum/core"
	"github.com/Pond-International/arb-go-ethereum/core/state"
	"github.com/Pond-International/arb-go-ethereum/core/types"
	"github.com/Pond-International/arb-go-ethereum/core/vm"
	"github.com/Pond-International/arb-go-ethereum/eth/tracers"
	"github.com/Pond-International/arb-go-ethereum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}

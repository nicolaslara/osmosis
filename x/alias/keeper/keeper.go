package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v10/x/alias/types"
)

type (
	Keeper struct {
		cdc      codec.Codec
		storeKey sdk.StoreKey
	}
)

// NewKeeper returns a new instance of the x/alias keeper
func NewKeeper(
	cdc codec.Codec,
	storeKey sdk.StoreKey,
) Keeper {

	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// Logger returns a logger for the x/alias module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetCreatorsPrefixStore returns the substore that contains the contract
// ToDo: This should probably be a param
func (k Keeper) GetContractStore(ctx sdk.Context) sdk.KVStore {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, []byte("Contract"))
}

func (k Keeper) Exec(ctx sdk.Context, sender sdk.AccAddress, as sdk.AccAddress, msgs []sdk.Msg) ([][]byte, error) {
	return [][]byte{{'g', 'o', 'l', 'a', 'n', 'g'}}, nil
}

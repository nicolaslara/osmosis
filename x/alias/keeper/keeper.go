package keeper

import (
	"fmt"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc        codec.Codec
		storeKey   sdk.StoreKey
		wasmKeeper wasmkeeper.Keeper
		router     baseapp.MsgServiceRouter
	}
)

// NewKeeper returns a new instance of the x/alias keeper
func NewKeeper(
	cdc codec.Codec,
	storeKey sdk.StoreKey,
	wasmKeeper wasmkeeper.Keeper,
	router baseapp.MsgServiceRouter,
) Keeper {

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		wasmKeeper: wasmKeeper,
		router:     router,
	}
}

// GetCreatorsPrefixStore returns the substore that contains the contract
// ToDo: This should probably be a param
func (k Keeper) GetContractStore(ctx sdk.Context) sdk.KVStore {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, []byte("Contract"))
}

func (k Keeper) Exec(ctx sdk.Context, sender sdk.AccAddress, as sdk.AccAddress, msg string) ([]byte, error) {
	// ToDo: Build auth message: "{sender: sender, msgs: [msg]}"
	fmt.Println("HERE")
	fmt.Println(msg)
	result, err := k.wasmKeeper.QuerySmart(ctx, as, []byte(msg))
	fmt.Println("HERE2")
	fmt.Println(result)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	// ToDo: Check result of the query and fail if it's not authorized

	// Experiment of how to send a custom message through the handler. If authorized, this should be the received message
	//testMsg := sdk.Msg(
	//	&banktypes.MsgSend{
	//		FromAddress: string(sender),
	//		ToAddress:   string(sender),
	//		Amount:      sdk.NewCoins(sdk.NewCoin("osmo", sdk.NewInt(1))),
	//	},
	//)
	//
	//handler := k.router.Handler(testMsg)
	//if handler == nil {
	//	return nil, sdkerrors.ErrUnknownRequest.Wrapf("unrecognized message route: %s", sdk.MsgTypeURL(testMsg))
	//}
	//
	//msgResp, err := handler(ctx, testMsg)
	//if err != nil {
	//	return nil, sdkerrors.Wrapf(err, "failed to execute message; message %v", msg)
	//}
	//
	//// Emit the data for testing
	//ctx.EventManager().EmitEvents(sdk.Events{
	//	sdk.NewEvent(
	//		types.TypeMsgExec,
	//		sdk.NewAttribute("data", string(msgResp.Data)),
	//	),
	//})
	//
	//// emit the events from the dispatched actions
	//events := msgResp.Events
	//sdkEvents := make([]sdk.Event, 0, len(events))
	//for i := 0; i < len(events); i++ {
	//	sdkEvents = append(sdkEvents, sdk.Event(events[i]))
	//}
	//ctx.EventManager().EmitEvents(sdkEvents)

	return result, nil
}

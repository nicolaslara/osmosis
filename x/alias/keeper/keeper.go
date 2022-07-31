package keeper

import (
	"encoding/json"
	"fmt"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	appparams "github.com/osmosis-labs/osmosis/v10/app/params"
	"github.com/osmosis-labs/osmosis/v10/x/alias/types"
)

type (
	Keeper struct {
		cdc            codec.Codec
		storeKey       sdk.StoreKey
		wasmKeeper     wasmkeeper.Keeper
		encodingConfig appparams.EncodingConfig
		router         baseapp.MsgServiceRouter
	}
)

// NewKeeper returns a new instance of the x/alias keeper
func NewKeeper(
	cdc codec.Codec,
	storeKey sdk.StoreKey,
	wasmKeeper wasmkeeper.Keeper,
	encodingConfig appparams.EncodingConfig,
	router baseapp.MsgServiceRouter,
) Keeper {

	return Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		wasmKeeper:     wasmKeeper,
		encodingConfig: encodingConfig,
		router:         router,
	}
}

// GetCreatorsPrefixStore returns the substore that contains the contract
// ToDo: This should probably be a param
func (k Keeper) GetContractStore(ctx sdk.Context) sdk.KVStore {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, []byte("Contract"))
}

func (k Keeper) Exec(ctx sdk.Context, sender sdk.AccAddress, msgType string, msgAsJson string, as sdk.AccAddress) ([]byte, error) {
	// ToDo: Currently using "as" as the contract addr. This should be extracted as config for this module.

	// Validate message type and message.
	msgMap := map[string]interface{}{}
	if err := json.Unmarshal([]byte(msgAsJson), &msgMap); err != nil {
		return nil, sdkerrors.Wrapf(err, "invalid message format (expected json); message %v", msgAsJson)
	}

	// ToDo: Validate msgType
	validatedType := msgType
	// Add message type to the original message
	msgMap["@type"] = validatedType

	// The full message includes the validated type
	validatedMsg, err := json.Marshal(msgMap)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "can't reconvert message to json; message %v", msgMap)
	}
	fmt.Println("validatedMsg: ", string(validatedMsg))

	registry := k.encodingConfig.InterfaceRegistry
	protoCdc := codec.NewProtoCodec(registry)

	var msg sdk.Msg
	err = protoCdc.UnmarshalInterfaceJSON([]byte(validatedMsg), &msg)
	fmt.Println("msg", msg)

	jsonMsg, err := protoCdc.MarshalInterfaceJSON(msg)
	if err != nil {
		return nil, err
	}

	fmt.Println("jsonMsg", string(jsonMsg))

	// ToDo: How do I convert msgTypes to the message format cosmwasm expects? (/cosmos.bank.v1beta1.MsgSend -> {"bank": "send": ...}})
	authMsg := fmt.Sprintf(`{"authorize": {"msgs": [{"bank": {"send": %s}}], "sender": "%s"}}`, jsonMsg, sender)
	fmt.Println(authMsg)
	response, err := k.wasmKeeper.QuerySmart(ctx, as, []byte(authMsg))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Failed to query authorization contract; query: %v", authMsg)
	}

	fmt.Println("Authorization Query response:", string(response))
	// Validate the query response
	responseMap := map[string]interface{}{}
	if err := json.Unmarshal(response, &responseMap); err != nil {
		return nil, sdkerrors.Wrapf(err, "invalid response format (expected json); response: %v", string(response))
	}

	// ToDo: Create a response type for this
	if !responseMap["authorized"].(bool) {
		// ToDo: Add events?
		return nil, sdkerrors.Wrapf(err, "The sender is not authorized to execute that message")
	}

	// // Consider using the same dispatcher as cosmwasm. This would need to be initialized with the keeper?
	//messager := wasmkeeper.NewDefaultMessageHandler(k.router, channelKeeper, capabilityKeeper, bankKeeper, cdc, portSource),

	handler := k.router.Handler(msg)
	if handler == nil {
		return nil, sdkerrors.ErrUnknownRequest.Wrapf("unrecognized message route: %s", sdk.MsgTypeURL(msg))
	}

	msgResponse, err := handler(ctx, msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to execute message; message %v", msg)
	}

	// Emit the data for testing
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeMsgExec,
			sdk.NewAttribute("data", string(msgResponse.Data)),
		),
	})

	// emit the events from the dispatched actions
	events := msgResponse.Events
	sdkEvents := make([]sdk.Event, 0, len(events))
	for i := 0; i < len(events); i++ {
		sdkEvents = append(sdkEvents, sdk.Event(events[i]))
	}
	ctx.EventManager().EmitEvents(sdkEvents)

	asBytes, err := msgResponse.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to marshal response; response %v", msgResponse)
	}

	return asBytes, nil
}

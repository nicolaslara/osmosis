package keeper

import (
	"encoding/json"
	"fmt"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
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

func (k Keeper) Exec(ctx sdk.Context, sender sdk.AccAddress, as sdk.AccAddress, msgAsJson string) ([]byte, error) {
	// ToDo: This should be done before getting here
	msgMap := map[string]interface{}{}
	if err := json.Unmarshal([]byte(msgAsJson), &msgMap); err != nil {
		return nil, sdkerrors.Wrapf(err, "invalid message format (expected json); message %v", msgAsJson)
	}
	validatedMsg, err := json.Marshal(msgMap)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "can't reconvert message to json; message %v", msgMap)
	}
	fmt.Println(string(validatedMsg))
	//err := wasmtypes.IsJSONObjectWithTopLevelKey([]byte(msgAsJson), [])

	authMsg := fmt.Sprintf("{\"authorize\": {\"msgs\": [%s], \"sender\": \"%s\"}}", validatedMsg, sender)
	fmt.Println(authMsg)
	result, err := k.wasmKeeper.QuerySmart(ctx, as, []byte(authMsg))

	if err != nil {
		return nil, err
	}

	// ToDo: Check result of the query and fail if it's not authorized
	fmt.Println(string(result))
	resultMap := map[string]interface{}{}
	if err := json.Unmarshal(result, &resultMap); err != nil {
		return nil, sdkerrors.Wrapf(err, "invalid result format (expected json); result %v", string(result))
	}

	// ToDo: Create an interface or type for this
	if !resultMap["authorized"].(bool) {
		// ToDo: Add events?
		return nil, sdkerrors.Wrapf(err, "The sender is not authorized to execute that message")
	}

	// Experiment of how to send a custom message through the handler. If authorized, this should be the received message

	//// This doesn't work because the value needs to be proto-encoded
	//testMsg := `{
	//	"type_url":"cosmos-sdk/MsgSubmitProposal",
	//	"value":{"content":{"type":"wasm/MigrateContractProposal","value":{"msg":{"foo":"bar"}}},"initial_deposit":[]}
	//}`
	//
	//var stargateMsg *wasmtypes.StargateMsg
	//if err := json.Unmarshal([]byte(testMsg), &stargateMsg); err != nil {
	//	return nil, err
	//}
	//
	//encoder := wasmkeeper.EncodeStargateMsg(types.ModuleCdc)
	//msgs, err := encoder(sender, stargateMsg)
	//if err != nil {
	//	return nil, err
	//}

	//
	//testMsg := sdk.Msg(
	//	&banktypes.MsgSend{
	//		FromAddress: sender.String(),
	//		ToAddress:   sender.String(),
	//		Amount:      sdk.NewCoins(sdk.NewCoin("uosmo", sdk.NewInt(1))),
	//	},
	//)

	var anyMsg cdctypes.Any
	testMsg2 := fmt.Sprintf(
		`{"@type":"/cosmos.bank.v1beta1.MsgSend",
	"from_address":"%s",
	"to_address":"%s",
	"amount":[{"denom":"uosmo","amount":"1"}]}`, sender, sender)
	if err := k.cdc.UnmarshalJSON([]byte(testMsg2), &anyMsg); err != nil {
		return nil, sdkerrors.Wrapf(err, "invalid result format (expected json); result %v", string(result))
	}

	fmt.Println("AnyMsg", anyMsg)

	//registry := cdctypes.NewInterfaceRegistry()
	//registry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	//registry := k.encodingConfig.InterfaceRegistry
	//protoCdc := codec.NewProtoCodec(registry)
	//
	//testMsg3 := sdk.Msg(&testdata.TestMsg{Signers: []string{"addr"}})
	//bz, err := protoCdc.MarshalInterface(testMsg3)
	//var msg2 sdk.Msg
	//err = protoCdc.UnmarshalInterface(bz, &msg2)

	registry := k.encodingConfig.InterfaceRegistry
	protoCdc := codec.NewProtoCodec(registry)
	var msg2 sdk.Msg
	err = protoCdc.UnmarshalInterface(anyMsg.Value, &msg2)

	fmt.Println("AnyMsg2", msg2)

	testMsg := fmt.Sprintf(
		`{"body":
{"messages":[
{"@type":"/cosmos.bank.v1beta1.MsgSend",
"from_address":"%s",
"to_address":"%s",
"amount":[{"denom":"uosmo","amount":"1"}]}
],
"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},
"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},
"signatures":[]}`, sender, sender)

	//encodingConfig := app.MakeEncodingConfig()
	// This works. Now just need to replace the decoder with my own
	tx, err := k.encodingConfig.TxConfig.TxJSONDecoder()([]byte(testMsg))

	if err != nil {
		return nil, err
	}

	msg := tx.GetMsgs()[0]
	fmt.Println(msg)

	bz, err := protoCdc.MarshalInterface(msg)
	fmt.Println("HERE", string(bz))
	var bz2 sdk.Msg
	err = protoCdc.UnmarshalInterface(bz, &bz2)
	fmt.Println("HERE", bz2)

	// // Consider using the same dispatcher as cosmwasm. This would need to be initialized with the keeper
	//messager := wasmkeeper.NewDefaultMessageHandler(k.router, channelKeeper, capabilityKeeper, bankKeeper, cdc, portSource),

	handler := k.router.Handler(msg)
	if handler == nil {
		return nil, sdkerrors.ErrUnknownRequest.Wrapf("unrecognized message route: %s", sdk.MsgTypeURL(msg))
	}

	msgResp, err := handler(ctx, msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to execute message; message %v", msg)
	}

	// Emit the data for testing
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeMsgExec,
			sdk.NewAttribute("data", string(msgResp.Data)),
		),
	})

	// emit the events from the dispatched actions
	events := msgResp.Events
	sdkEvents := make([]sdk.Event, 0, len(events))
	for i := 0; i < len(events); i++ {
		sdkEvents = append(sdkEvents, sdk.Event(events[i]))
	}
	ctx.EventManager().EmitEvents(sdkEvents)

	return result, nil
}

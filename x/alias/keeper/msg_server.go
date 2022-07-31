package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/v10/x/alias/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) Execute(goCtx context.Context, msg *types.MsgExec) (*types.MsgExecResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	as, err := sdk.AccAddressFromBech32(msg.As)
	if err != nil {
		return nil, err
	}

	// ToDo: Should we validate the message here or should we let it fail later?
	// If the message is not valid, it should fail either within cosmwasm, or
	// when sending decosing it before sending it if it were to be approved by the contract.
	// Though maybe just a bit of validation in ValidateBasic is in order (checking it's valid json)

	results, err := server.Keeper.Exec(ctx, sender, msg.MsgType, msg.Msg, as)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeMsgExec,
			sdk.NewAttribute(types.AttributeSender, msg.Sender),
			//sdk.NewAttribute("fullMsgs", string(msg.Msgs[0])),
			sdk.NewAttribute(types.AttributeResponse, string(results[:])),
		),
	})

	return &types.MsgExecResponse{results}, nil
}

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

	fullMsgs, err := msg.GetMessages()
	if err != nil {
		return nil, err
	}

	results, err := server.Keeper.Exec(ctx, sender, as, fullMsgs)
	if err != nil {
		return nil, err
	}

	return &types.MsgExecResponse{results}, nil
}

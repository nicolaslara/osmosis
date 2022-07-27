package keeper_test

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/v10/app"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"io/ioutil"
	"testing"
	"time"

	"github.com/osmosis-labs/osmosis/v10/app/apptesting"
	"github.com/osmosis-labs/osmosis/v10/x/alias/keeper"
	"github.com/osmosis-labs/osmosis/v10/x/alias/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	msgServer types.MsgServer
	contract  sdk.AccAddress
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

// Based on wasmbindings/test/custom_query_test.go
func SetupCustomApp(suite *KeeperTestSuite, addr sdk.AccAddress) (*app.OsmosisApp, sdk.AccAddress) {
	osmosis := app.Setup(false)
	suite.Ctx = osmosis.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "osmosis-1", Time: time.Now().UTC()})

	wasmKeeper := osmosis.WasmKeeper

	storeAuthCode(suite, osmosis, addr)

	cInfo := wasmKeeper.GetCodeInfo(suite.Ctx, 1)
	require.NotNil(suite.T(), cInfo)

	auth := instantiateAuthContract(suite, osmosis, addr)
	require.NotEmpty(suite.T(), auth)

	return osmosis, auth
}

func storeAuthCode(suite *KeeperTestSuite, osmosis *app.OsmosisApp, addr sdk.AccAddress) {
	govKeeper := osmosis.GovKeeper
	wasmCode, err := ioutil.ReadFile("../testdata/alias.wasm")
	require.NoError(suite.T(), err)

	src := wasmtypes.StoreCodeProposalFixture(func(p *wasmtypes.StoreCodeProposal) {
		p.RunAs = addr.String()
		p.WASMByteCode = wasmCode
	})

	// when stored
	storedProposal, err := govKeeper.SubmitProposal(suite.Ctx, src, false)
	require.NoError(suite.T(), err)

	// and proposal execute
	handler := govKeeper.Router().GetRoute(storedProposal.ProposalRoute())
	err = handler(suite.Ctx, storedProposal.GetContent())
	require.NoError(suite.T(), err)
}

func instantiateAuthContract(suite *KeeperTestSuite, osmosis *app.OsmosisApp, funder sdk.AccAddress) sdk.AccAddress {
	initMsgBz := []byte("{}")
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(osmosis.WasmKeeper)
	codeID := uint64(1)
	addr, _, err := contractKeeper.Instantiate(suite.Ctx, codeID, funder, funder, initMsgBz, "auth contract", nil)
	require.NoError(suite.T(), err)

	return addr
}

// The actual tests
func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.msgServer = keeper.NewMsgServerImpl(*suite.App.AliasKeeper)
	_, addr := SetupCustomApp(suite, suite.TestAccs[0])
	suite.contract = addr
}

func (suite *KeeperTestSuite) TestSimpleMsgExec() {
	suite.T().Log("TestSimpleMsgExec start")
	suite.T().Log("Instantiating contract")

	suite.T().Log("Running the test")

	// tmp message to match the interface. The sender will be set by the chain
	execMsg := types.NewMsgExec(suite.TestAccs[0].String(), suite.contract.String(), "{\"authorize\": {\"msgs\": [], \"sender\": \"osmo111111111111111111111111111111111111111\"}}")
	suite.T().Log("Created message")
	res, err := suite.msgServer.Execute(sdk.WrapSDKContext(suite.Ctx), execMsg)
	suite.T().Log("received response")
	suite.T().Log(res, err)
	suite.T().Log("TestSimpleMsgExec end")
	//suite.defaultDenom = res.GetNewTokenDenom()
}

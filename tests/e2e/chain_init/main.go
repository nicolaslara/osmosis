package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/osmosis-labs/osmosis/v7/tests/e2e/chain"
)

func main() {
	var (
		dataDir                string
		chainId                string
		numVal                 int
		pruning                string
		pruningKeepRecent      string
		pruningInterval        string
		snapshotInterval       uint64
		snapshotKeepRecent     uint
		telemetryEnabled       bool
		telemetryRetentionTime int64
		prometheusEnabled      bool
	)

	flag.StringVar(&dataDir, "data-dir", "", "chain data directory")
	flag.StringVar(&chainId, "chain-id", "", "chain ID")
	flag.IntVar(&numVal, "num-val", 2, "number of validators for each chain")
	flag.StringVar(&pruning, "pruning", "default", "pruning config")
	flag.StringVar(&pruningKeepRecent, "pruning-keep-recent", "0", "pruning keep recent config")
	flag.StringVar(&pruningInterval, "pruning-interval", "0", "pruning interval config")
	flag.Uint64Var(&snapshotInterval, "snapshot-interval", 1500, "state sync snapshot interval")
	flag.UintVar(&snapshotKeepRecent, "snapshot-keep-recent", 2, "number of state sync snapshots to keep and serve")
	flag.BoolVar(&telemetryEnabled, "telemetry-enabled", false, "enable telemetry in app.toml")
	flag.Int64Var(&telemetryRetentionTime, "telemetry-retention-time", 0, "retention time for telemetry in app.toml")
	flag.BoolVar(&prometheusEnabled, "prometheus-enabled", false, "enable prometheus in config.toml")
	flag.Parse()

	if len(dataDir) == 0 {
		panic("data-dir is required")
	}

	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		panic(err)
	}

	createdChain, err := chain.Init(chainId, dataDir, numVal, pruning, pruningKeepRecent, pruningInterval, snapshotInterval, uint32(snapshotKeepRecent), telemetryEnabled, telemetryRetentionTime, prometheusEnabled)
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(createdChain)
	fileName := fmt.Sprintf("%v/%v-encode", dataDir, chainId)
	if err = os.WriteFile(fileName, b, 0o777); err != nil {
		panic(err)
	}
}

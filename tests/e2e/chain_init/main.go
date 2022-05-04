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
		dataDir string
		chainId string
		numVal  int
	)

	flag.StringVar(&dataDir, "data-dir", "", "chain data directory")
	flag.StringVar(&chainId, "chain-id", "", "chain ID")
	flag.IntVar(&numVal, "num-val", 2, "number of validators for each chain")
	flag.Parse()

	if len(dataDir) == 0 {
		panic("data-dir is required")
	}

	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		panic(err)
	}

	createdChain, err := chain.Init(chainId, dataDir, numVal)
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(createdChain)
	fileName := fmt.Sprintf("%v/%v-encode", dataDir, chainId)
	if err = os.WriteFile(fileName, b, 0o777); err != nil {
		panic(err)
	}
}

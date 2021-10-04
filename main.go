package main

import (
	"fmt"
	"log"

	"github.com/aaronwinter/replay-celo-bug/counter"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/ethclient"
)

const RPC_ALFAJORES = "https://alfajores-forno.celo-testnet.org"
const CONTRACT = "0x8d3d74FB54780eA6c75e3fd9e79413dc71003A92"

func main() {
	contractAddress := common.HexToAddress(CONTRACT)
	alfajores, err := ethclient.Dial(RPC_ALFAJORES)
	if err != nil {
		log.Fatal(err.Error())
	}

	counter, err := counter.NewCounter(contractAddress, alfajores)
	if err != nil {
		log.Fatal(err.Error())
	}

	value, err := counter.Value(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Counter value:", value.String())
}

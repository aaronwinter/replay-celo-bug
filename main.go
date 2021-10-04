package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/aaronwinter/celo-blockchain/accounts/abi/bind_v2"
	"github.com/aaronwinter/celo-blockchain/common"
	"github.com/aaronwinter/celo-blockchain/ethclient"
	"github.com/aaronwinter/replay-celo-bug/counter"
)

// generate a new account using `geth account new`
// fund this account with alfajores celo
// paste path to keystore file here
const PATH_TO_SK = "/home/user/.celo/keystore/UTC--2021-10-04T14-10-50.370759409Z--6d214a0d085965b19f7e9ace6adc550094bdcaab"
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

	// Creating a keyed transactor to increment the counter (state changes)
	f, err := os.Open(PATH_TO_SK)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Fix with EIP155 signer:
	chainId := new(big.Int).SetInt64(44787)
	transactOpts, err := bind_v2.NewTransactorWithChainId(f, "", chainId)
	// transactOpts, err := bind_v2.NewTransactor(f, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	txObj := counter.Increment(transactOpts)
	txProm, err := txObj.Send()
	if err != nil {
		log.Fatal(err.Error())
	}

	txReceipt, err := txProm.WaitMined(context.TODO())
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("receipt - tx hash:", txReceipt.TxHash.Hex())
}

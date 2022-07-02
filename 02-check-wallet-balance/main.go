package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	// create our ethclient
	client, err := ethclient.DialContext(context.Background(), os.Getenv("INFURA_PROJECT_ENDPOINT"))
	if err != nil {
		log.Fatalf("Error creating Rinkeby ethclient: %v", err)
	}
	defer client.Close()

	addrToCheck := "0xcDA1048cf97B65ED9fb852AE677F02a28bd09ad3"
	checksumAddr := common.HexToAddress(addrToCheck)

	// use client.BalanceAt to check the balance of an Ethereum address
	balanceWei, err := client.BalanceAt(context.Background(), checksumAddr, nil)
	if err != nil {
		log.Fatalf("Error getting balance of address: %v", err)
	}

	fmt.Println(
		"Balance at",
		checksumAddr,
		"is",
		balanceWei,
		"Wei\n\tor",
		weiToEth(balanceWei),
		"Ether",
	)
}

func weiToEth(weiAmt *big.Int) *big.Float {
	weiAmt_float := new(big.Float)
	weiAmt_float.SetInt(weiAmt)

	// 1 Ether == 10^18 Wei
	ethAmt := new(big.Float).Quo(weiAmt_float, big.NewFloat(math.Pow10(18)))

	return ethAmt
}

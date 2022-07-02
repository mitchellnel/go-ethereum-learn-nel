package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	interactWithRinkeby()

	interactWithGanache()
}

func interactWithRinkeby() {
	// create our ethclient
	client, err := ethclient.DialContext(context.Background(), os.Getenv("INFURA_PROJECT_ENDPOINT"))
	if err != nil {
		log.Fatalf("Error creating Rinkeby ethclient: %v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(
		context.Background(),
		nil,
	) // pass nil to get current/most recent block
	if err != nil {
		log.Fatalf("Error getting current Rinkeby block: %v", err)
	}
	fmt.Println(block.Number())
}

func interactWithGanache() {
	// create out ethclient
	client, err := ethclient.DialContext(context.Background(), os.Getenv("GANACHE_CLI_ENDPOINT"))
	if err != nil {
		log.Fatalf("Error creating Ganache ethclient: %v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting current Ganache block: %v", err)
	}
	fmt.Println(block.Number())
}

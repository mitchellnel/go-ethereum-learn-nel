package main

import (
	"context"
	"crypto/ecdsa"
	todo "example/go-ethereum-learn-nel/todo-list-smart-contract/gen"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

/*
To build a binary file and the ABI from our smart contract, we run:
    $ solc --bin --abi contracts/Todo.sol -o build

Then, we use abigen to generate a .go file:
    $ abigen --bin=build/Todo.bin --abi=build/Todo.abi --pkg=todo --out=gen/todo.go
*/
func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	// create ethclient
	client, err := ethclient.Dial(os.Getenv("INFURA_PROJECT_ENDPOINT"))
	if err != nil {
		log.Fatalf("Error creating ethclient: %v", err)
	}
	defer client.Close()

	pvk, _, pubAddr := parseKeystore(
		"wallet/UTC--2022-07-02T16-44-15.195452000Z--d8eaef6563d22583bf98759088d3044c16c864c1",
	)

	// deploy the Todo contract
	deployContract(client, pvk, &pubAddr)
}

func parseKeystore(keystoreFile string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address) {
	// get key from wallet
	keystoreBytes, err := ioutil.ReadFile(keystoreFile)
	if err != nil {
		log.Fatalf("Error reading keystore file: %v", err)
	}

	key, err := keystore.DecryptKey(keystoreBytes, "password3")
	if err != nil {
		log.Fatalf("Error decrypting key: %v", err)
	}

	// derive key types
	pvk := key.PrivateKey
	pbk := pvk.PublicKey

	// derive public address
	pubAddr := crypto.PubkeyToAddress(pbk)

	return pvk, &pbk, pubAddr
}

func deployContract(client *ethclient.Client, pvk *ecdsa.PrivateKey, ownerAddress *common.Address) {
	// prepare deployment txn data
	// we need: nonce, gas limit, gas price, and chain ID
	nonce_uint64, err := client.PendingNonceAt(context.Background(), *ownerAddress)
	if err != nil {
		log.Fatalf("Error getting nonce for %v: %v", ownerAddress.Hex(), err)
	}
	var nonce big.Int
	nonce.SetUint64(nonce_uint64)

	var gasLimit uint64 = 3_000_000

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error getting suggested gas price: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error getting chain ID from ethclient: %v", err)
	}

	// create transaction options object
	auth, err := bind.NewKeyedTransactorWithChainID(pvk, chainID)
	if err != nil {
		log.Fatalf("Error creating transaction signer: %v", err)
	}

	auth.From = *ownerAddress
	auth.Nonce = &nonce
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	// use generated code from abigen in todo.go to deploy the contract
	contractAddr, txn, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatalf("Error deploying Todo contract: %v", err)
	} else {
		log.Printf(
			"Successfully deployed Todo contract -- transaction hash is %v",
			txn.Hash().Hex(),
		)
		log.Printf("Todo contract address is %v\n", common.HexToAddress(contractAddr.Hex()))
	}
}

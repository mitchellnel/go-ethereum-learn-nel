package main

import (
	"context"
	"crypto/ecdsa"
	todo "example/go-ethereum-learn-nel/todo-list-smart-contract/gen"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	contractAddr, _, _ := deployContract(client, pvk, &pubAddr)

	// interact with the todo contract
	// we could already pass the contract object, but to demonstrate creating an instance of the
	// object bound to an on-chain contract, we'll just pass the address
	interactWithContract(client, pvk, pubAddr, contractAddr)
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

func deployContract(
	client *ethclient.Client,
	pvk *ecdsa.PrivateKey,
	ownerAddress *common.Address,
) (common.Address, types.Transaction, todo.Todo) {
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
		log.Fatalf("Error creating transaction signer and options: %v", err)
	}

	auth.From = *ownerAddress
	auth.Nonce = &nonce
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	// use generated code from abigen in todo.go to deploy the contract
	contractAddr, txn, contract, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatalf("Error deploying Todo contract: %v", err)
	} else {
		log.Printf(
			"Successfully deployed Todo contract -- transaction hash is %v",
			txn.Hash().Hex(),
		)
		log.Printf("Todo contract address is %v\n", common.HexToAddress(contractAddr.Hex()))
	}

	// wait for deployment transaction to be mined
	_, err = bind.WaitMined(context.Background(), client, txn)
	if err != nil {
		log.Fatalf("Error mining Todo deployment transaction: %v", err)
	}

	return contractAddr, *txn, *contract
}

func interactWithContract(
	client *ethclient.Client,
	pvk *ecdsa.PrivateKey,
	pubAddr common.Address,
	contractAddr common.Address,
) {
	// sanitise addresses
	pubAddr = common.HexToAddress(pubAddr.Hex())
	contractAddr = common.HexToAddress(contractAddr.Hex())

	// create an instance of the Todo contract, bound to the on-chain contract

	contract, err := todo.NewTodo(contractAddr, client)
	if err != nil {
		log.Fatalf("Error creating an instance of the Todo contract: %v", err)
	}

	// then, we can access the functions of our smart contract by calling contract.FunctionName
	// e.g. contract.AddTask

	// suppose we wanted to add a task by calling AddTask
	// we need to provide contract.AddTask with: transaction options, and the actual parameter the
	// smart contract function needs to be passed, i.e. _content

	// prepare transcation options
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error getting chaind ID from ethclient: %v", err)
	}

	var gasLimit uint64 = 3_000_000

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error getting suggested gas price: %v", err)
	}

	txnOpts_AddTask, err := bind.NewKeyedTransactorWithChainID(pvk, chainID)
	if err != nil {
		log.Fatalf("Error creating transaction signer and options: %v", err)
	}
	txnOpts_AddTask.GasLimit = gasLimit
	txnOpts_AddTask.GasPrice = gasPrice

	// then call the smart contract function that we want to execute
	taskName := "Learn about go-ethereum"

	txn_AddTask, err := contract.AddTask(txnOpts_AddTask, taskName)
	if err != nil {
		log.Fatalf("Error creating transaction for AddTask(): %v", err)
	} else {
		log.Printf(
			"Successfully made transaction for AddTask() -- transaction has is %v",
			txn_AddTask.Hash().Hex(),
		)
	}

	// wait for transaction to be mined
	_, err = bind.WaitMined(context.Background(), client, txn_AddTask)
	if err != nil {
		log.Fatalf("Error mining AddTask() transaction: %v", err)
	}

	// then, suppose we wanted to get a list of our todo tasks by calling ListTasks
	// note that in ListTasks, we are not changing the state of the blockchain -- just reading it
	// this means that we will make a **call** not a transaction

	// so, we need not sign any transactions, we just make a call to the ListTasks function
	// (passing some caller options)
	tasks, err := contract.ListTasks(&bind.CallOpts{
		From: pubAddr,
	})
	if err != nil {
		log.Fatalf("Error making call to ListTasks(): %v", err)
	} else {
		log.Printf("Successfully made call to ListTasks()")
	}

	fmt.Println("Current Tasks:\n\t", tasks)
}

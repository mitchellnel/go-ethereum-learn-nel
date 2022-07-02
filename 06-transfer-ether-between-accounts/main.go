package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	var accounts []string

	numAccounts := 2
	password := "password3"

	if checkForExistingWallets("wallet", numAccounts) {
		accounts = getExistingAccounts("wallet", numAccounts, password)
		log.Println("Found existing wallets:", accounts)
	} else {
		accounts = createNewAccounts("wallet", numAccounts, password)
		log.Println("Created new wallets:", accounts)
	}

	// create ethclient
	client, err := ethclient.Dial(os.Getenv("INFURA_PROJECT_ENDPOINT"))
	if err != nil {
		log.Fatalf("Error creating ethclient: %v", err)
	}
	defer client.Close()

	// get current balances for all our accounts
	for _, account := range accounts {
		ethAddr := common.HexToAddress(account)
		balance, err := client.BalanceAt(context.Background(), ethAddr, nil)
		if err != nil {
			log.Fatalf("Error getting balance of %v: %v", account, err)
		}
		fmt.Println("Balance of", ethAddr, "is", weiToEth(balance), "Ether")
	}

	// can then go to a Rinkeby faucet and request Rinkeby ETH
	// then run the script
	// (I used https://rinkebyfaucet.com/)

	// now we want to transfer 0.01 ETH from accounts[0] to accounts[1]

	// prepare txn data
	toEthAddr := common.HexToAddress(accounts[1])

	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(accounts[0]))
	if err != nil {
		log.Fatalf("Error getting nonce for account %v: %v", accounts[0], err)
	}

	valueInEth := new(big.Float)
	valueInEth.SetFloat64(0.01)
	transactionValue := ethToWei(valueInEth)

	var gasLimit uint64 = 21000

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error getting suggested gas price: %v", err)
	}

	data := []byte{}

	// generate transaction
	txn := types.NewTx(&types.LegacyTx{
		To:       &toEthAddr,
		Nonce:    nonce,
		Value:    transactionValue,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	// sign transaction
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error getting chain ID from ethclient: %v", err)
	}

	// I can't be bothered programmatically importing private key for accounts[0]
	senderKeystorebytes, err := ioutil.ReadFile(
		filepath.Join(
			"wallet",
			"UTC--2022-07-02T16-44-15.195452000Z--d8eaef6563d22583bf98759088d3044c16c864c1",
		),
	)
	if err != nil {
		log.Fatalf("Error reading keystore file for accounts[0]: %v", err)
	}

	key, err := keystore.DecryptKey(senderKeystorebytes, password)
	if err != nil {
		log.Fatalf("Error decrypting key for accounts[0]: %v", err)
	}
	pvk := key.PrivateKey

	signedTxn, err := types.SignTx(txn, types.NewEIP155Signer(chainID), pvk)
	if err != nil {
		log.Fatalf("Error signing transaciton: %v", err)
	}

	// send the transaction to the network
	err = client.SendTransaction(context.Background(), signedTxn)
	if err != nil {
		log.Fatalf("Error sending transaction: %v", err)
	}
	fmt.Println("Transaction sent --", signedTxn.Hash().Hex())
}

func checkForExistingWallets(keystoreDir string, numAccounts int) bool {
	// check if keystoreDir exists
	if _, err := os.Stat(keystoreDir); os.IsNotExist(err) {
		// then create the directory
		err := os.Mkdir(keystoreDir, 0755)
		if err != nil {
			log.Fatalf("Error creating keystore directory %v: %v", keystoreDir, err)
		}

		return false
	}

	// check that numAccounts keystore files exist in keystoreDir
	// use filepath.Glob to get list of filenames matching given pattern
	filenames, err := filepath.Glob(filepath.Join(keystoreDir, "UTC*"))
	if err != nil {
		log.Fatalf("Error getting filenames from %v: %v", keystoreDir, err)
	}

	if len(filenames) != numAccounts {
		return false
	}

	return true
}

func getExistingAccounts(keystoreDir string, numAccounts int, password string) []string {
	accounts := []string{}

	// decrypt keystore files to get public addresses
	filenames, err := filepath.Glob(filepath.Join(keystoreDir, "UTC*"))
	if err != nil {
		log.Fatalf("Error getting filenames from %v: %v", keystoreDir, err)
	}

	for _, filename := range filenames {
		keystoreBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("Error reading file %v: %v", filename, err)
		}

		addr := getPublicAddressFromKeystore(keystoreBytes, password, filename)

		accounts = append(accounts, addr)
	}

	return accounts
}

func createNewAccounts(keystoreDir string, numAccounts int, password string) []string {
	accounts := []string{}

	// generate a new keystore
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// generate numAccounts new accounts
	for i := 0; i < numAccounts; i++ {
		account, err := ks.NewAccount(password)
		if err != nil {
			log.Fatalf("Error creating new account %v: %v", i, err)
		}
		accounts = append(accounts, account.Address.Hex())
	}

	return accounts
}

func getPublicAddressFromKeystore(keystoreBytes []byte, password, filename string) string {
	key, err := keystore.DecryptKey(keystoreBytes, password)
	if err != nil {
		log.Fatalf("Error decrypting keystore from %v: %v", filename, err)
	}

	pvk := key.PrivateKey
	pbk := pvk.PublicKey
	addr := crypto.PubkeyToAddress(pbk)

	return addr.Hex()
}

func weiToEth(weiAmt *big.Int) *big.Float {
	weiAmt_float := new(big.Float)
	weiAmt_float.SetInt(weiAmt)

	// 1 Ether == 10^18 Wei
	ethAmt := new(big.Float).Quo(weiAmt_float, big.NewFloat(math.Pow10(18)))

	return ethAmt
}

func ethToWei(ethAmt *big.Float) *big.Int {
	// 1 Ether == 10^18 Wei
	weiAmt_float := new(big.Float).Mul(ethAmt, big.NewFloat(math.Pow10(18)))
	weiAmt_int64, _ := weiAmt_float.Int64()

	weiAmt := new(big.Int)
	weiAmt.SetInt64(weiAmt_int64)

	return weiAmt
}

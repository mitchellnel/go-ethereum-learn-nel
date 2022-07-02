package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

/*
A keystore file (sometimes called a UTC file) in Ethereum is an encrypted version of your private
 key. They are generated using your private key and a password that you use to encrypt it. If you
 open up your keystore file in a text editor it contains data pertaining to the encryption of the
 private key.
*/

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	password := "password3"

	// generateKeystore(password)

	decryptKeystore(password)
}

func generateKeystore(password string) {
	// create a new keystore
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	// create new wallet -- passing a password
	addr, err := key.NewAccount(password)
	if err != nil {
		log.Fatalf("Error creating a new account: %v", err)
	}
	fmt.Println("Newly created wallet address is", addr.Address)
}

func decryptKeystore(password string) {
	// then we can decrypt the keystore file to get the private key
	keystoreData, err := ioutil.ReadFile(
		filepath.Join(
			"wallet",
			"UTC--2022-07-02T15-42-15.601407000Z--16628a1bdcd5e1fb37017b10044c4e7e6ee31adf",
		),
	)
	if err != nil {
		log.Fatalf("Error reading keystore file: %v", err)
	}

	key, err := keystore.DecryptKey(keystoreData, password)
	if err != nil {
		log.Fatalf("Error decrypting private key from keystore file: %v", err)
	}

	pvk := key.PrivateKey

	fmt.Println("Newly created wallet's private key is:", hexutil.Encode(crypto.FromECDSA(pvk)))

	// then get public key from private key
	pbk := pvk.PublicKey

	fmt.Println("Newly created wallet's public key is:", hexutil.Encode(crypto.FromECDSAPub(&pbk)))

	// and get public address from private key
	walletAddr := crypto.PubkeyToAddress(pbk)
	fmt.Println("Newly created wallet's public address is:", walletAddr.Hex())
}

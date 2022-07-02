package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	// generate wallet's private key
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}

	// convert private key to string
	pvk_bytes := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(pvk_bytes))

	// get public key from private key
	pbk := pvk.PublicKey

	// convert public key to string
	pbk_bytes := crypto.FromECDSAPub(&pbk)
	fmt.Println(hexutil.Encode(pbk_bytes))

	// get wallet (public) address from public key
	walletAddr := crypto.PubkeyToAddress(pbk)

	// convert to string
	fmt.Println(walletAddr.Hex())
}

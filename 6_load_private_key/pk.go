package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("INFURA_API_KEY")
	url := "https://mainnet.infura.io/v3/" + apiKey
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("could not connect to Infura with ethclient: %s", err)
	}
	ctx := context.Background()
	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("get chainID error: %s", err)
	}

	pk, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("load private key error: %s", err)
	}
	log.Printf("account load success, address: %s", crypto.PubkeyToAddress(pk.PublicKey))

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}

	_ = auth
}

package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	f, err := os.Open(os.Getenv("KEYSTORE_FILE"))
	if err != nil {
		log.Fatalf("open keystore file error: %s", err)
	}
	defer f.Close()

	auth, err := bind.NewTransactorWithChainID(f, os.Getenv("PASSWORD"), chainId)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}
	log.Printf("account load success, address: %s", auth.From)

	_ = auth
}

package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	url := "http://localhost:8545"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Could not connect to Infura with ethclient: %s", err)
	}
	log.Println("connect success")

	ctx := context.Background()
	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("get chainId error: %s", err)
	}
	log.Printf("chainId: %s", chainId)
	bn, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatalf("get chainId error: %s", err)
	}
	log.Printf("blocknumber: %d", bn)
}

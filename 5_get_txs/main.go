package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

func main() {
	apiKey := os.Getenv("INFURA_API_KEY")
	url := "https://mainnet.infura.io/v3/" + apiKey
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Could not connect to Infura with ethclient: %s", err)
	}
	ctx := context.Background()

	bn := big.NewInt(18382246)
	header, err := client.HeaderByNumber(ctx, bn)
	if err != nil {
		log.Fatalf("get header error: %s", err)
	}
	log.Printf("block %s, hash: %s, gas used: %d, base fee: %s", bn.String(), header.Hash(), header.GasUsed, header.BaseFee.String())

	block, err := client.BlockByHash(ctx, header.Hash())
	if err != nil {
		log.Fatalf("get block error: %s", err)
	}
	log.Printf("block %s, time: %d, txs: %d", block.Hash(), block.Time(), len(block.Transactions()))
}

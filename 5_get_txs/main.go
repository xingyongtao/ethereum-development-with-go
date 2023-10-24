package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
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

	tx, _, err := client.TransactionByHash(ctx, common.HexToHash("0x468f17f20488fa6011aae26fb59f9e14256ccf4a3835b28d5309b8d11b30656a"))
	if err != nil {
		log.Fatalf("get tx error: %s", err)
	}
	log.Printf("tx hash: %s", tx.Hash())
	log.Printf("chainID: %s", tx.ChainId())
	log.Printf("type: %d", tx.Type())
	log.Printf("gas limit: %d", tx.Gas())
	log.Printf("gas price: %d", tx.GasPrice())
	log.Printf("gas fee cap: %s", tx.GasFeeCap())
	log.Printf("gas tip cap: %s", tx.GasTipCap())
}

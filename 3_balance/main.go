package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
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
	addr := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045") // vitalik.eth
	bn := big.NewInt(18382246)
	balance, err := client.BalanceAt(ctx, addr, bn) // 935143001746486974073
	if err != nil {
		log.Fatalf("get chainId error: %s", err)
	}
	log.Printf("balance in Wei: %s", balance)
	bf := big.NewFloat(0).SetInt(balance)
	bf.Quo(bf, big.NewFloat(1e18)) // 935.1430017
	log.Printf("balance in Ether (converted by big.Float): %s", bf.String())

	bd := decimal.RequireFromString(balance.String())
	bd = bd.Div(decimal.NewFromFloat(1e18)) // 935.1430017464869741
	log.Printf("balance in Ether (converted decimal.Decimal): %s", bd.String())
}

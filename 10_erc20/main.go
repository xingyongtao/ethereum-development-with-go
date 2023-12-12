package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xingyongtao/ethereum-development-with-go/10_erc20/erc20"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("INFURA_API_KEY")
	url := "https://sepolia.infura.io/v3/" + apiKey
	ec, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("could not connect to Infura with ethclient: %s", err)
	}

	address := common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789")
	link, err := erc20.NewERC20(address, ec)
	if err != nil {
		log.Fatalf("new erc20 error: %s", err)
	}
	account := common.HexToAddress("0xAc25D7217FCCff38Ba52d3a7F453506522428713")
	balance, err := link.BalanceOf(nil, account)
	if err != nil {
		log.Fatalf("get balance error: %s", err)
	}
	log.Printf("account %s has balance %s", account, balance)
}

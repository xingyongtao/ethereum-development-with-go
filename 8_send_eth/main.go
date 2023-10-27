package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

func main() {
	apiKey := os.Getenv("INFURA_API_KEY")
	url := "https://sepolia.infura.io/v3/" + apiKey
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
	address := crypto.PubkeyToAddress(pk.PublicKey)
	log.Printf("account load success, address: %s", crypto.PubkeyToAddress(pk.PublicKey))

	nonce, err := client.NonceAt(ctx, address, nil)
	if err != nil {
		log.Fatalf("get nonce error: %v", err)
	}
	log.Printf("nonce: %d", nonce)
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("get header error: %v", err)
	}
	log.Printf("base fee: %s", header.BaseFee)
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatalf("get SuggestGasTipCap error: %v", err)
	}
	log.Printf("Suggested GasTipCap(maxPriorityFeePerGas): %s", gasTipCap)

	to := common.HexToAddress("0x26a1DDA0E911Ea245Fc3Fb7C5C10d18490942a60")
	amount := big.NewInt(100_000_000_000_000_000) // 0.1 ether
	txData := &types.DynamicFeeTx{
		ChainID: chainId,
		Nonce:   nonce,
		To:      &to,
		Value:   amount,
		Gas:     21000,

		GasFeeCap: header.BaseFee,
		GasTipCap: gasTipCap,
	}
	signedTx, err := types.SignNewTx(pk, types.LatestSignerForChainID(chainId), txData)
	if err != nil {
		log.Fatalf("sign tx error: %v", err)
	}
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatalf("sign tx error: %v", err)
	}
}

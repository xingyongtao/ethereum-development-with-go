package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	apiKey := os.Getenv("INFURA_API_KEY")
	url := "https://sepolia.infura.io/v3/" + apiKey
	ec, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("could not connect to Infura with ethclient: %s", err)
	}
	ctx := context.Background()

	pk, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("load private key error: %s", err)
	}

	to := common.HexToAddress("0x26a1DDA0E911Ea245Fc3Fb7C5C10d18490942a60")
	amount := big.NewInt(100_000_000_000_000_000) // 0.1 ether

	hash, err := sendEth(ctx, ec, pk, to, amount)
	if err != nil {
		log.Fatalf("send Ether error: %s", err)
	}
	log.Printf("tx sent, hash: %s", hash)
	if err = waitConfirm(ctx, ec, *hash, time.Minute*10); err != nil {
		log.Fatalf("wait confirmation error, please check the tx by yourself: %s", err)
	}
	log.Printf("tx %s confirmed", hash)
}

func sendEth(ctx context.Context, ec *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, amount *big.Int) (*common.Hash, error) {
	chainId, err := ec.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(pk.PublicKey)
	log.Printf("account: %s", address)

	nonce, err := ec.NonceAt(ctx, address, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("nonce: %d", nonce)
	header, err := ec.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("base fee: %s", header.BaseFee)
	gasTipCap, err := ec.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("Suggested GasTipCap(maxPriorityFeePerGas): %s", gasTipCap)
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
		return nil, err
	}
	err = ec.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}
	hash := signedTx.Hash()
	return &hash, nil
}

func waitConfirm(ctx context.Context, ec *ethclient.Client, txHash common.Hash, timeout time.Duration) error {
	pending := true
	for pending {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(timeout):
			return errors.New("timeout")
		case <-time.After(time.Second):
			_, isPending, err := ec.TransactionByHash(ctx, txHash)
			if err != nil {
				return err
			}
			if !isPending {
				pending = false // break `for`
			}
		}
	}
	receipt, err := ec.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		msg := fmt.Sprintf("transaction reverted, hash %s", receipt.TxHash.String())
		return errors.New(msg)
	}
	return nil
}

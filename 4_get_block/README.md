# 查询区块

对区块链数据的查询，除了之前提到的`ChainID`和当前的`BlockNumber`之外，对区块头（`Header`）的查询也比较常见，甚至有时候需要查询整个区块（`Block`）。

查询区块头和完整区块信息都有两种方法，按`Number`或者按`Hash`。

```go
bn := big.NewInt(18382246)
header, err := client.HeaderByNumber(ctx, bn)
if err != nil {
    log.Fatalf("get header error: %s", err)
}
log.Printf("block %s, hash: %s, gas used: %d, base fee: %s", bn.String(), header.Hash(), header.GasUsed, header.BaseFee.String())
```

一般情况下，重要的信息在`Header`里都已经有了，比如`BaseFee`等，不太需要检索整个区块。如果请求整个区块的话，数据量比较大，尽量要少用。

```go
block, err := client.BlockByHash(ctx, header.Hash())
if err != nil {
    log.Fatalf("get block error: %s", err)
}
log.Printf("block %s, time: %d, txs: %d", block.Hash(), block.Time(), len(block.Transactions()))
```

## 完整代码

```go
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
```

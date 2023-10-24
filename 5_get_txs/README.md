# 查询交易

比查询区块信息更常见的，是查询个别交易，这时候需要`TransactionByHash`。

查询时只需要知道`Hash`即可。

```go
tx, _, err := client.TransactionByHash(ctx, hash)
```

交易查询有时候用在确认自己的交易是否确认，有时候用在跟踪链上发生的事情，通过解析别人的交易知道发生了某些事。

## 完整代码
```go
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
```
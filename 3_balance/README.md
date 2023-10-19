# 账户余额

上一节已经获取到了以太坊公链的信息，这一节尝试获取某个账户的余额信息。（这里的余额指的是`ETH`余额，不是`ERC-20`代币余额。）

查询账户信息不需要私钥，只需要地址即可。地址的类型是`common.Address`，可以是普通的账户（`EOA`），也可以是智能合约地址（代币、金库、AA钱包、各种`Pool`等）。

接口除了接收地址参数以外，还有一个可选参数`blockNumber`，如果想查询指定区块上的余额，则传入区块号，查询最新余额的话，传入`nil`。

```go
ctx := context.Background()
addr := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045") // vitalik.eth
bn := big.NewInt(18382246)
balance, err := client.BalanceAt(ctx, addr, bn) // 935143001746486974073
```

返回值`balance`的单位是`Wei`，要转换成`Ether`需要除以1e18，详见[https://etherscan.io/unitconverter](https://etherscan.io/unitconverter)。

需要特别注意的是，账户余额可能是非常大的整数（uint256），**极易出现精度损失**，所以**非必要不转换，在本地最好存储成`Wei`**，以便和链上保持一致。

如果使用标准库的`big.Float`转换，可以这么做
```go
bf := big.NewFloat(0).SetInt(balance)
bf.Quo(bf, big.NewFloat(1e18)) // 935.1430017
```

如果使用另一个常用库[decimal](github.com/shopspring/decimal)转换，可以这么做
```go
bd := decimal.RequireFromString(balance.String())
bd = bd.Div(decimal.NewFromFloat(1e18)) // 935.1430017464869741
```

上述示例中，V神在区块`18382246`的余额是`935143001746486974073 Wei`，即`935.143001746486974073 Ether`，而使用Float转换得到的是`935.1430017 Ether`，使用`Decimal`转换后得到的是`935.1430017464869741 Ether`，都不准确。

## 完整代码

```go
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
```

如果有任何疑问，欢迎留言交流。
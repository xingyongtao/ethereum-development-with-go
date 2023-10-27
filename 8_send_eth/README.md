# 发送ETH

从这一节开始，因为涉及到提交交易，我们不再使用正式网，而是使用测试网`Sepolia`。

测试网开发也需要花费一些`Gas`，不过不用担心，可以免费领取。我推荐使用<https://sepoliafaucet.com/>水龙头领取。

确保账户里有一定的测试币后，就可以开始了。（本示例中发送的是`0.1 ether`，你可以根据账户余额情况随意调整。）

每笔交易都需要设置`Nonce`，所以先获取账户的`Nonce`，
```go
nonce, err := client.NonceAt(ctx, address, nil)
```
在从`Header`里读取现在整个网络的`baseFee`
```go
header, err := client.HeaderByNumber(ctx, nil)
// ...
log.Printf("base fee: %s", header.BaseFee)
```

后面我们直接用`baseFee`来设置`GasFeeCap`。

这么做的可能后果是，如果接下来的几个区块baseFee都比较高，我们的交易将会长时间等待，甚至被丢弃。不过因为是测试网，一般不会太繁忙，这么做问题不大。如果是实际开发，需要根据实际情况精心处理`GasFeeCap`。

同样的，我们获取网络推荐的`GasTipCap`
```go
gasTipCap, err := client.SuggestGasTipCap(ctx)
```

发送以太的交易，网络费用固定为21000，选择一个其他地址发送`0.1 ether`给他。
```go
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
```
然后对交易进行签名，注意这里使用了前面从环境变量加载进来的私钥
```go
signedTx, err := types.SignNewTx(pk, types.LatestSignerForChainID(chainId), txData)
```
签完名后的交易，就可以直接发送出去了。
```go
err = client.SendTransaction(ctx, signedTx)
```

## 完整代码
```go
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
```
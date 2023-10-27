# 导入钱包私钥

现在钱包管理的教程一般都是环境变量设置助记词或者私钥，程序导入使用。我这里也用这种方法来演示。

**注意：生产环境中，请根据实际情况，选择更稳妥的处理方式。这里的教程仅作技术演示，不是最佳方案。**

使用`go-ethereum` `crypto`包的`HexToECDSA`函数可以直接把16进制格式的私钥导入。
```go
pk, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
```

需要注意，`Go`语言导入私钥与其他语言略有不同，不需要私钥之前的`0x`前缀。

比如`anvil`测试账户`0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266`的私钥是`0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`，环境变量设置时只需要去掉`0x`的部分，即`'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`即可。

**注意：该私钥是测试工具`anvil`自动生成的，全网公开，任何人都能知道，账户里没有任何资产，也千万不要在任何正式网上尝试操作该账户。**

私钥读取成功后，可以使用`NewKeyedTransactorWithChainID`获得一个可写入的`transactor`，用于向区块链网络发起交易。
```go
auth, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
```

## 完整代码

```go
package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("INFURA_API_KEY")
	url := "https://mainnet.infura.io/v3/" + apiKey
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
	log.Printf("account load success, address: %s", crypto.PubkeyToAddress(pk.PublicKey))

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}

	_ = auth
}
```

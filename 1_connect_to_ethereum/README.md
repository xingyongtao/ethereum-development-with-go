# 连接以太坊

连接以太坊是和以太坊公链交互的第一步。所有EVM兼容链及其上的智能合约交互都需要这个基本步骤。

要连接以太坊节点，最常用的办法是使用[go-ethereum](https://github.com/ethereum/go-ethereum)库的`ethclient`。

你可以使用远程节点服务（如[Infura](https://app.infura.io/)，QuickNode，Alchemy等），也可以启动本地的调试节点（如Hardhat本地节点或者Foundry的[anvil](https://book.getfoundry.sh/reference/anvil/))。

```go
client, err := ethclient.Dial("https://mainnet.infura.io/v3/<API-KEY>")
```

## 使用远程节点（以Infura为例）

使用远程节点一般需要注册账户并获得一个私有的`API_KEY`。具体步骤可以参考对应网站介绍。获得`API_KEY`后就可以直接连接了。

### 完整代码
```go
package main

import (
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
	log.Println("connect success")

	_ = client
}
```

## 使用本地调试节点（以`anvil`为例）

`anvil`是`Foundry`工具集的一部分，详细使用方法可见 https://book.getfoundry.sh/reference/anvil/。

启动本地调试节点
```bash
anvil
```
默认监听端口是8545。只需要修改对应的连接地址即可。

```go
client, err := ethclient.Dial("http://localhost:8545")
```

当然，也可以尝试输出对应的`ChainId`和`BlockNumber`，你会发现和主网的区别。

### 完整代码

```go
package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	url := "http://localhost:8545"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Could not connect to Infura with ethclient: %s", err)
	}
	log.Println("connect success")

	ctx := context.Background()
	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("get chainId error: %s", err)
	}
	log.Printf("chainId: %s", chainId)
	bn, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatalf("get chainId error: %s", err)
	}
	log.Printf("blocknumber: %d", bn)
}
```

输出应该是类似如下的内容
```bash
2023/10/18 11:50:26 connect success
2023/10/18 11:50:26 chainId: 31337
2023/10/18 11:50:26 blocknumber: 0
```

恭喜你，已经学会了使用`Go`连接以太坊，可以开始神奇的以太坊开发旅程了。
# 同质化代币ERC20

这一节终于来到了同质化代币ERC20，有些场合也会被叫做Token，即令牌。

关于ERC20代币标准的资料已经非常多，这里不展开了。这一节主要看一下Go语言如何和ERC20代币交互。
我们先从订阅转账消息开始。

Go和以太坊智能合约交互，是使用了以太坊的JSON-RPC接口，即把所有读取写入请求转换成RPC请求，发送给以太坊节点，并等待回应。这和之前转账是完全一致的，只不过交互的地址从EOA变成了合约而已。

但是，合约交互不比简单转账，它是有参数的，而且参数的编码也要严格遵守规范。如果不借助工具，这个编码过程是非常繁琐的。幸运的是，有工具。

`go-ethereum`自带的`abigen`可以根据`ABI`文件，自动生成对应的Go文件。有了这些文件，我们不用费功夫编解码了，可以直接使用所有函数，跟本地函数类似。

获得`ABI`文件的方法有很多，可以从etherscan网站拷贝，可以从源码编译等。我随便举例其中的一种。

我先从OpenZeppelin库<https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v5.0.1/contracts/token/ERC20/IERC20.sol>下载接口文件，然后放到Remix里编译，拷贝出`ABI`文件。

假设保存的`ABI`文件名是`IERC20.json`。运行命令

```go
abigen --abi IERC20.json --type ERC20 --pkg er
c20 --out erc20.go
```

可以得到`Go`文件`erc20.go`，里边自动生成的类名是`ERC20`。

有了这个文件，我们可以找一个符合ERC20标准的测试网代币试一试了。我选中了`sepolia`网络的LINK代币，因为这个代币的水龙头容易找 <https://faucets.chain.link/sepolia>。它的合约地址是`0x779877A7B0D9E8603169DdbD7836e478b4624789`。

使用之前的`*ethclient.Client`，新建一个`ERC20`对象

```go
address := common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789")
link, err := erc20.NewERC20(address, ec)
```

查询某个账户持有的代币余额
```go
account := common.HexToAddress("0xAc25D7217FCCff38Ba52d3a7F453506522428713")
balance, err := link.BalanceOf(nil, account)
```

## 完整代码

```go
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
```
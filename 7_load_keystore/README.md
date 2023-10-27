# 导入`Keystore`文件

导入`Keystore`文件的教程网络上较少，但它其实更安全，而且`Go`处理起来也非常方便。

具体的步骤和导入私钥差不多，只是多了个密码而已。为了模拟更真实的使用场景，我们这次直接从文件中读入。

假设`Keystore`文件用环境变量`KEYSTORE_FILE`指定，先打开文件。
```go
f, err := os.Open(os.Getenv("KEYSTORE_FILE"))
if err != nil {
    log.Fatalf("open keystore file error: %s", err)
}
defer f.Close()
```

然后使用密码解锁文件，并生成可签名发送交易的`transactor`对象，这次使用的是`NewTransactorWithChainID`函数
```go
auth, err := bind.NewTransactorWithChainID(f, os.Getenv("PASSWORD"), chainId)
if err != nil {
    log.Fatalf("failed to create authorized transactor: %v", err)
}
```

## 完整代码

```go
package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	f, err := os.Open(os.Getenv("KEYSTORE_FILE"))
	if err != nil {
		log.Fatalf("open keystore file error: %s", err)
	}
	defer f.Close()

	auth, err := bind.NewTransactorWithChainID(f, os.Getenv("PASSWORD"), chainId)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}
	log.Printf("account load success, address: %s", auth.From)

	_ = auth
}
```

`Keystore`文件因为不明文存储私钥，所以更安全。没有密码光有文件解锁不了，而光有密码没有文件也不可能拿到私钥，所以给了我们**分开存储**的便利性。

如果有一天你怀疑密码或者`Keystore`文件可能被盗或者泄露了，还可以通过老密码和老文件，导入后重新生成新的`Keystore`文件和新密码，然后把老文件老密码彻底废弃不用。
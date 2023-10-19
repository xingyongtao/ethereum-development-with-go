# 获得ChainID和BlockNumber

以太坊节点提供`JSON-RPC`服务，如果开启了，就可以通过`HTTP/HTTPS`协议访问节点。

该服务是可选的，有一些自用节点不一定会开启。远程节点是专门提供该服务的，自然是开着的。本地调试节点（Hardhat、anvil等）默认都是开启的。

从本节开始后的章节，基本都使用远程节点，不再特殊说明。

在[连接到以太坊](../1_connect_to_ethereum/README.md)之后，我们就可以开始发起请求了。
常见请求有：
1. 获取ChainID用于核实网络是否匹配
2. 获取网络信息，如`BlockNumber`、账户余额等
3. 与智能合约交互
4. 监听网络事件

我们先从获取`ChainID`和`BlockNumber`开始讲起。

在`EVM`兼容网络之间，`ChainID`是区分网络的重要手段，每一个网络都有唯一的`ID`与其对应。
比如，以太坊的`ID`是`1`，Polygon的`ID`是`137`，而BSC的`ID`是`56`。想知道一个网络的具体`ID`，可以参考[https://chainlist.org/](https://chainlist.org/)。
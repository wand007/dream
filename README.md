
## dream 梦想

智能合约（Smart Contract）是运行在区块链上 的计算机程序。区块链技术让人们可以在去中心化 的情况下达成共识，而智能合约则决定了可以达成 什么样的共识。换言之，区块链只是一个分布式的 记账方式形成的公共账本，智能合约则进一步结合 千千万万个不同的应用场景和经济活动，约定了谁 与谁、在什么情况下记账，进而产生什么样的账本。


go mod init

GO111MODULE=on go mod vendor

区块链

https://github.com/hyperledger/fabric

https://github.com/hyperledger/fabric-samples

|  **系统工具** | **版本**     | 
| -------------|-------------|
| ubuntu       |20.04        | 
| java         |1.8          | 
| golang       |go1.15.2     |
| fabric       |2.1.0        |
| fabric-gateway-java |2.1.4 |


Fabric网络结构

|  **节点类型** | **节点名**      |  **所属组织** | **ip**     |  **服务端口** | 
| -------------|-------------| -------------|-------------| -------------|
| orderer	   |orderer.example.com	-	|org0   |192.168.0.2	|7050
| peer	       |peer0.org1.example.com	|org1	|192.168.0.2	|7051
| peer	       |peer1.org1.example.com	|org1	|192.168.0.2	|8051
| peer	       |peer0.org2.example.com	|org2	|192.168.0.2	|9051
| peer	       |peer1org2.example.com	|org2	|192.168.0.2	|10051
| ca	       |ca1.org1.example.com	|org1	|192.168.0.2	|7054
| ca	       |ca2.org1.example.com	|org2	|192.168.0.2	|8054
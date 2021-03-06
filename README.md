
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
| orderer	   |orderer1.org0.example.com	|org0   |192.168.0.2	|7050
| orderer	   |orderer2.org0.example.com	|org0   |192.168.0.2	|8050
| orderer	   |orderer3.org0.example.com	|org0   |192.168.0.2	|9050
| peer	       |peer0.org1.example.com	|org1	|192.168.0.2	|7051
| peer	       |peer1.org1.example.com	|org1	|192.168.0.2	|8051
| peer	       |peer0.org2.example.com	|org2	|192.168.0.2	|9051
| peer	       |peer1.org2.example.com	|org2	|192.168.0.2	|10051
| peer	       |peer0.org3.example.com	|org3	|192.168.0.2	|11051
| peer	       |peer1.org3.example.com	|org3	|192.168.0.2	|12051
| peer	       |peer0.org4.example.com	|org4	|192.168.0.2	|13051
| peer	       |peer1.org4.example.com	|org4	|192.168.0.2	|14051
| peer	       |peer0.org5.example.com	|org5	|192.168.0.2	|15051
| peer	       |peer1.org5.example.com	|org5	|192.168.0.2	|16051
| couchdb	   |couchdb0.org1.example.com	|org1	|192.168.0.2	|5984
| couchdb	   |couchdb1.org1.example.com	|org1	|192.168.0.2	|6984
| couchdb	   |couchdb0.org2.example.com	|org2	|192.168.0.2	|7984
| couchdb	   |couchdb1.org2.example.com	|org2	|192.168.0.2	|8984
| couchdb	   |couchdb0.org3.example.com	|org3	|192.168.0.2	|9984
| couchdb	   |couchdb1.org3.example.com	|org3	|192.168.0.2	|10984
| couchdb	   |couchdb0.org4.example.com	|org4	|192.168.0.2	|11984
| couchdb	   |couchdb1.org4.example.com	|org4	|192.168.0.2	|12984
| couchdb	   |couchdb0.org5.example.com	|org5	|192.168.0.2	|13984
| couchdb	   |couchdb1.org5.example.com	|org5	|192.168.0.2	|14984
| ca	       |ca-tls.ca.example.com	|tls	|192.168.0.2	|7052
| ca	       |ca-org0.ca.example.com	|org0	|192.168.0.2	|7053
| ca	       |ca-org1.ca.example.com	|org1	|192.168.0.2	|7054
| ca	       |ca-org2.ca.example.com	|org2	|192.168.0.2	|7055
| ca	       |ca-org3.ca.example.com	|org3	|192.168.0.2	|7056
| ca	       |ca-org4.ca.example.com	|org4	|192.168.0.2	|7057
| ca	       |ca-org5.ca.example.com	|org5	|192.168.0.2	|7058
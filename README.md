
## dream 梦想

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
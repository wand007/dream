
## 官方文档：https://hyperledger-fabric.readthedocs.io/zh_CN/release-2.2/private_data_tutorial.html

## 启动cli服务
docker-compose -f docker-compose-cli-peers.yaml up  -d 2>&1


docker exec -it cli-org1-peer0 bash
# 创建通道
peer channel create -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/channel.tx --outputBlock /usr/local/channel-artifacts/$CHANNEL_NAME.block --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE

exit

# 加入通道
docker exec -it cli-org1-peer0 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org1-peer1 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org2-peer0 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org2-peer1 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org3-peer0 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org3-peer1 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org4-peer0 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org4-peer1 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org5-peer0 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org5-peer1 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org6-peer0 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

# 加入通道
docker exec -it cli-org6-peer1 bash

peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit


# 更新锚节点
docker exec -it cli-org1-peer0 bash

peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org1MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE

exit

# 更新锚节点
docker exec -it cli-org2-peer0 bash

peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org2MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE

exit

# 更新锚节点
docker exec -it cli-org3-peer0 bash

peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org3MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE

exit

# 更新锚节点
docker exec -it cli-org4-peer0 bash

peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org4MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE

exit

# 更新锚节点
docker exec -it cli-org5-peer0 bash

peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org5MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE

exit

# 更新锚节点
docker exec -it cli-org6-peer0 bash

peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org6MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE

exit

## platform链码安装,只启动一个节点

# pee0-org1安装链码
docker exec -it cli-org1-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/main/
# 设置golang的环境变量
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

pushd $CC_SRC_PATH
GO111MODULE=on go mod vendor
popd

# 打包链码
peer lifecycle chaincode package /usr/local/chaincode-artifacts/platform.tar.gz --path $CC_SRC_PATH --lang golang --label platform_1

# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/platform.tar.gz

# 将链码id设置变量,便于我们后面的使用
export CC_PACKAGE_ID=platform_1:575bf94df6be77ff046e86ee94b6cb99116eb6b0c1a7faaf5ab179c3e382344c

# 查看peer0.org1.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --channelID $CHANNEL_NAME --name platform --version 1 --init-required --package-id $CC_PACKAGE_ID --sequence 1 --signature-policy "OR('Org1MSP.member')" --waitForEvent

# 查看链码认证结果 此时只有Org1MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name platform --version 1 --sequence 1 --output json --init-required  --signature-policy "OR('Org1MSP.member')"

exit


# pee0-org2安装链码
docker exec -it cli-org2-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/config/collections_config.json
# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/platform.tar.gz

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --channelID $CHANNEL_NAME --name platform --version 1 --init-required --sequence 1 --signature-policy "OR('Org1MSP.member')" --waitForEvent

# 查看链码认证结果 此时只有Org1MSP和Org2MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name platform --version 1 --sequence 1 --output json --init-required --signature-policy "OR('Org1MSP.member')"

exit


# pee0-org3安装链码
docker exec -it cli-org3-peer0 bash
# 重复pee0-org2安装链码

# pee0-org4安装链码
docker exec -it cli-org4-peer0 bash
# 重复pee0-org2安装链码

# pee0-org5安装链码
docker exec -it cli-org5-peer0 bash
# 重复pee0-org2安装链码

# pee0-org6安装链码
docker exec -it cli-org6-peer0 bash
# 重复pee0-org2安装链码

# pee1-org1安装链码
docker exec -it cli-org1-peer1 bash

# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/marbles02_private.tar.gz

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

exit


# 部署链码
docker exec -it cli-org1-peer0 bash

# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/config/collections_config.json

# 提交链码
peer lifecycle chaincode commit -o orderer1.org0.example.com:7050 --channelID $CHANNEL_NAME --name platform --version 1 --sequence 1 --init-required --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --signature-policy "OR('Org1MSP.member')"

# 查询已经提交的链码
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name platform

# 链码执行
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n platform --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --isInit  -c '{"function":"InitLedger","Args":[]}' --waitForEvent
# 查询
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n platform --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE    -c '{"function":"FindById","Args":["P768877118787432448"]}'

exit


## platform链码安装

# pee0-org1安装链码
docker exec -it cli-org1-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/config/collections_config.json
# 设置golang的环境变量
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

pushd $CC_SRC_PATH
GO111MODULE=on go mod vendor
popd

# 打包链码
peer lifecycle chaincode package /usr/local/chaincode-artifacts/platform.tar.gz --path $CC_SRC_PATH --lang golang --label platform_1

# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/platform.tar.gz

# 将链码id设置变量,便于我们后面的使用
export CC_PACKAGE_ID=platform_1:575bf94df6be77ff046e86ee94b6cb99116eb6b0c1a7faaf5ab179c3e382344c


# 查看peer0.org1.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --channelID $CHANNEL_NAME --name platform --version 1 --init-required --package-id $CC_PACKAGE_ID --sequence 1 --collections-config $CC_CC_PATH --waitForEvent

# 查看链码认证结果 此时只有Org1MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name platform --version 1 --sequence 1 --output json --init-required  --collections-config $CC_CC_PATH

exit


# pee0-org2安装链码
docker exec -it cli-org2-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/config/collections_config.json
# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/platform.tar.gz
# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --channelID $CHANNEL_NAME --name platform --version 1 --init-required --sequence 1 --collections-config $CC_CC_PATH --waitForEvent

# 查看链码认证结果 此时只有Org1MSP和Org2MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name platform --version 1 --sequence 1 --output json --init-required  --collections-config $CC_CC_PATH

exit


# pee0-org3安装链码
docker exec -it cli-org3-peer0 bash
# 重复pee0-org2安装链码

# pee0-org4安装链码
docker exec -it cli-org4-peer0 bash
# 重复pee0-org2安装链码

# pee0-org5安装链码
docker exec -it cli-org5-peer0 bash
# 重复pee0-org2安装链码

# pee0-org6安装链码
docker exec -it cli-org6-peer0 bash
# 重复pee0-org2安装链码

# pee1-org1安装链码
docker exec -it cli-org1-peer1 bash

# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/marbles02_private.tar.gz

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

exit


# 部署链码
docker exec -it cli-org1-peer0 bash

# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/platform/config/collections_config.json

# 提交链码
peer lifecycle chaincode commit -o orderer1.org0.example.com:7050 --channelID $CHANNEL_NAME --name platform --version 1 --sequence 1 --init-required --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --collections-config $CC_CC_PATH

# 查询已经提交的链码
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name platform

# 链码执行
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n platform --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE   --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --isInit  -c '{"Args":["InitLedger"]}' --waitForEvent
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n platform --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE   --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --isInit  -c '{"function":"InitLedger","Args":[]}' --waitForEvent
exit

## 测试链码
docker exec -it cli-org1-peer0 bash

# 创建一条记录marble1
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n marbles02_private --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE   --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE -c '{"Args":["initMarble"]}' --transient "{\"marble\":\"$MARBLE\"}" --waitForEvent

# 查询 Org1 被授权的私有数据
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarble","marble1"]}'
# 返回信息：{"color":"blue","docType":"marble","name":"marble1","owner":"tom","size":35}

# 查询 Org1 未被授权的私有数据
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarblePrivateDetails","marble1"]}'
# 返回信息：{"docType":"marblePrivateDetails","name":"marble1","price":99}

docker exec -it cli-org2-peer0 bash

# 查询 Org2 被授权的私有数据
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarble","marble1"]}'
# 返回信息：{"color":"blue","docType":"marble","name":"marble1","owner":"tom","size":35}

# 查询 Org2 未被授权的私有数据
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarblePrivateDetails","marble1"]}'
# 异常日志 tx creator does not have read access permission on privatedata in chaincodeName:marbles02_private collectionName: collectionMarblePrivateDetails

# 其他方法
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["getMarblesByRange","marble1","marble4"]}'
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["queryMarblesByOwner","tom"]}'
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["queryMarbles","{\"selector\":{\"owner\":\"tom\"}}"]}'




## 清除私有数据
docker exec -it cli-org1-peer0 bash


## 打开一个新终端窗口，通过运行如下命令来查看这个节点上私有数据日志
docker logs peer0.org1.example.com 2>&1 | grep -i -a -E 'private|pvt|privdata'


# 查询 marble1 的 price 数据（查询并不会产生一笔新的交易）
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarblePrivateDetails","marble1"]}'

# 创建一个新的 marble2。这个交易将在链上创建一个新区块
export MARBLE=$(echo -n "{\"name\":\"marble2\",\"color\":\"blue\",\"size\":35,\"owner\":\"tom\",\"price\":99}" | base64 | tr -d \\n)
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n marbles02_private --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE -c '{"Args":["initMarble"]}' --transient "{\"marble\":\"$MARBLE\"}"

# 查询 marble1 的价格数据
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarblePrivateDetails","marble1"]}'

# 私有数据没有被清除，查询结果也没有改变


# 将 marble2 转移给 “joe” 。这个交易将使链上增加第二个区块。
export MARBLE_OWNER=$(echo -n "{\"name\":\"marble2\",\"owner\":\"joe\"}" | base64 | tr -d \\n)
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n marbles02_private --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE -c '{"Args":["transferMarble"]}' --transient "{\"marble_owner\":\"$MARBLE_OWNER\"}" --waitForEvent

## 切换回终端窗口并查看节点的私有数据日志。你将看到区块高度增加了 1 。
docker logs peer0.org1.example.com 2>&1 | grep -i -a -E 'private|pvt|privdata'

# 再次运行如下命令查询 marble1 的价格数据
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarblePrivateDetails","marble1"]}'

# 将 marble2 转移给 “tom” 。这个交易将使链上增加第三个区块。

export MARBLE_OWNER=$(echo -n "{\"name\":\"marble2\",\"owner\":\"tom\"}" | base64 | tr -d \\n)
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n marbles02_private --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE -c '{"Args":["transferMarble"]}' --transient "{\"marble_owner\":\"$MARBLE_OWNER\"}" --waitForEvent

## 切换回终端窗口并查看节点的私有数据日志。你将看到区块高度增加了 1 。
docker logs peer0.org1.example.com 2>&1 | grep -i -a -E 'private|pvt|privdata'

# 回到节点容器，再次运行如下命令查询 marble1 的价格数据
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarblePrivateDetails","marble1"]}'

# 最后，运行下边的命令将 marble2 转移给 “jerry” 。这个交易将使链上增加第四个区块。在此次交易之后，price 私有数据将会被清除。
export MARBLE_OWNER=$(echo -n "{\"name\":\"marble2\",\"owner\":\"jerry\"}" | base64 | tr -d \\n)
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n marbles02_private --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE -c '{"Args":["transferMarble"]}' --transient "{\"marble_owner\":\"$MARBLE_OWNER\"}" --waitForEvent

## 再次切换回终端窗口并查看节点的私有数据日志。你将看到区块高度增加了 1 。
docker logs peer0.org1.example.com 2>&1 | grep -i -a -E 'private|pvt|privdata'

# 回到节点容器，再次运行如下命令查询 marble1 的价格数据
peer chaincode query -C $CHANNEL_NAME -n marbles02_private -c '{"Args":["readMarblePrivateDetails","marble1"]}'

# 因为价格数据已经被清除了，所以你就查询不到了。应该会看到类似下边的结果:
# Error: endorsement failure during query. response: status:500
# message:"{\"Error\":\"Marble private details does not exist: marble1\"}"



## individual链码安装

# pee0-org1安装链码
docker exec -it cli-org1-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/config/collections_config.json
# 设置golang的环境变量
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

pushd $CC_SRC_PATH
GO111MODULE=on go mod vendor
popd

# 打包链码
peer lifecycle chaincode package /usr/local/chaincode-artifacts/individual.tar.gz --path $CC_SRC_PATH --lang golang --label individual_1

# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/individual.tar.gz

# 将链码id设置变量,便于我们后面的使用
export CC_PACKAGE_ID=individual_1:521dfb811bc066a352126818252bb01f4571aed637eeb7575418bb3a0df630c5

# 查看peer0.org1.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --channelID $CHANNEL_NAME --name individual --version 1 --init-required --package-id $CC_PACKAGE_ID --sequence 1  --collections-config $CC_CC_PATH --signature-policy "OR('Org1MSP.member')" --waitForEvent

# 查看链码认证结果 此时只有Org1MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name individual --version 1 --sequence 1 --output json --init-required  --collections-config $CC_CC_PATH  --signature-policy "OR('Org1MSP.member')"

exit


# pee0-org2安装链码
docker exec -it cli-org2-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/config/collections_config.json
# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/individual.tar.gz

# 将链码id设置变量,便于我们后面的使用
#export CC_PACKAGE_ID=individual_1:521dfb811bc066a352126818252bb01f4571aed637eeb7575418bb3a0df630c5

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --channelID $CHANNEL_NAME --name individual --version 1 --init-required --sequence 1  --collections-config $CC_CC_PATH --signature-policy "OR('Org1MSP.member')" --waitForEvent

# 查看链码认证结果 此时只有Org1MSP和Org2MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name individual --version 1 --sequence 1 --output json --init-required  --collections-config $CC_CC_PATH --signature-policy "OR('Org1MSP.member')"

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
peer lifecycle chaincode install /usr/local/chaincode-artifacts/individual.tar.gz

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

exit


# 部署链码
docker exec -it cli-org1-peer0 bash

# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/config/collections_config.json

# 提交链码
peer lifecycle chaincode commit -o orderer1.org0.example.com:7050 --channelID $CHANNEL_NAME --name individual --version 1 --sequence 1 --init-required --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --collections-config $CC_CC_PATH --signature-policy "OR('Org1MSP.member')"

# 查询已经提交的链码
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name individual

# 链码执行
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n individual --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --isInit  -c '{"function":"InitLedger","Args":[]}' --waitForEvent
# 查询
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n individual --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE    -c '{"function":"FindById","Args":["P768877118787432448"]}'

exit


## individual链码安装

# pee0-org1安装链码
docker exec -it cli-org1-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/config/collections_config.json
# 设置golang的环境变量
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

pushd $CC_SRC_PATH
GO111MODULE=on go mod vendor
popd

# 打包链码
peer lifecycle chaincode package /usr/local/chaincode-artifacts/individual.tar.gz --path $CC_SRC_PATH --lang golang --label individual_1

# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/individual.tar.gz

# 将链码id设置变量,便于我们后面的使用
export CC_PACKAGE_ID=individual_1:521dfb811bc066a352126818252bb01f4571aed637eeb7575418bb3a0df630c5


# 查看peer0.org1.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --channelID $CHANNEL_NAME --name individual --version 1 --init-required --package-id $CC_PACKAGE_ID --sequence 1 --collections-config $CC_CC_PATH --waitForEvent

# 查看链码认证结果 此时只有Org1MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name individual --version 1 --sequence 1 --output json --init-required  --collections-config $CC_CC_PATH

exit


# pee0-org2安装链码
docker exec -it cli-org2-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/config/collections_config.json
# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/individual.tar.gz
# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --channelID $CHANNEL_NAME --name individual --version 1 --init-required --sequence 1 --collections-config $CC_CC_PATH --waitForEvent

# 查看链码认证结果 此时只有Org1MSP和Org2MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name individual --version 1 --sequence 1 --output json --init-required  --collections-config $CC_CC_PATH

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
peer lifecycle chaincode install /usr/local/chaincode-artifacts/individual.tar.gz

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

exit


# 部署链码
docker exec -it cli-org1-peer0 bash

# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/individual/config/collections_config.json

# 提交链码
peer lifecycle chaincode commit -o orderer1.org0.example.com:7050 --channelID $CHANNEL_NAME --name individual --version 1 --sequence 1 --init-required --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --collections-config $CC_CC_PATH

# 查询已经提交的链码
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name individual

# 链码执行
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n individual --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"Args":[]}' --waitForEvent

## 测试链码
# 初始化默认数据
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n individual --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"InitLedger","Args":[]}' --waitForEvent
# 查询默认公开数据
peer chaincode query -C $CHANNEL_NAME -n individual   -c '{"function":"FindById","Args":["IN760934239574175744"]}'
# 查询默认私有数据
peer chaincode query -C $CHANNEL_NAME -n individual   -c '{"function":"FindPrivateDataById","Args":["IN760934239574175744"]}'

# 新建个体
export MARBLE=$(echo -n "{\"id\":\"IN756579272398741516\",\"name\":\"新建个体1\",\"platformOrgID\":\"P768877118787432448\",\"certificateNo\":\"666666666666666666\",\"certificateType\":1,\"status\":1}" | base64 | tr -d \\n)
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n individual --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"Create","Args":[]}' --transient "{\"individual\":\"$MARBLE\"}" --waitForEvent

# 修改个体
export MARBLE=$(echo -n "{\"id\":\"IN756579272398741516\",\"name\":\"新建个体2\",\"platformOrgID\":\"P768877118787432448\",\"certificateNo\":\"666666666666666666\",\"certificateType\":1,\"status\":1}" | base64 | tr -d \\n)
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n individual --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE]  -c '{"function":"Update","Args":[]}' --transient "{\"individual\":\"$MARBLE\"}" --waitForEvent

exit

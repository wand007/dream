
## 分销商机构 financial链码安装

# pee0-org1安装链码
docker exec -it cli-org1-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial/config/collections_config.json
# 设置golang的环境变量
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

pushd $CC_SRC_PATH
GO111MODULE=on go mod vendor
popd

# 打包链码
peer lifecycle chaincode package /usr/local/chaincode-artifacts/financial.tar.gz --path $CC_SRC_PATH --lang golang --label financial_1

# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/financial.tar.gz

# 将链码id设置变量,便于我们后面的使用
export CC_PACKAGE_ID=financial_1:72f8d1522aeb44fd837163fd9d9ade5d8b20aafc25fa72ea7b35537710f8e24a

# 查看peer0.org1.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg --channelID $CHANNEL_NAME --name financial --version 1 --sequence 1 --init-required --package-id $CC_PACKAGE_ID -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE  --collections-config $CC_CC_PATH --signature-policy "AND('Org1MSP.member', 'Org2MSP.member')" --waitForEvent

# 查看链码认证结果 此时只有Org1MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name financial --version 1 --sequence 1 --output json --init-required  --collections-config $CC_CC_PATH  --signature-policy "AND('Org1MSP.member', 'Org2MSP.member')"

exit


# pee0-org2安装链码
docker exec -it cli-org2-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial/config/collections_config.json
# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/financial.tar.gz

# 将链码id设置变量,便于我们后面的使用
export CC_PACKAGE_ID=financial_1:72f8d1522aeb44fd837163fd9d9ade5d8b20aafc25fa72ea7b35537710f8e24a

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg --channelID $CHANNEL_NAME --name financial --version 1 --sequence 1 --init-required --package-id $CC_PACKAGE_ID -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --collections-config $CC_CC_PATH --signature-policy "AND('Org1MSP.member', 'Org2MSP.member')" --waitForEvent

# 查看链码认证结果 此时只有Org1MSP和Org2MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name financial --version 1 --sequence 1 --output json --init-required  --collections-config $CC_CC_PATH --signature-policy "AND('Org1MSP.member', 'Org2MSP.member')"

exit

# pee0-org3安装链码
docker exec -it cli-org3-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial/config/collections_config.json
# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/financial.tar.gz

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg --channelID $CHANNEL_NAME --name financial --version 1 --init-required --sequence 1 -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --collections-config $CC_CC_PATH --signature-policy "AND('Org1MSP.member', 'Org2MSP.member')" --waitForEvent

# 查看链码认证结果 此时只有Org1MSP和Org2MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name financial --version 1 --sequence 1 --output json --init-required --collections-config $CC_CC_PATH --signature-policy "AND('Org1MSP.member', 'Org2MSP.member')"

exit

# pee0-org4安装链码
docker exec -it cli-org4-peer0 bash
# 重复pee0-org3安装链码
exit

# pee0-org5安装链码
docker exec -it cli-org5-peer0 bash
# 重复pee0-org3安装链码
exit

# pee0-org6安装链码
docker exec -it cli-org6-peer0 bash
# 重复pee0-org3安装链码
exit



# 部署链码
docker exec -it cli-org1-peer0 bash

# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial/config/collections_config.json

# 提交链码
peer lifecycle chaincode commit -o orderer1.org0.example.com:7050 --channelID $CHANNEL_NAME --name financial --version 1 --sequence 1 --init-required --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --collections-config $CC_CC_PATH --signature-policy "AND('Org1MSP.member', 'Org2MSP.member')" --waitForEvent

# 查询已经提交的链码
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name financial

# 链码执行
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n financial --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --isInit  -c '{"Args":[]}' --waitForEvent

## 测试链码
# 初始化默认数据
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n financial --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE   --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"InitLedger","Args":[]}' --waitForEvent
# 查询默认公开数据
peer chaincode query -C $CHANNEL_NAME -n financial   -c '{"function":"FindById","Args":["F766005404604841984"]}'
# 查询默认私有数据
peer chaincode query -C $CHANNEL_NAME -n financial   -c '{"function":"FindPrivateDataById","Args":["F766005404604841984"]}'

# 新建分销商机构
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n financial --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE   --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"Create","Args":["F736182013215645696","新增分销商机构1","3","1"]}' --waitForEvent

# 一般账户向共管账户现金兑换票据
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n financial --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE   --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"TransferAsset","Args":["F766374712807800832","F766374712807800832","3"]}' --waitForEvent

exit

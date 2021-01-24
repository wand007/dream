
## 金融机构一般账户 financial_general_account链码安装

# pee0-org1安装链码
docker exec -it cli-org1-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial_general_account/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial_general_account/config/collections_config.json
# 设置golang的环境变量
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

pushd $CC_SRC_PATH
GO111MODULE=on go mod vendor
popd

# 打包链码
peer lifecycle chaincode package /usr/local/chaincode-artifacts/financial_general_account.tar.gz --path $CC_SRC_PATH --lang golang --label financial_general_account_1

# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/financial_general_account.tar.gz

# 将链码id设置变量,便于我们后面的使用
export CC_PACKAGE_ID=financial_general_account_1:57bc23c3802d0b86f7bc16ce140f64088d36486692bb50205d63f542c1a3c7f4

# 查看peer0.org1.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg --channelID $CHANNEL_NAME --name financial_general_account --version 1 --sequence 3 --init-required --package-id $CC_PACKAGE_ID -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE  --collections-config $CC_CC_PATH --waitForEvent

# 查看链码认证结果 此时只有Org1MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name financial_general_account --version 1 --sequence 3 --output json --init-required  --collections-config $CC_CC_PATH

exit


# pee0-org2安装链码
docker exec -it cli-org2-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial_general_account/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial_general_account/config/collections_config.json
# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/financial_general_account.tar.gz

# 将链码id设置变量,便于我们后面的使用
export CC_PACKAGE_ID=financial_general_account_1:57bc23c3802d0b86f7bc16ce140f64088d36486692bb50205d63f542c1a3c7f4

# 查看peer0.org2.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg --channelID $CHANNEL_NAME --name financial_general_account --version 1 --sequence 3 --init-required --package-id $CC_PACKAGE_ID -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --collections-config $CC_CC_PATH --waitForEvent

# 查看链码认证结果 此时只有Org1MSP和Org2MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name financial_general_account --version 1 --sequence 3 --output json --init-required  --collections-config $CC_CC_PATH

exit


# pee0-org3安装链码
docker exec -it cli-org3-peer0 bash

# 链码路径
export CC_SRC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial_general_account/main/
# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial_general_account/config/collections_config.json
# 安装链码
peer lifecycle chaincode install /usr/local/chaincode-artifacts/financial_general_account.tar.gz

# 查看peer0.org3.example.com链码安装结果
peer lifecycle chaincode queryinstalled

# 链码认证 根据设置的链码审批规则，只需要当前组织中的任意一个节点审批通过即可
peer lifecycle chaincode approveformyorg --channelID $CHANNEL_NAME --name financial_general_account --version 1 --sequence 3 --init-required -o orderer1.org0.example.com:7050 --ordererTLSHostnameOverride orderer1.org0.example.com --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --collections-config $CC_CC_PATH --waitForEvent

# 查看链码认证结果 此时只有Org1MSP和Org2MSP审核通过了
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name financial_general_account --version 1 --sequence 3 --output json --init-required  --collections-config $CC_CC_PATH

exit

# pee0-org4安装链码
docker exec -it cli-org4-peer0 bash
# 重复pee0-org3安装链码

# pee0-org5安装链码
docker exec -it cli-org5-peer0 bash
# 重复pee0-org3安装链码

# pee0-org6安装链码
docker exec -it cli-org6-peer0 bash
# 重复pee0-org3安装链码

exit



# 部署链码
docker exec -it cli-org1-peer0 bash

# 私有数据规则配置文件路径
export CC_CC_PATH=/opt/gopath/src/github.com/hyperledger/chaincode/financial_general_account/config/collections_config.json

# 提交链码
peer lifecycle chaincode commit -o orderer1.org0.example.com:7050 --channelID $CHANNEL_NAME --name financial_general_account --version 1 --sequence 3 --init-required --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org3.example.com:11051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org4.example.com:13051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0.org5.example.com:15051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org6.example.com:17051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --collections-config $CC_CC_PATH

# 查询已经提交的链码
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name financial_general_account

# 链码执行
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n financial_general_account --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --isInit  -c '{"Args":[]}' --waitForEvent

## 测试链码
# 查询默认私有数据
peer chaincode query -C $CHANNEL_NAME -n financial_general_account   -c '{"function":"FindPrivateDataById","Args":["6229486603953152819"]}'

# 新建金融机构一般账户
export MARBLE=$(echo -n "{\"cardNo\":\"6229486603953201814\",\"financialOrgID:\"F766005404604841984\",\"certificateNo\":\"91370181MA3D7J9W3W\",\"certificateType\":5,\"currentBalance\":0,\"voucherCurrentBalance\":0,\"accStatus\":1}" | base64 | tr -d \\n)
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n individual --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"Create","Args":[]}' --transient "{\"generalAccount\":\"$MARBLE\"}" --waitForEvent

# 零售商向零售商一般账户充值现金余额
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n financial --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"TransferCashAsset","Args":["6229486603953201814","101"]}' --waitForEvent
# 查询默认私有数据 ----现金余额变化 0--101
peer chaincode query -C $CHANNEL_NAME -n financial_general_account   -c '{"function":"FindPrivateDataById","Args":["6229486603953201814"]}'
# 票据交易
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n financial --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"TransferVoucherAsset","Args":["6229486603953201814","103"]}' --waitForEvent
# 查询默认私有数据 ----票据变化 0--103
peer chaincode query -C $CHANNEL_NAME -n financial_general_account   -c '{"function":"FindPrivateDataById","Args":["6229486603953201814"]}'
#  现金和票据交易 （票据提现）
peer chaincode invoke -o orderer1.org0.example.com:7050 --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE -C $CHANNEL_NAME -n financial --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  -c '{"function":"TransferAsset","Args":["6229486603953201814","102"]}' --waitForEvent
# 查询默认私有数据 ----票据变化 103--1，现金余额变化 102--203
peer chaincode query -C $CHANNEL_NAME -n financial_general_account   -c '{"function":"FindPrivateDataById","Args":["6229486603953201814"]}'


exit

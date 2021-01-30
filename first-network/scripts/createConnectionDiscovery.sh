#
# discover cli  服务发现命令行工具
# 发现服务是为了获取网络中最新的状态，使得客户端sdk能知道最新链码的背书策略，网络中各组织节点的endpoint等信息。
# https://hyperledger-fabric.readthedocs.io/en/release-2.1/discovery-cli.html

docker exec -it cli-org1-peer0 bash

# 生成配置文件
# 此处使用first-network做实验，根据实际情况设置环境变量
export USERKEY=/usr/local/home/org1/admin.org1.example.com/msp/keystore/key.pem

export USERCERT=/usr/local/home/org1/admin.org1.example.com/msp/signcerts/cert.pem

export PEERTLSCA=/usr/local/home/org1/peer0.org1.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem
# 可选
export TLSCERT=/usr/local/home/org1/peer0.org1.example.com/tls-msp/signcerts/cert.pem
# 可选
export TLSKEY=/usr/local/home/org1/peer0.org1.example.com/tls-msp/keystore/key.pem

# 生成配置文件
discover --configFile conf.yaml --peerTLSCA $TLSCERT --userKey $USERKEY --userCert $USERCERT  --MSP Org1MSP saveConfig
#discover --configFile conf.yaml --peerTLSCA $TLSCERT --tlsKey=$TLSKEY --tlsCert=$TLSCERT --userKey $USERKEY --userCert $USERCERT  --MSP Org1MSP saveConfig

# 查询peer membership
discover --configFile conf.yaml peers --channel mychannel  --server peer0.org1.example.com:7051
# 背书策略查询
# 查询背书需要指定参数
# --chaincode参数是必须的，指定了链码的名称
# --collection用于指定链码希望使用的私有数据collection
# discover --configFile conf.yaml endorsers --channel mychannel  --server peer0.org1.example.com:7051 --chaincode=agency --chaincode=financial --collection=agency:financial




## 官方文档：https://hyperledger-fabric.readthedocs.io/zh_CN/release-2.2/private_data_tutorial.html

## 生成组织节点通道配置
configtxgen -profile SampleMultiNodeEtcdRaft -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID mychannel -asOrg Org2MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors.tx -channelID mychannel -asOrg Org3MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org4MSPanchors.tx -channelID mychannel -asOrg Org4MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org5MSPanchors.tx -channelID mychannel -asOrg Org5MSP
#configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org6MSPanchors.tx -channelID mychannel -asOrg Org6MSP

## 启动节点服务
docker-compose -f docker-compose-etcdraft2.yaml up -d 2>&1


## 启动cli服务
docker-compose -f docker-compose-cli-peers.yaml up  -d 2>&1


docker exec -it cli-org1-peer0 bash
# 创建通道
peer channel create -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/channel.tx --outputBlock /usr/local/channel-artifacts/$CHANNEL_NAME.block --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE

exit


docker exec -it cli-org1-peer0 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block
# 更新锚节点 只需要当前组织中的任意一个节点更新即可
peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org1MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE
exit

docker exec -it cli-org1-peer1 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block


exit

docker exec -it cli-org2-peer0 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block
# 更新锚节点 只需要当前组织中的任意一个节点更新即可
peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org2MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE
exit

docker exec -it cli-org2-peer1 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

docker exec -it cli-org3-peer0 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block
# 更新锚节点 只需要当前组织中的任意一个节点更新即可
peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org3MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE
exit

docker exec -it cli-org3-peer1 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

docker exec -it cli-org4-peer0 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block
# 更新锚节点 只需要当前组织中的任意一个节点更新即可
peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org4MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE
exit

docker exec -it cli-org4-peer1 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

docker exec -it cli-org5-peer0 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block
# 更新锚节点 只需要当前组织中的任意一个节点更新即可
peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org5MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE
exit

docker exec -it cli-org5-peer1 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit

docker exec -it cli-org6-peer0 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block
# 更新锚节点 只需要当前组织中的任意一个节点更新即可
peer channel update -o orderer1.org0.example.com:7050 -c $CHANNEL_NAME -f /usr/local/channel-artifacts/Org6MSPanchors.tx --tls true --cafile $CORE_PEER_TLS_ROOTCERT_FILE
exit

docker exec -it cli-org6-peer1 bash

# 加入通道
peer channel join -b /usr/local/channel-artifacts/$CHANNEL_NAME.block

exit


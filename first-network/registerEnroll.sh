#!/bin/bash
## 工作目录,除非特殊说明，一般命令的执行都是在工作目录进行。
## http://localhost:5984/_utils/

## 首次运行清除docker影响

docker-compose -f docker-compose-ca.yaml  down --volumes --remove-orphans

## docker ps -a|awk '{print $1}'|xargs -i docker stop {}
## docker ps -a|awk '{print $1}'|xargs -i docker rm {}

## docker rm -f $(docker ps -a | awk '($2 ~ /dev-peer.*/) {print $1}')
## docker rmi -f $(docker images | awk '($1 ~ /dev-peer.*.mycc/) {print $3}')
## docker rmi -f $(docker images | awk '($1 ~ /dev-peer.*.marbles02_private/) {print $3}')
## docker volume prune

## 启动CA服务
docker-compose -f docker-compose-ca.yaml up -d 2>&1


## 登陆tls
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/tls-ca/admin
fabric-ca-client enroll -d -u https://tls-ca-admin:tls-ca-adminpw@0.0.0.0:7052

## 注册tls
# 为各个组件(组件指的是 peer,order和管理员)注册TLS证书,仅仅是注册了身份,并没有获取到证书;
fabric-ca-client register -d --id.name orderer1.org0.example.com --id.secret ordererPW --id.type orderer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name orderer2.org0.example.com --id.secret ordererPW --id.type orderer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name orderer3.org0.example.com --id.secret ordererPW --id.type orderer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name orderer4.org0.example.com --id.secret ordererPW --id.type orderer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name orderer5.org0.example.com --id.secret ordererPW --id.type orderer -u https://0.0.0.0:7052

fabric-ca-client register -d --id.name peer0.org1.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer1.org1.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer0.org2.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer1.org2.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer0.org3.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer1.org3.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer0.org4.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer1.org4.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer0.org5.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer1.org5.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer0.org6.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name peer1.org6.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7052

fabric-ca-client register -d --id.name admin.org0.example.com --id.secret org0AdminPW --id.type admin -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name admin.org1.example.com --id.secret org1AdminPW --id.type admin -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name admin.org2.example.com --id.secret org2AdminPW --id.type admin -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name admin.org3.example.com --id.secret org3AdminPW --id.type admin -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name admin.org4.example.com --id.secret org4AdminPW --id.type admin -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name admin.org5.example.com --id.secret org5AdminPW --id.type admin -u https://0.0.0.0:7052
fabric-ca-client register -d --id.name admin.org6.example.com --id.secret org6AdminPW --id.type admin -u https://0.0.0.0:7052


## 登录org0
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org0/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/ca/admin
fabric-ca-client enroll -d -u https://org0-admin:org0-adminpw@0.0.0.0:7053

## 注册org0
fabric-ca-client register -d --id.name orderer1.org0.example.com --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' -u https://0.0.0.0:7053
fabric-ca-client register -d --id.name orderer2.org0.example.com --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' -u https://0.0.0.0:7053
fabric-ca-client register -d --id.name orderer3.org0.example.com --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' -u https://0.0.0.0:7053
fabric-ca-client register -d --id.name orderer4.org0.example.com --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' -u https://0.0.0.0:7053
fabric-ca-client register -d --id.name orderer5.org0.example.com --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' -u https://0.0.0.0:7053
fabric-ca-client register -d --id.name admin.org0.example.com --id.secret org0adminpw --id.type admin --id.attrs "hf.Registrar.Roles=client,hf.Registrar.Attributes=*,hf.Revoker=true,hf.GenCRL=true,admin=true:ecert,abac.init=true:ecert" -u https://0.0.0.0:7053


## 登录org1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org1/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org1/ca/admin
fabric-ca-client enroll -d -u https://org1-admin:org1-adminpw@0.0.0.0:7054
## 注册org1
fabric-ca-client register -d --id.name peer0.org1.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7054
fabric-ca-client register -d --id.name peer1.org1.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7054
fabric-ca-client register -d --id.name admin.org1.example.com --id.secret org1AdminPW --id.type admin -u https://0.0.0.0:7054


## 登录org2
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org2/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org2/ca/admin
fabric-ca-client enroll -d -u https://org2-admin:org2-adminpw@0.0.0.0:7055
## 注册org2
fabric-ca-client register -d --id.name peer0.org2.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7055
fabric-ca-client register -d --id.name peer1.org2.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7055
fabric-ca-client register -d --id.name admin.org2.example.com --id.secret org2AdminPW --id.type admin -u https://0.0.0.0:7055


## 登录org3
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org3/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org3/ca/admin
fabric-ca-client enroll -d -u https://org3-admin:org3-adminpw@0.0.0.0:7056
## 注册org3
fabric-ca-client register -d --id.name peer0.org3.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7056
fabric-ca-client register -d --id.name peer1.org3.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7056
fabric-ca-client register -d --id.name admin.org3.example.com --id.secret org3AdminPW --id.type admin -u https://0.0.0.0:7056


## 登录org4
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org4/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org4/ca/admin
fabric-ca-client enroll -d -u https://org4-admin:org4-adminpw@0.0.0.0:7057
## 注册org4
fabric-ca-client register -d --id.name peer0.org4.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7057
fabric-ca-client register -d --id.name peer1.org4.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7057
fabric-ca-client register -d --id.name admin.org4.example.com --id.secret org4AdminPW --id.type admin -u https://0.0.0.0:7057


## 登录org5
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org5/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org5/ca/admin
fabric-ca-client enroll -d -u https://org5-admin:org5-adminpw@0.0.0.0:7058
## 注册org5
fabric-ca-client register -d --id.name peer0.org5.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7058
fabric-ca-client register -d --id.name peer1.org5.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7058
fabric-ca-client register -d --id.name admin.org5.example.com --id.secret org5AdminPW --id.type admin -u https://0.0.0.0:7058


## 登录org6
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org6/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org6/ca/admin
fabric-ca-client enroll -d -u https://org6-admin:org6-adminpw@0.0.0.0:7059
## 注册org6
fabric-ca-client register -d --id.name peer0.org6.example.com --id.secret peer1PW --id.type peer -u https://0.0.0.0:7059
fabric-ca-client register -d --id.name peer1.org6.example.com --id.secret peer2PW --id.type peer -u https://0.0.0.0:7059
fabric-ca-client register -d --id.name admin.org6.example.com --id.secret org6AdminPW --id.type admin -u https://0.0.0.0:7059


## 通过客户端为所有节点生成msp证书和tls证书


## org1内的peer节点制作证书

export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org1/admin.org1.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org1/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin.org1.example.com:org1AdminPW@0.0.0.0:7054


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org1/admin.org1.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://admin.org1.example.com:org1AdminPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts admin.org1.example.com

mv $PWD/crypto-config/org1/admin.org1.example.com/msp/keystore/*_sk $PWD/crypto-config/org1/admin.org1.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org1/admin.org1.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org1/admin.org1.example.com/tls-msp/keystore/key.pem



export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org1/peer0.org1.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org1/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer0.org1.example.com:peer1PW@0.0.0.0:7054


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org1/peer0.org1.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer0.org1.example.com:peer1PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer0.org1.example.com

mv $PWD/crypto-config/org1/peer0.org1.example.com/msp/keystore/*_sk $PWD/crypto-config/org1/peer0.org1.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org1/peer0.org1.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org1/peer0.org1.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org1/peer0.org1.example.com/msp/admincerts
cp $PWD/crypto-config/org1/admin.org1.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org1/peer0.org1.example.com/msp/admincerts/org1-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org1/peer1.org1.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org1/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1.org1.example.com:peer2PW@0.0.0.0:7054


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org1/peer1.org1.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer1.org1.example.com:peer2PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer1.org1.example.com

mv $PWD/crypto-config/org1/peer1.org1.example.com/msp/keystore/*_sk  $PWD/crypto-config/org1/peer1.org1.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org1/peer1.org1.example.com/tls-msp/keystore/*_sk  $PWD/crypto-config/org1/peer1.org1.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org1/peer1.org1.example.com/msp/admincerts
cp $PWD/crypto-config/org1/admin.org1.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org1/peer1.org1.example.com/msp/admincerts/org1-admin-cert.pem


## org2内的peer节点制作证书

export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org2/admin.org2.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org2/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin.org2.example.com:org2AdminPW@0.0.0.0:7055


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org2/admin.org2.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://admin.org2.example.com:org2AdminPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts admin.org2.example.com

mv $PWD/crypto-config/org2/admin.org2.example.com/msp/keystore/*_sk $PWD/crypto-config/org2/admin.org2.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org2/admin.org2.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org2/admin.org2.example.com/tls-msp/keystore/key.pem



export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org2/peer0.org2.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org2/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer0.org2.example.com:peer1PW@0.0.0.0:7055


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org2/peer0.org2.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer0.org2.example.com:peer1PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer0.org2.example.com

mv $PWD/crypto-config/org2/peer0.org2.example.com/msp/keystore/*_sk $PWD/crypto-config/org2/peer0.org2.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org2/peer0.org2.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org2/peer0.org2.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org2/peer0.org2.example.com/msp/admincerts
cp $PWD/crypto-config/org2/admin.org2.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org2/peer0.org2.example.com/msp/admincerts/org2-admin-cert.pem



export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org2/peer1.org2.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org2/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1.org2.example.com:peer2PW@0.0.0.0:7055


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org2/peer1.org2.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer1.org2.example.com:peer2PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer1.org2.example.com

mv $PWD/crypto-config/org2/peer1.org2.example.com/msp/keystore/*_sk $PWD/crypto-config/org2/peer1.org2.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org2/peer1.org2.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org2/peer1.org2.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org2/peer1.org2.example.com/msp/admincerts
cp $PWD/crypto-config/org2/admin.org2.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org2/peer1.org2.example.com/msp/admincerts/org2-admin-cert.pem





## org3内的peer节点制作证书

export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org3/admin.org3.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org3/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin.org3.example.com:org3AdminPW@0.0.0.0:7056


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org3/admin.org3.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://admin.org3.example.com:org3AdminPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts admin.org3.example.com

mv $PWD/crypto-config/org3/admin.org3.example.com/msp/keystore/*_sk $PWD/crypto-config/org3/admin.org3.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org3/admin.org3.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org3/admin.org3.example.com/tls-msp/keystore/key.pem



export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org3/peer0.org3.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org3/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer0.org3.example.com:peer1PW@0.0.0.0:7056


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org3/peer0.org3.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer0.org3.example.com:peer1PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer0.org3.example.com

mv $PWD/crypto-config/org3/peer0.org3.example.com/msp/keystore/*_sk $PWD/crypto-config/org3/peer0.org3.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org3/peer0.org3.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org3/peer0.org3.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org3/peer0.org3.example.com/msp/admincerts
cp $PWD/crypto-config/org3/admin.org3.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org3/peer0.org3.example.com/msp/admincerts/org3-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org3/peer1.org3.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org3/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1.org3.example.com:peer2PW@0.0.0.0:7056


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org3/peer1.org3.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer1.org3.example.com:peer2PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer1.org3.example.com

mv $PWD/crypto-config/org3/peer1.org3.example.com/msp/keystore/*_sk  $PWD/crypto-config/org3/peer1.org3.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org3/peer1.org3.example.com/tls-msp/keystore/*_sk  $PWD/crypto-config/org3/peer1.org3.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org3/peer1.org3.example.com/msp/admincerts
cp $PWD/crypto-config/org3/admin.org3.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org3/peer1.org3.example.com/msp/admincerts/org3-admin-cert.pem




## org4内的peer节点制作证书

export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org4/admin.org4.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org4/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin.org4.example.com:org4AdminPW@0.0.0.0:7057


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org4/admin.org4.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://admin.org4.example.com:org4AdminPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts admin.org4.example.com

mv $PWD/crypto-config/org4/admin.org4.example.com/msp/keystore/*_sk $PWD/crypto-config/org4/admin.org4.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org4/admin.org4.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org4/admin.org4.example.com/tls-msp/keystore/key.pem



export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org4/peer0.org4.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org4/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer0.org4.example.com:peer1PW@0.0.0.0:7057


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org4/peer0.org4.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer0.org4.example.com:peer1PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer0.org4.example.com

mv $PWD/crypto-config/org4/peer0.org4.example.com/msp/keystore/*_sk $PWD/crypto-config/org4/peer0.org4.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org4/peer0.org4.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org4/peer0.org4.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org4/peer0.org4.example.com/msp/admincerts
cp $PWD/crypto-config/org4/admin.org4.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org4/peer0.org4.example.com/msp/admincerts/org4-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org4/peer1.org4.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org4/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1.org4.example.com:peer2PW@0.0.0.0:7057


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org4/peer1.org4.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer1.org4.example.com:peer2PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer1.org4.example.com

mv $PWD/crypto-config/org4/peer1.org4.example.com/msp/keystore/*_sk  $PWD/crypto-config/org4/peer1.org4.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org4/peer1.org4.example.com/tls-msp/keystore/*_sk  $PWD/crypto-config/org4/peer1.org4.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org4/peer1.org4.example.com/msp/admincerts
cp $PWD/crypto-config/org4/admin.org4.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org4/peer1.org4.example.com/msp/admincerts/org4-admin-cert.pem




## org5内的peer节点制作证书

export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org5/admin.org5.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org5/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin.org5.example.com:org5AdminPW@0.0.0.0:7058


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org5/admin.org5.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://admin.org5.example.com:org5AdminPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts admin.org5.example.com

mv $PWD/crypto-config/org5/admin.org5.example.com/msp/keystore/*_sk $PWD/crypto-config/org5/admin.org5.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org5/admin.org5.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org5/admin.org5.example.com/tls-msp/keystore/key.pem



export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org5/peer0.org5.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org5/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer0.org5.example.com:peer1PW@0.0.0.0:7058


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org5/peer0.org5.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer0.org5.example.com:peer1PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer0.org5.example.com

mv $PWD/crypto-config/org5/peer0.org5.example.com/msp/keystore/*_sk $PWD/crypto-config/org5/peer0.org5.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org5/peer0.org5.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org5/peer0.org5.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org5/peer0.org5.example.com/msp/admincerts
cp $PWD/crypto-config/org5/admin.org5.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org5/peer0.org5.example.com/msp/admincerts/org5-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org5/peer1.org5.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org5/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1.org5.example.com:peer2PW@0.0.0.0:7058


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org5/peer1.org5.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer1.org5.example.com:peer2PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer1.org5.example.com

mv $PWD/crypto-config/org5/peer1.org5.example.com/msp/keystore/*_sk  $PWD/crypto-config/org5/peer1.org5.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org5/peer1.org5.example.com/tls-msp/keystore/*_sk  $PWD/crypto-config/org5/peer1.org5.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org5/peer1.org5.example.com/msp/admincerts
cp $PWD/crypto-config/org5/admin.org5.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org5/peer1.org5.example.com/msp/admincerts/org5-admin-cert.pem




## org6内的peer节点制作证书

export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org6/admin.org6.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org6/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin.org6.example.com:org6AdminPW@0.0.0.0:7059


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org6/admin.org6.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://admin.org6.example.com:org6AdminPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts admin.org6.example.com

mv $PWD/crypto-config/org6/admin.org6.example.com/msp/keystore/*_sk $PWD/crypto-config/org6/admin.org6.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org6/admin.org6.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org6/admin.org6.example.com/tls-msp/keystore/key.pem



export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org6/peer0.org6.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org6/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer0.org6.example.com:peer1PW@0.0.0.0:7059


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org6/peer0.org6.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer0.org6.example.com:peer1PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer0.org6.example.com

mv $PWD/crypto-config/org6/peer0.org6.example.com/msp/keystore/*_sk $PWD/crypto-config/org6/peer0.org6.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org6/peer0.org6.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org6/peer0.org6.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org6/peer0.org6.example.com/msp/admincerts
cp $PWD/crypto-config/org6/admin.org6.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org6/peer0.org6.example.com/msp/admincerts/org6-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org6/peer1.org6.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org6/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1.org6.example.com:peer2PW@0.0.0.0:7059


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org6/peer1.org6.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://peer1.org6.example.com:peer2PW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer1.org6.example.com

mv $PWD/crypto-config/org6/peer1.org6.example.com/msp/keystore/*_sk  $PWD/crypto-config/org6/peer1.org6.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org6/peer1.org6.example.com/tls-msp/keystore/*_sk  $PWD/crypto-config/org6/peer1.org6.example.com/tls-msp/keystore/key.pem
mkdir -p $PWD/crypto-config/org6/peer1.org6.example.com/msp/admincerts
cp $PWD/crypto-config/org6/admin.org6.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org6/peer1.org6.example.com/msp/admincerts/org6-admin-cert.pem



## org0内的orderer节点制作证书

export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/admin.org0.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org0/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin.org0.example.com:org0adminpw@0.0.0.0:7053


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/admin.org0.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
fabric-ca-client enroll -d -u https://admin.org0.example.com:org0AdminPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts admin.org0.example.com

mv $PWD/crypto-config/org0/admin.org0.example.com/msp/keystore/*_sk $PWD/crypto-config/org0/admin.org0.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org0/admin.org0.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org0/admin.org0.example.com/tls-msp/keystore/key.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer1.org0.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org0/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer1.org0.example.com:ordererpw@0.0.0.0:7053


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer1.org0.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://orderer1.org0.example.com:ordererPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts orderer1.org0.example.com

mv $PWD/crypto-config/org0/orderer1.org0.example.com/msp/keystore/*_sk $PWD/crypto-config/org0/orderer1.org0.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org0/orderer1.org0.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org0/orderer1.org0.example.com/tls-msp/keystore/key.pem
mkdir $PWD/crypto-config/org0/orderer1.org0.example.com/msp/admincerts
cp $PWD/crypto-config/org0/admin.org0.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org0/orderer1.org0.example.com/msp/admincerts/orderer-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer2.org0.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org0/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer2.org0.example.com:ordererpw@0.0.0.0:7053


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer2.org0.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://orderer2.org0.example.com:ordererPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts orderer2.org0.example.com

mv $PWD/crypto-config/org0/orderer2.org0.example.com/msp/keystore/*_sk $PWD/crypto-config/org0/orderer2.org0.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org0/orderer2.org0.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org0/orderer2.org0.example.com/tls-msp/keystore/key.pem
mkdir $PWD/crypto-config/org0/orderer2.org0.example.com/msp/admincerts
cp $PWD/crypto-config/org0/admin.org0.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org0/orderer2.org0.example.com/msp/admincerts/orderer-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer3.org0.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org0/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer3.org0.example.com:ordererpw@0.0.0.0:7053


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer3.org0.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://orderer3.org0.example.com:ordererPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts orderer3.org0.example.com

mv $PWD/crypto-config/org0/orderer3.org0.example.com/msp/keystore/*_sk $PWD/crypto-config/org0/orderer3.org0.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org0/orderer3.org0.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org0/orderer3.org0.example.com/tls-msp/keystore/key.pem
mkdir $PWD/crypto-config/org0/orderer3.org0.example.com/msp/admincerts
cp $PWD/crypto-config/org0/admin.org0.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org0/orderer3.org0.example.com/msp/admincerts/orderer-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer4.org0.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org0/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer4.org0.example.com:ordererpw@0.0.0.0:7053


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer4.org0.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://orderer4.org0.example.com:ordererPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts orderer4.org0.example.com

mv $PWD/crypto-config/org0/orderer4.org0.example.com/msp/keystore/*_sk $PWD/crypto-config/org0/orderer4.org0.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org0/orderer4.org0.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org0/orderer4.org0.example.com/tls-msp/keystore/key.pem
mkdir $PWD/crypto-config/org0/orderer4.org0.example.com/msp/admincerts
cp $PWD/crypto-config/org0/admin.org0.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org0/orderer4.org0.example.com/msp/admincerts/orderer-admin-cert.pem




export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer5.org0.example.com
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/org0/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer5.org0.example.com:ordererpw@0.0.0.0:7053


export FABRIC_CA_CLIENT_HOME=$PWD/crypto-config/org0/orderer5.org0.example.com
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$PWD/crypto-config/tls-ca/crypto/ca-cert.pem
fabric-ca-client enroll -d -u https://orderer5.org0.example.com:ordererPW@0.0.0.0:7052 --enrollment.profile tls --csr.hosts orderer5.org0.example.com

mv $PWD/crypto-config/org0/orderer5.org0.example.com/msp/keystore/*_sk $PWD/crypto-config/org0/orderer5.org0.example.com/msp/keystore/key.pem
mv $PWD/crypto-config/org0/orderer5.org0.example.com/tls-msp/keystore/*_sk $PWD/crypto-config/org0/orderer5.org0.example.com/tls-msp/keystore/key.pem
mkdir $PWD/crypto-config/org0/orderer5.org0.example.com/msp/admincerts
cp $PWD/crypto-config/org0/admin.org0.example.com/msp/signcerts/cert.pem $PWD/crypto-config/org0/orderer5.org0.example.com/msp/admincerts/orderer-admin-cert.pem



cp ./orderer-config/config.yaml ./crypto-config/org0/admin.org0.example.com/msp/
cp ./orderer-config/config.yaml ./crypto-config/org0/orderer1.org0.example.com/msp/
cp ./orderer-config/config.yaml ./crypto-config/org0/orderer2.org0.example.com/msp/
cp ./orderer-config/config.yaml ./crypto-config/org0/orderer3.org0.example.com/msp/
cp ./orderer-config/config.yaml ./crypto-config/org0/orderer4.org0.example.com/msp/
cp ./orderer-config/config.yaml ./crypto-config/org0/orderer5.org0.example.com/msp/

cp ./org1-config/config.yaml ./crypto-config/org1/admin.org1.example.com/msp/
cp ./org1-config/config.yaml ./crypto-config/org1/peer0.org1.example.com/msp/
cp ./org1-config/config.yaml ./crypto-config/org1/peer1.org1.example.com/msp/

cp ./org2-config/config.yaml ./crypto-config/org2/admin.org2.example.com/msp/
cp ./org2-config/config.yaml ./crypto-config/org2/peer0.org2.example.com/msp/
cp ./org2-config/config.yaml ./crypto-config/org2/peer1.org2.example.com/msp/

cp ./org3-config/config.yaml ./crypto-config/org3/admin.org3.example.com/msp/
cp ./org3-config/config.yaml ./crypto-config/org3/peer0.org3.example.com/msp/
cp ./org3-config/config.yaml ./crypto-config/org3/peer1.org3.example.com/msp/

cp ./org4-config/config.yaml ./crypto-config/org4/admin.org4.example.com/msp/
cp ./org4-config/config.yaml ./crypto-config/org4/peer0.org4.example.com/msp/
cp ./org4-config/config.yaml ./crypto-config/org4/peer1.org4.example.com/msp/

cp ./org5-config/config.yaml ./crypto-config/org5/admin.org5.example.com/msp/
cp ./org5-config/config.yaml ./crypto-config/org5/peer0.org5.example.com/msp/
cp ./org5-config/config.yaml ./crypto-config/org5/peer1.org5.example.com/msp/

cp ./org6-config/config.yaml ./crypto-config/org6/admin.org6.example.com/msp/
cp ./org6-config/config.yaml ./crypto-config/org6/peer0.org6.example.com/msp/
cp ./org6-config/config.yaml ./crypto-config/org6/peer1.org6.example.com/msp/


## org0的admin证书
mkdir $PWD/crypto-config/org0/admin.org0.example.com/msp/tlscacerts
cp  $PWD/crypto-config/org0/admin.org0.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem  $PWD/crypto-config/org0/admin.org0.example.com/msp/tlscacerts

## org1的admin证书
mkdir -p $PWD/crypto-config/org1/admin.org1.example.com/msp/tlscacerts
cp $PWD/crypto-config/org1/admin.org1.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem $PWD/crypto-config/org1/admin.org1.example.com/msp/tlscacerts

## org2的admin证书
mkdir -p $PWD/crypto-config/org2/admin.org2.example.com/msp/tlscacerts
cp $PWD/crypto-config/org2/admin.org2.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem $PWD/crypto-config/org2/admin.org2.example.com/msp/tlscacerts

## org3的admin证书
mkdir -p $PWD/crypto-config/org3/admin.org3.example.com/msp/tlscacerts
cp $PWD/crypto-config/org3/admin.org3.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem $PWD/crypto-config/org3/admin.org3.example.com/msp/tlscacerts

## org4的admin证书
mkdir -p $PWD/crypto-config/org4/admin.org4.example.com/msp/tlscacerts
cp $PWD/crypto-config/org4/admin.org4.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem $PWD/crypto-config/org4/admin.org4.example.com/msp/tlscacerts

## org5的admin证书
mkdir -p $PWD/crypto-config/org5/admin.org5.example.com/msp/tlscacerts
cp $PWD/crypto-config/org5/admin.org5.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem $PWD/crypto-config/org5/admin.org5.example.com/msp/tlscacerts

## org6的admin证书
mkdir -p $PWD/crypto-config/org6/admin.org6.example.com/msp/tlscacerts
cp $PWD/crypto-config/org6/admin.org6.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem $PWD/crypto-config/org6/admin.org6.example.com/msp/tlscacerts



configtxgen -profile SampleMultiNodeEtcdRaft -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID mychannel -asOrg Org2MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors.tx -channelID mychannel -asOrg Org3MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org4MSPanchors.tx -channelID mychannel -asOrg Org4MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org5MSPanchors.tx -channelID mychannel -asOrg Org5MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org6MSPanchors.tx -channelID mychannel -asOrg Org6MSP

##  docker-compose -f docker-compose-etcdraft2.yaml up -d 2>&1

## curl -X PUT http://localhost:5984/_users
## curl -X PUT http://localhost:5984/_replicator
## curl -X PUT http://localhost:5984/fabric__internal




## dream 梦想

go mod init

GO111MODULE=on go mod vendor

区块链

https://github.com/hyperledger/fabric

https://github.com/hyperledger/fabric-samples


组织机构定义：

|  **机构名称** | **机构名称英文** | **组织节点** |
| -------------|------------------------------|------------------|
| 平台（上帝/监管）机构 |platform | org0 | 
| 金融机构 |financial | org1 |
| 下发机构 | issue |org2 | 
| 代理商机构 | agency |org3 | 
| 商户机构 | merchant |org4 | 
| 个体 | individual|属于org0平台  | 


链码定义：

|  **链码名称** | **组织节点** | **链码描述** |  **链码名称** |
| -------------|------------------------------|------------------|------------------|
| platform_chaincode.go |org0 | 平台（上帝/监管）机构链码 | platformCC |
| individual_chaincode.go |org0 | 个体链码 | individualCC |
| distribution_record_chaincode.go |org0 | 派发记录链码 | distributionRecordCC |
| financial_chaincode.go |org1 | 金融机构链码 | financialCC |
| financial_managed_account_chaincode.go |org1 | 金融机构共管账户链码 | financialManagedAccountCC |
| financial_general_account_chaincode.go |org1 | 金融机构一般账户链码 | financialGeneralAccountCC |
| issue_chaincode.go |org2 | 下发机构链码 | issueCC |
| agency_chaincode.go |org3 | 代理商机构链码 | agencyCC |
| merchant_chaincode.go |org4 | 商户机构链码 | merchantCC |


私有数据集定义：

|  **私有数据集名称** | **组织节点** | **私有数据集描述** |  **包含节点** |
| -------------|------------------------------|------------------|------------------|
| collectionFinancial |org1 | 金融机构私有数据集 | org0 + org1 |
| collectionIssue |org2 | 下发机构私有数据集 | org0 + org2 |
| collectionMerchant |org3 | 代理商机构私有数据集 | org0 + org3 |
| collectionAgency |org4 | 商户机构私有数据集 | org0 + org4 |

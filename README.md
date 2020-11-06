
## dream 梦想

go mod init

GO111MODULE=on go mod vendor

区块链

https://github.com/hyperledger/fabric

https://github.com/hyperledger/fabric-samples


现金与票据兑换比例为：现金1元=100票据，即：1分=1票据
交易单位为"分"


组织机构定义：

|  **机构名称** | **机构名称英文** | **组织节点** |
| -------------|------------------------------|------------------|
| 平台（上帝/监管）机构 |platform | org1 | 
| 金融机构 |financial | org2 |
| 下发机构 | issue |org3 | 
| 代理商机构 | agency |org4 | 
| 商户机构 | merchant |org5 | 
| 个体 | individual|属于org1平台  | 


链码定义：

|  **链码名称** | **组织节点** | **链码描述** |  **链码名称** |
| -------------|------------------------------|------------------|------------------|
| platform_chaincode.go |org1 | 平台（上帝/监管）机构链码 | platformCC |
| individual_chaincode.go |org1 | 个体链码 | individualCC |
| distribution_record_chaincode.go |org1 | 派发记录链码 | distributionRecordCC |
| financial_chaincode.go |org2 | 金融机构链码 | financialCC |
| financial_managed_account_chaincode.go |org2 | 金融机构共管账户链码 | financialManagedAccountCC |
| financial_general_account_chaincode.go |org2 | 金融机构一般账户链码 | financialGeneralAccountCC |
| issue_chaincode.go |org3 | 下发机构链码 | issueCC |
| agency_chaincode.go |org4 | 代理商机构链码 | agencyCC |
| merchant_chaincode.go |org5 | 商户机构链码 | merchantCC |


私有数据集定义：

|  **私有数据集名称** | **组织节点** | **私有数据集描述** |  **包含节点** |
| -------------|------------------------------|------------------|------------------|
| collectionFinancial |org2 | 金融机构私有数据集 | org1 + org2 |
| collectionIssue |org3 | 下发机构私有数据集 | org1 + org3 |
| collectionMerchant |org4 | 代理商机构私有数据集 | org1 + org4 |
| collectionAgency |org5 | 商户机构私有数据集 | org1 + org5 |


完整交易流程：

充值流程
-->商户向商户一般账户充值现金余额

|  **角色名称** | **现金** | **票据** |  
| -------------|---------|---------|
|     商户     |    100   |    0    | 
|     个体     |   0      |    0   | 
|     代理商机构|   0      |    0   | 
|     下发机构  |   0      |    0   | 
|     金融机构  |   0     |     0    | 

-->商户用一般账户的现金余额向上级代理的上级下发机构的金融机构的共管账户充值，获取金融机构颁发的票据，共管账户增加票据余额，商户减少一般账户的现金余额，增加金融机构的现金余额和票据余额。

|  **角色名称** | **现金** | **票据** |  
| -------------|---------|---------|
|     商户     |    0    |    0    | 
|     个体     |   0      |    0   | 
|     代理商机构|   0      |    0   | 
|     下发机构  |   0      |    0   | 
|     金融机构  |   100   |    100   | 

下发流程
-->商户发起下发请求（下发记录ID，下发使用共管账户卡号，个体ID，个体收款一般账户卡号，派发金额，派发费率）
-->减少共管账户卡号的票据余额，增加个体一般账户的票据余额，增加下发机构的一般账户的佣金票据余额，增加代理商机构的一般账户的佣金票据余额，金融机构的现金余额和票据余额不变。

|  **角色名称** | **现金** | **票据** |  
| -------------|---------|---------|
|     商户     |    0     |    0    | 
|     个体      |   0     |    90   | 
|     代理商机构 |   0     |    5    | 
|     下发机构  |    0     |    5    | 
|     金融机构  |   100    |    100   | 

提现流程
-->个体/代理商/下发机构发起提现请求
-->减少金融机构的现金余额和票据余额，增加个体/代理商/下发机构一般账户的现金余额，减少个体/代理商/下发机构一般账户的票据余额
个体（现金:90,票据:0）
代理商机构（现金:5,票据:0）
下发机构（现金:5,票据:0）
金融机构（现金:0,票据:0）

|  **角色名称** | **现金** | **票据** |  
| -------------|---------|---------|
|     商户     |    0    |    0    | 
|     个体     |   90     |    0   | 
|   代理商机构  |   5      |    0   | 
|     下发机构  |   5     |     0   | 
|     金融机构  |   0     |     0   | 

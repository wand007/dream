package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DistributionRecordChaincode struct {
	contractapi.Contract
}

const CHANNEL_NAME string = "mychannel"
const CHAINCODE_NAME_ISSUE_ORG string = "issueOrgCC"

/**
 派发记录属性
 */
type DistributionRecordPrivateData struct {
	ID            string  `json:"id"`            //派发记录ID
	IndividualID  string  `json:"individualID"`  //个体ID Individual.ID
	MerchantOrgID string  `json:"merchantOrgID"` //商户机构ID MerchantOrg.ID
	AgencyOrgID   string  `json:"merchantOrgID"` //代理商机构ID AgencyOrg.ID
	IssueOrgID    string  `json:"issueOrgID"`    //下发机构ID IssueOrg.ID
	Amount        int     `json:"amount"`        //派发金额
	Rate          float64 `json:"rate"`          //派发费率
	CardNo        string  `json:"cardNo"`        //金融机构公管账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	CardCode      string  `json:"code"`          //金融机构代码 FinancialOrg.Code
	Status        int     `json:"status"`        //派发状态(启用/禁用)
}

/**
  新增金融机构共管账户私有数据
 */
func (t *DistributionRecordChaincode) Create(ctx contractapi.TransactionContextInterface) (string, error) {
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}

	financialPrivateDataJsonBytes, ok := transMap["distributionRecord"]
	if !ok {
		return "", errors.New("financial must be a key in the transient map")
	}

	if len(financialPrivateDataJsonBytes) == 0 {
		return "", errors.New("financial value in the transient map must be a non-empty JSON string")
	}
	var transientInput DistributionRecordPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	if len(transientInput.ID) == 0 {
		return "", errors.New("派发记录ID不能为空")
	}
	if len(transientInput.IndividualID) == 0 {
		return "", errors.New("个体ID不能为空")
	}
	if len(transientInput.MerchantOrgID) == 0 {
		return "", errors.New("商户机构ID不能为空")
	}
	if len(transientInput.AgencyOrgID) == 0 {
		return "", errors.New("代理商机构ID不能为空")
	}
	if len(transientInput.IssueOrgID) == 0 {
		return "", errors.New("下发机构ID不能为空")
	}
	if len(transientInput.CardNo) == 0 {
		return "", errors.New("金融机构公管账户账号不能为空")
	}
	if len(transientInput.CardCode) == 0 {
		return "", errors.New("金融机构代码不能为空")
	}
	if transientInput.Amount == 0 {
		return "", errors.New("派发金额不能为0")
	}
	if transientInput.Rate == 0 {
		return "", errors.New("派发费率不能为0")
	}

	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetPrivateData("collectionDistributionRecord", transientInput.ID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.ID + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.ID + "\"}"
		return "", errors.New(jsonResp)
	}

	carAsBytes, _ := json.Marshal(transientInput)

	err = ctx.GetStub().PutPrivateData("collectionFinancialIssue", transientInput.ID, carAsBytes)
	if err != nil {
		return "", errors.New("商户共管账户保存失败" + err.Error())
	}

	//todo 转账

	return transientInput.ID, nil
}

func TransferAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, amount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转出一般账户卡号不能为空")
	}
	if amount < 0 {
		return "", errors.New("转账金额不能小于0")
	}
	trans := [][]byte{[]byte("TransferAsset"), []byte("managedCardNo"), []byte(managedCardNo), []byte("generalCardNo"), []byte(generalCardNo), []byte("amount"), []byte(string(amount))}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to TransferAsset chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return "", fmt.Errorf(errStr)
	}

	fmt.Printf("FindIssueOrgById chaincode successful. Got response %s", string(response.Payload))
	return string(response.Payload), nil
}

func (t *DistributionRecordChaincode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("共管账户id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData("collectionDistributionRecord", id)
	if err != nil {
		return "", errors.New("共管账户查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("共管账户数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(DistributionRecordChaincode))
	if err != nil {
		fmt.Printf("Error create PlatformChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting PlatformChainCode chaincode: %s", err.Error())
	}
}

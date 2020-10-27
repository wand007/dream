package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type FinancialManagedAccountChaincode struct {
	contractapi.Contract
}

/**
   金融机构共管账户私有数据属性
 */
type FinancialOrgManagedAccountPrivateData struct {
	CardNo                string `json:"cardNo"`                //金融机构共管账户账号
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	PlatformOrgID         string `json:"platformOrgID"`         //平台机构ID PlatformOrg.ID
	MerchantOrgID         string `json:"merchantOrgID"`         //商户机构ID MerchantOrg.ID
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

/**
  新增金融机构共管账户私有数据
 */
func (t *FinancialManagedAccountChaincode) Create(ctx contractapi.TransactionContextInterface) (string, error) {

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}

	financialPrivateDataJsonBytes, ok := transMap["financial"]
	if !ok {
		return "", errors.New("financial must be a key in the transient map")
	}

	if len(financialPrivateDataJsonBytes) == 0 {
		return "", errors.New("financial value in the transient map must be a non-empty JSON string")
	}
	var transientInput FinancialOrgManagedAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	if len(transientInput.CardNo) == 0 {
		return "", errors.New("金融机构共管账户账号不能为空")
	}
	if len(transientInput.FinancialOrgID) == 0 {
		return "", errors.New("金融机构ID不能为空")
	}
	if len(transientInput.PlatformOrgID) == 0 {
		return "", errors.New("金平台机构ID不能为空")
	}
	if len(transientInput.MerchantOrgID) == 0 {
		return "", errors.New("商户机构ID不能为空")
	}
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetPrivateData("collectionFinancialMerchantPlatform", transientInput.CardNo)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}
	// Mongo Query string 语法见上文链接
	queryString := fmt.Sprintf(`{"selector":{"cardNo":"%s"}}`, transientInput.CardNo)
	// 富查询的返回结果可能为多条 所以这里返回的是一个迭代器 需要我们进一步的处理来获取需要的结果
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionFinancialMerchantPlatform", queryString)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	if resultsIterator != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	carAsBytes, _ := json.Marshal(transientInput)

	err = ctx.GetStub().PutPrivateData("collectionFinancialMerchantPlatform", transientInput.CardNo, carAsBytes)
	if err != nil {
		return "", errors.New("商户共管账户保存失败" + err.Error())
	}
	return string(Avalbytes), nil
}

func (t *FinancialManagedAccountChaincode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("共管账户id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData("collectionFinancialMerchantPlatform", id)
	if err != nil {
		return "", errors.New("共管账户查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("共管账户数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

func (t *FinancialManagedAccountChaincode) Recharge(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, amount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转出一般账户卡号不能为空")
	}
	if amount < 0 {
		return "",errors.New("转账金额不能小于0")
	}
	bytes, err := ctx.GetStub().GetPrivateData("collectionFinancial", managedCardNo)
	if err != nil {
		return "", errors.New("共管账户查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("共管账户数据不存在，读到的%s对应的数据为空！", managedCardNo)
	}
	return string(bytes), nil
}

func (t *FinancialManagedAccountChaincode) TransferAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, voucherAmount int) error {
	if len(managedCardNo) == 0 {
		return errors.New("共管账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData("collectionFinancialMerchantPlatform", managedCardNo)
	if err != nil {
		return errors.New("共管账户查询失败！")
	}
	if financialPrivateDataJsonBytes == nil {
		return fmt.Errorf("共管账户数据不存在，读到的%s对应的数据为空！", managedCardNo)
	}
	var transientInput FinancialOrgManagedAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	newCurrentBalance := transientInput.VoucherCurrentBalance + voucherAmount
	if newCurrentBalance < 0 {
		return errors.New("共管账户票据余额不足")
	}
	transientInput.VoucherCurrentBalance = newCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutState(managedCardNo, assetJSON)
}

func (t *FinancialManagedAccountChaincode) TransferVoucherAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, voucherAmount int) error {
	if len(managedCardNo) == 0 {
		return errors.New("共管账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData("collectionFinancialMerchantPlatform", managedCardNo)
	if err != nil {
		return errors.New("共管账户查询失败！")
	}
	if financialPrivateDataJsonBytes == nil {
		return fmt.Errorf("共管账户数据不存在，读到的%s对应的数据为空！", managedCardNo)
	}
	var transientInput FinancialOrgManagedAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	newCurrentBalance := transientInput.VoucherCurrentBalance + voucherAmount
	if newCurrentBalance < 0 {
		return errors.New("共管账户票据余额不足")
	}
	transientInput.VoucherCurrentBalance = newCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutState(managedCardNo, assetJSON)
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(FinancialManagedAccountChaincode))
	if err != nil {
		fmt.Printf("Error create PlatformChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting PlatformChainCode chaincode: %s", err.Error())
	}
}

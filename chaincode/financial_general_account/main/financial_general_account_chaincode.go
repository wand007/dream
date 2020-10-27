package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type FinancialGeneralAccountChaincode struct {
	contractapi.Contract
}

/**
   金融机构一般账户私有数据属性
 */
type FinancialOrgGeneralAccountPrivateData struct {
	CardNo                string `json:"cardNo"`                //金融机构公管账户账号(唯一不重复)
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	CertificateNo         string `json:"certificateNo"`         //持卡者证件号
	CertificateType       string `json:"certificateType"`       //持卡者证件类型 (身份证/港澳台证/护照/军官证)
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

//func (t *FinancialGeneralAccountChaincode) Create(ctx contractapi.TransactionContextInterface) (string, error) {
//
//	transMap, err := ctx.GetStub().GetTransient()
//	if err != nil {
//		return "", errors.New("Error getting transient: " + err.Error())
//	}
//
//	financialPrivateDataJsonBytes, ok := transMap["financial"]
//	if !ok {
//		return "", errors.New("financial must be a key in the transient map")
//	}
//
//	if len(financialPrivateDataJsonBytes) == 0 {
//		return "", errors.New("financial value in the transient map must be a non-empty JSON string")
//	}
//	var transientInput FinancialOrgGeneralAccountPrivateData
//	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
//	if err != nil {
//		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
//	}
//	if len(transientInput.CardNo) == 0 {
//		return "", errors.New("金融机构共管账户账号不能为空")
//	}
//	if len(transientInput.FinancialOrgID) == 0 {
//		return "", errors.New("金融机构ID不能为空")
//	}
//	if len(transientInput.CertificateNo) == 0 {
//		return "", errors.New("持卡者证件号不能为空")
//	}
//	if len(transientInput.CertificateType) == 0 {
//		return "", errors.New("持卡者证件类型不能为空")
//	}
//	// Get the state from the ledger
//	Avalbytes, err := ctx.GetStub().GetPrivateData("collectionFinancialMerchantPlatform", transientInput.CardNo)
//	if err != nil {
//		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
//		return "", errors.New(jsonResp)
//	}
//
//	if Avalbytes != nil {
//		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
//		return "", errors.New(jsonResp)
//	}
//	// Mongo Query string 语法见上文链接
//	queryString := fmt.Sprintf(`{"selector":{"cardNo":"%s"}}`, transientInput.CardNo)
//	// 富查询的返回结果可能为多条 所以这里返回的是一个迭代器 需要我们进一步的处理来获取需要的结果
//	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionFinancialMerchantPlatform", queryString)
//
//	if err != nil {
//		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
//		return "", errors.New(jsonResp)
//	}
//
//	if resultsIterator != nil {
//		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
//		return "", errors.New(jsonResp)
//	}
//
//	carAsBytes, _ := json.Marshal(transientInput)
//
//	err = ctx.GetStub().PutPrivateData("collectionFinancialMerchantPlatform", transientInput.CardNo, carAsBytes)
//	if err != nil {
//		return "", errors.New("商户共管账户保存失败" + err.Error())
//	}
//	return string(Avalbytes), nil
//}

/**
 新增金融机构一般账户私有数据
 */
func (t *FinancialGeneralAccountChaincode) Create(ctx contractapi.TransactionContextInterface, collectionName string) (string, error) {

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}

	financialPrivateDataJsonBytes, ok := transMap["general"]
	if !ok {
		return "", errors.New("financial must be a key in the transient map")
	}

	if len(financialPrivateDataJsonBytes) == 0 {
		return "", errors.New("financial value in the transient map must be a non-empty JSON string")
	}
	var transientInput FinancialOrgGeneralAccountPrivateData
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
	if len(transientInput.CertificateNo) == 0 {
		return "", errors.New("持卡者证件号不能为空")
	}
	if len(transientInput.CertificateType) == 0 {
		return "", errors.New("持卡者证件类型不能为空")
	}
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetPrivateData(collectionName, transientInput.CardNo)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}
	carAsBytes, _ := json.Marshal(transientInput)

	err = ctx.GetStub().PutPrivateData(collectionName, transientInput.CardNo, carAsBytes)
	if err != nil {
		return "", errors.New("商户共管账户保存失败" + err.Error())
	}
	return "", nil
}

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToMerchant(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, "collectionFinancialMerchant")
}

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToPlatform(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, "collectionFinancialPlatform")
}

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToAgency(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, "collectionFinancialAgency")
}

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToIssue(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, "collectionFinancialIssue")
}

func (t *FinancialGeneralAccountChaincode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
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

func (t *FinancialGeneralAccountChaincode) Recharge(ctx contractapi.TransactionContextInterface, generalCardNo string, amount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	if amount < 0 {
		return errors.New("一般账户充值金额不能小于0")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData("collectionFinancial", generalCardNo)
	if err != nil {
		return errors.New("一般账户查询失败！")
	}
	if financialPrivateDataJsonBytes == nil {
		return fmt.Errorf("一般账户数据不存在，读到的%s对应的数据为空！", generalCardNo)
	}
	var transientInput FinancialOrgGeneralAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	newVoucherCurrentBalance := transientInput.VoucherCurrentBalance + amount
	if newVoucherCurrentBalance < 0 {
		return errors.New("一般账户现金余额不足")
	}
	transientInput.VoucherCurrentBalance = newVoucherCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutState(generalCardNo, assetJSON)
}

func (t *FinancialGeneralAccountChaincode) TransferAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, voucherAmount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData("collectionFinancialMerchantPlatform", generalCardNo)
	if err != nil {
		return errors.New("一般账户查询失败！")
	}
	if financialPrivateDataJsonBytes == nil {
		return fmt.Errorf("一般账户数据不存在，读到的%s对应的数据为空！", generalCardNo)
	}
	var transientInput FinancialOrgGeneralAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	newVoucherCurrentBalance := transientInput.VoucherCurrentBalance + voucherAmount
	if newVoucherCurrentBalance < 0 {
		return errors.New("一般账户票据余额不足")
	}
	transientInput.VoucherCurrentBalance = newVoucherCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutState(generalCardNo, assetJSON)
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(FinancialGeneralAccountChaincode))
	if err != nil {
		fmt.Printf("Error create PlatformChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting PlatformChainCode chaincode: %s", err.Error())
	}
}

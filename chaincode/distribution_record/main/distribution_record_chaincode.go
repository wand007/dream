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
const CHAINCODE_NAME_FINANCIAL_ORG string = "financialOrgCC"
const CHAINCODE_NAME_AGENCY_ORG string = "agencyOrgCC"
const CHAINCODE_NAME_MERCHANT_ORG string = "merchantOrgCC"

/**
 派发记录属性
 */
type DistributionRecordPrivateData struct {
	ID              string  `json:"id"`            //派发记录ID
	IndividualID    string  `json:"individualID"`  //个体ID Individual.ID
	MerchantOrgID   string  `json:"merchantOrgID"` //商户机构ID MerchantOrg.ID
	AgencyOrgID     string  `json:"merchantOrgID"` //代理商机构ID AgencyOrg.ID
	IssueOrgID      string  `json:"issueOrgID"`    //下发机构ID IssueOrg.ID
	Amount          int     `json:"amount"`        //派发金额
	Rate            float64 `json:"rate"`          //派发费率
	ManagedCardNo   string  `json:"managedCardNo"` //金融机构公管账户账号 FinancialOrgManagedAccountPrivateData.CardNo
	ManagedCardCode string  `json:"managedCode"`   //金融机构代码 FinancialOrg.Code
	GeneralCardNo   string  `json:"generalCardNo"` //金融机构公管账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	GeneralCardCode string  `json:"generalCode"`   //金融机构代码 FinancialOrg.Code
	Status          int     `json:"status"`        //派发状态(0:未下发/1:已下发)
}

/**
 代理商机构私有数据属性
 */
type AgencyOrgPrivateData struct {
	ID        string  `json:"id"`        //代理商机构ID
	RateBasic float64 `json:"rateBasic"` //代理商机构基础费率
}

/**
 商户机构属性
 */
type MerchantOrgPrivateData struct {
	ID        string  `json:"id"`        //商户机构ID
	RateBasic float64 `json:"rateBasic"` //下发机构基础费率
}

/**
  新增派发记录属性数据
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
	if len(transientInput.ManagedCardNo) == 0 {
		return "", errors.New("公管账户金融机构账号不能为空")
	}
	if len(transientInput.ManagedCardCode) == 0 {
		return "", errors.New("公管账户金融机构代码不能为空")
	}
	if len(transientInput.GeneralCardNo) == 0 {
		return "", errors.New("一般账户金融机构账号不能为空")
	}
	if len(transientInput.GeneralCardCode) == 0 {
		return "", errors.New("一般账户金融机构代码不能为空")
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

	err = ctx.GetStub().PutPrivateData("collectionDistributionRecord", transientInput.ID, carAsBytes)
	if err != nil {
		return "", errors.New("商户共管账户保存失败" + err.Error())
	}

	return transientInput.ID, nil
}

/**
  派发
 */
func (t *DistributionRecordChaincode) Trade(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("派发记录ID不能为空")
	}

	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetPrivateData("collectionDistributionRecord", id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
		return "", errors.New(jsonResp)
	}
	var distributionRecordPrivateData DistributionRecordPrivateData
	err = json.Unmarshal(Avalbytes, &distributionRecordPrivateData)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	// 共管账户向个人一般账户转账票据
	_, err = TransferVoucherAsset(ctx, distributionRecordPrivateData.ManagedCardNo, distributionRecordPrivateData.GeneralCardNo, distributionRecordPrivateData.Amount)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	//商户机构
	agencyOrgPrivateData, err := findAgencyPrivateDataById(ctx, distributionRecordPrivateData.AgencyOrgID)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if agencyOrgPrivateData == nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}

	//代理商机构
	merchantOrgPrivateData, err := findMerchantPrivateDataById(ctx, distributionRecordPrivateData.MerchantOrgID)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if merchantOrgPrivateData == nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	//代理商佣金
	merchantBrokerage, err := CalculationBrokerage(merchantOrgPrivateData.RateBasic, agencyOrgPrivateData.RateBasic, distributionRecordPrivateData.Amount)
	if err != nil {
		return "", err
	}
	// 共管账户向代理商一般账户转账佣金票据
	_, err = TransferVoucherAsset(ctx, distributionRecordPrivateData.ManagedCardNo, distributionRecordPrivateData.GeneralCardNo, merchantBrokerage)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	//商户佣金
	agencyBrokerage, err := CalculationBrokerage(agencyOrgPrivateData.RateBasic, distributionRecordPrivateData.Rate, distributionRecordPrivateData.Amount)
	if err != nil {
		return "", err
	}
	// 共管账户向商户一般账户转账佣金票据
	_, err = TransferVoucherAsset(ctx, distributionRecordPrivateData.ManagedCardNo, distributionRecordPrivateData.GeneralCardNo, agencyBrokerage)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	//修改下发记录状态
	distributionRecordPrivateData.Status = 1

	carAsBytes, _ := json.Marshal(distributionRecordPrivateData)
	err = ctx.GetStub().PutPrivateData("collectionFinancialIssue", id, carAsBytes)
	if err != nil {
		return "", errors.New("商户共管账户保存失败" + err.Error())
	}
	return distributionRecordPrivateData.ID, nil
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

func findAgencyPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (*AgencyOrgPrivateData, error) {
	if len(id) == 0 {
		return nil, errors.New("代理商机构ID不能为空")
	}
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(id)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_AGENCY_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return nil, fmt.Errorf(errStr)
	}

	managedAccountPrivateData := new(AgencyOrgPrivateData)
	err := json.Unmarshal(response.Payload, &managedAccountPrivateData)
	if err != nil {
		return nil, errors.New("Failed to decode JSON of: " + string(response.Payload))
	}

	return managedAccountPrivateData, nil
}

func findMerchantPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (*MerchantOrgPrivateData, error) {
	if len(id) == 0 {
		return nil, errors.New("代理商机构ID不能为空")
	}
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(id)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_MERCHANT_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return nil, fmt.Errorf(errStr)
	}

	managedAccountPrivateData := new(MerchantOrgPrivateData)
	err := json.Unmarshal(response.Payload, &managedAccountPrivateData)
	if err != nil {
		return nil, errors.New("Failed to decode JSON of: " + string(response.Payload))
	}

	return managedAccountPrivateData, nil
}

func TransferVoucherAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, amount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转出一般账户卡号不能为空")
	}
	if amount < 0 {
		return "", errors.New("转账金额不能小于0")
	}
	trans := [][]byte{[]byte("TransferVoucherAsset"), []byte("managedCardNo"), []byte(managedCardNo), []byte("generalCardNo"), []byte(generalCardNo), []byte("amount"), []byte(string(amount))}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_FINANCIAL_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to TransferAsset chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return "", fmt.Errorf(errStr)
	}

	fmt.Printf("FindIssueOrgById chaincode successful. Got response %s", string(response.Payload))
	return string(response.Payload), nil
}

/**
  计算佣金收入
 */
func CalculationBrokerage(rateBasic float64, rate float64, amount int) (int, error) {
	if rateBasic == 0 {
		return 0, errors.New("上级费率不能为0")
	}
	if rate == 0 {
		return 0, errors.New("下发费率不能为0")
	}
	if amount < 0 {
		return 0, errors.New("转账金额不能小于0")
	}
	brokerage := (rate - rateBasic) * float64(amount)

	fmt.Printf("CalculationBrokerage chaincode successful. Got response %s", brokerage)
	return int(brokerage), nil
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

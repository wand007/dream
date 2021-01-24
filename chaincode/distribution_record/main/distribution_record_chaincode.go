package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"strings"
)

/**
新建派发记录并派发
 */

/**
派发记录链码
 */
type DistributionRecordChaincode struct {
	contractapi.Contract
}

//通道名称
const CHANNEL_NAME string = "mychannel"

//链码名称
const (
	//金融机构链码
	CHAINCODE_NAME_FINANCIAL_ORG string = "financialCC"
	//金融机构共管账户链码
	CHAINCODE_NAME_FINANCIAL_MANAGED_ACCOUNT string = "financialManagedAccountCC"
	//金融机构一般账户链码
	CHAINCODE_NAME_FINANCIAL_GENERAL_ACCOUNT string = "financialGeneralAccountCC"
	//下发机构链码
	CHAINCODE_NAME_ISSUE_ORG string = "issueCC"
	//分销商机构链码
	CHAINCODE_NAME_AGENCY_ORG string = "agencyCC"
	//零售商机构链码
	CHAINCODE_NAME_RETAILER_ORG string = "retailerCC"
)

//私有数据集名称
const COLLECTION_DISTRIBUTION_RECORD string = "collectionDistributionRecord"

/**
 派发记录属性
 */
type DistributionRecordPrivateData struct {
	ID                   string  `json:"id"`                   //派发记录ID
	PlatformOrgID        string  `json:"platformOrgID"`        //平台机构ID PlatformOrg.ID
	FinancialOrgID       string  `json:"financialOrgID"`       //金融机构ID FinancialOrg.ID
	IndividualID         string  `json:"individualID"`         //个体ID Individual.ID
	RetailerOrgID        string  `json:"retailerOrgID"`        //零售商机构ID RetailerOrg.ID
	AgencyOrgID          string  `json:"agencyOrgID"`          //分销商机构ID AgencyOrg.ID
	IssueOrgID           string  `json:"issueOrgID"`           //下发机构ID IssueOrg.ID
	ManagedAccountCardNo string  `json:"managedAccountCardNo"` //共管账户账号 FinancialOrgManagedAccountPrivateData.CardNo
	IssueCardNo          string  `json:"issueCardNo"`          //下发机构一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	IndividualCardNo     string  `json:"individualCardNo"`     //个体一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	AgencyCardNo         string  `json:"agencyCardNo"`        //分销商机构一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	ManagedCardNo        string  `json:"managedCardNo"`        //金融机构公管账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	GeneralCardNo        string  `json:"generalCardNo"`        //金融机构公管账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	Amount               int     `json:"amount"`               //派发金额
	Rate                 float64 `json:"rate"`                 //派发费率
	Status               int     `json:"status"`               //派发状态(0:未下发/1:已下发)
}

/**
   金融机构共管账户私有数据属性
 */
type FinancialOrgManagedAccountPrivateData struct {
	CardNo                string `json:"cardNo"`                //金融机构共管账户账号
	PlatformOrgID         string `json:"platformOrgID"`         //平台机构ID PlatformOrg.ID
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	IssueOrgID            string `json:"issueOrgID"`            //下发机构ID IssueOrg.ID
	RetailerOrgID         string `json:"retailerOrgID"`         //零售商机构ID RetailerOrg.ID
	AgencyOrgID           string `json:"agencyOrgID"`           //分销商机构ID AgencyOrg.ID
	IssueCardNo           string `json:"issueCardNo"`           //下发机构一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	AgencyCardNo          string `json:"managedCardNo"`         //分销商机构一般账户账号 FinancialOrgManagedAccountPrivateData.CardNo
	ManagedCardNo         string `json:"managedCardNo"`         //分销商机构一般账户账号 FinancialOrgManagedAccountPrivateData.CardNo
	GeneralCardNo         string `json:"generalCardNo"`         //零售商机构一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构零售商机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

/**
   金融机构一般账户私有数据属性
 */
type FinancialOrgGeneralAccountPrivateData struct {
	CardNo                string `json:"cardNo"`                //金融机构公管账户账号(唯一不重复)
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	CertificateNo         string `json:"certificateNo"`         //持卡者证件号
	CertificateType       int    `json:"certificateType"`       //持卡者证件类型 (身份证/港澳台证/护照/军官证/统一社会信用代码)
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构零售商机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

/**
   下发机构私有数据属性
 */
type IssueOrgPrivateData struct {
	ID        string  `json:"id"`        //下发机构ID IssueOrg.ID
	RateBasic float64 `json:"rateBasic"` //下发机构基础费率
}

/**
 分销商机构私有数据属性
 */
type AgencyOrgPrivateData struct {
	ID        string  `json:"id"`        //分销商机构ID
	RateBasic float64 `json:"rateBasic"` //分销商机构基础费率
}

/**
 零售商机构属性
 */
type RetailerOrgPrivateData struct {
	ID          string  `json:"id"`          //零售商机构ID
	AgencyOrgID string  `json:"agencyOrgID"` //分销商机构ID AgencyOrg.ID
	RateBasic   float64 `json:"rateBasic"`   //下发机构基础费率
}

/**
 派发记录属性
 */
type DistributionRecordTransientInput struct {
	ID                   string  `json:"id"`                   //派发记录ID
	IndividualID         string  `json:"individualID"`         //个体ID Individual.ID
	ManagedAccountCardNo string  `json:"managedAccountCardNo"` //共管账户账号 FinancialOrgManagedAccountPrivateData.CardNo
	IndividualCardNo     string  `json:"individualCardNo"`     //个体一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	Amount               int     `json:"amount"`               //派发金额
	Rate                 float64 `json:"rate"`                 //派发费率
}

/**
  新增派发记录属性数据
零售商发起下发请求（下发记录ID，下发使用共管账户卡号，个体ID，个体收款一般账户卡号，派发金额，派发费率）
减少共管账户卡号的票据余额，增加个体一般账户的票据余额，增加下发机构的一般账户的佣金票据余额，增加分销商机构的一般账户的佣金票据余额，金融机构的现金余额和票据余额不变。
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
		return "", errors.New("DistributionRecordChaincode value in the transient map must be a non-empty JSON string")
	}
	var transientInput DistributionRecordTransientInput
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
	if transientInput.Amount == 0 {
		return "", errors.New("派发金额不能为0")
	}
	if transientInput.Rate == 0 {
		return "", errors.New("派发费率不能为0")
	}

	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetPrivateData(COLLECTION_DISTRIBUTION_RECORD, transientInput.ID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.ID + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.ID + "\"}"
		return "", errors.New(jsonResp)
	}

	//金融机构共管账户私有数据
	managedAccountPrivateData, err := findManagedAccountPrivateDataById(ctx, transientInput.ManagedAccountCardNo)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if managedAccountPrivateData == nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}

	//个人一般账户
	generalAccountPrivateData, err := FindIndividualPrivateDataById(ctx, transientInput.ManagedAccountCardNo)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if generalAccountPrivateData == nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if !strings.Contains(generalAccountPrivateData.FinancialOrgID, managedAccountPrivateData.FinancialOrgID) {
		return "", errors.New("共管账户不属于当前金融机构" + managedAccountPrivateData.FinancialOrgID)
	}
	//零售商机构
	retailerOrgPrivateData, err := findRetailerPrivateDataById(ctx, managedAccountPrivateData.RetailerOrgID)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if retailerOrgPrivateData == nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if !strings.Contains(retailerOrgPrivateData.ID, managedAccountPrivateData.RetailerOrgID) {
		return "", errors.New("共管账户不属于当前零售商" + managedAccountPrivateData.RetailerOrgID)
	}
	//分销商机构
	agencyOrgPrivateData, err := findAgencyPrivateDataById(ctx, retailerOrgPrivateData.AgencyOrgID)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if agencyOrgPrivateData == nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if !strings.Contains(agencyOrgPrivateData.ID, managedAccountPrivateData.AgencyOrgID) {
		return "", errors.New("共管账户不属于当前分销商" + managedAccountPrivateData.AgencyOrgID)
	}
	//下发机构
	issueOrgPrivateData, err := findIssuePrivateDataById(ctx, managedAccountPrivateData.IssueOrgID)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if issueOrgPrivateData == nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	if !strings.Contains(issueOrgPrivateData.ID, managedAccountPrivateData.IssueOrgID) {
		return "", errors.New("共管账户不属于当前下发机构" + managedAccountPrivateData.IssueOrgID)
	}

	//共管账户向个人一般账户转账票据
	_, err = TransferVoucherAssetIndividual(ctx, managedAccountPrivateData.ManagedCardNo, transientInput.IndividualCardNo, transientInput.Amount)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	//共管账户减少票据
	_, err = TransferManagedVoucherAsset(ctx, managedAccountPrivateData.ManagedCardNo, -transientInput.Amount)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}

	//下发机构佣金
	issueBrokerage, err := CalculationBrokerage(issueOrgPrivateData.RateBasic, agencyOrgPrivateData.RateBasic, transientInput.Amount)
	if err != nil {
		return "", err
	}
	//共管账户向下发机构一般账户转账佣金票据
	_, err = TransferVoucherAssetIssue(ctx, managedAccountPrivateData.ManagedCardNo, managedAccountPrivateData.GeneralCardNo, issueBrokerage)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	//共管账户减少票据
	_, err = TransferManagedVoucherAsset(ctx, managedAccountPrivateData.ManagedCardNo, -issueBrokerage)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}

	//分销商佣金
	retailerBrokerage, err := CalculationBrokerage(agencyOrgPrivateData.RateBasic, retailerOrgPrivateData.RateBasic, transientInput.Amount)
	if err != nil {
		return "", err
	}
	//共管账户向分销商一般账户转账佣金票据
	_, err = TransferVoucherAssetAgency(ctx, managedAccountPrivateData.ManagedCardNo, managedAccountPrivateData.AgencyCardNo, retailerBrokerage)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	//共管账户减少票据
	_, err = TransferManagedVoucherAsset(ctx, managedAccountPrivateData.ManagedCardNo, -retailerBrokerage)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}

	//构造下发记录
	distributionRecordPrivateDat := DistributionRecordPrivateData{
		ID:                   transientInput.ID,
		PlatformOrgID:        managedAccountPrivateData.PlatformOrgID,
		FinancialOrgID:       managedAccountPrivateData.FinancialOrgID,
		IndividualID:         transientInput.IndividualID,
		RetailerOrgID:        managedAccountPrivateData.RetailerOrgID,
		AgencyOrgID:          managedAccountPrivateData.AgencyOrgID,
		IssueOrgID:           managedAccountPrivateData.IssueOrgID,
		ManagedAccountCardNo: transientInput.ManagedAccountCardNo,
		IssueCardNo:          managedAccountPrivateData.IssueCardNo,
		IndividualCardNo:     transientInput.IndividualCardNo,
		AgencyCardNo:         managedAccountPrivateData.AgencyCardNo,
		ManagedCardNo:        managedAccountPrivateData.ManagedCardNo,
		GeneralCardNo:        managedAccountPrivateData.GeneralCardNo,
		Amount:               transientInput.Amount,
		Rate:                 transientInput.Rate,
		Status:               1,
	}
	carAsBytes, _ := json.Marshal(distributionRecordPrivateDat)
	err = ctx.GetStub().PutPrivateData(COLLECTION_DISTRIBUTION_RECORD, distributionRecordPrivateDat.ID, carAsBytes)
	if err != nil {
		return "", errors.New("下发记录保存失败" + err.Error())
	}

	return transientInput.ID, nil
}

func (t *DistributionRecordChaincode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("下发记录id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_DISTRIBUTION_RECORD, id)
	if err != nil {
		return "", errors.New("下发记录查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("下发记录数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

func FindIndividualPrivateDataById(ctx contractapi.TransactionContextInterface, cardNo string) (*FinancialOrgGeneralAccountPrivateData, error) {
	if len(cardNo) == 0 {
		return nil, errors.New("金融机构一般账户账号不能为空")
	}
	trans := [][]byte{[]byte("FindIndividualPrivateDataById"), []byte("id"), []byte(cardNo)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_FINANCIAL_GENERAL_ACCOUNT, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return nil, fmt.Errorf(errStr)
	}

	generalAccountPrivateData := new(FinancialOrgGeneralAccountPrivateData)
	err := json.Unmarshal(response.Payload, &generalAccountPrivateData)
	if err != nil {
		return nil, errors.New("Failed to decode JSON of: " + string(response.Payload))
	}

	return generalAccountPrivateData, nil
}

func findManagedAccountPrivateDataById(ctx contractapi.TransactionContextInterface, cardNo string) (*FinancialOrgManagedAccountPrivateData, error) {
	if len(cardNo) == 0 {
		return nil, errors.New("金融机构共管账户账号不能为空")
	}
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(cardNo)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_FINANCIAL_MANAGED_ACCOUNT, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return nil, fmt.Errorf(errStr)
	}

	managedAccountPrivateData := new(FinancialOrgManagedAccountPrivateData)
	err := json.Unmarshal(response.Payload, &managedAccountPrivateData)
	if err != nil {
		return nil, errors.New("Failed to decode JSON of: " + string(response.Payload))
	}

	return managedAccountPrivateData, nil
}

func findIssuePrivateDataById(ctx contractapi.TransactionContextInterface, id string) (*IssueOrgPrivateData, error) {
	if len(id) == 0 {
		return nil, errors.New("下发机构ID不能为空")
	}
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(id)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return nil, fmt.Errorf(errStr)
	}

	issueOrgPrivateData := new(IssueOrgPrivateData)
	err := json.Unmarshal(response.Payload, &issueOrgPrivateData)
	if err != nil {
		return nil, errors.New("Failed to decode JSON of: " + string(response.Payload))
	}

	return issueOrgPrivateData, nil
}

func findAgencyPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (*AgencyOrgPrivateData, error) {
	if len(id) == 0 {
		return nil, errors.New("分销商机构ID不能为空")
	}
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(id)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_AGENCY_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return nil, fmt.Errorf(errStr)
	}

	agencyOrgPrivateData := new(AgencyOrgPrivateData)
	err := json.Unmarshal(response.Payload, &agencyOrgPrivateData)
	if err != nil {
		return nil, errors.New("Failed to decode JSON of: " + string(response.Payload))
	}

	return agencyOrgPrivateData, nil
}

func findRetailerPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (*RetailerOrgPrivateData, error) {
	if len(id) == 0 {
		return nil, errors.New("分销商机构ID不能为空")
	}
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(id)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_RETAILER_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return nil, fmt.Errorf(errStr)
	}

	retailerOrgPrivateData := new(RetailerOrgPrivateData)
	err := json.Unmarshal(response.Payload, &retailerOrgPrivateData)
	if err != nil {
		return nil, errors.New("Failed to decode JSON of: " + string(response.Payload))
	}

	return retailerOrgPrivateData, nil
}

/**
  票据交易
派发时增加下发机构的票据
 */
func TransferVoucherAssetIssue(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, amount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转出一般账户卡号不能为空")
	}
	if amount < 0 {
		return "", errors.New("转账金额不能小于0")
	}
	trans := [][]byte{[]byte("TransferVoucherAssetIssue"), []byte("managedCardNo"), []byte(managedCardNo), []byte("generalCardNo"), []byte(generalCardNo), []byte("amount"), []byte(string(amount))}
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
  票据交易
派发时增加分销商的票据
 */
func TransferVoucherAssetAgency(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, amount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转出一般账户卡号不能为空")
	}
	if amount < 0 {
		return "", errors.New("转账金额不能小于0")
	}
	trans := [][]byte{[]byte("TransferVoucherAssetAgency"), []byte("managedCardNo"), []byte(managedCardNo), []byte("generalCardNo"), []byte(generalCardNo), []byte("amount"), []byte(string(amount))}
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
  票据交易
派发时增加个体的票据
 */
func TransferVoucherAssetIndividual(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, amount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转出一般账户卡号不能为空")
	}
	if amount < 0 {
		return "", errors.New("转账金额不能小于0")
	}
	trans := [][]byte{[]byte("TransferVoucherAssetIndividual"), []byte("managedCardNo"), []byte(managedCardNo), []byte("generalCardNo"), []byte(generalCardNo), []byte("amount"), []byte(string(amount))}
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
  共管账户票据交易
 */
func TransferManagedVoucherAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, voucherAmount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	trans := [][]byte{[]byte("TransferVoucherAsset"), []byte("managedCardNo"), []byte(managedCardNo), []byte("voucherAmount"), []byte(string(voucherAmount))}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_FINANCIAL_MANAGED_ACCOUNT, trans, CHANNEL_NAME)

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
		fmt.Printf("Error create DistributionRecordChaincode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting DistributionRecordChaincode chaincode: %s", err.Error())
	}
}

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

const (
	//个人一般账户私有数据集名称
	COLLECTION_FINANCIAL_INDIVIDUAL string = "collectionFinancialIndividual"

	//商户机构一般账户私有数据集名称
	COLLECTION_FINANCIAL_MERCHANT string = "collectionFinancialMerchant"

	//平台机构一般账户私有数据集名称
	COLLECTION_FINANCIAL_PLATFORM string = "collectionFinancialPlatform"

	//代理机构一般账户私有数据集名称
	COLLECTION_FINANCIAL_AGENCY string = "collectionFinancialAgency"

	//下发机构一般账户私有数据集名称
	COLLECTION_FINANCIAL_ISSUE string = "collectionFinancialIssue"
)
const (
	//默认
	CERTIFICATE_TYPE_0 int = 0
	//身份证
	CERTIFICATE_TYPE_1 int = 1
	//港澳台证
	CERTIFICATE_TYPE_2 int = 2
	//护照
	CERTIFICATE_TYPE_3 int = 3
	//军官证
	CERTIFICATE_TYPE_4 int = 4
	//统一社会信用代码
	CERTIFICATE_TYPE_5 int = 5
)

/**
   金融机构一般账户私有数据属性
 */
type FinancialOrgGeneralAccountPrivateData struct {
	CardNo                string `json:"cardNo"`                //金融机构公管账户账号(唯一不重复)
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	CertificateNo         string `json:"certificateNo"`         //持卡者证件号
	CertificateType       int    `json:"certificateType"`       //持卡者证件类型 (身份证/港澳台证/护照/军官证/统一社会信用代码)
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

func (t *FinancialGeneralAccountChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("IssueChaincode Init")
	//个体
	individualOrg1 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7636941931858894848", FinancialOrgID: "F766005404604841984", CertificateNo: "888888888888888888", CertificateType: CERTIFICATE_TYPE_1, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	individualOrg2 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7636941931858894848", FinancialOrgID: "F766374712807800832", CertificateNo: "888888888888888888", CertificateType: CERTIFICATE_TYPE_1, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	//商户机构
	merchantOrg1 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7637019072143298560", FinancialOrgID: "F766005404604841984", CertificateNo: "92370112MA3F23MB5N", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	merchantOrg2 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7637019072143298560", FinancialOrgID: "F766374712807800832", CertificateNo: "92370112MA3F23MB5N", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	//代理商机构
	agencyOrg1 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7637019072143298560", FinancialOrgID: "F766005404604841984", CertificateNo: "92370104MA3DR08A4D", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	agencyOrg2 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7637019072143298560", FinancialOrgID: "F766374712807800832", CertificateNo: "92370104MA3DR08A4D", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	//下发机构
	issueOrg1 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7637019072143298560", FinancialOrgID: "F766005404604841984", CertificateNo: "91370181MA3D7J9W3W", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	issueOrg2 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7637019072143298560", FinancialOrgID: "F766374712807800832", CertificateNo: "91370181MA3D7J9W3W", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	//平台机构
	platformOrg1 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7637019072143298560", FinancialOrgID: "F766005404604841984", CertificateNo: "91370100MA3MP74K7A", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}
	platformOrg2 := FinancialOrgGeneralAccountPrivateData{CardNo: "I7637019072143298560", FinancialOrgID: "F766374712807800832", CertificateNo: "91370100MA3MP74K7A", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1}

	//个体私有数据
	individualOrg1JSON, err := json.Marshal(individualOrg1)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_INDIVIDUAL, individualOrg1.CardNo, individualOrg1JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	individualOrg2JSON, err := json.Marshal(individualOrg2)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_INDIVIDUAL, individualOrg2.CardNo, individualOrg2JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	//商户机构私有数据
	merchantOrg1JSON, err := json.Marshal(merchantOrg1)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_MERCHANT, merchantOrg1.CardNo, merchantOrg1JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	merchantOrg2JSON, err := json.Marshal(merchantOrg2)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_MERCHANT, merchantOrg2.CardNo, merchantOrg2JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	//代理商机构私有数据
	agencyOrg1JSON, err := json.Marshal(agencyOrg1)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_AGENCY, agencyOrg1.CardNo, agencyOrg1JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	agencyOrg2JSON, err := json.Marshal(agencyOrg2)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_AGENCY, agencyOrg2.CardNo, agencyOrg2JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	//下发机构私有数据
	issueOrg1JSON, err := json.Marshal(issueOrg1)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_ISSUE, issueOrg1.CardNo, issueOrg1JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	issueOrg2JSON, err := json.Marshal(issueOrg2)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_ISSUE, issueOrg2.CardNo, issueOrg2JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	//平台机构私有数据
	platformOrg1JSON, err := json.Marshal(platformOrg1)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_PLATFORM, platformOrg1.CardNo, platformOrg1JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	platformOrg2JSON, err := json.Marshal(platformOrg2)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_PLATFORM, platformOrg2.CardNo, platformOrg2JSON)
	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	return nil
}

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
	if transientInput.CertificateType == 0 {
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

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToIndividual(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, COLLECTION_FINANCIAL_INDIVIDUAL)
}

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToMerchant(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, COLLECTION_FINANCIAL_MERCHANT)
}

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToAgency(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, COLLECTION_FINANCIAL_AGENCY)
}

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToIssue(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, COLLECTION_FINANCIAL_ISSUE)
}

func (t *FinancialGeneralAccountChaincode) CreateGeneralAccountToPlatform(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.Create(ctx, COLLECTION_FINANCIAL_PLATFORM)
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

/**
  现金交易
 */
func (t *FinancialGeneralAccountChaincode) TransferAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, amount int) error {
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
	newCurrentBalance := transientInput.CurrentBalance + amount
	if newCurrentBalance < 0 {
		return errors.New("一般账户现金余额不足")
	}
	transientInput.CurrentBalance = newCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutState(generalCardNo, assetJSON)
}

/**
  票据交易
 */
func (t *FinancialGeneralAccountChaincode) TransferVoucherAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, voucherAmount int) error {
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

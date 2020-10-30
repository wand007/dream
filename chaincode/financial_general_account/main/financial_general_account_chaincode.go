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

//私有数据集名称
const COLLECTION_FINANCIAL_GENERAL_ACCOUNT string = "collectionFinancialOrgGeneralAccount"

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
	fmt.Println("FinancialGeneralAccountChaincode Init")

	generalAccountPrivateDatas := []FinancialOrgGeneralAccountPrivateData{
		//个体
		{CardNo: "6229486603953152819", FinancialOrgID: "F766005404604841984", CertificateNo: "888888888888888888", CertificateType: CERTIFICATE_TYPE_1, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953152825", FinancialOrgID: "F766374712807800832", CertificateNo: "888888888888888888", CertificateType: CERTIFICATE_TYPE_1, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		//商户机构
		{CardNo: "6229486603953174011", FinancialOrgID: "F766005404604841984", CertificateNo: "92370104MA3DR08A4D", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953174027", FinancialOrgID: "F766374712807800832", CertificateNo: "92370104MA3DR08A4D", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		//代理商机构
		{CardNo: "6229486603953188912", FinancialOrgID: "F766005404604841984", CertificateNo: "92370112MA3F23MB5N", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953188928", FinancialOrgID: "F766374712807800832", CertificateNo: "92370112MA3F23MB5N", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		//下发机构
		{CardNo: "6229486603953201814", FinancialOrgID: "F766005404604841984", CertificateNo: "91370181MA3D7J9W3W", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953201820", FinancialOrgID: "F766374712807800832", CertificateNo: "91370181MA3D7J9W3W", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		//平台机构
		{CardNo: "6229486603953250118", FinancialOrgID: "F766005404604841984", CertificateNo: "91370100MA3MP74K7A", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953250124", FinancialOrgID: "F766374712807800832", CertificateNo: "91370100MA3MP74K7A", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
	}
	//私有数据
	for _, asset := range generalAccountPrivateDatas {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, asset.CardNo, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}
	return nil
}

/**
 新增金融机构一般账户私有数据
 */
func (t *FinancialGeneralAccountChaincode) Create(ctx contractapi.TransactionContextInterface) (string, error) {

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}

	financialPrivateDataJsonBytes, ok := transMap["generalAccount"]
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
	Avalbytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, transientInput.CardNo)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}
	carAsBytes, _ := json.Marshal(transientInput)

	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, transientInput.CardNo, carAsBytes)
	if err != nil {
		return "", errors.New("商户共管账户保存失败" + err.Error())
	}
	return "", nil
}

/**
  现金交易--充值用
商户向商户一般账户充值现金余额
 */
func (t *FinancialGeneralAccountChaincode) TransferCashAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, amount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	if amount < 0 {
		return errors.New("一般账户充值金额不能小于0")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo)
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
  票据交易--下发用
 */
func (t *FinancialGeneralAccountChaincode) TransferVoucherAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, voucherAmount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo)
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

/**
  现金和票据交易 （票据提现和票据充值）
 */
func (t *FinancialGeneralAccountChaincode) TransferAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, voucherAmount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo)
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
	newCurrentBalance := transientInput.CurrentBalance - voucherAmount
	if newCurrentBalance < 0 {
		return errors.New("一般账户余额不足")
	}
	transientInput.VoucherCurrentBalance = newVoucherCurrentBalance
	transientInput.CurrentBalance = newCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutState(generalCardNo, assetJSON)
}

func (t *FinancialGeneralAccountChaincode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("共管账户id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, id)
	if err != nil {
		return "", errors.New("共管账户查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("共管账户数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
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

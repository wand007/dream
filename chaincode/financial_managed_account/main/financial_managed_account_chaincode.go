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

//私有数据集名称
const COLLECTION_FINANCIAL_MANAGED_ACCOUNT string = "collectionFinancialManagedAccount"

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

func (t *FinancialManagedAccountChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("FinancialManagedAccountChaincode Init")
	//私有数据
	managedAccountPrivateData := []FinancialOrgManagedAccountPrivateData{
		{CardNo: "3036603953562710", MerchantOrgID: "M766005404604841984", FinancialOrgID: "F766005404604841984", PlatformOrgID: "P768877118787432448", VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "3038603953562825", MerchantOrgID: "M764441096829812736", FinancialOrgID: "F766005404604841984", PlatformOrgID: "P768877118787432448", VoucherCurrentBalance: 0, AccStatus: 1},

		{CardNo: "3036603953578518", MerchantOrgID: "M766005404604841984", FinancialOrgID: "F766374712807800832", PlatformOrgID: "P768877118787432448", VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "3038603953578524", MerchantOrgID: "M764441096829812736", FinancialOrgID: "F766374712807800832", PlatformOrgID: "P768877118787432448", VoucherCurrentBalance: 0, AccStatus: 1},
	}
	for _, asset := range managedAccountPrivateData {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_MANAGED_ACCOUNT, asset.CardNo, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

/**
  新增金融机构共管账户私有数据
 */
func (t *FinancialManagedAccountChaincode) Create(ctx contractapi.TransactionContextInterface) (string, error) {
	//私有数据
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}

	financialPrivateDataJsonBytes, ok := transMap["financialManagedAccount"]
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
		return "", errors.New("平台机构ID不能为空")
	}
	if len(transientInput.MerchantOrgID) == 0 {
		return "", errors.New("商户机构ID不能为空")
	}
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_MANAGED_ACCOUNT, transientInput.CardNo)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	carAsBytes, _ := json.Marshal(transientInput)

	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_MANAGED_ACCOUNT, transientInput.CardNo, carAsBytes)
	if err != nil {
		return "", errors.New("商户共管账户保存失败" + err.Error())
	}
	return string(Avalbytes), nil
}

/**
	现金交易
 */
func (t *FinancialManagedAccountChaincode) TransferAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, voucherAmount int) error {
	if len(managedCardNo) == 0 {
		return errors.New("共管账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_MANAGED_ACCOUNT, managedCardNo)
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

/**
	票据交易
 */
func (t *FinancialManagedAccountChaincode) TransferVoucherAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, voucherAmount int) error {
	if len(managedCardNo) == 0 {
		return errors.New("共管账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_MANAGED_ACCOUNT, managedCardNo)
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

func (t *FinancialManagedAccountChaincode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("共管账户id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_MANAGED_ACCOUNT, id)
	if err != nil {
		return "", errors.New("共管账户查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("共管账户数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
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

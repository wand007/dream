package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/**
初始化共管账户记录
新建共管账户记录
一般账户票据交易
 */
/**
下发机构和零售商共同持有
 */
/**
共管账户链码
 */
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
	PlatformOrgID         string `json:"platformOrgID"`         //平台机构ID PlatformOrg.ID
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	IssueOrgID            string `json:"issueOrgID"`            //下发机构ID IssueOrg.ID
	RetailerOrgID         string `json:"retailerOrgID"`         //零售商机构ID RetailerOrg.ID
	AgencyOrgID           string `json:"agencyOrgID"`           //分销商机构ID AgencyOrg.ID
	IssueCardNo           string `json:"issueCardNo"`           //下发机构一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	AgencyCardNo          string `json:"agencyCardNo"`          //分销商机构一般账户账号 FinancialOrgManagedAccountPrivateData.CardNo
	ManagedCardNo         string `json:"managedCardNo"`         //分销商机构一般账户账号 FinancialOrgManagedAccountPrivateData.CardNo
	GeneralCardNo         string `json:"generalCardNo"`         //零售商机构一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构零售商机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

func (t *FinancialManagedAccountChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("FinancialManagedAccountChaincode Init")
	//私有数据
	managedAccountPrivateData := []FinancialOrgManagedAccountPrivateData{
		{CardNo: "3036603953562710", AgencyOrgID: "A766005404604841984", RetailerOrgID: "M766005404604841984", FinancialOrgID: "F766005404604841984", IssueOrgID: "I766005404604841984", PlatformOrgID: "P768877118787432448", IssueCardNo: "6229486603953201814", AgencyCardNo: "6229486603953188912", ManagedCardNo: "6229486603953188912", GeneralCardNo: "6229486603953174011", VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "3038603953562825", AgencyOrgID: "A766374712807800832", RetailerOrgID: "M764441096829812736", FinancialOrgID: "F766005404604841984", IssueOrgID: "I764441096829812736", PlatformOrgID: "P768877118787432448", IssueCardNo: "6229488603953201820", AgencyCardNo: "6229488603953188928", ManagedCardNo: "6229488603953188928", GeneralCardNo: "6229488603953174027", VoucherCurrentBalance: 0, AccStatus: 1},

		{CardNo: "3036603953578518", AgencyOrgID: "A766005404604841984", RetailerOrgID: "M766005404604841984", FinancialOrgID: "F766374712807800832", IssueOrgID: "I766005404604841984", PlatformOrgID: "P768877118787432448", IssueCardNo: "6229486603953201814", AgencyCardNo: "6229486603953188912", ManagedCardNo: "6229486603953188912", GeneralCardNo: "6229486603953174011", VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "3038603953578524", AgencyOrgID: "A766374712807800832", RetailerOrgID: "M764441096829812736", FinancialOrgID: "F766374712807800832", IssueOrgID: "I764441096829812736", PlatformOrgID: "P768877118787432448", IssueCardNo: "6229488603953201820", AgencyCardNo: "6229488603953188928", ManagedCardNo: "6229488603953188928", GeneralCardNo: "6229488603953174027", VoucherCurrentBalance: 0, AccStatus: 1},
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

	financialPrivateDataJsonBytes, ok := transMap["managedAccount"]
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
	if len(transientInput.IssueOrgID) == 0 {
		return "", errors.New("下发机构ID不能为空")
	}
	if len(transientInput.RetailerOrgID) == 0 {
		return "", errors.New("零售商机构ID不能为空")
	}
	if len(transientInput.AgencyOrgID) == 0 {
		return "", errors.New("分销商机构ID不能为空")
	}
	if len(transientInput.IssueCardNo) == 0 {
		return "", errors.New("下发机构一般账户账号不能为空")
	}
	if len(transientInput.ManagedCardNo) == 0 {
		return "", errors.New("分销商机构一般账户账号不能为空")
	}
	if len(transientInput.GeneralCardNo) == 0 {
		return "", errors.New("零售商机构一般账户账号不能为空")
	}
	if len(transientInput.AgencyCardNo) == 0 {
		return "", errors.New("分销商机构一般账户账号不能为空")
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
		return "", errors.New("零售商共管账户保存失败" + err.Error())
	}
	return string(Avalbytes), nil
}

/**
	票据交易
零售商向零售商共管账户充值现金时增加票据余额
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
	newVoucherCurrentBalance := transientInput.VoucherCurrentBalance + voucherAmount
	if newVoucherCurrentBalance < 0 {
		return errors.New("共管账户票据余额不足")
	}
	transientInput.VoucherCurrentBalance = newVoucherCurrentBalance
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
		fmt.Printf("Error create FinancialManagedAccountChaincode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting FinancialManagedAccountChaincode chaincode: %s", err.Error())
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type FinancialChainCode struct {
	contractapi.Contract
}

/**
   金融机构属性
 */
type FinancialOrg struct {
	ID     string `json:"id"`     //金融机构ID
	Name   string `json:"name"`   //金融机构名称
	Code   string `json:"code"`   //金融机构代码
	Status int    `json:"status"` //金融机构状态(启用/禁用)
}

/**
   金融机构一般账户私有数据属性
 */
type FinancialOrgGeneralAccountPrivateData struct {
	ID              string `json:"id"`              //金融机构ID
	CardNo          string `json:"cardNo"`          //金融机构共管账户账号
	FinancialOrgID  string `json:"financialOrgID"`  //金融机构ID FinancialOrg.ID
	CertificateNo   string `json:"certificateNo"`   //持卡者证件号
	CertificateType string `json:"certificateType"` //持卡者证件类型 (身份证/港澳台证/护照/军官证)¬
	CurrentBalance  int    `json:"currentBalance"`  //金融机构共管账户余额(现金)
	AccStatus       int    `json:"accStatus"`       //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

/**
   金融机构共管账户私有数据属性
 */
type FinancialOrgManagedAccountPrivateData struct {
	ID                    string `json:"id"`                    //金融机构ID
	CardNo                string `json:"cardNo"`                //金融机构共管账户账号
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	PlatformOrgID         string `json:"platformOrgID"`         //平台机构ID PlatformOrg.ID
	MerchantOrgID         string `json:"merchantOrgID"`         //商户机构ID MerchantOrg.ID
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

/**
   充值操作参数
 */
type RechargeTransientInput struct {
	ID            string `json:"id"`            //金融机构ID
	ManagedCardNo string `json:"managedCardNo"` //公管账户卡号 FinancialOrgManagedAccountPrivateData.CardNo
	GeneralCardNo string `json:"generalCardNo"` //一般账户卡号 FinancialOrgGeneralAccountPrivateData.CardNo
	Amount        int    `json:"amount"`        //充值金额
}

func (t *FinancialChainCode) Create(ctx contractapi.TransactionContextInterface, id string, name string,
	code string, status int) (string, error) {

	if len(id) == 0 {
		return "", errors.New("金融机构ID不能为空")
	}
	if len(name) == 0 {
		return "", errors.New("金融机构名称不能为空")
	}
	if len(code) == 0 {
		return "", errors.New("金融机构代码不能为空")
	}
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	queryString := fmt.Sprintf(`{"selector":{"name":"%s"}}`, name)    //Mongo Query string 语法见上文链接
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString) // 富查询的返回结果可能为多条 所以这里返回的是一个迭代器 需要我们进一步的处理来获取需要的结果

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	if resultsIterator != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	queryString = fmt.Sprintf(`{"selector":{"code":"%s"}}`, code)    //Mongo Query string 语法见上文链接
	resultsIterator, err = ctx.GetStub().GetQueryResult(queryString) // 富查询的返回结果可能为多条 所以这里返回的是一个迭代器 需要我们进一步的处理来获取需要的结果

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + code + "\"}"
		return "", errors.New(jsonResp)
	}

	if resultsIterator != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + code + "\"}"
		return "", errors.New(jsonResp)
	}

	financial := &FinancialOrg{
		ID:     id,
		Name:   name,
		Code:   code,
		Status: status,
	}
	carAsBytes, _ := json.Marshal(financial)

	err = ctx.GetStub().PutState(id, carAsBytes)
	if err != nil {
		return "", errors.New("金融机构保存失败" + err.Error())
	}
	return string(Avalbytes), nil
}

/**
 新增金融机构一般账户私有数据
 */
func (t *FinancialChainCode) CreateGeneralAccount(ctx contractapi.TransactionContextInterface, collectionName string) (string, error) {

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
	var transientInput FinancialOrgGeneralAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	id := transientInput.ID
	if len(id) == 0 {
		return "", errors.New("金融机构共管账户ID不能为空")
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
	Avalbytes, err := ctx.GetStub().GetPrivateData(collectionName, id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
		return "", errors.New(jsonResp)
	}
	// Mongo Query string 语法见上文链接
	queryString := fmt.Sprintf(`{"selector":{"cardNo":"%s"}}`, transientInput.CardNo)
	// 富查询的返回结果可能为多条 所以这里返回的是一个迭代器 需要我们进一步的处理来获取需要的结果
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	if resultsIterator != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	carAsBytes, _ := json.Marshal(transientInput)

	err = ctx.GetStub().PutPrivateData(collectionName, id, carAsBytes)
	if err != nil {
		return "", errors.New("商户共管账户保存失败" + err.Error())
	}
	return string(Avalbytes), nil
}

func (t *FinancialChainCode) Recharge(ctx contractapi.TransactionContextInterface) (string, error) {
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}

	individualPrivateDataJsonBytes, ok := transMap["individual"]
	if !ok {
		return "", errors.New("individual must be a key in the transient map")
	}

	if len(individualPrivateDataJsonBytes) == 0 {
		return "", errors.New("individual value in the transient map must be a non-empty JSON string")
	}
	//公开数据
	var rechargeTransientInput RechargeTransientInput
	err = json.Unmarshal(individualPrivateDataJsonBytes, &rechargeTransientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(individualPrivateDataJsonBytes))
	}
	id := rechargeTransientInput.ID
	if len(id) == 0 {
		return "", errors.New("转入操作ID不能为空")
	}
	managedCardNo := rechargeTransientInput.ManagedCardNo
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	generalCardNo := rechargeTransientInput.GeneralCardNo
	if len(generalCardNo) == 0 {
		return "", errors.New("转出一般账户卡号不能为空")
	}
	Amount := rechargeTransientInput.Amount

	generalAccountPrivateData, err := findGeneralAccountPrivateDataById(ctx, generalCardNo)
	if err != nil {
		return "", err
	}
	if generalAccountPrivateData == nil {
		return "", errors.New("一般账户卡号记录不存在")
	}
	//账户余额不能超过转账操作金额
	if generalAccountPrivateData.CurrentBalance < Amount {
		return "", errors.New("转出账户余额不足")
	}
	managedAccountPrivateData, err := findManagedAccountPrivateDataById(ctx, managedCardNo)
	if err != nil {
		return "", err
	}
	if managedAccountPrivateData == nil {
		return "", errors.New("共管账户卡号记录不存在")
	}

	//todo 增加共管账户
	//todo 减少一般账户

	return id, nil
}

func findManagedAccountPrivateDataById(ctx contractapi.TransactionContextInterface, managedCardNo string) (*FinancialOrgManagedAccountPrivateData, error) {
	if len(managedCardNo) == 0 {
		return nil, errors.New("一般账户卡号不能为空")
	}
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(managedCardNo)}
	response := ctx.GetStub().InvokeChaincode("financialManagedAccountCC", trans, "mychannel")

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

func findGeneralAccountPrivateDataById(ctx contractapi.TransactionContextInterface, generalCardNo string) (*FinancialOrgManagedAccountPrivateData, error) {
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(generalCardNo)}
	response := ctx.GetStub().InvokeChaincode("financialGeneralAccountCC", trans, "mychannel")

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

func (t *FinancialChainCode) CreateManagedAccountToMerchant(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.CreateGeneralAccount(ctx, "collectionFinancialMerchant")
}

func (t *FinancialChainCode) CreateManagedAccountToPlatform(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.CreateGeneralAccount(ctx, "collectionFinancialPlatform")
}

func (t *FinancialChainCode) CreateManagedAccountToAgency(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.CreateGeneralAccount(ctx, "collectionFinancialAgency")
}

func (t *FinancialChainCode) CreateManagedAccountToIssue(ctx contractapi.TransactionContextInterface) (string, error) {
	return t.CreateGeneralAccount(ctx, "collectionFinancialIssue")
}

func (t *FinancialChainCode) FindById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("金融机构id不能为空")
	}
	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", errors.New("金融机构查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("金融机构数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

func (t *FinancialChainCode) FindMerchantPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	return t.FindPrivateDataById(ctx, id, "collectionFinancialMerchant")
}

func (t *FinancialChainCode) FindPlatformPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	return t.FindPrivateDataById(ctx, id, "collectionFinancialPlatform")
}

func (t *FinancialChainCode) FindAgencyPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	return t.FindPrivateDataById(ctx, id, "collectionFinancialAgency")
}

func (t *FinancialChainCode) FindIssuePrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	return t.FindPrivateDataById(ctx, id, "collectionFinancialIssue")
}

func (t *FinancialChainCode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string, collectionName string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("金融机构id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData("collectionPlatform", id)
	if err != nil {
		return "", errors.New("金融机构私有数据查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("金融机构私有数据不存在，读到的%s对应的私有数据为空！", id)
	}
	return string(bytes), nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(FinancialChainCode))
	if err != nil {
		fmt.Printf("Error create PlatformChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting PlatformChainCode chaincode: %s", err.Error())
	}
}

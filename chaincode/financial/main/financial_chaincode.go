package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/**
初始化默认金融机构
新建金融机构
一般账户向共管账户现金兑换票据(充值)
一般账户向共管账户票据兑换现金(提现)
共管账户向一般账户发放票据(派发)
 */
/**
金融机构链码
 */
type FinancialChainCode struct {
	contractapi.Contract
}

//通道名称
const CHANNEL_NAME string = "mychannel"

//链码名称
const CHAINCODE_NAME_ISSUE_ORG string = "issueCC"

//私有数据集名称
const COLLECTION_FINANCIAL string = "collectionFinancial"

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
   金融机构私有数据属性
 */
type FinancialOrgPrivateData struct {
	ID                    string `json:"id"`                    //金融机构ID FinancialOrg.ID
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构零售商机构账户凭证(token)余额
}

/**
   金融机构共管账户私有数据属性
 */
type FinancialOrgManagedAccountPrivateData struct {
	ID                    string `json:"id"`                    //金融机构ID
	CardNo                string `json:"cardNo"`                //金融机构共管账户账号
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	PlatformOrgID         string `json:"platformOrgID"`         //平台机构ID PlatformOrg.ID
	RetailerOrgID         string `json:"retailerOrgID"`         //零售商机构ID RetailerOrg.ID
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构零售商机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

func (t *FinancialChainCode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("FinancialChainCode Init")
	//私有数据
	financialOrgs := []FinancialOrg{
		{ID: "F766005404604841984", Name: "默认金融机构1", Code: "1", Status: 1},
		{ID: "F766374712807800832", Name: "默认金融机构2", Code: "2", Status: 1},
	}
	for _, asset := range financialOrgs {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
		//私有数据
		financialOrgPrivateData := FinancialOrgPrivateData{ID: asset.ID, CurrentBalance: 0, VoucherCurrentBalance: 0}
		assetJSON, err = json.Marshal(financialOrgPrivateData)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL, asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

/**
   新增金融机构共管账户私有数据
 */
func (t *FinancialChainCode) Create(ctx contractapi.TransactionContextInterface, id string, name string, code string, status int) (string, error) {
	fmt.Println("id:" + id + ",name:" + name + ",code:" + code)
	//公有数据入参参数
	if len(id) == 0 {
		return "", errors.New("金融机构ID不能为空")
	}
	if len(name) == 0 {
		return "", errors.New("金融机构名称不能为空")
	}
	if len(code) == 0 {
		return "", errors.New("金融机构代码不能为空")
	}
	//防重复提交
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get id for " + id + "\"}"
		return "", errors.New(jsonResp)
	}
	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"金融机构ID:" + id + "不能重复\"}"
		return "", errors.New(jsonResp)
	}
	//银行名称防重复提交
	financialByName, err := t.queryFinancialByName(ctx, name)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get name for " + name + "\"}"
		return "", errors.New(jsonResp)
	}
	fmt.Printf("- financialByName queryResult:%#v", financialByName)
	if len(financialByName) != 0 {
		jsonResp := "{\"Error\":\"金融机构名称:" + name + "不能重复\"}"
		return "", errors.New(jsonResp)
	}
	//银行代码防重复提交
	financialByCode, err := t.queryFinancialByCode(ctx, code)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get code for " + code + "\"}"
		return "", errors.New(jsonResp)
	}
	fmt.Printf("- financialByCode queryResult:%#v", financialByCode)

	if len(financialByCode) != 0 {
		jsonResp := "{\"Error\":\"金融机构代码:" + code + "不能重复\"}"
		return "", errors.New(jsonResp)
	}
	//公开数据
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
	Avalbytes, err = ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL, id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"私有数据不为空 " + id + "\"}"
		return "", errors.New(jsonResp)
	}
	//私有数据
	financialPrivateData := &FinancialOrgPrivateData{
		ID:                    id,
		CurrentBalance:        0,
		VoucherCurrentBalance: 0,
	}
	financialPrivateDataAsBytes, _ := json.Marshal(financialPrivateData)

	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL, id, financialPrivateDataAsBytes)
	if err != nil {
		return "", errors.New("金融机构保存失败" + err.Error())
	}
	//发送事件通知
	ctx.GetStub().SetEvent("FinancialOrg",financialPrivateDataAsBytes)
	return string(Avalbytes), nil
}

/**
	一般账户向共管账户现金兑换票据
 */
func (t *FinancialChainCode) Grant(ctx contractapi.TransactionContextInterface, id string, amount int) (int, error) {
	if len(id) == 0 {
		return 0, errors.New("操作ID不能为空")
	}
	Avalbytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL, id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return 0, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
		return 0, errors.New(jsonResp)
	}
	//私有数据
	var financialOrgPrivateData FinancialOrgPrivateData
	err = json.Unmarshal(Avalbytes, &financialOrgPrivateData)
	if err != nil {
		return 0, errors.New("Failed to decode JSON of: " + string(Avalbytes))
	}
	financialOrgPrivateData.CurrentBalance = financialOrgPrivateData.CurrentBalance - amount
	financialOrgPrivateData.VoucherCurrentBalance = financialOrgPrivateData.VoucherCurrentBalance + amount
	carAsBytes, _ := json.Marshal(financialOrgPrivateData)

	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL, id, carAsBytes)

	if err != nil {
		return 0, errors.New("零售商共管账户保存失败" + err.Error())
	}
	return financialOrgPrivateData.VoucherCurrentBalance, nil
}

/**
  一般账户向共管账户票据兑换现金 (票据变现)
个体/分销商/下发机构发起提现请求
减少金融机构的现金余额和票据余额，增加个体/分销商/下发机构一般账户的现金余额，减少个体/分销商/下发机构一般账户的票据余额
 */
func (t *FinancialChainCode) Realization(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, voucherAmount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转出共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转入一般账户卡号不能为空")
	}
	if voucherAmount < 0 {
		return "", errors.New("兑换票据不能小于0")
	}

	//公管账户
	managedAccountPrivateData, err := findManagedAccountPrivateDataById(ctx, managedCardNo)
	if err != nil {
		return "", err
	}
	if managedAccountPrivateData == nil {
		return "", errors.New("共管账户卡号记录不存在")
	}
	//账户余额不能超过转账操作金额
	if managedAccountPrivateData.CurrentBalance < voucherAmount {
		return "", errors.New("共管账户余额不足")
	}
	//账户余额不能超过转账操作金额
	if managedAccountPrivateData.VoucherCurrentBalance < voucherAmount {
		return "", errors.New("共管账户票据余额不足")
	}
	//一般账户
	generalAccountPrivateData, err := findGeneralAccountPrivateDataById(ctx, generalCardNo)
	if err != nil {
		return "", err
	}
	if generalAccountPrivateData == nil {
		return "", errors.New("一般账户卡号记录不存在")
	}
	//账户余额不能超过转账操作金额
	if generalAccountPrivateData.VoucherCurrentBalance < voucherAmount {
		return "", errors.New("一般账户票据不足")
	}
	if managedAccountPrivateData.FinancialOrgID != generalAccountPrivateData.FinancialOrgID {
		return "", errors.New("请选择相同的金融机构")
	}
	//增加一般户现金余额并减少票据余额
	err = TransferGeneralAsset(ctx, generalCardNo, -voucherAmount)
	if err != nil {
		return "", err
	}
	//减少金融机构账户现金和票据余额
	financialOrgPrivateString, err := t.FindPrivateDataById(ctx, managedAccountPrivateData.FinancialOrgID)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + financialOrgPrivateString)
	}
	var financialOrgPrivateData FinancialOrgPrivateData
	err = json.Unmarshal([]byte(financialOrgPrivateString), &financialOrgPrivateData)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + financialOrgPrivateString)
	}
	financialOrgPrivateData.CurrentBalance = financialOrgPrivateData.CurrentBalance - voucherAmount
	financialOrgPrivateData.VoucherCurrentBalance = financialOrgPrivateData.VoucherCurrentBalance - voucherAmount
	carAsBytes, _ := json.Marshal(financialOrgPrivateData)

	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL, financialOrgPrivateData.ID, carAsBytes)

	if err != nil {
		return "", errors.New("零售商共管账户保存失败" + err.Error())
	}
	return "", nil
}

/**
  一般账户向共管账户现金兑换票据 (现金充值)
零售商用一般账户的现金余额向上级代理的上级下发机构的金融机构的共管账户充值，获取金融机构颁发的票据，共管账户增加票据余额，零售商减少一般账户的现金余额，增加金融机构的现金余额和票据余额。
 */
func (t *FinancialChainCode) TransferAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, amount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转入共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转出一般账户卡号不能为空")
	}
	if amount < 0 {
		return "", errors.New("转账金额不能小于0")
	}

	//公管账户
	managedAccountPrivateData, err := findManagedAccountPrivateDataById(ctx, managedCardNo)
	if err != nil {
		return "", err
	}
	if managedAccountPrivateData == nil {
		return "", errors.New("共管账户卡号记录不存在")
	}

	//一般账户
	generalAccountPrivateData, err := findGeneralAccountPrivateDataById(ctx, generalCardNo)
	if err != nil {
		return "", err
	}
	if generalAccountPrivateData == nil {
		return "", errors.New("一般账户卡号记录不存在")
	}
	//账户余额不能超过转账操作金额
	if generalAccountPrivateData.CurrentBalance < amount {
		return "", errors.New("一般账户余额不足")
	}
	if managedAccountPrivateData.FinancialOrgID != generalAccountPrivateData.FinancialOrgID {
		return "", errors.New("请选择相同的金融机构")
	}
	//减少一般户现金余额
	err = TransferGeneralAsset(ctx, generalCardNo, -amount)
	if err != nil {
		return "", err
	}
	//一般账户向经融机构现金兑换票据
	voucher, err := t.Grant(ctx, managedAccountPrivateData.FinancialOrgID, amount)
	if err != nil {
		return "", err
	}
	//增加共管账户票据余额
	err = TransferManagedVoucherAsset(ctx, managedCardNo, voucher)
	if err != nil {
		return "", err
	}
	return "", nil
}

/**
  共管账户向一般账户交易票据 (票据下发)
 */
func (t *FinancialChainCode) TransferVoucherAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, generalCardNo string, voucherAmount int) (string, error) {
	if len(managedCardNo) == 0 {
		return "", errors.New("转出共管账户卡号不能为空")
	}
	if len(generalCardNo) == 0 {
		return "", errors.New("转入一般账户卡号不能为空")
	}
	if voucherAmount < 0 {
		return "", errors.New("转账金额不能小于0")
	}

	//公管账户
	managedAccountPrivateData, err := findManagedAccountPrivateDataById(ctx, managedCardNo)
	if err != nil {
		return "", err
	}
	if managedAccountPrivateData == nil {
		return "", errors.New("共管账户卡号记录不存在")
	}
	//账户余额不能超过转账操作金额
	if managedAccountPrivateData.CurrentBalance < voucherAmount {
		return "", errors.New("共管账户票据余额不足")
	}

	//一般账户
	generalAccountPrivateData, err := findGeneralAccountPrivateDataById(ctx, generalCardNo)
	if err != nil {
		return "", err
	}
	if generalAccountPrivateData == nil {
		return "", errors.New("一般账户卡号记录不存在")
	}
	if managedAccountPrivateData.FinancialOrgID != generalAccountPrivateData.FinancialOrgID {
		return "", errors.New("请选择相同的金融机构")
	}

	//减少共管账户票据余额
	err = TransferManagedVoucherAsset(ctx, managedCardNo, -voucherAmount)
	if err != nil {
		return "", err
	}
	//增加一般户票据余额
	err = TransferGeneralVoucherAsset(ctx, generalCardNo, voucherAmount)
	if err != nil {
		return "", err
	}

	return "", nil
}

func TransferGeneralAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, voucherAmount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}

	trans := [][]byte{[]byte("TransferAssetRetailer"), []byte("generalCardNo"), []byte(generalCardNo), []byte("voucherAmount"), []byte(string(voucherAmount))}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return errors.New(errStr)
	}
	bytes := response.Payload
	if bytes != nil {
		return errors.New(string(bytes))
	}
	return nil
}

func TransferGeneralVoucherAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, amount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	trans := [][]byte{[]byte("TransferVoucherAsset"), []byte("generalCardNo"), []byte(generalCardNo), []byte("amount"), []byte(string(amount))}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return errors.New(errStr)
	}
	bytes := response.Payload
	if bytes != nil {
		return errors.New(string(bytes))
	}
	return nil
}

func TransferManagedVoucherAsset(ctx contractapi.TransactionContextInterface, managedCardNo string, amount int) error {
	if len(managedCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	trans := [][]byte{[]byte("TransferVoucherAsset"), []byte("managedCardNo"), []byte(managedCardNo), []byte("amount"), []byte(string(amount))}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to FindPrivateDataById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return errors.New(errStr)
	}
	bytes := response.Payload
	if bytes != nil {
		return errors.New(string(bytes))
	}
	return nil
}

func findManagedAccountPrivateDataById(ctx contractapi.TransactionContextInterface, managedCardNo string) (*FinancialOrgManagedAccountPrivateData, error) {
	if len(managedCardNo) == 0 {
		return nil, errors.New("一般账户卡号不能为空")
	}
	trans := [][]byte{[]byte("FindPrivateDataById"), []byte("id"), []byte(managedCardNo)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

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
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

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

func (t *FinancialChainCode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("金融机构id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL, id)
	if err != nil {
		return "", errors.New("金融机构私有数据查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("金融机构私有数据不存在，读到的%s对应的私有数据为空！", id)
	}
	return string(bytes), nil
}

/**
  根据金融机构名称查询
 */
func (t *FinancialChainCode) queryFinancialByName(ctx contractapi.TransactionContextInterface, name string) ([]*FinancialOrg, error) {

	queryString := fmt.Sprintf(`{"selector":{"name":"%s"}}`, name)

	queryResults, err := t.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, errors.New("金融机构名称查询失败" + err.Error())
	}

	return queryResults, nil
}

/**
  根据金融机构代码查询
 */
func (t *FinancialChainCode) queryFinancialByCode(ctx contractapi.TransactionContextInterface, code string) ([]*FinancialOrg, error) {

	queryString := fmt.Sprintf(`{"selector":{"code":"%s"}}`, code)

	queryResults, err := t.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, errors.New("金融机构代码查询失败" + err.Error())
	}

	return queryResults, nil
}
func (t *FinancialChainCode) getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*FinancialOrg, error) {
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	financialOrgs, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:%#v", financialOrgs)

	return financialOrgs, nil
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*FinancialOrg, error) {

	resp := []*FinancialOrg{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newRecord := new(FinancialOrg)
		err = json.Unmarshal(queryResponse.Value, newRecord)
		if err != nil {
			return nil, err
		}

		resp = append(resp, newRecord)
	}

	return resp, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(FinancialChainCode))
	if err != nil {
		fmt.Printf("Error create FinancialChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting FinancialChainCode chaincode: %s", err.Error())
	}
}

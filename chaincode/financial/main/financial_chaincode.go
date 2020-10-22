package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
   金融机构公管账户私有数据属性
 */
type FinancialOrgManagedAccountPrivateData struct {
	ID                    string  `json:"id"`                    //金融机构ID
	CardNo                string  `json:"cardNo"`                //金融机构公管账户账号
	FinancialOrgID        string  `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	PlatformOrgID         string  `json:"platformOrgID"`         //平台机构ID PlatformOrg.ID
	MerchantOrgID         string  `json:"merchantOrgID"`         //商户机构ID MerchantOrg.ID
	CurrentBalance        float64 `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance float64 `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int     `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

/**
   金融机构一般账户私有数据属性
 */
type FinancialOrgGeneralAccountPrivateData struct {
	ID              string  `json:"id"`              //金融机构ID
	CardNo          string  `json:"cardNo"`          //金融机构公管账户账号
	FinancialOrgID  string  `json:"financialOrgID"`  //金融机构ID FinancialOrg.ID
	CertificateNo   string  `json:"certificateNo"`   //个体证件号
	CertificateType string  `json:"certificateType"` //个体证件类型 (身份证/港澳台证/护照/军官证)¬
	CurrentBalance  float64 `json:"currentBalance"`  //金融机构共管账户余额(现金)
	AccStatus       int     `json:"accStatus"`       //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
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

func (t *FinancialChainCode) CreateManagedAccountToMerchant(ctx contractapi.TransactionContextInterface) (string, error) {

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
	var transientInput FinancialOrgManagedAccountPrivateData
	err = json.Unmarshal(individualPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(individualPrivateDataJsonBytes))
	}
	id := transientInput.ID
	if len(id) == 0 {
		return "", errors.New("金融机构ID不能为空")
	}
	if len(transientInput.CardNo) == 0 {
		return "", errors.New("金融机构公管账户账号不能为空")
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
	Avalbytes, err := ctx.GetStub().GetPrivateData("collectionFinancialMerchant", id)
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
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionPlatform", queryString)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	if resultsIterator != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	carAsBytes, _ := json.Marshal(transientInput)

	err = ctx.GetStub().PutPrivateData("transientInput", id, carAsBytes)
	if err != nil {
		return "", errors.New("商户公管账户保存失败" + err.Error())
	}
	return string(Avalbytes), nil
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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type AgencyOrgChainCode struct {
	contractapi.Contract
}

/**
 商户机构私有数据属性
 */
type MerchantOrg struct {
	ID          string `json:"id"`            //商户机构ID
	Name        string `json:"name"`          //商户机构名称
	AgencyOrgID string `json:"merchantOrgID"` //代理商机构ID AgencyOrg.ID
	Status      int    `json:"status"`        //金融机构状态(启用/禁用)
}

/**
 商户机构属性
 */
type MerchantOrgPrivateData struct {
	ID        string  `json:"id"`        //商户机构ID
	RateBasic float64 `json:"rateBasic"` //下发机构基础费率
}

func (t *AgencyOrgChainCode) Create(ctx contractapi.TransactionContextInterface, id string, name string, AgencyOrgID string) (string, error) {

	if len(id) == 0 {
		return "", errors.New("商户机构ID不能为空")
	}
	if len(name) == 0 {
		return "", errors.New("商户机构名称不能为空")
	}
	if len(AgencyOrgID) == 0 {
		return "", errors.New("代理商机构ID不能为空")
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
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}
	financialPrivateDataJsonBytes, ok := transMap["agency"]
	if !ok {
		return "", errors.New("financial must be a key in the transient map")
	}

	if len(financialPrivateDataJsonBytes) == 0 {
		return "", errors.New("financial value in the transient map must be a non-empty JSON string")
	}
	var transientInput MerchantOrgPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	if len(transientInput.ID) == 0 {
		return "", errors.New("下发机构ID不能为空")
	}
	if transientInput.RateBasic == 0 {
		return "", errors.New("下发机构基础费率不能为0")
	}
	err = ctx.GetStub().PutPrivateData("collectionAgency",id, financialPrivateDataJsonBytes)

	financial := &MerchantOrg{
		ID:     id,
		Name:   name,
		Status: 0,
	}
	carAsBytes, _ := json.Marshal(financial)

	err = ctx.GetStub().PutState(id, carAsBytes)

	if err != nil {
		return "", errors.New("金融机构保存失败" + err.Error())
	}
	return string(Avalbytes), nil
}

func (t *AgencyOrgChainCode) FindById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
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

func (t *AgencyOrgChainCode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string, collectionName string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("金融机构id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData("collectionAgency", id)
	if err != nil {
		return "", errors.New("金融机构私有数据查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("金融机构私有数据不存在，读到的%s对应的私有数据为空！", id)
	}
	return string(bytes), nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(AgencyOrgChainCode))
	if err != nil {
		fmt.Printf("Error create PlatformChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting PlatformChainCode chaincode: %s", err.Error())
	}
}

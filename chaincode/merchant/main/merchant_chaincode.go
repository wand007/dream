package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MerchantOrgChainCode struct {
	contractapi.Contract
}

//私有数据集名称
const COLLECTION_MERCHANT string = "collectionMerchant"

/**
 商户机构私有数据属性
 */
type MerchantOrg struct {
	ID     string `json:"id"`     //商户机构ID
	Name   string `json:"name"`   //商户机构名称
	Status int    `json:"status"` //金融机构状态(启用/禁用)
}

/**
 商户机构属性
 */
type MerchantOrgPrivateData struct {
	ID          string  `json:"id"`          //商户机构ID
	AgencyOrgID string  `json:"agencyOrgID"` //代理商机构ID AgencyOrg.ID
	RateBasic   float64 `json:"rateBasic"`   //下发机构基础费率
}

func (t *MerchantOrgChainCode) InitLedger(ctx contractapi.TransactionContextInterface) error {

	fmt.Println("IssueChaincode Init")
	//公开数据
	issueOrgs := []MerchantOrg{
		{ID: "M766005404604841984", Name: "默认商户机构1", Status: 1},
		{ID: "M764441096829812736", Name: "默认商户机构2", Status: 1},
	}
	for _, asset := range issueOrgs {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
		//私有数据
		agencyOrgPrivateDatas := []MerchantOrgPrivateData{
			{ID: asset.ID, AgencyOrgID: "A766005404604841984", RateBasic: 0.62},
			{ID: asset.ID, AgencyOrgID: "A766374712807800832", RateBasic: 0.62},
		}
		for _, privateData := range agencyOrgPrivateDatas {
			privateDataJSON, err := json.Marshal(privateData)
			if err != nil {
				return err
			}

			err = ctx.GetStub().PutPrivateData(COLLECTION_MERCHANT, asset.ID, privateDataJSON)
			if err != nil {
				return fmt.Errorf("Failed to put to world state. %s", err.Error())
			}
		}
	}

	return nil
}

/**
   新建商户机构
 */
func (t *MerchantOrgChainCode) Create(ctx contractapi.TransactionContextInterface, id string, name string, agencyOrgID string) (string, error) {
	//公有数据入参参数
	if len(id) == 0 {
		return "", errors.New("商户机构ID不能为空")
	}
	if len(name) == 0 {
		return "", errors.New("商户机构名称不能为空")
	}
	if len(agencyOrgID) == 0 {
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
	//私有数据入参参数
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}
	financialPrivateDataJsonBytes, ok := transMap["merchant"]
	if !ok {
		return "", errors.New("financial must be a key in the transient map")
	}

	if len(financialPrivateDataJsonBytes) == 0 {
		return "", errors.New("financial value in the transient map must be a non-empty JSON string")
	}
	var merchantOrgPrivateData MerchantOrgPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &merchantOrgPrivateData)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	if len(merchantOrgPrivateData.AgencyOrgID) == 0 {
		return "", errors.New("代理商机构ID不能为空")
	}
	if merchantOrgPrivateData.RateBasic == 0 {
		return "", errors.New("下发机构基础费率不能为0")
	}
	//防重复提交
	// Get the state from the ledger
	Avalbytes, err = ctx.GetStub().GetPrivateData(COLLECTION_MERCHANT, id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
		return "", errors.New(jsonResp)
	}
	//Mongo Query string 语法
	queryString := fmt.Sprintf(`{"selector":{"name":"%s"}}`, name)
	// 富查询的返回结果可能为多条 所以这里返回的是一个迭代器 需要我们进一步的处理来获取需要的结果
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	if resultsIterator != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	//公开数据
	merchantOrg := MerchantOrg{ID: id, Name: name, Status: 1}

	carAsBytes, _ := json.Marshal(merchantOrg)
	err = ctx.GetStub().PutState(merchantOrg.ID, carAsBytes)

	if err != nil {
		return "", err
	}

	//私有数据
	merchantOrgPrivateData.ID = id
	merchantOrgPrivateDataAsBytes, _ := json.Marshal(merchantOrgPrivateData)
	err = ctx.GetStub().PutPrivateData(COLLECTION_MERCHANT, merchantOrgPrivateData.ID, merchantOrgPrivateDataAsBytes)

	if err != nil {
		return "", err
	}
	return merchantOrg.ID, nil
}

func (t *MerchantOrgChainCode) FindById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
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

func (t *MerchantOrgChainCode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string, collectionName string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("金融机构id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_MERCHANT, id)
	if err != nil {
		return "", errors.New("金融机构私有数据查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("金融机构私有数据不存在，读到的%s对应的私有数据为空！", id)
	}
	return string(bytes), nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(MerchantOrgChainCode))
	if err != nil {
		fmt.Printf("Error create PlatformChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting PlatformChainCode chaincode: %s", err.Error())
	}
}

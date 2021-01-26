package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type IssueChaincode struct {
	contractapi.Contract
}

//私有数据集名称
const COLLECTION_ISSUE string = "collectionIssue"

/**
  下发机构属性
*/
type IssueOrg struct {
	ID     string `json:"id"`     //下发机构ID
	Name   string `json:"name"`   //下发机构名称
	Status int    `json:"status"` //下发机构状态(启用/禁用)
}

/**
  下发机构私有数据属性
*/
type IssueOrgPrivateData struct {
	ID        string  `json:"id"`        //下发机构ID IssueOrg.ID
	RateBasic float64 `json:"rateBasic"` //下发机构基础费率
}

func (t *IssueChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("IssueChaincode Init")
	//公开数据
	issueOrgs := []IssueOrg{
		{ID: "I766005404604841984", Name: "默认下发机构1", Status: 1},
		{ID: "I764441096829812736", Name: "默认下发机构2", Status: 1},
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
		issueOrgPrivateDatas := IssueOrgPrivateData{ID: asset.ID, RateBasic: 0.6}
		assetJSON, err = json.Marshal(issueOrgPrivateDatas)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutPrivateData(COLLECTION_ISSUE, asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

/**
  新增下发机构共管账户私有数据
*/
func (t *IssueChaincode) Create(ctx contractapi.TransactionContextInterface, id string, name string) (string, error) {
	//公有数据入参参数
	if len(id) == 0 {
		return "", errors.New("下发机构ID不能为空")
	}
	if len(name) == 0 {
		return "", errors.New("下发机构名称不能为空")
	}
	//私有数据入参参数
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}
	financialPrivateDataJsonBytes, ok := transMap["issue"]
	if !ok {
		return "", errors.New("financial must be a key in the transient map")
	}

	if len(financialPrivateDataJsonBytes) == 0 {
		return "", errors.New("financial value in the transient map must be a non-empty JSON string")
	}
	var transientInput IssueOrgPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	if transientInput.RateBasic == 0 {
		return "", errors.New("下发机构基础费率不能为0")
	}
	//防重复提交
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetPrivateData(COLLECTION_ISSUE, id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return "", errors.New(jsonResp)
	}
	fmt.Printf("Avalbytes:%v", Avalbytes)
	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	// Mongo Query string 语法见上文链接
	queryString := fmt.Sprintf(`{"selector":{"name":"%s"}}`, name)
	// 富查询的返回结果可能为多条 所以这里返回的是一个迭代器 需要我们进一步的处理来获取需要的结果
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + name + "\"}"
		return "", errors.New(jsonResp)
	}
	defer resultsIterator.Close()

	if resultsIterator.HasNext() {
		jsonResp := "{\"Error\":\"私有数据已存在 " + name + "\"}"
		return "", errors.New(jsonResp)
	}
	//公开数据
	issueOrg := IssueOrg{ID: id, Name: name, Status: 1}

	carAsBytes, _ := json.Marshal(issueOrg)
	err = ctx.GetStub().PutState(issueOrg.ID, carAsBytes)

	if err != nil {
		return "", fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	//私有数据
	transientInput.ID = id
	retailerOrgPrivateData := IssueOrgPrivateData{ID: issueOrg.ID, RateBasic: transientInput.RateBasic}

	retailerOrgPrivateDataAsBytes, _ := json.Marshal(retailerOrgPrivateData)
	err = ctx.GetStub().PutPrivateData(COLLECTION_ISSUE, retailerOrgPrivateData.ID, retailerOrgPrivateDataAsBytes)

	if err != nil {
		return "", fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	return id, nil
}
func (t *IssueChaincode) FindById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("下发机构id不能为空")
	}
	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", errors.New("下发机构查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("下发机构数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

func (t *IssueChaincode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("共管账户id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_ISSUE, id)
	if err != nil {
		return "", errors.New("共管账户查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("共管账户数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(IssueChaincode))
	if err != nil {
		fmt.Printf("Error create IssueChaincode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting IssueChaincode chaincode: %s", err.Error())
	}
}

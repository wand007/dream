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

//私有数据集名称
const COLLECTION_AGENCY string = "collectionAgency"

/**
分销商机构属性
*/
type AgencyOrg struct {
	ID                      string `json:"id"`                      //分销商机构ID
	Name                    string `json:"name"`                    //分销商机构名称
	UnifiedSocialCreditCode string `json:"unifiedSocialCreditCode"` //统一社会信用代码
	Status                  int    `json:"status"`                  //分销商机构状态(启用/禁用)
}

/**
分销商机构私有数据属性
*/
type AgencyOrgPrivateData struct {
	ID         string  `json:"id"`         //分销商机构ID
	IssueOrgID string  `json:"issueOrgID"` //下发机构ID IssueOrg.ID
	RateBasic  float64 `json:"rateBasic"`  //分销商机构基础费率
}

func (t *AgencyOrgChainCode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("AgencyOrgChainCode Init")
	//公开数据
	financialOrgs := []AgencyOrg{
		{ID: "A766005404604841984", Name: "分销商机构1", UnifiedSocialCreditCode: "92370112MA3F23MB5N", Status: 1},
		{ID: "A766374712807800832", Name: "分销商机构2", UnifiedSocialCreditCode: "92370104MA3DR08A4D", Status: 1},
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
		agencyOrgPrivateDatas := []AgencyOrgPrivateData{
			{ID: asset.ID, IssueOrgID: "I766005404604841984", RateBasic: 0.61},
			{ID: asset.ID, IssueOrgID: "I764441096829812736", RateBasic: 0.61},
		}
		for _, privateData := range agencyOrgPrivateDatas {
			privateDataJSON, err := json.Marshal(privateData)
			if err != nil {
				return err
			}

			err = ctx.GetStub().PutPrivateData(COLLECTION_AGENCY, asset.ID, privateDataJSON)
			if err != nil {
				return fmt.Errorf("Failed to put to world state. %s", err.Error())
			}
		}
	}
	return nil
}

/**
  新增代理机构共管账户私有数据
*/
func (t *AgencyOrgChainCode) Create(ctx contractapi.TransactionContextInterface, id string, name string, unifiedSocialCreditCode string) (string, error) {
	//公有数据入参参数
	if len(id) == 0 {
		return "", errors.New("分销商机构ID不能为空")
	}
	if len(name) == 0 {
		return "", errors.New("分销商机构名称不能为空")
	}
	if len(unifiedSocialCreditCode) == 0 {
		return "", errors.New("统一社会信用代码不能为空")
	}
	//私有数据入参参数
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
	var transientInput AgencyOrgPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	if len(transientInput.IssueOrgID) == 0 {
		return "", errors.New("下发机构ID不能为空")
	}
	if transientInput.RateBasic == 0 {
		return "", errors.New("分销商机构基础费率不能为0")
	}
	//防重复提交
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"公开数据已存在" + id + "\"}"
		return "", errors.New(jsonResp)
	}

	queryString := fmt.Sprintf(`{"selector":{"name":"%s"}}`, name)
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
	financial := &AgencyOrg{
		ID:                      id,
		Name:                    name,
		UnifiedSocialCreditCode: unifiedSocialCreditCode,
		Status:                  0,
	}

	carAsBytes, _ := json.Marshal(financial)
	err = ctx.GetStub().PutState(financial.ID, carAsBytes)

	if err != nil {
		return "", fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	//私有数据
	transientInput.ID = id
	retailerOrgPrivateDataAsBytes, _ := json.Marshal(transientInput)
	err = ctx.GetStub().PutPrivateData(COLLECTION_AGENCY, transientInput.ID, retailerOrgPrivateDataAsBytes)

	if err != nil {
		return "", fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	return id, nil
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

func (t *AgencyOrgChainCode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("金融机构id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_AGENCY, id)
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
		fmt.Printf("Error create AgencyOrgChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting AgencyOrgChainCode chaincode: %s", err.Error())
	}
}

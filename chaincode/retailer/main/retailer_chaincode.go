package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RetailerOrgChainCode struct {
	contractapi.Contract
}

//私有数据集名称
const COLLECTION_RETAILER string = "collectionRetailer"

/**
 零售商机构数据属性
 */
type RetailerOrg struct {
	ID                      string `json:"id"`                      //零售商机构ID
	Name                    string `json:"name"`                    //零售商机构名称
	UnifiedSocialCreditCode string `json:"unifiedSocialCreditCode"` //统一社会信用代码
	Status                  int    `json:"status"`                  //零售商机构状态(启用/禁用)
}

/**
 零售商机构私有属性
 */
type RetailerOrgPrivateData struct {
	ID          string  `json:"id"`          //零售商机构ID
	AgencyOrgID string  `json:"agencyOrgID"` //分销商机构ID AgencyOrg.ID
	RateBasic   float64 `json:"rateBasic"`   //下发机构基础费率
}

func (t *RetailerOrgChainCode) InitLedger(ctx contractapi.TransactionContextInterface) error {

	fmt.Println("RetailerOrgChainCode Init")
	//公开数据
	issueOrgs := []RetailerOrg{
		{ID: "M766005404604841984", Name: "默认零售商机构1", UnifiedSocialCreditCode: "91370181MA3D7J9W3W", Status: 1},
		{ID: "M764441096829812736", Name: "默认零售商机构2", UnifiedSocialCreditCode: "91370100MA3MP74K7A", Status: 1},
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
		agencyOrgPrivateDatas := []RetailerOrgPrivateData{
			{ID: asset.ID, AgencyOrgID: "A766005404604841984", RateBasic: 0.62},
			{ID: asset.ID, AgencyOrgID: "A766374712807800832", RateBasic: 0.62},
		}
		for _, privateData := range agencyOrgPrivateDatas {
			privateDataJSON, err := json.Marshal(privateData)
			if err != nil {
				return err
			}

			err = ctx.GetStub().PutPrivateData(COLLECTION_RETAILER, asset.ID, privateDataJSON)
			if err != nil {
				return fmt.Errorf("Failed to put to world state. %s", err.Error())
			}
		}
	}

	return nil
}

/**
   新建零售商机构
 */
func (t *RetailerOrgChainCode) Create(ctx contractapi.TransactionContextInterface, id string, name string, unifiedSocialCreditCode string, agencyOrgID string) (string, error) {
	//公有数据入参参数
	if len(id) == 0 {
		return "", errors.New("零售商机构ID不能为空")
	}
	if len(name) == 0 {
		return "", errors.New("零售商机构名称不能为空")
	}
	if len(agencyOrgID) == 0 {
		return "", errors.New("分销商机构ID不能为空")
	}
	if len(unifiedSocialCreditCode) == 0 {
		return "", errors.New("统一社会信用代码不能为空")
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
	financialPrivateDataJsonBytes, ok := transMap["retailer"]
	if !ok {
		return "", errors.New("financial must be a key in the transient map")
	}

	if len(financialPrivateDataJsonBytes) == 0 {
		return "", errors.New("RetailerOrgChainCode value in the transient map must be a non-empty JSON string")
	}
	var retailerOrgPrivateData RetailerOrgPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &retailerOrgPrivateData)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}
	if len(retailerOrgPrivateData.AgencyOrgID) == 0 {
		return "", errors.New("分销商机构ID不能为空")
	}
	if retailerOrgPrivateData.RateBasic == 0 {
		return "", errors.New("下发机构基础费率不能为0")
	}
	//防重复提交
	// Get the state from the ledger
	Avalbytes, err = ctx.GetStub().GetPrivateData(COLLECTION_RETAILER, id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + id + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
		return "", errors.New(jsonResp)
	}
	//Mongo Query string 语法
	queryString := fmt.Sprintf(`{"selector":{"unifiedSocialCreditCode":"%s"}}`, unifiedSocialCreditCode)
	// 富查询的返回结果可能为多条 所以这里返回的是一个迭代器 需要我们进一步的处理来获取需要的结果
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + unifiedSocialCreditCode + "\"}"
		return "", errors.New(jsonResp)
	}
	defer resultsIterator.Close()

	if resultsIterator.HasNext() {
		jsonResp := "{\"Error\":\"私有数据已存在 " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	//公开数据
	retailerOrg := RetailerOrg{ID: id, Name: name, UnifiedSocialCreditCode: unifiedSocialCreditCode, Status: 1}

	carAsBytes, _ := json.Marshal(retailerOrg)
	err = ctx.GetStub().PutState(retailerOrg.ID, carAsBytes)

	if err != nil {
		return "", err
	}

	//私有数据
	retailerOrgPrivateData.ID = id
	retailerOrgPrivateDataAsBytes, _ := json.Marshal(retailerOrgPrivateData)
	err = ctx.GetStub().PutPrivateData(COLLECTION_RETAILER, retailerOrgPrivateData.ID, retailerOrgPrivateDataAsBytes)

	if err != nil {
		return "", err
	}
	return retailerOrg.ID, nil
}

func (t *RetailerOrgChainCode) FindById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
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

func (t *RetailerOrgChainCode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("金融机构id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_RETAILER, id)
	if err != nil {
		return "", errors.New("金融机构私有数据查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("金融机构私有数据不存在，读到的%s对应的私有数据为空！", id)
	}
	return string(bytes), nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(RetailerOrgChainCode))
	if err != nil {
		fmt.Printf("Error create RetailerOrgChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting RetailerOrgChainCode chaincode: %s", err.Error())
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

type IndividualChainCode struct {
	contractapi.Contract
}

//私有数据集名称
const COLLECTION_INDIVIDUAL string = "collectionIndividual"

const (
	//默认
	CERTIFICATE_TYPE_0 int = 0
	//身份证
	CERTIFICATE_TYPE_1 int = 1
	//港澳台证
	CERTIFICATE_TYPE_2 int = 2
	//护照
	CERTIFICATE_TYPE_3 int = 3
	//军官证
	CERTIFICATE_TYPE_4 int = 4
	//统一社会信用代码
	CERTIFICATE_TYPE_5 int = 5
)

/**
 个体属性
 */
type Individual struct {
	ID            string `json:"id"`            //个体ID
	Name          string `json:"name"`          //个体名称
	PlatformOrgID string `json:"platformOrgID"` //平台机构ID PlatformOrg.ID
	Status        int    `json:"status"`        //个体状态(启用/禁用)
}

/**
 个体属性私有数据属性
 */
type IndividualPrivateData struct {
	ID              string `json:"id"`              //个体ID Individual.ID
	CertificateNo   string `json:"certificateNo"`   //个体证件号
	CertificateType int    `json:"certificateType"` //个体证件类型 (身份证/港澳台证/护照/军官证)
}

type IndividualTransientInput struct {
	ID              string `json:"id"`              //个体ID
	Name            string `json:"name"`            //个体名称
	PlatformOrgID   string `json:"platformOrgID"`   //平台机构ID PlatformOrg.ID
	CertificateNo   string `json:"certificateNo"`   //个体证件号
	CertificateType int    `json:"certificateType"` //个体证件类型 (身份证/港澳台证/护照/军官证)
	Status          int    `json:"status"`          //个体状态(启用/禁用)
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Record              *Individual
	TxId                string    `json:"txId"`
	Timestamp           time.Time `json:"timestamp"`
	FetchedRecordsCount int       `json:"fetchedRecordsCount"`
	Bookmark            string    `json:"bookmark"`
}

func (t *IndividualChainCode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("IndividualChainCode Init")
	//公开数据
	retailerOrg := Individual{ID: "IN760934239574175744", Name: "默认个体", PlatformOrgID: "P768877118787432448", Status: 1}

	carAsBytes, _ := json.Marshal(retailerOrg)
	err := ctx.GetStub().PutState(retailerOrg.ID, carAsBytes)

	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	//私有数据
	retailerOrgPrivateData := IndividualPrivateData{ID: retailerOrg.ID, CertificateNo: "888888888888888888", CertificateType: CERTIFICATE_TYPE_1}

	retailerOrgPrivateDataAsBytes, _ := json.Marshal(retailerOrgPrivateData)
	err = ctx.GetStub().PutPrivateData(COLLECTION_INDIVIDUAL, retailerOrgPrivateData.ID, retailerOrgPrivateDataAsBytes)

	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	return nil
}

/**
  新增个体数据
 */
func (t *IndividualChainCode) Create(ctx contractapi.TransactionContextInterface) (string, error) {

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

	var individualTransientInput IndividualTransientInput
	err = json.Unmarshal(individualPrivateDataJsonBytes, &individualTransientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(individualPrivateDataJsonBytes))
	}
	id := individualTransientInput.ID
	if len(id) == 0 {
		return "", errors.New("个体id不能为空")
	}
	if len(individualTransientInput.Name) == 0 {
		return "", errors.New("个体名称不能为空")
	}
	if len(individualTransientInput.PlatformOrgID) == 0 {
		return "", errors.New("平台机构ID不能为空")
	}
	//私有数据
	if len(individualTransientInput.CertificateNo) == 0 {
		return "", errors.New("个体证件号不能为空")
	}

	//公开数据防重复添加
	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", errors.New("个体查询失败！")
	}
	if bytes == nil {
		fmt.Println("个体公开数据已存在，读到的" + id + "对应的数据不为空！")
		// ==== Create marble object, marshal to JSON, and save to state ====
		individual := &Individual{
			ID:            individualTransientInput.ID,
			Name:          individualTransientInput.Name,
			PlatformOrgID: individualTransientInput.PlatformOrgID,
			Status:        individualTransientInput.Status,
		}

		carAsBytes, _ := json.Marshal(individual)
		err = ctx.GetStub().PutState(id, carAsBytes)
		if err != nil {
			return "", errors.New("个体公开数据保存失败" + err.Error())
		}
	}
	//私有数据防重复添加
	bytes, err = ctx.GetStub().GetPrivateData(COLLECTION_INDIVIDUAL, id)
	if err != nil {
		return "", errors.New("个体私有数据查询失败！")
	}
	if bytes != nil {
		return "", errors.New("个体私有数据已存在，读到的" + id + "对应的私有数据不为空！")
	}
	individualPrivateData := &IndividualPrivateData{
		ID:              individualTransientInput.ID,
		CertificateNo:   individualTransientInput.CertificateNo,
		CertificateType: individualTransientInput.CertificateType,
	}
	carAsBytes, _ := json.Marshal(individualPrivateData)
	err = ctx.GetStub().PutPrivateData(COLLECTION_INDIVIDUAL, id, carAsBytes)
	if err != nil {
		return "", errors.New("个体私有数据保存失败" + err.Error())
	}
	ctx.GetStub().SetEvent("RetailerOrg", carAsBytes)
	return id, nil
}

func (t *IndividualChainCode) Update(ctx contractapi.TransactionContextInterface) (string, error) {

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
	var individualTransientInput IndividualTransientInput
	err = json.Unmarshal(individualPrivateDataJsonBytes, &individualTransientInput)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(individualPrivateDataJsonBytes))
	}
	id := individualTransientInput.ID;
	if len(id) == 0 {
		return "", errors.New("个体id不能为空")
	}

	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", errors.New("个体查询失败！")
	}
	if bytes == nil {
		fmt.Println("个体公开数据不存在，读到的" + id + "对应的数据不存在！")
	}
	var individual Individual
	err = json.Unmarshal(bytes, &individual)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(individualPrivateDataJsonBytes))
	}

	if len(individualTransientInput.Name) != 0 {
		individual.Name = individualTransientInput.Name
	}
	if len(individualTransientInput.PlatformOrgID) == 0 {
		individual.PlatformOrgID = individualTransientInput.PlatformOrgID
	}
	if individualTransientInput.Status != 0 {
		individual.Status = individualTransientInput.Status
	}

	carAsBytes, _ := json.Marshal(individual)
	err = ctx.GetStub().PutState(id, carAsBytes)
	if err != nil {
		return "", errors.New("个体公开数据更新失败" + err.Error())
	}
	//私有数据
	bytes, err = ctx.GetStub().GetPrivateData(COLLECTION_INDIVIDUAL, id)
	if err != nil {
		return "", errors.New("个体私有数据查询失败！")
	}
	if bytes == nil {
		fmt.Println("个体私有数据不存在，读到的" + id + "对应的数据不存在！")
	}
	var individualPrivateData IndividualPrivateData
	err = json.Unmarshal(bytes, &individual)
	if err != nil {
		return "", errors.New("Failed to decode JSON of: " + string(individualPrivateDataJsonBytes))
	}

	if len(individualPrivateData.CertificateNo) == 0 {
		individualPrivateData.CertificateNo = individualTransientInput.CertificateNo
	}
	if individualTransientInput.CertificateType != 0 {
		individualPrivateData.CertificateType = individualTransientInput.CertificateType
	}

	carAsBytes, _ = json.Marshal(individualPrivateData)
	err = ctx.GetStub().PutPrivateData(COLLECTION_INDIVIDUAL, id, carAsBytes)
	if err != nil {
		return "", errors.New("个体私有数据更新失败" + err.Error())
	}
	return id, nil
}

func (t *IndividualChainCode) FindById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("个体id不能为空")
	}
	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", errors.New("个体查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("个体数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

func (t *IndividualChainCode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("个体id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_INDIVIDUAL, id)
	if err != nil {
		return "", errors.New("个体私有数据查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("个体私有数据不存在，读到的%s对应的私有数据为空！", id)
	}
	return string(bytes), nil
}

func (t *IndividualChainCode) QueryIndividualSimpleWithPagination(ctx contractapi.TransactionContextInterface, queryString, bookmark string, pageSize int) ([]*QueryResult, error) {
	fmt.Println("个体公开数据分页查询入参参数 %s,%s,%s", queryString, bookmark, pageSize)
	if len(queryString) == 0 {
		return nil, errors.New("查询条件不能为空")
	}

	return getQueryResultForQueryStringWithPagination(ctx, queryString, int32(pageSize), bookmark)
}

// getQueryResultForQueryStringWithPagination executes the passed in query string with
// pagination info. Result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) ([]*QueryResult, error) {

	resultsIterator, _, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

// constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*QueryResult, error) {

	resp := []*QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newRecord := new(QueryResult)
		err = json.Unmarshal(queryResponse.Value, newRecord)
		if err != nil {
			return nil, err
		}

		resp = append(resp, newRecord)
	}

	return resp, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(IndividualChainCode))
	if err != nil {
		fmt.Printf("Error create IndividualChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting IndividualChainCode chaincode: %s", err.Error())
	}
}

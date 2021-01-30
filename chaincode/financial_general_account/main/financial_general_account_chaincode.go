package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

/**
初始化全部角色的一般账户记录
新建一般账户记录
一般账户现金交易
一般账户票据交易
一般账户票据兑换现金交易
*/
/**
需要给不同的角色设置不同的私有数据集
*/
/**
一般账户链码
*/
type FinancialGeneralAccountChaincode struct {
	contractapi.Contract
}

//一般账户私有数据集名称
const COLLECTION_FINANCIAL_GENERAL_ACCOUNT string = "collectionFinancialGeneralAccount"

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
  金融机构一般账户私有数据属性
*/
type FinancialOrgGeneralAccountPrivateData struct {
	CardNo                string `json:"cardNo"`                //金融机构共管账户账号        //金融机构公管账户账号(唯一不重复)
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	CertificateNo         string `json:"certificateNo"`         //持卡者证件号
	CertificateType       int    `json:"certificateType"`       //持卡者证件类型 (身份证/港澳台证/护照/军官证/统一社会信用代码)
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构零售商机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Record              *FinancialOrgGeneralAccountPrivateData
	TxId                string    `json:"txId"`
	Timestamp           time.Time `json:"timestamp"`
	FetchedRecordsCount int       `json:"fetchedRecordsCount"`
	Bookmark            string    `json:"bookmark"`
}

/**
初始化金融机构
*/
func (t *FinancialGeneralAccountChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("FinancialGeneralAccountChaincode Init")
	err := t.InitIndividualsLedger(ctx)
	if err != nil {
		return fmt.Errorf("InitIndividualsLedger Failed to put to world state. %s", err.Error())
	}
	err = t.InitRetailersLedger(ctx)
	if err != nil {
		return fmt.Errorf("InitRetailersLedger Failed to put to world state. %s", err.Error())
	}
	err = t.InitAgencyLedger(ctx)
	if err != nil {
		return fmt.Errorf("InitAgencyLedger Failed to put to world state. %s", err.Error())
	}
	err = t.InitIssuesLedger(ctx)
	if err != nil {
		return fmt.Errorf("InitIssuesLedger Failed to put to world state. %s", err.Error())
	}

	return nil
}

/**
初始化金融机构一般账户----个体
*/
func (t *FinancialGeneralAccountChaincode) InitIndividualsLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("InitIndividualsLedger Init")

	//个体
	individuals := []FinancialOrgGeneralAccountPrivateData{
		{CardNo: "6229486603953152819", FinancialOrgID: "F766005404604841984", CertificateNo: "888888888888888888", CertificateType: CERTIFICATE_TYPE_1, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953152825", FinancialOrgID: "F766374712807800832", CertificateNo: "888888888888888888", CertificateType: CERTIFICATE_TYPE_1, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
	}
	//私有数据
	for _, asset := range individuals {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, asset.CardNo, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

/**
初始化金融机构一般账户 ----零售商机构
*/
func (t *FinancialGeneralAccountChaincode) InitRetailersLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("InitRetailersLedger Init")

	//零售商机构
	retailers := []FinancialOrgGeneralAccountPrivateData{
		{CardNo: "6229486603953174011", FinancialOrgID: "F766005404604841984", CertificateNo: "92370104MA3DR08A4D", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953174027", FinancialOrgID: "F766374712807800832", CertificateNo: "92370104MA3DR08A4D", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
	}
	//私有数据
	for _, asset := range retailers {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, asset.CardNo, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

/**
初始化金融机构一般账户 ----分销商机构
*/
func (t *FinancialGeneralAccountChaincode) InitAgencyLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("InitAgencyLedger Init")

	//分销商机构
	agencys := []FinancialOrgGeneralAccountPrivateData{
		{CardNo: "6229486603953188912", FinancialOrgID: "F766005404604841984", CertificateNo: "92370112MA3F23MB5N", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953188928", FinancialOrgID: "F766374712807800832", CertificateNo: "92370112MA3F23MB5N", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
	}
	//私有数据
	for _, asset := range agencys {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, asset.CardNo, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

/**
初始化金融机构一般账户 ----下发机构
*/
func (t *FinancialGeneralAccountChaincode) InitIssuesLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("InitIssuesLedger Init")

	//下发机构
	issues := []FinancialOrgGeneralAccountPrivateData{
		{CardNo: "6229486603953201814", FinancialOrgID: "F766005404604841984", CertificateNo: "91370181MA3D7J9W3W", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
		{CardNo: "6229488603953201820", FinancialOrgID: "F766374712807800832", CertificateNo: "91370181MA3D7J9W3W", CertificateType: CERTIFICATE_TYPE_5, CurrentBalance: 0, VoucherCurrentBalance: 0, AccStatus: 1},
	}
	//私有数据
	for _, asset := range issues {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, asset.CardNo, assetJSON)
		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

/**
新增金融机构一般账户私有数据
*/
func (t *FinancialGeneralAccountChaincode) Create(ctx contractapi.TransactionContextInterface) (string, error) {

	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", errors.New("Error getting transient: " + err.Error())
	}

	financialPrivateDataJsonBytes, ok := transMap["generalAccount"]
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
	if len(transientInput.CardNo) == 0 {
		return "", errors.New("金融机构一般账户账号不能为空")
	}
	if len(transientInput.FinancialOrgID) == 0 {
		return "", errors.New("金融机构ID不能为空")
	}
	if len(transientInput.CertificateNo) == 0 {
		return "", errors.New("持卡者证件号不能为空")
	}
	if transientInput.CertificateType == 0 {
		return "", errors.New("持卡者证件类型不能为空")
	}
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, transientInput.CardNo)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + transientInput.CardNo + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes != nil {
		jsonResp := "{\"Error\":\"一般账户账号:" + transientInput.CardNo + "不能重复\"}"
		return "", errors.New(jsonResp)
	}

	carAsBytes, _ := json.Marshal(transientInput)

	err = ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, transientInput.CardNo, carAsBytes)
	if err != nil {
		return "", errors.New("零售商共管账户保存失败" + err.Error())
	}
	return "", nil
}

/**
  现金交易
零售商向零售商一般账户充值现金余额
*/
func (t *FinancialGeneralAccountChaincode) TransferCashAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, amount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	if amount < 0 {
		return errors.New("一般账户充值金额不能小于0")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo)
	if err != nil {
		return errors.New("一般账户查询失败！")
	}
	if financialPrivateDataJsonBytes == nil {
		return fmt.Errorf("一般账户数据不存在，读到的%s对应的数据为空！", generalCardNo)
	}
	var transientInput FinancialOrgGeneralAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}

	newCurrentBalance := transientInput.CurrentBalance + amount
	if newCurrentBalance < 0 {
		return errors.New("一般账户现金余额不足")
	}
	transientInput.CurrentBalance = newCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo, assetJSON)
}

/**
  票据交易
增加个体/零售商/分销商的票据
*/
func (t *FinancialGeneralAccountChaincode) TransferVoucherAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, voucherAmount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo)
	if err != nil {
		return errors.New("一般账户查询失败！")
	}
	if financialPrivateDataJsonBytes == nil {
		return fmt.Errorf("一般账户数据不存在，读到的%s对应的数据为空！", generalCardNo)
	}
	var transientInput FinancialOrgGeneralAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}

	newVoucherCurrentBalance := transientInput.VoucherCurrentBalance + voucherAmount
	if newVoucherCurrentBalance < 0 {
		return errors.New("一般账户票据余额不足")
	}
	transientInput.VoucherCurrentBalance = newVoucherCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo, assetJSON)
}

/**
  现金和票据交易 （票据提现）
提现时增加个体/零售商/分销商的现金
提现时减少个体/零售商/分销商的票据
*/
func (t *FinancialGeneralAccountChaincode) TransferAsset(ctx contractapi.TransactionContextInterface, generalCardNo string, voucherAmount int) error {
	if len(generalCardNo) == 0 {
		return errors.New("一般账户卡号不能为空")
	}
	financialPrivateDataJsonBytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo)
	if err != nil {
		return errors.New("一般账户查询失败！")
	}
	if financialPrivateDataJsonBytes == nil {
		return fmt.Errorf("一般账户数据不存在，读到的%s对应的数据为空！", generalCardNo)
	}
	var transientInput FinancialOrgGeneralAccountPrivateData
	err = json.Unmarshal(financialPrivateDataJsonBytes, &transientInput)
	if err != nil {
		return errors.New("Failed to decode JSON of: " + string(financialPrivateDataJsonBytes))
	}

	newVoucherCurrentBalance := transientInput.VoucherCurrentBalance + voucherAmount
	if newVoucherCurrentBalance < 0 {
		return errors.New("一般账户票据余额不足")
	}
	newCurrentBalance := transientInput.CurrentBalance - voucherAmount
	if newCurrentBalance < 0 {
		return errors.New("一般账户余额不足")
	}
	transientInput.VoucherCurrentBalance = newVoucherCurrentBalance
	transientInput.CurrentBalance = newCurrentBalance
	assetJSON, _ := json.Marshal(transientInput)
	return ctx.GetStub().PutPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, generalCardNo, assetJSON)
}

/**
查询金融机构一般账户私有数据
*/
func (t *FinancialGeneralAccountChaincode) FindPrivateDataById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("一般账户id不能为空")
	}
	bytes, err := ctx.GetStub().GetPrivateData(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, id)
	if err != nil {
		return "", errors.New("一般账户查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("一般账户数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

/**
查询全部
*/
func (t *FinancialGeneralAccountChaincode) QueryFinancialGeneralWithPagination(ctx contractapi.TransactionContextInterface, financialOrgID, certificateNo string, bookmark string, pageSize int) ([]*QueryResult, error) {
	if len(certificateNo) == 0 {
		return nil, errors.New("证件号查询条件不能为空")
	}

	queryString := fmt.Sprintf(`{"selector":{"financialOrgID":"%s"} and "certificateNo":"%s"}}`, financialOrgID, certificateNo)

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

// verify client org id and matches peer org id.
func verifyClientOrgMatchesPeerOrg(clientOrgID string) error {
	peerOrgID, err := shim.GetMSPID()
	if err != nil {
		return fmt.Errorf("failed getting peer's orgID: %s", err.Error())
	}
	fmt.Println("client from org %s authorized to read or write private data from an org %s peer", clientOrgID, peerOrgID)
	if clientOrgID != peerOrgID {
		return fmt.Errorf("client from org %s is not authorized to read or write private data from an org %s peer", clientOrgID, peerOrgID)
	}

	return nil
}

/**
富查询 必须是CouchDB才行
*/
func (t *FinancialGeneralAccountChaincode) GetAllFinancialGenerals(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*QueryResult, error) {
	// range query with empty string for startKey and endKey does an open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange(COLLECTION_FINANCIAL_GENERAL_ACCOUNT, startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

/**
查询历史数据
*/
func (t *FinancialGeneralAccountChaincode) GetHistoryForMarble(ctx contractapi.TransactionContextInterface, cardNo string) ([]QueryResult, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(cardNo)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	records := []QueryResult{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		asset := new(FinancialOrgGeneralAccountPrivateData)
		err = json.Unmarshal(response.Value, asset)
		if err != nil {
			return nil, err
		}

		record := QueryResult{
			TxId:      response.TxId,
			Timestamp: time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)),
			Record:    asset,
		}
		records = append(records, record)
	}

	return records, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(FinancialGeneralAccountChaincode))
	if err != nil {
		fmt.Printf("Error create FinancialGeneralAccountChaincode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting FinancialGeneralAccountChaincode chaincode: %s", err.Error())
	}
}

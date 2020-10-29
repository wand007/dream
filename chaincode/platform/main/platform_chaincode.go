package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PlatformChainCode struct {
	contractapi.Contract
}

const CHANNEL_NAME string = "mychannel"
const CHAINCODE_NAME_ISSUE_ORG string = "issueCC"

/**
   平台机构属性
 */
type PlatformOrg struct {
	ID   string `json:"id"`   //平台机构ID
	Name string `json:"name"` //平台机构名称
}

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
 派发记录属性
 */
type DistributionRecordPrivateData struct {
	ID              string  `json:"id"`            //派发记录ID
	IndividualID    string  `json:"individualID"`  //个体ID Individual.ID
	MerchantOrgID   string  `json:"merchantOrgID"` //商户机构ID MerchantOrg.ID
	AgencyOrgID     string  `json:"merchantOrgID"` //代理商机构ID AgencyOrg.ID
	IssueOrgID      string  `json:"issueOrgID"`    //下发机构ID IssueOrg.ID
	Amount          int     `json:"amount"`        //派发金额
	Rate            float64 `json:"rate"`          //派发费率
	ManagedCardNo   string  `json:"managedCardNo"` //金融机构公管账户账号 FinancialOrgManagedAccountPrivateData.CardNo
	ManagedCardCode string  `json:"managedCode"`   //金融机构代码 FinancialOrg.Code
	GeneralCardNo   string  `json:"generalCardNo"` //金融机构公管账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
	GeneralCardCode string  `json:"generalCode"`   //金融机构代码 FinancialOrg.Code
	Status          int     `json:"status"`        //派发状态(0:未下发/1:已下发)
}

/**
 个体属性私有数据属性
 */
type IndividualPrivateData struct {
	ID              string `json:"id"`              //个体ID Individual.ID
	CertificateNo   string `json:"certificateNo"`   //个体证件号
	CertificateType int    `json:"certificateType"` //个体证件类型 (身份证/港澳台证/护照/军官证)
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
   金融机构私有数据属性
 */
type FinancialOrgPrivateData struct {
	ID                    string `json:"id"`                    //金融机构ID FinancialOrg.ID
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
}

/**
   金融机构公管账户私有数据属性
 */
type FinancialOrgManagedAccountPrivateData struct {
	CardNo                string `json:"cardNo"`                //金融机构公管账户账号(唯一不重复)
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	IssueOrgID            string `json:"issueOrgID"`            //下发机构ID IssueOrg.ID
	MerchantOrgID         string `json:"merchantOrgID"`         //商户机构ID MerchantOrg.ID
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

/**
   金融机构一般账户私有数据属性
 */
type FinancialOrgGeneralAccountPrivateData struct {
	CardNo                string `json:"cardNo"`                //金融机构公管账户账号(唯一不重复)
	FinancialOrgID        string `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	CertificateNo         string `json:"certificateNo"`         //持卡者证件号
	CertificateType       int    `json:"certificateType"`       //持卡者证件类型 (身份证/港澳台证/护照/军官证)
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

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

/**
   下发机构金融账户私有数据属性
 */
type IssueOrgFinancialAccountPrivateData struct {
	ID                    string `json:"id"`                    //下发机构账户ID
	IssueOrgID            string `json:"issueOrgID"`            //下发机构ID IssueOrg.ID
	FinancialCode         string `json:"financialCode"`         //金融机构代码
	FinancialCardNo       string `json:"financialCardNo"`       //金融机构下发机构账号
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构下发机构账户凭证余额
	AccStatus             int    `json:"accStatus"`             //系统账户状态(正常/冻结/禁用)
}

/**
 代理商机构属性
 */
type AgencyOrg struct {
	ID     string `json:"id"`     //代理商机构ID
	Name   string `json:"name"`   //代理商机构名称
	Status int    `json:"status"` //代理商机构状态(启用/禁用)
}

/**
 代理商机构私有数据属性
 */
type AgencyOrgPrivateData struct {
	ID        string  `json:"id"`        //代理商机构ID
	RateBasic float64 `json:"rateBasic"` //代理商机构基础费率
}

/**
   代理商机构金融账户私有数据属性
 */
type AgencyOrgFinancialAccountPrivateData struct {
	ID                    string `json:"id"`                    //商户机构账户ID
	AgencyOrgID           string `json:"merchantOrgID"`         //代理商机构ID AgencyOrg.ID
	FinancialCode         string `json:"financialCode"`         //金融机构代码
	FinancialCardNo       string `json:"financialCardNo"`       //金融机构代理商机构账号
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构代理商机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //系统账户状态(正常/冻结/禁用)
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

/**
   商户机构金融账户私有数据属性
 */
type MerchantOrgFinancialAccountPrivateData struct {
	ID                    string `json:"id"`                    //商户机构账户ID
	MerchantOrgID         string `json:"merchantOrgID"`         //商户机构ID MerchantOrg.ID
	FinancialCode         string `json:"financialCode"`         //金融机构代码
	FinancialCardNo       string `json:"financialCardNo"`       //金融机构商户机构账号
	CurrentBalance        int    `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance int    `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int    `json:"accStatus"`             //系统账户状态(正常/冻结/禁用)
}

func (t *PlatformChainCode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("PlatformChainCode Init")
	platformOrg := PlatformOrg{ID: "P768877118787432448", Name: "上帝监管平台"}

	carAsBytes, _ := json.Marshal(platformOrg)
	err := ctx.GetStub().PutState(platformOrg.ID, carAsBytes)

	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}
	return nil
}

func (t *PlatformChainCode) FindById(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("平台id不能为空")
	}
	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return "", errors.New("平台查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("平台数据不存在，读到的%s对应的数据为空！", id)
	}
	return string(bytes), nil
}

func (t *PlatformChainCode) FindIndividualById(ctx contractapi.TransactionContextInterface, issueOrgId string) (string, error) {
	if len(issueOrgId) == 0 {
		return "", errors.New("平台id不能为空")
	}
	trans := [][]byte{[]byte("findById"), []byte("id"), []byte(issueOrgId)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to findById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return "", fmt.Errorf(errStr)
	}

	fmt.Printf("FindIssueOrgById chaincode successful. Got response %s", string(response.Payload))
	return string(response.Payload), nil

}

func (t *PlatformChainCode) FindIssueOrgById(ctx contractapi.TransactionContextInterface, issueOrgId string) (string, error) {
	if len(issueOrgId) == 0 {
		return "", errors.New("平台id不能为空")
	}
	trans := [][]byte{[]byte("FindById"), []byte("id"), []byte(issueOrgId)}
	response := ctx.GetStub().InvokeChaincode(CHAINCODE_NAME_ISSUE_ORG, trans, CHANNEL_NAME)

	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to findById chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return "", fmt.Errorf(errStr)
	}

	fmt.Printf("FindIssueOrgById chaincode successful. Got response %s", string(response.Payload))
	return string(response.Payload), nil

}

func main() {
	chaincode, err := contractapi.NewChaincode(new(PlatformChainCode))
	if err != nil {
		fmt.Printf("Error create PlatformChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting PlatformChainCode chaincode: %s", err.Error())
	}

}

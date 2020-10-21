package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PlatformChainCode struct {
	contractapi.Contract
}

func main() {
	cc, err := contractapi.NewChaincode(new(PlatformChainCode))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting ABstore chaincode: %s", err)
	}
}

/**

角色：
平台（上帝/监管）机构----org0
金融机构----org1
下发机构----org2
代理商----org3
商户----org4
个体----属于org0平台，商户可见基本信息，没有单独的组织节点

 */

/**
   平台机构属性
 */
type PlatformOrg struct {
	ID   string `json:"id"`   //平台机构ID
	Name string `json:"name"` //平台机构名称
}

/**
   下发机构属性
 */
type IssueOrg struct {
	ID     string `json:"id"`     //下发机构ID
	Name   string `json:"name"`   //下发机构名称
	Status string `json:"status"` //下发机构状态(启用/禁用)
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
	ID                    string  `json:"id"`                    //下发机构账户ID
	IssueOrgID            string  `json:"issueOrgID"`            //下发机构ID IssueOrg.ID
	FinancialCode         string  `json:"financialCode"`         //金融机构代码
	FinancialCardNo       string  `json:"financialCardNo"`       //金融机构下发机构账号
	VoucherCurrentBalance float64 `json:"voucherCurrentBalance"` //金融机构下发机构账户凭证余额
	AccStatus             int     `json:"accStatus"`             //系统账户状态(正常/冻结/禁用)
}

/**
 个体属性
 */
type Individual struct {
	ID            string `json:"id"`         //个体ID
	Name          string `json:"name"`       //个体名称
	PlatformOrgID string `json:"issueOrgID"` //平台机构ID PlatformOrg.ID
	Status        string `json:"status"`     //个体状态(启用/禁用)
}

/**
 个体属性私有数据属性
 */
type IndividualPrivateData struct {
	ID              string `json:"id"`              //个体ID Individual.ID
	CertificateNo   string `json:"certificateNo"`   //个体证件号
	CertificateType string `json:"certificateType"` //个体证件类型 (身份证/港澳台证/护照/军官证)
}

/**
   金融机构属性
 */
type FinancialOrg struct {
	ID     string `json:"id"`     //金融机构ID
	Name   string `json:"name"`   //金融机构名称
	Code   string `json:"code"`   //金融机构代码
	Status string `json:"status"` //金融机构状态(启用/禁用)
}

/**
   金融机构公管账户私有数据属性
 */
type FinancialOrgManagedAccountPrivateData struct {
	ID                    string  `json:"id"`                    //金融机构ID
	CardNo                string  `json:"cardNo"`                //金融机构公管账户账号
	FinancialOrgID        string  `json:"financialOrgID"`        //金融机构ID FinancialOrg.ID
	PlatformOrgID         string  `json:"issueOrgID"`            //平台机构ID PlatformOrg.ID
	MerchantOrgID         string  `json:"merchantOrgID"`         //商户机构ID MerchantOrg.ID
	CurrentBalance        float64 `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance float64 `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	Status                string  `json:"status"`                //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
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
	Status          string  `json:"status"`          //金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
}

/**
 商户机构属性
 */
type MerchantOrg struct {
	ID     string `json:"id"`     //金融机构ID
	Name   string `json:"name"`   //金融机构名称
	Status string `json:"status"` //金融机构状态(启用/禁用)
}

/**
   商户机构金融账户私有数据属性
 */
type MerchantOrgFinancialAccountPrivateData struct {
	ID                    string  `json:"id"`                    //商户机构账户ID
	MerchantOrgID         string  `json:"merchantOrgID"`         //商户机构ID MerchantOrg.ID
	FinancialCode         string  `json:"financialCode"`         //金融机构代码
	FinancialCardNo       string  `json:"financialCardNo"`       //金融机构商户机构账号
	CurrentBalance        float64 `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance float64 `json:"voucherCurrentBalance"` //金融机构商户机构账户凭证(token)余额
	AccStatus             int     `json:"accStatus"`             //系统账户状态(正常/冻结/禁用)
}

/**
 代理商机构属性
 */
type AgencyOrg struct {
	ID     string `json:"id"`     //金融机构ID
	Name   string `json:"name"`   //金融机构名称
	Status string `json:"status"` //金融机构状态(启用/禁用)
}

/**
   代理商机构金融账户私有数据属性
 */
type AgencyOrgFinancialAccountPrivateData struct {
	ID                    string  `json:"id"`                    //商户机构账户ID
	AgencyOrgID           string  `json:"merchantOrgID"`         //商户机构ID AgencyOrg.ID
	FinancialCode         string  `json:"financialCode"`         //金融机构代码
	FinancialCardNo       string  `json:"financialCardNo"`       //金融机构代理商机构账号
	CurrentBalance        float64 `json:"currentBalance"`        //金融机构共管账户余额(现金)
	VoucherCurrentBalance float64 `json:"voucherCurrentBalance"` //金融机构代理商机构账户凭证(token)余额
	AccStatus             int     `json:"accStatus"`             //系统账户状态(正常/冻结/禁用)
}

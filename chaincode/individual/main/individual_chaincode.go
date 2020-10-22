package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type IndividualChainCode struct {
	contractapi.Contract
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
	id := individualTransientInput.ID;
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
	if bytes != nil {
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
	bytes, err = ctx.GetStub().GetPrivateData("collectionPlatform", id)
	if err != nil {
		return "", errors.New("个体私有数据查询失败！")
	}
	if bytes != nil {
		fmt.Println("个体私有数据已存在，读到的" + id + "对应的私有数据不为空！")
		individualPrivateData := &IndividualPrivateData{
			ID:              individualTransientInput.ID,
			CertificateNo:   individualTransientInput.CertificateNo,
			CertificateType: individualTransientInput.CertificateType,
		}
		carAsBytes, _ := json.Marshal(individualPrivateData)
		err = ctx.GetStub().PutPrivateData("collectionPlatform", id, carAsBytes)
		if err != nil {
			return "", errors.New("个体私有数据保存失败" + err.Error())
		}
	}

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
	bytes, err = ctx.GetStub().GetPrivateData("collectionPlatform", id)
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
	err = ctx.GetStub().PutPrivateData("collectionPlatform", id, carAsBytes)
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
	bytes, err := ctx.GetStub().GetPrivateData("collectionPlatform", id)
	if err != nil {
		return "", errors.New("个体私有数据查询失败！")
	}
	if bytes == nil {
		return "", fmt.Errorf("个体私有数据不存在，读到的%s对应的私有数据为空！", id)
	}
	return string(bytes), nil

}

func main() {
	chaincode, err := contractapi.NewChaincode(new(IndividualChainCode))
	if err != nil {
		fmt.Printf("Error create PlatformChainCode chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting PlatformChainCode chaincode: %s", err.Error())
	}
}

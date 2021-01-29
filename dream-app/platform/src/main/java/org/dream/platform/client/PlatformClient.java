package org.dream.platform.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.platform.param.rsp.Individual;
import org.dream.platform.param.rsp.IssueOrg;
import org.dream.platform.param.rsp.PlatformOrg;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.nio.charset.StandardCharsets;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 平台机构服务
 */
@Slf4j
@RestController
public class PlatformClient extends GlobalExceptionHandler {

    @Resource(name = "platform-contract")
    ContractImpl platformContract;

    /**
     * 平台机构公开数据查询
     *
     * @param id 平台机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findById"})
    public BusinessResponse findById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = platformContract.evaluateTransaction("FindById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), PlatformOrg.class));
    }

    /**
     * 下发机构公开数据查询
     *
     * @param issueOrgId 下发机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findIssueOrgById"})
    public BusinessResponse findIssueOrgById(@RequestParam(name = "issueOrgId") String issueOrgId) throws ContractException {
        byte[] bytes = platformContract.evaluateTransaction("FindIssueOrgById", issueOrgId);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), IssueOrg.class));
    }

    /**
     * 个体公开数据查询
     *
     * @param individualOrgId 个体ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findIndividualById"})
    public BusinessResponse findIndividualById(@RequestParam(name = "individualOrgId") String individualOrgId) throws ContractException {
        byte[] bytes = platformContract.evaluateTransaction("findIndividualById", individualOrgId);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), Individual.class));
    }


}

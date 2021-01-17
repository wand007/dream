package org.dream.platform.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.platform.param.rsp.Individual;
import org.dream.platform.param.rsp.IndividualPrivateData;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.nio.charset.StandardCharsets;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 个体服务
 */
@Slf4j
@RestController
@RequestMapping("individual")
public class IndividualClient extends GlobalExceptionHandler {
    @Resource
    Network network;
    @Resource(name = "individual-contract")
    Contract individualContract;

    /**
     * 个体公开数据查询
     *
     * @param id 个体ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"/findById"})
    public BusinessResponse findIndividualById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = individualContract.evaluateTransaction("FindById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), Individual.class));
    }

    /**
     * 个体私有数据查询
     *
     * @param id 个体ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"/findPrivateDataById"})
    public BusinessResponse findIndividualPrivateDataById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = individualContract.evaluateTransaction("FindPrivateDataById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), IndividualPrivateData.class));
    }


}

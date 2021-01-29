package org.dream.platform.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.platform.param.rqs.FinancialOrgCreate;
import org.dream.platform.param.rsp.FinancialOrg;
import org.dream.platform.param.rsp.FinancialOrgPrivateData;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
import org.hyperledger.fabric.sdk.Peer;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import javax.validation.Valid;
import java.nio.charset.StandardCharsets;
import java.util.EnumSet;
import java.util.concurrent.TimeoutException;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 金融机构链码客户端
 */
@Slf4j
@RestController
@RequestMapping("financial")
public class FinancialClient extends GlobalExceptionHandler {

    @Resource(name = "financial-contract")
    ContractImpl financialContract;


    /**
     * 金融机构公开数据查询
     *
     * @param id 金融机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findById"})
    public BusinessResponse findById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = financialContract.evaluateTransaction("FindById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), FinancialOrg.class));
    }

    /**
     * 金融机构私有数据查询
     *
     * @param id 金融机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findPrivateDataById"})
    public BusinessResponse FindPrivateDataById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = financialContract.evaluateTransaction("FindPrivateDataById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), FinancialOrgPrivateData.class));
    }

    /**
     * 金融机构新建
     *
     * @param param
     * @return
     * @throws ContractException
     */
    @PostMapping({"create"})
    public BusinessResponse create(@RequestBody @Valid FinancialOrgCreate param) throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = financialContract.createTransaction("Create")
                .setEndorsingPeers(financialContract.getNetwork().getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(param.getId(), param.getName(), param.getCode(), String.valueOf(param.getStatus()));
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
//        TransactionId 不知道在哪里存着的
//        354e879d0fecc640aa50f22f9c17486f63206b6226bf70ac9d76f3295eddbdc9
//        932c2bbfa6ee238370459a7689e42430c3353a72a60f5b4244f332041c7bd94c
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

}

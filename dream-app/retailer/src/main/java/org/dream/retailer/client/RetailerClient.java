package org.dream.retailer.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.retailer.handler.SampleCommitHandlerFactory;
import org.dream.retailer.param.rqs.RetailerOrgCreate;
import org.dream.retailer.param.rsp.RetailerOrg;
import org.dream.retailer.param.rsp.RetailerOrgPrivateData;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Transaction;
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
 * @description 零售商机构服务
 */
@Slf4j
@RestController
public class RetailerClient extends GlobalExceptionHandler {
    @Resource(name = "retailer-contract")
    ContractImpl retailerContract;

    /**
     * 金融机构公开数据查询
     *
     * @param id 金融机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findById"})
    public BusinessResponse findById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = retailerContract.evaluateTransaction("FindById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), RetailerOrg.class));
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
        byte[] bytes = retailerContract.evaluateTransaction("FindPrivateDataById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), RetailerOrgPrivateData.class));
    }

    /**
     * 金融机构新建
     *
     * @param param
     * @return
     * @throws ContractException
     */
    @PostMapping({"create"})
    public BusinessResponse create(@RequestBody @Valid RetailerOrgCreate param)
            throws ContractException, TimeoutException, InterruptedException {
        Transaction transaction = retailerContract.createTransaction("Create")
                .setEndorsingPeers(retailerContract.getNetwork().getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));
        byte[] bytes = transaction.submit(param.getId(), param.getName(), param.getAgencyOrgID(), String.valueOf(param.getUnifiedSocialCreditCode()));
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        //测试事件通知
        SampleCommitHandlerFactory commitHandlerFactory = SampleCommitHandlerFactory.INSTANCE;
        commitHandlerFactory.create(transaction.getTransactionId(), retailerContract.getNetwork());
        transaction.setCommitHandler(commitHandlerFactory);
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }


}

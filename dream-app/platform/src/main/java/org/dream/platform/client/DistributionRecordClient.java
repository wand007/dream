package org.dream.platform.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.platform.param.rqs.DistributionRecordCreate;
import org.dream.platform.param.rsp.DistributionRecordPrivateData;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.sdk.Peer;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import javax.validation.Valid;
import java.nio.charset.StandardCharsets;
import java.util.EnumSet;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeoutException;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 派发记录客户端
 */
@Slf4j
@RestController
public class DistributionRecordClient extends GlobalExceptionHandler {


    @Resource
    Network network;
    @Resource
    Contract contract;


    /**
     * 分销商机构私有数据查询
     *
     * @param id 分销商机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findPrivateDataById"})
    public BusinessResponse FindPrivateDataById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = contract.evaluateTransaction("FindPrivateDataById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), DistributionRecordPrivateData.class));
    }

    /**
     * 分销商机构新建
     *
     * @param param
     * @return
     * @throws ContractException
     */
    @PostMapping({"create"})
    public BusinessResponse create(@RequestBody @Valid DistributionRecordCreate param) throws ContractException, TimeoutException, InterruptedException {
        Map<String, byte[]> transienthMap = new HashMap<String, byte[]>() {
            {
                put("rateBasic", param.getAmount().toPlainString().getBytes());
                put("issueOrgID", param.getId().getBytes());
                put("id", param.getId().getBytes());
            }
        };
        byte[] bytes = contract.createTransaction("Create")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .setTransient(transienthMap)
                .submit();
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }
}

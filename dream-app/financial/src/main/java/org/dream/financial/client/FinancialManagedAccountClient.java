package org.dream.financial.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.financial.param.rqs.FinancialOrgManagedAccountPrivateDataCreate;
import org.dream.financial.param.rsp.FinancialOrgPrivateData;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.sdk.Peer;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import javax.validation.Valid;
import java.math.BigDecimal;
import java.nio.charset.StandardCharsets;
import java.util.EnumSet;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeoutException;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 金融机构链码客户端
 */
@Slf4j
@RestController
@RequestMapping("managed")
public class FinancialManagedAccountClient extends GlobalExceptionHandler {

    @Resource
    Network network;
    @Resource
    Contract contract;


    /**
     * 金融机构私有数据查询
     *
     * @param id 金融机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findPrivateDataById"})
    public BusinessResponse FindPrivateDataById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = contract.evaluateTransaction("FindPrivateDataById", id);
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
    public BusinessResponse create(@RequestBody @Valid FinancialOrgManagedAccountPrivateDataCreate param) throws ContractException, TimeoutException, InterruptedException {
        Map<String, byte[]> transienthMap = new HashMap<String, byte[]>(2) {
            {
                put("managedAccount", JSON.toJSONString(param).getBytes());
            }
        };
        byte[] bytes = contract.createTransaction("Create")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .setTransient(transienthMap)
                .submit();
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));

        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }


    /**
     * 票据交易
     * 零售商向零售商共管账户充值现金时增加票据余额
     *
     * @param managedCardNo
     * @param voucherAmount
     * @return
     * @throws ContractException
     * @throws TimeoutException
     * @throws InterruptedException
     */
    @PostMapping({"transferVoucherAsset"})
    public BusinessResponse transferVoucherAsset(@RequestParam(name = "managedCardNo") String managedCardNo,
                                                 @RequestParam(name = "voucherAmount") BigDecimal voucherAmount)
            throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("TransferVoucherAsset")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(managedCardNo, voucherAmount.toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }
}

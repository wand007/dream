package org.dream.financial.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.financial.param.rqs.FinancialOrgGeneralAccountCreate;
import org.dream.financial.param.rqs.FinancialOrgRealization;
import org.dream.financial.param.rsp.FinancialOrg;
import org.dream.financial.param.rsp.FinancialOrgGeneralAccountPrivateData;
import org.dream.financial.param.rsp.FinancialOrgPrivateData;
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
 * @description 金融机构链码客户端
 */
@Slf4j
@RestController
@RequestMapping("financial")
public class financialGeneralAccountClient extends GlobalExceptionHandler {

    @Resource
    Network network;
    @Resource
    Contract contract;

    /**
     * 金融机构公开数据查询
     *
     * @param id 金融机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findById"})
    public BusinessResponse findById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = contract.evaluateTransaction("FindById", id);
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
    public BusinessResponse create(@RequestBody @Valid FinancialOrgGeneralAccountCreate param) throws ContractException, TimeoutException, InterruptedException {
        Map<String, byte[]> transienthMap = new HashMap<String, byte[]>() {
            {
                put("currentBalance", param.getCurrentBalance().toPlainString().getBytes());
                put("voucherCurrentBalance", param.getVoucherCurrentBalance().toPlainString().getBytes());
                put("cardNo", param.getCardNo().getBytes());
                put("accStatus", String.valueOf(param.getAccStatus()).getBytes());
                put("certificateNo", param.getCertificateNo().getBytes());
                put("financialOrgID", param.getFinancialOrgID().getBytes());
                put("ownerOrg", param.getOwnerOrg().getBytes());
                put("certificateType", String.valueOf(param.getCertificateType()).getBytes());
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
     * 一般账户向共管账户票据兑换现金 (票据变现)
     *
     * @param param
     * @return
     * @throws ContractException
     */
    @PostMapping({"realization"})
    public BusinessResponse realization(@RequestBody @Valid FinancialOrgRealization param) throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("Realization")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(param.getManagedCardNo(), param.getGeneralCardNo(), param.getVoucherAmount().toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

    /**
     * 查询历史数据
     *
     * @param cardNo
     * @return
     * @throws ContractException
     */
    @PostMapping({"getHistoryForMarble"})
    public BusinessResponse getHistoryForMarble(@RequestParam(name = "cardNo") String cardNo) throws ContractException {
        byte[] bytes = contract.evaluateTransaction("GetHistoryForMarble", cardNo);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseArray(new String(bytes, StandardCharsets.UTF_8), FinancialOrgGeneralAccountPrivateData.class));
    }

}

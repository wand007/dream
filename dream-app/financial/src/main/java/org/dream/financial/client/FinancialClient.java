package org.dream.financial.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.financial.param.rqs.FinancialOrgCreate;
import org.dream.financial.param.rqs.FinancialOrgRealization;
import org.dream.financial.param.rsp.FinancialOrg;
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
import java.util.concurrent.TimeoutException;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 金融机构链码客户端
 */
@Slf4j
@RestController
public class FinancialClient extends GlobalExceptionHandler {

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
    public BusinessResponse create(@RequestBody @Valid FinancialOrgCreate param) throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("Create")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(param.getId(), param.getName(), param.getCode(), String.valueOf(param.getStatus()));
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
//        TransactionId 不知道在哪里存着的
//        354e879d0fecc640aa50f22f9c17486f63206b6226bf70ac9d76f3295eddbdc9
//        932c2bbfa6ee238370459a7689e42430c3353a72a60f5b4244f332041c7bd94c
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
    public BusinessResponse Realization(@RequestBody @Valid FinancialOrgRealization param) throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("Realization")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(param.getManagedCardNo(), param.getGeneralCardNo(), param.getVoucherAmount().toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

    /**
     * 一般账户向共管账户现金兑换票据
     *
     * @param id
     * @param amount
     * @return
     * @throws ContractException
     * @throws TimeoutException
     * @throws InterruptedException
     */
    @PostMapping({"grant"})
    public BusinessResponse grant(@RequestParam(name = "id") String id,
                                  @RequestParam(name = "amount") BigDecimal amount) throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("Grant")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(id, amount.toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

    /**
     * 一般账户向共管账户现金兑换票据 (现金充值)
     * 零售商用一般账户的现金余额向上级代理的上级下发机构的金融机构的共管账户充值，获取金融机构颁发的票据，共管账户增加票据余额，零售商减少一般账户的现金余额，增加金融机构的现金余额和票据余额。
     *
     * @param managedCardNo
     * @param generalCardNo
     * @param amount
     * @return
     * @throws ContractException
     * @throws TimeoutException
     * @throws InterruptedException
     */
    @PostMapping({"transferAsset"})
    public BusinessResponse transferAsset(@RequestParam(name = "managedCardNo") String managedCardNo,
                                          @RequestParam(name = "generalCardNo") String generalCardNo,
                                          @RequestParam(name = "amount") BigDecimal amount) throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("TransferAsset")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(managedCardNo, generalCardNo, amount.toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

    /**
     * 共管账户向一般账户交易票据 (票据下发)
     *
     * @param managedCardNo
     * @param generalCardNo
     * @param voucherAmount
     * @return
     * @throws ContractException
     * @throws TimeoutException
     * @throws InterruptedException
     */
    @PostMapping({"transferVoucherAsset"})
    public BusinessResponse transferVoucherAsset(@RequestParam(name = "managedCardNo") String managedCardNo,
                                                 @RequestParam(name = "generalCardNo") String generalCardNo,
                                                 @RequestParam(name = "voucherAmount") BigDecimal voucherAmount) throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("TransferVoucherAsset")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(managedCardNo, generalCardNo, voucherAmount.toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }


}

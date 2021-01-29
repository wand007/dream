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
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
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
@RequestMapping("general")
public class FinancialGeneralAccountClient extends GlobalExceptionHandler {

    @Resource(name = "financial-contract")
    ContractImpl contract;

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
    public BusinessResponse create(@RequestBody @Valid FinancialOrgGeneralAccountCreate param)
            throws ContractException, TimeoutException, InterruptedException {


        Map<String, byte[]> transienthMap = new HashMap<String, byte[]>(2) {
            {
                put("generalAccount", JSON.toJSONString(param).getBytes());
//                put("currentBalance", param.getCurrentBalance().toPlainString().getBytes());
//                put("voucherCurrentBalance", param.getVoucherCurrentBalance().toPlainString().getBytes());
//                put("cardNo", param.getCardNo().getBytes());
//                put("accStatus", String.valueOf(param.getAccStatus()).getBytes());
//                put("certificateNo", param.getCertificateNo().getBytes());
//                put("financialOrgID", param.getFinancialOrgID().getBytes());
//                put("ownerOrg", param.getOwnerOrg().getBytes());
//                put("certificateType", String.valueOf(param.getCertificateType()).getBytes());
            }
        };
        byte[] bytes = contract.createTransaction("Create")
                .setEndorsingPeers(contract.getNetwork().getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
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
    public BusinessResponse realization(@RequestBody @Valid FinancialOrgRealization param)
            throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("Realization")
                .setEndorsingPeers(contract.getNetwork().getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(param.getManagedCardNo(), param.getGeneralCardNo(), param.getVoucherAmount().toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

    /**
     * 现金交易
     * 零售商向零售商一般账户充值现金余额
     *
     * @param generalCardNo
     * @param voucherAmount
     * @return
     * @throws ContractException
     * @throws TimeoutException
     * @throws InterruptedException
     */
    @PostMapping({"transferCashAsset"})
    public BusinessResponse transferCashAsset(@RequestParam(name = "generalCardNo") String generalCardNo,
                                              @RequestParam(name = "voucherAmount") BigDecimal voucherAmount)
            throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("TransferCashAsset")
                .setEndorsingPeers(contract.getNetwork().getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(generalCardNo, voucherAmount.toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

    /**
     * 票据交易
     * 增加个体/零售商/分销商的票据
     *
     * @param generalCardNo
     * @param voucherAmount
     * @return
     * @throws ContractException
     * @throws TimeoutException
     * @throws InterruptedException
     */
    @PostMapping({"transferVoucherAsset"})
    public BusinessResponse transferVoucherAsset(@RequestParam(name = "generalCardNo") String generalCardNo,
                                                 @RequestParam(name = "voucherAmount") BigDecimal voucherAmount)
            throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("TransferVoucherAsset")
                .setEndorsingPeers(contract.getNetwork().getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(generalCardNo, voucherAmount.toPlainString());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

    /**
     * 现金和票据交易 （票据提现）
     * 提现时增加个体/零售商/分销商的现金
     * 提现时减少个体/零售商/分销商的票据
     *
     * @param generalCardNo
     * @param voucherAmount
     * @return
     * @throws ContractException
     * @throws TimeoutException
     * @throws InterruptedException
     */
    @PostMapping({"transferAsset"})
    public BusinessResponse transferAsset(@RequestParam(name = "generalCardNo") String generalCardNo,
                                          @RequestParam(name = "voucherAmount") BigDecimal voucherAmount)
            throws ContractException, TimeoutException, InterruptedException {
        byte[] bytes = contract.createTransaction("TransferAsset")
                .setEndorsingPeers(contract.getNetwork().getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .submit(generalCardNo, voucherAmount.toPlainString());
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
    @GetMapping({"getHistoryForMarble"})
    public BusinessResponse getHistoryForMarble(@RequestParam(name = "financialOrgID") String cardNo)
            throws ContractException {
        byte[] bytes = contract.evaluateTransaction("GetHistoryForMarble", cardNo);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseArray(new String(bytes, StandardCharsets.UTF_8), FinancialOrgGeneralAccountPrivateData.class));
    }

    /**
     * 查询全部
     *
     * @param financialOrgID
     * @param certificateNo
     * @param bookmark
     * @param pageSize
     * @return
     * @throws ContractException
     */
    @GetMapping({"queryFinancialGeneralByOwnerOrgWithPagination"})
    public BusinessResponse queryFinancialGeneralByOwnerOrgWithPagination(@RequestParam(name = "financialOrgID") String financialOrgID,
                                                                          @RequestParam(name = "certificateNo") String certificateNo,
                                                                          @RequestParam(name = "bookmark") String bookmark,
                                                                          @RequestParam(name = "pageSize") Integer pageSize)
            throws ContractException {
        byte[] bytes = contract.evaluateTransaction("QueryFinancialGeneralByOwnerOrgWithPagination",
                financialOrgID, certificateNo, bookmark, String.valueOf(pageSize));
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseArray(new String(bytes, StandardCharsets.UTF_8), FinancialOrgGeneralAccountPrivateData.class));
    }

    /**
     * 根据所属组织机构查询
     *
     * @param financialOrgID
     * @param certificateNo
     * @param bookmark
     * @param pageSize
     * @return
     * @throws ContractException
     */
    @GetMapping({"queryFinancialGeneralWithPagination"})
    public BusinessResponse queryFinancialGeneralWithPagination(@RequestParam(name = "financialOrgID") String financialOrgID,
                                                                @RequestParam(name = "certificateNo") String certificateNo,
                                                                @RequestParam(name = "bookmark") String bookmark,
                                                                @RequestParam(name = "pageSize") Integer pageSize)
            throws ContractException {
        byte[] bytes = contract.evaluateTransaction("QueryFinancialGeneralWithPagination",
                financialOrgID, certificateNo, bookmark, String.valueOf(pageSize));
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseArray(new String(bytes, StandardCharsets.UTF_8), FinancialOrgGeneralAccountPrivateData.class));
    }

}

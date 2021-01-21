package org.dream.platform.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.platform.param.rqs.IndividualCreate;
import org.dream.platform.param.rqs.IndividualUpdate;
import org.dream.platform.param.rsp.Individual;
import org.dream.platform.param.rsp.IndividualPrivateData;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.sdk.Peer;
import org.springframework.util.StringUtils;
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

    /**
     * 个体公开数据分页查询
     *
     * @param name
     * @param platformOrgID
     * @param certificateNo
     * @param certificateType
     * @param bookmark
     * @param pageSize
     * @return
     * @throws ContractException
     */
    @GetMapping({"/queryIndividualSimpleWithPagination"})
    public BusinessResponse queryIndividualSimpleWithPagination(@RequestParam(name = "name", required = false) String name,
                                                                @RequestParam(name = "platformOrgID", required = false) String platformOrgID,
                                                                @RequestParam(name = "certificateNo", required = false) String certificateNo,
                                                                @RequestParam(name = "certificateType", required = false) Integer certificateType,
                                                                @RequestParam(name = "certificateType") String bookmark,
                                                                @RequestParam(name = "pageSize") Integer pageSize) throws ContractException {
        String queryString = "{\"selector\":{\"status\":1}";
        if (!StringUtils.isEmpty(name)) {
            queryString = queryString.concat(",{name:" + name + "}");
        }
        if (!StringUtils.isEmpty(platformOrgID)) {
            queryString = queryString.concat(",{platformOrgID:" + platformOrgID + "}");
        }
        if (!StringUtils.isEmpty(certificateNo)) {
            queryString = queryString.concat(",{certificateNo:" + certificateNo + "}");
        }
        if (!StringUtils.isEmpty(certificateType)) {
            queryString = queryString.concat(",{certificateType:" + certificateType + "}");
        }
        queryString = queryString.concat("}");

        byte[] bytes = individualContract.evaluateTransaction("QueryIndividualSimpleWithPagination", queryString, bookmark, String.valueOf(pageSize));
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), IndividualPrivateData.class));
    }


    /**
     * 个体新建
     *
     * @param param
     * @return
     * @throws ContractException
     */
    @PostMapping({"create"})
    public BusinessResponse create(@RequestBody @Valid IndividualCreate param) throws ContractException, TimeoutException, InterruptedException {
        Map<String, byte[]> transienthMap = new HashMap<String, byte[]>() {
            {
                put("id", param.getId().getBytes());
                put("name", param.getName().getBytes());
                put("certificateNo", param.getCertificateNo().getBytes());
                put("platformOrgID", param.getPlatformOrgID().getBytes());
                put("certificateNo", param.getCertificateNo().getBytes());
                put("certificateType", String.valueOf(param.getCertificateType()).getBytes());
                put("status", String.valueOf(param.getStatus()).getBytes());
            }
        };
        byte[] bytes = individualContract.createTransaction("Create")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .setTransient(transienthMap)
                .submit();
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }

    /**
     * 个体修改
     *
     * @param param
     * @return
     * @throws ContractException
     */
    @PostMapping({"update"})
    public BusinessResponse update(@RequestBody @Valid IndividualUpdate param) throws ContractException, TimeoutException, InterruptedException {
        Map<String, byte[]> transienthMap = new HashMap<String, byte[]>() {
            {
                put("id", param.getId().getBytes());
                if (!StringUtils.isEmpty(param.getName())) {
                    put("name", param.getName().getBytes());
                }
                if (!StringUtils.isEmpty(param.getCertificateNo())) {
                    put("certificateNo", param.getCertificateNo().getBytes());
                }
                if (!StringUtils.isEmpty(param.getPlatformOrgID())) {
                    put("platformOrgID", param.getPlatformOrgID().getBytes());
                }
                if (!StringUtils.isEmpty(param.getCertificateNo())) {
                    put("certificateNo", param.getCertificateNo().getBytes());
                }
                if (!StringUtils.isEmpty(param.getCertificateType())) {
                    put("certificateType", param.getCertificateType().getBytes());
                }
                if (!StringUtils.isEmpty(param.getStatus())) {
                    put("status", param.getStatus().getBytes());
                }
            }
        };
        byte[] bytes = individualContract.createTransaction("Update")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .setTransient(transienthMap)
                .submit();
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }
}

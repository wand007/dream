package org.dream.issue.client;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessResponse;
import org.dream.core.base.GlobalExceptionHandler;
import org.dream.issue.param.rqs.IssueOrgCreate;
import org.dream.issue.param.rsp.IssueOrg;
import org.dream.issue.param.rsp.IssueOrgPrivateData;
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
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeoutException;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 零售商机构服务
 */
@Slf4j
@RestController
public class IssueClient extends GlobalExceptionHandler {

    @Resource(name = "issue-contract")
    ContractImpl issueContract;

    /**
     * 下发机构公开数据查询
     *
     * @param id 下发机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findById"})
    public BusinessResponse findById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = issueContract.evaluateTransaction("FindById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), IssueOrg.class));
    }

    /**
     * 下发机构私有数据查询
     *
     * @param id 下发机构ID
     * @return
     * @throws ContractException
     */
    @GetMapping({"findPrivateDataById"})
    public BusinessResponse FindPrivateDataById(@RequestParam(name = "id") String id) throws ContractException {
        byte[] bytes = issueContract.evaluateTransaction("FindPrivateDataById", id);
        System.out.println("查询结果：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(JSON.parseObject(new String(bytes, StandardCharsets.UTF_8), IssueOrgPrivateData.class));
    }

    /**
     * 下发机构新建
     *
     * @param param
     * @return
     * @throws ContractException
     */
    @PostMapping({"create"})
    public BusinessResponse create(@RequestBody @Valid IssueOrgCreate param)
            throws ContractException, TimeoutException, InterruptedException {

        Map<String, byte[]> transienthMap = new HashMap<String, byte[]>(2) {
            {
                put("issue", JSON.toJSONString(param).getBytes());
            }
        };
        byte[] bytes = issueContract.createTransaction("Create")
                .setEndorsingPeers(issueContract.getNetwork().getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                .setTransient(transienthMap)
                .submit(param.getId(), param.getName());
        System.out.println("返回值：" + new String(bytes, StandardCharsets.UTF_8));
        return BusinessResponse.success(new String(bytes, StandardCharsets.UTF_8));
    }


}

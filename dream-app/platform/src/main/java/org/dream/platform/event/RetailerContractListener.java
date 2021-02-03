package org.dream.platform.event;

import com.google.protobuf.ByteString;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
import org.hyperledger.fabric.sdk.BlockEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.util.List;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 零售商机构服务事件通知消费者
 */
@Slf4j
@Component
public class RetailerContractListener implements ApplicationListener<ContextRefreshedEvent> {

    @Resource(name = "retailer-contract")
    ContractImpl retailerContract;

    @Override
    public void onApplicationEvent(ContextRefreshedEvent contextRefreshedEvent) {
        log.info("Retailer Contract Listener Starting..");
        retailerContract.addContractListener(contractEvent -> {
            log.info("Retailer Contract ChaincodeId-------" + contractEvent.getChaincodeId());
            log.info("Retailer Contract Name-------" + contractEvent.getName());
            log.info("Retailer Contract Payload-------" + contractEvent.getPayload());
            BlockEvent blockEvent = contractEvent.getTransactionEvent().getBlockEvent();
            List<ByteString> dataList = blockEvent.getBlock().getData().getDataList();
            dataList.forEach(data -> {
                log.info("Retailer Contract blockEvent-----" + data.toStringUtf8());
            });
        });
    }
}

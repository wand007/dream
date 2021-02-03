package org.dream.platform.event;

import com.google.protobuf.ByteString;
import com.google.protobuf.InvalidProtocolBufferException;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
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
public class RetailerNetworkListener implements ApplicationListener<ContextRefreshedEvent> {

    @Resource(name = "retailer-contract")
    ContractImpl retailerContract;

    @Override
    public void onApplicationEvent(ContextRefreshedEvent contextRefreshedEvent) {
        log.info("Retailer Network Listener Starting..");
        retailerContract.getNetwork().addBlockListener(blockEvent -> {
            try {
                log.info("Retailer Contract ChaincodeId-------" + blockEvent.getChannelId());
                log.info("Retailer Contract PeerName-------" + blockEvent.getPeer().getName());
                
                List<ByteString> dataList = blockEvent.getBlock().getData().getDataList();
                dataList.forEach(data -> {
                    log.info("Retailer Network blockEvent-----" + data.toStringUtf8());
                });
            } catch (InvalidProtocolBufferException e) {
                e.printStackTrace();
            }
        });
    }
}

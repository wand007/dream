package org.dream.platform.param.rqs;

import lombok.Data;

import java.math.BigDecimal;

/**
 * @author 咚咚锵
 * @date 2021/1/22 8:01
 * @description 派发激励属性
 */
@Data
public class DistributionRecordCreate {
    /**
     * 派发记录ID
     */
    private String id;
    /***
     * 个体ID
     *
     */
    private String individualID;
    /**
     * 共管账户账号
     */
    private String managedAccountCardNo;
    /**
     * 个体一般账户账号
     */
    private String individualCardNo;
    /**
     * 派发金额
     */
    private BigDecimal amount;

    /**
     * 派发费率
     */
    private BigDecimal rate;
}

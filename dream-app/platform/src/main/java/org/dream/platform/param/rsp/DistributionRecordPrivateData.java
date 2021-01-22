package org.dream.platform.param.rsp;

import lombok.Data;

import java.math.BigDecimal;

/**
 * @author 咚咚锵
 * @date 2021/1/22 8:07
 * @description 派发属性私有数据
 */
@Data
public class DistributionRecordPrivateData {
    /**
     * 派发记录ID
     */
    String id;
    /**
     * 平台机构ID
     */
    String platformOrgID;
    /**
     * 金融机构ID
     */
    String financialOrgID;
    /**
     * 个体ID
     */
    String individualID;
    /**
     * 零售商机构ID
     */
    String retailerOrgID;
    /**
     * 分销商机构ID
     */
    String agencyOrgID;
    /**
     * 下发机构ID
     */
    String issueOrgID;
    /**
     * 共管账户账号
     */
    String managedAccountCardNo;
    /**
     * 下发机构一般账户账号
     */
    String issueCardNo;
    /**
     * 个体一般账户账号
     */
    String individualCardNo;
    /**
     * 分销商机构一般账户账号
     */
    String AgencyCardNo;
    /**
     * 金融机构公管账户账号
     */
    String managedCardNo;
    /**
     * 金融机构公管账户账号
     */
    String generalCardNo;
    /**
     * 派发金额
     */
    BigDecimal amount;
    /**
     * 派发费率
     */
    BigDecimal rate;
    /**
     * 派发状态(0:未下发/1:已下发)
     */
    BigDecimal status;
}

package org.dream.platform.param.rqs;

import lombok.Data;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 金融机构链码客户端
 */
@Data
public class FinancialOrgGeneralAccountCreate {
    /**
     * 金融机构公管账户账号(唯一不重复)
     */
    private String cardNo;
    /**
     * 金融机构ID FinancialOrg.ID
     */
    private String financialOrgID;
    /**
     * 持卡者证件号
     */
    private String certificateNo;
    /**
     * 持卡者证件类型 (身份证/港澳台证/护照/军官证/统一社会信用代码)
     */
    private String certificateType;
    /**
     * 金融机构共管账户余额(现金)
     */
    private String currentBalance;
    /**
     * 金融机构零售商机构账户凭证(token)余额
     */
    private String voucherCurrentBalance;
    /**
     * 所有权
     */
    private String ownerOrg;
    /**
     * 金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
     */
    private Integer accStatus;
}

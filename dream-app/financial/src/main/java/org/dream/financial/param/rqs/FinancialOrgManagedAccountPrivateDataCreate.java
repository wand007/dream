package org.dream.financial.param.rqs;

import lombok.Data;
import org.dream.core.enums.AccStatusEnum;

/**
 * @author 咚咚锵
 * @date 2021/1/24 10:34
 * @description 金融机构共管账户私有属性
 */
@Data
public class FinancialOrgManagedAccountPrivateDataCreate {
    /**
     * 金融机构共管账户账号
     */
    String cardNo;
    /**
     * 平台机构ID PlatformOrg.ID
     */
    String platformOrgID;
    /**
     * 金融机构ID FinancialOrg.ID
     */
    String financialOrgID;
    /**
     * 下发机构ID IssueOrg.ID
     */
    String issueOrgID;
    /**
     * 零售商机构ID RetailerOrg.ID
     */
    String retailerOrgID;
    /**
     * 分销商机构ID AgencyOrg.ID
     */
    String agencyOrgID;
    /**
     * 下发机构一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
     */
    String issueCardNo;
    /**
     * 分销商机构一般账户账号 FinancialOrgManagedAccountPrivateData.CardNo
     */
    String agencyCardNo;
    /**
     * 分销商机构一般账户账号 FinancialOrgManagedAccountPrivateData.CardNo
     */
    String managedCardNo;
    /**
     * 零售商机构一般账户账号 FinancialOrgGeneralAccountPrivateData.CardNo
     */
    String generalCardNo;
    /**
     * 金融机构零售商机构账户凭证(token)余额
     */
    Integer voucherCurrentBalance;
    /**
     * 金融机构共管账户状态(正常/冻结/黑名单/禁用/限制)
     *
     * @see AccStatusEnum
     */
    Integer accStatus;
}

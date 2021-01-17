package org.dream.platform.param.rsp;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.math.BigDecimal;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 金融机构私有属性
 */
@Getter
@Setter
@ToString
public class FinancialOrgPrivateData {
    /**
     * 金融机构ID
     */
    private String id;
    /**
     * 金融机构共管账户余额(现金)
     */
    private BigDecimal CurrentBalance;
    /**
     * 金融机构零售商机构账户凭证(token)余额
     */
    private BigDecimal voucherCurrentBalance;


    public FinancialOrgPrivateData() {

    }


}

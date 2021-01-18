package org.dream.retailer.param.rsp;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.math.BigDecimal;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 零售商机构私有属性
 */
@Getter
@Setter
@ToString
public class RetailerOrgPrivateData {
    /**
     * 零售商机构ID
     */
    private String id;
    /**
     * 分销商机构ID
     */
    private String agencyOrgID;
    /**
     * 下发机构基础费率
     */
    private BigDecimal rateBasic;


    public RetailerOrgPrivateData() {

    }


}

package org.dream.agency.param.rsp;

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
public class AgencyPrivateData {
    /**
     * 金融机构ID
     */
    private String id;
    /**
     * 下发机构ID
     */
    private String issueOrgID;
    /**
     * 分销商机构基础费率
     */
    private BigDecimal rateBasic;


    public AgencyPrivateData() {

    }


}

package org.dream.retailer.param.rsp;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 零售商机构公开属性
 */
@Getter
@Setter
@ToString
public class RetailerOrg {
    /**
     * 零售商机构ID
     */
    private String id;
    /**
     * 零售商机构名称
     */
    private String name;
    /**
     * 统一社会信用代码
     */
    private String unifiedSocialCreditCode;
    /**
     * 金融机构状态(启用/禁用)
     */
    private int status;


    public RetailerOrg() {

    }


}

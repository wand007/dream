package org.dream.platform.param.rsp;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 金融机构公开属性
 */
@Getter
@Setter
@ToString
public class FinancialOrg {
    /**
     * 金融机构ID
     */
    private String id;
    /**
     * 金融机构名称
     */
    private String name;
    /**
     * 金融机构代码
     */
    private String code;
    /**
     * 金融机构状态(启用/禁用)
     */
    private Integer status;


    public FinancialOrg() {

    }


}

package org.dream.platform.param.rqs;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 个体公开属性
 */
@Getter
@Setter
@ToString
public class IndividualCreate {

    /**
     * 个体ID
     */
    private String id;
    /**
     * 个体名称
     */
    private String name;
    /**
     * 平台机构ID
     */
    private String platformOrgID;
    /**
     * 个体证件号
     */
    private String certificateNo;
    /**
     * 个体证件类型 (身份证/港澳台证/护照/军官证)
     */
    private Integer certificateType;
    /**
     * 个体状态(启用/禁用)
     */
    private Integer status;
}

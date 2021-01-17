package org.dream.platform.param.rsp;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 个体私有属性
 */
@Getter
@Setter
@ToString
public class IndividualPrivateData {
    /**
     * 个体ID
     */
    private String id;
    /**
     * 个体证件号
     */
    private String certificateNo;
    /**
     * 个体证件类型 (身份证/港澳台证/护照/军官证)
     */
    private String certificateType;

}

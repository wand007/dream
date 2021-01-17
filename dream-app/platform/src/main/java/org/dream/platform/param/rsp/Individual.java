package org.dream.platform.param.rsp;

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
public class Individual {

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
     * 个体状态(启用/禁用)
     */
    private int status;
}

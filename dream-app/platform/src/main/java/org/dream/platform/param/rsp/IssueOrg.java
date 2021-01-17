package org.dream.platform.param.rsp;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 下发机构公开属性
 */
@Getter
@Setter
@ToString
public class IssueOrg {
    /**
     * 下发机构ID
     */
    private String id;
    /**
     * 下发机构名称
     */
    private String name;
    /**
     * 下发机构状态(启用/禁用)
     */
    private int status;

}

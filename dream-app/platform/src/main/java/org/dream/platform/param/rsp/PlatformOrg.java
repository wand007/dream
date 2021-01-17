package org.dream.platform.param.rsp;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 平台机构公开属性
 */
@Getter
@Setter
@ToString
public class PlatformOrg {
    /**
     * 平台机构ID
     */
    private String id;
    /**
     * 平台机构名称
     */
    private String name;


    public PlatformOrg() {

    }


}

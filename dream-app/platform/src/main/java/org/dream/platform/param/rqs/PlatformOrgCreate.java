package org.dream.platform.param.rqs;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 金融机构新建属性
 */
@Getter
@Setter
@ToString
public class PlatformOrgCreate {
    /**
     * 金融机构ID
     */
    @NotBlank(message = "金融机构ID不能为空")
    private String id;
    /**
     * 金融机构名称
     */
    @NotBlank(message = "金融机构名称不能为空")
    private String name;
    /**
     * 金融机构代码
     */
    @NotBlank(message = "金融机构代码不能为空")
    private String code;
    /**
     * 金融机构状态(启用/禁用)
     */
    @NotNull(message = "金融机构状态不能为空")
    private Integer status;


    public PlatformOrgCreate() {

    }


}

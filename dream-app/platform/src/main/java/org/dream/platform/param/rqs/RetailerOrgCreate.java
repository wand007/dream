package org.dream.platform.param.rqs;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 零售商机构机构新建属性
 */
@Getter
@Setter
@ToString
public class RetailerOrgCreate {
    /**
     * 零售商机构机构ID
     */
    @NotBlank(message = "零售商机构机构ID不能为空")
    private String id;
    /**
     * 零售商机构机构名称
     */
    @NotBlank(message = "零售商机构机构名称不能为空")
    private String name;
    /**
     * 零售商机构机构状态(启用/禁用)
     */
    @NotNull(message = "零售商机构机构状态不能为空")
    private int status;


    public RetailerOrgCreate() {

    }


}

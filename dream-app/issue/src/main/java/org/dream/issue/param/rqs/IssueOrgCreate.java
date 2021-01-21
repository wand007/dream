package org.dream.issue.param.rqs;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;
import java.math.BigDecimal;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 下发机构机构新建属性
 */
@Getter
@Setter
@ToString
public class IssueOrgCreate {
    /**
     * 下发机构机构ID
     */
    @NotBlank(message = "下发机构机构ID不能为空")
    private String id;
    /**
     * 下发机构机构名称
     */
    @NotBlank(message = "下发机构机构名称不能为空")
    private String name;
    /**
     * 下发机构基础费率
     */
    @NotBlank(message = "下发机构基础费率不能为空")
    private BigDecimal rateBasic;
    /**
     * 下发机构机构状态(启用/禁用)
     */
    @NotNull(message = "下发机构机构状态不能为空")
    private int status;


    public IssueOrgCreate() {

    }


}

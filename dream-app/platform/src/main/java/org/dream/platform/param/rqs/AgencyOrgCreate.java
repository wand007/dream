package org.dream.platform.param.rqs;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import javax.validation.constraints.DecimalMin;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;
import java.math.BigDecimal;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 分销商机构新建属性
 */
@Getter
@Setter
@ToString
public class AgencyOrgCreate {
    /**
     * 分销商机构ID
     */
    @NotBlank(message = "分销商机构ID不能为空")
    private String id;
    /**
     * 分销商机构名称
     */
    @NotBlank(message = "分销商机构名称不能为空")
    private String name;
    /**
     * 统一社会信用代码
     */
    @NotBlank(message = "统一社会信用代码不能为空")
    private String unifiedSocialCreditCode;
    /**
     * 下发机构ID
     */
    @NotBlank(message = "下发机构ID不能为空")
    private String issueOrgID;
    /**
     * 分销商机构基础费率
     */
    @NotBlank(message = "分销商机构基础费率不能为空")
    @DecimalMin(value = "0", message = "分销商机构基础费率不能小于0")
    private BigDecimal rateBasic;
    /**
     * 分销商机构状态(启用/禁用)
     */
    @NotNull(message = "分销商机构状态不能为空")
    private int status;


    public AgencyOrgCreate() {

    }


}

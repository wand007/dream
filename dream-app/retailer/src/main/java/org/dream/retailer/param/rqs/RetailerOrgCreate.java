package org.dream.retailer.param.rqs;

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
public class RetailerOrgCreate {
    /**
     * 零售商机构ID不
     */
    @NotBlank(message = "零售商机构ID不能为空")
    private String id;
    /**
     * 零售商机构名称
     */
    @NotBlank(message = "零售商机构名称不能为空")
    private String name;
    /**
     * 分销商机构ID
     */
    @NotBlank(message = "分销商机构ID不能为空")
    private String agencyOrgID;
    /**
     * 统一社会信用代码
     */
    @NotBlank(message = "统一社会信用代码不能为空")
    private String unifiedSocialCreditCode;


    public RetailerOrgCreate() {

    }


}

package org.dream.issue.param.rsp;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import javax.validation.constraints.NotBlank;
import java.math.BigDecimal;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 下发机构私有属性
 */
@Getter
@Setter
@ToString
public class IssueOrgPrivateData {
    /**
     * 下发机构ID
     */
    private String id;
    /**
     * 下发机构基础费率
     */
    @NotBlank(message = "下发机构基础费率不能为空")
    private BigDecimal rateBasic;


    public IssueOrgPrivateData() {

    }


}

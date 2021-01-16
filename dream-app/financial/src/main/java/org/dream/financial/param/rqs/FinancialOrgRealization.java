package org.dream.financial.param.rqs;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import javax.validation.constraints.DecimalMin;
import javax.validation.constraints.NotBlank;
import java.math.BigDecimal;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 一般账户向共管账户票据兑换现金 (票据变现)
 * 个体/分销商/下发机构发起提现请求
 * 减少金融机构的现金余额和票据余额，增加个体/分销商/下发机构一般账户的现金余额，减少个体/分销商/下发机构一般账户的票据余额
 */
@Getter
@Setter
@ToString
public class FinancialOrgRealization {
    /**
     * 转出共管账户卡号
     */
    @NotBlank(message = "转出共管账户卡号不能为空")
    private String managedCardNo;
    /**
     *
     * 转入一般账户卡号
     */
    @NotBlank(message = "转入一般账户卡号不能为空")
    private String generalCardNo;
    /**
     * 兑换现金金额
     */
    @NotBlank(message = "兑换现金金额不能为空")
    @DecimalMin(value = "0.01", message = "兑换现金金额最少为0.01元")
    private BigDecimal voucherAmount;


    public FinancialOrgRealization() {

    }


}

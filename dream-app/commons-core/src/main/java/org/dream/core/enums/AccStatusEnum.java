package org.dream.core.enums;

import lombok.Getter;
import lombok.ToString;
import org.dream.core.base.BusinessException;

/**
 * @author 咚咚锵
 * @date 2021/1/24 10:34
 * @description 金融机构共管账户状态(正常 / 冻结 / 黑名单 / 禁用 / 限制)
 */
@ToString
@Getter
public enum AccStatusEnum {
    //
    ACC_STATUS_1(1, "正常"),
    ACC_STATUS_2(2, "冻结"),
    ACC_STATUS_3(3, "黑名单"),
    ACC_STATUS_4(4, "禁用"),
    ACC_STATUS_5(5, "限制"),
    ;

    private int code;
    private String desc;

    AccStatusEnum(int code, String desc) {
        this.code = code;
        this.desc = desc;
    }

    public static AccStatusEnum parse(int val) {
        AccStatusEnum[] values = AccStatusEnum.values();
        for (AccStatusEnum anEnum : values) {
            if (anEnum.getCode() == val) {
                return anEnum;
            }
        }
        throw new BusinessException("没有找到此金融机构共管账户状态:" + val);
    }
}

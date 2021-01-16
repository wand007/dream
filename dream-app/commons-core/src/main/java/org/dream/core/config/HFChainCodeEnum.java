package org.dream.core.config;

import lombok.Getter;
import lombok.ToString;
import org.dream.core.base.BusinessException;

/**
 * @author 咚咚锵
 * @date 2021/1/16 22:56
 * @description Fabric链码地址
 */
@Getter
@ToString
public enum HFChainCodeEnum {
    //
    FINANCIAL_FIND_BY_ID("FindById", "金融机构公开数据查询"),
    FINANCIAL_FIND_PRIVATE_DATA_BY_ID("FindPrivateDataById", "金融机构私有数据查询"),
    FINANCIAL_CREATE("Create", "金融机构新建"),
    ;;
    private String name;
    private String desc;

    HFChainCodeEnum(String name, String desc) {
        this.name = name;
        this.desc = desc;
    }

    public static HFChainCodeEnum parse(String name) {
        HFChainCodeEnum[] values = HFChainCodeEnum.values();
        for (HFChainCodeEnum anEnum : values) {
            if (anEnum.getName().equalsIgnoreCase(name)) {
                return anEnum;
            }
        }
        throw new BusinessException("没有找到此链码方法 : " + name);
    }
}

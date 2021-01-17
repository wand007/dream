package org.dream.core.base;

import lombok.Getter;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 自定义状态码
 */
@Getter
public enum BusinessCode {


    /**
     * 成功状态码标志, 请求成功后必须返回该状态码
     * 可在代码中直接使用 SUCCESS, 如果有特殊需要, 自己定义状态码
     * 自己定义方式: COMMON_SUCCESS_XXX(I, "某某某")
     */
    SUCCESS(10000, "操作成功"),

    /**
     * 此状态码前端直接提示
     */
    ALERT_MESSAGE(12000, "前端直接提示的 指导用户进一步操作的信息"),

    /**
     * 基础失败状态码标志
     */
    ERROR(50000, "操作已受理,请稍后再试"),

    /**
     * 其他异常
     */
    ERROR_SYS_PARAMS(51000, "参数异常"),
    /**
     * Fabric区块链账本系统异常
     */
    ERROR_HF_SYS(52000, "区块链账本系统异常,请稍后再试"),


    ;

    private int code;
    private String desc;


    BusinessCode(int code, String desc) {
        this.code = code;
        this.desc = desc;
    }


    public static BusinessCode valueOf(int code) {
        BusinessCode[] values = BusinessCode.values();

        for (BusinessCode value : values) {
            if (code == value.getCode()) {
                return value;
            }
        }

        throw new BusinessException("没找到此业务代码 : " + code);
    }
}

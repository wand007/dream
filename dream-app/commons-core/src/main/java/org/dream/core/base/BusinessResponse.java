package org.dream.core.base;

import lombok.Getter;
import lombok.ToString;

import java.io.Serializable;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 通用返回对象
 */
@Getter
@ToString
public class BusinessResponse<T> implements Serializable {

    private int statusCode;
    private String statusText;
    private T data;
    private Long currentTimeMillis;

    public BusinessResponse(int statusCode, String statusText, T data, Long currentTimeMillis) {
        this.statusCode = statusCode;
        this.statusText = statusText;
        this.data = data;
        this.currentTimeMillis = currentTimeMillis;
    }


    //###############################  Success & Error  ################################################################


    public static <T> BusinessResponse<T> success() {
        return new BusinessResponse<>(BusinessCode.SUCCESS.getCode(), BusinessCode.SUCCESS.getDesc(), null, System.currentTimeMillis());
    }

    public static <T> BusinessResponse<T> success(T data) {
        return new BusinessResponse<>(BusinessCode.SUCCESS.getCode(), BusinessCode.SUCCESS.getDesc(), data, System.currentTimeMillis());
    }

    public static <T> BusinessResponse<T> error() {
        return new BusinessResponse<>(BusinessCode.ALERT_MESSAGE.getCode(), BusinessCode.ALERT_MESSAGE.getDesc(), null, System.currentTimeMillis());
    }

    public static <T> BusinessResponse<T> error(String errMsg) {
        return new BusinessResponse<>(BusinessCode.ALERT_MESSAGE.getCode(), errMsg, null, System.currentTimeMillis());
    }

    public static BusinessResponse fromBusinessCode(int statusCode, String statusText) {
        return new BusinessResponse(statusCode, statusText, null, System.currentTimeMillis());
    }

    public static BusinessResponse fromBusinessCode(BusinessCode businessCode) {
        return new BusinessResponse(businessCode.getCode(), businessCode.getDesc(), null, System.currentTimeMillis());
    }

    public static BusinessResponse fromBusinessCode(BusinessCode businessCode, String statusText) {
        return new BusinessResponse(businessCode.getCode(), statusText, null, System.currentTimeMillis());
    }


    //###############################  判断  ############################################################################


    public boolean successful() {
        return this.statusCode == BusinessCode.SUCCESS.getCode();
    }


}

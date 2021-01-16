package org.dream.core.base;

import lombok.Getter;
import lombok.ToString;

import java.io.Serializable;
import java.util.HashMap;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 通用返回对象
 */
@Getter
@ToString
public class BusinessResponse implements Serializable {

    private int statusCode;
    private String statusText;
    private Object data = new HashMap<>();
    private Long currentTimeMillis;


    public BusinessResponse() {
        this.statusCode = BusinessCode.SUCCESS.getCode();
        this.statusText = "success";
        this.currentTimeMillis = System.currentTimeMillis();
    }

    public BusinessResponse(int statusCode, String statusText, Object data, Long currentTimeMillis) {
        this.statusCode = statusCode;
        this.statusText = statusText;
        this.data = data;
        this.currentTimeMillis = currentTimeMillis;
    }

    public static BusinessResponse fromBusinessCode(int statusCode, String statusText) {
        return new BusinessResponse(statusCode, statusText, "{}", System.currentTimeMillis());
    }

    public static BusinessResponse fromBusinessCode(BusinessCode businessCode) {
        return new BusinessResponse(businessCode.getCode(), businessCode.getDesc(), "{}", System.currentTimeMillis());
    }

    public static BusinessResponse fromBusinessCode(BusinessCode businessCode, String statusText) {
        return new BusinessResponse(businessCode.getCode(), statusText, "{}", System.currentTimeMillis());
    }

    public static BusinessResponse success(Object data) {
        return new BusinessResponse(BusinessCode.SUCCESS.getCode(), "success", data, System.currentTimeMillis());
    }

    public static BusinessResponse success() {
        return new BusinessResponse(BusinessCode.SUCCESS.getCode(), "success", "{}", System.currentTimeMillis());
    }

    public static BusinessResponse error() {
        return new BusinessResponse(BusinessCode.ERROR.getCode(), BusinessCode.ERROR.getDesc(), "{}", System.currentTimeMillis());
    }

    public static BusinessResponse error(String statusText) {
        return new BusinessResponse(BusinessCode.ERROR.getCode(), statusText, "{}", System.currentTimeMillis());
    }


}

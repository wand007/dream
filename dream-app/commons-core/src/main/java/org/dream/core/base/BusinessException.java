package org.dream.core.base;

import lombok.Getter;

import java.io.Serializable;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 自定义异常
 */
@Getter
public class BusinessException extends RuntimeException implements Serializable {


    private Integer code;
    private String msg;


    public BusinessException(String msg) {
        super(msg);
        this.code = BusinessCode.ERROR.getCode();
        this.msg = msg;
    }

    public BusinessException(BusinessCode error) {
        super(error.getDesc());
        this.code = error.getCode();
        this.msg = error.getDesc();
    }

    public BusinessException(BusinessCode businessCode, String msg) {
        super(msg);
        this.code = businessCode.getCode();
        this.msg = msg;
    }
}

package org.dream.core.base;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.converter.HttpMessageConversionException;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.stereotype.Component;
import org.springframework.validation.BindException;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.MissingServletRequestParameterException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.validation.ConstraintViolation;
import javax.validation.ConstraintViolationException;
import javax.validation.ValidationException;
import java.util.Optional;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 全局异常
 */
@Slf4j
@Component
public class GlobalExceptionHandler extends BaseClient {

    @ExceptionHandler(Exception.class)
    @ResponseBody
    BusinessResponse exceptionHandler(Exception e) {
        log.error("SystemException[系统异常]", e);
        return BusinessResponse.fromBusinessCode(BusinessCode.ERROR);
    }

    @ExceptionHandler(BusinessException.class)
    @ResponseBody
    BusinessResponse exceptionHandler(BusinessException e) {
        log.error("BusinessException[业务异常] code=[{}] msg=[{}]", e.getCode(), e.getMsg(), e);
        return BusinessResponse.fromBusinessCode(e.getCode(), e.getMsg());
    }

    @ExceptionHandler(BindException.class)
    @ResponseBody
    BusinessResponse exceptionHandler(BindException e) {
        log.error("BindException[绑定异常]", e);
        return BusinessResponse.fromBusinessCode(BusinessCode.ERROR_SYS_PARAMS.getCode(), e.getBindingResult().getAllErrors().get(0).getDefaultMessage());
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    @ResponseBody
    BusinessResponse exceptionHandler(MethodArgumentNotValidException e) {
        log.error("MethodArgumentNotValidException[校验错误]", e);
        return BusinessResponse.fromBusinessCode(BusinessCode.ERROR_SYS_PARAMS.getCode(), e.getBindingResult().getAllErrors().get(0).getDefaultMessage());
    }

    @ExceptionHandler(MissingServletRequestParameterException.class)
    @ResponseBody
    BusinessResponse exceptionHandler(MissingServletRequestParameterException e) {
        log.error("MissingServletRequestParameterException[缺少参数]", e);
        return BusinessResponse.fromBusinessCode(BusinessCode.ERROR_SYS_PARAMS.getCode(), "请求参数 " + e.getParameterName() + " 不能为空");
    }

    @ExceptionHandler(HttpMessageNotReadableException.class)
    @ResponseBody
    BusinessResponse exceptionHandler(HttpMessageNotReadableException e) {
        log.error("HttpMessageNotReadableException[缺少请求体]", e);
        return BusinessResponse.fromBusinessCode(BusinessCode.ERROR_SYS_PARAMS.getCode(), "请求参数体不能为空");
    }

    @ExceptionHandler(HttpMessageConversionException.class)
    @ResponseBody
    BusinessResponse exceptionHandler(HttpMessageConversionException e) {
        log.error("HttpMessageConversionException[参数类型转换错误]", e);
        return BusinessResponse.fromBusinessCode(BusinessCode.ERROR_SYS_PARAMS.getCode(), "参数格式错误");
    }

    @ExceptionHandler
    @ResponseBody
    BusinessResponse exceptionHandler(ValidationException e) {
        log.error("ValidationException[校验错误]", e);
        String message = "校验错误";
        if (e instanceof ConstraintViolationException) {
            ConstraintViolationException exs = (ConstraintViolationException) e;
            Optional<ConstraintViolation<?>> optional = exs.getConstraintViolations().stream().findAny();
            message = optional.isPresent() ? optional.get().getMessage() : message;
        }
        return BusinessResponse.fromBusinessCode(BusinessCode.ERROR_SYS_PARAMS.getCode(), message);
    }

    @ExceptionHandler(IllegalArgumentException.class)
    @ResponseBody
    BusinessResponse exceptionHandler(IllegalArgumentException e) {
        log.error("IllegalArgumentException[校验错误]", e);
        return BusinessResponse.fromBusinessCode(BusinessCode.ERROR_SYS_PARAMS.getCode(), e.getMessage());
    }

}

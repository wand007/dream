package org.dream.platform.gateway;

import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessException;
import org.hyperledger.fabric.gateway.*;
import org.hyperledger.fabric.gateway.impl.GatewayImpl;
import org.springframework.beans.factory.annotation.Configurable;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.DependsOn;

import java.io.IOException;
import java.io.Reader;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.InvalidKeyException;
import java.security.PrivateKey;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;

import static org.dream.core.config.HFConfig.CHANNEL_NAME;

/**
 * @author 咚咚锵
 * @date 2021/1/16 21:11
 * @description 提交处理程序工厂实现，可创建配置了所有事件的示例提交处理程序实例
 */
@Slf4j
@Configurable
public class GatewayHandlerFactory {


}

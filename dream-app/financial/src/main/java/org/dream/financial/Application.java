package org.dream.financial;

import lombok.extern.slf4j.Slf4j;
import org.dream.core.base.BusinessException;
import org.hyperledger.fabric.gateway.*;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
import org.hyperledger.fabric.gateway.impl.GatewayImpl;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.DependsOn;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.scheduling.annotation.EnableScheduling;

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
 * @description 金融机构服务
 */
@Slf4j
@EnableAsync
@EnableScheduling
@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }


    @Bean("financial-gateway")
    public GatewayImpl platformGateway() {
        Path NETWORK_CONFIG_PATH = Paths.get("dream-app/financial/src/main/resources/connection.json");
        Path credentialPath = Paths.get("first-network/crypto-config/org2/admin.org2.example.com/msp");
        try {
            //使用org1中的user1初始化一个网关wallet账户用于连接网络
            Wallet wallet = Wallets.newInMemoryWallet();
            Path certificatePath = credentialPath.resolve(Paths.get("signcerts", "cert.pem"));

            X509Certificate certificate = readX509Certificate(certificatePath);

            Path privateKeyPath = credentialPath.resolve(Paths.get("keystore", "key.pem"));

            PrivateKey privateKey = getPrivateKey(privateKeyPath);

            wallet.put("user", Identities.newX509Identity("Org2MSP", certificate, privateKey));

            //根据connection.json 获取Fabric网络连接对象
            GatewayImpl.Builder builder = (GatewayImpl.Builder) Gateway.createBuilder();

            builder.identity(wallet, "user").networkConfig(NETWORK_CONFIG_PATH);
            //连接网关
            GatewayImpl gateway = builder.connect();
            return gateway;

        } catch (IOException e) {
            log.error("网关初始化文件失败", e);
            throw new BusinessException("网关初始化文件失败");
        } catch (CertificateException e) {
            log.error("网关初始化认证失败", e);
            throw new BusinessException("网关初始化认证失败");
        } catch (InvalidKeyException e) {
            log.error("网关初始化密钥失败", e);
            throw new BusinessException("网关初始化密钥失败");
        }
    }


    /**
     * 金融机构服务合约对象
     *
     * @return
     */
    @Bean("financial-contract")
    @DependsOn("financial-gateway")
    public ContractImpl platformContract(@Qualifier("financial-gateway") GatewayImpl gateway) {
        //获取mychannel通道
        Network network = gateway.getNetwork(CHANNEL_NAME);
        return (ContractImpl) network.getContract("financial");
    }

    /**
     * 金融机构一般账户合约对象
     *
     * @param gateway
     * @return
     */
    @Bean("financial-general-contract")
    @DependsOn("financial-gateway")
    public ContractImpl individualContract(@Qualifier("financial-gateway") GatewayImpl gateway) {
        //获取mychannel通道
        Network network = gateway.getNetwork(CHANNEL_NAME);
        return (ContractImpl) network.getContract("financial_general_account");
    }


    private static X509Certificate readX509Certificate(final Path certificatePath) throws IOException, CertificateException {
        try (Reader certificateReader = Files.newBufferedReader(certificatePath, StandardCharsets.UTF_8)) {
            return Identities.readX509Certificate(certificateReader);
        }
    }

    private static PrivateKey getPrivateKey(final Path privateKeyPath) throws IOException, InvalidKeyException {
        try (Reader privateKeyReader = Files.newBufferedReader(privateKeyPath, StandardCharsets.UTF_8)) {
            return Identities.readPrivateKey(privateKeyReader);
        }
    }

}

package org.dream.financial;

import org.hyperledger.fabric.gateway.*;
import org.hyperledger.fabric.gateway.impl.GatewayImpl;
import org.hyperledger.fabric.sdk.Peer;

import java.io.File;
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
import java.util.EnumSet;

/**
 * @author ; lidongdong
 * @Description
 * @Date 2021-01-09
 */
public class FinancialClientApp {


    public static void main(String[] args) {
        caConfig2();
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


    public static void caConfig2() {

        Path NETWORK_CONFIG_PATH = Paths.get("src/main/resources/connection.json");
        Path credentialPath = Paths.get("/src/main/resources/crypto-config/org1/admin.org1.example.com/msp");


        try {
            //使用org1中的user1初始化一个网关wallet账户用于连接网络
            Wallet wallet = Wallets.newInMemoryWallet();
            Path certificatePath = credentialPath.resolve(Paths.get("signcerts", "cert.pem"));

            X509Certificate certificate = readX509Certificate(certificatePath);

            Path privateKeyPath = credentialPath.resolve(Paths.get("keystore", "key.pem"));

            PrivateKey privateKey = getPrivateKey(privateKeyPath);

            wallet.put("user", Identities.newX509Identity("Org1MSP", certificate, privateKey));

            //根据connection.json 获取Fabric网络连接对象
            GatewayImpl.Builder builder = (GatewayImpl.Builder) Gateway.createBuilder();

            builder.identity(wallet, "user").networkConfig(NETWORK_CONFIG_PATH);
            //连接网关
            Gateway gateway = builder.connect();
            //获取mychannel通道
            Network network = gateway.getNetwork("mychannel");
            //获取合约对象
            Contract contract = network.getContract("financial");
            //查询a的余额
            byte[] queryAResultBefore = contract.evaluateTransaction("FindById", "F766005404604841984");
            System.out.println("交易前：" + new String(queryAResultBefore, StandardCharsets.UTF_8));

            // a转50给b
            byte[] invokeResult = contract.createTransaction("Create")
                    .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                    .submit("736182013215645696","新增金融机构1","3","1");
            System.out.println(new String(invokeResult, StandardCharsets.UTF_8));

            //查询交易结果
            byte[] queryAResultAfter = contract.evaluateTransaction("FindById", "736182013215645696");
            System.out.println("交易后：" + new String(queryAResultAfter, StandardCharsets.UTF_8));

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

package org.dream.platform.client;

import com.google.protobuf.InvalidProtocolBufferException;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
import org.hyperledger.fabric.gateway.impl.GatewayImpl;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.ChaincodeCollectionConfigurationException;
import org.hyperledger.fabric.sdk.exception.ChaincodeEndorsementPolicyParseException;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.io.File;
import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import java.util.function.Predicate;

import static java.lang.String.format;
import static org.dream.core.config.HFConfig.CHANNEL_NAME;

/**
 * @author 咚咚锵
 * @date 2021/1/18 22:05
 * @description 链码操作
 */
@Slf4j
@RestController
public class FabricLifecycle {
    @Resource(name = "platform-contract")
    ContractImpl platformContract;
    @Resource(name = "financial-contract")
    ContractImpl financialContract;
    @Resource(name = "agency-contract")
    ContractImpl agencyContract;
    @Resource(name = "retailer-contract")
    ContractImpl retailerContract;
    @Resource(name = "issue-contract")
    ContractImpl issueContract;


    private static final String DEFAULT_VALDITATION_PLUGIN = "vscc";
    private static final String DEFAULT_ENDORSMENT_PLUGIN = "escc";
    //public static final Path TEST_FIXTURE_PATH = Paths.get("D:/Work/GoLang/");
//    private static final String TEST_FIXTURES_PATH = "D:/Work/GoLang/src/dream/chaincode/financial/config/chaincodeendorsementpolicy.yaml";
//    public static final String TEST_PRIVATE_PATH = "D:/Work/GoLang/src/dream/chaincode/financial/config/PrivateDataIT.yaml";

    public static final Path TEST_FIXTURE_PATH = Paths.get("/Users/lidongdong/Private/Go_WorkSpace");
    private static final Path TEST_FIXTURES_PATH = TEST_FIXTURE_PATH.resolve(Paths.get("src", "/dream/chaincode/financial/config/chaincodeendorsementpolicy.yaml"));
    private static final Path TEST_PRIVATE_PATH = TEST_FIXTURE_PATH.resolve(Paths.get("src", "/dream/chaincode/financial/config/PrivateDataIT.yaml"));

    private static final String CHAIN_CODE_PATH = "/dream/chaincode/financial/main";
    private static final String CHAIN_CODE_VERSION = "1";
    private static final String ORG_1_MSP = "Org1MSP";
    private static final String ORG_2_MSP = "Org2MSP";
    private static final String ORG_3_MSP = "Org3MSP";
    private static final String ORG_4_MSP = "Org4MSP";
    private static final String ORG_5_MSP = "Org5MSP";
    private static final String ORG_6_MSP = "Org6MSP";

    static void out(String format, Object... args) {

        System.err.flush();
        System.out.flush();

        System.out.println(format(format, args));
        System.err.flush();
        System.out.flush();

    }

    @GetMapping({"runFabricLifecycle"})
    public void runFabricLifecycle() throws IOException, InvalidArgumentException, ChaincodeEndorsementPolicyParseException,
            ProposalException, ChaincodeCollectionConfigurationException, InterruptedException, ExecutionException, TimeoutException {

        GatewayImpl platformGateway = platformContract.getNetwork().getGateway();
        HFClient org1Client = platformGateway.getClient();
        Channel org1Channel = platformContract.getNetwork().getChannel();
        Collection<Peer> org1MyPeers = new ArrayList<>();
        for (Peer peer : org1Channel.getPeers()) {
            if ("peer0.org1.example.com".equalsIgnoreCase(peer.getName())) {
                org1MyPeers.add(peer);
            }
        }
        //校验链码安装
        verifyNoInstalledChaincodes(org1Client, org1MyPeers);

        GatewayImpl financialGateway = financialContract.getNetwork().getGateway();
        HFClient org2Client = financialGateway.getClient();
        Channel org2Channel = financialGateway.getNetwork(CHANNEL_NAME).getChannel();
        Collection<Peer> org2MyPeers = new ArrayList<>();
        for (Peer peer : org2Channel.getPeers()) {
            if ("peer0.org2.example.com".equalsIgnoreCase(peer.getName())) {
                org2MyPeers.add(peer);
            }
        }
        //校验链码安装
        verifyNoInstalledChaincodes(org2Client, org2MyPeers);

        GatewayImpl issueGateway = issueContract.getNetwork().getGateway();
        HFClient org3Client = issueGateway.getClient();
        Channel org3Channel = issueGateway.getNetwork(CHANNEL_NAME).getChannel();
        Collection<Peer> org3MyPeers = new ArrayList<>();
        for (Peer peer : org3Channel.getPeers()) {
            if ("peer0.org3.example.com".equalsIgnoreCase(peer.getName())) {
                org3MyPeers.add(peer);
            }
        }
        //校验链码安装
        verifyNoInstalledChaincodes(org3Client, org3MyPeers);

        GatewayImpl agencyGateway = agencyContract.getNetwork().getGateway();
        HFClient org4Client = agencyGateway.getClient();
        Channel org4Channel = agencyGateway.getNetwork(CHANNEL_NAME).getChannel();
        Collection<Peer> org4MyPeers = new ArrayList<>();
        for (Peer peer : org4Channel.getPeers()) {
            if ("peer0.org4.example.com".equalsIgnoreCase(peer.getName())) {
                org4MyPeers.add(peer);
            }
        }
        //校验链码安装
        verifyNoInstalledChaincodes(org4Client, org4MyPeers);

        GatewayImpl retailerGateway = retailerContract.getNetwork().getGateway();
        HFClient org5Client = retailerGateway.getClient();
        Channel org5Channel = retailerGateway.getNetwork(CHANNEL_NAME).getChannel();
        Collection<Peer> org5MyPeers = new ArrayList<>();
        for (Peer peer : org5Channel.getPeers()) {
            if ("peer0.org5.example.com".equalsIgnoreCase(peer.getName())) {
                org5MyPeers.add(peer);
            }
        }
        //校验链码安装
        verifyNoInstalledChaincodes(org5Client, org5MyPeers);


        //    verifyNotInstalledChaincode(org2Client, org2MyPeers, CHAIN_CODE_NAME, CHAIN_CODE_VERSION);

        //////////////
        ////  DO Go with our own endorsement policy

        final String goChaincodeName = "lc_example_cc_go";
        final String chaincodeVersion = "1";

        LifecycleChaincodePackage lifecycleChaincodePackage = createLifecycleChaincodePackage(
                goChaincodeName, // some label
                TransactionRequest.Type.GO_LANG,
                TEST_FIXTURE_PATH.toString(),
                CHAIN_CODE_PATH,
                null);

        //Org1 also creates the endorsement policy for the chaincode. // also known as validationParameter !
        LifecycleChaincodeEndorsementPolicy chaincodeEndorsementPolicy = LifecycleChaincodeEndorsementPolicy.fromSignaturePolicyYamlFile(Paths.get(TEST_FIXTURES_PATH.toString()));


        runChannelBack(org1Client, org1Channel, org1MyPeers,
                lifecycleChaincodePackage, goChaincodeName,
                chaincodeVersion, //Version - bump up next time.
                chaincodeEndorsementPolicy,
                ChaincodeCollectionConfiguration.fromYamlFile(new File(TEST_PRIVATE_PATH.toString())), // ChaincodeCollectionConfiguration
                true  // initRequired
        );
        runInitLedger(org1Client, org1Channel, org1MyPeers,
                org2Client, org2Channel, org2MyPeers,
                lifecycleChaincodePackage, goChaincodeName,
                chaincodeVersion, //Version - bump up next time.
                chaincodeEndorsementPolicy,
                ChaincodeCollectionConfiguration.fromYamlFile(new File(TEST_PRIVATE_PATH.toString())), // ChaincodeCollectionConfiguration
                true  // initRequired
        );

//        runChannelBack(org2Client, org2Channel, org2MyPeers,
//                lifecycleChaincodePackage, goChaincodeName,
//                chaincodeVersion, //Version - bump up next time.
//                chaincodeEndorsementPolicy,
//                ChaincodeCollectionConfiguration.fromYamlFile(new File(TEST_PRIVATE_PATH.toString())), // ChaincodeCollectionConfiguration
//                true,  // initRequired
//                new HashMap<String, Object>() {{
//                    put("sequence", 1L);  // this is an update sequence should be 2
//                    put("FindById", "F766005404604841984");   // init is run which set back to 300.  new chaincoode doubles the move of 10 to 20 so expect 320
//                }});
//
//        runChannelBack(org3Client, org3Channel, org3MyPeers,
//                lifecycleChaincodePackage, goChaincodeName,
//                chaincodeVersion, //Version - bump up next time.
//                chaincodeEndorsementPolicy,
//                ChaincodeCollectionConfiguration.fromYamlFile(new File(TEST_PRIVATE_PATH.toString())), // ChaincodeCollectionConfiguration
//                true,  // initRequired
//                new HashMap<String, Object>() {{
//                    put("sequence", 1L);  // this is an update sequence should be 2
//                    put("FindById", "F766005404604841984");   // init is run which set back to 300.  new chaincoode doubles the move of 10 to 20 so expect 320
//                }});
//
//        runChannelBack(org4Client, org4Channel, org4MyPeers,
//                lifecycleChaincodePackage, goChaincodeName,
//                chaincodeVersion, //Version - bump up next time.
//                chaincodeEndorsementPolicy,
//                ChaincodeCollectionConfiguration.fromYamlFile(new File(TEST_PRIVATE_PATH.toString())), // ChaincodeCollectionConfiguration
//                true,  // initRequired
//                new HashMap<String, Object>() {{
//                    put("sequence", 1L);  // this is an update sequence should be 2
//                    put("FindById", "F766005404604841984");   // init is run which set back to 300.  new chaincoode doubles the move of 10 to 20 so expect 320
//                }});
//
//        runChannelBack(org5Client, org5Channel, org5MyPeers,
//                lifecycleChaincodePackage, goChaincodeName,
//                chaincodeVersion, //Version - bump up next time.
//                chaincodeEndorsementPolicy,
//                ChaincodeCollectionConfiguration.fromYamlFile(new File(TEST_PRIVATE_PATH.toString())), // ChaincodeCollectionConfiguration
//                true,  // initRequired
//                new HashMap<String, Object>() {{
//                    put("sequence", 1L);  // this is an update sequence should be 2
//                    put("FindById", "F766005404604841984");   // init is run which set back to 300.  new chaincoode doubles the move of 10 to 20 so expect 320
//                }});

        //// Do Go update. Use same chaincode name, new version and chaincode package. This chaincode doubles move result so we know it changed.

//
//        LifecycleChaincodePackage lifecycleChaincodePackageUpdate = createLifecycleChaincodePackage(
//                goChaincodeName, // some label
//                TransactionRequest.Type.GO_LANG,
//                TEST_FIXTURE_PATH.toString(),
//                CHAIN_CODE_PATH,
//                null); // no metadata this time.
//
//        runChannel(org1Client, org1Channel, org1MyPeers,
//                org2Client, org2Channel, org2MyPeers,
//                lifecycleChaincodePackageUpdate, goChaincodeName,
//                "2", //version is 2 it's an update.
//                chaincodeEndorsementPolicy,
//                ChaincodeCollectionConfiguration.fromYamlFile(new File(TEST_PRIVATE_PATH.toString())),
//                true,  // initRequired
//                new HashMap<String, Object>() {{
//                    put("sequence", 2L);  // this is an update sequence should be 2
//                    put("queryBvalue", "320");  // init is run which set back to 300.  new chaincoode doubles the move of 10 to 20 so expect 320
//                }});
//
//        //////////////
//        ////  DO Go without any standard init required.
//
//        LifecycleChaincodePackage lifecycleChaincodePackageNoInit = createLifecycleChaincodePackage(
//                "lc_example_cc_go_1", // some label
//                TransactionRequest.Type.GO_LANG,
//                TEST_FIXTURE_PATH.toString(),
//                CHAIN_CODE_PATH,
//                null);
//
//        runChannel(org1Client, org1Channel, org1MyPeers,
//                org2Client, org2Channel, org2MyPeers,
//                lifecycleChaincodePackageNoInit,
//                "lc_example_cc_goNOIT", // chaincode name
//                CHAIN_CODE_VERSION,
//                null, // use default endorsement policy
//                null, // ChaincodeCollectionConfiguration
//                false,  // initRequired is now false
//                new HashMap<String, Object>() {{
//                    put("sequence", 2L);  // this is an update sequence should be 2
//                    put("queryBvalue", "320");  // init is run which set back to 300.  new chaincoode doubles the move of 10 to 20 so expect 320
//                }});


//        org1Channel.shutdown(true); // Force foo channel to shutdown clean up resources.

    }

    //CHECKSTYLE.OFF: ParameterNumber
    void runChannel(HFClient org1Client, Channel org1Channel, Collection<Peer> org1MyPeers,
                    HFClient org2Client, Channel org2Channel, Collection<Peer> org2MyPeers,
                    LifecycleChaincodePackage lifecycleChaincodePackage, String chaincodeName,
                    String chaincodeVersion, LifecycleChaincodeEndorsementPolicy lifecycleChaincodeEndorsementPolicy,
                    ChaincodeCollectionConfiguration chaincodeCollectionConfiguration, boolean initRequired)
            throws IOException, ProposalException, InvalidArgumentException, ExecutionException, InterruptedException,
            TimeoutException, ChaincodeCollectionConfigurationException {


        User org1 = org1Client.getUserContext();
        User org2 = org2Client.getUserContext();
        //Should be no chaincode installed at this time.


        final String chaincodeLabel = lifecycleChaincodePackage.getLabel();
        final TransactionRequest.Type chaincodeType = lifecycleChaincodePackage.getType();

        //Org1 installs the chaincode on its peers.
        out("Org1 installs the chaincode on its peers.");
        String org1ChaincodePackageID = lifecycleInstallChaincode(org1Client, org1MyPeers, lifecycleChaincodePackage);


        //Sanity check to see if chaincode really is on it's peers and has the hash as expected by querying all chaincodes.
        out("Org1 check installed chaincode on peers." + org1ChaincodePackageID);

        verifyByQueryInstalledChaincodes(org1Client, org1MyPeers, chaincodeLabel, org1ChaincodePackageID);
        // another query test if it works
        verifyByQueryInstalledChaincode(org1Client, org1MyPeers, org1ChaincodePackageID, chaincodeLabel);

        // Sequence  number increase with each change and is used to make sure you are referring to the same change.
        long sequence = -1L;
        final QueryLifecycleQueryChaincodeDefinitionRequest queryLifecycleQueryChaincodeDefinitionRequest = org1Client.newQueryLifecycleQueryChaincodeDefinitionRequest();
        queryLifecycleQueryChaincodeDefinitionRequest.setChaincodeName(chaincodeName);

        Collection<LifecycleQueryChaincodeDefinitionProposalResponse> firstQueryDefininitions = org1Channel.lifecycleQueryChaincodeDefinition(queryLifecycleQueryChaincodeDefinitionRequest, org1MyPeers);

        for (LifecycleQueryChaincodeDefinitionProposalResponse firstDefinition : firstQueryDefininitions) {
            if (firstDefinition.getStatus() == ProposalResponse.Status.SUCCESS) {
                sequence = firstDefinition.getSequence() + 1L; //Need to bump it up to the next.
                break;
            } else { //Failed but why?
                if (404 == firstDefinition.getChaincodeActionResponseStatus()) {
                    // not found .. done set sequence to 1;
                    sequence = 1;
                    break;
                }
            }
        }


        //     ChaincodeCollectionConfiguration chaincodeCollectionConfiguration = collectionConfiguration == null ? null : ChaincodeCollectionConfiguration.fromYamlFile(new File(collectionConfiguration));
//            // ChaincodeCollectionConfiguration chaincodeCollectionConfiguration = ChaincodeCollectionConfiguration.fromYamlFile(new File("src/test/fixture/collectionProperties/PrivateDataIT.yaml"));
//            chaincodeCollectionConfiguration = null;
        final Peer anOrg1Peer = org1MyPeers.iterator().next();
        out("Org1 approving chaincode definition for my org.");
        BlockEvent.TransactionEvent transactionEvent = lifecycleApproveChaincodeDefinitionForMyOrg(org1Client, org1Channel,
                Collections.singleton(anOrg1Peer),  //support approve on multiple peers but really today only need one. Go with minimum.
                sequence, chaincodeName, chaincodeVersion, lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1ChaincodePackageID)
                .get(10000, TimeUnit.SECONDS);
        out("Org1 approving chaincode TransactionEvent for my org." + transactionEvent.getBlockEvent().getBlock().getData().toString());

        verifyByCheckCommitReadinessStatus(org1Client, org1Channel, sequence, chaincodeName, chaincodeVersion,
                lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1MyPeers,
                new HashSet<>(Arrays.asList(ORG_1_MSP)), // Approved
                new HashSet<>(Arrays.asList(ORG_2_MSP))); // Un approved.

        //Serialize these to bytes to give to other organizations.
        byte[] chaincodePackageBtyes = lifecycleChaincodePackage.getAsBytes();
        final byte[] chaincodeEndorsementPolicyAsBytes = lifecycleChaincodeEndorsementPolicy == null ? null : lifecycleChaincodeEndorsementPolicy.getSerializedPolicyBytes();

        ///////////////////////////////////
        //org1 communicates to org2 out of bounds (email, floppy, etc) : CHAIN_CODE_NAME, CHAIN_CODE_VERSION, chaincodeHash, definitionSequence, chaincodePackageBtyes and chaincodeEndorsementPolicyAsBytes.
        ////  Now as org2
        LifecycleChaincodePackage org2LifecycleChaincodePackage = LifecycleChaincodePackage.fromBytes(chaincodePackageBtyes);
        LifecycleChaincodeEndorsementPolicy org2ChaincodeEndorsementPolicy = chaincodeEndorsementPolicyAsBytes == null ? null :
                LifecycleChaincodeEndorsementPolicy.fromBytes(chaincodeEndorsementPolicyAsBytes);

        //Org2 installs the chaincode on its peers
        out("Org2 installs the chaincode on its peers.");
        String org2ChaincodePackageID = lifecycleInstallChaincode(org2Client, org2MyPeers, org2LifecycleChaincodePackage);

        //Sanity check to see if chaincode really is on it's peers and has the hash as expected.
        out("Org2 check installed chaincode on peers.");
        verifyByQueryInstalledChaincodes(org2Client, org2MyPeers, chaincodeLabel, org2ChaincodePackageID);
        // check by querying for specific chaincode
        verifyByQueryInstalledChaincode(org2Client, org2MyPeers, org2ChaincodePackageID, chaincodeLabel);

        //Approve the chaincode for the peer's in org2
        out("Org2 approving chaincode definition for my org.");
        BlockEvent.TransactionEvent org2TransactionEvent = lifecycleApproveChaincodeDefinitionForMyOrg(org2Client, org2Channel,
                Collections.singleton(org2MyPeers.iterator().next()),  //support approve on multiple peers but really today only need one. Go with minimum.
                sequence, chaincodeName, chaincodeVersion, org2ChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org2ChaincodePackageID)
                .get(10000, TimeUnit.SECONDS);


        out("Checking on org2's network for approvals");
        verifyByCheckCommitReadinessStatus(org2Client, org2Channel, sequence, chaincodeName, chaincodeVersion, lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org2MyPeers,
                new HashSet<>(Arrays.asList(ORG_1_MSP, ORG_2_MSP)), // Approved
                Collections.emptySet()); // Un approved.

        out("Checking on org1's network for approvals");
        verifyByCheckCommitReadinessStatus(org1Client, org1Channel, sequence, chaincodeName, chaincodeVersion, lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1MyPeers,
                new HashSet<>(Arrays.asList(ORG_1_MSP, ORG_2_MSP)), // Approved
                Collections.emptySet()); // unapproved.

        // Org2 knows org1 has approved already so it does the chaincode definition commit, but this could be done by org1 too.


        // Get collection of one of org2 orgs peers and one from the other.

        Collection<Peer> org2EndorsingPeers = Arrays.asList(org2MyPeers.iterator().next());
        transactionEvent = commitChaincodeDefinitionRequest(org2Client, org2Channel, sequence, chaincodeName, chaincodeVersion, org2ChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org2EndorsingPeers)
                .get(300000, TimeUnit.SECONDS);


        verifyByQueryChaincodeDefinition(org2Client, org2Channel, chaincodeName, org2MyPeers, sequence, initRequired, chaincodeEndorsementPolicyAsBytes, chaincodeCollectionConfiguration);
        verifyByQueryChaincodeDefinition(org1Client, org1Channel, chaincodeName, org1MyPeers, sequence, initRequired, chaincodeEndorsementPolicyAsBytes, chaincodeCollectionConfiguration);

        verifyByQueryChaincodeDefinitions(org2Client, org2Channel, org2MyPeers, chaincodeName);
        verifyByQueryChaincodeDefinitions(org1Client, org1Channel, org1MyPeers, chaincodeName);

        //Now org2 could also do the init for the chaincode but it just informs org2 admin of the commit so it does it.

        transactionEvent = executeChaincode(org1Client, org1, org1Channel, "",
                initRequired ? true : null, // doInit don't even specify it has it should default to false
//                chaincodeName, chaincodeType, "a,", "100", "b", "300").get(10000, TimeUnit.SECONDS);
                chaincodeName, chaincodeType, "").get(300000, TimeUnit.SECONDS);


        transactionEvent = executeChaincode(org2Client, org2, org2Channel, "InitLedger",
                false, // doInit
                chaincodeName, chaincodeType, "").get(300000, TimeUnit.SECONDS);


        /// Upgrading chaincode is really the same processes as the initial install. Any change requires a new sequence.
        /// Upgrading the actual code will need same chaincode name,  new chaincode package and version.
        /// Cases where running init is never needed include updating the endorsement policy, or adding collections.
        // For that no chaincode install is needed. As always a new sequence is needed and the same chaincode name, version and hash would be used
        // in the ApproveChaincodeDefinitionForMyOrg and commitChaincodeDefinition operations.
        // If chaincode has been committed by other organizations, to run own your own organization peers besides installing it
        //  also the ApproveChaincodeDefinitionForMyOrg operation is needed which in this case would use the same sequence number since there is
        // no actual change to the definition.


    }

    //CHECKSTYLE.OFF: ParameterNumber
    void runChannelBack(HFClient org1Client, Channel org1Channel, Collection<Peer> org1MyPeers,
                        LifecycleChaincodePackage lifecycleChaincodePackage, String chaincodeName,
                        String chaincodeVersion, LifecycleChaincodeEndorsementPolicy lifecycleChaincodeEndorsementPolicy,
                        ChaincodeCollectionConfiguration chaincodeCollectionConfiguration, boolean initRequired) throws IOException, ProposalException, InvalidArgumentException, ExecutionException, InterruptedException, TimeoutException, ChaincodeCollectionConfigurationException {


        User org1 = org1Client.getUserContext();
        //Should be no chaincode installed at this time.


        final String chaincodeLabel = lifecycleChaincodePackage.getLabel();
        final TransactionRequest.Type chaincodeType = lifecycleChaincodePackage.getType();

        //Org1 installs the chaincode on its peers.
        out("Org1 installs the chaincode on its peers.");
        String org1ChaincodePackageID = lifecycleInstallChaincode(org1Client, org1MyPeers, lifecycleChaincodePackage);


        //Sanity check to see if chaincode really is on it's peers and has the hash as expected by querying all chaincodes.
        out("Org1 check installed chaincode on peers." + org1ChaincodePackageID);

        verifyByQueryInstalledChaincodes(org1Client, org1MyPeers, chaincodeLabel, org1ChaincodePackageID);
        // another query test if it works
        verifyByQueryInstalledChaincode(org1Client, org1MyPeers, org1ChaincodePackageID, chaincodeLabel);

        // Sequence  number increase with each change and is used to make sure you are referring to the same change.
        long sequence = -1L;
        final QueryLifecycleQueryChaincodeDefinitionRequest queryLifecycleQueryChaincodeDefinitionRequest = org1Client.newQueryLifecycleQueryChaincodeDefinitionRequest();
        queryLifecycleQueryChaincodeDefinitionRequest.setChaincodeName(chaincodeName);

        Collection<LifecycleQueryChaincodeDefinitionProposalResponse> firstQueryDefininitions = org1Channel.lifecycleQueryChaincodeDefinition(queryLifecycleQueryChaincodeDefinitionRequest, org1MyPeers);

        for (LifecycleQueryChaincodeDefinitionProposalResponse firstDefinition : firstQueryDefininitions) {
            if (firstDefinition.getStatus() == ProposalResponse.Status.SUCCESS) {
                sequence = firstDefinition.getSequence() + 1L; //Need to bump it up to the next.
                break;
            } else { //Failed but why?
                if (404 == firstDefinition.getChaincodeActionResponseStatus()) {
                    // not found .. done set sequence to 1;
                    sequence = 1;
                    break;
                }
            }
        }


        //     ChaincodeCollectionConfiguration chaincodeCollectionConfiguration = collectionConfiguration == null ? null : ChaincodeCollectionConfiguration.fromYamlFile(new File(collectionConfiguration));
//            // ChaincodeCollectionConfiguration chaincodeCollectionConfiguration = ChaincodeCollectionConfiguration.fromYamlFile(new File("src/test/fixture/collectionProperties/PrivateDataIT.yaml"));
//            chaincodeCollectionConfiguration = null;
        final Peer anOrg1Peer = org1MyPeers.iterator().next();
        out("Org1 approving chaincode definition for my org.");
        BlockEvent.TransactionEvent transactionEvent = lifecycleApproveChaincodeDefinitionForMyOrg(org1Client, org1Channel,
                Collections.singleton(anOrg1Peer),  //support approve on multiple peers but really today only need one. Go with minimum.
                sequence, chaincodeName, chaincodeVersion, lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1ChaincodePackageID)
                .get(30000, TimeUnit.SECONDS);
        out("Org1 approving chaincode TransactionEvent for my org." + transactionEvent.getBlockEvent().getChannelId());

        verifyByCheckCommitReadinessStatus(org1Client, org1Channel, sequence, chaincodeName, chaincodeVersion,
                lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1MyPeers,
                new HashSet<>(Arrays.asList(ORG_1_MSP)), // Approved
                new HashSet<>(Arrays.asList(ORG_2_MSP))); // Un approved.

        //Serialize these to bytes to give to other organizations.
        byte[] chaincodePackageBtyes = lifecycleChaincodePackage.getAsBytes();
        final byte[] chaincodeEndorsementPolicyAsBytes = lifecycleChaincodeEndorsementPolicy == null ? null : lifecycleChaincodeEndorsementPolicy.getSerializedPolicyBytes();

        ///////////////////////////////////
        //org1 communicates to org2 out of bounds (email, floppy, etc) : CHAIN_CODE_NAME, CHAIN_CODE_VERSION, chaincodeHash, definitionSequence, chaincodePackageBtyes and chaincodeEndorsementPolicyAsBytes.
        ////  Now as org2
        LifecycleChaincodePackage org2LifecycleChaincodePackage = LifecycleChaincodePackage.fromBytes(chaincodePackageBtyes);
        LifecycleChaincodeEndorsementPolicy org2ChaincodeEndorsementPolicy = chaincodeEndorsementPolicyAsBytes == null ? null :
                LifecycleChaincodeEndorsementPolicy.fromBytes(chaincodeEndorsementPolicyAsBytes);

        out("Checking on org1's network for approvals");
        verifyByCheckCommitReadinessStatus(org1Client, org1Channel, sequence, chaincodeName, chaincodeVersion, lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1MyPeers,
                new HashSet<>(Arrays.asList(ORG_1_MSP, ORG_2_MSP)), // Approved
                Collections.emptySet()); // unapproved.

        // Org2 knows org1 has approved already so it does the chaincode definition commit, but this could be done by org1 too.


        // Get collection of one of org2 orgs peers and one from the other.

//        Collection<Peer> org2EndorsingPeers = Arrays.asList(org1MyPeers.iterator().next());
//        transactionEvent = commitChaincodeDefinitionRequest(org1Client, org1Channel, sequence, chaincodeName, chaincodeVersion, org2ChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org2EndorsingPeers)
//                .get(300000, TimeUnit.SECONDS);
//
//        verifyByQueryChaincodeDefinition(org1Client, org1Channel, chaincodeName, org1MyPeers, sequence, initRequired, chaincodeEndorsementPolicyAsBytes, chaincodeCollectionConfiguration);
//
//        verifyByQueryChaincodeDefinitions(org1Client, org1Channel, org1MyPeers, chaincodeName);
//
//        //Now org2 could also do the init for the chaincode but it just informs org2 admin of the commit so it does it.
//
//        transactionEvent = executeChaincode(org1Client, org1, org1Channel, "",
//                initRequired ? true : null, // doInit don't even specify it has it should default to false
////                chaincodeName, chaincodeType, "a,", "100", "b", "300").get(10000, TimeUnit.SECONDS);
//                chaincodeName, chaincodeType, "").get(300000, TimeUnit.SECONDS);


        /// Upgrading chaincode is really the same processes as the initial install. Any change requires a new sequence.
        /// Upgrading the actual code will need same chaincode name,  new chaincode package and version.
        /// Cases where running init is never needed include updating the endorsement policy, or adding collections.
        // For that no chaincode install is needed. As always a new sequence is needed and the same chaincode name, version and hash would be used
        // in the ApproveChaincodeDefinitionForMyOrg and commitChaincodeDefinition operations.
        // If chaincode has been committed by other organizations, to run own your own organization peers besides installing it
        //  also the ApproveChaincodeDefinitionForMyOrg operation is needed which in this case would use the same sequence number since there is
        // no actual change to the definition.


    }

    //CHECKSTYLE.OFF: ParameterNumber
    void runInitLedger(HFClient org1Client, Channel org1Channel, Collection<Peer> org1MyPeers,
                       HFClient org2Client, Channel org2Channel, Collection<Peer> org2MyPeers,
                       LifecycleChaincodePackage lifecycleChaincodePackage, String chaincodeName,
                       String chaincodeVersion, LifecycleChaincodeEndorsementPolicy lifecycleChaincodeEndorsementPolicy,
                       ChaincodeCollectionConfiguration chaincodeCollectionConfiguration, boolean initRequired)
            throws IOException, ProposalException, InvalidArgumentException, ExecutionException, InterruptedException,
            TimeoutException, ChaincodeCollectionConfigurationException {


        User org1 = org1Client.getUserContext();
        User org2 = org2Client.getUserContext();
        //Should be no chaincode installed at this time.


        final String chaincodeLabel = lifecycleChaincodePackage.getLabel();
        final TransactionRequest.Type chaincodeType = lifecycleChaincodePackage.getType();

//        //Org1 installs the chaincode on its peers.
//        out("Org1 installs the chaincode on its peers.");
//        String org1ChaincodePackageID = lifecycleInstallChaincode(org1Client, org1MyPeers, lifecycleChaincodePackage);
//
//
//        //Sanity check to see if chaincode really is on it's peers and has the hash as expected by querying all chaincodes.
//        out("Org1 check installed chaincode on peers." + org1ChaincodePackageID);
//
//        verifyByQueryInstalledChaincodes(org1Client, org1MyPeers, chaincodeLabel, org1ChaincodePackageID);
//        // another query test if it works
//        verifyByQueryInstalledChaincode(org1Client, org1MyPeers, org1ChaincodePackageID, chaincodeLabel);

        // Sequence  number increase with each change and is used to make sure you are referring to the same change.
        long sequence = -1L;
//        final QueryLifecycleQueryChaincodeDefinitionRequest queryLifecycleQueryChaincodeDefinitionRequest = org1Client.newQueryLifecycleQueryChaincodeDefinitionRequest();
//        queryLifecycleQueryChaincodeDefinitionRequest.setChaincodeName(chaincodeName);
//
//        Collection<LifecycleQueryChaincodeDefinitionProposalResponse> firstQueryDefininitions = org1Channel.lifecycleQueryChaincodeDefinition(queryLifecycleQueryChaincodeDefinitionRequest, org1MyPeers);
//
//        for (LifecycleQueryChaincodeDefinitionProposalResponse firstDefinition : firstQueryDefininitions) {
//            if (firstDefinition.getStatus() == ProposalResponse.Status.SUCCESS) {
//                sequence = firstDefinition.getSequence() + 1L; //Need to bump it up to the next.
//                break;
//            } else { //Failed but why?
//                if (404 == firstDefinition.getChaincodeActionResponseStatus()) {
//                    // not found .. done set sequence to 1;
//                    sequence = 1;
//                    break;
//                }
//            }
//        }


        //     ChaincodeCollectionConfiguration chaincodeCollectionConfiguration = collectionConfiguration == null ? null : ChaincodeCollectionConfiguration.fromYamlFile(new File(collectionConfiguration));
//            // ChaincodeCollectionConfiguration chaincodeCollectionConfiguration = ChaincodeCollectionConfiguration.fromYamlFile(new File("src/test/fixture/collectionProperties/PrivateDataIT.yaml"));
//            chaincodeCollectionConfiguration = null;
        final Peer anOrg1Peer = org1MyPeers.iterator().next();
        out("Org1 approving chaincode definition for my org.");
//        BlockEvent.TransactionEvent transactionEvent = lifecycleApproveChaincodeDefinitionForMyOrg(org1Client, org1Channel,
//                Collections.singleton(anOrg1Peer),  //support approve on multiple peers but really today only need one. Go with minimum.
//                sequence, chaincodeName, chaincodeVersion, lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1ChaincodePackageID)
//                .get(10000, TimeUnit.SECONDS);
//        out("Org1 approving chaincode TransactionEvent for my org." + transactionEvent.getBlockEvent().getBlock().getData().toString());

//        verifyByCheckCommitReadinessStatus(org1Client, org1Channel, sequence, chaincodeName, chaincodeVersion,
//                lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1MyPeers,
//                new HashSet<>(Arrays.asList(ORG_1_MSP)), // Approved
//                new HashSet<>(Arrays.asList(ORG_2_MSP))); // Un approved.

        //Serialize these to bytes to give to other organizations.
//        byte[] chaincodePackageBtyes = lifecycleChaincodePackage.getAsBytes();
        final byte[] chaincodeEndorsementPolicyAsBytes = lifecycleChaincodeEndorsementPolicy == null ? null : lifecycleChaincodeEndorsementPolicy.getSerializedPolicyBytes();

        ///////////////////////////////////
        //org1 communicates to org2 out of bounds (email, floppy, etc) : CHAIN_CODE_NAME, CHAIN_CODE_VERSION, chaincodeHash, definitionSequence, chaincodePackageBtyes and chaincodeEndorsementPolicyAsBytes.
        ////  Now as org2
//        LifecycleChaincodePackage org2LifecycleChaincodePackage = LifecycleChaincodePackage.fromBytes(chaincodePackageBtyes);
        LifecycleChaincodeEndorsementPolicy org2ChaincodeEndorsementPolicy = chaincodeEndorsementPolicyAsBytes == null ? null :
                LifecycleChaincodeEndorsementPolicy.fromBytes(chaincodeEndorsementPolicyAsBytes);

//        //Org2 installs the chaincode on its peers
//        out("Org2 installs the chaincode on its peers.");
//        String org2ChaincodePackageID = lifecycleInstallChaincode(org2Client, org2MyPeers, org2LifecycleChaincodePackage);
//
//        //Sanity check to see if chaincode really is on it's peers and has the hash as expected.
//        out("Org2 check installed chaincode on peers.");
//        verifyByQueryInstalledChaincodes(org2Client, org2MyPeers, chaincodeLabel, org2ChaincodePackageID);
//        // check by querying for specific chaincode
//        verifyByQueryInstalledChaincode(org2Client, org2MyPeers, org2ChaincodePackageID, chaincodeLabel);
//
//        //Approve the chaincode for the peer's in org2
//        out("Org2 approving chaincode definition for my org.");
//        BlockEvent.TransactionEvent org2TransactionEvent = lifecycleApproveChaincodeDefinitionForMyOrg(org2Client, org2Channel,
//                Collections.singleton(org2MyPeers.iterator().next()),  //support approve on multiple peers but really today only need one. Go with minimum.
//                sequence, chaincodeName, chaincodeVersion, org2ChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org2ChaincodePackageID)
//                .get(10000, TimeUnit.SECONDS);


//        out("Checking on org2's network for approvals");
//        verifyByCheckCommitReadinessStatus(org2Client, org2Channel, sequence, chaincodeName, chaincodeVersion, lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org2MyPeers,
//                new HashSet<>(Arrays.asList(ORG_1_MSP, ORG_2_MSP)), // Approved
//                Collections.emptySet()); // Un approved.
//
//        out("Checking on org1's network for approvals");
//        verifyByCheckCommitReadinessStatus(org1Client, org1Channel, sequence, chaincodeName, chaincodeVersion, lifecycleChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org1MyPeers,
//                new HashSet<>(Arrays.asList(ORG_1_MSP, ORG_2_MSP)), // Approved
//                Collections.emptySet()); // unapproved.

        // Org2 knows org1 has approved already so it does the chaincode definition commit, but this could be done by org1 too.


        // Get collection of one of org2 orgs peers and one from the other.

        Collection<Peer> org2EndorsingPeers = Arrays.asList(org2MyPeers.iterator().next());
        BlockEvent.TransactionEvent transactionEvent = commitChaincodeDefinitionRequest(org2Client, org2Channel, sequence, chaincodeName, chaincodeVersion, org2ChaincodeEndorsementPolicy, chaincodeCollectionConfiguration, initRequired, org2EndorsingPeers)
                .get(300000, TimeUnit.SECONDS);


        verifyByQueryChaincodeDefinition(org2Client, org2Channel, chaincodeName, org2MyPeers, sequence, initRequired, chaincodeEndorsementPolicyAsBytes, chaincodeCollectionConfiguration);
        verifyByQueryChaincodeDefinition(org1Client, org1Channel, chaincodeName, org1MyPeers, sequence, initRequired, chaincodeEndorsementPolicyAsBytes, chaincodeCollectionConfiguration);

        verifyByQueryChaincodeDefinitions(org2Client, org2Channel, org2MyPeers, chaincodeName);
        verifyByQueryChaincodeDefinitions(org1Client, org1Channel, org1MyPeers, chaincodeName);

        //Now org2 could also do the init for the chaincode but it just informs org2 admin of the commit so it does it.

        transactionEvent = executeChaincode(org1Client, org1, org1Channel, "",
                initRequired ? true : null, // doInit don't even specify it has it should default to false
//                chaincodeName, chaincodeType, "a,", "100", "b", "300").get(10000, TimeUnit.SECONDS);
                chaincodeName, chaincodeType, "").get(300000, TimeUnit.SECONDS);


        transactionEvent = executeChaincode(org2Client, org2, org2Channel, "InitLedger",
                false, // doInit
                chaincodeName, chaincodeType, "").get(300000, TimeUnit.SECONDS);


        /// Upgrading chaincode is really the same processes as the initial install. Any change requires a new sequence.
        /// Upgrading the actual code will need same chaincode name,  new chaincode package and version.
        /// Cases where running init is never needed include updating the endorsement policy, or adding collections.
        // For that no chaincode install is needed. As always a new sequence is needed and the same chaincode name, version and hash would be used
        // in the ApproveChaincodeDefinitionForMyOrg and commitChaincodeDefinition operations.
        // If chaincode has been committed by other organizations, to run own your own organization peers besides installing it
        //  also the ApproveChaincodeDefinitionForMyOrg operation is needed which in this case would use the same sequence number since there is
        // no actual change to the definition.


    }

    /**
     * 智能合约打包
     *
     * @param chaincodeLabel
     * @param chaincodeType
     * @param chaincodeSourceLocation
     * @param chaincodePath
     * @param metadadataSource
     * @return
     * @throws IOException
     * @throws InvalidArgumentException
     */
    private LifecycleChaincodePackage createLifecycleChaincodePackage(String chaincodeLabel, TransactionRequest.Type chaincodeType,
                                                                      String chaincodeSourceLocation, String chaincodePath, String metadadataSource)
            throws IOException, InvalidArgumentException {
        log.info("creating install package %s.", chaincodeLabel);

        Path metadataSourcePath = null;
        if (metadadataSource != null) {
            metadataSourcePath = Paths.get(metadadataSource);
        }
        LifecycleChaincodePackage lifecycleChaincodePackage = LifecycleChaincodePackage.fromSource(chaincodeLabel, Paths.get(chaincodeSourceLocation),
                chaincodeType,
                chaincodePath, metadataSourcePath);
        log.info(chaincodeLabel + lifecycleChaincodePackage.getLabel()); // what we expect ?
        log.info(chaincodeType + "" + lifecycleChaincodePackage.getType());
        log.info(chaincodePath + lifecycleChaincodePackage.getPath());
//        log.info(chaincodePath + new String(lifecycleChaincodePackage.getAsBytes(), StandardCharsets.UTF_8));

        return lifecycleChaincodePackage;
    }

    private String lifecycleInstallChaincode(HFClient client, Collection<Peer> peers, LifecycleChaincodePackage lifecycleChaincodePackage) throws InvalidArgumentException, ProposalException, InvalidProtocolBufferException {

        int numInstallProposal = 0;

        numInstallProposal = numInstallProposal + peers.size();

        LifecycleInstallChaincodeRequest installProposalRequest = client.newLifecycleInstallChaincodeRequest();
        installProposalRequest.setLifecycleChaincodePackage(lifecycleChaincodePackage);
        installProposalRequest.setProposalWaitTime(300000);

        Collection<LifecycleInstallChaincodeProposalResponse> responses = client.sendLifecycleInstallChaincodeRequest(installProposalRequest, peers);
        log.info("" + responses);

        Collection<ProposalResponse> successful = new LinkedList<>();
        Collection<ProposalResponse> failed = new LinkedList<>();
        String packageID = null;
        for (LifecycleInstallChaincodeProposalResponse response : responses) {
            if (response.getStatus() == ProposalResponse.Status.SUCCESS) {
                log.info("Successful install proposal response Txid: %s from peer %s", response.getTransactionID(), response.getPeer().getName());
                successful.add(response);
                if (packageID == null) {
                    packageID = response.getPackageId();
                    log.info("Hashcode came back as null from peer: {} ", packageID);
                } else {
                    log.info("Miss match on what the peers returned back as the packageID" + packageID + response.getPackageId());
                }
            } else {
                failed.add(response);
                log.info("-----------------package失败:{}", response.getMessage());
            }
        }

        //   }
        log.info("Received %d install proposal responses. Successful+verified: %d . Failed: %d", numInstallProposal, successful.size(), failed.size());

        if (failed.size() > 0) {
            ProposalResponse first = failed.iterator().next();
            log.info("Not enough endorsers for install :" + successful.size() + ".  " + first.getMessage());
        }

        log.info(packageID);
//        log.info("" + packageID.isEmpty());

        return packageID;

    }

    CompletableFuture<BlockEvent.TransactionEvent> lifecycleApproveChaincodeDefinitionForMyOrg(HFClient client, Channel channel,
                                                                                               Collection<Peer> peers, long sequence,
                                                                                               String chaincodeName, String chaincodeVersion,
                                                                                               LifecycleChaincodeEndorsementPolicy chaincodeEndorsementPolicy,
                                                                                               ChaincodeCollectionConfiguration chaincodeCollectionConfiguration,
                                                                                               boolean initRequired, String org1ChaincodePackageID) throws InvalidArgumentException, ProposalException {

        LifecycleApproveChaincodeDefinitionForMyOrgRequest lifecycleApproveChaincodeDefinitionForMyOrgRequest = client.newLifecycleApproveChaincodeDefinitionForMyOrgRequest();
        lifecycleApproveChaincodeDefinitionForMyOrgRequest.setSequence(sequence);
        lifecycleApproveChaincodeDefinitionForMyOrgRequest.setChaincodeName(chaincodeName);
        lifecycleApproveChaincodeDefinitionForMyOrgRequest.setChaincodeVersion(chaincodeVersion);
        lifecycleApproveChaincodeDefinitionForMyOrgRequest.setInitRequired(initRequired);

        if (null != chaincodeCollectionConfiguration) {
            lifecycleApproveChaincodeDefinitionForMyOrgRequest.setChaincodeCollectionConfiguration(chaincodeCollectionConfiguration);
        }

        if (null != chaincodeEndorsementPolicy) {
            lifecycleApproveChaincodeDefinitionForMyOrgRequest.setChaincodeEndorsementPolicy(chaincodeEndorsementPolicy);
        }

        lifecycleApproveChaincodeDefinitionForMyOrgRequest.setPackageId(org1ChaincodePackageID);

        Collection<LifecycleApproveChaincodeDefinitionForMyOrgProposalResponse> lifecycleApproveChaincodeDefinitionForMyOrgProposalResponse = channel.sendLifecycleApproveChaincodeDefinitionForMyOrgProposal(lifecycleApproveChaincodeDefinitionForMyOrgRequest,
                peers);

        log.info("{}", peers.size() + lifecycleApproveChaincodeDefinitionForMyOrgProposalResponse.size());
        for (LifecycleApproveChaincodeDefinitionForMyOrgProposalResponse response : lifecycleApproveChaincodeDefinitionForMyOrgProposalResponse) {
            final Peer peer = response.getPeer();

            log.info("failure on {}  message is: {},{}" + response.getMessage(), ChaincodeResponse.Status.SUCCESS, response.getStatus());
            log.info(response.getMessage(), response.isInvalid());
            log.info(format("failure on "), response.isVerified());
        }

        return channel.sendTransaction(lifecycleApproveChaincodeDefinitionForMyOrgProposalResponse);

    }

    private CompletableFuture<BlockEvent.TransactionEvent> commitChaincodeDefinitionRequest(HFClient client, Channel channel, long definitionSequence, String chaincodeName, String chaincodeVersion,
                                                                                            LifecycleChaincodeEndorsementPolicy chaincodeEndorsementPolicy,
                                                                                            ChaincodeCollectionConfiguration chaincodeCollectionConfiguration,
                                                                                            boolean initRequired, Collection<Peer> endorsingPeers) throws ProposalException, InvalidArgumentException, InterruptedException, ExecutionException {
        LifecycleCommitChaincodeDefinitionRequest lifecycleCommitChaincodeDefinitionRequest = client.newLifecycleCommitChaincodeDefinitionRequest();

        lifecycleCommitChaincodeDefinitionRequest.setSequence(definitionSequence);
        lifecycleCommitChaincodeDefinitionRequest.setChaincodeName(chaincodeName);
        lifecycleCommitChaincodeDefinitionRequest.setChaincodeVersion(chaincodeVersion);
        if (null != chaincodeEndorsementPolicy) {
            lifecycleCommitChaincodeDefinitionRequest.setChaincodeEndorsementPolicy(chaincodeEndorsementPolicy);
        }
        if (null != chaincodeCollectionConfiguration) {
            lifecycleCommitChaincodeDefinitionRequest.setChaincodeCollectionConfiguration(chaincodeCollectionConfiguration);
        }
        lifecycleCommitChaincodeDefinitionRequest.setInitRequired(initRequired);

        Collection<LifecycleCommitChaincodeDefinitionProposalResponse> lifecycleCommitChaincodeDefinitionProposalResponses = channel.sendLifecycleCommitChaincodeDefinitionProposal(lifecycleCommitChaincodeDefinitionRequest,
                endorsingPeers);

        for (LifecycleCommitChaincodeDefinitionProposalResponse resp : lifecycleCommitChaincodeDefinitionProposalResponses) {

            final Peer peer = resp.getPeer();
            log.info(format("%s had unexpected status.", peer.toString()), ChaincodeResponse.Status.SUCCESS, resp.getStatus());
            log.info(format("%s not verified.", peer.toString()), resp.isVerified());
        }

        return channel.sendTransaction(lifecycleCommitChaincodeDefinitionProposalResponses);

    }

    // Lifecycle Queries to used to verify code...

    private void verifyByCheckCommitReadinessStatus(HFClient client, Channel channel, long definitionSequence, String chaincodeName,
                                                    String chaincodeVersion, LifecycleChaincodeEndorsementPolicy chaincodeEndorsementPolicy,
                                                    ChaincodeCollectionConfiguration chaincodeCollectionConfiguration, boolean initRequired, Collection<Peer> org1MyPeers,
                                                    Set<String> expectedApproved, Set<String> expectedUnApproved) throws InvalidArgumentException, ProposalException {
        LifecycleCheckCommitReadinessRequest lifecycleCheckCommitReadinessRequest = client.newLifecycleSimulateCommitChaincodeDefinitionRequest();
        lifecycleCheckCommitReadinessRequest.setSequence(definitionSequence);
        lifecycleCheckCommitReadinessRequest.setChaincodeName(chaincodeName);
        lifecycleCheckCommitReadinessRequest.setChaincodeVersion(chaincodeVersion);
        if (null != chaincodeEndorsementPolicy) {
            lifecycleCheckCommitReadinessRequest.setChaincodeEndorsementPolicy(chaincodeEndorsementPolicy);
        }
        if (null != chaincodeCollectionConfiguration) {
            lifecycleCheckCommitReadinessRequest.setChaincodeCollectionConfiguration(chaincodeCollectionConfiguration);
        }
        lifecycleCheckCommitReadinessRequest.setInitRequired(initRequired);

        Collection<LifecycleCheckCommitReadinessProposalResponse> responses = channel.sendLifecycleCheckCommitReadinessRequest(lifecycleCheckCommitReadinessRequest, org1MyPeers);
        for (LifecycleCheckCommitReadinessProposalResponse resp : responses) {
            final Peer peer = resp.getPeer();
            log.info("", ChaincodeResponse.Status.SUCCESS, resp.getStatus());
            log.info(format("Approved orgs failed on %s", peer), expectedApproved, resp.getApprovedOrgs());
            log.info(format("UnApproved orgs failed on %s", peer), expectedUnApproved, resp.getUnApprovedOrgs());
        }
    }

    private void verifyByQueryChaincodeDefinitions(HFClient client, Channel channel, Collection<Peer> peers, String expectChaincodeName) throws InvalidArgumentException, ProposalException {

        final LifecycleQueryChaincodeDefinitionsRequest request = client.newLifecycleQueryChaincodeDefinitionsRequest();

        Collection<LifecycleQueryChaincodeDefinitionsProposalResponse> proposalResponses = channel.lifecycleQueryChaincodeDefinitions(request, peers);
        for (LifecycleQueryChaincodeDefinitionsProposalResponse proposalResponse : proposalResponses) {
            Peer peer = proposalResponse.getPeer();

            log.info("", ChaincodeResponse.Status.SUCCESS, proposalResponse.getStatus());
            Collection<LifecycleQueryChaincodeDefinitionsResult> chaincodeDefinitions = proposalResponse.getLifecycleQueryChaincodeDefinitionsResult();

            Optional<String> matchingName = chaincodeDefinitions.stream()
                    .map(LifecycleQueryChaincodeDefinitionsResult::getName)
                    .filter(Predicate.isEqual(expectChaincodeName))
                    .findAny();
            log.info(format("On peer %s return namespace for chaincode %s", peer, expectChaincodeName), matchingName.isPresent());
        }
    }

    private void verifyByQueryChaincodeDefinition(HFClient client, Channel channel, String chaincodeName, Collection<Peer> peers, long expectedSequence, boolean expectedInitRequired, byte[] expectedValidationParameter,
                                                  ChaincodeCollectionConfiguration expectedChaincodeCollectionConfiguration) throws ProposalException, InvalidArgumentException, ChaincodeCollectionConfigurationException {

        final QueryLifecycleQueryChaincodeDefinitionRequest queryLifecycleQueryChaincodeDefinitionRequest = client.newQueryLifecycleQueryChaincodeDefinitionRequest();
        queryLifecycleQueryChaincodeDefinitionRequest.setChaincodeName(chaincodeName);

        Collection<LifecycleQueryChaincodeDefinitionProposalResponse> queryChaincodeDefinitionProposalResponses = channel.lifecycleQueryChaincodeDefinition(queryLifecycleQueryChaincodeDefinitionRequest, peers);

        log.info("{}", queryChaincodeDefinitionProposalResponses);
        log.info("{},{}", peers.size(), queryChaincodeDefinitionProposalResponses.size());
        for (LifecycleQueryChaincodeDefinitionProposalResponse response : queryChaincodeDefinitionProposalResponses) {
            log.info("{},{}", ChaincodeResponse.Status.SUCCESS, response.getStatus());
            log.info("{},{}", expectedSequence, response.getSequence());
            if (expectedValidationParameter != null) {
                byte[] validationParameter = response.getValidationParameter();
                log.info("{},{}", validationParameter);
                log.info("{},{}", expectedValidationParameter, validationParameter);
            }

            if (null != expectedChaincodeCollectionConfiguration) {
                final ChaincodeCollectionConfiguration chaincodeCollectionConfiguration = response.getChaincodeCollectionConfiguration();
                log.info("{},{}", chaincodeCollectionConfiguration);
                log.info("{},{}", expectedChaincodeCollectionConfiguration.getAsBytes(), chaincodeCollectionConfiguration.getAsBytes());
            }

            ChaincodeCollectionConfiguration collections = response.getChaincodeCollectionConfiguration();
            log.info("{},{}", expectedInitRequired, response.getInitRequired());
            log.info("{},{}", DEFAULT_ENDORSMENT_PLUGIN, response.getEndorsementPlugin());
            log.info("{},{}", DEFAULT_VALDITATION_PLUGIN, response.getValidationPlugin());
        }
    }

    private void verifyByQueryInstalledChaincode(HFClient client, Collection<Peer> peers, String packageId, String expectedLabel) throws ProposalException, InvalidArgumentException {

        final LifecycleQueryInstalledChaincodeRequest lifecycleQueryInstalledChaincodeRequest = client.newLifecycleQueryInstalledChaincodeRequest();
        lifecycleQueryInstalledChaincodeRequest.setPackageID(packageId);

        Collection<LifecycleQueryInstalledChaincodeProposalResponse> responses = client.sendLifecycleQueryInstalledChaincode(lifecycleQueryInstalledChaincodeRequest, peers);
        log.info("{},{}", responses);
        log.info("responses not same as peers", peers.size(), responses.size());

        for (LifecycleQueryInstalledChaincodeProposalResponse response : responses) {
            log.info("{},{}", ChaincodeResponse.Status.SUCCESS, response.getStatus());
            String peerName = response.getPeer().getName();
            log.info(format("Peer %s returned back bad status code", peerName), ChaincodeResponse.Status.SUCCESS, response.getStatus());
            log.info(format("Peer %s returned back different label", peerName), expectedLabel, response.getLabel());

        }
    }

    private void verifyByQueryInstalledChaincodes(HFClient client, Collection<Peer> peers, String excpectedChaincodeLabel, String excpectedPackageId) throws ProposalException, InvalidArgumentException {

        Collection<LifecycleQueryInstalledChaincodesProposalResponse> results = client.sendLifecycleQueryInstalledChaincodes(client.newLifecycleQueryInstalledChaincodesRequest(), peers);
        log.info("{},{}", results);
        log.info("{},{}", peers.size(), results.size());

        for (LifecycleQueryInstalledChaincodesProposalResponse peerResults : results) {
            boolean found = false;
            final String peerName = peerResults.getPeer().getName();

            log.info(format("Peer returned back bad status %s", peerName), peerResults.getStatus(), ChaincodeResponse.Status.SUCCESS);

            for (LifecycleQueryInstalledChaincodesProposalResponse.LifecycleQueryInstalledChaincodesResult lifecycleQueryInstalledChaincodesResult : peerResults.getLifecycleQueryInstalledChaincodesResult()) {

                if (excpectedPackageId.equals(lifecycleQueryInstalledChaincodesResult.getPackageId())) {
                    log.info(format("Peer %s had chaincode lable mismatch", peerName), excpectedChaincodeLabel, lifecycleQueryInstalledChaincodesResult.getLabel());
                    found = true;
                    break;
                }

            }
            log.info(format("Chaincode label %s, packageId %s not found on peer %s ", excpectedChaincodeLabel, excpectedPackageId, peerName), found);

        }
        return;

    }

    private void verifyNoInstalledChaincodes(HFClient client, Collection<Peer> peers) throws ProposalException, InvalidArgumentException {

        Collection<LifecycleQueryInstalledChaincodesProposalResponse> results = client.sendLifecycleQueryInstalledChaincodes(client.newLifecycleQueryInstalledChaincodesRequest(), peers);
        log.info("{},{}", results);
        log.info("{},{}", peers.size(), results.size());

        for (LifecycleQueryInstalledChaincodesProposalResponse result : results) {

            final String peerName = result.getPeer().getName();
            log.info(format("Peer returned back bad status %s", peerName), result.getStatus(), ChaincodeResponse.Status.SUCCESS);
            Collection<LifecycleQueryInstalledChaincodesProposalResponse.LifecycleQueryInstalledChaincodesResult> lifecycleQueryInstalledChaincodesResult = result.getLifecycleQueryInstalledChaincodesResult();
            log.info(format("Peer %s returned back null result.", peerName), lifecycleQueryInstalledChaincodesResult);
            log.info(format("Peer %s returned back result with chaincode installed.", peerName), lifecycleQueryInstalledChaincodesResult.isEmpty());

        }

    }

    // Not new =================

    CompletableFuture<BlockEvent.TransactionEvent> executeChaincode(HFClient client, User userContext, Channel channel, String fcn, Boolean doInit, String chaincodeName, TransactionRequest.Type chaincodeType, String... args) throws InvalidArgumentException, ProposalException {

        final ExecutionException[] executionExceptions = new ExecutionException[1];

        Collection<ProposalResponse> successful = new LinkedList<>();
        Collection<ProposalResponse> failed = new LinkedList<>();

        TransactionProposalRequest transactionProposalRequest = client.newTransactionProposalRequest();
        transactionProposalRequest.setChaincodeName(chaincodeName);
        transactionProposalRequest.setChaincodeLanguage(chaincodeType);
        transactionProposalRequest.setUserContext(userContext);

        transactionProposalRequest.setFcn(fcn);
        transactionProposalRequest.setProposalWaitTime(30000);
        transactionProposalRequest.setArgs(args);
        if (null != doInit) {
            transactionProposalRequest.setInit(doInit);
        }

        //  Collection<ProposalResponse> transactionPropResp = channel.sendTransactionProposalToEndorsers(transactionProposalRequest);
        Collection<ProposalResponse> transactionPropResp = channel.sendTransactionProposal(transactionProposalRequest, channel.getPeers());
        for (ProposalResponse response : transactionPropResp) {
            if (response.getStatus() == ProposalResponse.Status.SUCCESS) {
                log.info("Successful transaction proposal response Txid: %s from peer %s", response.getTransactionID(), response.getPeer().getName());
                successful.add(response);
            } else {
                failed.add(response);
            }
        }

        log.info("Received %d transaction proposal responses. Successful+verified: %d . Failed: %d",
                transactionPropResp.size(), successful.size(), failed.size());
        if (failed.size() > 0) {
            ProposalResponse firstTransactionProposalResponse = failed.iterator().next();
            log.info("Not enough endorsers for executeChaincode(move a,b,100):" + failed.size() + " endorser error: " +
                    firstTransactionProposalResponse.getMessage() +
                    ". Was verified: " + firstTransactionProposalResponse.isVerified());
        }
        log.info("Successfully received transaction proposal responses.");

        //  System.exit(10);

        ////////////////////////////
        // Send Transaction Transaction to orderer
        log.info("Sending chaincode transaction(move a,b,100) to orderer.");
        return channel.sendTransaction(successful);

    }

    void executeVerifyByQuery(HFClient client, Channel channel, String chaincodeName, String expect) throws ProposalException, InvalidArgumentException {
        log.info("Now query chaincode for the value of b.");
        QueryByChaincodeRequest queryByChaincodeRequest = client.newQueryProposalRequest();
        queryByChaincodeRequest.setArgs("F766005404604841984");
        queryByChaincodeRequest.setFcn("FindById");
        queryByChaincodeRequest.setChaincodeName(chaincodeName);

        Collection<ProposalResponse> queryProposals = channel.queryByChaincode(queryByChaincodeRequest, channel.getPeers());
        for (ProposalResponse proposalResponse : queryProposals) {
            if (!proposalResponse.isVerified() || proposalResponse.getStatus() != ProposalResponse.Status.SUCCESS) {
                log.info("Failed query proposal from peer " + proposalResponse.getPeer().getName() + " status: " + proposalResponse.getStatus() +
                        ". Messages: " + proposalResponse.getMessage()
                        + ". Was verified : " + proposalResponse.isVerified());
            } else {
                String payload = proposalResponse.getProposalResponse().getResponse().getPayload().toStringUtf8();
                log.info("Query payload of b from peer %s returned %s", proposalResponse.getPeer().getName(), payload);

            }
        }

    }
}

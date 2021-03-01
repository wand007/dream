package org.dream.platform.client;


/*
 *
 *  Copyright 2016,2017 DTCC, Fujitsu Australia Software Technology, IBM - All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */


import lombok.extern.slf4j.Slf4j;
import org.apache.http.HttpEntity;
import org.apache.http.HttpResponse;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.ByteArrayEntity;
import org.apache.http.entity.ContentType;
import org.apache.http.entity.StringEntity;
import org.apache.http.entity.mime.HttpMultipartMode;
import org.apache.http.entity.mime.MultipartEntityBuilder;
import org.apache.http.util.EntityUtils;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
import org.hyperledger.fabric.gateway.impl.GatewayImpl;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.BlockEvent.TransactionEvent;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.io.IOException;
import java.util.*;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.TimeUnit;

import static java.lang.String.format;
import static org.hyperledger.fabric.sdk.Channel.PeerOptions.createPeerOptions;


/**
 * Update channel scenario
 * See http://hyperledger-fabric.readthedocs.io/en/master/configtxlator.html
 * for details.
 */
@Slf4j
@RestController
public class UpdateChannel {
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


    private static final String ORG_HYPERLEDGER_FABRIC_SDK_TEST_FABRIC_HOST = "ORG_HYPERLEDGER_FABRIC_SDK_TEST_FABRIC_HOST";
    private static final String LOCALHOST = //Change test to reference another host .. easier config for my testing on Windows !
            System.getenv(ORG_HYPERLEDGER_FABRIC_SDK_TEST_FABRIC_HOST) == null ? "localhost" : System.getenv(ORG_HYPERLEDGER_FABRIC_SDK_TEST_FABRIC_HOST);

    private static final String CONFIGTXLATOR_LOCATION = "http://" + LOCALHOST + ":7059";

    private static final String ORIGINAL_BATCH_TIMEOUT = "\"timeout\": \"2s\""; // Batch time out in configtx.yaml
    private static final String UPDATED_BATCH_TIMEOUT = "\"timeout\": \"5s\"";  // What we want to change it to.

    //  private static final String FOO_CHANNEL_NAME = "systemOrdererChannel";
    private static final String FOO_CHANNEL_NAME = "foo";
    private static final String SYSTEM_CHANNEL_NAME = "systemordererchannel";
    private static final String PEER_0_ORG_1_EXAMPLE_COM_7051 = "peer0.org1.example.com:7051";
    private static final String REGX_S_HOST_PEER_0_ORG_1_EXAMPLE_COM = "(?s).*\"host\":[ \t]*\"peer0\\.org1\\.example\\.com\".*";
    private static final String REGX_S_ANCHOR_PEERS = "(?s).*\"*AnchorPeers\":[ \t]*\\{.*";

    // "Consortiums": { "groups": { "SampleConsortium": {

    private static final String REGX_IS_SYSTEM_CHANNEL = "(?s).*\"Consortiums\":[ \\t\\s]*\\{[ \\s\\t]*\"groups\":[ \\t\\s]*\\{[ \\t\\s]*\"SampleConsortium\":[ \\t\\s]*\\{.*";


//    private Collection<SampleOrg> testSampleOrgs;

    //    SampleStore sampleStore;
//    HFClient client;
    //    SampleUser ordererAdmin;
    HttpClient httpclient;
    //    SampleOrg sampleOrg;
//    User baduser;


//    public void checkConfig() throws Exception {
//
//        out("\n\n\nRUNNING: UpdateChannelIT\n");
//
////        log.info(256, Config.getConfig().getSecurityLevel());
//
//        GatewayImpl platformGateway = platformContract.getNetwork().getGateway();
//        HFClient org1Client = platformGateway.getClient();
//        Channel org1Channel = platformContract.getNetwork().getChannel();
//        Collection<Peer> org1MyPeers = new ArrayList<>();
//        for (Peer peer : org1Channel.getPeers()) {
//            if ("peer0.org1.example.com".equalsIgnoreCase(peer.getName())) {
//                org1MyPeers.add(peer);
//            }
//        }
//        User org2 = org1Client.getUserContext();
//        ////////////////////////////
//        // Setup client
//
//        //Create instance of client.
//        client = HFClient.createNewInstance();
//
//        client.setCryptoSuite(CryptoSuite.Factory.getCryptoSuite());
//
//        ////////////////////////////
//        //Set up USERS
//
//        //Persistence is not part of SDK. Sample file store is for demonstration purposes only!
//        //   MUST be replaced with more robust application implementation  (Database, LDAP)
//        File sampleStoreFile = new File(System.getProperty("java.io.tmpdir") + "/HFCSampletest.properties");
//        //    sampleStoreFile.deleteOnExit();
//
//        sampleStore = new SampleStore(sampleStoreFile);
//
//        //SampleUser can be any implementation that implements org.hyperledger.fabric.sdk.User Interface
//
//        ////////////////////////////
//        // get users for all orgs
//
//        for (SampleOrg sampleOrg : testSampleOrgs) {
//
//            final String orgName = sampleOrg.getName();
//            sampleOrg.setPeerAdmin(sampleStore.getMember(orgName + "Admin", orgName));
//        }
//
//        sampleOrg = testConfig.getIntegrationTestsSampleOrg("peerOrg1");
//
//        SampleOrg sampleOrg2 = testConfig.getIntegrationTestsSampleOrg("peerOrg2");
//        baduser = sampleOrg2.getUser("user1");
//
//        final String sampleOrgName = sampleOrg.getName();
//
//        ordererAdmin = sampleStore.getMember(sampleOrgName + "OrderAdmin", sampleOrgName, "OrdererMSP",
//                Util.findFileSk(Paths.get("src/test/fixture/sdkintegration/e2e-2Orgs/" + testConfig.getFabricConfigGenVers() + "/crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp/keystore/").toFile()),
//                Paths.get("src/test/fixture/sdkintegration/e2e-2Orgs/" + testConfig.getFabricConfigGenVers() + "/crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp/signcerts/Admin@example.com-cert.pem").toFile());
//
//        httpclient = HttpClients.createDefault();
//    }createDefaulthttpclient

    @GetMapping({"test01UserChannel"})
    public void test01UserChannel() {

        try {
            GatewayImpl platformGateway = platformContract.getNetwork().getGateway();
            HFClient org1Client = platformGateway.getClient();
            Channel org1Channel = platformContract.getNetwork().getChannel();
            Collection<Peer> org1MyPeers = new ArrayList<>();
            for (Peer peer : org1Channel.getPeers()) {
                if ("peer0.org1.example.com".equalsIgnoreCase(peer.getName())) {
                    org1MyPeers.add(peer);
                }
            }
            User user = org1Client.getUserContext();
            Collection<Orderer> orderers = org1Channel.getOrderers();
            Orderer ordererAdmin = orderers.iterator().next();
            ////////////////////////////
            //Reconstruct and run the channels

            Channel fooChannel = reconstructChannel(false, FOO_CHANNEL_NAME, org1Client, orderers);

            // Getting foo channels current configuration bytes.
            byte[] channelConfigurationBytes = fooChannel.getChannelConfigurationBytes();

            String originalConfigJson = configTxlatorDecode(httpclient, channelConfigurationBytes);

            //responseAsString is JSON but use just string operations for this test.

            if (!originalConfigJson.contains(ORIGINAL_BATCH_TIMEOUT)) {

                log.info(format("Did not find expected batch timeout '%s', in:%s", ORIGINAL_BATCH_TIMEOUT, originalConfigJson));
            }

            byte[] reEncodedOriginalConfig = configTxLatorEncode(httpclient, originalConfigJson); // we need to get this to make sure the compare has encoding in the same way!

            //Now modify the batch timeout
            String updateString = originalConfigJson.replace(ORIGINAL_BATCH_TIMEOUT, UPDATED_BATCH_TIMEOUT);

            byte[] updatedConfigBytes = configTxLatorEncode(httpclient, updateString);

            byte[] updateBytes = getChannelUpdateBytes(fooChannel, reEncodedOriginalConfig, updatedConfigBytes);

            UpdateChannelConfiguration updateChannelConfiguration = new UpdateChannelConfiguration(updateBytes);

            //To change the channel we need to sign with orderer admin certs which crypto gen stores:

            // private key: src/test/fixture/sdkintegration/e2e-2Orgs/channel/crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp/keystore/f1a9a940f57419a18a83a852884790d59b378281347dd3d4a88c2b820a0f70c9_sk
            //certificate:  src/test/fixture/sdkintegration/e2e-2Orgs/channel/crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp/signcerts/Admin@example.com-cert.pem

            //Ok now do actual channel update.
            fooChannel.updateChannelConfiguration(updateChannelConfiguration, org1Client.getUpdateChannelConfigurationSignature(updateChannelConfiguration, user));

            Thread.sleep(3000); // give time for events to happen

            //Let's add some additional verification...

            // client.setUserContext(sampleOrg.getPeerAdmin());

            final byte[] modChannelBytes = fooChannel.getChannelConfigurationBytes();

            originalConfigJson = configTxlatorDecode(httpclient, modChannelBytes);

            if (!originalConfigJson.contains(UPDATED_BATCH_TIMEOUT)) {
                //If it doesn't have the updated time out it log.infoed.
                log.info(format("Did not find updated expected batch timeout '%s', in:%s", UPDATED_BATCH_TIMEOUT, originalConfigJson));
            }

            if (originalConfigJson.contains(ORIGINAL_BATCH_TIMEOUT)) { //Should not have been there anymore!

                log.info(format("Found original batch timeout '%s', when it was not expected in:%s", ORIGINAL_BATCH_TIMEOUT, originalConfigJson));
            }

            log.info("" + eventCountFilteredBlock); // make sure we got blockevent that were tested.updateChannelConfiguration
            log.info("" + eventCountBlock); // make sure we got blockevent that were tested.

            //Should be no anchor peers defined.
            log.info("" + originalConfigJson.matches(REGX_S_HOST_PEER_0_ORG_1_EXAMPLE_COM));
            log.info("" + originalConfigJson.matches(REGX_S_ANCHOR_PEERS));

            // Get config update for adding an anchor peer.
            Channel.AnchorPeersConfigUpdateResult configUpdateAnchorPeers = fooChannel.getConfigUpdateAnchorPeers(fooChannel.getPeers().iterator().next(), user,
                    Arrays.asList(PEER_0_ORG_1_EXAMPLE_COM_7051), null);

            log.info("" + configUpdateAnchorPeers.getUpdateChannelConfiguration());
            log.info("" + configUpdateAnchorPeers.getPeersAdded().contains(PEER_0_ORG_1_EXAMPLE_COM_7051));

            //Now add anchor peer to channel configuration.
            fooChannel.updateChannelConfiguration(configUpdateAnchorPeers.getUpdateChannelConfiguration(),
                    org1Client.getUpdateChannelConfigurationSignature(configUpdateAnchorPeers.getUpdateChannelConfiguration(), user));
            Thread.sleep(3000); // give time for events to happen

            // Getting foo channels current configuration bytes to check with configtxlator
            channelConfigurationBytes = fooChannel.getChannelConfigurationBytes();
            originalConfigJson = configTxlatorDecode(httpclient, channelConfigurationBytes);

            // Check is anchor peer in config block?
            log.info("" + originalConfigJson.matches(REGX_S_HOST_PEER_0_ORG_1_EXAMPLE_COM));
            log.info("" + originalConfigJson.matches(REGX_S_ANCHOR_PEERS));

            //Should see what's there.
            configUpdateAnchorPeers = fooChannel.getConfigUpdateAnchorPeers(fooChannel.getPeers().iterator().next(), user,
                    null, null);

            log.info("" + configUpdateAnchorPeers.getUpdateChannelConfiguration()); // not updating anything.
            log.info("" + configUpdateAnchorPeers.getCurrentPeers().contains(PEER_0_ORG_1_EXAMPLE_COM_7051)); // peer should   be there.
            log.info("" + configUpdateAnchorPeers.getPeersRemoved().isEmpty()); // not removing any
            log.info("" + configUpdateAnchorPeers.getPeersAdded().isEmpty()); // not adding anything.
            log.info("" + configUpdateAnchorPeers.getUpdatedPeers().isEmpty()); // not updating anyting.

            //Now remove the anchor peer -- get the config update block.
            configUpdateAnchorPeers = fooChannel.getConfigUpdateAnchorPeers(fooChannel.getPeers().iterator().next(), user,
                    null, Arrays.asList(PEER_0_ORG_1_EXAMPLE_COM_7051));

            log.info("" + configUpdateAnchorPeers.getUpdateChannelConfiguration());
            log.info("" + configUpdateAnchorPeers.getCurrentPeers().contains(PEER_0_ORG_1_EXAMPLE_COM_7051)); // peer should still be there.
            log.info("" + configUpdateAnchorPeers.getPeersRemoved().contains(PEER_0_ORG_1_EXAMPLE_COM_7051)); // peer to remove.
            log.info("" + configUpdateAnchorPeers.getPeersAdded().isEmpty()); // not adding anything.
            log.info("" + configUpdateAnchorPeers.getUpdatedPeers().isEmpty());  // no peers should be left.

            // Now do the actual update.
            fooChannel.updateChannelConfiguration(configUpdateAnchorPeers.getUpdateChannelConfiguration(),
                    org1Client.getUpdateChannelConfigurationSignature(configUpdateAnchorPeers.getUpdateChannelConfiguration(), user));
            Thread.sleep(3000); // give time for events to happen
            // Getting foo channels current configuration bytes to check with configtxlator.
            channelConfigurationBytes = fooChannel.getChannelConfigurationBytes(user, fooChannel.getPeers().iterator().next());
            originalConfigJson = configTxlatorDecode(httpclient, channelConfigurationBytes);

            log.info("" + originalConfigJson.matches(REGX_S_HOST_PEER_0_ORG_1_EXAMPLE_COM)); // should be gone!
            log.info("" + originalConfigJson.matches(REGX_S_ANCHOR_PEERS)); //ODDLY we still want this even if it's empty!

            //Should see what's there.
            configUpdateAnchorPeers = fooChannel.getConfigUpdateAnchorPeers(fooChannel.getPeers().iterator().next(), user,
                    null, null);

            // processing of queued blocks should be done on a separate thread and processed relatively quickly to avoid queues from becoming full,
            // But we're just testing/demoing here.
            log.info("" + fooChannel.getBlockListenerHandles().size());  // 1 event type block listener and 2 queued type.
            fooChannel.unregisterBlockListener(listenerHandler1);
            fooChannel.unregisterBlockListener(listenerHandler2);
            log.info("" + fooChannel.getBlockListenerHandles().size()); // now there's only one.

            log.info("" + blockingQueue1.size());
            log.info("" + blockingQueue2.size());
            log.info("" + eventQueueCaputure.size());

            Collection<QueuedBlockEvent> drain1 = new ArrayList<>();
            blockingQueue1.drainTo(drain1);
            Collection<? super QueuedBlockEvent> drain2 = new ArrayList<>();
            blockingQueue2.drainTo(drain2);

            Collection<? super BlockEvent> eventQDrain = new ArrayList<>();
            eventQueueCaputure.drainTo(eventQDrain);

            QueuedBlockEvent[] drain1Array = drain1.toArray(new QueuedBlockEvent[drain1.size()]);
            QueuedBlockEvent[] drain2Array = drain2.toArray(new QueuedBlockEvent[drain2.size()]);
            BlockEvent[] drainEventQArray = eventQDrain.toArray(new BlockEvent[eventQDrain.size()]);

            for (int i = drain1Array.length - 1; i > -1; --i) {
                final long blockNumber = drain1Array[i].getBlockEvent().getBlockNumber();
                final String url = drain1Array[i].getBlockEvent().getPeer().getUrl();

                log.info("" + blockNumber, drain2Array[i].getBlockEvent().getBlockNumber());
                log.info("" + url, drain2Array[i].getBlockEvent().getPeer().getUrl());
                log.info("" + blockNumber, drainEventQArray[i].getBlockNumber());
                log.info("" + url, drainEventQArray[i].getPeer().getUrl());
            }

            log.info("" + configUpdateAnchorPeers.getUpdateChannelConfiguration()); // not updating anything.
            log.info("" + configUpdateAnchorPeers.getCurrentPeers().isEmpty()); // peer should be now gone.
            log.info("" + configUpdateAnchorPeers.getPeersRemoved().isEmpty()); // not removing any
            log.info("" + configUpdateAnchorPeers.getPeersAdded().isEmpty()); // not adding anything.
            log.info("" + configUpdateAnchorPeers.getUpdatedPeers().isEmpty());  // no peers should be left

            out("That's all folks!");

        } catch (Exception e) {
            e.printStackTrace();
            log.info(e.getMessage());
        }
    }

    @GetMapping({"test02SystemChannel"})
    public void test02SystemChannel() {

        try {
            GatewayImpl platformGateway = platformContract.getNetwork().getGateway();
            HFClient org1Client = platformGateway.getClient();
            Channel org1Channel = platformContract.getNetwork().getChannel();
            Collection<Peer> org1MyPeers = new ArrayList<>();
            for (Peer peer : org1Channel.getPeers()) {
                if ("peer0.org1.example.com".equalsIgnoreCase(peer.getName())) {
                    org1MyPeers.add(peer);
                }
            }
            User user = org1Client.getUserContext();
            Collection<Orderer> orderers = org1Channel.getOrderers();
            Orderer ordererAdmin = orderers.iterator().next();
            ////////////////////////////
            //Reconstruct and run the channels
            //    SampleOrg sampleOrg = testConfig.getIntegrationTestsSampleOrg("peerOrg1");
            Channel channel = reconstructChannel(true, SYSTEM_CHANNEL_NAME, org1Client, orderers);

            log.info("" + channel.getPeers().isEmpty()); // no peers

            // Getting foo channels current configuration bytes.
            byte[] channelConfigurationBytes = channel.getChannelConfigurationBytes(user, channel.getOrderers().iterator().next());

            String originalConfigJson = configTxlatorDecode(httpclient, channelConfigurationBytes);

            log.info("" + originalConfigJson.matches(REGX_IS_SYSTEM_CHANNEL));  // verify is system channel

            //responseAsString is JSON but use just string operations for this test.

            if (!originalConfigJson.contains(ORIGINAL_BATCH_TIMEOUT)) {

                log.info(format("Did not find expected batch timeout '%s', in:%s", ORIGINAL_BATCH_TIMEOUT, originalConfigJson));
            }

            byte[] reEncodedOriginalConfig = configTxLatorEncode(httpclient, originalConfigJson); // we need to get this to make sure the compare has encoding in the same way!

            //Now modify the batch timeout
            String updateString = originalConfigJson.replace(ORIGINAL_BATCH_TIMEOUT, UPDATED_BATCH_TIMEOUT);

            byte[] updatedConfigBytes = configTxLatorEncode(httpclient, updateString);

            // Now send to configtxlator multipart form post with original config bytes, updated config bytes and channel name.
            byte[] updateBytes = getChannelUpdateBytes(channel, reEncodedOriginalConfig, updatedConfigBytes);

            UpdateChannelConfiguration updateChannelConfiguration = new UpdateChannelConfiguration(updateBytes);

            //To change the channel we need to sign with orderer admin certs which crypto gen stores:

            // client.setUserContext(ordererAdmin);
            //Ok now do actual channel update.
            channel.updateChannelConfiguration(user, updateChannelConfiguration,
                    channel.getOrderers().iterator().next(),
                    org1Client.getUpdateChannelConfigurationSignature(updateChannelConfiguration, user));

            Thread.sleep(3000); // give time for events to happen

            //Let's add some additional verification...

            // client.setUserContext(sampleOrg.getPeerAdmin());

            final byte[] modChannelBytes = channel.getChannelConfigurationBytes(user);

            originalConfigJson = configTxlatorDecode(httpclient, modChannelBytes);

            if (!originalConfigJson.contains(UPDATED_BATCH_TIMEOUT)) {
                //If it doesn't have the updated time out it log.infoed.
                log.info(format("Did not find updated expected batch timeout '%s', in:%s", UPDATED_BATCH_TIMEOUT, originalConfigJson));
            }

            if (originalConfigJson.contains(ORIGINAL_BATCH_TIMEOUT)) { //Should not have been there anymore!

                log.info(format("Found original batch timeout '%s', when it was not expected in:%s", ORIGINAL_BATCH_TIMEOUT, originalConfigJson));
            }

            out("That's all folks!");

        } catch (Exception e) {
            e.printStackTrace();
            log.info(e.getMessage());
        }
    }

    private byte[] getChannelUpdateBytes(Channel fooChannel, byte[] reEncodedOriginalConfig, byte[] updatedConfigBytes) throws IOException {
        HttpPost httppost;
        HttpResponse response;

        // Now send to configtxlator multipart form post with original config bytes, updated config bytes and channel name.
        httppost = new HttpPost(CONFIGTXLATOR_LOCATION + "/configtxlator/compute/update-from-configs");

        HttpEntity multipartEntity = MultipartEntityBuilder.create()
                .setMode(HttpMultipartMode.BROWSER_COMPATIBLE)
                .addBinaryBody("original", reEncodedOriginalConfig, ContentType.APPLICATION_OCTET_STREAM, "originalFakeFilename")
                .addBinaryBody("updated", updatedConfigBytes, ContentType.APPLICATION_OCTET_STREAM, "updatedFakeFilename")
                .addBinaryBody("channel", fooChannel.getName().getBytes()).build();

        httppost.setEntity(multipartEntity);

        response = httpclient.execute(httppost);
        int statuscode = response.getStatusLine().getStatusCode();
        out("Got %s status for updated config bytes needed for updateChannelConfiguration ", statuscode);
        log.info("" + statuscode);

        return EntityUtils.toByteArray(response.getEntity());
    }

    private byte[] configTxLatorEncode(HttpClient httpclient, String jsonEncoded) throws IOException {
        HttpPost httppost = new HttpPost(CONFIGTXLATOR_LOCATION + "/protolator/encode/common.Config");
        httppost.setEntity(new StringEntity(jsonEncoded));

        HttpResponse response = httpclient.execute(httppost);

        int statuscode = response.getStatusLine().getStatusCode();
        out("Got %s status for encoding the new desired channel config bytes", statuscode);
        log.info("" + statuscode);
        return EntityUtils.toByteArray(response.getEntity());
    }

    private String configTxlatorDecode(HttpClient httpclient, byte[] channelConfigurationBytes) throws IOException {
        HttpPost httppost = new HttpPost(CONFIGTXLATOR_LOCATION + "/protolator/decode/common.Config");
        httppost.setEntity(new ByteArrayEntity(channelConfigurationBytes));

        HttpResponse response = httpclient.execute(httppost);
        int statuscode = response.getStatusLine().getStatusCode();
        //  out("Got %s status for decoding current channel config bytes", statuscode);
        log.info("" + statuscode);
        return EntityUtils.toString(response.getEntity());
    }

    int eventCountFilteredBlock = 0;
    int eventCountBlock = 0;

    private Channel reconstructChannel(final boolean isSystemChannel, String name, HFClient client, Collection<Orderer> orderers) throws Exception {

        Channel newChannel = client.newChannel(name);

        for (Orderer orderName : orderers) {
            newChannel.addOrderer(orderName);
        }
        if (isSystemChannel) { // done
            newChannel.initialize();
            return newChannel;
        }

        int i = 0;
        for (Peer peer : newChannel.getPeers()) {

            //Query the actual peer for which channels it belongs to and check it belongs to this channel
            Set<String> channels = client.queryChannels(peer);
            if (!channels.contains(name)) {
                throw new AssertionError(format("Peer %s does not appear to belong to channel %s", peer.getName(), name));
            }
            Channel.PeerOptions peerOptions = createPeerOptions().setPeerRoles(EnumSet.of(Peer.PeerRole.CHAINCODE_QUERY,
                    Peer.PeerRole.ENDORSING_PEER, Peer.PeerRole.LEDGER_QUERY, Peer.PeerRole.EVENT_SOURCE));

            if (i % 2 == 0) {
                peerOptions.registerEventsForFilteredBlocks(); // we need a mix of each type for testing.
            } else {
                peerOptions.registerEventsForBlocks();
            }
            ++i;

            newChannel.addPeer(peer, peerOptions);
        }

        //For testing of blocks which are not transactions.
        newChannel.registerBlockListener(blockEvent -> {
            eventQueueCaputure.add(blockEvent); // used with the other queued to make sure same.
            // Note peer eventing will always start with sending the last block so this will get the last endorser block
            int transactions = 0;
            int nonTransactions = 0;
            for (BlockInfo.EnvelopeInfo envelopeInfo : blockEvent.getEnvelopeInfos()) {

                if (BlockInfo.EnvelopeType.TRANSACTION_ENVELOPE == envelopeInfo.getType()) {
                    ++transactions;
                } else {
                    log.info("" + BlockInfo.EnvelopeType.ENVELOPE, envelopeInfo.getType());
                    ++nonTransactions;
                }

            }
            log.info(format("nontransactions %d, transactions %d", nonTransactions, transactions), nonTransactions < 2); // non transaction blocks only have one envelope
            log.info(format("nontransactions %d, transactions %d", nonTransactions, transactions), nonTransactions + transactions > 0); // has to be one.
            log.info(format("nontransactions %d, transactions %d", nonTransactions, transactions), nonTransactions > 0 && transactions > 0); // can't have both.

            if (nonTransactions > 0) { // this is an update block -- don't care about others here.

                if (blockEvent.isFiltered()) {
                    ++eventCountFilteredBlock; // make sure we're seeing non transaction events.
                } else {
                    ++eventCountBlock;
                }
                log.info("" + blockEvent.getTransactionCount());
                log.info("" + blockEvent.getEnvelopeCount());
                for (TransactionEvent transactionEvent : blockEvent.getTransactionEvents()) {
                    log.info("Got transaction event in a block update"); // only events for update should not have transactions.
                }
            }
        });

        // Register Queued block listeners just for testing use both ways.
        // Ideally an application would have it's own independent thread to monitor and take off elements as fast as they can.
        // This would wait forever however if event could not be put in the queue like if the capacity is at a maximum. For LinkedBlockingQueue so unlikely
        listenerHandler1 = newChannel.registerBlockListener(blockingQueue1);
        log.info(listenerHandler1);
        // This is the same but put a timeout on it.  If its not queued in time like if the queue is full it would generate a log warning and ignore the event.
        listenerHandler2 = newChannel.registerBlockListener(blockingQueue2, 1L, TimeUnit.SECONDS);
        log.info(listenerHandler2);

        newChannel.initialize();

        return newChannel;
    }

    // Handles to unregister handlers.
    String listenerHandler1;
    String listenerHandler2;

    BlockingQueue<QueuedBlockEvent> blockingQueue1 = new LinkedBlockingQueue<>(); // really this is unbounded.
    BlockingQueue<QueuedBlockEvent> blockingQueue2 = new ArrayBlockingQueue<>(1000); // application should  pull off queue so not to go full.

    // Have the event handler put into this queue so we can compare.
    BlockingQueue<BlockEvent> eventQueueCaputure = new LinkedBlockingQueue<>();

    static void out(String format, Object... args) {

        System.err.flush();
        System.out.flush();

        System.out.println(format(format, args));
        System.err.flush();
        System.out.flush();

    }

}

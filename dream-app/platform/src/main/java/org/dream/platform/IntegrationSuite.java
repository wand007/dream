/*
 Copyright IBM Corp. All Rights Reserved.

 SPDX-License-Identifier: Apache-2.0
*/
package org.dream.platform;

import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


public class IntegrationSuite {
    public static final Path TEST_FIXTURE_PATH = Paths.get("src", "test", "fixture");

    private static final int fabricMajorVersion = 0;
    private static final int fabricMinorVersion = 1;
    private static final Map<String, List<Class>> runmap = new HashMap<>();
    private static final Path SDK_INTEGRATION_PATH = TEST_FIXTURE_PATH.resolve("sdkintegration");
    private static final Path GO_CHAINCODE_PATH = SDK_INTEGRATION_PATH.resolve("gocc");
    private static final Path NODE_CHAINCODE_PATH = SDK_INTEGRATION_PATH.resolve("nodecc");
    private static final Path JAVA_CHAINCODE_PATH = SDK_INTEGRATION_PATH.resolve("javacc");


    public static Path getGoChaincodePath(String chaincodeName) {
        return GO_CHAINCODE_PATH.resolve(chaincodeName);
    }

    public static Path getNodeChaincodePath(String chaincodeName) {
        return NODE_CHAINCODE_PATH.resolve(chaincodeName);
    }

    public static Path getJavaChaincodePath(String chaincodeName) {
        final Path chaincodeRootPath;
        if (fabricMajorVersion == 1 && fabricMinorVersion == 4) {
            chaincodeRootPath = JAVA_CHAINCODE_PATH.resolve("1.4");
        } else if (fabricMajorVersion == 2) {
            chaincodeRootPath = JAVA_CHAINCODE_PATH.resolve("2.1");
        } else {
            throw new RuntimeException(String.format("Unexpected Fabric version for Java chaincode: %d.%d",
                    fabricMajorVersion, fabricMinorVersion));
        }

        return chaincodeRootPath.resolve(chaincodeName);
    }

    private void checkStyleWorkAround() {  //avoid utility class issue
    }

}
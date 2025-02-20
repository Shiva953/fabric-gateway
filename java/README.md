# Hyperledger Fabric Gateway Client API for Java

The Fabric Gateway client API allows applications to interact with a Hyperledger Fabric blockchain network. It implements the Fabric programming model, providing a simple API to submit transactions to a ledger or query the contents of a ledger with minimal code.

## How to use

Samples showing how to create client applications that connect to and interact with a Hyperledger Fabric network, are available in the [fabric-samples](https://github.com/hyperledger/fabric-samples) repository:

- [fabric-samples/asset-transfer-basic](https://github.com/hyperledger/fabric-samples/tree/main/asset-transfer-basic) for examples of transaction submit and evaluate.
- [fabric-samples/asset-transfer-events](https://github.com/hyperledger/fabric-samples/tree/main/asset-transfer-events) for examples of chaincode eventing.
- [fabric-samples/off_chain_data](https://github.com/hyperledger/fabric-samples/tree/main/off_chain_data) for examples of block eventing.

## API documentation

The Gateway client API documentation for Java is available here:

- https://hyperledger.github.io/fabric-gateway/main/api/java/

## Installation

The Fabric Gateway client API package is published to [Maven Central](https://search.maven.org/artifact/org.hyperledger.fabric/fabric-gateway).

### Maven

Add the following dependency to your project's `pom.xml` file:

```xml
<dependency>
    <groupId>org.hyperledger.fabric</groupId>
    <artifactId>fabric-gateway</artifactId>
    <version>1.3.0</version>
</dependency>
```

A suitable gRPC channel service provider must also be specified (as described in the [gRPC security documentation](https://github.com/grpc/grpc-java/blob/master/SECURITY.md#transport-security-tls)), such as:

```xml
<dependency>
    <groupId>io.grpc</groupId>
    <artifactId>grpc-netty-shaded</artifactId>
    <version>1.56.0</version>
    <scope>runtime</scope>
</dependency>
```

### Gradle

Add the following dependency to your project's `build.gradle` file:

```groovy
implementation 'org.hyperledger.fabric:fabric-gateway:1.3.0'
```

A suitable gRPC channel service provider must also be specified (as described in the [gRPC security documentation](https://github.com/grpc/grpc-java/blob/master/SECURITY.md#transport-security-tls)), such as:

```groovy
runtimeOnly 'io.grpc:grpc-netty-shaded:1.56.0'
```

## Compatibility

This API requires Fabric v2.4 (or later) with a Gateway enabled Peer. Additional compatibility information is available in the documentation:

- https://hyperledger.github.io/fabric-gateway/

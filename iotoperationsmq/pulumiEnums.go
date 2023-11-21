// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperationsmq

// Memory profile of broker.
type BrokerMemoryProfile string

const (
	// Tiny memory profile.
	BrokerMemoryProfileTiny = BrokerMemoryProfile("tiny")
	// Low memory profile.
	BrokerMemoryProfileLow = BrokerMemoryProfile("low")
	// Medium memory profile.
	BrokerMemoryProfileMedium = BrokerMemoryProfile("medium")
	// High memory profile.
	BrokerMemoryProfileHigh = BrokerMemoryProfile("high")
)

// DataLake database format to use.
type DataLakeDatabaseFormat string

const (
	// Delta format.
	DataLakeDatabaseFormatDelta = DataLakeDatabaseFormat("delta")
	// Parquet format.
	DataLakeDatabaseFormatParquet = DataLakeDatabaseFormat("parquet")
)

// Delta table format supported.
type DeltaTableFormatEnum string

const (
	// Bool format
	DeltaTableFormatEnumBoolean = DeltaTableFormatEnum("boolean")
	// Signed integer 8
	DeltaTableFormatEnumInt8 = DeltaTableFormatEnum("int8")
	// Signed integer 16
	DeltaTableFormatEnumInt16 = DeltaTableFormatEnum("int16")
	// Signed integer 32
	DeltaTableFormatEnumInt32 = DeltaTableFormatEnum("int32")
	// Unsigned integer 8
	DeltaTableFormatEnumUInt8 = DeltaTableFormatEnum("uInt8")
	// Unsigned integer 16
	DeltaTableFormatEnumUInt16 = DeltaTableFormatEnum("uInt16")
	// Unsigned integer 32
	DeltaTableFormatEnumUInt32 = DeltaTableFormatEnum("uInt32")
	// Unsigned integer 64
	DeltaTableFormatEnumUInt64 = DeltaTableFormatEnum("uInt64")
	// Float 16
	DeltaTableFormatEnumFloat16 = DeltaTableFormatEnum("float16")
	// Float 32
	DeltaTableFormatEnumFloat32 = DeltaTableFormatEnum("float32")
	// Float 64
	DeltaTableFormatEnumFloat64 = DeltaTableFormatEnum("float64")
	// Date 32
	DeltaTableFormatEnumDate32 = DeltaTableFormatEnum("date32")
	// Date 64
	DeltaTableFormatEnumDate64 = DeltaTableFormatEnum("date64")
	// Binary data
	DeltaTableFormatEnumBinary = DeltaTableFormatEnum("binary")
	// UTF8 format
	DeltaTableFormatEnumUtf8 = DeltaTableFormatEnum("utf8")
)

// Type of ExtendedLocation.
type ExtendedLocationType string

const (
	// CustomLocation type
	ExtendedLocationTypeCustomLocation = ExtendedLocationType("CustomLocation")
)

// Fabric path type to use.
type FabricPathType string

const (
	// Fabric path type is Files.
	FabricPathTypeFiles = FabricPathType("files")
	// Fabric path type is Tables.
	FabricPathTypeTables = FabricPathType("tables")
)

// The kafka acks to use.
type KafkaAcks string

const (
	// Kafka acks zero.
	KafkaAcksZero = KafkaAcks("zero")
	// Kafka acks one.
	KafkaAcksOne = KafkaAcks("one")
	// Kafka acks all.
	KafkaAcksAll = KafkaAcks("all")
)

// The compression to use for kafka messages.
type KafkaMessageCompressionType string

const (
	// No Kafka message compression.
	KafkaMessageCompressionTypeNone = KafkaMessageCompressionType("none")
	// Gzip Kafka message compression.
	KafkaMessageCompressionTypeGzip = KafkaMessageCompressionType("gzip")
	// Snappy Kafka message compression.
	KafkaMessageCompressionTypeSnappy = KafkaMessageCompressionType("snappy")
	// Lz4 Kafka message compression.
	KafkaMessageCompressionTypeLz4 = KafkaMessageCompressionType("lz4")
)

// The partition strategy to use for Kafka.
type KafkaPartitionStrategy string

const (
	// Default partition strategy.
	KafkaPartitionStrategyDefault = KafkaPartitionStrategy("default")
	// Static partition strategy.
	KafkaPartitionStrategyStatic = KafkaPartitionStrategy("static")
	// Topic partition strategy.
	KafkaPartitionStrategyTopic = KafkaPartitionStrategy("topic")
	// Property partition strategy.
	KafkaPartitionStrategyProperty = KafkaPartitionStrategy("property")
)

// Sasl Mechanism for remote broker authentication.
type KafkaSaslType string

const (
	// Sasl Plain authentication.
	KafkaSaslTypePlain = KafkaSaslType("plain")
	// Sasl ScramSha256 authentication.
	KafkaSaslTypeScramSha256 = KafkaSaslType("scramSha256")
	// Sasl ScramSha512 authentication.
	KafkaSaslTypeScramSha512 = KafkaSaslType("scramSha512")
)

// Protocol for remote connection.
type MqttBridgeRemoteBrokerProtocol string

const (
	// MQTT protocol.
	MqttBridgeRemoteBrokerProtocolMqtt = MqttBridgeRemoteBrokerProtocol("mqtt")
	// MQTT over WebSocket protocol.
	MqttBridgeRemoteBrokerProtocolWebSocket = MqttBridgeRemoteBrokerProtocol("webSocket")
)

// Direction of the route.
type MqttBridgeRouteDirection string

const (
	// Remote to Local Broker.
	MqttBridgeRouteDirectionRemoteToLocal = MqttBridgeRouteDirection("remote-to-local")
	// Local to Remote Broker.
	MqttBridgeRouteDirectionLocalToRemote = MqttBridgeRouteDirection("local-to-remote")
)

// The protocol to use for connecting with Brokers.
type MqttProtocol string

const (
	// Mqttv3
	MqttProtocolV3 = MqttProtocol("v3")
	// Mqttv5
	MqttProtocolV5 = MqttProtocol("v5")
)

// The type of action that the clients can perform on the broker: Connect, Publish or Subscribe.
type ResourceInfoDefinitionMethods string

const (
	// Allowed Connecting to Broker
	ResourceInfoDefinitionMethodsConnect = ResourceInfoDefinitionMethods("Connect")
	// Allowed Publishing to Broker
	ResourceInfoDefinitionMethodsPublish = ResourceInfoDefinitionMethods("Publish")
	// Allowed Subscribing to Broker
	ResourceInfoDefinitionMethodsSubscribe = ResourceInfoDefinitionMethods("Subscribe")
)

// The Running Mode of the Broker Deployment.
type RunMode string

const (
	// Automatically provision Frontend and Backend pods.
	RunModeAuto = RunMode("auto")
	// Use Cardinality to set Frontend and Backend pods.
	RunModeDistributed = RunMode("distributed")
)

// The Kubernetes Service type to deploy for Listener.
type ServiceType string

const (
	// Cluster IP Service.
	ServiceTypeClusterIp = ServiceType("clusterIp")
	// Load Balancer Service.
	ServiceTypeLoadBalancer = ServiceType("loadBalancer")
	// Node Port Service.
	ServiceTypeNodePort = ServiceType("nodePort")
)

func init() {
}
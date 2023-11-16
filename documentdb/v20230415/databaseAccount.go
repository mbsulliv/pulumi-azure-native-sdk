// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB database account.
type DatabaseAccount struct {
	pulumi.CustomResourceState

	// Analytical storage specific properties.
	AnalyticalStorageConfiguration AnalyticalStorageConfigurationResponsePtrOutput `pulumi:"analyticalStorageConfiguration"`
	// API specific properties.
	ApiProperties ApiPropertiesResponsePtrOutput `pulumi:"apiProperties"`
	// The object representing the policy for taking backups on an account.
	BackupPolicy pulumi.AnyOutput `pulumi:"backupPolicy"`
	// List of Cosmos DB capabilities for the account
	Capabilities CapabilityResponseArrayOutput `pulumi:"capabilities"`
	// The object that represents all properties related to capacity enforcement on an account.
	Capacity CapacityResponsePtrOutput `pulumi:"capacity"`
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer pulumi.StringPtrOutput `pulumi:"connectorOffer"`
	// The consistency policy for the Cosmos DB database account.
	ConsistencyPolicy ConsistencyPolicyResponsePtrOutput `pulumi:"consistencyPolicy"`
	// The CORS policy for the Cosmos DB database account.
	Cors CorsPolicyResponseArrayOutput `pulumi:"cors"`
	// Enum to indicate the mode of account creation.
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// The offer type for the Cosmos DB database account. Default value: Standard.
	DatabaseAccountOfferType pulumi.StringOutput `pulumi:"databaseAccountOfferType"`
	// The default identity for accessing key vault used in features like customer managed keys. The default identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity pulumi.StringPtrOutput `pulumi:"defaultIdentity"`
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess pulumi.BoolPtrOutput `pulumi:"disableKeyBasedMetadataWriteAccess"`
	// Opt-out of local authentication and ensure only MSI and AAD can be used exclusively for authentication.
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// The connection endpoint for the Cosmos DB database account.
	DocumentEndpoint pulumi.StringOutput `pulumi:"documentEndpoint"`
	// Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage pulumi.BoolPtrOutput `pulumi:"enableAnalyticalStorage"`
	// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
	EnableAutomaticFailover pulumi.BoolPtrOutput `pulumi:"enableAutomaticFailover"`
	// Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector pulumi.BoolPtrOutput `pulumi:"enableCassandraConnector"`
	// Flag to indicate whether Free Tier is enabled.
	EnableFreeTier pulumi.BoolPtrOutput `pulumi:"enableFreeTier"`
	// Enables the account to write in multiple locations
	EnableMultipleWriteLocations pulumi.BoolPtrOutput `pulumi:"enableMultipleWriteLocations"`
	// Flag to indicate enabling/disabling of Partition Merge feature on the account
	EnablePartitionMerge pulumi.BoolPtrOutput `pulumi:"enablePartitionMerge"`
	// An array that contains the regions ordered by their failover priorities.
	FailoverPolicies FailoverPolicyResponseArrayOutput `pulumi:"failoverPolicies"`
	// Identity for the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// A unique identifier assigned to the database account
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// List of IpRules.
	IpRules IpAddressOrRangeResponseArrayOutput `pulumi:"ipRules"`
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrOutput `pulumi:"isVirtualNetworkFilterEnabled"`
	// The URI of the key vault
	KeyVaultKeyUri pulumi.StringPtrOutput `pulumi:"keyVaultKeyUri"`
	// The object that represents the metadata for the Account Keys of the Cosmos DB account.
	KeysMetadata DatabaseAccountKeysMetadataResponseOutput `pulumi:"keysMetadata"`
	// Indicates the type of database account. This can only be set at database account creation.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// An array that contains all of the locations enabled for the Cosmos DB account.
	Locations LocationResponseArrayOutput `pulumi:"locations"`
	// Indicates the minimum allowed Tls version. The default value is Tls 1.2. Cassandra and Mongo APIs only work with Tls 1.2.
	MinimalTlsVersion pulumi.StringPtrOutput `pulumi:"minimalTlsVersion"`
	// The name of the ARM resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass pulumi.StringPtrOutput `pulumi:"networkAclBypass"`
	// An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds pulumi.StringArrayOutput `pulumi:"networkAclBypassResourceIds"`
	// List of Private Endpoint Connections configured for the Cosmos DB account.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Whether requests from Public Network are allowed
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// An array that contains of the read locations enabled for the Cosmos DB account.
	ReadLocations LocationResponseArrayOutput `pulumi:"readLocations"`
	// Parameters to indicate the information about the restore.
	RestoreParameters RestoreParametersResponsePtrOutput `pulumi:"restoreParameters"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules VirtualNetworkRuleResponseArrayOutput `pulumi:"virtualNetworkRules"`
	// An array that contains the write location for the Cosmos DB account.
	WriteLocations LocationResponseArrayOutput `pulumi:"writeLocations"`
}

// NewDatabaseAccount registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccount(ctx *pulumi.Context,
	name string, args *DatabaseAccountArgs, opts ...pulumi.ResourceOption) (*DatabaseAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseAccountOfferType == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseAccountOfferType'")
	}
	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.CreateMode == nil {
		args.CreateMode = pulumi.StringPtr("Default")
	}
	if args.Kind == nil {
		args.Kind = pulumi.StringPtr("GlobalDocumentDB")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:DatabaseAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DatabaseAccount
	err := ctx.RegisterResource("azure-native:documentdb/v20230415:DatabaseAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccount gets an existing DatabaseAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountState, opts ...pulumi.ResourceOption) (*DatabaseAccount, error) {
	var resource DatabaseAccount
	err := ctx.ReadResource("azure-native:documentdb/v20230415:DatabaseAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccount resources.
type databaseAccountState struct {
}

type DatabaseAccountState struct {
}

func (DatabaseAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountState)(nil)).Elem()
}

type databaseAccountArgs struct {
	// Cosmos DB database account name.
	AccountName *string `pulumi:"accountName"`
	// Analytical storage specific properties.
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration `pulumi:"analyticalStorageConfiguration"`
	// API specific properties. Currently, supported only for MongoDB API.
	ApiProperties *ApiProperties `pulumi:"apiProperties"`
	// The object representing the policy for taking backups on an account.
	BackupPolicy interface{} `pulumi:"backupPolicy"`
	// List of Cosmos DB capabilities for the account
	Capabilities []Capability `pulumi:"capabilities"`
	// The object that represents all properties related to capacity enforcement on an account.
	Capacity *Capacity `pulumi:"capacity"`
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer *string `pulumi:"connectorOffer"`
	// The consistency policy for the Cosmos DB account.
	ConsistencyPolicy *ConsistencyPolicy `pulumi:"consistencyPolicy"`
	// The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicy `pulumi:"cors"`
	// Enum to indicate the mode of account creation.
	CreateMode *string `pulumi:"createMode"`
	// The offer type for the database
	DatabaseAccountOfferType DatabaseAccountOfferType `pulumi:"databaseAccountOfferType"`
	// The default identity for accessing key vault used in features like customer managed keys. The default identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity *string `pulumi:"defaultIdentity"`
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess *bool `pulumi:"disableKeyBasedMetadataWriteAccess"`
	// Opt-out of local authentication and ensure only MSI and AAD can be used exclusively for authentication.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `pulumi:"enableAnalyticalStorage"`
	// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
	EnableAutomaticFailover *bool `pulumi:"enableAutomaticFailover"`
	// Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector *bool `pulumi:"enableCassandraConnector"`
	// Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `pulumi:"enableFreeTier"`
	// Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `pulumi:"enableMultipleWriteLocations"`
	// Flag to indicate enabling/disabling of Partition Merge feature on the account
	EnablePartitionMerge *bool `pulumi:"enablePartitionMerge"`
	// Identity for the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// List of IpRules.
	IpRules []IpAddressOrRange `pulumi:"ipRules"`
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// The URI of the key vault
	KeyVaultKeyUri *string `pulumi:"keyVaultKeyUri"`
	// Indicates the type of database account. This can only be set at database account creation.
	Kind *string `pulumi:"kind"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// An array that contains the georeplication locations enabled for the Cosmos DB account.
	Locations []Location `pulumi:"locations"`
	// Indicates the minimum allowed Tls version. The default value is Tls 1.2. Cassandra and Mongo APIs only work with Tls 1.2.
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass *NetworkAclBypass `pulumi:"networkAclBypass"`
	// An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds []string `pulumi:"networkAclBypassResourceIds"`
	// Whether requests from Public Network are allowed
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Parameters to indicate the information about the restore.
	RestoreParameters *RestoreParameters `pulumi:"restoreParameters"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}

// The set of arguments for constructing a DatabaseAccount resource.
type DatabaseAccountArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringPtrInput
	// Analytical storage specific properties.
	AnalyticalStorageConfiguration AnalyticalStorageConfigurationPtrInput
	// API specific properties. Currently, supported only for MongoDB API.
	ApiProperties ApiPropertiesPtrInput
	// The object representing the policy for taking backups on an account.
	BackupPolicy pulumi.Input
	// List of Cosmos DB capabilities for the account
	Capabilities CapabilityArrayInput
	// The object that represents all properties related to capacity enforcement on an account.
	Capacity CapacityPtrInput
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer pulumi.StringPtrInput
	// The consistency policy for the Cosmos DB account.
	ConsistencyPolicy ConsistencyPolicyPtrInput
	// The CORS policy for the Cosmos DB database account.
	Cors CorsPolicyArrayInput
	// Enum to indicate the mode of account creation.
	CreateMode pulumi.StringPtrInput
	// The offer type for the database
	DatabaseAccountOfferType DatabaseAccountOfferTypeInput
	// The default identity for accessing key vault used in features like customer managed keys. The default identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity pulumi.StringPtrInput
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess pulumi.BoolPtrInput
	// Opt-out of local authentication and ensure only MSI and AAD can be used exclusively for authentication.
	DisableLocalAuth pulumi.BoolPtrInput
	// Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage pulumi.BoolPtrInput
	// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
	EnableAutomaticFailover pulumi.BoolPtrInput
	// Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector pulumi.BoolPtrInput
	// Flag to indicate whether Free Tier is enabled.
	EnableFreeTier pulumi.BoolPtrInput
	// Enables the account to write in multiple locations
	EnableMultipleWriteLocations pulumi.BoolPtrInput
	// Flag to indicate enabling/disabling of Partition Merge feature on the account
	EnablePartitionMerge pulumi.BoolPtrInput
	// Identity for the resource.
	Identity ManagedServiceIdentityPtrInput
	// List of IpRules.
	IpRules IpAddressOrRangeArrayInput
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrInput
	// The URI of the key vault
	KeyVaultKeyUri pulumi.StringPtrInput
	// Indicates the type of database account. This can only be set at database account creation.
	Kind pulumi.StringPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// An array that contains the georeplication locations enabled for the Cosmos DB account.
	Locations LocationArrayInput
	// Indicates the minimum allowed Tls version. The default value is Tls 1.2. Cassandra and Mongo APIs only work with Tls 1.2.
	MinimalTlsVersion pulumi.StringPtrInput
	// Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass NetworkAclBypassPtrInput
	// An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds pulumi.StringArrayInput
	// Whether requests from Public Network are allowed
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Parameters to indicate the information about the restore.
	RestoreParameters RestoreParametersPtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules VirtualNetworkRuleArrayInput
}

func (DatabaseAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountArgs)(nil)).Elem()
}

type DatabaseAccountInput interface {
	pulumi.Input

	ToDatabaseAccountOutput() DatabaseAccountOutput
	ToDatabaseAccountOutputWithContext(ctx context.Context) DatabaseAccountOutput
}

func (*DatabaseAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccount)(nil)).Elem()
}

func (i *DatabaseAccount) ToDatabaseAccountOutput() DatabaseAccountOutput {
	return i.ToDatabaseAccountOutputWithContext(context.Background())
}

func (i *DatabaseAccount) ToDatabaseAccountOutputWithContext(ctx context.Context) DatabaseAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountOutput)
}

type DatabaseAccountOutput struct{ *pulumi.OutputState }

func (DatabaseAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccount)(nil)).Elem()
}

func (o DatabaseAccountOutput) ToDatabaseAccountOutput() DatabaseAccountOutput {
	return o
}

func (o DatabaseAccountOutput) ToDatabaseAccountOutputWithContext(ctx context.Context) DatabaseAccountOutput {
	return o
}

// Analytical storage specific properties.
func (o DatabaseAccountOutput) AnalyticalStorageConfiguration() AnalyticalStorageConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) AnalyticalStorageConfigurationResponsePtrOutput {
		return v.AnalyticalStorageConfiguration
	}).(AnalyticalStorageConfigurationResponsePtrOutput)
}

// API specific properties.
func (o DatabaseAccountOutput) ApiProperties() ApiPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) ApiPropertiesResponsePtrOutput { return v.ApiProperties }).(ApiPropertiesResponsePtrOutput)
}

// The object representing the policy for taking backups on an account.
func (o DatabaseAccountOutput) BackupPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.AnyOutput { return v.BackupPolicy }).(pulumi.AnyOutput)
}

// List of Cosmos DB capabilities for the account
func (o DatabaseAccountOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) CapabilityResponseArrayOutput { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

// The object that represents all properties related to capacity enforcement on an account.
func (o DatabaseAccountOutput) Capacity() CapacityResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) CapacityResponsePtrOutput { return v.Capacity }).(CapacityResponsePtrOutput)
}

// The cassandra connector offer type for the Cosmos DB database C* account.
func (o DatabaseAccountOutput) ConnectorOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.ConnectorOffer }).(pulumi.StringPtrOutput)
}

// The consistency policy for the Cosmos DB database account.
func (o DatabaseAccountOutput) ConsistencyPolicy() ConsistencyPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) ConsistencyPolicyResponsePtrOutput { return v.ConsistencyPolicy }).(ConsistencyPolicyResponsePtrOutput)
}

// The CORS policy for the Cosmos DB database account.
func (o DatabaseAccountOutput) Cors() CorsPolicyResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) CorsPolicyResponseArrayOutput { return v.Cors }).(CorsPolicyResponseArrayOutput)
}

// Enum to indicate the mode of account creation.
func (o DatabaseAccountOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.CreateMode }).(pulumi.StringPtrOutput)
}

// The offer type for the Cosmos DB database account. Default value: Standard.
func (o DatabaseAccountOutput) DatabaseAccountOfferType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.DatabaseAccountOfferType }).(pulumi.StringOutput)
}

// The default identity for accessing key vault used in features like customer managed keys. The default identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
func (o DatabaseAccountOutput) DefaultIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.DefaultIdentity }).(pulumi.StringPtrOutput)
}

// Disable write operations on metadata resources (databases, containers, throughput) via account keys
func (o DatabaseAccountOutput) DisableKeyBasedMetadataWriteAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.DisableKeyBasedMetadataWriteAccess }).(pulumi.BoolPtrOutput)
}

// Opt-out of local authentication and ensure only MSI and AAD can be used exclusively for authentication.
func (o DatabaseAccountOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// The connection endpoint for the Cosmos DB database account.
func (o DatabaseAccountOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

// Flag to indicate whether to enable storage analytics.
func (o DatabaseAccountOutput) EnableAnalyticalStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableAnalyticalStorage }).(pulumi.BoolPtrOutput)
}

// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
func (o DatabaseAccountOutput) EnableAutomaticFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableAutomaticFailover }).(pulumi.BoolPtrOutput)
}

// Enables the cassandra connector on the Cosmos DB C* account
func (o DatabaseAccountOutput) EnableCassandraConnector() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableCassandraConnector }).(pulumi.BoolPtrOutput)
}

// Flag to indicate whether Free Tier is enabled.
func (o DatabaseAccountOutput) EnableFreeTier() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableFreeTier }).(pulumi.BoolPtrOutput)
}

// Enables the account to write in multiple locations
func (o DatabaseAccountOutput) EnableMultipleWriteLocations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableMultipleWriteLocations }).(pulumi.BoolPtrOutput)
}

// Flag to indicate enabling/disabling of Partition Merge feature on the account
func (o DatabaseAccountOutput) EnablePartitionMerge() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnablePartitionMerge }).(pulumi.BoolPtrOutput)
}

// An array that contains the regions ordered by their failover priorities.
func (o DatabaseAccountOutput) FailoverPolicies() FailoverPolicyResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) FailoverPolicyResponseArrayOutput { return v.FailoverPolicies }).(FailoverPolicyResponseArrayOutput)
}

// Identity for the resource.
func (o DatabaseAccountOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// A unique identifier assigned to the database account
func (o DatabaseAccountOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// List of IpRules.
func (o DatabaseAccountOutput) IpRules() IpAddressOrRangeResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) IpAddressOrRangeResponseArrayOutput { return v.IpRules }).(IpAddressOrRangeResponseArrayOutput)
}

// Flag to indicate whether to enable/disable Virtual Network ACL rules.
func (o DatabaseAccountOutput) IsVirtualNetworkFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.IsVirtualNetworkFilterEnabled }).(pulumi.BoolPtrOutput)
}

// The URI of the key vault
func (o DatabaseAccountOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

// The object that represents the metadata for the Account Keys of the Cosmos DB account.
func (o DatabaseAccountOutput) KeysMetadata() DatabaseAccountKeysMetadataResponseOutput {
	return o.ApplyT(func(v *DatabaseAccount) DatabaseAccountKeysMetadataResponseOutput { return v.KeysMetadata }).(DatabaseAccountKeysMetadataResponseOutput)
}

// Indicates the type of database account. This can only be set at database account creation.
func (o DatabaseAccountOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o DatabaseAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// An array that contains all of the locations enabled for the Cosmos DB account.
func (o DatabaseAccountOutput) Locations() LocationResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) LocationResponseArrayOutput { return v.Locations }).(LocationResponseArrayOutput)
}

// Indicates the minimum allowed Tls version. The default value is Tls 1.2. Cassandra and Mongo APIs only work with Tls 1.2.
func (o DatabaseAccountOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o DatabaseAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates what services are allowed to bypass firewall checks.
func (o DatabaseAccountOutput) NetworkAclBypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.NetworkAclBypass }).(pulumi.StringPtrOutput)
}

// An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
func (o DatabaseAccountOutput) NetworkAclBypassResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringArrayOutput { return v.NetworkAclBypassResourceIds }).(pulumi.StringArrayOutput)
}

// List of Private Endpoint Connections configured for the Cosmos DB account.
func (o DatabaseAccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
func (o DatabaseAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether requests from Public Network are allowed
func (o DatabaseAccountOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// An array that contains of the read locations enabled for the Cosmos DB account.
func (o DatabaseAccountOutput) ReadLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) LocationResponseArrayOutput { return v.ReadLocations }).(LocationResponseArrayOutput)
}

// Parameters to indicate the information about the restore.
func (o DatabaseAccountOutput) RestoreParameters() RestoreParametersResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) RestoreParametersResponsePtrOutput { return v.RestoreParameters }).(RestoreParametersResponsePtrOutput)
}

// The system meta data relating to this resource.
func (o DatabaseAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DatabaseAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o DatabaseAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o DatabaseAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// List of Virtual Network ACL rules configured for the Cosmos DB account.
func (o DatabaseAccountOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) VirtualNetworkRuleResponseArrayOutput { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

// An array that contains the write location for the Cosmos DB account.
func (o DatabaseAccountOutput) WriteLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) LocationResponseArrayOutput { return v.WriteLocations }).(LocationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountOutput{})
}

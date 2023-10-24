// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the properties of an existing Azure Cosmos DB database account.
// Azure REST API version: 2023-04-15.
//
// Other available API versions: 2020-03-01, 2020-06-01-preview, 2020-09-01, 2021-04-01-preview, 2023-03-15-preview, 2023-09-15, 2023-09-15-preview.
func LookupDatabaseAccount(ctx *pulumi.Context, args *LookupDatabaseAccountArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseAccountResult
	err := ctx.Invoke("azure-native:documentdb:getDatabaseAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDatabaseAccountArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB database account.
type LookupDatabaseAccountResult struct {
	// Analytical storage specific properties.
	AnalyticalStorageConfiguration *AnalyticalStorageConfigurationResponse `pulumi:"analyticalStorageConfiguration"`
	// API specific properties.
	ApiProperties *ApiPropertiesResponse `pulumi:"apiProperties"`
	// The object representing the policy for taking backups on an account.
	BackupPolicy interface{} `pulumi:"backupPolicy"`
	// List of Cosmos DB capabilities for the account
	Capabilities []CapabilityResponse `pulumi:"capabilities"`
	// The object that represents all properties related to capacity enforcement on an account.
	Capacity *CapacityResponse `pulumi:"capacity"`
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer *string `pulumi:"connectorOffer"`
	// The consistency policy for the Cosmos DB database account.
	ConsistencyPolicy *ConsistencyPolicyResponse `pulumi:"consistencyPolicy"`
	// The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicyResponse `pulumi:"cors"`
	// Enum to indicate the mode of account creation.
	CreateMode *string `pulumi:"createMode"`
	// The offer type for the Cosmos DB database account. Default value: Standard.
	DatabaseAccountOfferType string `pulumi:"databaseAccountOfferType"`
	// The default identity for accessing key vault used in features like customer managed keys. The default identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
	DefaultIdentity *string `pulumi:"defaultIdentity"`
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess *bool `pulumi:"disableKeyBasedMetadataWriteAccess"`
	// Opt-out of local authentication and ensure only MSI and AAD can be used exclusively for authentication.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// The connection endpoint for the Cosmos DB database account.
	DocumentEndpoint string `pulumi:"documentEndpoint"`
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
	// An array that contains the regions ordered by their failover priorities.
	FailoverPolicies []FailoverPolicyResponse `pulumi:"failoverPolicies"`
	// The unique resource identifier of the ARM resource.
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// A unique identifier assigned to the database account
	InstanceId string `pulumi:"instanceId"`
	// List of IpRules.
	IpRules []IpAddressOrRangeResponse `pulumi:"ipRules"`
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// The URI of the key vault
	KeyVaultKeyUri *string `pulumi:"keyVaultKeyUri"`
	// The object that represents the metadata for the Account Keys of the Cosmos DB account.
	KeysMetadata DatabaseAccountKeysMetadataResponse `pulumi:"keysMetadata"`
	// Indicates the type of database account. This can only be set at database account creation.
	Kind *string `pulumi:"kind"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// An array that contains all of the locations enabled for the Cosmos DB account.
	Locations []LocationResponse `pulumi:"locations"`
	// Indicates the minimum allowed Tls version. The default value is Tls 1.2. Cassandra and Mongo APIs only work with Tls 1.2.
	MinimalTlsVersion *string `pulumi:"minimalTlsVersion"`
	// The name of the ARM resource.
	Name string `pulumi:"name"`
	// Indicates what services are allowed to bypass firewall checks.
	NetworkAclBypass *string `pulumi:"networkAclBypass"`
	// An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
	NetworkAclBypassResourceIds []string `pulumi:"networkAclBypassResourceIds"`
	// List of Private Endpoint Connections configured for the Cosmos DB account.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
	ProvisioningState string `pulumi:"provisioningState"`
	// Whether requests from Public Network are allowed
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// An array that contains of the read locations enabled for the Cosmos DB account.
	ReadLocations []LocationResponse `pulumi:"readLocations"`
	// Parameters to indicate the information about the restore.
	RestoreParameters *RestoreParametersResponse `pulumi:"restoreParameters"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
	// An array that contains the write location for the Cosmos DB account.
	WriteLocations []LocationResponse `pulumi:"writeLocations"`
}

// Defaults sets the appropriate defaults for LookupDatabaseAccountResult
func (val *LookupDatabaseAccountResult) Defaults() *LookupDatabaseAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.CreateMode == nil {
		createMode_ := "Default"
		tmp.CreateMode = &createMode_
	}
	if tmp.Kind == nil {
		kind_ := "GlobalDocumentDB"
		tmp.Kind = &kind_
	}
	return &tmp
}

func LookupDatabaseAccountOutput(ctx *pulumi.Context, args LookupDatabaseAccountOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseAccountResult, error) {
			args := v.(LookupDatabaseAccountArgs)
			r, err := LookupDatabaseAccount(ctx, &args, opts...)
			var s LookupDatabaseAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseAccountResultOutput)
}

type LookupDatabaseAccountOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountArgs)(nil)).Elem()
}

// An Azure Cosmos DB database account.
type LookupDatabaseAccountResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseAccountResult)(nil)).Elem()
}

func (o LookupDatabaseAccountResultOutput) ToLookupDatabaseAccountResultOutput() LookupDatabaseAccountResultOutput {
	return o
}

func (o LookupDatabaseAccountResultOutput) ToLookupDatabaseAccountResultOutputWithContext(ctx context.Context) LookupDatabaseAccountResultOutput {
	return o
}

func (o LookupDatabaseAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDatabaseAccountResult] {
	return pulumix.Output[LookupDatabaseAccountResult]{
		OutputState: o.OutputState,
	}
}

// Analytical storage specific properties.
func (o LookupDatabaseAccountResultOutput) AnalyticalStorageConfiguration() AnalyticalStorageConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *AnalyticalStorageConfigurationResponse {
		return v.AnalyticalStorageConfiguration
	}).(AnalyticalStorageConfigurationResponsePtrOutput)
}

// API specific properties.
func (o LookupDatabaseAccountResultOutput) ApiProperties() ApiPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *ApiPropertiesResponse { return v.ApiProperties }).(ApiPropertiesResponsePtrOutput)
}

// The object representing the policy for taking backups on an account.
func (o LookupDatabaseAccountResultOutput) BackupPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) interface{} { return v.BackupPolicy }).(pulumi.AnyOutput)
}

// List of Cosmos DB capabilities for the account
func (o LookupDatabaseAccountResultOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []CapabilityResponse { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

// The object that represents all properties related to capacity enforcement on an account.
func (o LookupDatabaseAccountResultOutput) Capacity() CapacityResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *CapacityResponse { return v.Capacity }).(CapacityResponsePtrOutput)
}

// The cassandra connector offer type for the Cosmos DB database C* account.
func (o LookupDatabaseAccountResultOutput) ConnectorOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.ConnectorOffer }).(pulumi.StringPtrOutput)
}

// The consistency policy for the Cosmos DB database account.
func (o LookupDatabaseAccountResultOutput) ConsistencyPolicy() ConsistencyPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *ConsistencyPolicyResponse { return v.ConsistencyPolicy }).(ConsistencyPolicyResponsePtrOutput)
}

// The CORS policy for the Cosmos DB database account.
func (o LookupDatabaseAccountResultOutput) Cors() CorsPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []CorsPolicyResponse { return v.Cors }).(CorsPolicyResponseArrayOutput)
}

// Enum to indicate the mode of account creation.
func (o LookupDatabaseAccountResultOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

// The offer type for the Cosmos DB database account. Default value: Standard.
func (o LookupDatabaseAccountResultOutput) DatabaseAccountOfferType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.DatabaseAccountOfferType }).(pulumi.StringOutput)
}

// The default identity for accessing key vault used in features like customer managed keys. The default identity needs to be explicitly set by the users. It can be "FirstPartyIdentity", "SystemAssignedIdentity" and more.
func (o LookupDatabaseAccountResultOutput) DefaultIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.DefaultIdentity }).(pulumi.StringPtrOutput)
}

// Disable write operations on metadata resources (databases, containers, throughput) via account keys
func (o LookupDatabaseAccountResultOutput) DisableKeyBasedMetadataWriteAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.DisableKeyBasedMetadataWriteAccess }).(pulumi.BoolPtrOutput)
}

// Opt-out of local authentication and ensure only MSI and AAD can be used exclusively for authentication.
func (o LookupDatabaseAccountResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// The connection endpoint for the Cosmos DB database account.
func (o LookupDatabaseAccountResultOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

// Flag to indicate whether to enable storage analytics.
func (o LookupDatabaseAccountResultOutput) EnableAnalyticalStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableAnalyticalStorage }).(pulumi.BoolPtrOutput)
}

// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
func (o LookupDatabaseAccountResultOutput) EnableAutomaticFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableAutomaticFailover }).(pulumi.BoolPtrOutput)
}

// Enables the cassandra connector on the Cosmos DB C* account
func (o LookupDatabaseAccountResultOutput) EnableCassandraConnector() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableCassandraConnector }).(pulumi.BoolPtrOutput)
}

// Flag to indicate whether Free Tier is enabled.
func (o LookupDatabaseAccountResultOutput) EnableFreeTier() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableFreeTier }).(pulumi.BoolPtrOutput)
}

// Enables the account to write in multiple locations
func (o LookupDatabaseAccountResultOutput) EnableMultipleWriteLocations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnableMultipleWriteLocations }).(pulumi.BoolPtrOutput)
}

// Flag to indicate enabling/disabling of Partition Merge feature on the account
func (o LookupDatabaseAccountResultOutput) EnablePartitionMerge() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.EnablePartitionMerge }).(pulumi.BoolPtrOutput)
}

// An array that contains the regions ordered by their failover priorities.
func (o LookupDatabaseAccountResultOutput) FailoverPolicies() FailoverPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []FailoverPolicyResponse { return v.FailoverPolicies }).(FailoverPolicyResponseArrayOutput)
}

// The unique resource identifier of the ARM resource.
func (o LookupDatabaseAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupDatabaseAccountResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// A unique identifier assigned to the database account
func (o LookupDatabaseAccountResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// List of IpRules.
func (o LookupDatabaseAccountResultOutput) IpRules() IpAddressOrRangeResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []IpAddressOrRangeResponse { return v.IpRules }).(IpAddressOrRangeResponseArrayOutput)
}

// Flag to indicate whether to enable/disable Virtual Network ACL rules.
func (o LookupDatabaseAccountResultOutput) IsVirtualNetworkFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *bool { return v.IsVirtualNetworkFilterEnabled }).(pulumi.BoolPtrOutput)
}

// The URI of the key vault
func (o LookupDatabaseAccountResultOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

// The object that represents the metadata for the Account Keys of the Cosmos DB account.
func (o LookupDatabaseAccountResultOutput) KeysMetadata() DatabaseAccountKeysMetadataResponseOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) DatabaseAccountKeysMetadataResponse { return v.KeysMetadata }).(DatabaseAccountKeysMetadataResponseOutput)
}

// Indicates the type of database account. This can only be set at database account creation.
func (o LookupDatabaseAccountResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The location of the resource group to which the resource belongs.
func (o LookupDatabaseAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// An array that contains all of the locations enabled for the Cosmos DB account.
func (o LookupDatabaseAccountResultOutput) Locations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.Locations }).(LocationResponseArrayOutput)
}

// Indicates the minimum allowed Tls version. The default value is Tls 1.2. Cassandra and Mongo APIs only work with Tls 1.2.
func (o LookupDatabaseAccountResultOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the ARM resource.
func (o LookupDatabaseAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates what services are allowed to bypass firewall checks.
func (o LookupDatabaseAccountResultOutput) NetworkAclBypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.NetworkAclBypass }).(pulumi.StringPtrOutput)
}

// An array that contains the Resource Ids for Network Acl Bypass for the Cosmos DB account.
func (o LookupDatabaseAccountResultOutput) NetworkAclBypassResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []string { return v.NetworkAclBypassResourceIds }).(pulumi.StringArrayOutput)
}

// List of Private Endpoint Connections configured for the Cosmos DB account.
func (o LookupDatabaseAccountResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
func (o LookupDatabaseAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether requests from Public Network are allowed
func (o LookupDatabaseAccountResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// An array that contains of the read locations enabled for the Cosmos DB account.
func (o LookupDatabaseAccountResultOutput) ReadLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.ReadLocations }).(LocationResponseArrayOutput)
}

// Parameters to indicate the information about the restore.
func (o LookupDatabaseAccountResultOutput) RestoreParameters() RestoreParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) *RestoreParametersResponse { return v.RestoreParameters }).(RestoreParametersResponsePtrOutput)
}

// The system meta data relating to this resource.
func (o LookupDatabaseAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
func (o LookupDatabaseAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of Azure resource.
func (o LookupDatabaseAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

// List of Virtual Network ACL rules configured for the Cosmos DB account.
func (o LookupDatabaseAccountResultOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

// An array that contains the write location for the Cosmos DB account.
func (o LookupDatabaseAccountResultOutput) WriteLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseAccountResult) []LocationResponse { return v.WriteLocations }).(LocationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseAccountResultOutput{})
}

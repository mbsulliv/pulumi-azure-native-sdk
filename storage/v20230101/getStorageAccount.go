// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the properties for the specified storage account including but not limited to name, SKU name, location, and account status. The ListKeys operation should be used to retrieve storage keys.
func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:storage/v20230101:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupStorageAccountArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// May be used to expand the properties within account's properties. By default, data is not included when fetching properties. Currently we only support geoReplicationStats and blobRestoreStatus.
	Expand *string `pulumi:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The storage account.
type LookupStorageAccountResult struct {
	// Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium' access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium block blobs storage account type.
	AccessTier string `pulumi:"accessTier"`
	// If customer initiated account migration is in progress, the value will be true else it will be null.
	AccountMigrationInProgress bool `pulumi:"accountMigrationInProgress"`
	// Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is false for this property.
	AllowBlobPublicAccess *bool `pulumi:"allowBlobPublicAccess"`
	// Allow or disallow cross AAD tenant object replication. Set this property to true for new or existing accounts only if object replication policies will involve storage accounts in different AAD tenants. The default interpretation is false for new accounts to follow best security practices by default.
	AllowCrossTenantReplication *bool `pulumi:"allowCrossTenantReplication"`
	// Indicates whether the storage account permits requests to be authorized with the account access key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess *bool `pulumi:"allowSharedKeyAccess"`
	// Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
	AllowedCopyScope *string `pulumi:"allowedCopyScope"`
	// Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthenticationResponse `pulumi:"azureFilesIdentityBasedAuthentication"`
	// Blob restore status
	BlobRestoreStatus BlobRestoreStatusResponse `pulumi:"blobRestoreStatus"`
	// Gets the creation date and time of the storage account in UTC.
	CreationTime string `pulumi:"creationTime"`
	// Gets the custom domain the user assigned to this storage account.
	CustomDomain CustomDomainResponse `pulumi:"customDomain"`
	// A boolean flag which indicates whether the default authentication is OAuth or not. The default interpretation is false for this property.
	DefaultToOAuthAuthentication *bool `pulumi:"defaultToOAuthAuthentication"`
	// Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an alphanumeric DNS Zone identifier.
	DnsEndpointType *string `pulumi:"dnsEndpointType"`
	// Allows https traffic only to storage service if sets to true.
	EnableHttpsTrafficOnly *bool `pulumi:"enableHttpsTrafficOnly"`
	// NFS 3.0 protocol support enabled if set to true.
	EnableNfsV3 *bool `pulumi:"enableNfsV3"`
	// Encryption settings to be used for server-side encryption for the storage account.
	Encryption EncryptionResponse `pulumi:"encryption"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// If the failover is in progress, the value will be true, otherwise, it will be null.
	FailoverInProgress bool `pulumi:"failoverInProgress"`
	// Geo Replication Stats
	GeoReplicationStats GeoReplicationStatsResponse `pulumi:"geoReplicationStats"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The property is immutable and can only be set to true at the account creation time. When set to true, it enables object level immutability for all the containers in the account by default.
	ImmutableStorageWithVersioning *ImmutableStorageAccountResponse `pulumi:"immutableStorageWithVersioning"`
	// Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled *bool `pulumi:"isHnsEnabled"`
	// Enables local users feature, if set to true
	IsLocalUserEnabled *bool `pulumi:"isLocalUserEnabled"`
	// Enables Secure File Transfer Protocol, if set to true
	IsSftpEnabled *bool `pulumi:"isSftpEnabled"`
	// This property will be set to true or false on an event of ongoing migration. Default value is null.
	IsSkuConversionBlocked bool `pulumi:"isSkuConversionBlocked"`
	// Storage account keys creation time.
	KeyCreationTime KeyCreationTimeResponse `pulumi:"keyCreationTime"`
	// KeyPolicy assigned to the storage account.
	KeyPolicy KeyPolicyResponse `pulumi:"keyPolicy"`
	// Gets the Kind.
	Kind string `pulumi:"kind"`
	// Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState *string `pulumi:"largeFileSharesState"`
	// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime string `pulumi:"lastGeoFailoverTime"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Network rule set
	NetworkRuleSet NetworkRuleSetResponse `pulumi:"networkRuleSet"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints EndpointsResponse `pulumi:"primaryEndpoints"`
	// Gets the location of the primary data center for the storage account.
	PrimaryLocation string `pulumi:"primaryLocation"`
	// List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Gets the status of the storage account at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// Allow or disallow public network access to Storage Account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Maintains information about the network routing choice opted by the user for data transfer
	RoutingPreference *RoutingPreferenceResponse `pulumi:"routingPreference"`
	// SasPolicy assigned to the storage account.
	SasPolicy SasPolicyResponse `pulumi:"sasPolicy"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints EndpointsResponse `pulumi:"secondaryEndpoints"`
	// Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation string `pulumi:"secondaryLocation"`
	// Gets the SKU.
	Sku SkuResponse `pulumi:"sku"`
	// Gets the status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary string `pulumi:"statusOfPrimary"`
	// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
	StatusOfSecondary string `pulumi:"statusOfSecondary"`
	// This property is readOnly and is set by server during asynchronous storage account sku conversion operations.
	StorageAccountSkuConversionStatus *StorageAccountSkuConversionStatusResponse `pulumi:"storageAccountSkuConversionStatus"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupStorageAccountResult
func (val *LookupStorageAccountResult) Defaults() *LookupStorageAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = *tmp.Encryption.Defaults()

	tmp.NetworkRuleSet = *tmp.NetworkRuleSet.Defaults()

	tmp.SasPolicy = *tmp.SasPolicy.Defaults()

	return &tmp
}

func LookupStorageAccountOutput(ctx *pulumi.Context, args LookupStorageAccountOutputArgs, opts ...pulumi.InvokeOption) LookupStorageAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageAccountResult, error) {
			args := v.(LookupStorageAccountArgs)
			r, err := LookupStorageAccount(ctx, &args, opts...)
			var s LookupStorageAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageAccountResultOutput)
}

type LookupStorageAccountOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// May be used to expand the properties within account's properties. By default, data is not included when fetching properties. Currently we only support geoReplicationStats and blobRestoreStatus.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStorageAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountArgs)(nil)).Elem()
}

// The storage account.
type LookupStorageAccountResultOutput struct{ *pulumi.OutputState }

func (LookupStorageAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountResult)(nil)).Elem()
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutput() LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutputWithContext(ctx context.Context) LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStorageAccountResult] {
	return pulumix.Output[LookupStorageAccountResult]{
		OutputState: o.OutputState,
	}
}

// Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium' access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium block blobs storage account type.
func (o LookupStorageAccountResultOutput) AccessTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.AccessTier }).(pulumi.StringOutput)
}

// If customer initiated account migration is in progress, the value will be true else it will be null.
func (o LookupStorageAccountResultOutput) AccountMigrationInProgress() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) bool { return v.AccountMigrationInProgress }).(pulumi.BoolOutput)
}

// Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is false for this property.
func (o LookupStorageAccountResultOutput) AllowBlobPublicAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.AllowBlobPublicAccess }).(pulumi.BoolPtrOutput)
}

// Allow or disallow cross AAD tenant object replication. Set this property to true for new or existing accounts only if object replication policies will involve storage accounts in different AAD tenants. The default interpretation is false for new accounts to follow best security practices by default.
func (o LookupStorageAccountResultOutput) AllowCrossTenantReplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.AllowCrossTenantReplication }).(pulumi.BoolPtrOutput)
}

// Indicates whether the storage account permits requests to be authorized with the account access key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure Active Directory (Azure AD). The default value is null, which is equivalent to true.
func (o LookupStorageAccountResultOutput) AllowSharedKeyAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.AllowSharedKeyAccess }).(pulumi.BoolPtrOutput)
}

// Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
func (o LookupStorageAccountResultOutput) AllowedCopyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.AllowedCopyScope }).(pulumi.StringPtrOutput)
}

// Provides the identity based authentication settings for Azure Files.
func (o LookupStorageAccountResultOutput) AzureFilesIdentityBasedAuthentication() AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *AzureFilesIdentityBasedAuthenticationResponse {
		return v.AzureFilesIdentityBasedAuthentication
	}).(AzureFilesIdentityBasedAuthenticationResponsePtrOutput)
}

// Blob restore status
func (o LookupStorageAccountResultOutput) BlobRestoreStatus() BlobRestoreStatusResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) BlobRestoreStatusResponse { return v.BlobRestoreStatus }).(BlobRestoreStatusResponseOutput)
}

// Gets the creation date and time of the storage account in UTC.
func (o LookupStorageAccountResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets the custom domain the user assigned to this storage account.
func (o LookupStorageAccountResultOutput) CustomDomain() CustomDomainResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) CustomDomainResponse { return v.CustomDomain }).(CustomDomainResponseOutput)
}

// A boolean flag which indicates whether the default authentication is OAuth or not. The default interpretation is false for this property.
func (o LookupStorageAccountResultOutput) DefaultToOAuthAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.DefaultToOAuthAuthentication }).(pulumi.BoolPtrOutput)
}

// Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an alphanumeric DNS Zone identifier.
func (o LookupStorageAccountResultOutput) DnsEndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.DnsEndpointType }).(pulumi.StringPtrOutput)
}

// Allows https traffic only to storage service if sets to true.
func (o LookupStorageAccountResultOutput) EnableHttpsTrafficOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.EnableHttpsTrafficOnly }).(pulumi.BoolPtrOutput)
}

// NFS 3.0 protocol support enabled if set to true.
func (o LookupStorageAccountResultOutput) EnableNfsV3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.EnableNfsV3 }).(pulumi.BoolPtrOutput)
}

// Encryption settings to be used for server-side encryption for the storage account.
func (o LookupStorageAccountResultOutput) Encryption() EncryptionResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) EncryptionResponse { return v.Encryption }).(EncryptionResponseOutput)
}

// The extendedLocation of the resource.
func (o LookupStorageAccountResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// If the failover is in progress, the value will be true, otherwise, it will be null.
func (o LookupStorageAccountResultOutput) FailoverInProgress() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) bool { return v.FailoverInProgress }).(pulumi.BoolOutput)
}

// Geo Replication Stats
func (o LookupStorageAccountResultOutput) GeoReplicationStats() GeoReplicationStatsResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) GeoReplicationStatsResponse { return v.GeoReplicationStats }).(GeoReplicationStatsResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupStorageAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupStorageAccountResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The property is immutable and can only be set to true at the account creation time. When set to true, it enables object level immutability for all the containers in the account by default.
func (o LookupStorageAccountResultOutput) ImmutableStorageWithVersioning() ImmutableStorageAccountResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *ImmutableStorageAccountResponse {
		return v.ImmutableStorageWithVersioning
	}).(ImmutableStorageAccountResponsePtrOutput)
}

// Account HierarchicalNamespace enabled if sets to true.
func (o LookupStorageAccountResultOutput) IsHnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.IsHnsEnabled }).(pulumi.BoolPtrOutput)
}

// Enables local users feature, if set to true
func (o LookupStorageAccountResultOutput) IsLocalUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.IsLocalUserEnabled }).(pulumi.BoolPtrOutput)
}

// Enables Secure File Transfer Protocol, if set to true
func (o LookupStorageAccountResultOutput) IsSftpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.IsSftpEnabled }).(pulumi.BoolPtrOutput)
}

// This property will be set to true or false on an event of ongoing migration. Default value is null.
func (o LookupStorageAccountResultOutput) IsSkuConversionBlocked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) bool { return v.IsSkuConversionBlocked }).(pulumi.BoolOutput)
}

// Storage account keys creation time.
func (o LookupStorageAccountResultOutput) KeyCreationTime() KeyCreationTimeResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) KeyCreationTimeResponse { return v.KeyCreationTime }).(KeyCreationTimeResponseOutput)
}

// KeyPolicy assigned to the storage account.
func (o LookupStorageAccountResultOutput) KeyPolicy() KeyPolicyResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) KeyPolicyResponse { return v.KeyPolicy }).(KeyPolicyResponseOutput)
}

// Gets the Kind.
func (o LookupStorageAccountResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
func (o LookupStorageAccountResultOutput) LargeFileSharesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.LargeFileSharesState }).(pulumi.StringPtrOutput)
}

// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
func (o LookupStorageAccountResultOutput) LastGeoFailoverTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.LastGeoFailoverTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupStorageAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

// Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
func (o LookupStorageAccountResultOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupStorageAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network rule set
func (o LookupStorageAccountResultOutput) NetworkRuleSet() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) NetworkRuleSetResponse { return v.NetworkRuleSet }).(NetworkRuleSetResponseOutput)
}

// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
func (o LookupStorageAccountResultOutput) PrimaryEndpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) EndpointsResponse { return v.PrimaryEndpoints }).(EndpointsResponseOutput)
}

// Gets the location of the primary data center for the storage account.
func (o LookupStorageAccountResultOutput) PrimaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.PrimaryLocation }).(pulumi.StringOutput)
}

// List of private endpoint connection associated with the specified storage account
func (o LookupStorageAccountResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Gets the status of the storage account at the time the operation was called.
func (o LookupStorageAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Allow or disallow public network access to Storage Account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
func (o LookupStorageAccountResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Maintains information about the network routing choice opted by the user for data transfer
func (o LookupStorageAccountResultOutput) RoutingPreference() RoutingPreferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *RoutingPreferenceResponse { return v.RoutingPreference }).(RoutingPreferenceResponsePtrOutput)
}

// SasPolicy assigned to the storage account.
func (o LookupStorageAccountResultOutput) SasPolicy() SasPolicyResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) SasPolicyResponse { return v.SasPolicy }).(SasPolicyResponseOutput)
}

// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
func (o LookupStorageAccountResultOutput) SecondaryEndpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) EndpointsResponse { return v.SecondaryEndpoints }).(EndpointsResponseOutput)
}

// Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
func (o LookupStorageAccountResultOutput) SecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.SecondaryLocation }).(pulumi.StringOutput)
}

// Gets the SKU.
func (o LookupStorageAccountResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Gets the status indicating whether the primary location of the storage account is available or unavailable.
func (o LookupStorageAccountResultOutput) StatusOfPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.StatusOfPrimary }).(pulumi.StringOutput)
}

// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
func (o LookupStorageAccountResultOutput) StatusOfSecondary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.StatusOfSecondary }).(pulumi.StringOutput)
}

// This property is readOnly and is set by server during asynchronous storage account sku conversion operations.
func (o LookupStorageAccountResultOutput) StorageAccountSkuConversionStatus() StorageAccountSkuConversionStatusResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *StorageAccountSkuConversionStatusResponse {
		return v.StorageAccountSkuConversionStatus
	}).(StorageAccountSkuConversionStatusResponsePtrOutput)
}

// Resource tags.
func (o LookupStorageAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStorageAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountResultOutput{})
}

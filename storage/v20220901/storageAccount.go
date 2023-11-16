// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The storage account.
type StorageAccount struct {
	pulumi.CustomResourceState

	// Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium' access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium block blobs storage account type.
	AccessTier pulumi.StringOutput `pulumi:"accessTier"`
	// Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is true for this property.
	AllowBlobPublicAccess pulumi.BoolPtrOutput `pulumi:"allowBlobPublicAccess"`
	// Allow or disallow cross AAD tenant object replication. The default interpretation is true for this property.
	AllowCrossTenantReplication pulumi.BoolPtrOutput `pulumi:"allowCrossTenantReplication"`
	// Indicates whether the storage account permits requests to be authorized with the account access key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess pulumi.BoolPtrOutput `pulumi:"allowSharedKeyAccess"`
	// Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
	AllowedCopyScope pulumi.StringPtrOutput `pulumi:"allowedCopyScope"`
	// Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication AzureFilesIdentityBasedAuthenticationResponsePtrOutput `pulumi:"azureFilesIdentityBasedAuthentication"`
	// Blob restore status
	BlobRestoreStatus BlobRestoreStatusResponseOutput `pulumi:"blobRestoreStatus"`
	// Gets the creation date and time of the storage account in UTC.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets the custom domain the user assigned to this storage account.
	CustomDomain CustomDomainResponseOutput `pulumi:"customDomain"`
	// A boolean flag which indicates whether the default authentication is OAuth or not. The default interpretation is false for this property.
	DefaultToOAuthAuthentication pulumi.BoolPtrOutput `pulumi:"defaultToOAuthAuthentication"`
	// Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an alphanumeric DNS Zone identifier.
	DnsEndpointType pulumi.StringPtrOutput `pulumi:"dnsEndpointType"`
	// Allows https traffic only to storage service if sets to true.
	EnableHttpsTrafficOnly pulumi.BoolPtrOutput `pulumi:"enableHttpsTrafficOnly"`
	// NFS 3.0 protocol support enabled if set to true.
	EnableNfsV3 pulumi.BoolPtrOutput `pulumi:"enableNfsV3"`
	// Encryption settings to be used for server-side encryption for the storage account.
	Encryption EncryptionResponseOutput `pulumi:"encryption"`
	// The extendedLocation of the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// If the failover is in progress, the value will be true, otherwise, it will be null.
	FailoverInProgress pulumi.BoolOutput `pulumi:"failoverInProgress"`
	// Geo Replication Stats
	GeoReplicationStats GeoReplicationStatsResponseOutput `pulumi:"geoReplicationStats"`
	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The property is immutable and can only be set to true at the account creation time. When set to true, it enables object level immutability for all the containers in the account by default.
	ImmutableStorageWithVersioning ImmutableStorageAccountResponsePtrOutput `pulumi:"immutableStorageWithVersioning"`
	// Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled pulumi.BoolPtrOutput `pulumi:"isHnsEnabled"`
	// Enables local users feature, if set to true
	IsLocalUserEnabled pulumi.BoolPtrOutput `pulumi:"isLocalUserEnabled"`
	// Enables Secure File Transfer Protocol, if set to true
	IsSftpEnabled pulumi.BoolPtrOutput `pulumi:"isSftpEnabled"`
	// Storage account keys creation time.
	KeyCreationTime KeyCreationTimeResponseOutput `pulumi:"keyCreationTime"`
	// KeyPolicy assigned to the storage account.
	KeyPolicy KeyPolicyResponseOutput `pulumi:"keyPolicy"`
	// Gets the Kind.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState pulumi.StringPtrOutput `pulumi:"largeFileSharesState"`
	// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime pulumi.StringOutput `pulumi:"lastGeoFailoverTime"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Network rule set
	NetworkRuleSet NetworkRuleSetResponseOutput `pulumi:"networkRuleSet"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints EndpointsResponseOutput `pulumi:"primaryEndpoints"`
	// Gets the location of the primary data center for the storage account.
	PrimaryLocation pulumi.StringOutput `pulumi:"primaryLocation"`
	// List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Gets the status of the storage account at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Allow or disallow public network access to Storage Account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Maintains information about the network routing choice opted by the user for data transfer
	RoutingPreference RoutingPreferenceResponsePtrOutput `pulumi:"routingPreference"`
	// SasPolicy assigned to the storage account.
	SasPolicy SasPolicyResponseOutput `pulumi:"sasPolicy"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints EndpointsResponseOutput `pulumi:"secondaryEndpoints"`
	// Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation pulumi.StringOutput `pulumi:"secondaryLocation"`
	// Gets the SKU.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Gets the status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary pulumi.StringOutput `pulumi:"statusOfPrimary"`
	// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
	StatusOfSecondary pulumi.StringOutput `pulumi:"statusOfSecondary"`
	// This property is readOnly and is set by server during asynchronous storage account sku conversion operations.
	StorageAccountSkuConversionStatus StorageAccountSkuConversionStatusResponsePtrOutput `pulumi:"storageAccountSkuConversionStatus"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageAccount registers a new resource with the given unique name, arguments, and options.
func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.Encryption != nil {
		args.Encryption = args.Encryption.ToEncryptionPtrOutput().ApplyT(func(v *Encryption) *Encryption { return v.Defaults() }).(EncryptionPtrOutput)
	}
	if args.NetworkRuleSet != nil {
		args.NetworkRuleSet = args.NetworkRuleSet.ToNetworkRuleSetPtrOutput().ApplyT(func(v *NetworkRuleSet) *NetworkRuleSet { return v.Defaults() }).(NetworkRuleSetPtrOutput)
	}
	if args.SasPolicy != nil {
		args.SasPolicy = args.SasPolicy.ToSasPolicyPtrOutput().ApplyT(func(v *SasPolicy) *SasPolicy { return v.Defaults() }).(SasPolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20150501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20150615:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20160101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20160501:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20161201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20170601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20171001:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180301preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180701:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230101:StorageAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StorageAccount
	err := ctx.RegisterResource("azure-native:storage/v20220901:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageAccount gets an existing StorageAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azure-native:storage/v20220901:StorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageAccount resources.
type storageAccountState struct {
}

type StorageAccountState struct {
}

func (StorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountState)(nil)).Elem()
}

type storageAccountArgs struct {
	// Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium' access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium block blobs storage account type.
	AccessTier *AccessTier `pulumi:"accessTier"`
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName *string `pulumi:"accountName"`
	// Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is true for this property.
	AllowBlobPublicAccess *bool `pulumi:"allowBlobPublicAccess"`
	// Allow or disallow cross AAD tenant object replication. The default interpretation is true for this property.
	AllowCrossTenantReplication *bool `pulumi:"allowCrossTenantReplication"`
	// Indicates whether the storage account permits requests to be authorized with the account access key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess *bool `pulumi:"allowSharedKeyAccess"`
	// Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
	AllowedCopyScope *string `pulumi:"allowedCopyScope"`
	// Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication `pulumi:"azureFilesIdentityBasedAuthentication"`
	// User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
	CustomDomain *CustomDomain `pulumi:"customDomain"`
	// A boolean flag which indicates whether the default authentication is OAuth or not. The default interpretation is false for this property.
	DefaultToOAuthAuthentication *bool `pulumi:"defaultToOAuthAuthentication"`
	// Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an alphanumeric DNS Zone identifier.
	DnsEndpointType *string `pulumi:"dnsEndpointType"`
	// Allows https traffic only to storage service if sets to true. The default value is true since API version 2019-04-01.
	EnableHttpsTrafficOnly *bool `pulumi:"enableHttpsTrafficOnly"`
	// NFS 3.0 protocol support enabled if set to true.
	EnableNfsV3 *bool `pulumi:"enableNfsV3"`
	// Encryption settings to be used for server-side encryption for the storage account.
	Encryption *Encryption `pulumi:"encryption"`
	// Optional. Set the extended location of the resource. If not set, the storage account will be created in Azure main region. Otherwise it will be created in the specified extended location
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// The property is immutable and can only be set to true at the account creation time. When set to true, it enables object level immutability for all the new containers in the account by default.
	ImmutableStorageWithVersioning *ImmutableStorageAccount `pulumi:"immutableStorageWithVersioning"`
	// Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled *bool `pulumi:"isHnsEnabled"`
	// Enables local users feature, if set to true
	IsLocalUserEnabled *bool `pulumi:"isLocalUserEnabled"`
	// Enables Secure File Transfer Protocol, if set to true
	IsSftpEnabled *bool `pulumi:"isSftpEnabled"`
	// KeyPolicy assigned to the storage account.
	KeyPolicy *KeyPolicy `pulumi:"keyPolicy"`
	// Required. Indicates the type of storage account.
	Kind string `pulumi:"kind"`
	// Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState *string `pulumi:"largeFileSharesState"`
	// Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
	Location *string `pulumi:"location"`
	// Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// Network rule set
	NetworkRuleSet *NetworkRuleSet `pulumi:"networkRuleSet"`
	// Allow or disallow public network access to Storage Account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Maintains information about the network routing choice opted by the user for data transfer
	RoutingPreference *RoutingPreference `pulumi:"routingPreference"`
	// SasPolicy assigned to the storage account.
	SasPolicy *SasPolicy `pulumi:"sasPolicy"`
	// Required. Gets or sets the SKU name.
	Sku Sku `pulumi:"sku"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageAccount resource.
type StorageAccountArgs struct {
	// Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium' access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium block blobs storage account type.
	AccessTier AccessTierPtrInput
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringPtrInput
	// Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is true for this property.
	AllowBlobPublicAccess pulumi.BoolPtrInput
	// Allow or disallow cross AAD tenant object replication. The default interpretation is true for this property.
	AllowCrossTenantReplication pulumi.BoolPtrInput
	// Indicates whether the storage account permits requests to be authorized with the account access key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess pulumi.BoolPtrInput
	// Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
	AllowedCopyScope pulumi.StringPtrInput
	// Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication AzureFilesIdentityBasedAuthenticationPtrInput
	// User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
	CustomDomain CustomDomainPtrInput
	// A boolean flag which indicates whether the default authentication is OAuth or not. The default interpretation is false for this property.
	DefaultToOAuthAuthentication pulumi.BoolPtrInput
	// Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an alphanumeric DNS Zone identifier.
	DnsEndpointType pulumi.StringPtrInput
	// Allows https traffic only to storage service if sets to true. The default value is true since API version 2019-04-01.
	EnableHttpsTrafficOnly pulumi.BoolPtrInput
	// NFS 3.0 protocol support enabled if set to true.
	EnableNfsV3 pulumi.BoolPtrInput
	// Encryption settings to be used for server-side encryption for the storage account.
	Encryption EncryptionPtrInput
	// Optional. Set the extended location of the resource. If not set, the storage account will be created in Azure main region. Otherwise it will be created in the specified extended location
	ExtendedLocation ExtendedLocationPtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// The property is immutable and can only be set to true at the account creation time. When set to true, it enables object level immutability for all the new containers in the account by default.
	ImmutableStorageWithVersioning ImmutableStorageAccountPtrInput
	// Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled pulumi.BoolPtrInput
	// Enables local users feature, if set to true
	IsLocalUserEnabled pulumi.BoolPtrInput
	// Enables Secure File Transfer Protocol, if set to true
	IsSftpEnabled pulumi.BoolPtrInput
	// KeyPolicy assigned to the storage account.
	KeyPolicy KeyPolicyPtrInput
	// Required. Indicates the type of storage account.
	Kind pulumi.StringInput
	// Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState pulumi.StringPtrInput
	// Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
	Location pulumi.StringPtrInput
	// Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
	MinimumTlsVersion pulumi.StringPtrInput
	// Network rule set
	NetworkRuleSet NetworkRuleSetPtrInput
	// Allow or disallow public network access to Storage Account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Maintains information about the network routing choice opted by the user for data transfer
	RoutingPreference RoutingPreferencePtrInput
	// SasPolicy assigned to the storage account.
	SasPolicy SasPolicyPtrInput
	// Required. Gets or sets the SKU name.
	Sku SkuInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountArgs)(nil)).Elem()
}

type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput
}

func (*StorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *StorageAccount) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i *StorageAccount) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

// Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium' access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium block blobs storage account type.
func (o StorageAccountOutput) AccessTier() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.AccessTier }).(pulumi.StringOutput)
}

// Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is true for this property.
func (o StorageAccountOutput) AllowBlobPublicAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.AllowBlobPublicAccess }).(pulumi.BoolPtrOutput)
}

// Allow or disallow cross AAD tenant object replication. The default interpretation is true for this property.
func (o StorageAccountOutput) AllowCrossTenantReplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.AllowCrossTenantReplication }).(pulumi.BoolPtrOutput)
}

// Indicates whether the storage account permits requests to be authorized with the account access key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure Active Directory (Azure AD). The default value is null, which is equivalent to true.
func (o StorageAccountOutput) AllowSharedKeyAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.AllowSharedKeyAccess }).(pulumi.BoolPtrOutput)
}

// Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
func (o StorageAccountOutput) AllowedCopyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.AllowedCopyScope }).(pulumi.StringPtrOutput)
}

// Provides the identity based authentication settings for Azure Files.
func (o StorageAccountOutput) AzureFilesIdentityBasedAuthentication() AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
		return v.AzureFilesIdentityBasedAuthentication
	}).(AzureFilesIdentityBasedAuthenticationResponsePtrOutput)
}

// Blob restore status
func (o StorageAccountOutput) BlobRestoreStatus() BlobRestoreStatusResponseOutput {
	return o.ApplyT(func(v *StorageAccount) BlobRestoreStatusResponseOutput { return v.BlobRestoreStatus }).(BlobRestoreStatusResponseOutput)
}

// Gets the creation date and time of the storage account in UTC.
func (o StorageAccountOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets the custom domain the user assigned to this storage account.
func (o StorageAccountOutput) CustomDomain() CustomDomainResponseOutput {
	return o.ApplyT(func(v *StorageAccount) CustomDomainResponseOutput { return v.CustomDomain }).(CustomDomainResponseOutput)
}

// A boolean flag which indicates whether the default authentication is OAuth or not. The default interpretation is false for this property.
func (o StorageAccountOutput) DefaultToOAuthAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.DefaultToOAuthAuthentication }).(pulumi.BoolPtrOutput)
}

// Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an alphanumeric DNS Zone identifier.
func (o StorageAccountOutput) DnsEndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.DnsEndpointType }).(pulumi.StringPtrOutput)
}

// Allows https traffic only to storage service if sets to true.
func (o StorageAccountOutput) EnableHttpsTrafficOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.EnableHttpsTrafficOnly }).(pulumi.BoolPtrOutput)
}

// NFS 3.0 protocol support enabled if set to true.
func (o StorageAccountOutput) EnableNfsV3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.EnableNfsV3 }).(pulumi.BoolPtrOutput)
}

// Encryption settings to be used for server-side encryption for the storage account.
func (o StorageAccountOutput) Encryption() EncryptionResponseOutput {
	return o.ApplyT(func(v *StorageAccount) EncryptionResponseOutput { return v.Encryption }).(EncryptionResponseOutput)
}

// The extendedLocation of the resource.
func (o StorageAccountOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// If the failover is in progress, the value will be true, otherwise, it will be null.
func (o StorageAccountOutput) FailoverInProgress() pulumi.BoolOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolOutput { return v.FailoverInProgress }).(pulumi.BoolOutput)
}

// Geo Replication Stats
func (o StorageAccountOutput) GeoReplicationStats() GeoReplicationStatsResponseOutput {
	return o.ApplyT(func(v *StorageAccount) GeoReplicationStatsResponseOutput { return v.GeoReplicationStats }).(GeoReplicationStatsResponseOutput)
}

// The identity of the resource.
func (o StorageAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The property is immutable and can only be set to true at the account creation time. When set to true, it enables object level immutability for all the containers in the account by default.
func (o StorageAccountOutput) ImmutableStorageWithVersioning() ImmutableStorageAccountResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) ImmutableStorageAccountResponsePtrOutput {
		return v.ImmutableStorageWithVersioning
	}).(ImmutableStorageAccountResponsePtrOutput)
}

// Account HierarchicalNamespace enabled if sets to true.
func (o StorageAccountOutput) IsHnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.IsHnsEnabled }).(pulumi.BoolPtrOutput)
}

// Enables local users feature, if set to true
func (o StorageAccountOutput) IsLocalUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.IsLocalUserEnabled }).(pulumi.BoolPtrOutput)
}

// Enables Secure File Transfer Protocol, if set to true
func (o StorageAccountOutput) IsSftpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.IsSftpEnabled }).(pulumi.BoolPtrOutput)
}

// Storage account keys creation time.
func (o StorageAccountOutput) KeyCreationTime() KeyCreationTimeResponseOutput {
	return o.ApplyT(func(v *StorageAccount) KeyCreationTimeResponseOutput { return v.KeyCreationTime }).(KeyCreationTimeResponseOutput)
}

// KeyPolicy assigned to the storage account.
func (o StorageAccountOutput) KeyPolicy() KeyPolicyResponseOutput {
	return o.ApplyT(func(v *StorageAccount) KeyPolicyResponseOutput { return v.KeyPolicy }).(KeyPolicyResponseOutput)
}

// Gets the Kind.
func (o StorageAccountOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
func (o StorageAccountOutput) LargeFileSharesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.LargeFileSharesState }).(pulumi.StringPtrOutput)
}

// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
func (o StorageAccountOutput) LastGeoFailoverTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.LastGeoFailoverTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o StorageAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
func (o StorageAccountOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o StorageAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network rule set
func (o StorageAccountOutput) NetworkRuleSet() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *StorageAccount) NetworkRuleSetResponseOutput { return v.NetworkRuleSet }).(NetworkRuleSetResponseOutput)
}

// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
func (o StorageAccountOutput) PrimaryEndpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v *StorageAccount) EndpointsResponseOutput { return v.PrimaryEndpoints }).(EndpointsResponseOutput)
}

// Gets the location of the primary data center for the storage account.
func (o StorageAccountOutput) PrimaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.PrimaryLocation }).(pulumi.StringOutput)
}

// List of private endpoint connection associated with the specified storage account
func (o StorageAccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *StorageAccount) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Gets the status of the storage account at the time the operation was called.
func (o StorageAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Allow or disallow public network access to Storage Account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
func (o StorageAccountOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Maintains information about the network routing choice opted by the user for data transfer
func (o StorageAccountOutput) RoutingPreference() RoutingPreferenceResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) RoutingPreferenceResponsePtrOutput { return v.RoutingPreference }).(RoutingPreferenceResponsePtrOutput)
}

// SasPolicy assigned to the storage account.
func (o StorageAccountOutput) SasPolicy() SasPolicyResponseOutput {
	return o.ApplyT(func(v *StorageAccount) SasPolicyResponseOutput { return v.SasPolicy }).(SasPolicyResponseOutput)
}

// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
func (o StorageAccountOutput) SecondaryEndpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v *StorageAccount) EndpointsResponseOutput { return v.SecondaryEndpoints }).(EndpointsResponseOutput)
}

// Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
func (o StorageAccountOutput) SecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.SecondaryLocation }).(pulumi.StringOutput)
}

// Gets the SKU.
func (o StorageAccountOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *StorageAccount) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Gets the status indicating whether the primary location of the storage account is available or unavailable.
func (o StorageAccountOutput) StatusOfPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.StatusOfPrimary }).(pulumi.StringOutput)
}

// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
func (o StorageAccountOutput) StatusOfSecondary() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.StatusOfSecondary }).(pulumi.StringOutput)
}

// This property is readOnly and is set by server during asynchronous storage account sku conversion operations.
func (o StorageAccountOutput) StorageAccountSkuConversionStatus() StorageAccountSkuConversionStatusResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) StorageAccountSkuConversionStatusResponsePtrOutput {
		return v.StorageAccountSkuConversionStatus
	}).(StorageAccountSkuConversionStatusResponsePtrOutput)
}

// Resource tags.
func (o StorageAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221229

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a Kusto cluster.
type Cluster struct {
	pulumi.CustomResourceState

	// The cluster's accepted audiences.
	AcceptedAudiences AcceptedAudiencesResponseArrayOutput `pulumi:"acceptedAudiences"`
	// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
	AllowedFqdnList pulumi.StringArrayOutput `pulumi:"allowedFqdnList"`
	// The list of ips in the format of CIDR allowed to connect to the cluster.
	AllowedIpRangeList pulumi.StringArrayOutput `pulumi:"allowedIpRangeList"`
	// The cluster data ingestion URI.
	DataIngestionUri pulumi.StringOutput `pulumi:"dataIngestionUri"`
	// A boolean value that indicates if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	EnableAutoStop pulumi.BoolPtrOutput `pulumi:"enableAutoStop"`
	// A boolean value that indicates if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrOutput `pulumi:"enableDiskEncryption"`
	// A boolean value that indicates if double encryption is enabled.
	EnableDoubleEncryption pulumi.BoolPtrOutput `pulumi:"enableDoubleEncryption"`
	// A boolean value that indicates if the purge operations are enabled.
	EnablePurge pulumi.BoolPtrOutput `pulumi:"enablePurge"`
	// A boolean value that indicates if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrOutput `pulumi:"enableStreamingIngest"`
	// The engine type
	EngineType pulumi.StringPtrOutput `pulumi:"engineType"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The identity of the cluster, if configured.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// KeyVault properties for the cluster encryption.
	KeyVaultProperties KeyVaultPropertiesResponsePtrOutput `pulumi:"keyVaultProperties"`
	// List of the cluster's language extensions.
	LanguageExtensions LanguageExtensionsListResponsePtrOutput `pulumi:"languageExtensions"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Optimized auto scale definition.
	OptimizedAutoscale OptimizedAutoscaleResponsePtrOutput `pulumi:"optimizedAutoscale"`
	// A list of private endpoint connections.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6)
	PublicIPType pulumi.StringPtrOutput `pulumi:"publicIPType"`
	// Public network access to the cluster is enabled by default. When disabled, only private endpoint connection to the cluster is allowed
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Whether or not to restrict outbound network access.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess pulumi.StringPtrOutput `pulumi:"restrictOutboundNetworkAccess"`
	// The SKU of the cluster.
	Sku AzureSkuResponseOutput `pulumi:"sku"`
	// The state of the resource.
	State pulumi.StringOutput `pulumi:"state"`
	// The reason for the cluster's current state.
	StateReason pulumi.StringOutput `pulumi:"stateReason"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The cluster's external tenants.
	TrustedExternalTenants TrustedExternalTenantResponseArrayOutput `pulumi:"trustedExternalTenants"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The cluster URI.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// Virtual network definition.
	VirtualNetworkConfiguration VirtualNetworkConfigurationResponsePtrOutput `pulumi:"virtualNetworkConfiguration"`
	// The availability zones of the cluster.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.EnableAutoStop == nil {
		args.EnableAutoStop = pulumi.BoolPtr(true)
	}
	if args.EnableDiskEncryption == nil {
		args.EnableDiskEncryption = pulumi.BoolPtr(false)
	}
	if args.EnableDoubleEncryption == nil {
		args.EnableDoubleEncryption = pulumi.BoolPtr(false)
	}
	if args.EnablePurge == nil {
		args.EnablePurge = pulumi.BoolPtr(false)
	}
	if args.EnableStreamingIngest == nil {
		args.EnableStreamingIngest = pulumi.BoolPtr(false)
	}
	if args.EngineType == nil {
		args.EngineType = pulumi.StringPtr("V3")
	}
	if args.PublicIPType == nil {
		args.PublicIPType = pulumi.StringPtr("IPv4")
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	if args.RestrictOutboundNetworkAccess == nil {
		args.RestrictOutboundNetworkAccess = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230815:Cluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:kusto/v20221229:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:kusto/v20221229:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The cluster's accepted audiences.
	AcceptedAudiences []AcceptedAudiences `pulumi:"acceptedAudiences"`
	// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
	AllowedFqdnList []string `pulumi:"allowedFqdnList"`
	// The list of ips in the format of CIDR allowed to connect to the cluster.
	AllowedIpRangeList []string `pulumi:"allowedIpRangeList"`
	// The name of the Kusto cluster.
	ClusterName *string `pulumi:"clusterName"`
	// A boolean value that indicates if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	EnableAutoStop *bool `pulumi:"enableAutoStop"`
	// A boolean value that indicates if the cluster's disks are encrypted.
	EnableDiskEncryption *bool `pulumi:"enableDiskEncryption"`
	// A boolean value that indicates if double encryption is enabled.
	EnableDoubleEncryption *bool `pulumi:"enableDoubleEncryption"`
	// A boolean value that indicates if the purge operations are enabled.
	EnablePurge *bool `pulumi:"enablePurge"`
	// A boolean value that indicates if the streaming ingest is enabled.
	EnableStreamingIngest *bool `pulumi:"enableStreamingIngest"`
	// The engine type
	EngineType *string `pulumi:"engineType"`
	// The identity of the cluster, if configured.
	Identity *Identity `pulumi:"identity"`
	// KeyVault properties for the cluster encryption.
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	// List of the cluster's language extensions.
	LanguageExtensions *LanguageExtensionsList `pulumi:"languageExtensions"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Optimized auto scale definition.
	OptimizedAutoscale *OptimizedAutoscale `pulumi:"optimizedAutoscale"`
	// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6)
	PublicIPType *string `pulumi:"publicIPType"`
	// Public network access to the cluster is enabled by default. When disabled, only private endpoint connection to the cluster is allowed
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Whether or not to restrict outbound network access.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess *string `pulumi:"restrictOutboundNetworkAccess"`
	// The SKU of the cluster.
	Sku AzureSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The cluster's external tenants.
	TrustedExternalTenants []TrustedExternalTenant `pulumi:"trustedExternalTenants"`
	// Virtual Cluster graduation properties
	VirtualClusterGraduationProperties *string `pulumi:"virtualClusterGraduationProperties"`
	// Virtual network definition.
	VirtualNetworkConfiguration *VirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// The availability zones of the cluster.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The cluster's accepted audiences.
	AcceptedAudiences AcceptedAudiencesArrayInput
	// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
	AllowedFqdnList pulumi.StringArrayInput
	// The list of ips in the format of CIDR allowed to connect to the cluster.
	AllowedIpRangeList pulumi.StringArrayInput
	// The name of the Kusto cluster.
	ClusterName pulumi.StringPtrInput
	// A boolean value that indicates if the cluster could be automatically stopped (due to lack of data or no activity for many days).
	EnableAutoStop pulumi.BoolPtrInput
	// A boolean value that indicates if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrInput
	// A boolean value that indicates if double encryption is enabled.
	EnableDoubleEncryption pulumi.BoolPtrInput
	// A boolean value that indicates if the purge operations are enabled.
	EnablePurge pulumi.BoolPtrInput
	// A boolean value that indicates if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrInput
	// The engine type
	EngineType pulumi.StringPtrInput
	// The identity of the cluster, if configured.
	Identity IdentityPtrInput
	// KeyVault properties for the cluster encryption.
	KeyVaultProperties KeyVaultPropertiesPtrInput
	// List of the cluster's language extensions.
	LanguageExtensions LanguageExtensionsListPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Optimized auto scale definition.
	OptimizedAutoscale OptimizedAutoscalePtrInput
	// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6)
	PublicIPType pulumi.StringPtrInput
	// Public network access to the cluster is enabled by default. When disabled, only private endpoint connection to the cluster is allowed
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// Whether or not to restrict outbound network access.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess pulumi.StringPtrInput
	// The SKU of the cluster.
	Sku AzureSkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The cluster's external tenants.
	TrustedExternalTenants TrustedExternalTenantArrayInput
	// Virtual Cluster graduation properties
	VirtualClusterGraduationProperties pulumi.StringPtrInput
	// Virtual network definition.
	VirtualNetworkConfiguration VirtualNetworkConfigurationPtrInput
	// The availability zones of the cluster.
	Zones pulumi.StringArrayInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The cluster's accepted audiences.
func (o ClusterOutput) AcceptedAudiences() AcceptedAudiencesResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) AcceptedAudiencesResponseArrayOutput { return v.AcceptedAudiences }).(AcceptedAudiencesResponseArrayOutput)
}

// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
func (o ClusterOutput) AllowedFqdnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.AllowedFqdnList }).(pulumi.StringArrayOutput)
}

// The list of ips in the format of CIDR allowed to connect to the cluster.
func (o ClusterOutput) AllowedIpRangeList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.AllowedIpRangeList }).(pulumi.StringArrayOutput)
}

// The cluster data ingestion URI.
func (o ClusterOutput) DataIngestionUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DataIngestionUri }).(pulumi.StringOutput)
}

// A boolean value that indicates if the cluster could be automatically stopped (due to lack of data or no activity for many days).
func (o ClusterOutput) EnableAutoStop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EnableAutoStop }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates if the cluster's disks are encrypted.
func (o ClusterOutput) EnableDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EnableDiskEncryption }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates if double encryption is enabled.
func (o ClusterOutput) EnableDoubleEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EnableDoubleEncryption }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates if the purge operations are enabled.
func (o ClusterOutput) EnablePurge() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EnablePurge }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates if the streaming ingest is enabled.
func (o ClusterOutput) EnableStreamingIngest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EnableStreamingIngest }).(pulumi.BoolPtrOutput)
}

// The engine type
func (o ClusterOutput) EngineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.EngineType }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ClusterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The identity of the cluster, if configured.
func (o ClusterOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// KeyVault properties for the cluster encryption.
func (o ClusterOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) KeyVaultPropertiesResponsePtrOutput { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

// List of the cluster's language extensions.
func (o ClusterOutput) LanguageExtensions() LanguageExtensionsListResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) LanguageExtensionsListResponsePtrOutput { return v.LanguageExtensions }).(LanguageExtensionsListResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optimized auto scale definition.
func (o ClusterOutput) OptimizedAutoscale() OptimizedAutoscaleResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) OptimizedAutoscaleResponsePtrOutput { return v.OptimizedAutoscale }).(OptimizedAutoscaleResponsePtrOutput)
}

// A list of private endpoint connections.
func (o ClusterOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioned state of the resource.
func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6)
func (o ClusterOutput) PublicIPType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.PublicIPType }).(pulumi.StringPtrOutput)
}

// Public network access to the cluster is enabled by default. When disabled, only private endpoint connection to the cluster is allowed
func (o ClusterOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Whether or not to restrict outbound network access.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
func (o ClusterOutput) RestrictOutboundNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.RestrictOutboundNetworkAccess }).(pulumi.StringPtrOutput)
}

// The SKU of the cluster.
func (o ClusterOutput) Sku() AzureSkuResponseOutput {
	return o.ApplyT(func(v *Cluster) AzureSkuResponseOutput { return v.Sku }).(AzureSkuResponseOutput)
}

// The state of the resource.
func (o ClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The reason for the cluster's current state.
func (o ClusterOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.StateReason }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The cluster's external tenants.
func (o ClusterOutput) TrustedExternalTenants() TrustedExternalTenantResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) TrustedExternalTenantResponseArrayOutput { return v.TrustedExternalTenants }).(TrustedExternalTenantResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The cluster URI.
func (o ClusterOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

// Virtual network definition.
func (o ClusterOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) VirtualNetworkConfigurationResponsePtrOutput { return v.VirtualNetworkConfiguration }).(VirtualNetworkConfigurationResponsePtrOutput)
}

// The availability zones of the cluster.
func (o ClusterOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}

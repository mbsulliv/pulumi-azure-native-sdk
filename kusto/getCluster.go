// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Kusto cluster.
// Azure REST API version: 2022-12-29.
//
// Other available API versions: 2022-07-07, 2023-05-02, 2023-08-15.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:kusto:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupClusterArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a Kusto cluster.
type LookupClusterResult struct {
	// The cluster's accepted audiences.
	AcceptedAudiences []AcceptedAudiencesResponse `pulumi:"acceptedAudiences"`
	// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
	AllowedFqdnList []string `pulumi:"allowedFqdnList"`
	// The list of ips in the format of CIDR allowed to connect to the cluster.
	AllowedIpRangeList []string `pulumi:"allowedIpRangeList"`
	// The cluster data ingestion URI.
	DataIngestionUri string `pulumi:"dataIngestionUri"`
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
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The identity of the cluster, if configured.
	Identity *IdentityResponse `pulumi:"identity"`
	// KeyVault properties for the cluster encryption.
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	// List of the cluster's language extensions.
	LanguageExtensions *LanguageExtensionsListResponse `pulumi:"languageExtensions"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Optimized auto scale definition.
	OptimizedAutoscale *OptimizedAutoscaleResponse `pulumi:"optimizedAutoscale"`
	// A list of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6)
	PublicIPType *string `pulumi:"publicIPType"`
	// Public network access to the cluster is enabled by default. When disabled, only private endpoint connection to the cluster is allowed
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Whether or not to restrict outbound network access.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	RestrictOutboundNetworkAccess *string `pulumi:"restrictOutboundNetworkAccess"`
	// The SKU of the cluster.
	Sku AzureSkuResponse `pulumi:"sku"`
	// The state of the resource.
	State string `pulumi:"state"`
	// The reason for the cluster's current state.
	StateReason string `pulumi:"stateReason"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The cluster's external tenants.
	TrustedExternalTenants []TrustedExternalTenantResponse `pulumi:"trustedExternalTenants"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The cluster URI.
	Uri string `pulumi:"uri"`
	// Virtual network definition.
	VirtualNetworkConfiguration *VirtualNetworkConfigurationResponse `pulumi:"virtualNetworkConfiguration"`
	// The availability zones of the cluster.
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupClusterResult
func (val *LookupClusterResult) Defaults() *LookupClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EnableAutoStop == nil {
		enableAutoStop_ := true
		tmp.EnableAutoStop = &enableAutoStop_
	}
	if tmp.EnableDiskEncryption == nil {
		enableDiskEncryption_ := false
		tmp.EnableDiskEncryption = &enableDiskEncryption_
	}
	if tmp.EnableDoubleEncryption == nil {
		enableDoubleEncryption_ := false
		tmp.EnableDoubleEncryption = &enableDoubleEncryption_
	}
	if tmp.EnablePurge == nil {
		enablePurge_ := false
		tmp.EnablePurge = &enablePurge_
	}
	if tmp.EnableStreamingIngest == nil {
		enableStreamingIngest_ := false
		tmp.EnableStreamingIngest = &enableStreamingIngest_
	}
	if tmp.EngineType == nil {
		engineType_ := "V3"
		tmp.EngineType = &engineType_
	}
	if tmp.PublicIPType == nil {
		publicIPType_ := "IPv4"
		tmp.PublicIPType = &publicIPType_
	}
	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	if tmp.RestrictOutboundNetworkAccess == nil {
		restrictOutboundNetworkAccess_ := "Disabled"
		tmp.RestrictOutboundNetworkAccess = &restrictOutboundNetworkAccess_
	}
	return &tmp
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// Class representing a Kusto cluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupClusterResult] {
	return pulumix.Output[LookupClusterResult]{
		OutputState: o.OutputState,
	}
}

// The cluster's accepted audiences.
func (o LookupClusterResultOutput) AcceptedAudiences() AcceptedAudiencesResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []AcceptedAudiencesResponse { return v.AcceptedAudiences }).(AcceptedAudiencesResponseArrayOutput)
}

// List of allowed FQDNs(Fully Qualified Domain Name) for egress from Cluster.
func (o LookupClusterResultOutput) AllowedFqdnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.AllowedFqdnList }).(pulumi.StringArrayOutput)
}

// The list of ips in the format of CIDR allowed to connect to the cluster.
func (o LookupClusterResultOutput) AllowedIpRangeList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.AllowedIpRangeList }).(pulumi.StringArrayOutput)
}

// The cluster data ingestion URI.
func (o LookupClusterResultOutput) DataIngestionUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DataIngestionUri }).(pulumi.StringOutput)
}

// A boolean value that indicates if the cluster could be automatically stopped (due to lack of data or no activity for many days).
func (o LookupClusterResultOutput) EnableAutoStop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnableAutoStop }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates if the cluster's disks are encrypted.
func (o LookupClusterResultOutput) EnableDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnableDiskEncryption }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates if double encryption is enabled.
func (o LookupClusterResultOutput) EnableDoubleEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnableDoubleEncryption }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates if the purge operations are enabled.
func (o LookupClusterResultOutput) EnablePurge() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnablePurge }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates if the streaming ingest is enabled.
func (o LookupClusterResultOutput) EnableStreamingIngest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.EnableStreamingIngest }).(pulumi.BoolPtrOutput)
}

// The engine type
func (o LookupClusterResultOutput) EngineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.EngineType }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupClusterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the cluster, if configured.
func (o LookupClusterResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// KeyVault properties for the cluster encryption.
func (o LookupClusterResultOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

// List of the cluster's language extensions.
func (o LookupClusterResultOutput) LanguageExtensions() LanguageExtensionsListResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *LanguageExtensionsListResponse { return v.LanguageExtensions }).(LanguageExtensionsListResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optimized auto scale definition.
func (o LookupClusterResultOutput) OptimizedAutoscale() OptimizedAutoscaleResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *OptimizedAutoscaleResponse { return v.OptimizedAutoscale }).(OptimizedAutoscaleResponsePtrOutput)
}

// A list of private endpoint connections.
func (o LookupClusterResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioned state of the resource.
func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Indicates what public IP type to create - IPv4 (default), or DualStack (both IPv4 and IPv6)
func (o LookupClusterResultOutput) PublicIPType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.PublicIPType }).(pulumi.StringPtrOutput)
}

// Public network access to the cluster is enabled by default. When disabled, only private endpoint connection to the cluster is allowed
func (o LookupClusterResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Whether or not to restrict outbound network access.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
func (o LookupClusterResultOutput) RestrictOutboundNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.RestrictOutboundNetworkAccess }).(pulumi.StringPtrOutput)
}

// The SKU of the cluster.
func (o LookupClusterResultOutput) Sku() AzureSkuResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) AzureSkuResponse { return v.Sku }).(AzureSkuResponseOutput)
}

// The state of the resource.
func (o LookupClusterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.State }).(pulumi.StringOutput)
}

// The reason for the cluster's current state.
func (o LookupClusterResultOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.StateReason }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The cluster's external tenants.
func (o LookupClusterResultOutput) TrustedExternalTenants() TrustedExternalTenantResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []TrustedExternalTenantResponse { return v.TrustedExternalTenants }).(TrustedExternalTenantResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// The cluster URI.
func (o LookupClusterResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Uri }).(pulumi.StringOutput)
}

// Virtual network definition.
func (o LookupClusterResultOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *VirtualNetworkConfigurationResponse { return v.VirtualNetworkConfiguration }).(VirtualNetworkConfigurationResponsePtrOutput)
}

// The availability zones of the cluster.
func (o LookupClusterResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}

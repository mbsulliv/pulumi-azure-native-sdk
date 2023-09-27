// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a description for the specified namespace.
// Azure REST API version: 2022-01-01-preview.
func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:servicebus:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNamespaceArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a namespace resource.
type LookupNamespaceResult struct {
	// Alternate name for namespace
	AlternateName *string `pulumi:"alternateName"`
	// The time the namespace was created
	CreatedAt string `pulumi:"createdAt"`
	// This property disables SAS authentication for the Service Bus namespace.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Properties of BYOK Encryption description
	Encryption *EncryptionResponse `pulumi:"encryption"`
	// Resource Id
	Id string `pulumi:"id"`
	// Properties of BYOK Identity description
	Identity *IdentityResponse `pulumi:"identity"`
	// The Geo-location where the resource lives
	Location string `pulumi:"location"`
	// Identifier for Azure Insights metrics
	MetricId string `pulumi:"metricId"`
	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// Resource name
	Name string `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the namespace.
	ProvisioningState string `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint string `pulumi:"serviceBusEndpoint"`
	// Properties of SKU
	Sku *SBSkuResponse `pulumi:"sku"`
	// Status of the namespace.
	Status string `pulumi:"status"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt string `pulumi:"updatedAt"`
	// Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// Defaults sets the appropriate defaults for LookupNamespaceResult
func (val *LookupNamespaceResult) Defaults() *LookupNamespaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = tmp.Encryption.Defaults()

	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}

func LookupNamespaceOutput(ctx *pulumi.Context, args LookupNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceResult, error) {
			args := v.(LookupNamespaceArgs)
			r, err := LookupNamespace(ctx, &args, opts...)
			var s LookupNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceResultOutput)
}

type LookupNamespaceOutputArgs struct {
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceArgs)(nil)).Elem()
}

// Description of a namespace resource.
type LookupNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceResult)(nil)).Elem()
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutput() LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutputWithContext(ctx context.Context) LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNamespaceResult] {
	return pulumix.Output[LookupNamespaceResult]{
		OutputState: o.OutputState,
	}
}

// Alternate name for namespace
func (o LookupNamespaceResultOutput) AlternateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.AlternateName }).(pulumi.StringPtrOutput)
}

// The time the namespace was created
func (o LookupNamespaceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// This property disables SAS authentication for the Service Bus namespace.
func (o LookupNamespaceResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Properties of BYOK Encryption description
func (o LookupNamespaceResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// Resource Id
func (o LookupNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Properties of BYOK Identity description
func (o LookupNamespaceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The Geo-location where the resource lives
func (o LookupNamespaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Identifier for Azure Insights metrics
func (o LookupNamespaceResultOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.MetricId }).(pulumi.StringOutput)
}

// The minimum TLS version for the cluster to support, e.g. '1.2'
func (o LookupNamespaceResultOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections.
func (o LookupNamespaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupNamespaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the namespace.
func (o LookupNamespaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
func (o LookupNamespaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Endpoint you can use to perform Service Bus operations.
func (o LookupNamespaceResultOutput) ServiceBusEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.ServiceBusEndpoint }).(pulumi.StringOutput)
}

// Properties of SKU
func (o LookupNamespaceResultOutput) Sku() SBSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *SBSkuResponse { return v.Sku }).(SBSkuResponsePtrOutput)
}

// Status of the namespace.
func (o LookupNamespaceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Status }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o LookupNamespaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNamespaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The time the namespace was updated.
func (o LookupNamespaceResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
func (o LookupNamespaceResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceResultOutput{})
}

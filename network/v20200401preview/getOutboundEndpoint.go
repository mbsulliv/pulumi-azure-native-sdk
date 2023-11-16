// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of an outbound endpoint for a DNS resolver.
func LookupOutboundEndpoint(ctx *pulumi.Context, args *LookupOutboundEndpointArgs, opts ...pulumi.InvokeOption) (*LookupOutboundEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOutboundEndpointResult
	err := ctx.Invoke("azure-native:network/v20200401preview:getOutboundEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutboundEndpointArgs struct {
	// The name of the DNS resolver.
	DnsResolverName string `pulumi:"dnsResolverName"`
	// The name of the outbound endpoint for the DNS resolver.
	OutboundEndpointName string `pulumi:"outboundEndpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes an outbound endpoint for a DNS resolver.
type LookupOutboundEndpointResult struct {
	// ETag of the outbound endpoint.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The current provisioning state of the outbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resourceGuid property of the outbound endpoint resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The reference to the subnet used for the outbound endpoint.
	Subnet *SubResourceResponse `pulumi:"subnet"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupOutboundEndpointOutput(ctx *pulumi.Context, args LookupOutboundEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupOutboundEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOutboundEndpointResult, error) {
			args := v.(LookupOutboundEndpointArgs)
			r, err := LookupOutboundEndpoint(ctx, &args, opts...)
			var s LookupOutboundEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOutboundEndpointResultOutput)
}

type LookupOutboundEndpointOutputArgs struct {
	// The name of the DNS resolver.
	DnsResolverName pulumi.StringInput `pulumi:"dnsResolverName"`
	// The name of the outbound endpoint for the DNS resolver.
	OutboundEndpointName pulumi.StringInput `pulumi:"outboundEndpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOutboundEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutboundEndpointArgs)(nil)).Elem()
}

// Describes an outbound endpoint for a DNS resolver.
type LookupOutboundEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupOutboundEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutboundEndpointResult)(nil)).Elem()
}

func (o LookupOutboundEndpointResultOutput) ToLookupOutboundEndpointResultOutput() LookupOutboundEndpointResultOutput {
	return o
}

func (o LookupOutboundEndpointResultOutput) ToLookupOutboundEndpointResultOutputWithContext(ctx context.Context) LookupOutboundEndpointResultOutput {
	return o
}

// ETag of the outbound endpoint.
func (o LookupOutboundEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOutboundEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupOutboundEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOutboundEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the outbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
func (o LookupOutboundEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid property of the outbound endpoint resource.
func (o LookupOutboundEndpointResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The reference to the subnet used for the outbound endpoint.
func (o LookupOutboundEndpointResultOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupOutboundEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupOutboundEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOutboundEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOutboundEndpointResultOutput{})
}

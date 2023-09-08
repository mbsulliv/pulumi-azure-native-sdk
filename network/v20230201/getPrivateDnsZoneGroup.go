// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the private dns zone group resource by specified private dns zone group name.
func LookupPrivateDnsZoneGroup(ctx *pulumi.Context, args *LookupPrivateDnsZoneGroupArgs, opts ...pulumi.InvokeOption) (*LookupPrivateDnsZoneGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateDnsZoneGroupResult
	err := ctx.Invoke("azure-native:network/v20230201:getPrivateDnsZoneGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateDnsZoneGroupArgs struct {
	// The name of the private dns zone group.
	PrivateDnsZoneGroupName string `pulumi:"privateDnsZoneGroupName"`
	// The name of the private endpoint.
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Private dns zone group resource.
type LookupPrivateDnsZoneGroupResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// A collection of private dns zone configurations of the private dns zone group.
	PrivateDnsZoneConfigs []PrivateDnsZoneConfigResponse `pulumi:"privateDnsZoneConfigs"`
	// The provisioning state of the private dns zone group resource.
	ProvisioningState string `pulumi:"provisioningState"`
}

func LookupPrivateDnsZoneGroupOutput(ctx *pulumi.Context, args LookupPrivateDnsZoneGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateDnsZoneGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateDnsZoneGroupResult, error) {
			args := v.(LookupPrivateDnsZoneGroupArgs)
			r, err := LookupPrivateDnsZoneGroup(ctx, &args, opts...)
			var s LookupPrivateDnsZoneGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateDnsZoneGroupResultOutput)
}

type LookupPrivateDnsZoneGroupOutputArgs struct {
	// The name of the private dns zone group.
	PrivateDnsZoneGroupName pulumi.StringInput `pulumi:"privateDnsZoneGroupName"`
	// The name of the private endpoint.
	PrivateEndpointName pulumi.StringInput `pulumi:"privateEndpointName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateDnsZoneGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDnsZoneGroupArgs)(nil)).Elem()
}

// Private dns zone group resource.
type LookupPrivateDnsZoneGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateDnsZoneGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDnsZoneGroupResult)(nil)).Elem()
}

func (o LookupPrivateDnsZoneGroupResultOutput) ToLookupPrivateDnsZoneGroupResultOutput() LookupPrivateDnsZoneGroupResultOutput {
	return o
}

func (o LookupPrivateDnsZoneGroupResultOutput) ToLookupPrivateDnsZoneGroupResultOutputWithContext(ctx context.Context) LookupPrivateDnsZoneGroupResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupPrivateDnsZoneGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupPrivateDnsZoneGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupPrivateDnsZoneGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A collection of private dns zone configurations of the private dns zone group.
func (o LookupPrivateDnsZoneGroupResultOutput) PrivateDnsZoneConfigs() PrivateDnsZoneConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) []PrivateDnsZoneConfigResponse { return v.PrivateDnsZoneConfigs }).(PrivateDnsZoneConfigResponseArrayOutput)
}

// The provisioning state of the private dns zone group resource.
func (o LookupPrivateDnsZoneGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateDnsZoneGroupResultOutput{})
}
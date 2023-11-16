// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a VirtualHubRouteTableV2.
func LookupVirtualHubRouteTableV2(ctx *pulumi.Context, args *LookupVirtualHubRouteTableV2Args, opts ...pulumi.InvokeOption) (*LookupVirtualHubRouteTableV2Result, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualHubRouteTableV2Result
	err := ctx.Invoke("azure-native:network/v20230401:getVirtualHubRouteTableV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubRouteTableV2Args struct {
	// The resource group name of the VirtualHubRouteTableV2.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHubRouteTableV2.
	RouteTableName string `pulumi:"routeTableName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// VirtualHubRouteTableV2 Resource.
type LookupVirtualHubRouteTableV2Result struct {
	// List of all connections attached to this route table v2.
	AttachedConnections []string `pulumi:"attachedConnections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the virtual hub route table v2 resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of all routes.
	Routes []VirtualHubRouteV2Response `pulumi:"routes"`
}

func LookupVirtualHubRouteTableV2Output(ctx *pulumi.Context, args LookupVirtualHubRouteTableV2OutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubRouteTableV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubRouteTableV2Result, error) {
			args := v.(LookupVirtualHubRouteTableV2Args)
			r, err := LookupVirtualHubRouteTableV2(ctx, &args, opts...)
			var s LookupVirtualHubRouteTableV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubRouteTableV2ResultOutput)
}

type LookupVirtualHubRouteTableV2OutputArgs struct {
	// The resource group name of the VirtualHubRouteTableV2.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VirtualHubRouteTableV2.
	RouteTableName pulumi.StringInput `pulumi:"routeTableName"`
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubRouteTableV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubRouteTableV2Args)(nil)).Elem()
}

// VirtualHubRouteTableV2 Resource.
type LookupVirtualHubRouteTableV2ResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubRouteTableV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubRouteTableV2Result)(nil)).Elem()
}

func (o LookupVirtualHubRouteTableV2ResultOutput) ToLookupVirtualHubRouteTableV2ResultOutput() LookupVirtualHubRouteTableV2ResultOutput {
	return o
}

func (o LookupVirtualHubRouteTableV2ResultOutput) ToLookupVirtualHubRouteTableV2ResultOutputWithContext(ctx context.Context) LookupVirtualHubRouteTableV2ResultOutput {
	return o
}

// List of all connections attached to this route table v2.
func (o LookupVirtualHubRouteTableV2ResultOutput) AttachedConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) []string { return v.AttachedConnections }).(pulumi.StringArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualHubRouteTableV2ResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVirtualHubRouteTableV2ResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupVirtualHubRouteTableV2ResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the virtual hub route table v2 resource.
func (o LookupVirtualHubRouteTableV2ResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of all routes.
func (o LookupVirtualHubRouteTableV2ResultOutput) Routes() VirtualHubRouteV2ResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) []VirtualHubRouteV2Response { return v.Routes }).(VirtualHubRouteV2ResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubRouteTableV2ResultOutput{})
}

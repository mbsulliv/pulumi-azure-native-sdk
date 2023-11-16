// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified route from a route table.
func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteResult
	err := ctx.Invoke("azure-native:network/v20230501:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route.
	RouteName string `pulumi:"routeName"`
	// The name of the route table.
	RouteTableName string `pulumi:"routeTableName"`
}

// Route resource.
type LookupRouteResult struct {
	// The destination CIDR to which the route applies.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
	HasBgpOverride *bool `pulumi:"hasBgpOverride"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress *string `pulumi:"nextHopIpAddress"`
	// The type of Azure hop the packet should be sent to.
	NextHopType string `pulumi:"nextHopType"`
	// The provisioning state of the route resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

type LookupRouteOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the route.
	RouteName pulumi.StringInput `pulumi:"routeName"`
	// The name of the route table.
	RouteTableName pulumi.StringInput `pulumi:"routeTableName"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}

// Route resource.
type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

// The destination CIDR to which the route applies.
func (o LookupRouteResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupRouteResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Etag }).(pulumi.StringOutput)
}

// A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
func (o LookupRouteResultOutput) HasBgpOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *bool { return v.HasBgpOverride }).(pulumi.BoolPtrOutput)
}

// Resource ID.
func (o LookupRouteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupRouteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
func (o LookupRouteResultOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

// The type of Azure hop the packet should be sent to.
func (o LookupRouteResultOutput) NextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopType }).(pulumi.StringOutput)
}

// The provisioning state of the route resource.
func (o LookupRouteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource.
func (o LookupRouteResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}

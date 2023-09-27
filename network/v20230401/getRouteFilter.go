// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified route filter.
func LookupRouteFilter(ctx *pulumi.Context, args *LookupRouteFilterArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteFilterResult
	err := ctx.Invoke("azure-native:network/v20230401:getRouteFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteFilterArgs struct {
	// Expands referenced express route bgp peering resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName string `pulumi:"routeFilterName"`
}

// Route Filter Resource.
type LookupRouteFilterResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A collection of references to express route circuit ipv6 peerings.
	Ipv6Peerings []ExpressRouteCircuitPeeringResponse `pulumi:"ipv6Peerings"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// A collection of references to express route circuit peerings.
	Peerings []ExpressRouteCircuitPeeringResponse `pulumi:"peerings"`
	// The provisioning state of the route filter resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Collection of RouteFilterRules contained within a route filter.
	Rules []RouteFilterRuleResponse `pulumi:"rules"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupRouteFilterOutput(ctx *pulumi.Context, args LookupRouteFilterOutputArgs, opts ...pulumi.InvokeOption) LookupRouteFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteFilterResult, error) {
			args := v.(LookupRouteFilterArgs)
			r, err := LookupRouteFilter(ctx, &args, opts...)
			var s LookupRouteFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteFilterResultOutput)
}

type LookupRouteFilterOutputArgs struct {
	// Expands referenced express route bgp peering resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName pulumi.StringInput `pulumi:"routeFilterName"`
}

func (LookupRouteFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterArgs)(nil)).Elem()
}

// Route Filter Resource.
type LookupRouteFilterResultOutput struct{ *pulumi.OutputState }

func (LookupRouteFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteFilterResult)(nil)).Elem()
}

func (o LookupRouteFilterResultOutput) ToLookupRouteFilterResultOutput() LookupRouteFilterResultOutput {
	return o
}

func (o LookupRouteFilterResultOutput) ToLookupRouteFilterResultOutputWithContext(ctx context.Context) LookupRouteFilterResultOutput {
	return o
}

func (o LookupRouteFilterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRouteFilterResult] {
	return pulumix.Output[LookupRouteFilterResult]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupRouteFilterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupRouteFilterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A collection of references to express route circuit ipv6 peerings.
func (o LookupRouteFilterResultOutput) Ipv6Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) []ExpressRouteCircuitPeeringResponse { return v.Ipv6Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

// Resource location.
func (o LookupRouteFilterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupRouteFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of references to express route circuit peerings.
func (o LookupRouteFilterResultOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) []ExpressRouteCircuitPeeringResponse { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

// The provisioning state of the route filter resource.
func (o LookupRouteFilterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Collection of RouteFilterRules contained within a route filter.
func (o LookupRouteFilterResultOutput) Rules() RouteFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) []RouteFilterRuleResponse { return v.Rules }).(RouteFilterRuleResponseArrayOutput)
}

// Resource tags.
func (o LookupRouteFilterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupRouteFilterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteFilterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteFilterResultOutput{})
}

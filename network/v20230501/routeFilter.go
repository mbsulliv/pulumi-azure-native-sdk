// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Route Filter Resource.
type RouteFilter struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A collection of references to express route circuit ipv6 peerings.
	Ipv6Peerings ExpressRouteCircuitPeeringResponseArrayOutput `pulumi:"ipv6Peerings"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A collection of references to express route circuit peerings.
	Peerings ExpressRouteCircuitPeeringResponseArrayOutput `pulumi:"peerings"`
	// The provisioning state of the route filter resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Collection of RouteFilterRules contained within a route filter.
	Rules RouteFilterRuleResponseArrayOutput `pulumi:"rules"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRouteFilter registers a new resource with the given unique name, arguments, and options.
func NewRouteFilter(ctx *pulumi.Context,
	name string, args *RouteFilterArgs, opts ...pulumi.ResourceOption) (*RouteFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:RouteFilter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:RouteFilter"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RouteFilter
	err := ctx.RegisterResource("azure-native:network/v20230501:RouteFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteFilter gets an existing RouteFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteFilterState, opts ...pulumi.ResourceOption) (*RouteFilter, error) {
	var resource RouteFilter
	err := ctx.ReadResource("azure-native:network/v20230501:RouteFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteFilter resources.
type routeFilterState struct {
}

type RouteFilterState struct {
}

func (RouteFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterState)(nil)).Elem()
}

type routeFilterArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName *string `pulumi:"routeFilterName"`
	// Collection of RouteFilterRules contained within a route filter.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Rules []RouteFilterRuleType `pulumi:"rules"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RouteFilter resource.
type RouteFilterArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the route filter.
	RouteFilterName pulumi.StringPtrInput
	// Collection of RouteFilterRules contained within a route filter.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Rules RouteFilterRuleTypeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (RouteFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterArgs)(nil)).Elem()
}

type RouteFilterInput interface {
	pulumi.Input

	ToRouteFilterOutput() RouteFilterOutput
	ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput
}

func (*RouteFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilter)(nil)).Elem()
}

func (i *RouteFilter) ToRouteFilterOutput() RouteFilterOutput {
	return i.ToRouteFilterOutputWithContext(context.Background())
}

func (i *RouteFilter) ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterOutput)
}

func (i *RouteFilter) ToOutput(ctx context.Context) pulumix.Output[*RouteFilter] {
	return pulumix.Output[*RouteFilter]{
		OutputState: i.ToRouteFilterOutputWithContext(ctx).OutputState,
	}
}

type RouteFilterOutput struct{ *pulumi.OutputState }

func (RouteFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilter)(nil)).Elem()
}

func (o RouteFilterOutput) ToRouteFilterOutput() RouteFilterOutput {
	return o
}

func (o RouteFilterOutput) ToRouteFilterOutputWithContext(ctx context.Context) RouteFilterOutput {
	return o
}

func (o RouteFilterOutput) ToOutput(ctx context.Context) pulumix.Output[*RouteFilter] {
	return pulumix.Output[*RouteFilter]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o RouteFilterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A collection of references to express route circuit ipv6 peerings.
func (o RouteFilterOutput) Ipv6Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilter) ExpressRouteCircuitPeeringResponseArrayOutput { return v.Ipv6Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

// Resource location.
func (o RouteFilterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o RouteFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A collection of references to express route circuit peerings.
func (o RouteFilterOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilter) ExpressRouteCircuitPeeringResponseArrayOutput { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

// The provisioning state of the route filter resource.
func (o RouteFilterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Collection of RouteFilterRules contained within a route filter.
func (o RouteFilterOutput) Rules() RouteFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilter) RouteFilterRuleResponseArrayOutput { return v.Rules }).(RouteFilterRuleResponseArrayOutput)
}

// Resource tags.
func (o RouteFilterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o RouteFilterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteFilterOutput{})
}

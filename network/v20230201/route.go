// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Route resource.
type Route struct {
	pulumi.CustomResourceState

	// The destination CIDR to which the route applies.
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
	HasBgpOverride pulumi.BoolPtrOutput `pulumi:"hasBgpOverride"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress pulumi.StringPtrOutput `pulumi:"nextHopIpAddress"`
	// The type of Azure hop the packet should be sent to.
	NextHopType pulumi.StringOutput `pulumi:"nextHopType"`
	// The provisioning state of the route resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the resource.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NextHopType == nil {
		return nil, errors.New("invalid value for required argument 'NextHopType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RouteTableName == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:Route"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:Route"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("azure-native:network/v20230201:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("azure-native:network/v20230201:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The destination CIDR to which the route applies.
	AddressPrefix *string `pulumi:"addressPrefix"`
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
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route.
	RouteName *string `pulumi:"routeName"`
	// The name of the route table.
	RouteTableName string `pulumi:"routeTableName"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The destination CIDR to which the route applies.
	AddressPrefix pulumi.StringPtrInput
	// A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
	HasBgpOverride pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress pulumi.StringPtrInput
	// The type of Azure hop the packet should be sent to.
	NextHopType pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the route.
	RouteName pulumi.StringPtrInput
	// The name of the route table.
	RouteTableName pulumi.StringInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i *Route) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: i.ToRouteOutputWithContext(ctx).OutputState,
	}
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: o.OutputState,
	}
}

// The destination CIDR to which the route applies.
func (o RouteOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o RouteOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
func (o RouteOutput) HasBgpOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.BoolPtrOutput { return v.HasBgpOverride }).(pulumi.BoolPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o RouteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
func (o RouteOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

// The type of Azure hop the packet should be sent to.
func (o RouteOutput) NextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.NextHopType }).(pulumi.StringOutput)
}

// The provisioning state of the route resource.
func (o RouteOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource.
func (o RouteOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}

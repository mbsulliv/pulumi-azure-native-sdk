// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Route table resource.
type RouteTable struct {
	pulumi.CustomResourceState

	// Whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation pulumi.BoolPtrOutput `pulumi:"disableBgpRoutePropagation"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the route table resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the route table.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Collection of routes contained within a route table.
	Routes RouteResponseArrayOutput `pulumi:"routes"`
	// A collection of references to subnets.
	Subnets SubnetResponseArrayOutput `pulumi:"subnets"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:RouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:RouteTable"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RouteTable
	err := ctx.RegisterResource("azure-native:network/v20230501:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("azure-native:network/v20230501:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
}

type RouteTableState struct {
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	// Whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `pulumi:"disableBgpRoutePropagation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route table.
	RouteTableName *string `pulumi:"routeTableName"`
	// Collection of routes contained within a route table.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Routes []RouteType `pulumi:"routes"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// Whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the route table.
	RouteTableName pulumi.StringPtrInput
	// Collection of routes contained within a route table.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Routes RouteTypeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (*RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (i *RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

type RouteTableOutput struct{ *pulumi.OutputState }

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

// Whether to disable the routes learned by BGP on that route table. True means disable.
func (o RouteTableOutput) DisableBgpRoutePropagation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.BoolPtrOutput { return v.DisableBgpRoutePropagation }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o RouteTableOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o RouteTableOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o RouteTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the route table resource.
func (o RouteTableOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the route table.
func (o RouteTableOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Collection of routes contained within a route table.
func (o RouteTableOutput) Routes() RouteResponseArrayOutput {
	return o.ApplyT(func(v *RouteTable) RouteResponseArrayOutput { return v.Routes }).(RouteResponseArrayOutput)
}

// A collection of references to subnets.
func (o RouteTableOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *RouteTable) SubnetResponseArrayOutput { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Resource tags.
func (o RouteTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o RouteTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteTableOutput{})
}

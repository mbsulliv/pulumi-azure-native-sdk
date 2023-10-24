// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ExpressRoutePort resource definition.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2019-08-01, 2023-04-01, 2023-05-01.
type ExpressRoutePort struct {
	pulumi.CustomResourceState

	// Date of the physical port allocation to be used in Letter of Authorization.
	AllocationDate pulumi.StringOutput `pulumi:"allocationDate"`
	// Bandwidth of procured ports in Gbps.
	BandwidthInGbps pulumi.IntPtrOutput `pulumi:"bandwidthInGbps"`
	// The billing type of the ExpressRoutePort resource.
	BillingType pulumi.StringPtrOutput `pulumi:"billingType"`
	// Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
	Circuits SubResourceResponseArrayOutput `pulumi:"circuits"`
	// Encapsulation method on physical ports.
	Encapsulation pulumi.StringPtrOutput `pulumi:"encapsulation"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Ether type of the physical port.
	EtherType pulumi.StringOutput `pulumi:"etherType"`
	// The identity of ExpressRoutePort, if configured.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The set of physical links of the ExpressRoutePort resource.
	Links ExpressRouteLinkResponseArrayOutput `pulumi:"links"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Maximum transmission unit of the physical port pair(s).
	Mtu pulumi.StringOutput `pulumi:"mtu"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the peering location that the ExpressRoutePort is mapped to physically.
	PeeringLocation pulumi.StringPtrOutput `pulumi:"peeringLocation"`
	// Aggregate Gbps of associated circuit bandwidths.
	ProvisionedBandwidthInGbps pulumi.Float64Output `pulumi:"provisionedBandwidthInGbps"`
	// The provisioning state of the express route port resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the express route port resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExpressRoutePort registers a new resource with the given unique name, arguments, and options.
func NewExpressRoutePort(ctx *pulumi.Context,
	name string, args *ExpressRoutePortArgs, opts ...pulumi.ResourceOption) (*ExpressRoutePort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ExpressRoutePort"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ExpressRoutePort"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ExpressRoutePort
	err := ctx.RegisterResource("azure-native:network:ExpressRoutePort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRoutePort gets an existing ExpressRoutePort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRoutePort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRoutePortState, opts ...pulumi.ResourceOption) (*ExpressRoutePort, error) {
	var resource ExpressRoutePort
	err := ctx.ReadResource("azure-native:network:ExpressRoutePort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRoutePort resources.
type expressRoutePortState struct {
}

type ExpressRoutePortState struct {
}

func (ExpressRoutePortState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRoutePortState)(nil)).Elem()
}

type expressRoutePortArgs struct {
	// Bandwidth of procured ports in Gbps.
	BandwidthInGbps *int `pulumi:"bandwidthInGbps"`
	// The billing type of the ExpressRoutePort resource.
	BillingType *string `pulumi:"billingType"`
	// Encapsulation method on physical ports.
	Encapsulation *string `pulumi:"encapsulation"`
	// The name of the ExpressRoutePort resource.
	ExpressRoutePortName *string `pulumi:"expressRoutePortName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The identity of ExpressRoutePort, if configured.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The set of physical links of the ExpressRoutePort resource.
	Links []ExpressRouteLink `pulumi:"links"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the peering location that the ExpressRoutePort is mapped to physically.
	PeeringLocation *string `pulumi:"peeringLocation"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ExpressRoutePort resource.
type ExpressRoutePortArgs struct {
	// Bandwidth of procured ports in Gbps.
	BandwidthInGbps pulumi.IntPtrInput
	// The billing type of the ExpressRoutePort resource.
	BillingType pulumi.StringPtrInput
	// Encapsulation method on physical ports.
	Encapsulation pulumi.StringPtrInput
	// The name of the ExpressRoutePort resource.
	ExpressRoutePortName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The identity of ExpressRoutePort, if configured.
	Identity ManagedServiceIdentityPtrInput
	// The set of physical links of the ExpressRoutePort resource.
	Links ExpressRouteLinkArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the peering location that the ExpressRoutePort is mapped to physically.
	PeeringLocation pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ExpressRoutePortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRoutePortArgs)(nil)).Elem()
}

type ExpressRoutePortInput interface {
	pulumi.Input

	ToExpressRoutePortOutput() ExpressRoutePortOutput
	ToExpressRoutePortOutputWithContext(ctx context.Context) ExpressRoutePortOutput
}

func (*ExpressRoutePort) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePort)(nil)).Elem()
}

func (i *ExpressRoutePort) ToExpressRoutePortOutput() ExpressRoutePortOutput {
	return i.ToExpressRoutePortOutputWithContext(context.Background())
}

func (i *ExpressRoutePort) ToExpressRoutePortOutputWithContext(ctx context.Context) ExpressRoutePortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRoutePortOutput)
}

func (i *ExpressRoutePort) ToOutput(ctx context.Context) pulumix.Output[*ExpressRoutePort] {
	return pulumix.Output[*ExpressRoutePort]{
		OutputState: i.ToExpressRoutePortOutputWithContext(ctx).OutputState,
	}
}

type ExpressRoutePortOutput struct{ *pulumi.OutputState }

func (ExpressRoutePortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePort)(nil)).Elem()
}

func (o ExpressRoutePortOutput) ToExpressRoutePortOutput() ExpressRoutePortOutput {
	return o
}

func (o ExpressRoutePortOutput) ToExpressRoutePortOutputWithContext(ctx context.Context) ExpressRoutePortOutput {
	return o
}

func (o ExpressRoutePortOutput) ToOutput(ctx context.Context) pulumix.Output[*ExpressRoutePort] {
	return pulumix.Output[*ExpressRoutePort]{
		OutputState: o.OutputState,
	}
}

// Date of the physical port allocation to be used in Letter of Authorization.
func (o ExpressRoutePortOutput) AllocationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.AllocationDate }).(pulumi.StringOutput)
}

// Bandwidth of procured ports in Gbps.
func (o ExpressRoutePortOutput) BandwidthInGbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.IntPtrOutput { return v.BandwidthInGbps }).(pulumi.IntPtrOutput)
}

// The billing type of the ExpressRoutePort resource.
func (o ExpressRoutePortOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringPtrOutput { return v.BillingType }).(pulumi.StringPtrOutput)
}

// Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
func (o ExpressRoutePortOutput) Circuits() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRoutePort) SubResourceResponseArrayOutput { return v.Circuits }).(SubResourceResponseArrayOutput)
}

// Encapsulation method on physical ports.
func (o ExpressRoutePortOutput) Encapsulation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringPtrOutput { return v.Encapsulation }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ExpressRoutePortOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Ether type of the physical port.
func (o ExpressRoutePortOutput) EtherType() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.EtherType }).(pulumi.StringOutput)
}

// The identity of ExpressRoutePort, if configured.
func (o ExpressRoutePortOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The set of physical links of the ExpressRoutePort resource.
func (o ExpressRoutePortOutput) Links() ExpressRouteLinkResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRoutePort) ExpressRouteLinkResponseArrayOutput { return v.Links }).(ExpressRouteLinkResponseArrayOutput)
}

// Resource location.
func (o ExpressRoutePortOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Maximum transmission unit of the physical port pair(s).
func (o ExpressRoutePortOutput) Mtu() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.Mtu }).(pulumi.StringOutput)
}

// Resource name.
func (o ExpressRoutePortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the peering location that the ExpressRoutePort is mapped to physically.
func (o ExpressRoutePortOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringPtrOutput { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

// Aggregate Gbps of associated circuit bandwidths.
func (o ExpressRoutePortOutput) ProvisionedBandwidthInGbps() pulumi.Float64Output {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.Float64Output { return v.ProvisionedBandwidthInGbps }).(pulumi.Float64Output)
}

// The provisioning state of the express route port resource.
func (o ExpressRoutePortOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the express route port resource.
func (o ExpressRoutePortOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o ExpressRoutePortOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ExpressRoutePortOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePort) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRoutePortOutput{})
}

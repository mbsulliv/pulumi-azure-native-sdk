// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A virtual network.
// Azure REST API version: 2018-09-15. Prior API version in Azure Native 1.x: 2018-09-15
type VirtualNetwork struct {
	pulumi.CustomResourceState

	// The allowed subnets of the virtual network.
	AllowedSubnets SubnetResponseArrayOutput `pulumi:"allowedSubnets"`
	// The creation date of the virtual network.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The description of the virtual network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId pulumi.StringPtrOutput `pulumi:"externalProviderResourceId"`
	// The external subnet properties.
	ExternalSubnets ExternalSubnetResponseArrayOutput `pulumi:"externalSubnets"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The subnet overrides of the virtual network.
	SubnetOverrides SubnetOverrideResponseArrayOutput `pulumi:"subnetOverrides"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
}

// NewVirtualNetwork registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:VirtualNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure-native:devtestlab:VirtualNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetwork gets an existing VirtualNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkState, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	var resource VirtualNetwork
	err := ctx.ReadResource("azure-native:devtestlab:VirtualNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetwork resources.
type virtualNetworkState struct {
}

type VirtualNetworkState struct {
}

func (VirtualNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkState)(nil)).Elem()
}

type virtualNetworkArgs struct {
	// The allowed subnets of the virtual network.
	AllowedSubnets []Subnet `pulumi:"allowedSubnets"`
	// The description of the virtual network.
	Description *string `pulumi:"description"`
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId *string `pulumi:"externalProviderResourceId"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the virtual network.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subnet overrides of the virtual network.
	SubnetOverrides []SubnetOverride `pulumi:"subnetOverrides"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualNetwork resource.
type VirtualNetworkArgs struct {
	// The allowed subnets of the virtual network.
	AllowedSubnets SubnetArrayInput
	// The description of the virtual network.
	Description pulumi.StringPtrInput
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId pulumi.StringPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the virtual network.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The subnet overrides of the virtual network.
	SubnetOverrides SubnetOverrideArrayInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
}

func (VirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkArgs)(nil)).Elem()
}

type VirtualNetworkInput interface {
	pulumi.Input

	ToVirtualNetworkOutput() VirtualNetworkOutput
	ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput
}

func (*VirtualNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetwork)(nil)).Elem()
}

func (i *VirtualNetwork) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return i.ToVirtualNetworkOutputWithContext(context.Background())
}

func (i *VirtualNetwork) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkOutput)
}

func (i *VirtualNetwork) ToOutput(ctx context.Context) pulumix.Output[*VirtualNetwork] {
	return pulumix.Output[*VirtualNetwork]{
		OutputState: i.ToVirtualNetworkOutputWithContext(ctx).OutputState,
	}
}

type VirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return o
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return o
}

func (o VirtualNetworkOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualNetwork] {
	return pulumix.Output[*VirtualNetwork]{
		OutputState: o.OutputState,
	}
}

// The allowed subnets of the virtual network.
func (o VirtualNetworkOutput) AllowedSubnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) SubnetResponseArrayOutput { return v.AllowedSubnets }).(SubnetResponseArrayOutput)
}

// The creation date of the virtual network.
func (o VirtualNetworkOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// The description of the virtual network.
func (o VirtualNetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Microsoft.Network resource identifier of the virtual network.
func (o VirtualNetworkOutput) ExternalProviderResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.ExternalProviderResourceId }).(pulumi.StringPtrOutput)
}

// The external subnet properties.
func (o VirtualNetworkOutput) ExternalSubnets() ExternalSubnetResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) ExternalSubnetResponseArrayOutput { return v.ExternalSubnets }).(ExternalSubnetResponseArrayOutput)
}

// The location of the resource.
func (o VirtualNetworkOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o VirtualNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o VirtualNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The subnet overrides of the virtual network.
func (o VirtualNetworkOutput) SubnetOverrides() SubnetOverrideResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) SubnetOverrideResponseArrayOutput { return v.SubnetOverrides }).(SubnetOverrideResponseArrayOutput)
}

// The tags of the resource.
func (o VirtualNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o VirtualNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o VirtualNetworkOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
}

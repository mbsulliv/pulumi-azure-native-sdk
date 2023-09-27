// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Virtual Network resource.
type VirtualNetwork struct {
	pulumi.CustomResourceState

	// The AddressSpace that contains an array of IP address ranges that can be used by subnets.
	AddressSpace AddressSpaceResponsePtrOutput `pulumi:"addressSpace"`
	// The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan SubResourceResponsePtrOutput `pulumi:"ddosProtectionPlan"`
	// The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
	DhcpOptions DhcpOptionsResponsePtrOutput `pulumi:"dhcpOptions"`
	// Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
	EnableDdosProtection pulumi.BoolPtrOutput `pulumi:"enableDdosProtection"`
	// Indicates if VM protection is enabled for all the subnets in the virtual network.
	EnableVmProtection pulumi.BoolPtrOutput `pulumi:"enableVmProtection"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The resourceGuid property of the Virtual Network resource.
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// A list of subnets in a Virtual Network.
	Subnets SubnetResponseArrayOutput `pulumi:"subnets"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of peerings in a Virtual Network.
	VirtualNetworkPeerings VirtualNetworkPeeringResponseArrayOutput `pulumi:"virtualNetworkPeerings"`
}

// NewVirtualNetwork registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.EnableDdosProtection == nil {
		args.EnableDdosProtection = pulumi.BoolPtr(false)
	}
	if args.EnableVmProtection == nil {
		args.EnableVmProtection = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VirtualNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure-native:network/v20190601:VirtualNetwork", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:network/v20190601:VirtualNetwork", name, id, state, &resource, opts...)
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
	// The AddressSpace that contains an array of IP address ranges that can be used by subnets.
	AddressSpace *AddressSpace `pulumi:"addressSpace"`
	// The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan *SubResource `pulumi:"ddosProtectionPlan"`
	// The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
	DhcpOptions *DhcpOptions `pulumi:"dhcpOptions"`
	// Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
	EnableDdosProtection *bool `pulumi:"enableDdosProtection"`
	// Indicates if VM protection is enabled for all the subnets in the virtual network.
	EnableVmProtection *bool `pulumi:"enableVmProtection"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resourceGuid property of the Virtual Network resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// A list of subnets in a Virtual Network.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Subnets []SubnetType `pulumi:"subnets"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the virtual network.
	VirtualNetworkName *string `pulumi:"virtualNetworkName"`
	// A list of peerings in a Virtual Network.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	VirtualNetworkPeerings []VirtualNetworkPeeringType `pulumi:"virtualNetworkPeerings"`
}

// The set of arguments for constructing a VirtualNetwork resource.
type VirtualNetworkArgs struct {
	// The AddressSpace that contains an array of IP address ranges that can be used by subnets.
	AddressSpace AddressSpacePtrInput
	// The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan SubResourcePtrInput
	// The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
	DhcpOptions DhcpOptionsPtrInput
	// Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
	EnableDdosProtection pulumi.BoolPtrInput
	// Indicates if VM protection is enabled for all the subnets in the virtual network.
	EnableVmProtection pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resourceGuid property of the Virtual Network resource.
	ResourceGuid pulumi.StringPtrInput
	// A list of subnets in a Virtual Network.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Subnets SubnetTypeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringPtrInput
	// A list of peerings in a Virtual Network.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	VirtualNetworkPeerings VirtualNetworkPeeringTypeArrayInput
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

// The AddressSpace that contains an array of IP address ranges that can be used by subnets.
func (o VirtualNetworkOutput) AddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) AddressSpaceResponsePtrOutput { return v.AddressSpace }).(AddressSpaceResponsePtrOutput)
}

// The DDoS protection plan associated with the virtual network.
func (o VirtualNetworkOutput) DdosProtectionPlan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) SubResourceResponsePtrOutput { return v.DdosProtectionPlan }).(SubResourceResponsePtrOutput)
}

// The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
func (o VirtualNetworkOutput) DhcpOptions() DhcpOptionsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) DhcpOptionsResponsePtrOutput { return v.DhcpOptions }).(DhcpOptionsResponsePtrOutput)
}

// Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
func (o VirtualNetworkOutput) EnableDdosProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.BoolPtrOutput { return v.EnableDdosProtection }).(pulumi.BoolPtrOutput)
}

// Indicates if VM protection is enabled for all the subnets in the virtual network.
func (o VirtualNetworkOutput) EnableVmProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.BoolPtrOutput { return v.EnableVmProtection }).(pulumi.BoolPtrOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o VirtualNetworkOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o VirtualNetworkOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o VirtualNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o VirtualNetworkOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The resourceGuid property of the Virtual Network resource.
func (o VirtualNetworkOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

// A list of subnets in a Virtual Network.
func (o VirtualNetworkOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) SubnetResponseArrayOutput { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Resource tags.
func (o VirtualNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VirtualNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of peerings in a Virtual Network.
func (o VirtualNetworkOutput) VirtualNetworkPeerings() VirtualNetworkPeeringResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) VirtualNetworkPeeringResponseArrayOutput { return v.VirtualNetworkPeerings }).(VirtualNetworkPeeringResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
}

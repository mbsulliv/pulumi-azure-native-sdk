// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220521preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The VirtualNetworks resource definition.
type VirtualNetwork struct {
	pulumi.CustomResourceState

	// The extended location.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId pulumi.StringPtrOutput `pulumi:"inventoryItemId"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the virtual network in vmmServer.
	NetworkName pulumi.StringOutput `pulumi:"networkName"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique ID of the virtual network.
	Uuid pulumi.StringPtrOutput `pulumi:"uuid"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrOutput `pulumi:"vmmServerId"`
}

// NewVirtualNetwork registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20200605preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20230401preview:VirtualNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure-native:scvmm/v20220521preview:VirtualNetwork", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:scvmm/v20220521preview:VirtualNetwork", name, id, state, &resource, opts...)
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
	// The extended location.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Unique ID of the virtual network.
	Uuid *string `pulumi:"uuid"`
	// Name of the VirtualNetwork.
	VirtualNetworkName *string `pulumi:"virtualNetworkName"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId *string `pulumi:"vmmServerId"`
}

// The set of arguments for constructing a VirtualNetwork resource.
type VirtualNetworkArgs struct {
	// The extended location.
	ExtendedLocation ExtendedLocationInput
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Unique ID of the virtual network.
	Uuid pulumi.StringPtrInput
	// Name of the VirtualNetwork.
	VirtualNetworkName pulumi.StringPtrInput
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId pulumi.StringPtrInput
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

// The extended location.
func (o VirtualNetworkOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VirtualNetwork) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Gets or sets the inventory Item ID for the resource.
func (o VirtualNetworkOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o VirtualNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o VirtualNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of the virtual network in vmmServer.
func (o VirtualNetworkOutput) NetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.NetworkName }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o VirtualNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system data.
func (o VirtualNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o VirtualNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o VirtualNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique ID of the virtual network.
func (o VirtualNetworkOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

// ARM Id of the vmmServer resource in which this resource resides.
func (o VirtualNetworkOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
}

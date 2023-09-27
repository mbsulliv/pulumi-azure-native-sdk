// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The virtualNetworks resource definition.
type VirtualNetworkRetrieve struct {
	pulumi.CustomResourceState

	ExtendedLocation VirtualNetworksResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
	Properties VirtualNetworksPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualNetworkRetrieve registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkRetrieve(ctx *pulumi.Context,
	name string, args *VirtualNetworkRetrieveArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220901preview:virtualNetworkRetrieve"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice:VirtualNetworkRetrieve"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice:virtualNetworkRetrieve"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220501preview:VirtualNetworkRetrieve"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220501preview:virtualNetworkRetrieve"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualNetworkRetrieve
	err := ctx.RegisterResource("azure-native:hybridcontainerservice/v20220901preview:VirtualNetworkRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkRetrieve gets an existing VirtualNetworkRetrieve resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkRetrieveState, opts ...pulumi.ResourceOption) (*VirtualNetworkRetrieve, error) {
	var resource VirtualNetworkRetrieve
	err := ctx.ReadResource("azure-native:hybridcontainerservice/v20220901preview:VirtualNetworkRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkRetrieve resources.
type virtualNetworkRetrieveState struct {
}

type VirtualNetworkRetrieveState struct {
}

func (VirtualNetworkRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRetrieveState)(nil)).Elem()
}

type virtualNetworkRetrieveArgs struct {
	ExtendedLocation *VirtualNetworksExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
	Properties *VirtualNetworksProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Parameter for the name of the virtual network
	VirtualNetworksName *string `pulumi:"virtualNetworksName"`
}

// The set of arguments for constructing a VirtualNetworkRetrieve resource.
type VirtualNetworkRetrieveArgs struct {
	ExtendedLocation VirtualNetworksExtendedLocationPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
	Properties VirtualNetworksPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Parameter for the name of the virtual network
	VirtualNetworksName pulumi.StringPtrInput
}

func (VirtualNetworkRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRetrieveArgs)(nil)).Elem()
}

type VirtualNetworkRetrieveInput interface {
	pulumi.Input

	ToVirtualNetworkRetrieveOutput() VirtualNetworkRetrieveOutput
	ToVirtualNetworkRetrieveOutputWithContext(ctx context.Context) VirtualNetworkRetrieveOutput
}

func (*VirtualNetworkRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRetrieve)(nil)).Elem()
}

func (i *VirtualNetworkRetrieve) ToVirtualNetworkRetrieveOutput() VirtualNetworkRetrieveOutput {
	return i.ToVirtualNetworkRetrieveOutputWithContext(context.Background())
}

func (i *VirtualNetworkRetrieve) ToVirtualNetworkRetrieveOutputWithContext(ctx context.Context) VirtualNetworkRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRetrieveOutput)
}

func (i *VirtualNetworkRetrieve) ToOutput(ctx context.Context) pulumix.Output[*VirtualNetworkRetrieve] {
	return pulumix.Output[*VirtualNetworkRetrieve]{
		OutputState: i.ToVirtualNetworkRetrieveOutputWithContext(ctx).OutputState,
	}
}

type VirtualNetworkRetrieveOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRetrieve)(nil)).Elem()
}

func (o VirtualNetworkRetrieveOutput) ToVirtualNetworkRetrieveOutput() VirtualNetworkRetrieveOutput {
	return o
}

func (o VirtualNetworkRetrieveOutput) ToVirtualNetworkRetrieveOutputWithContext(ctx context.Context) VirtualNetworkRetrieveOutput {
	return o
}

func (o VirtualNetworkRetrieveOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualNetworkRetrieve] {
	return pulumix.Output[*VirtualNetworkRetrieve]{
		OutputState: o.OutputState,
	}
}

func (o VirtualNetworkRetrieveOutput) ExtendedLocation() VirtualNetworksResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) VirtualNetworksResponseExtendedLocationPtrOutput {
		return v.ExtendedLocation
	}).(VirtualNetworksResponseExtendedLocationPtrOutput)
}

// The geo-location where the resource lives
func (o VirtualNetworkRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o VirtualNetworkRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
func (o VirtualNetworkRetrieveOutput) Properties() VirtualNetworksPropertiesResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) VirtualNetworksPropertiesResponseOutput { return v.Properties }).(VirtualNetworksPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o VirtualNetworkRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o VirtualNetworkRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VirtualNetworkRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkRetrieveOutput{})
}

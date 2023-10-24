// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mobilenetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Network slice resource. Must be created in the same location as its parent mobile network.
// Azure REST API version: 2023-06-01. Prior API version in Azure Native 1.x: 2022-04-01-preview.
//
// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-09-01.
type Slice struct {
	pulumi.CustomResourceState

	// An optional description for this network slice.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the network slice resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Single-network slice selection assistance information (S-NSSAI). Unique at the scope of a mobile network.
	Snssai SnssaiResponseOutput `pulumi:"snssai"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSlice registers a new resource with the given unique name, arguments, and options.
func NewSlice(ctx *pulumi.Context,
	name string, args *SliceArgs, opts ...pulumi.ResourceOption) (*Slice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Snssai == nil {
		return nil, errors.New("invalid value for required argument 'Snssai'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:Slice"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:Slice"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:Slice"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230601:Slice"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230901:Slice"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Slice
	err := ctx.RegisterResource("azure-native:mobilenetwork:Slice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlice gets an existing Slice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SliceState, opts ...pulumi.ResourceOption) (*Slice, error) {
	var resource Slice
	err := ctx.ReadResource("azure-native:mobilenetwork:Slice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Slice resources.
type sliceState struct {
}

type SliceState struct {
}

func (SliceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sliceState)(nil)).Elem()
}

type sliceArgs struct {
	// An optional description for this network slice.
	Description *string `pulumi:"description"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network slice.
	SliceName *string `pulumi:"sliceName"`
	// Single-network slice selection assistance information (S-NSSAI). Unique at the scope of a mobile network.
	Snssai Snssai `pulumi:"snssai"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Slice resource.
type SliceArgs struct {
	// An optional description for this network slice.
	Description pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the network slice.
	SliceName pulumi.StringPtrInput
	// Single-network slice selection assistance information (S-NSSAI). Unique at the scope of a mobile network.
	Snssai SnssaiInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SliceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sliceArgs)(nil)).Elem()
}

type SliceInput interface {
	pulumi.Input

	ToSliceOutput() SliceOutput
	ToSliceOutputWithContext(ctx context.Context) SliceOutput
}

func (*Slice) ElementType() reflect.Type {
	return reflect.TypeOf((**Slice)(nil)).Elem()
}

func (i *Slice) ToSliceOutput() SliceOutput {
	return i.ToSliceOutputWithContext(context.Background())
}

func (i *Slice) ToSliceOutputWithContext(ctx context.Context) SliceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SliceOutput)
}

func (i *Slice) ToOutput(ctx context.Context) pulumix.Output[*Slice] {
	return pulumix.Output[*Slice]{
		OutputState: i.ToSliceOutputWithContext(ctx).OutputState,
	}
}

type SliceOutput struct{ *pulumi.OutputState }

func (SliceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Slice)(nil)).Elem()
}

func (o SliceOutput) ToSliceOutput() SliceOutput {
	return o
}

func (o SliceOutput) ToSliceOutputWithContext(ctx context.Context) SliceOutput {
	return o
}

func (o SliceOutput) ToOutput(ctx context.Context) pulumix.Output[*Slice] {
	return pulumix.Output[*Slice]{
		OutputState: o.OutputState,
	}
}

// An optional description for this network slice.
func (o SliceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o SliceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SliceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the network slice resource.
func (o SliceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Single-network slice selection assistance information (S-NSSAI). Unique at the scope of a mobile network.
func (o SliceOutput) Snssai() SnssaiResponseOutput {
	return o.ApplyT(func(v *Slice) SnssaiResponseOutput { return v.Snssai }).(SnssaiResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SliceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Slice) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SliceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SliceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SliceOutput{})
}

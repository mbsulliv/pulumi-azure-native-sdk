// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Subvolume Information properties
type Subvolume struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// parent path to the subvolume
	ParentPath pulumi.StringPtrOutput `pulumi:"parentPath"`
	// Path to the subvolume
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSubvolume registers a new resource with the given unique name, arguments, and options.
func NewSubvolume(ctx *pulumi.Context,
	name string, args *SubvolumeArgs, opts ...pulumi.ResourceOption) (*Subvolume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101preview:Subvolume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230501:Subvolume"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Subvolume
	err := ctx.RegisterResource("azure-native:netapp/v20221101:Subvolume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubvolume gets an existing Subvolume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubvolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubvolumeState, opts ...pulumi.ResourceOption) (*Subvolume, error) {
	var resource Subvolume
	err := ctx.ReadResource("azure-native:netapp/v20221101:Subvolume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subvolume resources.
type subvolumeState struct {
}

type SubvolumeState struct {
}

func (SubvolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*subvolumeState)(nil)).Elem()
}

type subvolumeArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// parent path to the subvolume
	ParentPath *string `pulumi:"parentPath"`
	// Path to the subvolume
	Path *string `pulumi:"path"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Truncate subvolume to the provided size in bytes
	Size *float64 `pulumi:"size"`
	// The name of the subvolume.
	SubvolumeName *string `pulumi:"subvolumeName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// The set of arguments for constructing a Subvolume resource.
type SubvolumeArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// parent path to the subvolume
	ParentPath pulumi.StringPtrInput
	// Path to the subvolume
	Path pulumi.StringPtrInput
	// The name of the capacity pool
	PoolName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Truncate subvolume to the provided size in bytes
	Size pulumi.Float64PtrInput
	// The name of the subvolume.
	SubvolumeName pulumi.StringPtrInput
	// The name of the volume
	VolumeName pulumi.StringInput
}

func (SubvolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subvolumeArgs)(nil)).Elem()
}

type SubvolumeInput interface {
	pulumi.Input

	ToSubvolumeOutput() SubvolumeOutput
	ToSubvolumeOutputWithContext(ctx context.Context) SubvolumeOutput
}

func (*Subvolume) ElementType() reflect.Type {
	return reflect.TypeOf((**Subvolume)(nil)).Elem()
}

func (i *Subvolume) ToSubvolumeOutput() SubvolumeOutput {
	return i.ToSubvolumeOutputWithContext(context.Background())
}

func (i *Subvolume) ToSubvolumeOutputWithContext(ctx context.Context) SubvolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubvolumeOutput)
}

func (i *Subvolume) ToOutput(ctx context.Context) pulumix.Output[*Subvolume] {
	return pulumix.Output[*Subvolume]{
		OutputState: i.ToSubvolumeOutputWithContext(ctx).OutputState,
	}
}

type SubvolumeOutput struct{ *pulumi.OutputState }

func (SubvolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subvolume)(nil)).Elem()
}

func (o SubvolumeOutput) ToSubvolumeOutput() SubvolumeOutput {
	return o
}

func (o SubvolumeOutput) ToSubvolumeOutputWithContext(ctx context.Context) SubvolumeOutput {
	return o
}

func (o SubvolumeOutput) ToOutput(ctx context.Context) pulumix.Output[*Subvolume] {
	return pulumix.Output[*Subvolume]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o SubvolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// parent path to the subvolume
func (o SubvolumeOutput) ParentPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringPtrOutput { return v.ParentPath }).(pulumi.StringPtrOutput)
}

// Path to the subvolume
func (o SubvolumeOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// Azure lifecycle management
func (o SubvolumeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SubvolumeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Subvolume) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SubvolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Subvolume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubvolumeOutput{})
}

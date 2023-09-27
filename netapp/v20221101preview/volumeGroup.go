// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Volume group resource for create
type VolumeGroup struct {
	pulumi.CustomResourceState

	// Volume group details
	GroupMetaData VolumeGroupMetaDataResponsePtrOutput `pulumi:"groupMetaData"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// List of volumes from group
	Volumes VolumeGroupVolumePropertiesResponseArrayOutput `pulumi:"volumes"`
}

// NewVolumeGroup registers a new resource with the given unique name, arguments, and options.
func NewVolumeGroup(ctx *pulumi.Context,
	name string, args *VolumeGroupArgs, opts ...pulumi.ResourceOption) (*VolumeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230501:VolumeGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VolumeGroup
	err := ctx.RegisterResource("azure-native:netapp/v20221101preview:VolumeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeGroup gets an existing VolumeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeGroupState, opts ...pulumi.ResourceOption) (*VolumeGroup, error) {
	var resource VolumeGroup
	err := ctx.ReadResource("azure-native:netapp/v20221101preview:VolumeGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeGroup resources.
type volumeGroupState struct {
}

type VolumeGroupState struct {
}

func (VolumeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeGroupState)(nil)).Elem()
}

type volumeGroupArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Volume group details
	GroupMetaData *VolumeGroupMetaData `pulumi:"groupMetaData"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the volumeGroup
	VolumeGroupName *string `pulumi:"volumeGroupName"`
	// List of volumes from group
	Volumes []VolumeGroupVolumeProperties `pulumi:"volumes"`
}

// The set of arguments for constructing a VolumeGroup resource.
type VolumeGroupArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// Volume group details
	GroupMetaData VolumeGroupMetaDataPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the volumeGroup
	VolumeGroupName pulumi.StringPtrInput
	// List of volumes from group
	Volumes VolumeGroupVolumePropertiesArrayInput
}

func (VolumeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeGroupArgs)(nil)).Elem()
}

type VolumeGroupInput interface {
	pulumi.Input

	ToVolumeGroupOutput() VolumeGroupOutput
	ToVolumeGroupOutputWithContext(ctx context.Context) VolumeGroupOutput
}

func (*VolumeGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeGroup)(nil)).Elem()
}

func (i *VolumeGroup) ToVolumeGroupOutput() VolumeGroupOutput {
	return i.ToVolumeGroupOutputWithContext(context.Background())
}

func (i *VolumeGroup) ToVolumeGroupOutputWithContext(ctx context.Context) VolumeGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupOutput)
}

func (i *VolumeGroup) ToOutput(ctx context.Context) pulumix.Output[*VolumeGroup] {
	return pulumix.Output[*VolumeGroup]{
		OutputState: i.ToVolumeGroupOutputWithContext(ctx).OutputState,
	}
}

type VolumeGroupOutput struct{ *pulumi.OutputState }

func (VolumeGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeGroup)(nil)).Elem()
}

func (o VolumeGroupOutput) ToVolumeGroupOutput() VolumeGroupOutput {
	return o
}

func (o VolumeGroupOutput) ToVolumeGroupOutputWithContext(ctx context.Context) VolumeGroupOutput {
	return o
}

func (o VolumeGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*VolumeGroup] {
	return pulumix.Output[*VolumeGroup]{
		OutputState: o.OutputState,
	}
}

// Volume group details
func (o VolumeGroupOutput) GroupMetaData() VolumeGroupMetaDataResponsePtrOutput {
	return o.ApplyT(func(v *VolumeGroup) VolumeGroupMetaDataResponsePtrOutput { return v.GroupMetaData }).(VolumeGroupMetaDataResponsePtrOutput)
}

// Resource location
func (o VolumeGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o VolumeGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o VolumeGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type
func (o VolumeGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// List of volumes from group
func (o VolumeGroupOutput) Volumes() VolumeGroupVolumePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *VolumeGroup) VolumeGroupVolumePropertiesResponseArrayOutput { return v.Volumes }).(VolumeGroupVolumePropertiesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeGroupOutput{})
}

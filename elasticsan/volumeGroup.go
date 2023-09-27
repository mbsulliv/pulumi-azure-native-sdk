// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsan

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response for Volume Group request.
// Azure REST API version: 2021-11-20-preview. Prior API version in Azure Native 1.x: 2021-11-20-preview
type VolumeGroup struct {
	pulumi.CustomResourceState

	// Type of encryption
	Encryption pulumi.StringPtrOutput `pulumi:"encryption"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A collection of rules governing the accessibility from specific network locations.
	NetworkAcls NetworkRuleSetResponsePtrOutput `pulumi:"networkAcls"`
	// Type of storage target
	ProtocolType pulumi.StringPtrOutput `pulumi:"protocolType"`
	// State of the operation on the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource metadata required by ARM RPC
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVolumeGroup registers a new resource with the given unique name, arguments, and options.
func NewVolumeGroup(ctx *pulumi.Context,
	name string, args *VolumeGroupArgs, opts ...pulumi.ResourceOption) (*VolumeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ElasticSanName == nil {
		return nil, errors.New("invalid value for required argument 'ElasticSanName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:elasticsan/v20211120preview:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:elasticsan/v20221201preview:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:elasticsan/v20230101:VolumeGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VolumeGroup
	err := ctx.RegisterResource("azure-native:elasticsan:VolumeGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:elasticsan:VolumeGroup", name, id, state, &resource, opts...)
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
	// The name of the ElasticSan.
	ElasticSanName string `pulumi:"elasticSanName"`
	// Type of encryption
	Encryption *string `pulumi:"encryption"`
	// A collection of rules governing the accessibility from specific network locations.
	NetworkAcls *NetworkRuleSet `pulumi:"networkAcls"`
	// Type of storage target
	ProtocolType *string `pulumi:"protocolType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the VolumeGroup.
	VolumeGroupName *string `pulumi:"volumeGroupName"`
}

// The set of arguments for constructing a VolumeGroup resource.
type VolumeGroupArgs struct {
	// The name of the ElasticSan.
	ElasticSanName pulumi.StringInput
	// Type of encryption
	Encryption pulumi.StringPtrInput
	// A collection of rules governing the accessibility from specific network locations.
	NetworkAcls NetworkRuleSetPtrInput
	// Type of storage target
	ProtocolType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// The name of the VolumeGroup.
	VolumeGroupName pulumi.StringPtrInput
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

// Type of encryption
func (o VolumeGroupOutput) Encryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringPtrOutput { return v.Encryption }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o VolumeGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A collection of rules governing the accessibility from specific network locations.
func (o VolumeGroupOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v *VolumeGroup) NetworkRuleSetResponsePtrOutput { return v.NetworkAcls }).(NetworkRuleSetResponsePtrOutput)
}

// Type of storage target
func (o VolumeGroupOutput) ProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringPtrOutput { return v.ProtocolType }).(pulumi.StringPtrOutput)
}

// State of the operation on the resource.
func (o VolumeGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource metadata required by ARM RPC
func (o VolumeGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VolumeGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o VolumeGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o VolumeGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeGroupOutput{})
}

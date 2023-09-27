// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerbidedicated

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents an instance of an auto scale v-core resource.
// Azure REST API version: 2021-01-01. Prior API version in Azure Native 1.x: 2021-01-01
type AutoScaleVCore struct {
	pulumi.CustomResourceState

	// The maximum capacity of an auto scale v-core resource.
	CapacityLimit pulumi.IntPtrOutput `pulumi:"capacityLimit"`
	// The object ID of the capacity resource associated with the auto scale v-core resource.
	CapacityObjectId pulumi.StringPtrOutput `pulumi:"capacityObjectId"`
	// Location of the PowerBI Dedicated resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the PowerBI Dedicated resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment state of an auto scale v-core resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU of the auto scale v-core resource.
	Sku AutoScaleVCoreSkuResponseOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponsePtrOutput `pulumi:"systemData"`
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the PowerBI Dedicated resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAutoScaleVCore registers a new resource with the given unique name, arguments, and options.
func NewAutoScaleVCore(ctx *pulumi.Context,
	name string, args *AutoScaleVCoreArgs, opts ...pulumi.ResourceOption) (*AutoScaleVCore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:powerbidedicated/v20210101:AutoScaleVCore"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AutoScaleVCore
	err := ctx.RegisterResource("azure-native:powerbidedicated:AutoScaleVCore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoScaleVCore gets an existing AutoScaleVCore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoScaleVCore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoScaleVCoreState, opts ...pulumi.ResourceOption) (*AutoScaleVCore, error) {
	var resource AutoScaleVCore
	err := ctx.ReadResource("azure-native:powerbidedicated:AutoScaleVCore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoScaleVCore resources.
type autoScaleVCoreState struct {
}

type AutoScaleVCoreState struct {
}

func (AutoScaleVCoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoScaleVCoreState)(nil)).Elem()
}

type autoScaleVCoreArgs struct {
	// The maximum capacity of an auto scale v-core resource.
	CapacityLimit *int `pulumi:"capacityLimit"`
	// The object ID of the capacity resource associated with the auto scale v-core resource.
	CapacityObjectId *string `pulumi:"capacityObjectId"`
	// Location of the PowerBI Dedicated resource.
	Location *string `pulumi:"location"`
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the auto scale v-core resource.
	Sku AutoScaleVCoreSku `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `pulumi:"systemData"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
	// The name of the auto scale v-core. It must be a minimum of 3 characters, and a maximum of 63.
	VcoreName *string `pulumi:"vcoreName"`
}

// The set of arguments for constructing a AutoScaleVCore resource.
type AutoScaleVCoreArgs struct {
	// The maximum capacity of an auto scale v-core resource.
	CapacityLimit pulumi.IntPtrInput
	// The object ID of the capacity resource associated with the auto scale v-core resource.
	CapacityObjectId pulumi.StringPtrInput
	// Location of the PowerBI Dedicated resource.
	Location pulumi.StringPtrInput
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName pulumi.StringInput
	// The SKU of the auto scale v-core resource.
	Sku AutoScaleVCoreSkuInput
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataPtrInput
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapInput
	// The name of the auto scale v-core. It must be a minimum of 3 characters, and a maximum of 63.
	VcoreName pulumi.StringPtrInput
}

func (AutoScaleVCoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoScaleVCoreArgs)(nil)).Elem()
}

type AutoScaleVCoreInput interface {
	pulumi.Input

	ToAutoScaleVCoreOutput() AutoScaleVCoreOutput
	ToAutoScaleVCoreOutputWithContext(ctx context.Context) AutoScaleVCoreOutput
}

func (*AutoScaleVCore) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleVCore)(nil)).Elem()
}

func (i *AutoScaleVCore) ToAutoScaleVCoreOutput() AutoScaleVCoreOutput {
	return i.ToAutoScaleVCoreOutputWithContext(context.Background())
}

func (i *AutoScaleVCore) ToAutoScaleVCoreOutputWithContext(ctx context.Context) AutoScaleVCoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleVCoreOutput)
}

func (i *AutoScaleVCore) ToOutput(ctx context.Context) pulumix.Output[*AutoScaleVCore] {
	return pulumix.Output[*AutoScaleVCore]{
		OutputState: i.ToAutoScaleVCoreOutputWithContext(ctx).OutputState,
	}
}

type AutoScaleVCoreOutput struct{ *pulumi.OutputState }

func (AutoScaleVCoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleVCore)(nil)).Elem()
}

func (o AutoScaleVCoreOutput) ToAutoScaleVCoreOutput() AutoScaleVCoreOutput {
	return o
}

func (o AutoScaleVCoreOutput) ToAutoScaleVCoreOutputWithContext(ctx context.Context) AutoScaleVCoreOutput {
	return o
}

func (o AutoScaleVCoreOutput) ToOutput(ctx context.Context) pulumix.Output[*AutoScaleVCore] {
	return pulumix.Output[*AutoScaleVCore]{
		OutputState: o.OutputState,
	}
}

// The maximum capacity of an auto scale v-core resource.
func (o AutoScaleVCoreOutput) CapacityLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleVCore) pulumi.IntPtrOutput { return v.CapacityLimit }).(pulumi.IntPtrOutput)
}

// The object ID of the capacity resource associated with the auto scale v-core resource.
func (o AutoScaleVCoreOutput) CapacityObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleVCore) pulumi.StringPtrOutput { return v.CapacityObjectId }).(pulumi.StringPtrOutput)
}

// Location of the PowerBI Dedicated resource.
func (o AutoScaleVCoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoScaleVCore) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the PowerBI Dedicated resource.
func (o AutoScaleVCoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoScaleVCore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment state of an auto scale v-core resource. The provisioningState is to indicate states for resource provisioning.
func (o AutoScaleVCoreOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoScaleVCore) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the auto scale v-core resource.
func (o AutoScaleVCoreOutput) Sku() AutoScaleVCoreSkuResponseOutput {
	return o.ApplyT(func(v *AutoScaleVCore) AutoScaleVCoreSkuResponseOutput { return v.Sku }).(AutoScaleVCoreSkuResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o AutoScaleVCoreOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v *AutoScaleVCore) SystemDataResponsePtrOutput { return v.SystemData }).(SystemDataResponsePtrOutput)
}

// Key-value pairs of additional resource provisioning properties.
func (o AutoScaleVCoreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AutoScaleVCore) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the PowerBI Dedicated resource.
func (o AutoScaleVCoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoScaleVCore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoScaleVCoreOutput{})
}

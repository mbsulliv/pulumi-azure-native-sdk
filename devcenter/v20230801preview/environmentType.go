// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents an environment type.
type EnvironmentType struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnvironmentType registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentType(ctx *pulumi.Context,
	name string, args *EnvironmentTypeArgs, opts ...pulumi.ResourceOption) (*EnvironmentType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230401:EnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20231001preview:EnvironmentType"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnvironmentType
	err := ctx.RegisterResource("azure-native:devcenter/v20230801preview:EnvironmentType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentType gets an existing EnvironmentType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentTypeState, opts ...pulumi.ResourceOption) (*EnvironmentType, error) {
	var resource EnvironmentType
	err := ctx.ReadResource("azure-native:devcenter/v20230801preview:EnvironmentType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentType resources.
type environmentTypeState struct {
}

type EnvironmentTypeState struct {
}

func (EnvironmentTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentTypeState)(nil)).Elem()
}

type environmentTypeArgs struct {
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the environment type.
	EnvironmentTypeName *string `pulumi:"environmentTypeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EnvironmentType resource.
type EnvironmentTypeArgs struct {
	// The name of the devcenter.
	DevCenterName pulumi.StringInput
	// The name of the environment type.
	EnvironmentTypeName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (EnvironmentTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentTypeArgs)(nil)).Elem()
}

type EnvironmentTypeInput interface {
	pulumi.Input

	ToEnvironmentTypeOutput() EnvironmentTypeOutput
	ToEnvironmentTypeOutputWithContext(ctx context.Context) EnvironmentTypeOutput
}

func (*EnvironmentType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentType)(nil)).Elem()
}

func (i *EnvironmentType) ToEnvironmentTypeOutput() EnvironmentTypeOutput {
	return i.ToEnvironmentTypeOutputWithContext(context.Background())
}

func (i *EnvironmentType) ToEnvironmentTypeOutputWithContext(ctx context.Context) EnvironmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentTypeOutput)
}

func (i *EnvironmentType) ToOutput(ctx context.Context) pulumix.Output[*EnvironmentType] {
	return pulumix.Output[*EnvironmentType]{
		OutputState: i.ToEnvironmentTypeOutputWithContext(ctx).OutputState,
	}
}

type EnvironmentTypeOutput struct{ *pulumi.OutputState }

func (EnvironmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentType)(nil)).Elem()
}

func (o EnvironmentTypeOutput) ToEnvironmentTypeOutput() EnvironmentTypeOutput {
	return o
}

func (o EnvironmentTypeOutput) ToEnvironmentTypeOutputWithContext(ctx context.Context) EnvironmentTypeOutput {
	return o
}

func (o EnvironmentTypeOutput) ToOutput(ctx context.Context) pulumix.Output[*EnvironmentType] {
	return pulumix.Output[*EnvironmentType]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o EnvironmentTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o EnvironmentTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EnvironmentTypeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnvironmentType) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o EnvironmentTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnvironmentType) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EnvironmentTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentTypeOutput{})
}
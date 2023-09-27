// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// MachinePool represents a MachinePool
type MachinePool struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name      pulumi.StringOutput    `pulumi:"name"`
	Resources pulumi.StringPtrOutput `pulumi:"resources"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMachinePool registers a new resource with the given unique name, arguments, and options.
func NewMachinePool(ctx *pulumi.Context,
	name string, args *MachinePoolArgs, opts ...pulumi.ResourceOption) (*MachinePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:redhatopenshift:MachinePool"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20220904:MachinePool"),
		},
		{
			Type: pulumi.String("azure-native:redhatopenshift/v20230401:MachinePool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MachinePool
	err := ctx.RegisterResource("azure-native:redhatopenshift/v20230701preview:MachinePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachinePool gets an existing MachinePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachinePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachinePoolState, opts ...pulumi.ResourceOption) (*MachinePool, error) {
	var resource MachinePool
	err := ctx.ReadResource("azure-native:redhatopenshift/v20230701preview:MachinePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachinePool resources.
type machinePoolState struct {
}

type MachinePoolState struct {
}

func (MachinePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*machinePoolState)(nil)).Elem()
}

type machinePoolArgs struct {
	// The name of the MachinePool resource.
	ChildResourceName *string `pulumi:"childResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName string  `pulumi:"resourceName"`
	Resources    *string `pulumi:"resources"`
}

// The set of arguments for constructing a MachinePool resource.
type MachinePoolArgs struct {
	// The name of the MachinePool resource.
	ChildResourceName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the OpenShift cluster resource.
	ResourceName pulumi.StringInput
	Resources    pulumi.StringPtrInput
}

func (MachinePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machinePoolArgs)(nil)).Elem()
}

type MachinePoolInput interface {
	pulumi.Input

	ToMachinePoolOutput() MachinePoolOutput
	ToMachinePoolOutputWithContext(ctx context.Context) MachinePoolOutput
}

func (*MachinePool) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePool)(nil)).Elem()
}

func (i *MachinePool) ToMachinePoolOutput() MachinePoolOutput {
	return i.ToMachinePoolOutputWithContext(context.Background())
}

func (i *MachinePool) ToMachinePoolOutputWithContext(ctx context.Context) MachinePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePoolOutput)
}

func (i *MachinePool) ToOutput(ctx context.Context) pulumix.Output[*MachinePool] {
	return pulumix.Output[*MachinePool]{
		OutputState: i.ToMachinePoolOutputWithContext(ctx).OutputState,
	}
}

type MachinePoolOutput struct{ *pulumi.OutputState }

func (MachinePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePool)(nil)).Elem()
}

func (o MachinePoolOutput) ToMachinePoolOutput() MachinePoolOutput {
	return o
}

func (o MachinePoolOutput) ToMachinePoolOutputWithContext(ctx context.Context) MachinePoolOutput {
	return o
}

func (o MachinePoolOutput) ToOutput(ctx context.Context) pulumix.Output[*MachinePool] {
	return pulumix.Output[*MachinePool]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o MachinePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachinePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachinePoolOutput) Resources() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePool) pulumi.StringPtrOutput { return v.Resources }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MachinePoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MachinePool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MachinePoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachinePool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachinePoolOutput{})
}

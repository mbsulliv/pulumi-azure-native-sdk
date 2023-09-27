// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Storage resource for connectedEnvironment.
type ConnectedEnvironmentsStorage struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Storage properties
	Properties ConnectedEnvironmentStorageResponsePropertiesOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectedEnvironmentsStorage registers a new resource with the given unique name, arguments, and options.
func NewConnectedEnvironmentsStorage(ctx *pulumi.Context,
	name string, args *ConnectedEnvironmentsStorageArgs, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectedEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedEnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ConnectedEnvironmentsStorage"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ConnectedEnvironmentsStorage"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ConnectedEnvironmentsStorage"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:ConnectedEnvironmentsStorage"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230401preview:ConnectedEnvironmentsStorage"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230502preview:ConnectedEnvironmentsStorage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectedEnvironmentsStorage
	err := ctx.RegisterResource("azure-native:app/v20230501:ConnectedEnvironmentsStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectedEnvironmentsStorage gets an existing ConnectedEnvironmentsStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectedEnvironmentsStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedEnvironmentsStorageState, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsStorage, error) {
	var resource ConnectedEnvironmentsStorage
	err := ctx.ReadResource("azure-native:app/v20230501:ConnectedEnvironmentsStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectedEnvironmentsStorage resources.
type connectedEnvironmentsStorageState struct {
}

type ConnectedEnvironmentsStorageState struct {
}

func (ConnectedEnvironmentsStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsStorageState)(nil)).Elem()
}

type connectedEnvironmentsStorageArgs struct {
	// Name of the Environment.
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	// Storage properties
	Properties *ConnectedEnvironmentStorageProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the storage.
	StorageName *string `pulumi:"storageName"`
}

// The set of arguments for constructing a ConnectedEnvironmentsStorage resource.
type ConnectedEnvironmentsStorageArgs struct {
	// Name of the Environment.
	ConnectedEnvironmentName pulumi.StringInput
	// Storage properties
	Properties ConnectedEnvironmentStoragePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of the storage.
	StorageName pulumi.StringPtrInput
}

func (ConnectedEnvironmentsStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsStorageArgs)(nil)).Elem()
}

type ConnectedEnvironmentsStorageInput interface {
	pulumi.Input

	ToConnectedEnvironmentsStorageOutput() ConnectedEnvironmentsStorageOutput
	ToConnectedEnvironmentsStorageOutputWithContext(ctx context.Context) ConnectedEnvironmentsStorageOutput
}

func (*ConnectedEnvironmentsStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsStorage)(nil)).Elem()
}

func (i *ConnectedEnvironmentsStorage) ToConnectedEnvironmentsStorageOutput() ConnectedEnvironmentsStorageOutput {
	return i.ToConnectedEnvironmentsStorageOutputWithContext(context.Background())
}

func (i *ConnectedEnvironmentsStorage) ToConnectedEnvironmentsStorageOutputWithContext(ctx context.Context) ConnectedEnvironmentsStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedEnvironmentsStorageOutput)
}

func (i *ConnectedEnvironmentsStorage) ToOutput(ctx context.Context) pulumix.Output[*ConnectedEnvironmentsStorage] {
	return pulumix.Output[*ConnectedEnvironmentsStorage]{
		OutputState: i.ToConnectedEnvironmentsStorageOutputWithContext(ctx).OutputState,
	}
}

type ConnectedEnvironmentsStorageOutput struct{ *pulumi.OutputState }

func (ConnectedEnvironmentsStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsStorage)(nil)).Elem()
}

func (o ConnectedEnvironmentsStorageOutput) ToConnectedEnvironmentsStorageOutput() ConnectedEnvironmentsStorageOutput {
	return o
}

func (o ConnectedEnvironmentsStorageOutput) ToConnectedEnvironmentsStorageOutputWithContext(ctx context.Context) ConnectedEnvironmentsStorageOutput {
	return o
}

func (o ConnectedEnvironmentsStorageOutput) ToOutput(ctx context.Context) pulumix.Output[*ConnectedEnvironmentsStorage] {
	return pulumix.Output[*ConnectedEnvironmentsStorage]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o ConnectedEnvironmentsStorageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsStorage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Storage properties
func (o ConnectedEnvironmentsStorageOutput) Properties() ConnectedEnvironmentStorageResponsePropertiesOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsStorage) ConnectedEnvironmentStorageResponsePropertiesOutput {
		return v.Properties
	}).(ConnectedEnvironmentStorageResponsePropertiesOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConnectedEnvironmentsStorageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsStorage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConnectedEnvironmentsStorageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsStorage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedEnvironmentsStorageOutput{})
}

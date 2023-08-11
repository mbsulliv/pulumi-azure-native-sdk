// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagetasks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Storage Task.
// Azure REST API version: 2023-01-01.
type StorageTask struct {
	pulumi.CustomResourceState

	// The storage task action that is executed
	Action StorageTaskActionResponseOutput `pulumi:"action"`
	// The creation date and time of the storage task in UTC.
	CreationTimeInUtc pulumi.StringOutput `pulumi:"creationTimeInUtc"`
	// Text that describes the purpose of the storage task
	Description pulumi.StringOutput `pulumi:"description"`
	// Storage Task is enabled when set to true and disabled when set to false
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The managed service identity of the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Represents the provisioning state of the storage task.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Storage task version.
	TaskVersion pulumi.Float64Output `pulumi:"taskVersion"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageTask registers a new resource with the given unique name, arguments, and options.
func NewStorageTask(ctx *pulumi.Context,
	name string, args *StorageTaskArgs, opts ...pulumi.ResourceOption) (*StorageTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagetasks/v20230101:StorageTask"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageTask
	err := ctx.RegisterResource("azure-native:storagetasks:StorageTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageTask gets an existing StorageTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageTaskState, opts ...pulumi.ResourceOption) (*StorageTask, error) {
	var resource StorageTask
	err := ctx.ReadResource("azure-native:storagetasks:StorageTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageTask resources.
type storageTaskState struct {
}

type StorageTaskState struct {
}

func (StorageTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageTaskState)(nil)).Elem()
}

type storageTaskArgs struct {
	// The storage task action that is executed
	Action StorageTaskAction `pulumi:"action"`
	// Text that describes the purpose of the storage task
	Description string `pulumi:"description"`
	// Storage Task is enabled when set to true and disabled when set to false
	Enabled bool `pulumi:"enabled"`
	// The managed service identity of the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the storage task within the specified resource group. Storage task names must be between 3 and 18 characters in length and use numbers and lower-case letters only.
	StorageTaskName *string `pulumi:"storageTaskName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageTask resource.
type StorageTaskArgs struct {
	// The storage task action that is executed
	Action StorageTaskActionInput
	// Text that describes the purpose of the storage task
	Description pulumi.StringInput
	// Storage Task is enabled when set to true and disabled when set to false
	Enabled pulumi.BoolInput
	// The managed service identity of the resource.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the storage task within the specified resource group. Storage task names must be between 3 and 18 characters in length and use numbers and lower-case letters only.
	StorageTaskName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (StorageTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageTaskArgs)(nil)).Elem()
}

type StorageTaskInput interface {
	pulumi.Input

	ToStorageTaskOutput() StorageTaskOutput
	ToStorageTaskOutputWithContext(ctx context.Context) StorageTaskOutput
}

func (*StorageTask) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageTask)(nil)).Elem()
}

func (i *StorageTask) ToStorageTaskOutput() StorageTaskOutput {
	return i.ToStorageTaskOutputWithContext(context.Background())
}

func (i *StorageTask) ToStorageTaskOutputWithContext(ctx context.Context) StorageTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageTaskOutput)
}

type StorageTaskOutput struct{ *pulumi.OutputState }

func (StorageTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageTask)(nil)).Elem()
}

func (o StorageTaskOutput) ToStorageTaskOutput() StorageTaskOutput {
	return o
}

func (o StorageTaskOutput) ToStorageTaskOutputWithContext(ctx context.Context) StorageTaskOutput {
	return o
}

// The storage task action that is executed
func (o StorageTaskOutput) Action() StorageTaskActionResponseOutput {
	return o.ApplyT(func(v *StorageTask) StorageTaskActionResponseOutput { return v.Action }).(StorageTaskActionResponseOutput)
}

// The creation date and time of the storage task in UTC.
func (o StorageTaskOutput) CreationTimeInUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTask) pulumi.StringOutput { return v.CreationTimeInUtc }).(pulumi.StringOutput)
}

// Text that describes the purpose of the storage task
func (o StorageTaskOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTask) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Storage Task is enabled when set to true and disabled when set to false
func (o StorageTaskOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *StorageTask) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The managed service identity of the resource.
func (o StorageTaskOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *StorageTask) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o StorageTaskOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTask) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o StorageTaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTask) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Represents the provisioning state of the storage task.
func (o StorageTaskOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTask) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o StorageTaskOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageTask) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o StorageTaskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageTask) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Storage task version.
func (o StorageTaskOutput) TaskVersion() pulumi.Float64Output {
	return o.ApplyT(func(v *StorageTask) pulumi.Float64Output { return v.TaskVersion }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StorageTaskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageTask) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageTaskOutput{})
}

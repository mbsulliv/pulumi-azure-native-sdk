// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The lock information.
// Azure REST API version: 2020-05-01. Prior API version in Azure Native 1.x: 2017-04-01.
//
// Other available API versions: 2015-01-01.
type ManagementLockAtResourceGroupLevel struct {
	pulumi.CustomResourceState

	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level pulumi.StringOutput `pulumi:"level"`
	// The name of the lock.
	Name pulumi.StringOutput `pulumi:"name"`
	// Notes about the lock. Maximum of 512 characters.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The owners of the lock.
	Owners ManagementLockOwnerResponseArrayOutput `pulumi:"owners"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The resource type of the lock - Microsoft.Authorization/locks.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementLockAtResourceGroupLevel registers a new resource with the given unique name, arguments, and options.
func NewManagementLockAtResourceGroupLevel(ctx *pulumi.Context,
	name string, args *ManagementLockAtResourceGroupLevelArgs, opts ...pulumi.ResourceOption) (*ManagementLockAtResourceGroupLevel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Level == nil {
		return nil, errors.New("invalid value for required argument 'Level'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20150101:ManagementLockAtResourceGroupLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160901:ManagementLockAtResourceGroupLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20170401:ManagementLockAtResourceGroupLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200501:ManagementLockAtResourceGroupLevel"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagementLockAtResourceGroupLevel
	err := ctx.RegisterResource("azure-native:authorization:ManagementLockAtResourceGroupLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementLockAtResourceGroupLevel gets an existing ManagementLockAtResourceGroupLevel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementLockAtResourceGroupLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementLockAtResourceGroupLevelState, opts ...pulumi.ResourceOption) (*ManagementLockAtResourceGroupLevel, error) {
	var resource ManagementLockAtResourceGroupLevel
	err := ctx.ReadResource("azure-native:authorization:ManagementLockAtResourceGroupLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementLockAtResourceGroupLevel resources.
type managementLockAtResourceGroupLevelState struct {
}

type ManagementLockAtResourceGroupLevelState struct {
}

func (ManagementLockAtResourceGroupLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtResourceGroupLevelState)(nil)).Elem()
}

type managementLockAtResourceGroupLevelArgs struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level string `pulumi:"level"`
	// The lock name. The lock name can be a maximum of 260 characters. It cannot contain <, > %, &, :, \, ?, /, or any control characters.
	LockName *string `pulumi:"lockName"`
	// Notes about the lock. Maximum of 512 characters.
	Notes *string `pulumi:"notes"`
	// The owners of the lock.
	Owners []ManagementLockOwner `pulumi:"owners"`
	// The name of the resource group to lock.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagementLockAtResourceGroupLevel resource.
type ManagementLockAtResourceGroupLevelArgs struct {
	// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
	Level pulumi.StringInput
	// The lock name. The lock name can be a maximum of 260 characters. It cannot contain <, > %, &, :, \, ?, /, or any control characters.
	LockName pulumi.StringPtrInput
	// Notes about the lock. Maximum of 512 characters.
	Notes pulumi.StringPtrInput
	// The owners of the lock.
	Owners ManagementLockOwnerArrayInput
	// The name of the resource group to lock.
	ResourceGroupName pulumi.StringInput
}

func (ManagementLockAtResourceGroupLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtResourceGroupLevelArgs)(nil)).Elem()
}

type ManagementLockAtResourceGroupLevelInput interface {
	pulumi.Input

	ToManagementLockAtResourceGroupLevelOutput() ManagementLockAtResourceGroupLevelOutput
	ToManagementLockAtResourceGroupLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceGroupLevelOutput
}

func (*ManagementLockAtResourceGroupLevel) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLockAtResourceGroupLevel)(nil)).Elem()
}

func (i *ManagementLockAtResourceGroupLevel) ToManagementLockAtResourceGroupLevelOutput() ManagementLockAtResourceGroupLevelOutput {
	return i.ToManagementLockAtResourceGroupLevelOutputWithContext(context.Background())
}

func (i *ManagementLockAtResourceGroupLevel) ToManagementLockAtResourceGroupLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceGroupLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockAtResourceGroupLevelOutput)
}

type ManagementLockAtResourceGroupLevelOutput struct{ *pulumi.OutputState }

func (ManagementLockAtResourceGroupLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLockAtResourceGroupLevel)(nil)).Elem()
}

func (o ManagementLockAtResourceGroupLevelOutput) ToManagementLockAtResourceGroupLevelOutput() ManagementLockAtResourceGroupLevelOutput {
	return o
}

func (o ManagementLockAtResourceGroupLevelOutput) ToManagementLockAtResourceGroupLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceGroupLevelOutput {
	return o
}

// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
func (o ManagementLockAtResourceGroupLevelOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceGroupLevel) pulumi.StringOutput { return v.Level }).(pulumi.StringOutput)
}

// The name of the lock.
func (o ManagementLockAtResourceGroupLevelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceGroupLevel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Notes about the lock. Maximum of 512 characters.
func (o ManagementLockAtResourceGroupLevelOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceGroupLevel) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The owners of the lock.
func (o ManagementLockAtResourceGroupLevelOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceGroupLevel) ManagementLockOwnerResponseArrayOutput { return v.Owners }).(ManagementLockOwnerResponseArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ManagementLockAtResourceGroupLevelOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceGroupLevel) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type of the lock - Microsoft.Authorization/locks.
func (o ManagementLockAtResourceGroupLevelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceGroupLevel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementLockAtResourceGroupLevelOutput{})
}

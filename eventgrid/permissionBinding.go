// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Permission binding resource.
// Azure REST API version: 2023-06-01-preview.
//
// Other available API versions: 2023-12-15-preview.
type PermissionBinding struct {
	pulumi.CustomResourceState

	// The name of the client group resource that the permission is bound to.
	// The client group needs to be a resource under the same namespace the permission binding is a part of.
	ClientGroupName pulumi.StringPtrOutput `pulumi:"clientGroupName"`
	// Description for the Permission Binding resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The allowed permission.
	Permission pulumi.StringPtrOutput `pulumi:"permission"`
	// Provisioning state of the PermissionBinding resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system metadata relating to the PermissionBinding resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The name of the Topic Space resource that the permission is bound to.
	// The Topic space needs to be a resource under the same namespace the permission binding is a part of.
	TopicSpaceName pulumi.StringPtrOutput `pulumi:"topicSpaceName"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPermissionBinding registers a new resource with the given unique name, arguments, and options.
func NewPermissionBinding(ctx *pulumi.Context,
	name string, args *PermissionBindingArgs, opts ...pulumi.ResourceOption) (*PermissionBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid/v20230601preview:PermissionBinding"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20231215preview:PermissionBinding"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PermissionBinding
	err := ctx.RegisterResource("azure-native:eventgrid:PermissionBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissionBinding gets an existing PermissionBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissionBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionBindingState, opts ...pulumi.ResourceOption) (*PermissionBinding, error) {
	var resource PermissionBinding
	err := ctx.ReadResource("azure-native:eventgrid:PermissionBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PermissionBinding resources.
type permissionBindingState struct {
}

type PermissionBindingState struct {
}

func (PermissionBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionBindingState)(nil)).Elem()
}

type permissionBindingArgs struct {
	// The name of the client group resource that the permission is bound to.
	// The client group needs to be a resource under the same namespace the permission binding is a part of.
	ClientGroupName *string `pulumi:"clientGroupName"`
	// Description for the Permission Binding resource.
	Description *string `pulumi:"description"`
	// Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The allowed permission.
	Permission *string `pulumi:"permission"`
	// The permission binding name.
	PermissionBindingName *string `pulumi:"permissionBindingName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Topic Space resource that the permission is bound to.
	// The Topic space needs to be a resource under the same namespace the permission binding is a part of.
	TopicSpaceName *string `pulumi:"topicSpaceName"`
}

// The set of arguments for constructing a PermissionBinding resource.
type PermissionBindingArgs struct {
	// The name of the client group resource that the permission is bound to.
	// The client group needs to be a resource under the same namespace the permission binding is a part of.
	ClientGroupName pulumi.StringPtrInput
	// Description for the Permission Binding resource.
	Description pulumi.StringPtrInput
	// Name of the namespace.
	NamespaceName pulumi.StringInput
	// The allowed permission.
	Permission pulumi.StringPtrInput
	// The permission binding name.
	PermissionBindingName pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the Topic Space resource that the permission is bound to.
	// The Topic space needs to be a resource under the same namespace the permission binding is a part of.
	TopicSpaceName pulumi.StringPtrInput
}

func (PermissionBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionBindingArgs)(nil)).Elem()
}

type PermissionBindingInput interface {
	pulumi.Input

	ToPermissionBindingOutput() PermissionBindingOutput
	ToPermissionBindingOutputWithContext(ctx context.Context) PermissionBindingOutput
}

func (*PermissionBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionBinding)(nil)).Elem()
}

func (i *PermissionBinding) ToPermissionBindingOutput() PermissionBindingOutput {
	return i.ToPermissionBindingOutputWithContext(context.Background())
}

func (i *PermissionBinding) ToPermissionBindingOutputWithContext(ctx context.Context) PermissionBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionBindingOutput)
}

type PermissionBindingOutput struct{ *pulumi.OutputState }

func (PermissionBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionBinding)(nil)).Elem()
}

func (o PermissionBindingOutput) ToPermissionBindingOutput() PermissionBindingOutput {
	return o
}

func (o PermissionBindingOutput) ToPermissionBindingOutputWithContext(ctx context.Context) PermissionBindingOutput {
	return o
}

// The name of the client group resource that the permission is bound to.
// The client group needs to be a resource under the same namespace the permission binding is a part of.
func (o PermissionBindingOutput) ClientGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PermissionBinding) pulumi.StringPtrOutput { return v.ClientGroupName }).(pulumi.StringPtrOutput)
}

// Description for the Permission Binding resource.
func (o PermissionBindingOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PermissionBinding) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o PermissionBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PermissionBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The allowed permission.
func (o PermissionBindingOutput) Permission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PermissionBinding) pulumi.StringPtrOutput { return v.Permission }).(pulumi.StringPtrOutput)
}

// Provisioning state of the PermissionBinding resource.
func (o PermissionBindingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PermissionBinding) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to the PermissionBinding resource.
func (o PermissionBindingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PermissionBinding) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The name of the Topic Space resource that the permission is bound to.
// The Topic space needs to be a resource under the same namespace the permission binding is a part of.
func (o PermissionBindingOutput) TopicSpaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PermissionBinding) pulumi.StringPtrOutput { return v.TopicSpaceName }).(pulumi.StringPtrOutput)
}

// Type of the resource.
func (o PermissionBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PermissionBinding) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PermissionBindingOutput{})
}

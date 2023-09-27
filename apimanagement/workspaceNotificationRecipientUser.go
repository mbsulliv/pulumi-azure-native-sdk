// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Recipient User details.
// Azure REST API version: 2022-09-01-preview.
type WorkspaceNotificationRecipientUser struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// API Management UserId subscribed to notification.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewWorkspaceNotificationRecipientUser registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceNotificationRecipientUser(ctx *pulumi.Context,
	name string, args *WorkspaceNotificationRecipientUserArgs, opts ...pulumi.ResourceOption) (*WorkspaceNotificationRecipientUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NotificationName == nil {
		return nil, errors.New("invalid value for required argument 'NotificationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceNotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceNotificationRecipientUser"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceNotificationRecipientUser
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspaceNotificationRecipientUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceNotificationRecipientUser gets an existing WorkspaceNotificationRecipientUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceNotificationRecipientUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceNotificationRecipientUserState, opts ...pulumi.ResourceOption) (*WorkspaceNotificationRecipientUser, error) {
	var resource WorkspaceNotificationRecipientUser
	err := ctx.ReadResource("azure-native:apimanagement:WorkspaceNotificationRecipientUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceNotificationRecipientUser resources.
type workspaceNotificationRecipientUserState struct {
}

type WorkspaceNotificationRecipientUserState struct {
}

func (WorkspaceNotificationRecipientUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceNotificationRecipientUserState)(nil)).Elem()
}

type workspaceNotificationRecipientUserArgs struct {
	// Notification Name Identifier.
	NotificationName string `pulumi:"notificationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId *string `pulumi:"userId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceNotificationRecipientUser resource.
type WorkspaceNotificationRecipientUserArgs struct {
	// Notification Name Identifier.
	NotificationName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// User identifier. Must be unique in the current API Management service instance.
	UserId pulumi.StringPtrInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceNotificationRecipientUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceNotificationRecipientUserArgs)(nil)).Elem()
}

type WorkspaceNotificationRecipientUserInput interface {
	pulumi.Input

	ToWorkspaceNotificationRecipientUserOutput() WorkspaceNotificationRecipientUserOutput
	ToWorkspaceNotificationRecipientUserOutputWithContext(ctx context.Context) WorkspaceNotificationRecipientUserOutput
}

func (*WorkspaceNotificationRecipientUser) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceNotificationRecipientUser)(nil)).Elem()
}

func (i *WorkspaceNotificationRecipientUser) ToWorkspaceNotificationRecipientUserOutput() WorkspaceNotificationRecipientUserOutput {
	return i.ToWorkspaceNotificationRecipientUserOutputWithContext(context.Background())
}

func (i *WorkspaceNotificationRecipientUser) ToWorkspaceNotificationRecipientUserOutputWithContext(ctx context.Context) WorkspaceNotificationRecipientUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceNotificationRecipientUserOutput)
}

func (i *WorkspaceNotificationRecipientUser) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceNotificationRecipientUser] {
	return pulumix.Output[*WorkspaceNotificationRecipientUser]{
		OutputState: i.ToWorkspaceNotificationRecipientUserOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceNotificationRecipientUserOutput struct{ *pulumi.OutputState }

func (WorkspaceNotificationRecipientUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceNotificationRecipientUser)(nil)).Elem()
}

func (o WorkspaceNotificationRecipientUserOutput) ToWorkspaceNotificationRecipientUserOutput() WorkspaceNotificationRecipientUserOutput {
	return o
}

func (o WorkspaceNotificationRecipientUserOutput) ToWorkspaceNotificationRecipientUserOutputWithContext(ctx context.Context) WorkspaceNotificationRecipientUserOutput {
	return o
}

func (o WorkspaceNotificationRecipientUserOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceNotificationRecipientUser] {
	return pulumix.Output[*WorkspaceNotificationRecipientUser]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o WorkspaceNotificationRecipientUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceNotificationRecipientUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceNotificationRecipientUserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceNotificationRecipientUser) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// API Management UserId subscribed to notification.
func (o WorkspaceNotificationRecipientUserOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceNotificationRecipientUser) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceNotificationRecipientUserOutput{})
}

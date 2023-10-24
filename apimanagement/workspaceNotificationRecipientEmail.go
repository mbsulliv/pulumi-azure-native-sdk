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

// Recipient Email details.
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
type WorkspaceNotificationRecipientEmail struct {
	pulumi.CustomResourceState

	// User Email subscribed to notification.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceNotificationRecipientEmail registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceNotificationRecipientEmail(ctx *pulumi.Context,
	name string, args *WorkspaceNotificationRecipientEmailArgs, opts ...pulumi.ResourceOption) (*WorkspaceNotificationRecipientEmail, error) {
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
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceNotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceNotificationRecipientEmail"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceNotificationRecipientEmail
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspaceNotificationRecipientEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceNotificationRecipientEmail gets an existing WorkspaceNotificationRecipientEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceNotificationRecipientEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceNotificationRecipientEmailState, opts ...pulumi.ResourceOption) (*WorkspaceNotificationRecipientEmail, error) {
	var resource WorkspaceNotificationRecipientEmail
	err := ctx.ReadResource("azure-native:apimanagement:WorkspaceNotificationRecipientEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceNotificationRecipientEmail resources.
type workspaceNotificationRecipientEmailState struct {
}

type WorkspaceNotificationRecipientEmailState struct {
}

func (WorkspaceNotificationRecipientEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceNotificationRecipientEmailState)(nil)).Elem()
}

type workspaceNotificationRecipientEmailArgs struct {
	// Email identifier.
	Email *string `pulumi:"email"`
	// Notification Name Identifier.
	NotificationName string `pulumi:"notificationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceNotificationRecipientEmail resource.
type WorkspaceNotificationRecipientEmailArgs struct {
	// Email identifier.
	Email pulumi.StringPtrInput
	// Notification Name Identifier.
	NotificationName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceNotificationRecipientEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceNotificationRecipientEmailArgs)(nil)).Elem()
}

type WorkspaceNotificationRecipientEmailInput interface {
	pulumi.Input

	ToWorkspaceNotificationRecipientEmailOutput() WorkspaceNotificationRecipientEmailOutput
	ToWorkspaceNotificationRecipientEmailOutputWithContext(ctx context.Context) WorkspaceNotificationRecipientEmailOutput
}

func (*WorkspaceNotificationRecipientEmail) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceNotificationRecipientEmail)(nil)).Elem()
}

func (i *WorkspaceNotificationRecipientEmail) ToWorkspaceNotificationRecipientEmailOutput() WorkspaceNotificationRecipientEmailOutput {
	return i.ToWorkspaceNotificationRecipientEmailOutputWithContext(context.Background())
}

func (i *WorkspaceNotificationRecipientEmail) ToWorkspaceNotificationRecipientEmailOutputWithContext(ctx context.Context) WorkspaceNotificationRecipientEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceNotificationRecipientEmailOutput)
}

func (i *WorkspaceNotificationRecipientEmail) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceNotificationRecipientEmail] {
	return pulumix.Output[*WorkspaceNotificationRecipientEmail]{
		OutputState: i.ToWorkspaceNotificationRecipientEmailOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceNotificationRecipientEmailOutput struct{ *pulumi.OutputState }

func (WorkspaceNotificationRecipientEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceNotificationRecipientEmail)(nil)).Elem()
}

func (o WorkspaceNotificationRecipientEmailOutput) ToWorkspaceNotificationRecipientEmailOutput() WorkspaceNotificationRecipientEmailOutput {
	return o
}

func (o WorkspaceNotificationRecipientEmailOutput) ToWorkspaceNotificationRecipientEmailOutputWithContext(ctx context.Context) WorkspaceNotificationRecipientEmailOutput {
	return o
}

func (o WorkspaceNotificationRecipientEmailOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceNotificationRecipientEmail] {
	return pulumix.Output[*WorkspaceNotificationRecipientEmail]{
		OutputState: o.OutputState,
	}
}

// User Email subscribed to notification.
func (o WorkspaceNotificationRecipientEmailOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceNotificationRecipientEmail) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkspaceNotificationRecipientEmailOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceNotificationRecipientEmail) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceNotificationRecipientEmailOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceNotificationRecipientEmail) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceNotificationRecipientEmailOutput{})
}

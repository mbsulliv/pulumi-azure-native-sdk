// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Recipient User details.
type NotificationRecipientUser struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// API Management UserId subscribed to notification.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewNotificationRecipientUser registers a new resource with the given unique name, arguments, and options.
func NewNotificationRecipientUser(ctx *pulumi.Context,
	name string, args *NotificationRecipientUserArgs, opts ...pulumi.ResourceOption) (*NotificationRecipientUser, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:NotificationRecipientUser"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:NotificationRecipientUser"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NotificationRecipientUser
	err := ctx.RegisterResource("azure-native:apimanagement/v20220801:NotificationRecipientUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationRecipientUser gets an existing NotificationRecipientUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationRecipientUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationRecipientUserState, opts ...pulumi.ResourceOption) (*NotificationRecipientUser, error) {
	var resource NotificationRecipientUser
	err := ctx.ReadResource("azure-native:apimanagement/v20220801:NotificationRecipientUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationRecipientUser resources.
type notificationRecipientUserState struct {
}

type NotificationRecipientUserState struct {
}

func (NotificationRecipientUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientUserState)(nil)).Elem()
}

type notificationRecipientUserArgs struct {
	// Notification Name Identifier.
	NotificationName string `pulumi:"notificationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a NotificationRecipientUser resource.
type NotificationRecipientUserArgs struct {
	// Notification Name Identifier.
	NotificationName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// User identifier. Must be unique in the current API Management service instance.
	UserId pulumi.StringPtrInput
}

func (NotificationRecipientUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientUserArgs)(nil)).Elem()
}

type NotificationRecipientUserInput interface {
	pulumi.Input

	ToNotificationRecipientUserOutput() NotificationRecipientUserOutput
	ToNotificationRecipientUserOutputWithContext(ctx context.Context) NotificationRecipientUserOutput
}

func (*NotificationRecipientUser) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRecipientUser)(nil)).Elem()
}

func (i *NotificationRecipientUser) ToNotificationRecipientUserOutput() NotificationRecipientUserOutput {
	return i.ToNotificationRecipientUserOutputWithContext(context.Background())
}

func (i *NotificationRecipientUser) ToNotificationRecipientUserOutputWithContext(ctx context.Context) NotificationRecipientUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRecipientUserOutput)
}

func (i *NotificationRecipientUser) ToOutput(ctx context.Context) pulumix.Output[*NotificationRecipientUser] {
	return pulumix.Output[*NotificationRecipientUser]{
		OutputState: i.ToNotificationRecipientUserOutputWithContext(ctx).OutputState,
	}
}

type NotificationRecipientUserOutput struct{ *pulumi.OutputState }

func (NotificationRecipientUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRecipientUser)(nil)).Elem()
}

func (o NotificationRecipientUserOutput) ToNotificationRecipientUserOutput() NotificationRecipientUserOutput {
	return o
}

func (o NotificationRecipientUserOutput) ToNotificationRecipientUserOutputWithContext(ctx context.Context) NotificationRecipientUserOutput {
	return o
}

func (o NotificationRecipientUserOutput) ToOutput(ctx context.Context) pulumix.Output[*NotificationRecipientUser] {
	return pulumix.Output[*NotificationRecipientUser]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o NotificationRecipientUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRecipientUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NotificationRecipientUserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRecipientUser) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// API Management UserId subscribed to notification.
func (o NotificationRecipientUserOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationRecipientUser) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationRecipientUserOutput{})
}

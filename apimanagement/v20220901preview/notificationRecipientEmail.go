// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Recipient Email details.
type NotificationRecipientEmail struct {
	pulumi.CustomResourceState

	// User Email subscribed to notification.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNotificationRecipientEmail registers a new resource with the given unique name, arguments, and options.
func NewNotificationRecipientEmail(ctx *pulumi.Context,
	name string, args *NotificationRecipientEmailArgs, opts ...pulumi.ResourceOption) (*NotificationRecipientEmail, error) {
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
			Type: pulumi.String("azure-native:apimanagement:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:NotificationRecipientEmail"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:NotificationRecipientEmail"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NotificationRecipientEmail
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:NotificationRecipientEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationRecipientEmail gets an existing NotificationRecipientEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationRecipientEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationRecipientEmailState, opts ...pulumi.ResourceOption) (*NotificationRecipientEmail, error) {
	var resource NotificationRecipientEmail
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:NotificationRecipientEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationRecipientEmail resources.
type notificationRecipientEmailState struct {
}

type NotificationRecipientEmailState struct {
}

func (NotificationRecipientEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientEmailState)(nil)).Elem()
}

type notificationRecipientEmailArgs struct {
	// Email identifier.
	Email *string `pulumi:"email"`
	// Notification Name Identifier.
	NotificationName string `pulumi:"notificationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a NotificationRecipientEmail resource.
type NotificationRecipientEmailArgs struct {
	// Email identifier.
	Email pulumi.StringPtrInput
	// Notification Name Identifier.
	NotificationName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (NotificationRecipientEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRecipientEmailArgs)(nil)).Elem()
}

type NotificationRecipientEmailInput interface {
	pulumi.Input

	ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput
	ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput
}

func (*NotificationRecipientEmail) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRecipientEmail)(nil)).Elem()
}

func (i *NotificationRecipientEmail) ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput {
	return i.ToNotificationRecipientEmailOutputWithContext(context.Background())
}

func (i *NotificationRecipientEmail) ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRecipientEmailOutput)
}

func (i *NotificationRecipientEmail) ToOutput(ctx context.Context) pulumix.Output[*NotificationRecipientEmail] {
	return pulumix.Output[*NotificationRecipientEmail]{
		OutputState: i.ToNotificationRecipientEmailOutputWithContext(ctx).OutputState,
	}
}

type NotificationRecipientEmailOutput struct{ *pulumi.OutputState }

func (NotificationRecipientEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRecipientEmail)(nil)).Elem()
}

func (o NotificationRecipientEmailOutput) ToNotificationRecipientEmailOutput() NotificationRecipientEmailOutput {
	return o
}

func (o NotificationRecipientEmailOutput) ToNotificationRecipientEmailOutputWithContext(ctx context.Context) NotificationRecipientEmailOutput {
	return o
}

func (o NotificationRecipientEmailOutput) ToOutput(ctx context.Context) pulumix.Output[*NotificationRecipientEmail] {
	return pulumix.Output[*NotificationRecipientEmail]{
		OutputState: o.OutputState,
	}
}

// User Email subscribed to notification.
func (o NotificationRecipientEmailOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationRecipientEmail) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o NotificationRecipientEmailOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRecipientEmail) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NotificationRecipientEmailOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRecipientEmail) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationRecipientEmailOutput{})
}

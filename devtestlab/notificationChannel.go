// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A notification.
// Azure REST API version: 2018-09-15. Prior API version in Azure Native 1.x: 2018-09-15.
//
// Other available API versions: 2016-05-15.
type NotificationChannel struct {
	pulumi.CustomResourceState

	// The creation date of the notification channel.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Description of notification.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
	EmailRecipient pulumi.StringPtrOutput `pulumi:"emailRecipient"`
	// The list of event for which this notification is enabled.
	Events EventResponseArrayOutput `pulumi:"events"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The locale to use when sending a notification (fallback for unsupported languages is EN).
	NotificationLocale pulumi.StringPtrOutput `pulumi:"notificationLocale"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
	// The webhook URL to send notifications to.
	WebHookUrl pulumi.StringPtrOutput `pulumi:"webHookUrl"`
}

// NewNotificationChannel registers a new resource with the given unique name, arguments, and options.
func NewNotificationChannel(ctx *pulumi.Context,
	name string, args *NotificationChannelArgs, opts ...pulumi.ResourceOption) (*NotificationChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:NotificationChannel"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:NotificationChannel"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NotificationChannel
	err := ctx.RegisterResource("azure-native:devtestlab:NotificationChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationChannel gets an existing NotificationChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationChannelState, opts ...pulumi.ResourceOption) (*NotificationChannel, error) {
	var resource NotificationChannel
	err := ctx.ReadResource("azure-native:devtestlab:NotificationChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationChannel resources.
type notificationChannelState struct {
}

type NotificationChannelState struct {
}

func (NotificationChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelState)(nil)).Elem()
}

type notificationChannelArgs struct {
	// Description of notification.
	Description *string `pulumi:"description"`
	// The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
	EmailRecipient *string `pulumi:"emailRecipient"`
	// The list of event for which this notification is enabled.
	Events []Event `pulumi:"events"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the notification channel.
	Name *string `pulumi:"name"`
	// The locale to use when sending a notification (fallback for unsupported languages is EN).
	NotificationLocale *string `pulumi:"notificationLocale"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The webhook URL to send notifications to.
	WebHookUrl *string `pulumi:"webHookUrl"`
}

// The set of arguments for constructing a NotificationChannel resource.
type NotificationChannelArgs struct {
	// Description of notification.
	Description pulumi.StringPtrInput
	// The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
	EmailRecipient pulumi.StringPtrInput
	// The list of event for which this notification is enabled.
	Events EventArrayInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the notification channel.
	Name pulumi.StringPtrInput
	// The locale to use when sending a notification (fallback for unsupported languages is EN).
	NotificationLocale pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The webhook URL to send notifications to.
	WebHookUrl pulumi.StringPtrInput
}

func (NotificationChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelArgs)(nil)).Elem()
}

type NotificationChannelInput interface {
	pulumi.Input

	ToNotificationChannelOutput() NotificationChannelOutput
	ToNotificationChannelOutputWithContext(ctx context.Context) NotificationChannelOutput
}

func (*NotificationChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannel)(nil)).Elem()
}

func (i *NotificationChannel) ToNotificationChannelOutput() NotificationChannelOutput {
	return i.ToNotificationChannelOutputWithContext(context.Background())
}

func (i *NotificationChannel) ToNotificationChannelOutputWithContext(ctx context.Context) NotificationChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelOutput)
}

func (i *NotificationChannel) ToOutput(ctx context.Context) pulumix.Output[*NotificationChannel] {
	return pulumix.Output[*NotificationChannel]{
		OutputState: i.ToNotificationChannelOutputWithContext(ctx).OutputState,
	}
}

type NotificationChannelOutput struct{ *pulumi.OutputState }

func (NotificationChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannel)(nil)).Elem()
}

func (o NotificationChannelOutput) ToNotificationChannelOutput() NotificationChannelOutput {
	return o
}

func (o NotificationChannelOutput) ToNotificationChannelOutputWithContext(ctx context.Context) NotificationChannelOutput {
	return o
}

func (o NotificationChannelOutput) ToOutput(ctx context.Context) pulumix.Output[*NotificationChannel] {
	return pulumix.Output[*NotificationChannel]{
		OutputState: o.OutputState,
	}
}

// The creation date of the notification channel.
func (o NotificationChannelOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// Description of notification.
func (o NotificationChannelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
func (o NotificationChannelOutput) EmailRecipient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.EmailRecipient }).(pulumi.StringPtrOutput)
}

// The list of event for which this notification is enabled.
func (o NotificationChannelOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v *NotificationChannel) EventResponseArrayOutput { return v.Events }).(EventResponseArrayOutput)
}

// The location of the resource.
func (o NotificationChannelOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o NotificationChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The locale to use when sending a notification (fallback for unsupported languages is EN).
func (o NotificationChannelOutput) NotificationLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.NotificationLocale }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o NotificationChannelOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o NotificationChannelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o NotificationChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o NotificationChannelOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// The webhook URL to send notifications to.
func (o NotificationChannelOutput) WebHookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannel) pulumi.StringPtrOutput { return v.WebHookUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationChannelOutput{})
}

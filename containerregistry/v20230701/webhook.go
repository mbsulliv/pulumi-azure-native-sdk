// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents a webhook for a container registry.
type Webhook struct {
	pulumi.CustomResourceState

	// The list of actions that trigger the webhook to post notifications.
	Actions pulumi.StringArrayOutput `pulumi:"actions"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the webhook at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The status of the webhook at the time the operation was called.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceUri == nil {
		return nil, errors.New("invalid value for required argument 'ServiceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20171001:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210901:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20221201:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230101preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230801preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20231101preview:Webhook"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Webhook
	err := ctx.RegisterResource("azure-native:containerregistry/v20230701:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("azure-native:containerregistry/v20230701:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
}

type WebhookState struct {
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// The list of actions that trigger the webhook to post notifications.
	Actions []string `pulumi:"actions"`
	// Custom headers that will be added to the webhook notifications.
	CustomHeaders map[string]string `pulumi:"customHeaders"`
	// The location of the webhook. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
	Scope *string `pulumi:"scope"`
	// The service URI for the webhook to post notifications.
	ServiceUri string `pulumi:"serviceUri"`
	// The status of the webhook at the time the operation was called.
	Status *string `pulumi:"status"`
	// The tags for the webhook.
	Tags map[string]string `pulumi:"tags"`
	// The name of the webhook.
	WebhookName *string `pulumi:"webhookName"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// The list of actions that trigger the webhook to post notifications.
	Actions pulumi.StringArrayInput
	// Custom headers that will be added to the webhook notifications.
	CustomHeaders pulumi.StringMapInput
	// The location of the webhook. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
	Scope pulumi.StringPtrInput
	// The service URI for the webhook to post notifications.
	ServiceUri pulumi.StringInput
	// The status of the webhook at the time the operation was called.
	Status pulumi.StringPtrInput
	// The tags for the webhook.
	Tags pulumi.StringMapInput
	// The name of the webhook.
	WebhookName pulumi.StringPtrInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

type WebhookOutput struct{ *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

// The list of actions that trigger the webhook to post notifications.
func (o WebhookOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringArrayOutput { return v.Actions }).(pulumi.StringArrayOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o WebhookOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o WebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the webhook at the time the operation was called.
func (o WebhookOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
func (o WebhookOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The status of the webhook at the time the operation was called.
func (o WebhookOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o WebhookOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Webhook) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the resource.
func (o WebhookOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o WebhookOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebhookOutput{})
}

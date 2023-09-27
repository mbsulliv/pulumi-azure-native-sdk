// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A tenant action group resource.
type TenantActionGroup struct {
	pulumi.CustomResourceState

	// The list of AzureAppPush receivers that are part of this tenant action group.
	AzureAppPushReceivers AzureAppPushReceiverResponseArrayOutput `pulumi:"azureAppPushReceivers"`
	// The list of email receivers that are part of this tenant action group.
	EmailReceivers EmailReceiverResponseArrayOutput `pulumi:"emailReceivers"`
	// Indicates whether this tenant action group is enabled. If a tenant action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringOutput `pulumi:"groupShortName"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of SMS receivers that are part of this tenant action group.
	SmsReceivers SmsReceiverResponseArrayOutput `pulumi:"smsReceivers"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of voice receivers that are part of this tenant action group.
	VoiceReceivers VoiceReceiverResponseArrayOutput `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this tenant action group.
	WebhookReceivers WebhookReceiverResponseArrayOutput `pulumi:"webhookReceivers"`
}

// NewTenantActionGroup registers a new resource with the given unique name, arguments, and options.
func NewTenantActionGroup(ctx *pulumi.Context,
	name string, args *TenantActionGroupArgs, opts ...pulumi.ResourceOption) (*TenantActionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupShortName == nil {
		return nil, errors.New("invalid value for required argument 'GroupShortName'")
	}
	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	if args.Enabled == nil {
		args.Enabled = pulumi.Bool(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:TenantActionGroup"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20230301preview:TenantActionGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TenantActionGroup
	err := ctx.RegisterResource("azure-native:insights/v20230501preview:TenantActionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenantActionGroup gets an existing TenantActionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenantActionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantActionGroupState, opts ...pulumi.ResourceOption) (*TenantActionGroup, error) {
	var resource TenantActionGroup
	err := ctx.ReadResource("azure-native:insights/v20230501preview:TenantActionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TenantActionGroup resources.
type tenantActionGroupState struct {
}

type TenantActionGroupState struct {
}

func (TenantActionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantActionGroupState)(nil)).Elem()
}

type tenantActionGroupArgs struct {
	// The list of AzureAppPush receivers that are part of this tenant action group.
	AzureAppPushReceivers []AzureAppPushReceiver `pulumi:"azureAppPushReceivers"`
	// The list of email receivers that are part of this tenant action group.
	EmailReceivers []EmailReceiver `pulumi:"emailReceivers"`
	// Indicates whether this tenant action group is enabled. If a tenant action group is not enabled, then none of its receivers will receive communications.
	Enabled bool `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName string `pulumi:"groupShortName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The management group id.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The list of SMS receivers that are part of this tenant action group.
	SmsReceivers []SmsReceiver `pulumi:"smsReceivers"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The name of the action group.
	TenantActionGroupName *string `pulumi:"tenantActionGroupName"`
	// The list of voice receivers that are part of this tenant action group.
	VoiceReceivers []VoiceReceiver `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this tenant action group.
	WebhookReceivers []WebhookReceiver `pulumi:"webhookReceivers"`
}

// The set of arguments for constructing a TenantActionGroup resource.
type TenantActionGroupArgs struct {
	// The list of AzureAppPush receivers that are part of this tenant action group.
	AzureAppPushReceivers AzureAppPushReceiverArrayInput
	// The list of email receivers that are part of this tenant action group.
	EmailReceivers EmailReceiverArrayInput
	// Indicates whether this tenant action group is enabled. If a tenant action group is not enabled, then none of its receivers will receive communications.
	Enabled pulumi.BoolInput
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// The management group id.
	ManagementGroupId pulumi.StringInput
	// The list of SMS receivers that are part of this tenant action group.
	SmsReceivers SmsReceiverArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The name of the action group.
	TenantActionGroupName pulumi.StringPtrInput
	// The list of voice receivers that are part of this tenant action group.
	VoiceReceivers VoiceReceiverArrayInput
	// The list of webhook receivers that are part of this tenant action group.
	WebhookReceivers WebhookReceiverArrayInput
}

func (TenantActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantActionGroupArgs)(nil)).Elem()
}

type TenantActionGroupInput interface {
	pulumi.Input

	ToTenantActionGroupOutput() TenantActionGroupOutput
	ToTenantActionGroupOutputWithContext(ctx context.Context) TenantActionGroupOutput
}

func (*TenantActionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantActionGroup)(nil)).Elem()
}

func (i *TenantActionGroup) ToTenantActionGroupOutput() TenantActionGroupOutput {
	return i.ToTenantActionGroupOutputWithContext(context.Background())
}

func (i *TenantActionGroup) ToTenantActionGroupOutputWithContext(ctx context.Context) TenantActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantActionGroupOutput)
}

func (i *TenantActionGroup) ToOutput(ctx context.Context) pulumix.Output[*TenantActionGroup] {
	return pulumix.Output[*TenantActionGroup]{
		OutputState: i.ToTenantActionGroupOutputWithContext(ctx).OutputState,
	}
}

type TenantActionGroupOutput struct{ *pulumi.OutputState }

func (TenantActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TenantActionGroup)(nil)).Elem()
}

func (o TenantActionGroupOutput) ToTenantActionGroupOutput() TenantActionGroupOutput {
	return o
}

func (o TenantActionGroupOutput) ToTenantActionGroupOutputWithContext(ctx context.Context) TenantActionGroupOutput {
	return o
}

func (o TenantActionGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*TenantActionGroup] {
	return pulumix.Output[*TenantActionGroup]{
		OutputState: o.OutputState,
	}
}

// The list of AzureAppPush receivers that are part of this tenant action group.
func (o TenantActionGroupOutput) AzureAppPushReceivers() AzureAppPushReceiverResponseArrayOutput {
	return o.ApplyT(func(v *TenantActionGroup) AzureAppPushReceiverResponseArrayOutput { return v.AzureAppPushReceivers }).(AzureAppPushReceiverResponseArrayOutput)
}

// The list of email receivers that are part of this tenant action group.
func (o TenantActionGroupOutput) EmailReceivers() EmailReceiverResponseArrayOutput {
	return o.ApplyT(func(v *TenantActionGroup) EmailReceiverResponseArrayOutput { return v.EmailReceivers }).(EmailReceiverResponseArrayOutput)
}

// Indicates whether this tenant action group is enabled. If a tenant action group is not enabled, then none of its receivers will receive communications.
func (o TenantActionGroupOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *TenantActionGroup) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The short name of the action group. This will be used in SMS messages.
func (o TenantActionGroupOutput) GroupShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantActionGroup) pulumi.StringOutput { return v.GroupShortName }).(pulumi.StringOutput)
}

// Resource location
func (o TenantActionGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantActionGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o TenantActionGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantActionGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of SMS receivers that are part of this tenant action group.
func (o TenantActionGroupOutput) SmsReceivers() SmsReceiverResponseArrayOutput {
	return o.ApplyT(func(v *TenantActionGroup) SmsReceiverResponseArrayOutput { return v.SmsReceivers }).(SmsReceiverResponseArrayOutput)
}

// Resource tags
func (o TenantActionGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TenantActionGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o TenantActionGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TenantActionGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of voice receivers that are part of this tenant action group.
func (o TenantActionGroupOutput) VoiceReceivers() VoiceReceiverResponseArrayOutput {
	return o.ApplyT(func(v *TenantActionGroup) VoiceReceiverResponseArrayOutput { return v.VoiceReceivers }).(VoiceReceiverResponseArrayOutput)
}

// The list of webhook receivers that are part of this tenant action group.
func (o TenantActionGroupOutput) WebhookReceivers() WebhookReceiverResponseArrayOutput {
	return o.ApplyT(func(v *TenantActionGroup) WebhookReceiverResponseArrayOutput { return v.WebhookReceivers }).(WebhookReceiverResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(TenantActionGroupOutput{})
}

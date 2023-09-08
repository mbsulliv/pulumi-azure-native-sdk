// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a tenant action group.
func LookupTenantActionGroup(ctx *pulumi.Context, args *LookupTenantActionGroupArgs, opts ...pulumi.InvokeOption) (*LookupTenantActionGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTenantActionGroupResult
	err := ctx.Invoke("azure-native:insights/v20230501preview:getTenantActionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTenantActionGroupArgs struct {
	// The management group id.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The name of the action group.
	TenantActionGroupName string `pulumi:"tenantActionGroupName"`
}

// A tenant action group resource.
type LookupTenantActionGroupResult struct {
	// The list of AzureAppPush receivers that are part of this tenant action group.
	AzureAppPushReceivers []AzureAppPushReceiverResponse `pulumi:"azureAppPushReceivers"`
	// The list of email receivers that are part of this tenant action group.
	EmailReceivers []EmailReceiverResponse `pulumi:"emailReceivers"`
	// Indicates whether this tenant action group is enabled. If a tenant action group is not enabled, then none of its receivers will receive communications.
	Enabled bool `pulumi:"enabled"`
	// The short name of the action group. This will be used in SMS messages.
	GroupShortName string `pulumi:"groupShortName"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// The list of SMS receivers that are part of this tenant action group.
	SmsReceivers []SmsReceiverResponse `pulumi:"smsReceivers"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type string `pulumi:"type"`
	// The list of voice receivers that are part of this tenant action group.
	VoiceReceivers []VoiceReceiverResponse `pulumi:"voiceReceivers"`
	// The list of webhook receivers that are part of this tenant action group.
	WebhookReceivers []WebhookReceiverResponse `pulumi:"webhookReceivers"`
}

// Defaults sets the appropriate defaults for LookupTenantActionGroupResult
func (val *LookupTenantActionGroupResult) Defaults() *LookupTenantActionGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if utilities.IsZero(tmp.Enabled) {
		tmp.Enabled = true
	}
	return &tmp
}

func LookupTenantActionGroupOutput(ctx *pulumi.Context, args LookupTenantActionGroupOutputArgs, opts ...pulumi.InvokeOption) LookupTenantActionGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTenantActionGroupResult, error) {
			args := v.(LookupTenantActionGroupArgs)
			r, err := LookupTenantActionGroup(ctx, &args, opts...)
			var s LookupTenantActionGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTenantActionGroupResultOutput)
}

type LookupTenantActionGroupOutputArgs struct {
	// The management group id.
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	// The name of the action group.
	TenantActionGroupName pulumi.StringInput `pulumi:"tenantActionGroupName"`
}

func (LookupTenantActionGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantActionGroupArgs)(nil)).Elem()
}

// A tenant action group resource.
type LookupTenantActionGroupResultOutput struct{ *pulumi.OutputState }

func (LookupTenantActionGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantActionGroupResult)(nil)).Elem()
}

func (o LookupTenantActionGroupResultOutput) ToLookupTenantActionGroupResultOutput() LookupTenantActionGroupResultOutput {
	return o
}

func (o LookupTenantActionGroupResultOutput) ToLookupTenantActionGroupResultOutputWithContext(ctx context.Context) LookupTenantActionGroupResultOutput {
	return o
}

// The list of AzureAppPush receivers that are part of this tenant action group.
func (o LookupTenantActionGroupResultOutput) AzureAppPushReceivers() AzureAppPushReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) []AzureAppPushReceiverResponse { return v.AzureAppPushReceivers }).(AzureAppPushReceiverResponseArrayOutput)
}

// The list of email receivers that are part of this tenant action group.
func (o LookupTenantActionGroupResultOutput) EmailReceivers() EmailReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) []EmailReceiverResponse { return v.EmailReceivers }).(EmailReceiverResponseArrayOutput)
}

// Indicates whether this tenant action group is enabled. If a tenant action group is not enabled, then none of its receivers will receive communications.
func (o LookupTenantActionGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The short name of the action group. This will be used in SMS messages.
func (o LookupTenantActionGroupResultOutput) GroupShortName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) string { return v.GroupShortName }).(pulumi.StringOutput)
}

// Azure resource Id
func (o LookupTenantActionGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupTenantActionGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o LookupTenantActionGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of SMS receivers that are part of this tenant action group.
func (o LookupTenantActionGroupResultOutput) SmsReceivers() SmsReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) []SmsReceiverResponse { return v.SmsReceivers }).(SmsReceiverResponseArrayOutput)
}

// Resource tags
func (o LookupTenantActionGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o LookupTenantActionGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The list of voice receivers that are part of this tenant action group.
func (o LookupTenantActionGroupResultOutput) VoiceReceivers() VoiceReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) []VoiceReceiverResponse { return v.VoiceReceivers }).(VoiceReceiverResponseArrayOutput)
}

// The list of webhook receivers that are part of this tenant action group.
func (o LookupTenantActionGroupResultOutput) WebhookReceivers() WebhookReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupTenantActionGroupResult) []WebhookReceiverResponse { return v.WebhookReceivers }).(WebhookReceiverResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTenantActionGroupResultOutput{})
}
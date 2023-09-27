// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the active diagnostic setting for AadIam.
func LookupDiagnosticSetting(ctx *pulumi.Context, args *LookupDiagnosticSettingArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticSettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDiagnosticSettingResult
	err := ctx.Invoke("azure-native:aadiam/v20170401:getDiagnosticSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiagnosticSettingArgs struct {
	// The name of the diagnostic setting.
	Name string `pulumi:"name"`
}

// The diagnostic setting resource.
type LookupDiagnosticSettingResult struct {
	// The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId *string `pulumi:"eventHubAuthorizationRuleId"`
	// The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName *string `pulumi:"eventHubName"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// The list of logs settings.
	Logs []LogSettingsResponse `pulumi:"logs"`
	// Azure resource name
	Name string `pulumi:"name"`
	// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId *string `pulumi:"serviceBusRuleId"`
	// The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Azure resource type
	Type string `pulumi:"type"`
	// The workspace ID (resource ID of a Log Analytics workspace) for a Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId *string `pulumi:"workspaceId"`
}

func LookupDiagnosticSettingOutput(ctx *pulumi.Context, args LookupDiagnosticSettingOutputArgs, opts ...pulumi.InvokeOption) LookupDiagnosticSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiagnosticSettingResult, error) {
			args := v.(LookupDiagnosticSettingArgs)
			r, err := LookupDiagnosticSetting(ctx, &args, opts...)
			var s LookupDiagnosticSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiagnosticSettingResultOutput)
}

type LookupDiagnosticSettingOutputArgs struct {
	// The name of the diagnostic setting.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDiagnosticSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticSettingArgs)(nil)).Elem()
}

// The diagnostic setting resource.
type LookupDiagnosticSettingResultOutput struct{ *pulumi.OutputState }

func (LookupDiagnosticSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticSettingResult)(nil)).Elem()
}

func (o LookupDiagnosticSettingResultOutput) ToLookupDiagnosticSettingResultOutput() LookupDiagnosticSettingResultOutput {
	return o
}

func (o LookupDiagnosticSettingResultOutput) ToLookupDiagnosticSettingResultOutputWithContext(ctx context.Context) LookupDiagnosticSettingResultOutput {
	return o
}

func (o LookupDiagnosticSettingResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDiagnosticSettingResult] {
	return pulumix.Output[LookupDiagnosticSettingResult]{
		OutputState: o.OutputState,
	}
}

// The resource Id for the event hub authorization rule.
func (o LookupDiagnosticSettingResultOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

// The name of the event hub. If none is specified, the default event hub will be selected.
func (o LookupDiagnosticSettingResultOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

// Azure resource Id
func (o LookupDiagnosticSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of logs settings.
func (o LookupDiagnosticSettingResultOutput) Logs() LogSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) []LogSettingsResponse { return v.Logs }).(LogSettingsResponseArrayOutput)
}

// Azure resource name
func (o LookupDiagnosticSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
func (o LookupDiagnosticSettingResultOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

// The resource ID of the storage account to which you would like to send Diagnostic Logs.
func (o LookupDiagnosticSettingResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

// Azure resource type
func (o LookupDiagnosticSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

// The workspace ID (resource ID of a Log Analytics workspace) for a Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
func (o LookupDiagnosticSettingResultOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiagnosticSettingResultOutput{})
}

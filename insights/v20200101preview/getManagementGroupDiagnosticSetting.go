// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the active management group diagnostic settings for the specified resource.
func LookupManagementGroupDiagnosticSetting(ctx *pulumi.Context, args *LookupManagementGroupDiagnosticSettingArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupDiagnosticSettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagementGroupDiagnosticSettingResult
	err := ctx.Invoke("azure-native:insights/v20200101preview:getManagementGroupDiagnosticSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupDiagnosticSettingArgs struct {
	// The management group id.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The name of the diagnostic setting.
	Name string `pulumi:"name"`
}

// The management group diagnostic setting resource.
type LookupManagementGroupDiagnosticSettingResult struct {
	// The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId *string `pulumi:"eventHubAuthorizationRuleId"`
	// The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName *string `pulumi:"eventHubName"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// Location of the resource
	Location *string `pulumi:"location"`
	// The list of logs settings.
	Logs []ManagementGroupLogSettingsResponse `pulumi:"logs"`
	// Azure resource name
	Name string `pulumi:"name"`
	// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId *string `pulumi:"serviceBusRuleId"`
	// The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Azure resource type
	Type string `pulumi:"type"`
	// The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId *string `pulumi:"workspaceId"`
}

func LookupManagementGroupDiagnosticSettingOutput(ctx *pulumi.Context, args LookupManagementGroupDiagnosticSettingOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupDiagnosticSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupDiagnosticSettingResult, error) {
			args := v.(LookupManagementGroupDiagnosticSettingArgs)
			r, err := LookupManagementGroupDiagnosticSetting(ctx, &args, opts...)
			var s LookupManagementGroupDiagnosticSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementGroupDiagnosticSettingResultOutput)
}

type LookupManagementGroupDiagnosticSettingOutputArgs struct {
	// The management group id.
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	// The name of the diagnostic setting.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupManagementGroupDiagnosticSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupDiagnosticSettingArgs)(nil)).Elem()
}

// The management group diagnostic setting resource.
type LookupManagementGroupDiagnosticSettingResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupDiagnosticSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupDiagnosticSettingResult)(nil)).Elem()
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) ToLookupManagementGroupDiagnosticSettingResultOutput() LookupManagementGroupDiagnosticSettingResultOutput {
	return o
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) ToLookupManagementGroupDiagnosticSettingResultOutputWithContext(ctx context.Context) LookupManagementGroupDiagnosticSettingResultOutput {
	return o
}

// The resource Id for the event hub authorization rule.
func (o LookupManagementGroupDiagnosticSettingResultOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

// The name of the event hub. If none is specified, the default event hub will be selected.
func (o LookupManagementGroupDiagnosticSettingResultOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

// Azure resource Id
func (o LookupManagementGroupDiagnosticSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the resource
func (o LookupManagementGroupDiagnosticSettingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The list of logs settings.
func (o LookupManagementGroupDiagnosticSettingResultOutput) Logs() ManagementGroupLogSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) []ManagementGroupLogSettingsResponse {
		return v.Logs
	}).(ManagementGroupLogSettingsResponseArrayOutput)
}

// Azure resource name
func (o LookupManagementGroupDiagnosticSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
func (o LookupManagementGroupDiagnosticSettingResultOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

// The resource ID of the storage account to which you would like to send Diagnostic Logs.
func (o LookupManagementGroupDiagnosticSettingResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

// Azure resource type
func (o LookupManagementGroupDiagnosticSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

// The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
func (o LookupManagementGroupDiagnosticSettingResultOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupDiagnosticSettingResultOutput{})
}

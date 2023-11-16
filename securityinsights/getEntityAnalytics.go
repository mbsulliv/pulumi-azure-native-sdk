// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a setting.
// Azure REST API version: 2023-06-01-preview.
func LookupEntityAnalytics(ctx *pulumi.Context, args *LookupEntityAnalyticsArgs, opts ...pulumi.InvokeOption) (*LookupEntityAnalyticsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEntityAnalyticsResult
	err := ctx.Invoke("azure-native:securityinsights:getEntityAnalytics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntityAnalyticsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName string `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Settings with single toggle.
type LookupEntityAnalyticsResult struct {
	// The relevant entity providers that are synced
	EntityProviders []string `pulumi:"entityProviders"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the setting
	// Expected value is 'EntityAnalytics'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEntityAnalyticsOutput(ctx *pulumi.Context, args LookupEntityAnalyticsOutputArgs, opts ...pulumi.InvokeOption) LookupEntityAnalyticsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEntityAnalyticsResult, error) {
			args := v.(LookupEntityAnalyticsArgs)
			r, err := LookupEntityAnalytics(ctx, &args, opts...)
			var s LookupEntityAnalyticsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEntityAnalyticsResultOutput)
}

type LookupEntityAnalyticsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName pulumi.StringInput `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEntityAnalyticsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityAnalyticsArgs)(nil)).Elem()
}

// Settings with single toggle.
type LookupEntityAnalyticsResultOutput struct{ *pulumi.OutputState }

func (LookupEntityAnalyticsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityAnalyticsResult)(nil)).Elem()
}

func (o LookupEntityAnalyticsResultOutput) ToLookupEntityAnalyticsResultOutput() LookupEntityAnalyticsResultOutput {
	return o
}

func (o LookupEntityAnalyticsResultOutput) ToLookupEntityAnalyticsResultOutputWithContext(ctx context.Context) LookupEntityAnalyticsResultOutput {
	return o
}

// The relevant entity providers that are synced
func (o LookupEntityAnalyticsResultOutput) EntityProviders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) []string { return v.EntityProviders }).(pulumi.StringArrayOutput)
}

// Etag of the azure resource
func (o LookupEntityAnalyticsResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEntityAnalyticsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the setting
// Expected value is 'EntityAnalytics'.
func (o LookupEntityAnalyticsResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEntityAnalyticsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEntityAnalyticsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEntityAnalyticsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEntityAnalyticsResultOutput{})
}

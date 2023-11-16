// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a setting.
func LookupAnomalies(ctx *pulumi.Context, args *LookupAnomaliesArgs, opts ...pulumi.InvokeOption) (*LookupAnomaliesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAnomaliesResult
	err := ctx.Invoke("azure-native:securityinsights/v20230801preview:getAnomalies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnomaliesArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName string `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Settings with single toggle.
type LookupAnomaliesResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Determines whether the setting is enable or disabled.
	IsEnabled bool `pulumi:"isEnabled"`
	// The kind of the setting
	// Expected value is 'Anomalies'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAnomaliesOutput(ctx *pulumi.Context, args LookupAnomaliesOutputArgs, opts ...pulumi.InvokeOption) LookupAnomaliesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnomaliesResult, error) {
			args := v.(LookupAnomaliesArgs)
			r, err := LookupAnomalies(ctx, &args, opts...)
			var s LookupAnomaliesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnomaliesResultOutput)
}

type LookupAnomaliesOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName pulumi.StringInput `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAnomaliesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomaliesArgs)(nil)).Elem()
}

// Settings with single toggle.
type LookupAnomaliesResultOutput struct{ *pulumi.OutputState }

func (LookupAnomaliesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomaliesResult)(nil)).Elem()
}

func (o LookupAnomaliesResultOutput) ToLookupAnomaliesResultOutput() LookupAnomaliesResultOutput {
	return o
}

func (o LookupAnomaliesResultOutput) ToLookupAnomaliesResultOutputWithContext(ctx context.Context) LookupAnomaliesResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupAnomaliesResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAnomaliesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Determines whether the setting is enable or disabled.
func (o LookupAnomaliesResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

// The kind of the setting
// Expected value is 'Anomalies'.
func (o LookupAnomaliesResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAnomaliesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAnomaliesResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAnomaliesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnomaliesResultOutput{})
}

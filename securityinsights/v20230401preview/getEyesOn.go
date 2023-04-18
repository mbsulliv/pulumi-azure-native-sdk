// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a setting.
func LookupEyesOn(ctx *pulumi.Context, args *LookupEyesOnArgs, opts ...pulumi.InvokeOption) (*LookupEyesOnResult, error) {
	var rv LookupEyesOnResult
	err := ctx.Invoke("azure-native:securityinsights/v20230401preview:getEyesOn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEyesOnArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName string `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Settings with single toggle.
type LookupEyesOnResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Determines whether the setting is enable or disabled.
	IsEnabled bool `pulumi:"isEnabled"`
	// The kind of the setting
	// Expected value is 'EyesOn'.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEyesOnOutput(ctx *pulumi.Context, args LookupEyesOnOutputArgs, opts ...pulumi.InvokeOption) LookupEyesOnResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEyesOnResult, error) {
			args := v.(LookupEyesOnArgs)
			r, err := LookupEyesOn(ctx, &args, opts...)
			var s LookupEyesOnResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEyesOnResultOutput)
}

type LookupEyesOnOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
	SettingsName pulumi.StringInput `pulumi:"settingsName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEyesOnOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEyesOnArgs)(nil)).Elem()
}

// Settings with single toggle.
type LookupEyesOnResultOutput struct{ *pulumi.OutputState }

func (LookupEyesOnResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEyesOnResult)(nil)).Elem()
}

func (o LookupEyesOnResultOutput) ToLookupEyesOnResultOutput() LookupEyesOnResultOutput {
	return o
}

func (o LookupEyesOnResultOutput) ToLookupEyesOnResultOutputWithContext(ctx context.Context) LookupEyesOnResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupEyesOnResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEyesOnResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEyesOnResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEyesOnResult) string { return v.Id }).(pulumi.StringOutput)
}

// Determines whether the setting is enable or disabled.
func (o LookupEyesOnResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEyesOnResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

// The kind of the setting
// Expected value is 'EyesOn'.
func (o LookupEyesOnResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEyesOnResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEyesOnResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEyesOnResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEyesOnResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEyesOnResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEyesOnResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEyesOnResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEyesOnResultOutput{})
}

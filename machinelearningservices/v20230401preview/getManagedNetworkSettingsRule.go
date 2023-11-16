// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an outbound rule from the managed network of a machine learning workspace.
func LookupManagedNetworkSettingsRule(ctx *pulumi.Context, args *LookupManagedNetworkSettingsRuleArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkSettingsRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedNetworkSettingsRuleResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getManagedNetworkSettingsRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkSettingsRuleArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the workspace managed network outbound rule
	RuleName string `pulumi:"ruleName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Outbound Rule Basic Resource for the managed network of a machine learning workspace.
type LookupManagedNetworkSettingsRuleResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Outbound Rule for the managed network of a machine learning workspace.
	Properties interface{} `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupManagedNetworkSettingsRuleOutput(ctx *pulumi.Context, args LookupManagedNetworkSettingsRuleOutputArgs, opts ...pulumi.InvokeOption) LookupManagedNetworkSettingsRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedNetworkSettingsRuleResult, error) {
			args := v.(LookupManagedNetworkSettingsRuleArgs)
			r, err := LookupManagedNetworkSettingsRule(ctx, &args, opts...)
			var s LookupManagedNetworkSettingsRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedNetworkSettingsRuleResultOutput)
}

type LookupManagedNetworkSettingsRuleOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the workspace managed network outbound rule
	RuleName pulumi.StringInput `pulumi:"ruleName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupManagedNetworkSettingsRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkSettingsRuleArgs)(nil)).Elem()
}

// Outbound Rule Basic Resource for the managed network of a machine learning workspace.
type LookupManagedNetworkSettingsRuleResultOutput struct{ *pulumi.OutputState }

func (LookupManagedNetworkSettingsRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkSettingsRuleResult)(nil)).Elem()
}

func (o LookupManagedNetworkSettingsRuleResultOutput) ToLookupManagedNetworkSettingsRuleResultOutput() LookupManagedNetworkSettingsRuleResultOutput {
	return o
}

func (o LookupManagedNetworkSettingsRuleResultOutput) ToLookupManagedNetworkSettingsRuleResultOutputWithContext(ctx context.Context) LookupManagedNetworkSettingsRuleResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedNetworkSettingsRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkSettingsRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagedNetworkSettingsRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkSettingsRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Outbound Rule for the managed network of a machine learning workspace.
func (o LookupManagedNetworkSettingsRuleResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupManagedNetworkSettingsRuleResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupManagedNetworkSettingsRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedNetworkSettingsRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagedNetworkSettingsRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkSettingsRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedNetworkSettingsRuleResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response for POST requests that return single SharedAccessAuthorizationRule.
func LookupNamespaceAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamespaceAuthorizationRuleResult
	err := ctx.Invoke("azure-native:notificationhubs/v20230101preview:getNamespaceAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceAuthorizationRuleArgs struct {
	// Authorization Rule Name
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for POST requests that return single SharedAccessAuthorizationRule.
type LookupNamespaceAuthorizationRuleResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Deprecated - only for compatibility.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// SharedAccessAuthorizationRule properties.
	Properties SharedAccessAuthorizationRulePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Deprecated - only for compatibility.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNamespaceAuthorizationRuleOutput(ctx *pulumi.Context, args LookupNamespaceAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceAuthorizationRuleResult, error) {
			args := v.(LookupNamespaceAuthorizationRuleArgs)
			r, err := LookupNamespaceAuthorizationRule(ctx, &args, opts...)
			var s LookupNamespaceAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceAuthorizationRuleResultOutput)
}

type LookupNamespaceAuthorizationRuleOutputArgs struct {
	// Authorization Rule Name
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleArgs)(nil)).Elem()
}

// Response for POST requests that return single SharedAccessAuthorizationRule.
type LookupNamespaceAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToLookupNamespaceAuthorizationRuleResultOutput() LookupNamespaceAuthorizationRuleResultOutput {
	return o
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToLookupNamespaceAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupNamespaceAuthorizationRuleResultOutput {
	return o
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNamespaceAuthorizationRuleResult] {
	return pulumix.Output[LookupNamespaceAuthorizationRuleResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNamespaceAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Deprecated - only for compatibility.
func (o LookupNamespaceAuthorizationRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupNamespaceAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// SharedAccessAuthorizationRule properties.
func (o LookupNamespaceAuthorizationRuleResultOutput) Properties() SharedAccessAuthorizationRulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) SharedAccessAuthorizationRulePropertiesResponse {
		return v.Properties
	}).(SharedAccessAuthorizationRulePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNamespaceAuthorizationRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Deprecated - only for compatibility.
func (o LookupNamespaceAuthorizationRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNamespaceAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceAuthorizationRuleResultOutput{})
}

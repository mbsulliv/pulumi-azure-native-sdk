// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an AuthorizationRule for a Namespace by rule name.
func LookupNamespaceAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamespaceAuthorizationRuleResult
	err := ctx.Invoke("azure-native:eventhub/v20230101preview:getNamespaceAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in a List or Get AuthorizationRule operation
type LookupNamespaceAuthorizationRuleResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
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
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleArgs)(nil)).Elem()
}

// Single item in a List or Get AuthorizationRule operation
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

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupNamespaceAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNamespaceAuthorizationRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNamespaceAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The rights associated with the rule.
func (o LookupNamespaceAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

// The system meta data relating to this resource.
func (o LookupNamespaceAuthorizationRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o LookupNamespaceAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceAuthorizationRuleResultOutput{})
}

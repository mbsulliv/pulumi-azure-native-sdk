// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Hybrid connection authorization rule for a hybrid connection by name.
func LookupHybridConnectionAuthorizationRule(ctx *pulumi.Context, args *LookupHybridConnectionAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupHybridConnectionAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHybridConnectionAuthorizationRuleResult
	err := ctx.Invoke("azure-native:relay/v20211101:getHybridConnectionAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridConnectionAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The hybrid connection name.
	HybridConnectionName string `pulumi:"hybridConnectionName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in a List or Get AuthorizationRule operation
type LookupHybridConnectionAuthorizationRuleResult struct {
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

func LookupHybridConnectionAuthorizationRuleOutput(ctx *pulumi.Context, args LookupHybridConnectionAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupHybridConnectionAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridConnectionAuthorizationRuleResult, error) {
			args := v.(LookupHybridConnectionAuthorizationRuleArgs)
			r, err := LookupHybridConnectionAuthorizationRule(ctx, &args, opts...)
			var s LookupHybridConnectionAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridConnectionAuthorizationRuleResultOutput)
}

type LookupHybridConnectionAuthorizationRuleOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The hybrid connection name.
	HybridConnectionName pulumi.StringInput `pulumi:"hybridConnectionName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHybridConnectionAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridConnectionAuthorizationRuleArgs)(nil)).Elem()
}

// Single item in a List or Get AuthorizationRule operation
type LookupHybridConnectionAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupHybridConnectionAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridConnectionAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) ToLookupHybridConnectionAuthorizationRuleResultOutput() LookupHybridConnectionAuthorizationRuleResultOutput {
	return o
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) ToLookupHybridConnectionAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupHybridConnectionAuthorizationRuleResultOutput {
	return o
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupHybridConnectionAuthorizationRuleResult] {
	return pulumix.Output[LookupHybridConnectionAuthorizationRuleResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupHybridConnectionAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupHybridConnectionAuthorizationRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupHybridConnectionAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The rights associated with the rule.
func (o LookupHybridConnectionAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

// The system meta data relating to this resource.
func (o LookupHybridConnectionAuthorizationRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o LookupHybridConnectionAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridConnectionAuthorizationRuleResultOutput{})
}

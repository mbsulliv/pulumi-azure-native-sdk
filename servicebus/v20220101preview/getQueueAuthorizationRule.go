// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an authorization rule for a queue by rule name.
func LookupQueueAuthorizationRule(ctx *pulumi.Context, args *LookupQueueAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupQueueAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupQueueAuthorizationRuleResult
	err := ctx.Invoke("azure-native:servicebus/v20220101preview:getQueueAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The queue name.
	QueueName string `pulumi:"queueName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a namespace authorization rule.
type LookupQueueAuthorizationRuleResult struct {
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

func LookupQueueAuthorizationRuleOutput(ctx *pulumi.Context, args LookupQueueAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupQueueAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueAuthorizationRuleResult, error) {
			args := v.(LookupQueueAuthorizationRuleArgs)
			r, err := LookupQueueAuthorizationRule(ctx, &args, opts...)
			var s LookupQueueAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueueAuthorizationRuleResultOutput)
}

type LookupQueueAuthorizationRuleOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The queue name.
	QueueName pulumi.StringInput `pulumi:"queueName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueueAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueAuthorizationRuleArgs)(nil)).Elem()
}

// Description of a namespace authorization rule.
type LookupQueueAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupQueueAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupQueueAuthorizationRuleResultOutput) ToLookupQueueAuthorizationRuleResultOutput() LookupQueueAuthorizationRuleResultOutput {
	return o
}

func (o LookupQueueAuthorizationRuleResultOutput) ToLookupQueueAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupQueueAuthorizationRuleResultOutput {
	return o
}

func (o LookupQueueAuthorizationRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupQueueAuthorizationRuleResult] {
	return pulumix.Output[LookupQueueAuthorizationRuleResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupQueueAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupQueueAuthorizationRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupQueueAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The rights associated with the rule.
func (o LookupQueueAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

// The system meta data relating to this resource.
func (o LookupQueueAuthorizationRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o LookupQueueAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueAuthorizationRuleResultOutput{})
}

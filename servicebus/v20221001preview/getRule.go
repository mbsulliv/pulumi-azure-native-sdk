// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the description for the specified rule.
func LookupRule(ctx *pulumi.Context, args *LookupRuleArgs, opts ...pulumi.InvokeOption) (*LookupRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRuleResult
	err := ctx.Invoke("azure-native:servicebus/v20221001preview:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRuleArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rule name.
	RuleName string `pulumi:"ruleName"`
	// The subscription name.
	SubscriptionName string `pulumi:"subscriptionName"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// Description of Rule Resource.
type LookupRuleResult struct {
	// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
	Action *ActionResponse `pulumi:"action"`
	// Properties of correlationFilter
	CorrelationFilter *CorrelationFilterResponse `pulumi:"correlationFilter"`
	// Filter type that is evaluated against a BrokeredMessage.
	FilterType *string `pulumi:"filterType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of sqlFilter
	SqlFilter *SqlFilterResponse `pulumi:"sqlFilter"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRuleResult
func (val *LookupRuleResult) Defaults() *LookupRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Action = tmp.Action.Defaults()

	tmp.CorrelationFilter = tmp.CorrelationFilter.Defaults()

	tmp.SqlFilter = tmp.SqlFilter.Defaults()

	return &tmp
}

func LookupRuleOutput(ctx *pulumi.Context, args LookupRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleResult, error) {
			args := v.(LookupRuleArgs)
			r, err := LookupRule(ctx, &args, opts...)
			var s LookupRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleResultOutput)
}

type LookupRuleOutputArgs struct {
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The rule name.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
	// The subscription name.
	SubscriptionName pulumi.StringInput `pulumi:"subscriptionName"`
	// The topic name.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (LookupRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleArgs)(nil)).Elem()
}

// Description of Rule Resource.
type LookupRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleResult)(nil)).Elem()
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutput() LookupRuleResultOutput {
	return o
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutputWithContext(ctx context.Context) LookupRuleResultOutput {
	return o
}

// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
func (o LookupRuleResultOutput) Action() ActionResponsePtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *ActionResponse { return v.Action }).(ActionResponsePtrOutput)
}

// Properties of correlationFilter
func (o LookupRuleResultOutput) CorrelationFilter() CorrelationFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *CorrelationFilterResponse { return v.CorrelationFilter }).(CorrelationFilterResponsePtrOutput)
}

// Filter type that is evaluated against a BrokeredMessage.
func (o LookupRuleResultOutput) FilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *string { return v.FilterType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of sqlFilter
func (o LookupRuleResultOutput) SqlFilter() SqlFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *SqlFilterResponse { return v.SqlFilter }).(SqlFilterResponsePtrOutput)
}

// The system meta data relating to this resource.
func (o LookupRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o LookupRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleResultOutput{})
}

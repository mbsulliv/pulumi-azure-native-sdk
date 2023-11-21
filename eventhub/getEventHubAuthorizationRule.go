// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an AuthorizationRule for an Event Hub by rule name.
// Azure REST API version: 2022-10-01-preview.
//
// Other available API versions: 2015-08-01, 2023-01-01-preview, 2024-01-01.
func LookupEventHubAuthorizationRule(ctx *pulumi.Context, args *LookupEventHubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupEventHubAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEventHubAuthorizationRuleResult
	err := ctx.Invoke("azure-native:eventhub:getEventHubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single item in a List or Get AuthorizationRule operation
type LookupEventHubAuthorizationRuleResult struct {
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

func LookupEventHubAuthorizationRuleOutput(ctx *pulumi.Context, args LookupEventHubAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubAuthorizationRuleResult, error) {
			args := v.(LookupEventHubAuthorizationRuleArgs)
			r, err := LookupEventHubAuthorizationRule(ctx, &args, opts...)
			var s LookupEventHubAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubAuthorizationRuleResultOutput)
}

type LookupEventHubAuthorizationRuleOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The Event Hub name
	EventHubName pulumi.StringInput `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubAuthorizationRuleArgs)(nil)).Elem()
}

// Single item in a List or Get AuthorizationRule operation
type LookupEventHubAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupEventHubAuthorizationRuleResultOutput) ToLookupEventHubAuthorizationRuleResultOutput() LookupEventHubAuthorizationRuleResultOutput {
	return o
}

func (o LookupEventHubAuthorizationRuleResultOutput) ToLookupEventHubAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupEventHubAuthorizationRuleResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEventHubAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupEventHubAuthorizationRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEventHubAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The rights associated with the rule.
func (o LookupEventHubAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

// The system meta data relating to this resource.
func (o LookupEventHubAuthorizationRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o LookupEventHubAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubAuthorizationRuleResultOutput{})
}

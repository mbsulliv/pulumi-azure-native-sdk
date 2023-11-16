// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get authorizationRule for a WCF relay by name.
func LookupWCFRelayAuthorizationRule(ctx *pulumi.Context, args *LookupWCFRelayAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupWCFRelayAuthorizationRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWCFRelayAuthorizationRuleResult
	err := ctx.Invoke("azure-native:relay/v20170401:getWCFRelayAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWCFRelayAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The relay name.
	RelayName string `pulumi:"relayName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a namespace authorization rule.
type LookupWCFRelayAuthorizationRuleResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWCFRelayAuthorizationRuleOutput(ctx *pulumi.Context, args LookupWCFRelayAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupWCFRelayAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWCFRelayAuthorizationRuleResult, error) {
			args := v.(LookupWCFRelayAuthorizationRuleArgs)
			r, err := LookupWCFRelayAuthorizationRule(ctx, &args, opts...)
			var s LookupWCFRelayAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWCFRelayAuthorizationRuleResultOutput)
}

type LookupWCFRelayAuthorizationRuleOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The relay name.
	RelayName pulumi.StringInput `pulumi:"relayName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWCFRelayAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWCFRelayAuthorizationRuleArgs)(nil)).Elem()
}

// Description of a namespace authorization rule.
type LookupWCFRelayAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupWCFRelayAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWCFRelayAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupWCFRelayAuthorizationRuleResultOutput) ToLookupWCFRelayAuthorizationRuleResultOutput() LookupWCFRelayAuthorizationRuleResultOutput {
	return o
}

func (o LookupWCFRelayAuthorizationRuleResultOutput) ToLookupWCFRelayAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupWCFRelayAuthorizationRuleResultOutput {
	return o
}

// Resource ID.
func (o LookupWCFRelayAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupWCFRelayAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The rights associated with the rule.
func (o LookupWCFRelayAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWCFRelayAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

// Resource type.
func (o LookupWCFRelayAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWCFRelayAuthorizationRuleResultOutput{})
}

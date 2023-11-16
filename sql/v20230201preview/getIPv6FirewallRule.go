// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an IPv6 firewall rule.
func LookupIPv6FirewallRule(ctx *pulumi.Context, args *LookupIPv6FirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupIPv6FirewallRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIPv6FirewallRuleResult
	err := ctx.Invoke("azure-native:sql/v20230201preview:getIPv6FirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPv6FirewallRuleArgs struct {
	// The name of the firewall rule.
	FirewallRuleName string `pulumi:"firewallRuleName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// An IPv6 server firewall rule.
type LookupIPv6FirewallRuleResult struct {
	// The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to startIpAddress.
	EndIPv6Address *string `pulumi:"endIPv6Address"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The start IP address of the firewall rule. Must be IPv6 format.
	StartIPv6Address *string `pulumi:"startIPv6Address"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupIPv6FirewallRuleOutput(ctx *pulumi.Context, args LookupIPv6FirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupIPv6FirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPv6FirewallRuleResult, error) {
			args := v.(LookupIPv6FirewallRuleArgs)
			r, err := LookupIPv6FirewallRule(ctx, &args, opts...)
			var s LookupIPv6FirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIPv6FirewallRuleResultOutput)
}

type LookupIPv6FirewallRuleOutputArgs struct {
	// The name of the firewall rule.
	FirewallRuleName pulumi.StringInput `pulumi:"firewallRuleName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupIPv6FirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPv6FirewallRuleArgs)(nil)).Elem()
}

// An IPv6 server firewall rule.
type LookupIPv6FirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupIPv6FirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPv6FirewallRuleResult)(nil)).Elem()
}

func (o LookupIPv6FirewallRuleResultOutput) ToLookupIPv6FirewallRuleResultOutput() LookupIPv6FirewallRuleResultOutput {
	return o
}

func (o LookupIPv6FirewallRuleResultOutput) ToLookupIPv6FirewallRuleResultOutputWithContext(ctx context.Context) LookupIPv6FirewallRuleResultOutput {
	return o
}

// The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to startIpAddress.
func (o LookupIPv6FirewallRuleResultOutput) EndIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) *string { return v.EndIPv6Address }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupIPv6FirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupIPv6FirewallRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The start IP address of the firewall rule. Must be IPv6 format.
func (o LookupIPv6FirewallRuleResultOutput) StartIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) *string { return v.StartIPv6Address }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupIPv6FirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPv6FirewallRuleResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datalakeanalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Data Lake Analytics firewall rule.
// Azure REST API version: 2019-11-01-preview.
func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("azure-native:datalakeanalytics:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallRuleArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName string `pulumi:"accountName"`
	// The name of the firewall rule to retrieve.
	FirewallRuleName string `pulumi:"firewallRuleName"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Data Lake Analytics firewall rule information.
type LookupFirewallRuleResult struct {
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress string `pulumi:"endIpAddress"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress string `pulumi:"startIpAddress"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupFirewallRuleOutput(ctx *pulumi.Context, args LookupFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallRuleResult, error) {
			args := v.(LookupFirewallRuleArgs)
			r, err := LookupFirewallRule(ctx, &args, opts...)
			var s LookupFirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallRuleResultOutput)
}

type LookupFirewallRuleOutputArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the firewall rule to retrieve.
	FirewallRuleName pulumi.StringInput `pulumi:"firewallRuleName"`
	// The name of the Azure resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallRuleArgs)(nil)).Elem()
}

// Data Lake Analytics firewall rule information.
type LookupFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallRuleResult)(nil)).Elem()
}

func (o LookupFirewallRuleResultOutput) ToLookupFirewallRuleResultOutput() LookupFirewallRuleResultOutput {
	return o
}

func (o LookupFirewallRuleResultOutput) ToLookupFirewallRuleResultOutputWithContext(ctx context.Context) LookupFirewallRuleResultOutput {
	return o
}

// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
func (o LookupFirewallRuleResultOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.EndIpAddress }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
func (o LookupFirewallRuleResultOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.StartIpAddress }).(pulumi.StringOutput)
}

// The resource type.
func (o LookupFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallRuleResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a firewall rule.
func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("azure-native:sql/v20140401:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallRuleArgs struct {
	// The name of the firewall rule.
	FirewallRuleName string `pulumi:"firewallRuleName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a server firewall rule.
type LookupFirewallRuleResult struct {
	// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress. Use value '0.0.0.0' to represent all Azure-internal IP addresses.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Kind of server that contains this firewall rule.
	Kind string `pulumi:"kind"`
	// Location of the server that contains this firewall rule.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' to represent all Azure-internal IP addresses.
	StartIpAddress string `pulumi:"startIpAddress"`
	// Resource type.
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
	// The name of the firewall rule.
	FirewallRuleName pulumi.StringInput `pulumi:"firewallRuleName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallRuleArgs)(nil)).Elem()
}

// Represents a server firewall rule.
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

func (o LookupFirewallRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFirewallRuleResult] {
	return pulumix.Output[LookupFirewallRuleResult]{
		OutputState: o.OutputState,
	}
}

// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress. Use value '0.0.0.0' to represent all Azure-internal IP addresses.
func (o LookupFirewallRuleResultOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.EndIpAddress }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of server that contains this firewall rule.
func (o LookupFirewallRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Location of the server that contains this firewall rule.
func (o LookupFirewallRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' to represent all Azure-internal IP addresses.
func (o LookupFirewallRuleResultOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.StartIpAddress }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallRuleResultOutput{})
}

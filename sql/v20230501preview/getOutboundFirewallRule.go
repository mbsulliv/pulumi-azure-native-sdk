// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an outbound firewall rule.
func LookupOutboundFirewallRule(ctx *pulumi.Context, args *LookupOutboundFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupOutboundFirewallRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOutboundFirewallRuleResult
	err := ctx.Invoke("azure-native:sql/v20230501preview:getOutboundFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutboundFirewallRuleArgs struct {
	OutboundRuleFqdn string `pulumi:"outboundRuleFqdn"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// An Azure SQL DB Server Outbound Firewall Rule.
type LookupOutboundFirewallRuleResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The state of the outbound rule.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupOutboundFirewallRuleOutput(ctx *pulumi.Context, args LookupOutboundFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupOutboundFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOutboundFirewallRuleResult, error) {
			args := v.(LookupOutboundFirewallRuleArgs)
			r, err := LookupOutboundFirewallRule(ctx, &args, opts...)
			var s LookupOutboundFirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOutboundFirewallRuleResultOutput)
}

type LookupOutboundFirewallRuleOutputArgs struct {
	OutboundRuleFqdn pulumi.StringInput `pulumi:"outboundRuleFqdn"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupOutboundFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutboundFirewallRuleArgs)(nil)).Elem()
}

// An Azure SQL DB Server Outbound Firewall Rule.
type LookupOutboundFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupOutboundFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutboundFirewallRuleResult)(nil)).Elem()
}

func (o LookupOutboundFirewallRuleResultOutput) ToLookupOutboundFirewallRuleResultOutput() LookupOutboundFirewallRuleResultOutput {
	return o
}

func (o LookupOutboundFirewallRuleResultOutput) ToLookupOutboundFirewallRuleResultOutputWithContext(ctx context.Context) LookupOutboundFirewallRuleResultOutput {
	return o
}

// Resource ID.
func (o LookupOutboundFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupOutboundFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the outbound rule.
func (o LookupOutboundFirewallRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundFirewallRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupOutboundFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOutboundFirewallRuleResultOutput{})
}

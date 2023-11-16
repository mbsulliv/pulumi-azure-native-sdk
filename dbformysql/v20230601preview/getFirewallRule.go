// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a server firewall rule.
func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("azure-native:dbformysql/v20230601preview:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallRuleArgs struct {
	// The name of the server firewall rule.
	FirewallRuleName string `pulumi:"firewallRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a server firewall rule.
type LookupFirewallRuleResult struct {
	// The end IP address of the server firewall rule. Must be IPv4 format.
	EndIpAddress string `pulumi:"endIpAddress"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The start IP address of the server firewall rule. Must be IPv4 format.
	StartIpAddress string `pulumi:"startIpAddress"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
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
	// The name of the server firewall rule.
	FirewallRuleName pulumi.StringInput `pulumi:"firewallRuleName"`
	// The name of the resource group. The name is case insensitive.
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

// The end IP address of the server firewall rule. Must be IPv4 format.
func (o LookupFirewallRuleResultOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.EndIpAddress }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The start IP address of the server firewall rule. Must be IPv4 format.
func (o LookupFirewallRuleResultOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.StartIpAddress }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFirewallRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallRuleResultOutput{})
}

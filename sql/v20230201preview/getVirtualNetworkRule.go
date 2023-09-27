// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a virtual network rule.
func LookupVirtualNetworkRule(ctx *pulumi.Context, args *LookupVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:sql/v20230201preview:getVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkRuleArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the virtual network rule.
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}

// A virtual network rule.
type LookupVirtualNetworkRuleResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// Resource name.
	Name string `pulumi:"name"`
	// Virtual Network Rule State
	State string `pulumi:"state"`
	// Resource type.
	Type string `pulumi:"type"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId string `pulumi:"virtualNetworkSubnetId"`
}

func LookupVirtualNetworkRuleOutput(ctx *pulumi.Context, args LookupVirtualNetworkRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkRuleResult, error) {
			args := v.(LookupVirtualNetworkRuleArgs)
			r, err := LookupVirtualNetworkRule(ctx, &args, opts...)
			var s LookupVirtualNetworkRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkRuleResultOutput)
}

type LookupVirtualNetworkRuleOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// The name of the virtual network rule.
	VirtualNetworkRuleName pulumi.StringInput `pulumi:"virtualNetworkRuleName"`
}

func (LookupVirtualNetworkRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkRuleArgs)(nil)).Elem()
}

// A virtual network rule.
type LookupVirtualNetworkRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkRuleResult)(nil)).Elem()
}

func (o LookupVirtualNetworkRuleResultOutput) ToLookupVirtualNetworkRuleResultOutput() LookupVirtualNetworkRuleResultOutput {
	return o
}

func (o LookupVirtualNetworkRuleResultOutput) ToLookupVirtualNetworkRuleResultOutputWithContext(ctx context.Context) LookupVirtualNetworkRuleResultOutput {
	return o
}

func (o LookupVirtualNetworkRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualNetworkRuleResult] {
	return pulumix.Output[LookupVirtualNetworkRuleResult]{
		OutputState: o.OutputState,
	}
}

// Resource ID.
func (o LookupVirtualNetworkRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Create firewall rule before the virtual network has vnet service endpoint enabled.
func (o LookupVirtualNetworkRuleResultOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

// Resource name.
func (o LookupVirtualNetworkRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Virtual Network Rule State
func (o LookupVirtualNetworkRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.State }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupVirtualNetworkRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

// The ARM resource id of the virtual network subnet.
func (o LookupVirtualNetworkRuleResultOutput) VirtualNetworkSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.VirtualNetworkSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkRuleResultOutput{})
}

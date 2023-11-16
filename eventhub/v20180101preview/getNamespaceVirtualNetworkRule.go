// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an VirtualNetworkRule for a Namespace by rule name.
func LookupNamespaceVirtualNetworkRule(ctx *pulumi.Context, args *LookupNamespaceVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceVirtualNetworkRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamespaceVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:eventhub/v20180101preview:getNamespaceVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceVirtualNetworkRuleArgs struct {
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Virtual Network Rule name.
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}

// Single item in a List or Get VirtualNetworkRules operation
type LookupNamespaceVirtualNetworkRuleResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// ARM ID of Virtual Network Subnet
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}

func LookupNamespaceVirtualNetworkRuleOutput(ctx *pulumi.Context, args LookupNamespaceVirtualNetworkRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceVirtualNetworkRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceVirtualNetworkRuleResult, error) {
			args := v.(LookupNamespaceVirtualNetworkRuleArgs)
			r, err := LookupNamespaceVirtualNetworkRule(ctx, &args, opts...)
			var s LookupNamespaceVirtualNetworkRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceVirtualNetworkRuleResultOutput)
}

type LookupNamespaceVirtualNetworkRuleOutputArgs struct {
	// The Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Virtual Network Rule name.
	VirtualNetworkRuleName pulumi.StringInput `pulumi:"virtualNetworkRuleName"`
}

func (LookupNamespaceVirtualNetworkRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceVirtualNetworkRuleArgs)(nil)).Elem()
}

// Single item in a List or Get VirtualNetworkRules operation
type LookupNamespaceVirtualNetworkRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceVirtualNetworkRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceVirtualNetworkRuleResult)(nil)).Elem()
}

func (o LookupNamespaceVirtualNetworkRuleResultOutput) ToLookupNamespaceVirtualNetworkRuleResultOutput() LookupNamespaceVirtualNetworkRuleResultOutput {
	return o
}

func (o LookupNamespaceVirtualNetworkRuleResultOutput) ToLookupNamespaceVirtualNetworkRuleResultOutputWithContext(ctx context.Context) LookupNamespaceVirtualNetworkRuleResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupNamespaceVirtualNetworkRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceVirtualNetworkRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNamespaceVirtualNetworkRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceVirtualNetworkRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNamespaceVirtualNetworkRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceVirtualNetworkRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

// ARM ID of Virtual Network Subnet
func (o LookupNamespaceVirtualNetworkRuleResultOutput) VirtualNetworkSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceVirtualNetworkRuleResult) *string { return v.VirtualNetworkSubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceVirtualNetworkRuleResultOutput{})
}

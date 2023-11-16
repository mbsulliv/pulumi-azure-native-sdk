// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the current filter values for the signatures overrides
func ListFirewallPolicyIdpsSignaturesFilterValue(ctx *pulumi.Context, args *ListFirewallPolicyIdpsSignaturesFilterValueArgs, opts ...pulumi.InvokeOption) (*ListFirewallPolicyIdpsSignaturesFilterValueResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListFirewallPolicyIdpsSignaturesFilterValueResult
	err := ctx.Invoke("azure-native:network/v20230201:listFirewallPolicyIdpsSignaturesFilterValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFirewallPolicyIdpsSignaturesFilterValueArgs struct {
	// Describes the name of the column which values will be returned
	FilterName *string `pulumi:"filterName"`
	// The name of the Firewall Policy.
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes the list of all possible values for a specific filter value
type ListFirewallPolicyIdpsSignaturesFilterValueResult struct {
	// Describes the possible values
	FilterValues []string `pulumi:"filterValues"`
}

func ListFirewallPolicyIdpsSignaturesFilterValueOutput(ctx *pulumi.Context, args ListFirewallPolicyIdpsSignaturesFilterValueOutputArgs, opts ...pulumi.InvokeOption) ListFirewallPolicyIdpsSignaturesFilterValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFirewallPolicyIdpsSignaturesFilterValueResult, error) {
			args := v.(ListFirewallPolicyIdpsSignaturesFilterValueArgs)
			r, err := ListFirewallPolicyIdpsSignaturesFilterValue(ctx, &args, opts...)
			var s ListFirewallPolicyIdpsSignaturesFilterValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFirewallPolicyIdpsSignaturesFilterValueResultOutput)
}

type ListFirewallPolicyIdpsSignaturesFilterValueOutputArgs struct {
	// Describes the name of the column which values will be returned
	FilterName pulumi.StringPtrInput `pulumi:"filterName"`
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringInput `pulumi:"firewallPolicyName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListFirewallPolicyIdpsSignaturesFilterValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallPolicyIdpsSignaturesFilterValueArgs)(nil)).Elem()
}

// Describes the list of all possible values for a specific filter value
type ListFirewallPolicyIdpsSignaturesFilterValueResultOutput struct{ *pulumi.OutputState }

func (ListFirewallPolicyIdpsSignaturesFilterValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallPolicyIdpsSignaturesFilterValueResult)(nil)).Elem()
}

func (o ListFirewallPolicyIdpsSignaturesFilterValueResultOutput) ToListFirewallPolicyIdpsSignaturesFilterValueResultOutput() ListFirewallPolicyIdpsSignaturesFilterValueResultOutput {
	return o
}

func (o ListFirewallPolicyIdpsSignaturesFilterValueResultOutput) ToListFirewallPolicyIdpsSignaturesFilterValueResultOutputWithContext(ctx context.Context) ListFirewallPolicyIdpsSignaturesFilterValueResultOutput {
	return o
}

// Describes the possible values
func (o ListFirewallPolicyIdpsSignaturesFilterValueResultOutput) FilterValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListFirewallPolicyIdpsSignaturesFilterValueResult) []string { return v.FilterValues }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFirewallPolicyIdpsSignaturesFilterValueResultOutput{})
}

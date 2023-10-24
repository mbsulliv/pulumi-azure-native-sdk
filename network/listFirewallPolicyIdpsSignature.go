// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the current status of IDPS signatures for the relevant policy
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2021-08-01, 2023-04-01, 2023-05-01.
func ListFirewallPolicyIdpsSignature(ctx *pulumi.Context, args *ListFirewallPolicyIdpsSignatureArgs, opts ...pulumi.InvokeOption) (*ListFirewallPolicyIdpsSignatureResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListFirewallPolicyIdpsSignatureResult
	err := ctx.Invoke("azure-native:network:listFirewallPolicyIdpsSignature", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFirewallPolicyIdpsSignatureArgs struct {
	// Contain all filters names and values
	Filters []FilterItems `pulumi:"filters"`
	// The name of the Firewall Policy.
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	// Column to sort response by
	OrderBy *OrderBy `pulumi:"orderBy"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of the results to return in each page
	ResultsPerPage *int `pulumi:"resultsPerPage"`
	// Search term in all columns
	Search *string `pulumi:"search"`
	// The number of records matching the filter to skip
	Skip *int `pulumi:"skip"`
}

// Query result
type ListFirewallPolicyIdpsSignatureResult struct {
	// Number of total records matching the query.
	MatchingRecordsCount *float64 `pulumi:"matchingRecordsCount"`
	// Array containing the results of the query
	Signatures []SingleQueryResultResponse `pulumi:"signatures"`
}

func ListFirewallPolicyIdpsSignatureOutput(ctx *pulumi.Context, args ListFirewallPolicyIdpsSignatureOutputArgs, opts ...pulumi.InvokeOption) ListFirewallPolicyIdpsSignatureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFirewallPolicyIdpsSignatureResult, error) {
			args := v.(ListFirewallPolicyIdpsSignatureArgs)
			r, err := ListFirewallPolicyIdpsSignature(ctx, &args, opts...)
			var s ListFirewallPolicyIdpsSignatureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFirewallPolicyIdpsSignatureResultOutput)
}

type ListFirewallPolicyIdpsSignatureOutputArgs struct {
	// Contain all filters names and values
	Filters FilterItemsArrayInput `pulumi:"filters"`
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringInput `pulumi:"firewallPolicyName"`
	// Column to sort response by
	OrderBy OrderByPtrInput `pulumi:"orderBy"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The number of the results to return in each page
	ResultsPerPage pulumi.IntPtrInput `pulumi:"resultsPerPage"`
	// Search term in all columns
	Search pulumi.StringPtrInput `pulumi:"search"`
	// The number of records matching the filter to skip
	Skip pulumi.IntPtrInput `pulumi:"skip"`
}

func (ListFirewallPolicyIdpsSignatureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallPolicyIdpsSignatureArgs)(nil)).Elem()
}

// Query result
type ListFirewallPolicyIdpsSignatureResultOutput struct{ *pulumi.OutputState }

func (ListFirewallPolicyIdpsSignatureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallPolicyIdpsSignatureResult)(nil)).Elem()
}

func (o ListFirewallPolicyIdpsSignatureResultOutput) ToListFirewallPolicyIdpsSignatureResultOutput() ListFirewallPolicyIdpsSignatureResultOutput {
	return o
}

func (o ListFirewallPolicyIdpsSignatureResultOutput) ToListFirewallPolicyIdpsSignatureResultOutputWithContext(ctx context.Context) ListFirewallPolicyIdpsSignatureResultOutput {
	return o
}

func (o ListFirewallPolicyIdpsSignatureResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListFirewallPolicyIdpsSignatureResult] {
	return pulumix.Output[ListFirewallPolicyIdpsSignatureResult]{
		OutputState: o.OutputState,
	}
}

// Number of total records matching the query.
func (o ListFirewallPolicyIdpsSignatureResultOutput) MatchingRecordsCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListFirewallPolicyIdpsSignatureResult) *float64 { return v.MatchingRecordsCount }).(pulumi.Float64PtrOutput)
}

// Array containing the results of the query
func (o ListFirewallPolicyIdpsSignatureResultOutput) Signatures() SingleQueryResultResponseArrayOutput {
	return o.ApplyT(func(v ListFirewallPolicyIdpsSignatureResult) []SingleQueryResultResponse { return v.Signatures }).(SingleQueryResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFirewallPolicyIdpsSignatureResultOutput{})
}

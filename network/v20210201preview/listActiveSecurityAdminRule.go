// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists active security admin rules in a network manager.
func ListActiveSecurityAdminRule(ctx *pulumi.Context, args *ListActiveSecurityAdminRuleArgs, opts ...pulumi.InvokeOption) (*ListActiveSecurityAdminRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListActiveSecurityAdminRuleResult
	err := ctx.Invoke("azure-native:network/v20210201preview:listActiveSecurityAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveSecurityAdminRuleArgs struct {
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// List of regions.
	Regions []string `pulumi:"regions"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
}

// Result of the request to list active security admin rules. It contains a list of active security admin rules and a skiptoken to get the next set of results.
type ListActiveSecurityAdminRuleResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of active security admin rules.
	Value []interface{} `pulumi:"value"`
}

func ListActiveSecurityAdminRuleOutput(ctx *pulumi.Context, args ListActiveSecurityAdminRuleOutputArgs, opts ...pulumi.InvokeOption) ListActiveSecurityAdminRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveSecurityAdminRuleResult, error) {
			args := v.(ListActiveSecurityAdminRuleArgs)
			r, err := ListActiveSecurityAdminRule(ctx, &args, opts...)
			var s ListActiveSecurityAdminRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveSecurityAdminRuleResultOutput)
}

type ListActiveSecurityAdminRuleOutputArgs struct {
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// List of regions.
	Regions pulumi.StringArrayInput `pulumi:"regions"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListActiveSecurityAdminRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityAdminRuleArgs)(nil)).Elem()
}

// Result of the request to list active security admin rules. It contains a list of active security admin rules and a skiptoken to get the next set of results.
type ListActiveSecurityAdminRuleResultOutput struct{ *pulumi.OutputState }

func (ListActiveSecurityAdminRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityAdminRuleResult)(nil)).Elem()
}

func (o ListActiveSecurityAdminRuleResultOutput) ToListActiveSecurityAdminRuleResultOutput() ListActiveSecurityAdminRuleResultOutput {
	return o
}

func (o ListActiveSecurityAdminRuleResultOutput) ToListActiveSecurityAdminRuleResultOutputWithContext(ctx context.Context) ListActiveSecurityAdminRuleResultOutput {
	return o
}

// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
func (o ListActiveSecurityAdminRuleResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveSecurityAdminRuleResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// Gets a page of active security admin rules.
func (o ListActiveSecurityAdminRuleResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListActiveSecurityAdminRuleResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveSecurityAdminRuleResultOutput{})
}

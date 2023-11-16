// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List all effective security admin rules applied on a virtual network.
func ListNetworkManagerEffectiveSecurityAdminRules(ctx *pulumi.Context, args *ListNetworkManagerEffectiveSecurityAdminRulesArgs, opts ...pulumi.InvokeOption) (*ListNetworkManagerEffectiveSecurityAdminRulesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListNetworkManagerEffectiveSecurityAdminRulesResult
	err := ctx.Invoke("azure-native:network/v20230401:listNetworkManagerEffectiveSecurityAdminRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNetworkManagerEffectiveSecurityAdminRulesArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// An optional query parameter which specifies the maximum number of records to be returned by the server.
	Top *int `pulumi:"top"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// Result of the request to list networkManagerEffectiveSecurityAdminRules. It contains a list of groups and a skiptoken to get the next set of results.
type ListNetworkManagerEffectiveSecurityAdminRulesResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of NetworkManagerEffectiveSecurityAdminRules
	Value []interface{} `pulumi:"value"`
}

func ListNetworkManagerEffectiveSecurityAdminRulesOutput(ctx *pulumi.Context, args ListNetworkManagerEffectiveSecurityAdminRulesOutputArgs, opts ...pulumi.InvokeOption) ListNetworkManagerEffectiveSecurityAdminRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNetworkManagerEffectiveSecurityAdminRulesResult, error) {
			args := v.(ListNetworkManagerEffectiveSecurityAdminRulesArgs)
			r, err := ListNetworkManagerEffectiveSecurityAdminRules(ctx, &args, opts...)
			var s ListNetworkManagerEffectiveSecurityAdminRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNetworkManagerEffectiveSecurityAdminRulesResultOutput)
}

type ListNetworkManagerEffectiveSecurityAdminRulesOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
	// An optional query parameter which specifies the maximum number of records to be returned by the server.
	Top pulumi.IntPtrInput `pulumi:"top"`
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (ListNetworkManagerEffectiveSecurityAdminRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveSecurityAdminRulesArgs)(nil)).Elem()
}

// Result of the request to list networkManagerEffectiveSecurityAdminRules. It contains a list of groups and a skiptoken to get the next set of results.
type ListNetworkManagerEffectiveSecurityAdminRulesResultOutput struct{ *pulumi.OutputState }

func (ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveSecurityAdminRulesResult)(nil)).Elem()
}

func (o ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) ToListNetworkManagerEffectiveSecurityAdminRulesResultOutput() ListNetworkManagerEffectiveSecurityAdminRulesResultOutput {
	return o
}

func (o ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) ToListNetworkManagerEffectiveSecurityAdminRulesResultOutputWithContext(ctx context.Context) ListNetworkManagerEffectiveSecurityAdminRulesResultOutput {
	return o
}

// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
func (o ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveSecurityAdminRulesResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// Gets a page of NetworkManagerEffectiveSecurityAdminRules
func (o ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveSecurityAdminRulesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNetworkManagerEffectiveSecurityAdminRulesResultOutput{})
}

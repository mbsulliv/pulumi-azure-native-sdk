// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists active security admin rules in a network manager.
func ListActiveSecurityAdminRules(ctx *pulumi.Context, args *ListActiveSecurityAdminRulesArgs, opts ...pulumi.InvokeOption) (*ListActiveSecurityAdminRulesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListActiveSecurityAdminRulesResult
	err := ctx.Invoke("azure-native:network/v20230501:listActiveSecurityAdminRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveSecurityAdminRulesArgs struct {
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// List of regions.
	Regions []string `pulumi:"regions"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// An optional query parameter which specifies the maximum number of records to be returned by the server.
	Top *int `pulumi:"top"`
}

// Result of the request to list active security admin rules. It contains a list of active security admin rules and a skiptoken to get the next set of results.
type ListActiveSecurityAdminRulesResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of active security admin rules.
	Value []interface{} `pulumi:"value"`
}

func ListActiveSecurityAdminRulesOutput(ctx *pulumi.Context, args ListActiveSecurityAdminRulesOutputArgs, opts ...pulumi.InvokeOption) ListActiveSecurityAdminRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveSecurityAdminRulesResult, error) {
			args := v.(ListActiveSecurityAdminRulesArgs)
			r, err := ListActiveSecurityAdminRules(ctx, &args, opts...)
			var s ListActiveSecurityAdminRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveSecurityAdminRulesResultOutput)
}

type ListActiveSecurityAdminRulesOutputArgs struct {
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// List of regions.
	Regions pulumi.StringArrayInput `pulumi:"regions"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
	// An optional query parameter which specifies the maximum number of records to be returned by the server.
	Top pulumi.IntPtrInput `pulumi:"top"`
}

func (ListActiveSecurityAdminRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityAdminRulesArgs)(nil)).Elem()
}

// Result of the request to list active security admin rules. It contains a list of active security admin rules and a skiptoken to get the next set of results.
type ListActiveSecurityAdminRulesResultOutput struct{ *pulumi.OutputState }

func (ListActiveSecurityAdminRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityAdminRulesResult)(nil)).Elem()
}

func (o ListActiveSecurityAdminRulesResultOutput) ToListActiveSecurityAdminRulesResultOutput() ListActiveSecurityAdminRulesResultOutput {
	return o
}

func (o ListActiveSecurityAdminRulesResultOutput) ToListActiveSecurityAdminRulesResultOutputWithContext(ctx context.Context) ListActiveSecurityAdminRulesResultOutput {
	return o
}

// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
func (o ListActiveSecurityAdminRulesResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveSecurityAdminRulesResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// Gets a page of active security admin rules.
func (o ListActiveSecurityAdminRulesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListActiveSecurityAdminRulesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveSecurityAdminRulesResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package search

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the primary and secondary admin API keys for the specified Azure Cognitive Search service.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2015-02-28, 2021-04-01-preview, 2023-11-01.
func ListAdminKey(ctx *pulumi.Context, args *ListAdminKeyArgs, opts ...pulumi.InvokeOption) (*ListAdminKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAdminKeyResult
	err := ctx.Invoke("azure-native:search:listAdminKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAdminKeyArgs struct {
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Cognitive Search service associated with the specified resource group.
	SearchServiceName string `pulumi:"searchServiceName"`
}

// Response containing the primary and secondary admin API keys for a given Azure Cognitive Search service.
type ListAdminKeyResult struct {
	// The primary admin API key of the search service.
	PrimaryKey string `pulumi:"primaryKey"`
	// The secondary admin API key of the search service.
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListAdminKeyOutput(ctx *pulumi.Context, args ListAdminKeyOutputArgs, opts ...pulumi.InvokeOption) ListAdminKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAdminKeyResult, error) {
			args := v.(ListAdminKeyArgs)
			r, err := ListAdminKey(ctx, &args, opts...)
			var s ListAdminKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAdminKeyResultOutput)
}

type ListAdminKeyOutputArgs struct {
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Azure Cognitive Search service associated with the specified resource group.
	SearchServiceName pulumi.StringInput `pulumi:"searchServiceName"`
}

func (ListAdminKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAdminKeyArgs)(nil)).Elem()
}

// Response containing the primary and secondary admin API keys for a given Azure Cognitive Search service.
type ListAdminKeyResultOutput struct{ *pulumi.OutputState }

func (ListAdminKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAdminKeyResult)(nil)).Elem()
}

func (o ListAdminKeyResultOutput) ToListAdminKeyResultOutput() ListAdminKeyResultOutput {
	return o
}

func (o ListAdminKeyResultOutput) ToListAdminKeyResultOutputWithContext(ctx context.Context) ListAdminKeyResultOutput {
	return o
}

// The primary admin API key of the search service.
func (o ListAdminKeyResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListAdminKeyResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The secondary admin API key of the search service.
func (o ListAdminKeyResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListAdminKeyResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAdminKeyResultOutput{})
}

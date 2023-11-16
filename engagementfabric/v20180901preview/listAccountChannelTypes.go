// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of the EngagementFabric channel descriptions
func ListAccountChannelTypes(ctx *pulumi.Context, args *ListAccountChannelTypesArgs, opts ...pulumi.InvokeOption) (*ListAccountChannelTypesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAccountChannelTypesResult
	err := ctx.Invoke("azure-native:engagementfabric/v20180901preview:listAccountChannelTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountChannelTypesArgs struct {
	// Account Name
	AccountName string `pulumi:"accountName"`
	// Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of the EngagementFabric channel descriptions
type ListAccountChannelTypesResult struct {
	// Channel descriptions
	Value []ChannelTypeDescriptionResponse `pulumi:"value"`
}

func ListAccountChannelTypesOutput(ctx *pulumi.Context, args ListAccountChannelTypesOutputArgs, opts ...pulumi.InvokeOption) ListAccountChannelTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAccountChannelTypesResult, error) {
			args := v.(ListAccountChannelTypesArgs)
			r, err := ListAccountChannelTypes(ctx, &args, opts...)
			var s ListAccountChannelTypesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAccountChannelTypesResultOutput)
}

type ListAccountChannelTypesOutputArgs struct {
	// Account Name
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Resource Group Name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAccountChannelTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountChannelTypesArgs)(nil)).Elem()
}

// List of the EngagementFabric channel descriptions
type ListAccountChannelTypesResultOutput struct{ *pulumi.OutputState }

func (ListAccountChannelTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountChannelTypesResult)(nil)).Elem()
}

func (o ListAccountChannelTypesResultOutput) ToListAccountChannelTypesResultOutput() ListAccountChannelTypesResultOutput {
	return o
}

func (o ListAccountChannelTypesResultOutput) ToListAccountChannelTypesResultOutputWithContext(ctx context.Context) ListAccountChannelTypesResultOutput {
	return o
}

// Channel descriptions
func (o ListAccountChannelTypesResultOutput) Value() ChannelTypeDescriptionResponseArrayOutput {
	return o.ApplyT(func(v ListAccountChannelTypesResult) []ChannelTypeDescriptionResponse { return v.Value }).(ChannelTypeDescriptionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountChannelTypesResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240125

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists all pending flows for a connection.
func ListListPendingFlow(ctx *pulumi.Context, args *ListListPendingFlowArgs, opts ...pulumi.InvokeOption) (*ListListPendingFlowResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListListPendingFlowResult
	err := ctx.Invoke("azure-native:azuredatatransfer/v20240125:listListPendingFlow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListPendingFlowArgs struct {
	// The name for the connection that is to be requested.
	ConnectionName string `pulumi:"connectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connections list result.
type ListListPendingFlowResult struct {
	// Link to next results
	NextLink *string `pulumi:"nextLink"`
	// flows array.
	Value []PendingFlowResponse `pulumi:"value"`
}

func ListListPendingFlowOutput(ctx *pulumi.Context, args ListListPendingFlowOutputArgs, opts ...pulumi.InvokeOption) ListListPendingFlowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListListPendingFlowResult, error) {
			args := v.(ListListPendingFlowArgs)
			r, err := ListListPendingFlow(ctx, &args, opts...)
			var s ListListPendingFlowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListListPendingFlowResultOutput)
}

type ListListPendingFlowOutputArgs struct {
	// The name for the connection that is to be requested.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListListPendingFlowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListPendingFlowArgs)(nil)).Elem()
}

// The connections list result.
type ListListPendingFlowResultOutput struct{ *pulumi.OutputState }

func (ListListPendingFlowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListPendingFlowResult)(nil)).Elem()
}

func (o ListListPendingFlowResultOutput) ToListListPendingFlowResultOutput() ListListPendingFlowResultOutput {
	return o
}

func (o ListListPendingFlowResultOutput) ToListListPendingFlowResultOutputWithContext(ctx context.Context) ListListPendingFlowResultOutput {
	return o
}

// Link to next results
func (o ListListPendingFlowResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListListPendingFlowResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// flows array.
func (o ListListPendingFlowResultOutput) Value() PendingFlowResponseArrayOutput {
	return o.ApplyT(func(v ListListPendingFlowResult) []PendingFlowResponse { return v.Value }).(PendingFlowResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListPendingFlowResultOutput{})
}

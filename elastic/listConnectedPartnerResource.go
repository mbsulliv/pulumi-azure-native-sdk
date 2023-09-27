// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elastic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

//	List of all active elastic deployments.
//
// Azure REST API version: 2023-07-01-preview.
func ListConnectedPartnerResource(ctx *pulumi.Context, args *ListConnectedPartnerResourceArgs, opts ...pulumi.InvokeOption) (*ListConnectedPartnerResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListConnectedPartnerResourceResult
	err := ctx.Invoke("azure-native:elastic:listConnectedPartnerResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectedPartnerResourceArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of all active elastic deployments.
type ListConnectedPartnerResourceResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// Results of a list operation.
	Value []ConnectedPartnerResourcesListFormatResponse `pulumi:"value"`
}

func ListConnectedPartnerResourceOutput(ctx *pulumi.Context, args ListConnectedPartnerResourceOutputArgs, opts ...pulumi.InvokeOption) ListConnectedPartnerResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConnectedPartnerResourceResult, error) {
			args := v.(ListConnectedPartnerResourceArgs)
			r, err := ListConnectedPartnerResource(ctx, &args, opts...)
			var s ListConnectedPartnerResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConnectedPartnerResourceResultOutput)
}

type ListConnectedPartnerResourceOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListConnectedPartnerResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedPartnerResourceArgs)(nil)).Elem()
}

// List of all active elastic deployments.
type ListConnectedPartnerResourceResultOutput struct{ *pulumi.OutputState }

func (ListConnectedPartnerResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectedPartnerResourceResult)(nil)).Elem()
}

func (o ListConnectedPartnerResourceResultOutput) ToListConnectedPartnerResourceResultOutput() ListConnectedPartnerResourceResultOutput {
	return o
}

func (o ListConnectedPartnerResourceResultOutput) ToListConnectedPartnerResourceResultOutputWithContext(ctx context.Context) ListConnectedPartnerResourceResultOutput {
	return o
}

func (o ListConnectedPartnerResourceResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListConnectedPartnerResourceResult] {
	return pulumix.Output[ListConnectedPartnerResourceResult]{
		OutputState: o.OutputState,
	}
}

// Link to the next set of results, if any.
func (o ListConnectedPartnerResourceResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListConnectedPartnerResourceResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Results of a list operation.
func (o ListConnectedPartnerResourceResultOutput) Value() ConnectedPartnerResourcesListFormatResponseArrayOutput {
	return o.ApplyT(func(v ListConnectedPartnerResourceResult) []ConnectedPartnerResourcesListFormatResponse {
		return v.Value
	}).(ConnectedPartnerResourcesListFormatResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectedPartnerResourceResultOutput{})
}

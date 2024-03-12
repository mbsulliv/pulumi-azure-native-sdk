// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package orbital

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a list of L2 Connections attached to an edge site.
// Azure REST API version: 2024-03-01-preview.
//
// Other available API versions: 2024-03-01.
func ListEdgeSiteL2Connections(ctx *pulumi.Context, args *ListEdgeSiteL2ConnectionsArgs, opts ...pulumi.InvokeOption) (*ListEdgeSiteL2ConnectionsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListEdgeSiteL2ConnectionsResult
	err := ctx.Invoke("azure-native:orbital:listEdgeSiteL2Connections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEdgeSiteL2ConnectionsArgs struct {
	// Edge site name.
	EdgeSiteName string `pulumi:"edgeSiteName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for an API service call that lists the resource IDs of resources associated with another resource.
type ListEdgeSiteL2ConnectionsResult struct {
	// The URL to get the next set of results.
	NextLink string `pulumi:"nextLink"`
	// A list of Azure Resource IDs.
	Value []ResourceIdListResultResponseValue `pulumi:"value"`
}

func ListEdgeSiteL2ConnectionsOutput(ctx *pulumi.Context, args ListEdgeSiteL2ConnectionsOutputArgs, opts ...pulumi.InvokeOption) ListEdgeSiteL2ConnectionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEdgeSiteL2ConnectionsResult, error) {
			args := v.(ListEdgeSiteL2ConnectionsArgs)
			r, err := ListEdgeSiteL2Connections(ctx, &args, opts...)
			var s ListEdgeSiteL2ConnectionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEdgeSiteL2ConnectionsResultOutput)
}

type ListEdgeSiteL2ConnectionsOutputArgs struct {
	// Edge site name.
	EdgeSiteName pulumi.StringInput `pulumi:"edgeSiteName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListEdgeSiteL2ConnectionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEdgeSiteL2ConnectionsArgs)(nil)).Elem()
}

// Response for an API service call that lists the resource IDs of resources associated with another resource.
type ListEdgeSiteL2ConnectionsResultOutput struct{ *pulumi.OutputState }

func (ListEdgeSiteL2ConnectionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEdgeSiteL2ConnectionsResult)(nil)).Elem()
}

func (o ListEdgeSiteL2ConnectionsResultOutput) ToListEdgeSiteL2ConnectionsResultOutput() ListEdgeSiteL2ConnectionsResultOutput {
	return o
}

func (o ListEdgeSiteL2ConnectionsResultOutput) ToListEdgeSiteL2ConnectionsResultOutputWithContext(ctx context.Context) ListEdgeSiteL2ConnectionsResultOutput {
	return o
}

// The URL to get the next set of results.
func (o ListEdgeSiteL2ConnectionsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListEdgeSiteL2ConnectionsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// A list of Azure Resource IDs.
func (o ListEdgeSiteL2ConnectionsResultOutput) Value() ResourceIdListResultResponseValueArrayOutput {
	return o.ApplyT(func(v ListEdgeSiteL2ConnectionsResult) []ResourceIdListResultResponseValue { return v.Value }).(ResourceIdListResultResponseValueArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEdgeSiteL2ConnectionsResultOutput{})
}

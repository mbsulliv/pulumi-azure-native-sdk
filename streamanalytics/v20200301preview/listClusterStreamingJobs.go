// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Lists all of the streaming jobs in the given cluster.
func ListClusterStreamingJobs(ctx *pulumi.Context, args *ListClusterStreamingJobsArgs, opts ...pulumi.InvokeOption) (*ListClusterStreamingJobsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListClusterStreamingJobsResult
	err := ctx.Invoke("azure-native:streamanalytics/v20200301preview:listClusterStreamingJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterStreamingJobsArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A list of streaming jobs. Populated by a List operation.
type ListClusterStreamingJobsResult struct {
	// The URL to fetch the next set of streaming jobs.
	NextLink string `pulumi:"nextLink"`
	// A list of streaming jobs.
	Value []ClusterJobResponse `pulumi:"value"`
}

func ListClusterStreamingJobsOutput(ctx *pulumi.Context, args ListClusterStreamingJobsOutputArgs, opts ...pulumi.InvokeOption) ListClusterStreamingJobsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListClusterStreamingJobsResult, error) {
			args := v.(ListClusterStreamingJobsArgs)
			r, err := ListClusterStreamingJobs(ctx, &args, opts...)
			var s ListClusterStreamingJobsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListClusterStreamingJobsResultOutput)
}

type ListClusterStreamingJobsOutputArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListClusterStreamingJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterStreamingJobsArgs)(nil)).Elem()
}

// A list of streaming jobs. Populated by a List operation.
type ListClusterStreamingJobsResultOutput struct{ *pulumi.OutputState }

func (ListClusterStreamingJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterStreamingJobsResult)(nil)).Elem()
}

func (o ListClusterStreamingJobsResultOutput) ToListClusterStreamingJobsResultOutput() ListClusterStreamingJobsResultOutput {
	return o
}

func (o ListClusterStreamingJobsResultOutput) ToListClusterStreamingJobsResultOutputWithContext(ctx context.Context) ListClusterStreamingJobsResultOutput {
	return o
}

func (o ListClusterStreamingJobsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListClusterStreamingJobsResult] {
	return pulumix.Output[ListClusterStreamingJobsResult]{
		OutputState: o.OutputState,
	}
}

// The URL to fetch the next set of streaming jobs.
func (o ListClusterStreamingJobsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListClusterStreamingJobsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// A list of streaming jobs.
func (o ListClusterStreamingJobsResultOutput) Value() ClusterJobResponseArrayOutput {
	return o.ApplyT(func(v ListClusterStreamingJobsResult) []ClusterJobResponse { return v.Value }).(ClusterJobResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterStreamingJobsResultOutput{})
}

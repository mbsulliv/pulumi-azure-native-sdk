// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response of a list operation.
// Azure REST API version: 2022-01-01-preview.
func ListSubAccountMonitoredResources(ctx *pulumi.Context, args *ListSubAccountMonitoredResourcesArgs, opts ...pulumi.InvokeOption) (*ListSubAccountMonitoredResourcesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSubAccountMonitoredResourcesResult
	err := ctx.Invoke("azure-native:logz:listSubAccountMonitoredResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubAccountMonitoredResourcesArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sub Account resource name
	SubAccountName string `pulumi:"subAccountName"`
}

// Response of a list operation.
type ListSubAccountMonitoredResourcesResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// Results of a list operation.
	Value []MonitoredResourceResponse `pulumi:"value"`
}

func ListSubAccountMonitoredResourcesOutput(ctx *pulumi.Context, args ListSubAccountMonitoredResourcesOutputArgs, opts ...pulumi.InvokeOption) ListSubAccountMonitoredResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSubAccountMonitoredResourcesResult, error) {
			args := v.(ListSubAccountMonitoredResourcesArgs)
			r, err := ListSubAccountMonitoredResources(ctx, &args, opts...)
			var s ListSubAccountMonitoredResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSubAccountMonitoredResourcesResultOutput)
}

type ListSubAccountMonitoredResourcesOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Sub Account resource name
	SubAccountName pulumi.StringInput `pulumi:"subAccountName"`
}

func (ListSubAccountMonitoredResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubAccountMonitoredResourcesArgs)(nil)).Elem()
}

// Response of a list operation.
type ListSubAccountMonitoredResourcesResultOutput struct{ *pulumi.OutputState }

func (ListSubAccountMonitoredResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubAccountMonitoredResourcesResult)(nil)).Elem()
}

func (o ListSubAccountMonitoredResourcesResultOutput) ToListSubAccountMonitoredResourcesResultOutput() ListSubAccountMonitoredResourcesResultOutput {
	return o
}

func (o ListSubAccountMonitoredResourcesResultOutput) ToListSubAccountMonitoredResourcesResultOutputWithContext(ctx context.Context) ListSubAccountMonitoredResourcesResultOutput {
	return o
}

func (o ListSubAccountMonitoredResourcesResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListSubAccountMonitoredResourcesResult] {
	return pulumix.Output[ListSubAccountMonitoredResourcesResult]{
		OutputState: o.OutputState,
	}
}

// Link to the next set of results, if any.
func (o ListSubAccountMonitoredResourcesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSubAccountMonitoredResourcesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Results of a list operation.
func (o ListSubAccountMonitoredResourcesResultOutput) Value() MonitoredResourceResponseArrayOutput {
	return o.ApplyT(func(v ListSubAccountMonitoredResourcesResult) []MonitoredResourceResponse { return v.Value }).(MonitoredResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSubAccountMonitoredResourcesResultOutput{})
}

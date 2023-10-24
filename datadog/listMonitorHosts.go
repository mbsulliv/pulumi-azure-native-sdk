// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datadog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response of a list operation.
// Azure REST API version: 2022-06-01.
//
// Other available API versions: 2022-08-01, 2023-01-01.
func ListMonitorHosts(ctx *pulumi.Context, args *ListMonitorHostsArgs, opts ...pulumi.InvokeOption) (*ListMonitorHostsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListMonitorHostsResult
	err := ctx.Invoke("azure-native:datadog:listMonitorHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorHostsArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response of a list operation.
type ListMonitorHostsResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// Results of a list operation.
	Value []DatadogHostResponse `pulumi:"value"`
}

func ListMonitorHostsOutput(ctx *pulumi.Context, args ListMonitorHostsOutputArgs, opts ...pulumi.InvokeOption) ListMonitorHostsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorHostsResult, error) {
			args := v.(ListMonitorHostsArgs)
			r, err := ListMonitorHosts(ctx, &args, opts...)
			var s ListMonitorHostsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorHostsResultOutput)
}

type ListMonitorHostsOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitorHostsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorHostsArgs)(nil)).Elem()
}

// Response of a list operation.
type ListMonitorHostsResultOutput struct{ *pulumi.OutputState }

func (ListMonitorHostsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorHostsResult)(nil)).Elem()
}

func (o ListMonitorHostsResultOutput) ToListMonitorHostsResultOutput() ListMonitorHostsResultOutput {
	return o
}

func (o ListMonitorHostsResultOutput) ToListMonitorHostsResultOutputWithContext(ctx context.Context) ListMonitorHostsResultOutput {
	return o
}

func (o ListMonitorHostsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListMonitorHostsResult] {
	return pulumix.Output[ListMonitorHostsResult]{
		OutputState: o.OutputState,
	}
}

// Link to the next set of results, if any.
func (o ListMonitorHostsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorHostsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Results of a list operation.
func (o ListMonitorHostsResultOutput) Value() DatadogHostResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorHostsResult) []DatadogHostResponse { return v.Value }).(DatadogHostResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorHostsResultOutput{})
}

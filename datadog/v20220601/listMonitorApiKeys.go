// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response of a list operation.
func ListMonitorApiKeys(ctx *pulumi.Context, args *ListMonitorApiKeysArgs, opts ...pulumi.InvokeOption) (*ListMonitorApiKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListMonitorApiKeysResult
	err := ctx.Invoke("azure-native:datadog/v20220601:listMonitorApiKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorApiKeysArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response of a list operation.
type ListMonitorApiKeysResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// Results of a list operation.
	Value []DatadogApiKeyResponse `pulumi:"value"`
}

func ListMonitorApiKeysOutput(ctx *pulumi.Context, args ListMonitorApiKeysOutputArgs, opts ...pulumi.InvokeOption) ListMonitorApiKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorApiKeysResult, error) {
			args := v.(ListMonitorApiKeysArgs)
			r, err := ListMonitorApiKeys(ctx, &args, opts...)
			var s ListMonitorApiKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorApiKeysResultOutput)
}

type ListMonitorApiKeysOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitorApiKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorApiKeysArgs)(nil)).Elem()
}

// Response of a list operation.
type ListMonitorApiKeysResultOutput struct{ *pulumi.OutputState }

func (ListMonitorApiKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorApiKeysResult)(nil)).Elem()
}

func (o ListMonitorApiKeysResultOutput) ToListMonitorApiKeysResultOutput() ListMonitorApiKeysResultOutput {
	return o
}

func (o ListMonitorApiKeysResultOutput) ToListMonitorApiKeysResultOutputWithContext(ctx context.Context) ListMonitorApiKeysResultOutput {
	return o
}

func (o ListMonitorApiKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListMonitorApiKeysResult] {
	return pulumix.Output[ListMonitorApiKeysResult]{
		OutputState: o.OutputState,
	}
}

// Link to the next set of results, if any.
func (o ListMonitorApiKeysResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorApiKeysResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Results of a list operation.
func (o ListMonitorApiKeysResultOutput) Value() DatadogApiKeyResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorApiKeysResult) []DatadogApiKeyResponse { return v.Value }).(DatadogApiKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorApiKeysResultOutput{})
}

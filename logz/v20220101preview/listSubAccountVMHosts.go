// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response of a list VM Host Update Operation.
func ListSubAccountVMHosts(ctx *pulumi.Context, args *ListSubAccountVMHostsArgs, opts ...pulumi.InvokeOption) (*ListSubAccountVMHostsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSubAccountVMHostsResult
	err := ctx.Invoke("azure-native:logz/v20220101preview:listSubAccountVMHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubAccountVMHostsArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sub Account resource name
	SubAccountName string `pulumi:"subAccountName"`
}

// Response of a list VM Host Update Operation.
type ListSubAccountVMHostsResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// Response of a list vm host update operation.
	Value []VMResourcesResponse `pulumi:"value"`
}

func ListSubAccountVMHostsOutput(ctx *pulumi.Context, args ListSubAccountVMHostsOutputArgs, opts ...pulumi.InvokeOption) ListSubAccountVMHostsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSubAccountVMHostsResult, error) {
			args := v.(ListSubAccountVMHostsArgs)
			r, err := ListSubAccountVMHosts(ctx, &args, opts...)
			var s ListSubAccountVMHostsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSubAccountVMHostsResultOutput)
}

type ListSubAccountVMHostsOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Sub Account resource name
	SubAccountName pulumi.StringInput `pulumi:"subAccountName"`
}

func (ListSubAccountVMHostsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubAccountVMHostsArgs)(nil)).Elem()
}

// Response of a list VM Host Update Operation.
type ListSubAccountVMHostsResultOutput struct{ *pulumi.OutputState }

func (ListSubAccountVMHostsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubAccountVMHostsResult)(nil)).Elem()
}

func (o ListSubAccountVMHostsResultOutput) ToListSubAccountVMHostsResultOutput() ListSubAccountVMHostsResultOutput {
	return o
}

func (o ListSubAccountVMHostsResultOutput) ToListSubAccountVMHostsResultOutputWithContext(ctx context.Context) ListSubAccountVMHostsResultOutput {
	return o
}

func (o ListSubAccountVMHostsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListSubAccountVMHostsResult] {
	return pulumix.Output[ListSubAccountVMHostsResult]{
		OutputState: o.OutputState,
	}
}

// Link to the next set of results, if any.
func (o ListSubAccountVMHostsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSubAccountVMHostsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Response of a list vm host update operation.
func (o ListSubAccountVMHostsResultOutput) Value() VMResourcesResponseArrayOutput {
	return o.ApplyT(func(v ListSubAccountVMHostsResult) []VMResourcesResponse { return v.Value }).(VMResourcesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSubAccountVMHostsResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List of elastic traffic filters in the account
func ListAllTrafficFilter(ctx *pulumi.Context, args *ListAllTrafficFilterArgs, opts ...pulumi.InvokeOption) (*ListAllTrafficFilterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAllTrafficFilterResult
	err := ctx.Invoke("azure-native:elastic/v20230601:listAllTrafficFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAllTrafficFilterArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of elastic traffic filters in the account
type ListAllTrafficFilterResult struct {
	// List of elastic traffic filters in the account
	Rulesets []ElasticTrafficFilterResponse `pulumi:"rulesets"`
}

func ListAllTrafficFilterOutput(ctx *pulumi.Context, args ListAllTrafficFilterOutputArgs, opts ...pulumi.InvokeOption) ListAllTrafficFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAllTrafficFilterResult, error) {
			args := v.(ListAllTrafficFilterArgs)
			r, err := ListAllTrafficFilter(ctx, &args, opts...)
			var s ListAllTrafficFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAllTrafficFilterResultOutput)
}

type ListAllTrafficFilterOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAllTrafficFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAllTrafficFilterArgs)(nil)).Elem()
}

// List of elastic traffic filters in the account
type ListAllTrafficFilterResultOutput struct{ *pulumi.OutputState }

func (ListAllTrafficFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAllTrafficFilterResult)(nil)).Elem()
}

func (o ListAllTrafficFilterResultOutput) ToListAllTrafficFilterResultOutput() ListAllTrafficFilterResultOutput {
	return o
}

func (o ListAllTrafficFilterResultOutput) ToListAllTrafficFilterResultOutputWithContext(ctx context.Context) ListAllTrafficFilterResultOutput {
	return o
}

func (o ListAllTrafficFilterResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListAllTrafficFilterResult] {
	return pulumix.Output[ListAllTrafficFilterResult]{
		OutputState: o.OutputState,
	}
}

// List of elastic traffic filters in the account
func (o ListAllTrafficFilterResultOutput) Rulesets() ElasticTrafficFilterResponseArrayOutput {
	return o.ApplyT(func(v ListAllTrafficFilterResult) []ElasticTrafficFilterResponse { return v.Rulesets }).(ElasticTrafficFilterResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAllTrafficFilterResultOutput{})
}

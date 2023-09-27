// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2022-04-01-preview.
func GetBillingHubServiceFreeHourBalance(ctx *pulumi.Context, args *GetBillingHubServiceFreeHourBalanceArgs, opts ...pulumi.InvokeOption) (*GetBillingHubServiceFreeHourBalanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetBillingHubServiceFreeHourBalanceResult
	err := ctx.Invoke("azure-native:testbase:getBillingHubServiceFreeHourBalance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBillingHubServiceFreeHourBalanceArgs struct {
	// The name of the resource group that contains the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

type GetBillingHubServiceFreeHourBalanceResult struct {
	IncrementEntries        []BillingHubFreeHourIncrementEntryResponse `pulumi:"incrementEntries"`
	TotalRemainingFreeHours *float64                                   `pulumi:"totalRemainingFreeHours"`
}

func GetBillingHubServiceFreeHourBalanceOutput(ctx *pulumi.Context, args GetBillingHubServiceFreeHourBalanceOutputArgs, opts ...pulumi.InvokeOption) GetBillingHubServiceFreeHourBalanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBillingHubServiceFreeHourBalanceResult, error) {
			args := v.(GetBillingHubServiceFreeHourBalanceArgs)
			r, err := GetBillingHubServiceFreeHourBalance(ctx, &args, opts...)
			var s GetBillingHubServiceFreeHourBalanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBillingHubServiceFreeHourBalanceResultOutput)
}

type GetBillingHubServiceFreeHourBalanceOutputArgs struct {
	// The name of the resource group that contains the resource.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (GetBillingHubServiceFreeHourBalanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingHubServiceFreeHourBalanceArgs)(nil)).Elem()
}

type GetBillingHubServiceFreeHourBalanceResultOutput struct{ *pulumi.OutputState }

func (GetBillingHubServiceFreeHourBalanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingHubServiceFreeHourBalanceResult)(nil)).Elem()
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) ToGetBillingHubServiceFreeHourBalanceResultOutput() GetBillingHubServiceFreeHourBalanceResultOutput {
	return o
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) ToGetBillingHubServiceFreeHourBalanceResultOutputWithContext(ctx context.Context) GetBillingHubServiceFreeHourBalanceResultOutput {
	return o
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetBillingHubServiceFreeHourBalanceResult] {
	return pulumix.Output[GetBillingHubServiceFreeHourBalanceResult]{
		OutputState: o.OutputState,
	}
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) IncrementEntries() BillingHubFreeHourIncrementEntryResponseArrayOutput {
	return o.ApplyT(func(v GetBillingHubServiceFreeHourBalanceResult) []BillingHubFreeHourIncrementEntryResponse {
		return v.IncrementEntries
	}).(BillingHubFreeHourIncrementEntryResponseArrayOutput)
}

func (o GetBillingHubServiceFreeHourBalanceResultOutput) TotalRemainingFreeHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetBillingHubServiceFreeHourBalanceResult) *float64 { return v.TotalRemainingFreeHours }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBillingHubServiceFreeHourBalanceResultOutput{})
}

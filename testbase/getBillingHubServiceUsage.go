// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure REST API version: 2022-04-01-preview.
func GetBillingHubServiceUsage(ctx *pulumi.Context, args *GetBillingHubServiceUsageArgs, opts ...pulumi.InvokeOption) (*GetBillingHubServiceUsageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetBillingHubServiceUsageResult
	err := ctx.Invoke("azure-native:testbase:getBillingHubServiceUsage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBillingHubServiceUsageArgs struct {
	EndTimeStamp string `pulumi:"endTimeStamp"`
	PageIndex    *int   `pulumi:"pageIndex"`
	PageSize     *int   `pulumi:"pageSize"`
	// The name of the resource group that contains the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StartTimeStamp    string `pulumi:"startTimeStamp"`
	// The resource name of the Test Base Account.
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

type GetBillingHubServiceUsageResult struct {
	NextRequest            *BillingHubGetUsageRequestResponse `pulumi:"nextRequest"`
	PackageUsageEntries    []BillingHubPackageUsageResponse   `pulumi:"packageUsageEntries"`
	TotalCharges           *float64                           `pulumi:"totalCharges"`
	TotalUsedBillableHours *float64                           `pulumi:"totalUsedBillableHours"`
	TotalUsedFreeHours     *float64                           `pulumi:"totalUsedFreeHours"`
}

func GetBillingHubServiceUsageOutput(ctx *pulumi.Context, args GetBillingHubServiceUsageOutputArgs, opts ...pulumi.InvokeOption) GetBillingHubServiceUsageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBillingHubServiceUsageResult, error) {
			args := v.(GetBillingHubServiceUsageArgs)
			r, err := GetBillingHubServiceUsage(ctx, &args, opts...)
			var s GetBillingHubServiceUsageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBillingHubServiceUsageResultOutput)
}

type GetBillingHubServiceUsageOutputArgs struct {
	EndTimeStamp pulumi.StringInput `pulumi:"endTimeStamp"`
	PageIndex    pulumi.IntPtrInput `pulumi:"pageIndex"`
	PageSize     pulumi.IntPtrInput `pulumi:"pageSize"`
	// The name of the resource group that contains the resource.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StartTimeStamp    pulumi.StringInput `pulumi:"startTimeStamp"`
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (GetBillingHubServiceUsageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingHubServiceUsageArgs)(nil)).Elem()
}

type GetBillingHubServiceUsageResultOutput struct{ *pulumi.OutputState }

func (GetBillingHubServiceUsageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingHubServiceUsageResult)(nil)).Elem()
}

func (o GetBillingHubServiceUsageResultOutput) ToGetBillingHubServiceUsageResultOutput() GetBillingHubServiceUsageResultOutput {
	return o
}

func (o GetBillingHubServiceUsageResultOutput) ToGetBillingHubServiceUsageResultOutputWithContext(ctx context.Context) GetBillingHubServiceUsageResultOutput {
	return o
}

func (o GetBillingHubServiceUsageResultOutput) NextRequest() BillingHubGetUsageRequestResponsePtrOutput {
	return o.ApplyT(func(v GetBillingHubServiceUsageResult) *BillingHubGetUsageRequestResponse { return v.NextRequest }).(BillingHubGetUsageRequestResponsePtrOutput)
}

func (o GetBillingHubServiceUsageResultOutput) PackageUsageEntries() BillingHubPackageUsageResponseArrayOutput {
	return o.ApplyT(func(v GetBillingHubServiceUsageResult) []BillingHubPackageUsageResponse { return v.PackageUsageEntries }).(BillingHubPackageUsageResponseArrayOutput)
}

func (o GetBillingHubServiceUsageResultOutput) TotalCharges() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetBillingHubServiceUsageResult) *float64 { return v.TotalCharges }).(pulumi.Float64PtrOutput)
}

func (o GetBillingHubServiceUsageResultOutput) TotalUsedBillableHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetBillingHubServiceUsageResult) *float64 { return v.TotalUsedBillableHours }).(pulumi.Float64PtrOutput)
}

func (o GetBillingHubServiceUsageResultOutput) TotalUsedFreeHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetBillingHubServiceUsageResult) *float64 { return v.TotalUsedFreeHours }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBillingHubServiceUsageResultOutput{})
}

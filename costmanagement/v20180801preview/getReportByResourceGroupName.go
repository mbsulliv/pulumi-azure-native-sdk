// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the report for a resource group under a subscription by report name.
func LookupReportByResourceGroupName(ctx *pulumi.Context, args *LookupReportByResourceGroupNameArgs, opts ...pulumi.InvokeOption) (*LookupReportByResourceGroupNameResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReportByResourceGroupNameResult
	err := ctx.Invoke("azure-native:costmanagement/v20180801preview:getReportByResourceGroupName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportByResourceGroupNameArgs struct {
	// Report Name.
	ReportName string `pulumi:"reportName"`
	// Azure Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A report resource.
type LookupReportByResourceGroupNameResult struct {
	// Has definition for the report.
	Definition ReportDefinitionResponse `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Has schedule information for the report.
	Schedule *ReportScheduleResponse `pulumi:"schedule"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupReportByResourceGroupNameOutput(ctx *pulumi.Context, args LookupReportByResourceGroupNameOutputArgs, opts ...pulumi.InvokeOption) LookupReportByResourceGroupNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportByResourceGroupNameResult, error) {
			args := v.(LookupReportByResourceGroupNameArgs)
			r, err := LookupReportByResourceGroupName(ctx, &args, opts...)
			var s LookupReportByResourceGroupNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportByResourceGroupNameResultOutput)
}

type LookupReportByResourceGroupNameOutputArgs struct {
	// Report Name.
	ReportName pulumi.StringInput `pulumi:"reportName"`
	// Azure Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReportByResourceGroupNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportByResourceGroupNameArgs)(nil)).Elem()
}

// A report resource.
type LookupReportByResourceGroupNameResultOutput struct{ *pulumi.OutputState }

func (LookupReportByResourceGroupNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportByResourceGroupNameResult)(nil)).Elem()
}

func (o LookupReportByResourceGroupNameResultOutput) ToLookupReportByResourceGroupNameResultOutput() LookupReportByResourceGroupNameResultOutput {
	return o
}

func (o LookupReportByResourceGroupNameResultOutput) ToLookupReportByResourceGroupNameResultOutputWithContext(ctx context.Context) LookupReportByResourceGroupNameResultOutput {
	return o
}

// Has definition for the report.
func (o LookupReportByResourceGroupNameResultOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) ReportDefinitionResponse { return v.Definition }).(ReportDefinitionResponseOutput)
}

// Has delivery information for the report.
func (o LookupReportByResourceGroupNameResultOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) ReportDeliveryInfoResponse { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

// The format of the report being delivered.
func (o LookupReportByResourceGroupNameResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupReportByResourceGroupNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupReportByResourceGroupNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) string { return v.Name }).(pulumi.StringOutput)
}

// Has schedule information for the report.
func (o LookupReportByResourceGroupNameResultOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) *ReportScheduleResponse { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

// Resource tags.
func (o LookupReportByResourceGroupNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupReportByResourceGroupNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportByResourceGroupNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportByResourceGroupNameResultOutput{})
}

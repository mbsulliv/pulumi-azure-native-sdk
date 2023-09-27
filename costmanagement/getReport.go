// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the report for a subscription by report name.
// Azure REST API version: 2018-08-01-preview.
func LookupReport(ctx *pulumi.Context, args *LookupReportArgs, opts ...pulumi.InvokeOption) (*LookupReportResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReportResult
	err := ctx.Invoke("azure-native:costmanagement:getReport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportArgs struct {
	// Report Name.
	ReportName string `pulumi:"reportName"`
}

// A report resource.
type LookupReportResult struct {
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

func LookupReportOutput(ctx *pulumi.Context, args LookupReportOutputArgs, opts ...pulumi.InvokeOption) LookupReportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportResult, error) {
			args := v.(LookupReportArgs)
			r, err := LookupReport(ctx, &args, opts...)
			var s LookupReportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportResultOutput)
}

type LookupReportOutputArgs struct {
	// Report Name.
	ReportName pulumi.StringInput `pulumi:"reportName"`
}

func (LookupReportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportArgs)(nil)).Elem()
}

// A report resource.
type LookupReportResultOutput struct{ *pulumi.OutputState }

func (LookupReportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportResult)(nil)).Elem()
}

func (o LookupReportResultOutput) ToLookupReportResultOutput() LookupReportResultOutput {
	return o
}

func (o LookupReportResultOutput) ToLookupReportResultOutputWithContext(ctx context.Context) LookupReportResultOutput {
	return o
}

func (o LookupReportResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReportResult] {
	return pulumix.Output[LookupReportResult]{
		OutputState: o.OutputState,
	}
}

// Has definition for the report.
func (o LookupReportResultOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportResult) ReportDefinitionResponse { return v.Definition }).(ReportDefinitionResponseOutput)
}

// Has delivery information for the report.
func (o LookupReportResultOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportResult) ReportDeliveryInfoResponse { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

// The format of the report being delivered.
func (o LookupReportResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupReportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupReportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Name }).(pulumi.StringOutput)
}

// Has schedule information for the report.
func (o LookupReportResultOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportResult) *ReportScheduleResponse { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

// Resource tags.
func (o LookupReportResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupReportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportResultOutput{})
}

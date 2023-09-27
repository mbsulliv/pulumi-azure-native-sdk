// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221116preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the AppComplianceAutomation report and its properties.
func LookupReport(ctx *pulumi.Context, args *LookupReportArgs, opts ...pulumi.InvokeOption) (*LookupReportResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReportResult
	err := ctx.Invoke("azure-native:appcomplianceautomation/v20221116preview:getReport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportArgs struct {
	// Report Name.
	ReportName string `pulumi:"reportName"`
}

// A class represent an AppComplianceAutomation report resource.
type LookupReportResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Report property.
	Properties ReportPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
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

// A class represent an AppComplianceAutomation report resource.
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

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupReportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupReportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Name }).(pulumi.StringOutput)
}

// Report property.
func (o LookupReportResultOutput) Properties() ReportPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReportResult) ReportPropertiesResponse { return v.Properties }).(ReportPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupReportResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupReportResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupReportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportResultOutput{})
}

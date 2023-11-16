// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation to get the export for the defined scope by export name.
// Azure REST API version: 2023-03-01.
//
// Other available API versions: 2019-10-01, 2023-04-01-preview, 2023-08-01.
func LookupExport(ctx *pulumi.Context, args *LookupExportArgs, opts ...pulumi.InvokeOption) (*LookupExportResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExportResult
	err := ctx.Invoke("azure-native:costmanagement:getExport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportArgs struct {
	// May be used to expand the properties within an export. Currently only 'runHistory' is supported and will return information for the last 10 runs of the export.
	Expand *string `pulumi:"expand"`
	// Export Name.
	ExportName string `pulumi:"exportName"`
	// The scope associated with export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
	Scope string `pulumi:"scope"`
}

// An export resource.
type LookupExportResult struct {
	// Has the definition for the export.
	Definition ExportDefinitionResponse `pulumi:"definition"`
	// Has delivery information for the export.
	DeliveryInfo ExportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// The format of the export being delivered. Currently only 'Csv' is supported.
	Format *string `pulumi:"format"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// If the export has an active schedule, provides an estimate of the next run time.
	NextRunTimeEstimate string `pulumi:"nextRunTimeEstimate"`
	// If set to true, exported data will be partitioned by size and placed in a blob directory together with a manifest file. Note: this option is currently available only for Microsoft Customer Agreement commerce scopes.
	PartitionData *bool `pulumi:"partitionData"`
	// If requested, has the most recent run history for the export.
	RunHistory *ExportExecutionListResultResponse `pulumi:"runHistory"`
	// Has schedule information for the export.
	Schedule *ExportScheduleResponse `pulumi:"schedule"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupExportOutput(ctx *pulumi.Context, args LookupExportOutputArgs, opts ...pulumi.InvokeOption) LookupExportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExportResult, error) {
			args := v.(LookupExportArgs)
			r, err := LookupExport(ctx, &args, opts...)
			var s LookupExportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExportResultOutput)
}

type LookupExportOutputArgs struct {
	// May be used to expand the properties within an export. Currently only 'runHistory' is supported and will return information for the last 10 runs of the export.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// Export Name.
	ExportName pulumi.StringInput `pulumi:"exportName"`
	// The scope associated with export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupExportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportArgs)(nil)).Elem()
}

// An export resource.
type LookupExportResultOutput struct{ *pulumi.OutputState }

func (LookupExportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportResult)(nil)).Elem()
}

func (o LookupExportResultOutput) ToLookupExportResultOutput() LookupExportResultOutput {
	return o
}

func (o LookupExportResultOutput) ToLookupExportResultOutputWithContext(ctx context.Context) LookupExportResultOutput {
	return o
}

// Has the definition for the export.
func (o LookupExportResultOutput) Definition() ExportDefinitionResponseOutput {
	return o.ApplyT(func(v LookupExportResult) ExportDefinitionResponse { return v.Definition }).(ExportDefinitionResponseOutput)
}

// Has delivery information for the export.
func (o LookupExportResultOutput) DeliveryInfo() ExportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupExportResult) ExportDeliveryInfoResponse { return v.DeliveryInfo }).(ExportDeliveryInfoResponseOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o LookupExportResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// The format of the export being delivered. Currently only 'Csv' is supported.
func (o LookupExportResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupExportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupExportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Name }).(pulumi.StringOutput)
}

// If the export has an active schedule, provides an estimate of the next run time.
func (o LookupExportResultOutput) NextRunTimeEstimate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.NextRunTimeEstimate }).(pulumi.StringOutput)
}

// If set to true, exported data will be partitioned by size and placed in a blob directory together with a manifest file. Note: this option is currently available only for Microsoft Customer Agreement commerce scopes.
func (o LookupExportResultOutput) PartitionData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExportResult) *bool { return v.PartitionData }).(pulumi.BoolPtrOutput)
}

// If requested, has the most recent run history for the export.
func (o LookupExportResultOutput) RunHistory() ExportExecutionListResultResponsePtrOutput {
	return o.ApplyT(func(v LookupExportResult) *ExportExecutionListResultResponse { return v.RunHistory }).(ExportExecutionListResultResponsePtrOutput)
}

// Has schedule information for the export.
func (o LookupExportResultOutput) Schedule() ExportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupExportResult) *ExportScheduleResponse { return v.Schedule }).(ExportScheduleResponsePtrOutput)
}

// Resource type.
func (o LookupExportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExportResultOutput{})
}

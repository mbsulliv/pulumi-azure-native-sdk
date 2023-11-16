// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the view by view name.
// Azure REST API version: 2023-03-01.
//
// Other available API versions: 2019-11-01, 2020-06-01, 2022-10-01, 2022-10-05-preview, 2023-04-01-preview, 2023-08-01.
func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupViewResult
	err := ctx.Invoke("azure-native:costmanagement:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	// View name
	ViewName string `pulumi:"viewName"`
}

// States and configurations of Cost Analysis.
type LookupViewResult struct {
	// Show costs accumulated over time.
	Accumulated *string `pulumi:"accumulated"`
	// Chart type of the main view in Cost Analysis. Required.
	Chart *string `pulumi:"chart"`
	// Date the user created this view.
	CreatedOn string `pulumi:"createdOn"`
	// Currency of the current view.
	Currency string `pulumi:"currency"`
	// Has definition for data in this report config.
	DataSet *ReportConfigDatasetResponse `pulumi:"dataSet"`
	// Date range of the current view.
	DateRange *string `pulumi:"dateRange"`
	// User input name of the view. Required.
	DisplayName *string `pulumi:"displayName"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// Resource Id.
	Id string `pulumi:"id"`
	// If true, report includes monetary commitment.
	IncludeMonetaryCommitment *bool `pulumi:"includeMonetaryCommitment"`
	// List of KPIs to show in Cost Analysis UI.
	Kpis []KpiPropertiesResponse `pulumi:"kpis"`
	// Metric to use when displaying costs.
	Metric *string `pulumi:"metric"`
	// Date when the user last modified this view.
	ModifiedOn *string `pulumi:"modifiedOn"`
	// Resource name.
	Name string `pulumi:"name"`
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots []PivotPropertiesResponse `pulumi:"pivots"`
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope *string `pulumi:"scope"`
	// Has time period for pulling data for the report.
	TimePeriod *ReportConfigTimePeriodResponse `pulumi:"timePeriod"`
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe string `pulumi:"timeframe"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupViewOutput(ctx *pulumi.Context, args LookupViewOutputArgs, opts ...pulumi.InvokeOption) LookupViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupViewResult, error) {
			args := v.(LookupViewArgs)
			r, err := LookupView(ctx, &args, opts...)
			var s LookupViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupViewResultOutput)
}

type LookupViewOutputArgs struct {
	// View name
	ViewName pulumi.StringInput `pulumi:"viewName"`
}

func (LookupViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewArgs)(nil)).Elem()
}

// States and configurations of Cost Analysis.
type LookupViewResultOutput struct{ *pulumi.OutputState }

func (LookupViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewResult)(nil)).Elem()
}

func (o LookupViewResultOutput) ToLookupViewResultOutput() LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) ToLookupViewResultOutputWithContext(ctx context.Context) LookupViewResultOutput {
	return o
}

// Show costs accumulated over time.
func (o LookupViewResultOutput) Accumulated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Accumulated }).(pulumi.StringPtrOutput)
}

// Chart type of the main view in Cost Analysis. Required.
func (o LookupViewResultOutput) Chart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Chart }).(pulumi.StringPtrOutput)
}

// Date the user created this view.
func (o LookupViewResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// Currency of the current view.
func (o LookupViewResultOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Currency }).(pulumi.StringOutput)
}

// Has definition for data in this report config.
func (o LookupViewResultOutput) DataSet() ReportConfigDatasetResponsePtrOutput {
	return o.ApplyT(func(v LookupViewResult) *ReportConfigDatasetResponse { return v.DataSet }).(ReportConfigDatasetResponsePtrOutput)
}

// Date range of the current view.
func (o LookupViewResultOutput) DateRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.DateRange }).(pulumi.StringPtrOutput)
}

// User input name of the view. Required.
func (o LookupViewResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o LookupViewResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupViewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Id }).(pulumi.StringOutput)
}

// If true, report includes monetary commitment.
func (o LookupViewResultOutput) IncludeMonetaryCommitment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *bool { return v.IncludeMonetaryCommitment }).(pulumi.BoolPtrOutput)
}

// List of KPIs to show in Cost Analysis UI.
func (o LookupViewResultOutput) Kpis() KpiPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupViewResult) []KpiPropertiesResponse { return v.Kpis }).(KpiPropertiesResponseArrayOutput)
}

// Metric to use when displaying costs.
func (o LookupViewResultOutput) Metric() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Metric }).(pulumi.StringPtrOutput)
}

// Date when the user last modified this view.
func (o LookupViewResultOutput) ModifiedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.ModifiedOn }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupViewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Name }).(pulumi.StringOutput)
}

// Configuration of 3 sub-views in the Cost Analysis UI.
func (o LookupViewResultOutput) Pivots() PivotPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupViewResult) []PivotPropertiesResponse { return v.Pivots }).(PivotPropertiesResponseArrayOutput)
}

// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
func (o LookupViewResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// Has time period for pulling data for the report.
func (o LookupViewResultOutput) TimePeriod() ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v LookupViewResult) *ReportConfigTimePeriodResponse { return v.TimePeriod }).(ReportConfigTimePeriodResponsePtrOutput)
}

// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
func (o LookupViewResultOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Timeframe }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupViewResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupViewResultOutput{})
}

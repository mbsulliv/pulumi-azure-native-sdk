// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the view for the defined scope by view name.
func LookupViewByScope(ctx *pulumi.Context, args *LookupViewByScopeArgs, opts ...pulumi.InvokeOption) (*LookupViewByScopeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupViewByScopeResult
	err := ctx.Invoke("azure-native:costmanagement/v20230801:getViewByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewByScopeArgs struct {
	// The scope associated with view operations. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External Billing Account scope and 'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope.
	Scope string `pulumi:"scope"`
	// View name
	ViewName string `pulumi:"viewName"`
}

// States and configurations of Cost Analysis.
type LookupViewByScopeResult struct {
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

func LookupViewByScopeOutput(ctx *pulumi.Context, args LookupViewByScopeOutputArgs, opts ...pulumi.InvokeOption) LookupViewByScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupViewByScopeResult, error) {
			args := v.(LookupViewByScopeArgs)
			r, err := LookupViewByScope(ctx, &args, opts...)
			var s LookupViewByScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupViewByScopeResultOutput)
}

type LookupViewByScopeOutputArgs struct {
	// The scope associated with view operations. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External Billing Account scope and 'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope.
	Scope pulumi.StringInput `pulumi:"scope"`
	// View name
	ViewName pulumi.StringInput `pulumi:"viewName"`
}

func (LookupViewByScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewByScopeArgs)(nil)).Elem()
}

// States and configurations of Cost Analysis.
type LookupViewByScopeResultOutput struct{ *pulumi.OutputState }

func (LookupViewByScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewByScopeResult)(nil)).Elem()
}

func (o LookupViewByScopeResultOutput) ToLookupViewByScopeResultOutput() LookupViewByScopeResultOutput {
	return o
}

func (o LookupViewByScopeResultOutput) ToLookupViewByScopeResultOutputWithContext(ctx context.Context) LookupViewByScopeResultOutput {
	return o
}

func (o LookupViewByScopeResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupViewByScopeResult] {
	return pulumix.Output[LookupViewByScopeResult]{
		OutputState: o.OutputState,
	}
}

// Show costs accumulated over time.
func (o LookupViewByScopeResultOutput) Accumulated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.Accumulated }).(pulumi.StringPtrOutput)
}

// Chart type of the main view in Cost Analysis. Required.
func (o LookupViewByScopeResultOutput) Chart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.Chart }).(pulumi.StringPtrOutput)
}

// Date the user created this view.
func (o LookupViewByScopeResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// Currency of the current view.
func (o LookupViewByScopeResultOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Currency }).(pulumi.StringOutput)
}

// Has definition for data in this report config.
func (o LookupViewByScopeResultOutput) DataSet() ReportConfigDatasetResponsePtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *ReportConfigDatasetResponse { return v.DataSet }).(ReportConfigDatasetResponsePtrOutput)
}

// Date range of the current view.
func (o LookupViewByScopeResultOutput) DateRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.DateRange }).(pulumi.StringPtrOutput)
}

// User input name of the view. Required.
func (o LookupViewByScopeResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o LookupViewByScopeResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupViewByScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

// If true, report includes monetary commitment.
func (o LookupViewByScopeResultOutput) IncludeMonetaryCommitment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *bool { return v.IncludeMonetaryCommitment }).(pulumi.BoolPtrOutput)
}

// List of KPIs to show in Cost Analysis UI.
func (o LookupViewByScopeResultOutput) Kpis() KpiPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) []KpiPropertiesResponse { return v.Kpis }).(KpiPropertiesResponseArrayOutput)
}

// Metric to use when displaying costs.
func (o LookupViewByScopeResultOutput) Metric() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.Metric }).(pulumi.StringPtrOutput)
}

// Date when the user last modified this view.
func (o LookupViewByScopeResultOutput) ModifiedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.ModifiedOn }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupViewByScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Configuration of 3 sub-views in the Cost Analysis UI.
func (o LookupViewByScopeResultOutput) Pivots() PivotPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) []PivotPropertiesResponse { return v.Pivots }).(PivotPropertiesResponseArrayOutput)
}

// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
func (o LookupViewByScopeResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// Has time period for pulling data for the report.
func (o LookupViewByScopeResultOutput) TimePeriod() ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *ReportConfigTimePeriodResponse { return v.TimePeriod }).(ReportConfigTimePeriodResponsePtrOutput)
}

// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
func (o LookupViewByScopeResultOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Timeframe }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupViewByScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupViewByScopeResultOutput{})
}

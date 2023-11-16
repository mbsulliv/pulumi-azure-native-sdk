// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// States and configurations of Cost Analysis.
type ViewByScope struct {
	pulumi.CustomResourceState

	// Show costs accumulated over time.
	Accumulated pulumi.StringPtrOutput `pulumi:"accumulated"`
	// Chart type of the main view in Cost Analysis. Required.
	Chart pulumi.StringPtrOutput `pulumi:"chart"`
	// Date the user created this view.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Currency of the current view.
	Currency pulumi.StringOutput `pulumi:"currency"`
	// Has definition for data in this report config.
	DataSet ReportConfigDatasetResponsePtrOutput `pulumi:"dataSet"`
	// Date range of the current view.
	DateRange pulumi.StringPtrOutput `pulumi:"dateRange"`
	// User input name of the view. Required.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// If true, report includes monetary commitment.
	IncludeMonetaryCommitment pulumi.BoolPtrOutput `pulumi:"includeMonetaryCommitment"`
	// List of KPIs to show in Cost Analysis UI.
	Kpis KpiPropertiesResponseArrayOutput `pulumi:"kpis"`
	// Metric to use when displaying costs.
	Metric pulumi.StringPtrOutput `pulumi:"metric"`
	// Date when the user last modified this view.
	ModifiedOn pulumi.StringPtrOutput `pulumi:"modifiedOn"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots PivotPropertiesResponseArrayOutput `pulumi:"pivots"`
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Has time period for pulling data for the report.
	TimePeriod ReportConfigTimePeriodResponsePtrOutput `pulumi:"timePeriod"`
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe pulumi.StringOutput `pulumi:"timeframe"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewViewByScope registers a new resource with the given unique name, arguments, and options.
func NewViewByScope(ctx *pulumi.Context,
	name string, args *ViewByScopeArgs, opts ...pulumi.ResourceOption) (*ViewByScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Timeframe == nil {
		return nil, errors.New("invalid value for required argument 'Timeframe'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190401preview:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20191101:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20200601:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20211001:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20220801preview:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001preview:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221005preview:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230301:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230801:ViewByScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ViewByScope
	err := ctx.RegisterResource("azure-native:costmanagement/v20230401preview:ViewByScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetViewByScope gets an existing ViewByScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetViewByScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewByScopeState, opts ...pulumi.ResourceOption) (*ViewByScope, error) {
	var resource ViewByScope
	err := ctx.ReadResource("azure-native:costmanagement/v20230401preview:ViewByScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ViewByScope resources.
type viewByScopeState struct {
}

type ViewByScopeState struct {
}

func (ViewByScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewByScopeState)(nil)).Elem()
}

type viewByScopeArgs struct {
	// Show costs accumulated over time.
	Accumulated *string `pulumi:"accumulated"`
	// Chart type of the main view in Cost Analysis. Required.
	Chart *string `pulumi:"chart"`
	// Has definition for data in this report config.
	DataSet *ReportConfigDataset `pulumi:"dataSet"`
	// Date range of the current view.
	DateRange *string `pulumi:"dateRange"`
	// User input name of the view. Required.
	DisplayName *string `pulumi:"displayName"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// If true, report includes monetary commitment.
	IncludeMonetaryCommitment *bool `pulumi:"includeMonetaryCommitment"`
	// List of KPIs to show in Cost Analysis UI.
	Kpis []KpiProperties `pulumi:"kpis"`
	// Metric to use when displaying costs.
	Metric *string `pulumi:"metric"`
	// Date when the user last modified this view.
	ModifiedOn *string `pulumi:"modifiedOn"`
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots []PivotProperties `pulumi:"pivots"`
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope string `pulumi:"scope"`
	// Has time period for pulling data for the report.
	TimePeriod *ReportConfigTimePeriod `pulumi:"timePeriod"`
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe string `pulumi:"timeframe"`
	// The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
	Type string `pulumi:"type"`
	// View name
	ViewName *string `pulumi:"viewName"`
}

// The set of arguments for constructing a ViewByScope resource.
type ViewByScopeArgs struct {
	// Show costs accumulated over time.
	Accumulated pulumi.StringPtrInput
	// Chart type of the main view in Cost Analysis. Required.
	Chart pulumi.StringPtrInput
	// Has definition for data in this report config.
	DataSet ReportConfigDatasetPtrInput
	// Date range of the current view.
	DateRange pulumi.StringPtrInput
	// User input name of the view. Required.
	DisplayName pulumi.StringPtrInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// If true, report includes monetary commitment.
	IncludeMonetaryCommitment pulumi.BoolPtrInput
	// List of KPIs to show in Cost Analysis UI.
	Kpis KpiPropertiesArrayInput
	// Metric to use when displaying costs.
	Metric pulumi.StringPtrInput
	// Date when the user last modified this view.
	ModifiedOn pulumi.StringPtrInput
	// Configuration of 3 sub-views in the Cost Analysis UI.
	Pivots PivotPropertiesArrayInput
	// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringInput
	// Has time period for pulling data for the report.
	TimePeriod ReportConfigTimePeriodPtrInput
	// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
	Timeframe pulumi.StringInput
	// The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
	Type pulumi.StringInput
	// View name
	ViewName pulumi.StringPtrInput
}

func (ViewByScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewByScopeArgs)(nil)).Elem()
}

type ViewByScopeInput interface {
	pulumi.Input

	ToViewByScopeOutput() ViewByScopeOutput
	ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput
}

func (*ViewByScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ViewByScope)(nil)).Elem()
}

func (i *ViewByScope) ToViewByScopeOutput() ViewByScopeOutput {
	return i.ToViewByScopeOutputWithContext(context.Background())
}

func (i *ViewByScope) ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewByScopeOutput)
}

type ViewByScopeOutput struct{ *pulumi.OutputState }

func (ViewByScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ViewByScope)(nil)).Elem()
}

func (o ViewByScopeOutput) ToViewByScopeOutput() ViewByScopeOutput {
	return o
}

func (o ViewByScopeOutput) ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput {
	return o
}

// Show costs accumulated over time.
func (o ViewByScopeOutput) Accumulated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.Accumulated }).(pulumi.StringPtrOutput)
}

// Chart type of the main view in Cost Analysis. Required.
func (o ViewByScopeOutput) Chart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.Chart }).(pulumi.StringPtrOutput)
}

// Date the user created this view.
func (o ViewByScopeOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

// Currency of the current view.
func (o ViewByScopeOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.Currency }).(pulumi.StringOutput)
}

// Has definition for data in this report config.
func (o ViewByScopeOutput) DataSet() ReportConfigDatasetResponsePtrOutput {
	return o.ApplyT(func(v *ViewByScope) ReportConfigDatasetResponsePtrOutput { return v.DataSet }).(ReportConfigDatasetResponsePtrOutput)
}

// Date range of the current view.
func (o ViewByScopeOutput) DateRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.DateRange }).(pulumi.StringPtrOutput)
}

// User input name of the view. Required.
func (o ViewByScopeOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o ViewByScopeOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// If true, report includes monetary commitment.
func (o ViewByScopeOutput) IncludeMonetaryCommitment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.BoolPtrOutput { return v.IncludeMonetaryCommitment }).(pulumi.BoolPtrOutput)
}

// List of KPIs to show in Cost Analysis UI.
func (o ViewByScopeOutput) Kpis() KpiPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *ViewByScope) KpiPropertiesResponseArrayOutput { return v.Kpis }).(KpiPropertiesResponseArrayOutput)
}

// Metric to use when displaying costs.
func (o ViewByScopeOutput) Metric() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.Metric }).(pulumi.StringPtrOutput)
}

// Date when the user last modified this view.
func (o ViewByScopeOutput) ModifiedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.ModifiedOn }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ViewByScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configuration of 3 sub-views in the Cost Analysis UI.
func (o ViewByScopeOutput) Pivots() PivotPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *ViewByScope) PivotPropertiesResponseArrayOutput { return v.Pivots }).(PivotPropertiesResponseArrayOutput)
}

// Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
func (o ViewByScopeOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Has time period for pulling data for the report.
func (o ViewByScopeOutput) TimePeriod() ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ViewByScope) ReportConfigTimePeriodResponsePtrOutput { return v.TimePeriod }).(ReportConfigTimePeriodResponsePtrOutput)
}

// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
func (o ViewByScopeOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.Timeframe }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ViewByScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ViewByScopeOutput{})
}

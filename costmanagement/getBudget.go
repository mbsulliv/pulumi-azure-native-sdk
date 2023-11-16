// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the budget for the scope by budget name.
// Azure REST API version: 2023-04-01-preview.
//
// Other available API versions: 2019-04-01-preview, 2023-08-01.
func LookupBudget(ctx *pulumi.Context, args *LookupBudgetArgs, opts ...pulumi.InvokeOption) (*LookupBudgetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBudgetResult
	err := ctx.Invoke("azure-native:costmanagement:getBudget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBudgetArgs struct {
	// Budget Name.
	BudgetName string `pulumi:"budgetName"`
	// The scope associated with budget operations.
	//
	//  Supported scopes for **CategoryType: Cost**
	//
	//  Azure RBAC Scopes:
	// - '/subscriptions/{subscriptionId}/' for subscription scope
	// - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope
	// - '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope
	//
	//  EA (Enterprise Agreement) Scopes:
	//
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope
	//
	//  MCA (Modern Customer Agreement) Scopes:
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
	//
	//  Supported scopes for **CategoryType: ReservationUtilization**
	//
	//  EA (Enterprise Agreement) Scopes:
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account Scope
	//
	// MCA (Modern Customer Agreement) Scopes:
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope (non-CSP only)
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
	Scope string `pulumi:"scope"`
}

// A budget resource.
type LookupBudgetResult struct {
	// The total amount of cost to track with the budget.
	//
	//  Supported for CategoryType(s): Cost.
	//
	//  Required for CategoryType(s): Cost.
	Amount *float64 `pulumi:"amount"`
	// The category of the budget.
	// - 'Cost' defines a Budget.
	// - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
	Category string `pulumi:"category"`
	// The current amount of cost which is being tracked for a budget.
	//
	//  Supported for CategoryType(s): Cost.
	CurrentSpend CurrentSpendResponse `pulumi:"currentSpend"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// May be used to filter budgets by user-specified dimensions and/or tags.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	Filter *BudgetFilterResponse `pulumi:"filter"`
	// The forecasted cost which is being tracked for a budget.
	//
	//  Supported for CategoryType(s): Cost.
	ForecastSpend ForecastSpendResponse `pulumi:"forecastSpend"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Dictionary of notifications associated with the budget.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	//
	// - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
	// - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
	Notifications map[string]NotificationResponse `pulumi:"notifications"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
	//
	// Supported for CategoryType(s): Cost, ReservationUtilization.
	//
	//  Supported timeGrainTypes for **CategoryType: Cost**
	//
	// - Monthly
	// - Quarterly
	// - Annually
	// - BillingMonth*
	// - BillingQuarter*
	// - BillingAnnual*
	//
	//   *only supported for Web Direct customers.
	//
	//  Supported timeGrainTypes for **CategoryType: ReservationUtilization**
	// - Last7Days
	// - Last30Days
	//
	//  Required for CategoryType(s): Cost, ReservationUtilization.
	TimeGrain string `pulumi:"timeGrain"`
	// The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	//
	//  Required for CategoryType(s): Cost, ReservationUtilization.
	TimePeriod BudgetTimePeriodResponse `pulumi:"timePeriod"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBudgetOutput(ctx *pulumi.Context, args LookupBudgetOutputArgs, opts ...pulumi.InvokeOption) LookupBudgetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBudgetResult, error) {
			args := v.(LookupBudgetArgs)
			r, err := LookupBudget(ctx, &args, opts...)
			var s LookupBudgetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBudgetResultOutput)
}

type LookupBudgetOutputArgs struct {
	// Budget Name.
	BudgetName pulumi.StringInput `pulumi:"budgetName"`
	// The scope associated with budget operations.
	//
	//  Supported scopes for **CategoryType: Cost**
	//
	//  Azure RBAC Scopes:
	// - '/subscriptions/{subscriptionId}/' for subscription scope
	// - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope
	// - '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope
	//
	//  EA (Enterprise Agreement) Scopes:
	//
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope
	//
	//  MCA (Modern Customer Agreement) Scopes:
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
	//
	//  Supported scopes for **CategoryType: ReservationUtilization**
	//
	//  EA (Enterprise Agreement) Scopes:
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account Scope
	//
	// MCA (Modern Customer Agreement) Scopes:
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope (non-CSP only)
	// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupBudgetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetArgs)(nil)).Elem()
}

// A budget resource.
type LookupBudgetResultOutput struct{ *pulumi.OutputState }

func (LookupBudgetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBudgetResult)(nil)).Elem()
}

func (o LookupBudgetResultOutput) ToLookupBudgetResultOutput() LookupBudgetResultOutput {
	return o
}

func (o LookupBudgetResultOutput) ToLookupBudgetResultOutputWithContext(ctx context.Context) LookupBudgetResultOutput {
	return o
}

// The total amount of cost to track with the budget.
//
//	Supported for CategoryType(s): Cost.
//
//	Required for CategoryType(s): Cost.
func (o LookupBudgetResultOutput) Amount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupBudgetResult) *float64 { return v.Amount }).(pulumi.Float64PtrOutput)
}

// The category of the budget.
// - 'Cost' defines a Budget.
// - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
func (o LookupBudgetResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Category }).(pulumi.StringOutput)
}

// The current amount of cost which is being tracked for a budget.
//
//	Supported for CategoryType(s): Cost.
func (o LookupBudgetResultOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) CurrentSpendResponse { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o LookupBudgetResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBudgetResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// May be used to filter budgets by user-specified dimensions and/or tags.
//
//	Supported for CategoryType(s): Cost, ReservationUtilization.
func (o LookupBudgetResultOutput) Filter() BudgetFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupBudgetResult) *BudgetFilterResponse { return v.Filter }).(BudgetFilterResponsePtrOutput)
}

// The forecasted cost which is being tracked for a budget.
//
//	Supported for CategoryType(s): Cost.
func (o LookupBudgetResultOutput) ForecastSpend() ForecastSpendResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) ForecastSpendResponse { return v.ForecastSpend }).(ForecastSpendResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBudgetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBudgetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Dictionary of notifications associated with the budget.
//
//	Supported for CategoryType(s): Cost, ReservationUtilization.
//
// - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
// - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
func (o LookupBudgetResultOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v LookupBudgetResult) map[string]NotificationResponse { return v.Notifications }).(NotificationResponseMapOutput)
}

// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
//
// Supported for CategoryType(s): Cost, ReservationUtilization.
//
//	Supported timeGrainTypes for **CategoryType: Cost**
//
// - Monthly
// - Quarterly
// - Annually
// - BillingMonth*
// - BillingQuarter*
// - BillingAnnual*
//
//	 *only supported for Web Direct customers.
//
//	Supported timeGrainTypes for **CategoryType: ReservationUtilization**
//
// - Last7Days
// - Last30Days
//
//	Required for CategoryType(s): Cost, ReservationUtilization.
func (o LookupBudgetResultOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.TimeGrain }).(pulumi.StringOutput)
}

// The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
//
//	Supported for CategoryType(s): Cost, ReservationUtilization.
//
//	Required for CategoryType(s): Cost, ReservationUtilization.
func (o LookupBudgetResultOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) BudgetTimePeriodResponse { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBudgetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBudgetResultOutput{})
}

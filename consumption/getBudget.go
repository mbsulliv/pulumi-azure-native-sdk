// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package consumption

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the budget for the scope by budget name.
// Azure REST API version: 2023-05-01.
func LookupBudget(ctx *pulumi.Context, args *LookupBudgetArgs, opts ...pulumi.InvokeOption) (*LookupBudgetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBudgetResult
	err := ctx.Invoke("azure-native:consumption:getBudget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBudgetArgs struct {
	// Budget Name.
	BudgetName string `pulumi:"budgetName"`
	// The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
	Scope string `pulumi:"scope"`
}

// A budget resource.
type LookupBudgetResult struct {
	// The total amount of cost to track with the budget
	Amount float64 `pulumi:"amount"`
	// The category of the budget, whether the budget tracks cost or usage.
	Category string `pulumi:"category"`
	// The current amount of cost which is being tracked for a budget.
	CurrentSpend CurrentSpendResponse `pulumi:"currentSpend"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// May be used to filter budgets by user-specified dimensions and/or tags.
	Filter *BudgetFilterResponse `pulumi:"filter"`
	// The forecasted cost which is being tracked for a budget.
	ForecastSpend ForecastSpendResponse `pulumi:"forecastSpend"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications map[string]NotificationResponse `pulumi:"notifications"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
	TimeGrain string `pulumi:"timeGrain"`
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriodResponse `pulumi:"timePeriod"`
	// Resource type.
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
	// The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
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

func (o LookupBudgetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBudgetResult] {
	return pulumix.Output[LookupBudgetResult]{
		OutputState: o.OutputState,
	}
}

// The total amount of cost to track with the budget
func (o LookupBudgetResultOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBudgetResult) float64 { return v.Amount }).(pulumi.Float64Output)
}

// The category of the budget, whether the budget tracks cost or usage.
func (o LookupBudgetResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Category }).(pulumi.StringOutput)
}

// The current amount of cost which is being tracked for a budget.
func (o LookupBudgetResultOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) CurrentSpendResponse { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o LookupBudgetResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBudgetResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// May be used to filter budgets by user-specified dimensions and/or tags.
func (o LookupBudgetResultOutput) Filter() BudgetFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupBudgetResult) *BudgetFilterResponse { return v.Filter }).(BudgetFilterResponsePtrOutput)
}

// The forecasted cost which is being tracked for a budget.
func (o LookupBudgetResultOutput) ForecastSpend() ForecastSpendResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) ForecastSpendResponse { return v.ForecastSpend }).(ForecastSpendResponseOutput)
}

// Resource Id.
func (o LookupBudgetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupBudgetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
func (o LookupBudgetResultOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v LookupBudgetResult) map[string]NotificationResponse { return v.Notifications }).(NotificationResponseMapOutput)
}

// The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
func (o LookupBudgetResultOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.TimeGrain }).(pulumi.StringOutput)
}

// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
func (o LookupBudgetResultOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v LookupBudgetResult) BudgetTimePeriodResponse { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

// Resource type.
func (o LookupBudgetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBudgetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBudgetResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A budget resource.
type Budget struct {
	pulumi.CustomResourceState

	// The total amount of cost to track with the budget.
	//
	//  Supported for CategoryType(s): Cost.
	//
	//  Required for CategoryType(s): Cost.
	Amount pulumi.Float64PtrOutput `pulumi:"amount"`
	// The category of the budget.
	// - 'Cost' defines a Budget.
	// - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
	Category pulumi.StringOutput `pulumi:"category"`
	// The current amount of cost which is being tracked for a budget.
	//
	//  Supported for CategoryType(s): Cost.
	CurrentSpend CurrentSpendResponseOutput `pulumi:"currentSpend"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// May be used to filter budgets by user-specified dimensions and/or tags.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	Filter BudgetFilterResponsePtrOutput `pulumi:"filter"`
	// The forecasted cost which is being tracked for a budget.
	//
	//  Supported for CategoryType(s): Cost.
	ForecastSpend ForecastSpendResponseOutput `pulumi:"forecastSpend"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Dictionary of notifications associated with the budget.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	//
	// - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
	// - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
	Notifications NotificationResponseMapOutput `pulumi:"notifications"`
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
	TimeGrain pulumi.StringOutput `pulumi:"timeGrain"`
	// The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	//
	//  Required for CategoryType(s): Cost, ReservationUtilization.
	TimePeriod BudgetTimePeriodResponseOutput `pulumi:"timePeriod"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBudget registers a new resource with the given unique name, arguments, and options.
func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOption) (*Budget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.TimeGrain == nil {
		return nil, errors.New("invalid value for required argument 'TimeGrain'")
	}
	if args.TimePeriod == nil {
		return nil, errors.New("invalid value for required argument 'TimePeriod'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:Budget"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190401preview:Budget"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230801:Budget"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Budget
	err := ctx.RegisterResource("azure-native:costmanagement/v20230401preview:Budget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudget gets an existing Budget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetState, opts ...pulumi.ResourceOption) (*Budget, error) {
	var resource Budget
	err := ctx.ReadResource("azure-native:costmanagement/v20230401preview:Budget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Budget resources.
type budgetState struct {
}

type BudgetState struct {
}

func (BudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetState)(nil)).Elem()
}

type budgetArgs struct {
	// The total amount of cost to track with the budget.
	//
	//  Supported for CategoryType(s): Cost.
	//
	//  Required for CategoryType(s): Cost.
	Amount *float64 `pulumi:"amount"`
	// Budget Name.
	BudgetName *string `pulumi:"budgetName"`
	// The category of the budget.
	// - 'Cost' defines a Budget.
	// - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
	Category string `pulumi:"category"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// May be used to filter budgets by user-specified dimensions and/or tags.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	Filter *BudgetFilter `pulumi:"filter"`
	// Dictionary of notifications associated with the budget.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	//
	// - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
	// - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
	Notifications map[string]Notification `pulumi:"notifications"`
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
	TimePeriod BudgetTimePeriod `pulumi:"timePeriod"`
}

// The set of arguments for constructing a Budget resource.
type BudgetArgs struct {
	// The total amount of cost to track with the budget.
	//
	//  Supported for CategoryType(s): Cost.
	//
	//  Required for CategoryType(s): Cost.
	Amount pulumi.Float64PtrInput
	// Budget Name.
	BudgetName pulumi.StringPtrInput
	// The category of the budget.
	// - 'Cost' defines a Budget.
	// - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
	Category pulumi.StringInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// May be used to filter budgets by user-specified dimensions and/or tags.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	Filter BudgetFilterPtrInput
	// Dictionary of notifications associated with the budget.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	//
	// - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
	// - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
	Notifications NotificationMapInput
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
	Scope pulumi.StringInput
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
	TimeGrain pulumi.StringInput
	// The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
	//
	//  Supported for CategoryType(s): Cost, ReservationUtilization.
	//
	//  Required for CategoryType(s): Cost, ReservationUtilization.
	TimePeriod BudgetTimePeriodInput
}

func (BudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetArgs)(nil)).Elem()
}

type BudgetInput interface {
	pulumi.Input

	ToBudgetOutput() BudgetOutput
	ToBudgetOutputWithContext(ctx context.Context) BudgetOutput
}

func (*Budget) ElementType() reflect.Type {
	return reflect.TypeOf((**Budget)(nil)).Elem()
}

func (i *Budget) ToBudgetOutput() BudgetOutput {
	return i.ToBudgetOutputWithContext(context.Background())
}

func (i *Budget) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetOutput)
}

func (i *Budget) ToOutput(ctx context.Context) pulumix.Output[*Budget] {
	return pulumix.Output[*Budget]{
		OutputState: i.ToBudgetOutputWithContext(ctx).OutputState,
	}
}

type BudgetOutput struct{ *pulumi.OutputState }

func (BudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Budget)(nil)).Elem()
}

func (o BudgetOutput) ToBudgetOutput() BudgetOutput {
	return o
}

func (o BudgetOutput) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return o
}

func (o BudgetOutput) ToOutput(ctx context.Context) pulumix.Output[*Budget] {
	return pulumix.Output[*Budget]{
		OutputState: o.OutputState,
	}
}

// The total amount of cost to track with the budget.
//
//	Supported for CategoryType(s): Cost.
//
//	Required for CategoryType(s): Cost.
func (o BudgetOutput) Amount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Budget) pulumi.Float64PtrOutput { return v.Amount }).(pulumi.Float64PtrOutput)
}

// The category of the budget.
// - 'Cost' defines a Budget.
// - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
func (o BudgetOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// The current amount of cost which is being tracked for a budget.
//
//	Supported for CategoryType(s): Cost.
func (o BudgetOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v *Budget) CurrentSpendResponseOutput { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o BudgetOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// May be used to filter budgets by user-specified dimensions and/or tags.
//
//	Supported for CategoryType(s): Cost, ReservationUtilization.
func (o BudgetOutput) Filter() BudgetFilterResponsePtrOutput {
	return o.ApplyT(func(v *Budget) BudgetFilterResponsePtrOutput { return v.Filter }).(BudgetFilterResponsePtrOutput)
}

// The forecasted cost which is being tracked for a budget.
//
//	Supported for CategoryType(s): Cost.
func (o BudgetOutput) ForecastSpend() ForecastSpendResponseOutput {
	return o.ApplyT(func(v *Budget) ForecastSpendResponseOutput { return v.ForecastSpend }).(ForecastSpendResponseOutput)
}

// The name of the resource
func (o BudgetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Dictionary of notifications associated with the budget.
//
//	Supported for CategoryType(s): Cost, ReservationUtilization.
//
// - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
// - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
func (o BudgetOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v *Budget) NotificationResponseMapOutput { return v.Notifications }).(NotificationResponseMapOutput)
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
func (o BudgetOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.TimeGrain }).(pulumi.StringOutput)
}

// The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
//
//	Supported for CategoryType(s): Cost, ReservationUtilization.
//
//	Required for CategoryType(s): Cost, ReservationUtilization.
func (o BudgetOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v *Budget) BudgetTimePeriodResponseOutput { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BudgetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BudgetOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package consumption

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A budget resource.
// Azure REST API version: 2023-05-01. Prior API version in Azure Native 1.x: 2019-10-01.
//
// Other available API versions: 2017-12-30-preview, 2018-10-01, 2019-05-01, 2019-06-01, 2023-11-01.
type Budget struct {
	pulumi.CustomResourceState

	// The total amount of cost to track with the budget
	Amount pulumi.Float64Output `pulumi:"amount"`
	// The category of the budget, whether the budget tracks cost or usage.
	Category pulumi.StringOutput `pulumi:"category"`
	// The current amount of cost which is being tracked for a budget.
	CurrentSpend CurrentSpendResponseOutput `pulumi:"currentSpend"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// May be used to filter budgets by user-specified dimensions and/or tags.
	Filter BudgetFilterResponsePtrOutput `pulumi:"filter"`
	// The forecasted cost which is being tracked for a budget.
	ForecastSpend ForecastSpendResponseOutput `pulumi:"forecastSpend"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications NotificationResponseMapOutput `pulumi:"notifications"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
	TimeGrain pulumi.StringOutput `pulumi:"timeGrain"`
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriodResponseOutput `pulumi:"timePeriod"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBudget registers a new resource with the given unique name, arguments, and options.
func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOption) (*Budget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Amount == nil {
		return nil, errors.New("invalid value for required argument 'Amount'")
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
			Type: pulumi.String("azure-native:consumption/v20190101:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20190401preview:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20190501:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20190501preview:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20190601:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20191001:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20191101:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20210501:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20211001:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20220901:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20230301:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20230501:Budget"),
		},
		{
			Type: pulumi.String("azure-native:consumption/v20231101:Budget"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Budget
	err := ctx.RegisterResource("azure-native:consumption:Budget", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:consumption:Budget", name, id, state, &resource, opts...)
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
	// The total amount of cost to track with the budget
	Amount float64 `pulumi:"amount"`
	// Budget Name.
	BudgetName *string `pulumi:"budgetName"`
	// The category of the budget, whether the budget tracks cost or usage.
	Category string `pulumi:"category"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// May be used to filter budgets by user-specified dimensions and/or tags.
	Filter *BudgetFilter `pulumi:"filter"`
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications map[string]Notification `pulumi:"notifications"`
	// The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
	Scope string `pulumi:"scope"`
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
	TimeGrain string `pulumi:"timeGrain"`
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
	TimePeriod BudgetTimePeriod `pulumi:"timePeriod"`
}

// The set of arguments for constructing a Budget resource.
type BudgetArgs struct {
	// The total amount of cost to track with the budget
	Amount pulumi.Float64Input
	// Budget Name.
	BudgetName pulumi.StringPtrInput
	// The category of the budget, whether the budget tracks cost or usage.
	Category pulumi.StringInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// May be used to filter budgets by user-specified dimensions and/or tags.
	Filter BudgetFilterPtrInput
	// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
	Notifications NotificationMapInput
	// The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
	Scope pulumi.StringInput
	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
	TimeGrain pulumi.StringInput
	// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
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

// The total amount of cost to track with the budget
func (o BudgetOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v *Budget) pulumi.Float64Output { return v.Amount }).(pulumi.Float64Output)
}

// The category of the budget, whether the budget tracks cost or usage.
func (o BudgetOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// The current amount of cost which is being tracked for a budget.
func (o BudgetOutput) CurrentSpend() CurrentSpendResponseOutput {
	return o.ApplyT(func(v *Budget) CurrentSpendResponseOutput { return v.CurrentSpend }).(CurrentSpendResponseOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o BudgetOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// May be used to filter budgets by user-specified dimensions and/or tags.
func (o BudgetOutput) Filter() BudgetFilterResponsePtrOutput {
	return o.ApplyT(func(v *Budget) BudgetFilterResponsePtrOutput { return v.Filter }).(BudgetFilterResponsePtrOutput)
}

// The forecasted cost which is being tracked for a budget.
func (o BudgetOutput) ForecastSpend() ForecastSpendResponseOutput {
	return o.ApplyT(func(v *Budget) ForecastSpendResponseOutput { return v.ForecastSpend }).(ForecastSpendResponseOutput)
}

// Resource name.
func (o BudgetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
func (o BudgetOutput) Notifications() NotificationResponseMapOutput {
	return o.ApplyT(func(v *Budget) NotificationResponseMapOutput { return v.Notifications }).(NotificationResponseMapOutput)
}

// The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
func (o BudgetOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.TimeGrain }).(pulumi.StringOutput)
}

// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
func (o BudgetOutput) TimePeriod() BudgetTimePeriodResponseOutput {
	return o.ApplyT(func(v *Budget) BudgetTimePeriodResponseOutput { return v.TimePeriod }).(BudgetTimePeriodResponseOutput)
}

// Resource type.
func (o BudgetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Budget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BudgetOutput{})
}

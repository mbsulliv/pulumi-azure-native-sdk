// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the shared scheduled action from the given scope by name.
func LookupScheduledActionByScope(ctx *pulumi.Context, args *LookupScheduledActionByScopeArgs, opts ...pulumi.InvokeOption) (*LookupScheduledActionByScopeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupScheduledActionByScopeResult
	err := ctx.Invoke("azure-native:costmanagement/v20230401preview:getScheduledActionByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledActionByScopeArgs struct {
	// Scheduled action name.
	Name string `pulumi:"name"`
	// The scope associated with scheduled action operations. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External Billing Account scope and 'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope. Note: Insight Alerts are only available on subscription scope.
	Scope string `pulumi:"scope"`
}

// Scheduled action definition.
type LookupScheduledActionByScopeResult struct {
	// Scheduled action name.
	DisplayName string `pulumi:"displayName"`
	// Resource Etag. For update calls, eTag is mandatory. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
	ETag string `pulumi:"eTag"`
	// Destination format of the view data. This is optional.
	FileDestination *FileDestinationResponse `pulumi:"fileDestination"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Kind of the scheduled action.
	Kind *string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Notification properties based on scheduled action kind.
	Notification NotificationPropertiesResponse `pulumi:"notification"`
	// Email address of the point of contact that should get the unsubscribe requests and notification emails.
	NotificationEmail *string `pulumi:"notificationEmail"`
	// Schedule of the scheduled action.
	Schedule SchedulePropertiesResponse `pulumi:"schedule"`
	// Cost Management scope like 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope *string `pulumi:"scope"`
	// Status of the scheduled action.
	Status string `pulumi:"status"`
	// Kind of the scheduled action.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
	ViewId string `pulumi:"viewId"`
}

func LookupScheduledActionByScopeOutput(ctx *pulumi.Context, args LookupScheduledActionByScopeOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledActionByScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledActionByScopeResult, error) {
			args := v.(LookupScheduledActionByScopeArgs)
			r, err := LookupScheduledActionByScope(ctx, &args, opts...)
			var s LookupScheduledActionByScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledActionByScopeResultOutput)
}

type LookupScheduledActionByScopeOutputArgs struct {
	// Scheduled action name.
	Name pulumi.StringInput `pulumi:"name"`
	// The scope associated with scheduled action operations. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External Billing Account scope and 'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope. Note: Insight Alerts are only available on subscription scope.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupScheduledActionByScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionByScopeArgs)(nil)).Elem()
}

// Scheduled action definition.
type LookupScheduledActionByScopeResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledActionByScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionByScopeResult)(nil)).Elem()
}

func (o LookupScheduledActionByScopeResultOutput) ToLookupScheduledActionByScopeResultOutput() LookupScheduledActionByScopeResultOutput {
	return o
}

func (o LookupScheduledActionByScopeResultOutput) ToLookupScheduledActionByScopeResultOutputWithContext(ctx context.Context) LookupScheduledActionByScopeResultOutput {
	return o
}

// Scheduled action name.
func (o LookupScheduledActionByScopeResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource Etag. For update calls, eTag is mandatory. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
func (o LookupScheduledActionByScopeResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.ETag }).(pulumi.StringOutput)
}

// Destination format of the view data. This is optional.
func (o LookupScheduledActionByScopeResultOutput) FileDestination() FileDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) *FileDestinationResponse { return v.FileDestination }).(FileDestinationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupScheduledActionByScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of the scheduled action.
func (o LookupScheduledActionByScopeResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupScheduledActionByScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Notification properties based on scheduled action kind.
func (o LookupScheduledActionByScopeResultOutput) Notification() NotificationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) NotificationPropertiesResponse { return v.Notification }).(NotificationPropertiesResponseOutput)
}

// Email address of the point of contact that should get the unsubscribe requests and notification emails.
func (o LookupScheduledActionByScopeResultOutput) NotificationEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) *string { return v.NotificationEmail }).(pulumi.StringPtrOutput)
}

// Schedule of the scheduled action.
func (o LookupScheduledActionByScopeResultOutput) Schedule() SchedulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) SchedulePropertiesResponse { return v.Schedule }).(SchedulePropertiesResponseOutput)
}

// Cost Management scope like 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
func (o LookupScheduledActionByScopeResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// Status of the scheduled action.
func (o LookupScheduledActionByScopeResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.Status }).(pulumi.StringOutput)
}

// Kind of the scheduled action.
func (o LookupScheduledActionByScopeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupScheduledActionByScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
func (o LookupScheduledActionByScopeResultOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionByScopeResult) string { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledActionByScopeResultOutput{})
}
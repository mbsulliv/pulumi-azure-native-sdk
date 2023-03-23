// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the private scheduled action by name.
func LookupScheduledAction(ctx *pulumi.Context, args *LookupScheduledActionArgs, opts ...pulumi.InvokeOption) (*LookupScheduledActionResult, error) {
	var rv LookupScheduledActionResult
	err := ctx.Invoke("azure-native:costmanagement/v20221001:getScheduledAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledActionArgs struct {
	// Scheduled action name.
	Name string `pulumi:"name"`
}

// Scheduled action definition.
type LookupScheduledActionResult struct {
	// Scheduled action name.
	DisplayName string `pulumi:"displayName"`
	// Resource Etag. For update calls, eTag is optional and can be specified to achieve optimistic concurrency. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
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

func LookupScheduledActionOutput(ctx *pulumi.Context, args LookupScheduledActionOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledActionResult, error) {
			args := v.(LookupScheduledActionArgs)
			r, err := LookupScheduledAction(ctx, &args, opts...)
			var s LookupScheduledActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledActionResultOutput)
}

type LookupScheduledActionOutputArgs struct {
	// Scheduled action name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupScheduledActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionArgs)(nil)).Elem()
}

// Scheduled action definition.
type LookupScheduledActionResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionResult)(nil)).Elem()
}

func (o LookupScheduledActionResultOutput) ToLookupScheduledActionResultOutput() LookupScheduledActionResultOutput {
	return o
}

func (o LookupScheduledActionResultOutput) ToLookupScheduledActionResultOutputWithContext(ctx context.Context) LookupScheduledActionResultOutput {
	return o
}

// Scheduled action name.
func (o LookupScheduledActionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource Etag. For update calls, eTag is optional and can be specified to achieve optimistic concurrency. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
func (o LookupScheduledActionResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.ETag }).(pulumi.StringOutput)
}

// Destination format of the view data. This is optional.
func (o LookupScheduledActionResultOutput) FileDestination() FileDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *FileDestinationResponse { return v.FileDestination }).(FileDestinationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupScheduledActionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of the scheduled action.
func (o LookupScheduledActionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupScheduledActionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Notification properties based on scheduled action kind.
func (o LookupScheduledActionResultOutput) Notification() NotificationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) NotificationPropertiesResponse { return v.Notification }).(NotificationPropertiesResponseOutput)
}

// Email address of the point of contact that should get the unsubscribe requests and notification emails.
func (o LookupScheduledActionResultOutput) NotificationEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.NotificationEmail }).(pulumi.StringPtrOutput)
}

// Schedule of the scheduled action.
func (o LookupScheduledActionResultOutput) Schedule() SchedulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) SchedulePropertiesResponse { return v.Schedule }).(SchedulePropertiesResponseOutput)
}

// Cost Management scope like 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
func (o LookupScheduledActionResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// Status of the scheduled action.
func (o LookupScheduledActionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.Status }).(pulumi.StringOutput)
}

// Kind of the scheduled action.
func (o LookupScheduledActionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupScheduledActionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
func (o LookupScheduledActionResultOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) string { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledActionResultOutput{})
}

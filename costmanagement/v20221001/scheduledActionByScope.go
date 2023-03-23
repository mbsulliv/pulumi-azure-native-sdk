// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Scheduled action definition.
type ScheduledActionByScope struct {
	pulumi.CustomResourceState

	// Scheduled action name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource Etag. For update calls, eTag is optional and can be specified to achieve optimistic concurrency. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// Destination format of the view data. This is optional.
	FileDestination FileDestinationResponsePtrOutput `pulumi:"fileDestination"`
	// Kind of the scheduled action.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification properties based on scheduled action kind.
	Notification NotificationPropertiesResponseOutput `pulumi:"notification"`
	// Email address of the point of contact that should get the unsubscribe requests and notification emails.
	NotificationEmail pulumi.StringPtrOutput `pulumi:"notificationEmail"`
	// Schedule of the scheduled action.
	Schedule SchedulePropertiesResponseOutput `pulumi:"schedule"`
	// Cost Management scope like 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Status of the scheduled action.
	Status pulumi.StringOutput `pulumi:"status"`
	// Kind of the scheduled action.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
	ViewId pulumi.StringOutput `pulumi:"viewId"`
}

// NewScheduledActionByScope registers a new resource with the given unique name, arguments, and options.
func NewScheduledActionByScope(ctx *pulumi.Context,
	name string, args *ScheduledActionByScopeArgs, opts ...pulumi.ResourceOption) (*ScheduledActionByScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Notification == nil {
		return nil, errors.New("invalid value for required argument 'Notification'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.ViewId == nil {
		return nil, errors.New("invalid value for required argument 'ViewId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:ScheduledActionByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20220401preview:ScheduledActionByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20220601preview:ScheduledActionByScope"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledActionByScope
	err := ctx.RegisterResource("azure-native:costmanagement/v20221001:ScheduledActionByScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledActionByScope gets an existing ScheduledActionByScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledActionByScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledActionByScopeState, opts ...pulumi.ResourceOption) (*ScheduledActionByScope, error) {
	var resource ScheduledActionByScope
	err := ctx.ReadResource("azure-native:costmanagement/v20221001:ScheduledActionByScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledActionByScope resources.
type scheduledActionByScopeState struct {
}

type ScheduledActionByScopeState struct {
}

func (ScheduledActionByScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionByScopeState)(nil)).Elem()
}

type scheduledActionByScopeArgs struct {
	// Scheduled action name.
	DisplayName string `pulumi:"displayName"`
	// Destination format of the view data. This is optional.
	FileDestination *FileDestination `pulumi:"fileDestination"`
	// Kind of the scheduled action.
	Kind *string `pulumi:"kind"`
	// Scheduled action name.
	Name *string `pulumi:"name"`
	// Notification properties based on scheduled action kind.
	Notification NotificationProperties `pulumi:"notification"`
	// Email address of the point of contact that should get the unsubscribe requests and notification emails.
	NotificationEmail *string `pulumi:"notificationEmail"`
	// Schedule of the scheduled action.
	Schedule ScheduleProperties `pulumi:"schedule"`
	// Cost Management scope like 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope string `pulumi:"scope"`
	// Status of the scheduled action.
	Status string `pulumi:"status"`
	// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
	ViewId string `pulumi:"viewId"`
}

// The set of arguments for constructing a ScheduledActionByScope resource.
type ScheduledActionByScopeArgs struct {
	// Scheduled action name.
	DisplayName pulumi.StringInput
	// Destination format of the view data. This is optional.
	FileDestination FileDestinationPtrInput
	// Kind of the scheduled action.
	Kind pulumi.StringPtrInput
	// Scheduled action name.
	Name pulumi.StringPtrInput
	// Notification properties based on scheduled action kind.
	Notification NotificationPropertiesInput
	// Email address of the point of contact that should get the unsubscribe requests and notification emails.
	NotificationEmail pulumi.StringPtrInput
	// Schedule of the scheduled action.
	Schedule SchedulePropertiesInput
	// Cost Management scope like 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringInput
	// Status of the scheduled action.
	Status pulumi.StringInput
	// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
	ViewId pulumi.StringInput
}

func (ScheduledActionByScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionByScopeArgs)(nil)).Elem()
}

type ScheduledActionByScopeInput interface {
	pulumi.Input

	ToScheduledActionByScopeOutput() ScheduledActionByScopeOutput
	ToScheduledActionByScopeOutputWithContext(ctx context.Context) ScheduledActionByScopeOutput
}

func (*ScheduledActionByScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledActionByScope)(nil)).Elem()
}

func (i *ScheduledActionByScope) ToScheduledActionByScopeOutput() ScheduledActionByScopeOutput {
	return i.ToScheduledActionByScopeOutputWithContext(context.Background())
}

func (i *ScheduledActionByScope) ToScheduledActionByScopeOutputWithContext(ctx context.Context) ScheduledActionByScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionByScopeOutput)
}

type ScheduledActionByScopeOutput struct{ *pulumi.OutputState }

func (ScheduledActionByScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledActionByScope)(nil)).Elem()
}

func (o ScheduledActionByScopeOutput) ToScheduledActionByScopeOutput() ScheduledActionByScopeOutput {
	return o
}

func (o ScheduledActionByScopeOutput) ToScheduledActionByScopeOutputWithContext(ctx context.Context) ScheduledActionByScopeOutput {
	return o
}

// Scheduled action name.
func (o ScheduledActionByScopeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource Etag. For update calls, eTag is optional and can be specified to achieve optimistic concurrency. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
func (o ScheduledActionByScopeOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// Destination format of the view data. This is optional.
func (o ScheduledActionByScopeOutput) FileDestination() FileDestinationResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) FileDestinationResponsePtrOutput { return v.FileDestination }).(FileDestinationResponsePtrOutput)
}

// Kind of the scheduled action.
func (o ScheduledActionByScopeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ScheduledActionByScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Notification properties based on scheduled action kind.
func (o ScheduledActionByScopeOutput) Notification() NotificationPropertiesResponseOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) NotificationPropertiesResponseOutput { return v.Notification }).(NotificationPropertiesResponseOutput)
}

// Email address of the point of contact that should get the unsubscribe requests and notification emails.
func (o ScheduledActionByScopeOutput) NotificationEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringPtrOutput { return v.NotificationEmail }).(pulumi.StringPtrOutput)
}

// Schedule of the scheduled action.
func (o ScheduledActionByScopeOutput) Schedule() SchedulePropertiesResponseOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) SchedulePropertiesResponseOutput { return v.Schedule }).(SchedulePropertiesResponseOutput)
}

// Cost Management scope like 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
func (o ScheduledActionByScopeOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Status of the scheduled action.
func (o ScheduledActionByScopeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Kind of the scheduled action.
func (o ScheduledActionByScopeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScheduledActionByScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
func (o ScheduledActionByScopeOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledActionByScopeOutput{})
}

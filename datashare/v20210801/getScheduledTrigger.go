// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Trigger in a shareSubscription
func LookupScheduledTrigger(ctx *pulumi.Context, args *LookupScheduledTriggerArgs, opts ...pulumi.InvokeOption) (*LookupScheduledTriggerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupScheduledTriggerResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getScheduledTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledTriggerArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
	// The name of the trigger.
	TriggerName string `pulumi:"triggerName"`
}

// A type of trigger based on schedule
type LookupScheduledTriggerResult struct {
	// Time at which the trigger was created.
	CreatedAt string `pulumi:"createdAt"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of synchronization on trigger.
	// Expected value is 'ScheduleBased'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Gets the provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// Recurrence Interval
	RecurrenceInterval string `pulumi:"recurrenceInterval"`
	// Synchronization mode
	SynchronizationMode *string `pulumi:"synchronizationMode"`
	// Synchronization time
	SynchronizationTime string `pulumi:"synchronizationTime"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Gets the trigger state
	TriggerStatus string `pulumi:"triggerStatus"`
	// Type of the azure resource
	Type string `pulumi:"type"`
	// Name of the user who created the trigger.
	UserName string `pulumi:"userName"`
}

func LookupScheduledTriggerOutput(ctx *pulumi.Context, args LookupScheduledTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledTriggerResult, error) {
			args := v.(LookupScheduledTriggerArgs)
			r, err := LookupScheduledTrigger(ctx, &args, opts...)
			var s LookupScheduledTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledTriggerResultOutput)
}

type LookupScheduledTriggerOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
	// The name of the trigger.
	TriggerName pulumi.StringInput `pulumi:"triggerName"`
}

func (LookupScheduledTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledTriggerArgs)(nil)).Elem()
}

// A type of trigger based on schedule
type LookupScheduledTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledTriggerResult)(nil)).Elem()
}

func (o LookupScheduledTriggerResultOutput) ToLookupScheduledTriggerResultOutput() LookupScheduledTriggerResultOutput {
	return o
}

func (o LookupScheduledTriggerResultOutput) ToLookupScheduledTriggerResultOutputWithContext(ctx context.Context) LookupScheduledTriggerResultOutput {
	return o
}

func (o LookupScheduledTriggerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupScheduledTriggerResult] {
	return pulumix.Output[LookupScheduledTriggerResult]{
		OutputState: o.OutputState,
	}
}

// Time at which the trigger was created.
func (o LookupScheduledTriggerResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupScheduledTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of synchronization on trigger.
// Expected value is 'ScheduleBased'.
func (o LookupScheduledTriggerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupScheduledTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state
func (o LookupScheduledTriggerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Recurrence Interval
func (o LookupScheduledTriggerResultOutput) RecurrenceInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.RecurrenceInterval }).(pulumi.StringOutput)
}

// Synchronization mode
func (o LookupScheduledTriggerResultOutput) SynchronizationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) *string { return v.SynchronizationMode }).(pulumi.StringPtrOutput)
}

// Synchronization time
func (o LookupScheduledTriggerResultOutput) SynchronizationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.SynchronizationTime }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupScheduledTriggerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets the trigger state
func (o LookupScheduledTriggerResultOutput) TriggerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.TriggerStatus }).(pulumi.StringOutput)
}

// Type of the azure resource
func (o LookupScheduledTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

// Name of the user who created the trigger.
func (o LookupScheduledTriggerResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledTriggerResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledTriggerResultOutput{})
}

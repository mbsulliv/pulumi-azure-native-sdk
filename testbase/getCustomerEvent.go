// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Test Base CustomerEvent.
// Azure REST API version: 2022-04-01-preview.
func LookupCustomerEvent(ctx *pulumi.Context, args *LookupCustomerEventArgs, opts ...pulumi.InvokeOption) (*LookupCustomerEventResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomerEventResult
	err := ctx.Invoke("azure-native:testbase:getCustomerEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomerEventArgs struct {
	// The resource name of the Test Base Customer event.
	CustomerEventName string `pulumi:"customerEventName"`
	// The name of the resource group that contains the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

// The Customer Notification Event resource.
type LookupCustomerEventResult struct {
	// The name of the event subscribed to.
	EventName string `pulumi:"eventName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The notification event receivers.
	Receivers []NotificationEventReceiverResponse `pulumi:"receivers"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupCustomerEventOutput(ctx *pulumi.Context, args LookupCustomerEventOutputArgs, opts ...pulumi.InvokeOption) LookupCustomerEventResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomerEventResult, error) {
			args := v.(LookupCustomerEventArgs)
			r, err := LookupCustomerEvent(ctx, &args, opts...)
			var s LookupCustomerEventResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomerEventResultOutput)
}

type LookupCustomerEventOutputArgs struct {
	// The resource name of the Test Base Customer event.
	CustomerEventName pulumi.StringInput `pulumi:"customerEventName"`
	// The name of the resource group that contains the resource.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (LookupCustomerEventOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomerEventArgs)(nil)).Elem()
}

// The Customer Notification Event resource.
type LookupCustomerEventResultOutput struct{ *pulumi.OutputState }

func (LookupCustomerEventResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomerEventResult)(nil)).Elem()
}

func (o LookupCustomerEventResultOutput) ToLookupCustomerEventResultOutput() LookupCustomerEventResultOutput {
	return o
}

func (o LookupCustomerEventResultOutput) ToLookupCustomerEventResultOutputWithContext(ctx context.Context) LookupCustomerEventResultOutput {
	return o
}

func (o LookupCustomerEventResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCustomerEventResult] {
	return pulumix.Output[LookupCustomerEventResult]{
		OutputState: o.OutputState,
	}
}

// The name of the event subscribed to.
func (o LookupCustomerEventResultOutput) EventName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) string { return v.EventName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupCustomerEventResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupCustomerEventResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) string { return v.Name }).(pulumi.StringOutput)
}

// The notification event receivers.
func (o LookupCustomerEventResultOutput) Receivers() NotificationEventReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) []NotificationEventReceiverResponse { return v.Receivers }).(NotificationEventReceiverResponseArrayOutput)
}

// The system metadata relating to this resource
func (o LookupCustomerEventResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupCustomerEventResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomerEventResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomerEventResultOutput{})
}

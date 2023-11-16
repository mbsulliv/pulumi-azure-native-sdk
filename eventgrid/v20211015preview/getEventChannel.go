// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of an event channel.
func LookupEventChannel(ctx *pulumi.Context, args *LookupEventChannelArgs, opts ...pulumi.InvokeOption) (*LookupEventChannelResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEventChannelResult
	err := ctx.Invoke("azure-native:eventgrid/v20211015preview:getEventChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEventChannelArgs struct {
	// Name of the event channel.
	EventChannelName string `pulumi:"eventChannelName"`
	// Name of the partner namespace.
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Event Channel.
type LookupEventChannelResult struct {
	// Represents the destination of an event channel.
	Destination *EventChannelDestinationResponse `pulumi:"destination"`
	// Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
	// the event channel and corresponding partner topic are deleted.
	ExpirationTimeIfNotActivatedUtc *string `pulumi:"expirationTimeIfNotActivatedUtc"`
	// Information about the filter for the event channel.
	Filter *EventChannelFilterResponse `pulumi:"filter"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription *string `pulumi:"partnerTopicFriendlyDescription"`
	// The readiness state of the corresponding partner topic.
	PartnerTopicReadinessState string `pulumi:"partnerTopicReadinessState"`
	// Provisioning state of the event channel.
	ProvisioningState string `pulumi:"provisioningState"`
	// Source of the event channel. This represents a unique resource in the partner's resource model.
	Source *EventChannelSourceResponse `pulumi:"source"`
	// The system metadata relating to Event Channel resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupEventChannelResult
func (val *LookupEventChannelResult) Defaults() *LookupEventChannelResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Filter = tmp.Filter.Defaults()

	return &tmp
}

func LookupEventChannelOutput(ctx *pulumi.Context, args LookupEventChannelOutputArgs, opts ...pulumi.InvokeOption) LookupEventChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventChannelResult, error) {
			args := v.(LookupEventChannelArgs)
			r, err := LookupEventChannel(ctx, &args, opts...)
			var s LookupEventChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventChannelResultOutput)
}

type LookupEventChannelOutputArgs struct {
	// Name of the event channel.
	EventChannelName pulumi.StringInput `pulumi:"eventChannelName"`
	// Name of the partner namespace.
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventChannelArgs)(nil)).Elem()
}

// Event Channel.
type LookupEventChannelResultOutput struct{ *pulumi.OutputState }

func (LookupEventChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventChannelResult)(nil)).Elem()
}

func (o LookupEventChannelResultOutput) ToLookupEventChannelResultOutput() LookupEventChannelResultOutput {
	return o
}

func (o LookupEventChannelResultOutput) ToLookupEventChannelResultOutputWithContext(ctx context.Context) LookupEventChannelResultOutput {
	return o
}

// Represents the destination of an event channel.
func (o LookupEventChannelResultOutput) Destination() EventChannelDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *EventChannelDestinationResponse { return v.Destination }).(EventChannelDestinationResponsePtrOutput)
}

// Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
// the event channel and corresponding partner topic are deleted.
func (o LookupEventChannelResultOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *string { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

// Information about the filter for the event channel.
func (o LookupEventChannelResultOutput) Filter() EventChannelFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *EventChannelFilterResponse { return v.Filter }).(EventChannelFilterResponsePtrOutput)
}

// Fully qualified identifier of the resource.
func (o LookupEventChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LookupEventChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
func (o LookupEventChannelResultOutput) PartnerTopicFriendlyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *string { return v.PartnerTopicFriendlyDescription }).(pulumi.StringPtrOutput)
}

// The readiness state of the corresponding partner topic.
func (o LookupEventChannelResultOutput) PartnerTopicReadinessState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.PartnerTopicReadinessState }).(pulumi.StringOutput)
}

// Provisioning state of the event channel.
func (o LookupEventChannelResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Source of the event channel. This represents a unique resource in the partner's resource model.
func (o LookupEventChannelResultOutput) Source() EventChannelSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupEventChannelResult) *EventChannelSourceResponse { return v.Source }).(EventChannelSourceResponsePtrOutput)
}

// The system metadata relating to Event Channel resource.
func (o LookupEventChannelResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEventChannelResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource.
func (o LookupEventChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventChannelResultOutput{})
}

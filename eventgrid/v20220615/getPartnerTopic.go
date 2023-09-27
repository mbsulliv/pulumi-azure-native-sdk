// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of a partner topic.
func LookupPartnerTopic(ctx *pulumi.Context, args *LookupPartnerTopicArgs, opts ...pulumi.InvokeOption) (*LookupPartnerTopicResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPartnerTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getPartnerTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerTopicArgs struct {
	// Name of the partner topic.
	PartnerTopicName string `pulumi:"partnerTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Event Grid Partner Topic.
type LookupPartnerTopicResult struct {
	// Activation state of the partner topic.
	ActivationState *string `pulumi:"activationState"`
	// Event Type information from the corresponding event channel.
	EventTypeInfo *EventTypeInfoResponse `pulumi:"eventTypeInfo"`
	// Expiration time of the partner topic. If this timer expires while the partner topic is still never activated,
	// the partner topic and corresponding event channel are deleted.
	ExpirationTimeIfNotActivatedUtc *string `pulumi:"expirationTimeIfNotActivatedUtc"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Identity information for the Partner Topic resource.
	Identity *IdentityInfoResponse `pulumi:"identity"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Context or helpful message that can be used during the approval process by the subscriber.
	MessageForActivation *string `pulumi:"messageForActivation"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// The immutableId of the corresponding partner registration.
	PartnerRegistrationImmutableId *string `pulumi:"partnerRegistrationImmutableId"`
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription *string `pulumi:"partnerTopicFriendlyDescription"`
	// Provisioning state of the partner topic.
	ProvisioningState string `pulumi:"provisioningState"`
	// Source associated with this partner topic. This represents a unique partner resource.
	Source *string `pulumi:"source"`
	// The system metadata relating to Partner Topic resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupPartnerTopicOutput(ctx *pulumi.Context, args LookupPartnerTopicOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerTopicResult, error) {
			args := v.(LookupPartnerTopicArgs)
			r, err := LookupPartnerTopic(ctx, &args, opts...)
			var s LookupPartnerTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerTopicResultOutput)
}

type LookupPartnerTopicOutputArgs struct {
	// Name of the partner topic.
	PartnerTopicName pulumi.StringInput `pulumi:"partnerTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerTopicArgs)(nil)).Elem()
}

// Event Grid Partner Topic.
type LookupPartnerTopicResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerTopicResult)(nil)).Elem()
}

func (o LookupPartnerTopicResultOutput) ToLookupPartnerTopicResultOutput() LookupPartnerTopicResultOutput {
	return o
}

func (o LookupPartnerTopicResultOutput) ToLookupPartnerTopicResultOutputWithContext(ctx context.Context) LookupPartnerTopicResultOutput {
	return o
}

func (o LookupPartnerTopicResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPartnerTopicResult] {
	return pulumix.Output[LookupPartnerTopicResult]{
		OutputState: o.OutputState,
	}
}

// Activation state of the partner topic.
func (o LookupPartnerTopicResultOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.ActivationState }).(pulumi.StringPtrOutput)
}

// Event Type information from the corresponding event channel.
func (o LookupPartnerTopicResultOutput) EventTypeInfo() EventTypeInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *EventTypeInfoResponse { return v.EventTypeInfo }).(EventTypeInfoResponsePtrOutput)
}

// Expiration time of the partner topic. If this timer expires while the partner topic is still never activated,
// the partner topic and corresponding event channel are deleted.
func (o LookupPartnerTopicResultOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

// Fully qualified identifier of the resource.
func (o LookupPartnerTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity information for the Partner Topic resource.
func (o LookupPartnerTopicResultOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *IdentityInfoResponse { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

// Location of the resource.
func (o LookupPartnerTopicResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.Location }).(pulumi.StringOutput)
}

// Context or helpful message that can be used during the approval process by the subscriber.
func (o LookupPartnerTopicResultOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupPartnerTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

// The immutableId of the corresponding partner registration.
func (o LookupPartnerTopicResultOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
func (o LookupPartnerTopicResultOutput) PartnerTopicFriendlyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.PartnerTopicFriendlyDescription }).(pulumi.StringPtrOutput)
}

// Provisioning state of the partner topic.
func (o LookupPartnerTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Source associated with this partner topic. This represents a unique partner resource.
func (o LookupPartnerTopicResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

// The system metadata relating to Partner Topic resource.
func (o LookupPartnerTopicResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o LookupPartnerTopicResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o LookupPartnerTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerTopicResultOutput{})
}

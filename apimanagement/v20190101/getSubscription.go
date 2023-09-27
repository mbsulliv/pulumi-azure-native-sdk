// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified Subscription entity.
func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSubscriptionResult
	err := ctx.Invoke("azure-native:apimanagement/v20190101:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid string `pulumi:"sid"`
}

// Subscription details.
type LookupSubscriptionResult struct {
	// Determines whether tracing is enabled
	AllowTracing *bool `pulumi:"allowTracing"`
	// Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedDate string `pulumi:"createdDate"`
	// The name of the subscription, or null if the subscription has no name.
	DisplayName *string `pulumi:"displayName"`
	// Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	EndDate *string `pulumi:"endDate"`
	// Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	ExpirationDate *string `pulumi:"expirationDate"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	NotificationDate *string `pulumi:"notificationDate"`
	// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{userId} where {userId} is a user identifier.
	OwnerId *string `pulumi:"ownerId"`
	// Subscription primary key.
	PrimaryKey string `pulumi:"primaryKey"`
	// Scope like /products/{productId} or /apis or /apis/{apiId}.
	Scope string `pulumi:"scope"`
	// Subscription secondary key.
	SecondaryKey string `pulumi:"secondaryKey"`
	// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	StartDate *string `pulumi:"startDate"`
	// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State string `pulumi:"state"`
	// Optional subscription comment added by an administrator.
	StateComment *string `pulumi:"stateComment"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}

func LookupSubscriptionOutput(ctx *pulumi.Context, args LookupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionResult, error) {
			args := v.(LookupSubscriptionArgs)
			r, err := LookupSubscription(ctx, &args, opts...)
			var s LookupSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionResultOutput)
}

type LookupSubscriptionOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid pulumi.StringInput `pulumi:"sid"`
}

func (LookupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionArgs)(nil)).Elem()
}

// Subscription details.
type LookupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionResult)(nil)).Elem()
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutput() LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutputWithContext(ctx context.Context) LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSubscriptionResult] {
	return pulumix.Output[LookupSubscriptionResult]{
		OutputState: o.OutputState,
	}
}

// Determines whether tracing is enabled
func (o LookupSubscriptionResultOutput) AllowTracing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *bool { return v.AllowTracing }).(pulumi.BoolPtrOutput)
}

// Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupSubscriptionResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

// The name of the subscription, or null if the subscription has no name.
func (o LookupSubscriptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupSubscriptionResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

// Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupSubscriptionResultOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupSubscriptionResultOutput) NotificationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.NotificationDate }).(pulumi.StringPtrOutput)
}

// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{userId} where {userId} is a user identifier.
func (o LookupSubscriptionResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Subscription primary key.
func (o LookupSubscriptionResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Scope like /products/{productId} or /apis or /apis/{apiId}.
func (o LookupSubscriptionResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Subscription secondary key.
func (o LookupSubscriptionResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupSubscriptionResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
func (o LookupSubscriptionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.State }).(pulumi.StringOutput)
}

// Optional subscription comment added by an administrator.
func (o LookupSubscriptionResultOutput) StateComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.StateComment }).(pulumi.StringPtrOutput)
}

// Resource type for API Management resource.
func (o LookupSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionResultOutput{})
}

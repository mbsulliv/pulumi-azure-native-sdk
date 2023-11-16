// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Subscription entity.
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
func LookupWorkspaceSubscription(ctx *pulumi.Context, args *LookupWorkspaceSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceSubscriptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceSubscriptionResult
	err := ctx.Invoke("azure-native:apimanagement:getWorkspaceSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceSubscriptionArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid string `pulumi:"sid"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Subscription details.
type LookupWorkspaceSubscriptionResult struct {
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
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	NotificationDate *string `pulumi:"notificationDate"`
	// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{userId} where {userId} is a user identifier.
	OwnerId *string `pulumi:"ownerId"`
	// Subscription primary key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Scope like /products/{productId} or /apis or /apis/{apiId}.
	Scope string `pulumi:"scope"`
	// Subscription secondary key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	StartDate *string `pulumi:"startDate"`
	// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State string `pulumi:"state"`
	// Optional subscription comment added by an administrator when the state is changed to the 'rejected'.
	StateComment *string `pulumi:"stateComment"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceSubscriptionOutput(ctx *pulumi.Context, args LookupWorkspaceSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceSubscriptionResult, error) {
			args := v.(LookupWorkspaceSubscriptionArgs)
			r, err := LookupWorkspaceSubscription(ctx, &args, opts...)
			var s LookupWorkspaceSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceSubscriptionResultOutput)
}

type LookupWorkspaceSubscriptionOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid pulumi.StringInput `pulumi:"sid"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSubscriptionArgs)(nil)).Elem()
}

// Subscription details.
type LookupWorkspaceSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceSubscriptionResult)(nil)).Elem()
}

func (o LookupWorkspaceSubscriptionResultOutput) ToLookupWorkspaceSubscriptionResultOutput() LookupWorkspaceSubscriptionResultOutput {
	return o
}

func (o LookupWorkspaceSubscriptionResultOutput) ToLookupWorkspaceSubscriptionResultOutputWithContext(ctx context.Context) LookupWorkspaceSubscriptionResultOutput {
	return o
}

// Determines whether tracing is enabled
func (o LookupWorkspaceSubscriptionResultOutput) AllowTracing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *bool { return v.AllowTracing }).(pulumi.BoolPtrOutput)
}

// Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupWorkspaceSubscriptionResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

// The name of the subscription, or null if the subscription has no name.
func (o LookupWorkspaceSubscriptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupWorkspaceSubscriptionResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

// Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupWorkspaceSubscriptionResultOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupWorkspaceSubscriptionResultOutput) NotificationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.NotificationDate }).(pulumi.StringPtrOutput)
}

// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{userId} where {userId} is a user identifier.
func (o LookupWorkspaceSubscriptionResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Subscription primary key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o LookupWorkspaceSubscriptionResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// Scope like /products/{productId} or /apis or /apis/{apiId}.
func (o LookupWorkspaceSubscriptionResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Subscription secondary key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o LookupWorkspaceSubscriptionResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupWorkspaceSubscriptionResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
func (o LookupWorkspaceSubscriptionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) string { return v.State }).(pulumi.StringOutput)
}

// Optional subscription comment added by an administrator when the state is changed to the 'rejected'.
func (o LookupWorkspaceSubscriptionResultOutput) StateComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) *string { return v.StateComment }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceSubscriptionResultOutput{})
}

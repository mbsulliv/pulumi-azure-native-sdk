// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Subscription details.
type Subscription struct {
	pulumi.CustomResourceState

	// Determines whether tracing is enabled
	AllowTracing pulumi.BoolPtrOutput `pulumi:"allowTracing"`
	// Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The name of the subscription, or null if the subscription has no name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	ExpirationDate pulumi.StringPtrOutput `pulumi:"expirationDate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	NotificationDate pulumi.StringPtrOutput `pulumi:"notificationDate"`
	// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{userId} where {userId} is a user identifier.
	OwnerId pulumi.StringPtrOutput `pulumi:"ownerId"`
	// Subscription primary key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	PrimaryKey pulumi.StringPtrOutput `pulumi:"primaryKey"`
	// Scope like /products/{productId} or /apis or /apis/{apiId}.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Subscription secondary key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	SecondaryKey pulumi.StringPtrOutput `pulumi:"secondaryKey"`
	// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	StartDate pulumi.StringPtrOutput `pulumi:"startDate"`
	// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State pulumi.StringOutput `pulumi:"state"`
	// Optional subscription comment added by an administrator when the state is changed to the 'rejected'.
	StateComment pulumi.StringPtrOutput `pulumi:"stateComment"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:Subscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Subscription
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
}

type SubscriptionState struct {
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// Determines whether tracing can be enabled
	AllowTracing *bool `pulumi:"allowTracing"`
	// Determines the type of application which send the create user request. Default is legacy publisher portal.
	AppType *string `pulumi:"appType"`
	// Subscription name.
	DisplayName string `pulumi:"displayName"`
	// Notify change in Subscription State.
	//  - If false, do not send any email notification for change of state of subscription
	//  - If true, send email notification of change of state of subscription
	Notify *bool `pulumi:"notify"`
	// User (user id path) for whom subscription is being created in form /users/{userId}
	OwnerId *string `pulumi:"ownerId"`
	// Primary subscription key. If not specified during request key will be generated automatically.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Scope like /products/{productId} or /apis or /apis/{apiId}.
	Scope string `pulumi:"scope"`
	// Secondary subscription key. If not specified during request key will be generated automatically.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid *string `pulumi:"sid"`
	// Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State *SubscriptionStateEnum `pulumi:"state"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// Determines whether tracing can be enabled
	AllowTracing pulumi.BoolPtrInput
	// Determines the type of application which send the create user request. Default is legacy publisher portal.
	AppType pulumi.StringPtrInput
	// Subscription name.
	DisplayName pulumi.StringInput
	// Notify change in Subscription State.
	//  - If false, do not send any email notification for change of state of subscription
	//  - If true, send email notification of change of state of subscription
	Notify pulumi.BoolPtrInput
	// User (user id path) for whom subscription is being created in form /users/{userId}
	OwnerId pulumi.StringPtrInput
	// Primary subscription key. If not specified during request key will be generated automatically.
	PrimaryKey pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Scope like /products/{productId} or /apis or /apis/{apiId}.
	Scope pulumi.StringInput
	// Secondary subscription key. If not specified during request key will be generated automatically.
	SecondaryKey pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid pulumi.StringPtrInput
	// Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State SubscriptionStateEnumPtrInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

func (i *Subscription) ToOutput(ctx context.Context) pulumix.Output[*Subscription] {
	return pulumix.Output[*Subscription]{
		OutputState: i.ToSubscriptionOutputWithContext(ctx).OutputState,
	}
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*Subscription] {
	return pulumix.Output[*Subscription]{
		OutputState: o.OutputState,
	}
}

// Determines whether tracing is enabled
func (o SubscriptionOutput) AllowTracing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.AllowTracing }).(pulumi.BoolPtrOutput)
}

// Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o SubscriptionOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// The name of the subscription, or null if the subscription has no name.
func (o SubscriptionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o SubscriptionOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

// Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o SubscriptionOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o SubscriptionOutput) NotificationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.NotificationDate }).(pulumi.StringPtrOutput)
}

// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{userId} where {userId} is a user identifier.
func (o SubscriptionOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Subscription primary key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o SubscriptionOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// Scope like /products/{productId} or /apis or /apis/{apiId}.
func (o SubscriptionOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// Subscription secondary key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o SubscriptionOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o SubscriptionOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
func (o SubscriptionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Optional subscription comment added by an administrator when the state is changed to the 'rejected'.
func (o SubscriptionOutput) StateComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.StateComment }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
}

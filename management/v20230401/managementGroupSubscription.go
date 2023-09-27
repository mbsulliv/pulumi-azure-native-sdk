// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The details of subscription under management group.
type ManagementGroupSubscription struct {
	pulumi.CustomResourceState

	// The friendly name of the subscription.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the parent management group.
	Parent DescendantParentGroupInfoResponsePtrOutput `pulumi:"parent"`
	// The state of the subscription.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
	Tenant pulumi.StringPtrOutput `pulumi:"tenant"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementGroupSubscription registers a new resource with the given unique name, arguments, and options.
func NewManagementGroupSubscription(ctx *pulumi.Context,
	name string, args *ManagementGroupSubscriptionArgs, opts ...pulumi.ResourceOption) (*ManagementGroupSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:management:ManagementGroupSubscription"),
		},
		{
			Type: pulumi.String("azure-native:management/v20200501:ManagementGroupSubscription"),
		},
		{
			Type: pulumi.String("azure-native:management/v20201001:ManagementGroupSubscription"),
		},
		{
			Type: pulumi.String("azure-native:management/v20210401:ManagementGroupSubscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagementGroupSubscription
	err := ctx.RegisterResource("azure-native:management/v20230401:ManagementGroupSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementGroupSubscription gets an existing ManagementGroupSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementGroupSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupSubscriptionState, opts ...pulumi.ResourceOption) (*ManagementGroupSubscription, error) {
	var resource ManagementGroupSubscription
	err := ctx.ReadResource("azure-native:management/v20230401:ManagementGroupSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementGroupSubscription resources.
type managementGroupSubscriptionState struct {
}

type ManagementGroupSubscriptionState struct {
}

func (ManagementGroupSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupSubscriptionState)(nil)).Elem()
}

type managementGroupSubscriptionArgs struct {
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
	// Subscription ID.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a ManagementGroupSubscription resource.
type ManagementGroupSubscriptionArgs struct {
	// Management Group ID.
	GroupId pulumi.StringInput
	// Subscription ID.
	SubscriptionId pulumi.StringPtrInput
}

func (ManagementGroupSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupSubscriptionArgs)(nil)).Elem()
}

type ManagementGroupSubscriptionInput interface {
	pulumi.Input

	ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput
	ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput
}

func (*ManagementGroupSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupSubscription)(nil)).Elem()
}

func (i *ManagementGroupSubscription) ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput {
	return i.ToManagementGroupSubscriptionOutputWithContext(context.Background())
}

func (i *ManagementGroupSubscription) ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupSubscriptionOutput)
}

func (i *ManagementGroupSubscription) ToOutput(ctx context.Context) pulumix.Output[*ManagementGroupSubscription] {
	return pulumix.Output[*ManagementGroupSubscription]{
		OutputState: i.ToManagementGroupSubscriptionOutputWithContext(ctx).OutputState,
	}
}

type ManagementGroupSubscriptionOutput struct{ *pulumi.OutputState }

func (ManagementGroupSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupSubscription)(nil)).Elem()
}

func (o ManagementGroupSubscriptionOutput) ToManagementGroupSubscriptionOutput() ManagementGroupSubscriptionOutput {
	return o
}

func (o ManagementGroupSubscriptionOutput) ToManagementGroupSubscriptionOutputWithContext(ctx context.Context) ManagementGroupSubscriptionOutput {
	return o
}

func (o ManagementGroupSubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagementGroupSubscription] {
	return pulumix.Output[*ManagementGroupSubscription]{
		OutputState: o.OutputState,
	}
}

// The friendly name of the subscription.
func (o ManagementGroupSubscriptionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
func (o ManagementGroupSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the parent management group.
func (o ManagementGroupSubscriptionOutput) Parent() DescendantParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) DescendantParentGroupInfoResponsePtrOutput { return v.Parent }).(DescendantParentGroupInfoResponsePtrOutput)
}

// The state of the subscription.
func (o ManagementGroupSubscriptionOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
func (o ManagementGroupSubscriptionOutput) Tenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringPtrOutput { return v.Tenant }).(pulumi.StringPtrOutput)
}

// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
func (o ManagementGroupSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupSubscriptionOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Network Manager Connection resource
type SubscriptionNetworkManagerConnection struct {
	pulumi.CustomResourceState

	// A description of the network manager connection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network Manager Id.
	NetworkManagerId pulumi.StringPtrOutput `pulumi:"networkManagerId"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSubscriptionNetworkManagerConnection registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionNetworkManagerConnection(ctx *pulumi.Context,
	name string, args *SubscriptionNetworkManagerConnectionArgs, opts ...pulumi.ResourceOption) (*SubscriptionNetworkManagerConnection, error) {
	if args == nil {
		args = &SubscriptionNetworkManagerConnectionArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:SubscriptionNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:SubscriptionNetworkManagerConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SubscriptionNetworkManagerConnection
	err := ctx.RegisterResource("azure-native:network/v20230401:SubscriptionNetworkManagerConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionNetworkManagerConnection gets an existing SubscriptionNetworkManagerConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionNetworkManagerConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionNetworkManagerConnectionState, opts ...pulumi.ResourceOption) (*SubscriptionNetworkManagerConnection, error) {
	var resource SubscriptionNetworkManagerConnection
	err := ctx.ReadResource("azure-native:network/v20230401:SubscriptionNetworkManagerConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionNetworkManagerConnection resources.
type subscriptionNetworkManagerConnectionState struct {
}

type SubscriptionNetworkManagerConnectionState struct {
}

func (SubscriptionNetworkManagerConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionNetworkManagerConnectionState)(nil)).Elem()
}

type subscriptionNetworkManagerConnectionArgs struct {
	// A description of the network manager connection.
	Description *string `pulumi:"description"`
	// Name for the network manager connection.
	NetworkManagerConnectionName *string `pulumi:"networkManagerConnectionName"`
	// Network Manager Id.
	NetworkManagerId *string `pulumi:"networkManagerId"`
}

// The set of arguments for constructing a SubscriptionNetworkManagerConnection resource.
type SubscriptionNetworkManagerConnectionArgs struct {
	// A description of the network manager connection.
	Description pulumi.StringPtrInput
	// Name for the network manager connection.
	NetworkManagerConnectionName pulumi.StringPtrInput
	// Network Manager Id.
	NetworkManagerId pulumi.StringPtrInput
}

func (SubscriptionNetworkManagerConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionNetworkManagerConnectionArgs)(nil)).Elem()
}

type SubscriptionNetworkManagerConnectionInput interface {
	pulumi.Input

	ToSubscriptionNetworkManagerConnectionOutput() SubscriptionNetworkManagerConnectionOutput
	ToSubscriptionNetworkManagerConnectionOutputWithContext(ctx context.Context) SubscriptionNetworkManagerConnectionOutput
}

func (*SubscriptionNetworkManagerConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionNetworkManagerConnection)(nil)).Elem()
}

func (i *SubscriptionNetworkManagerConnection) ToSubscriptionNetworkManagerConnectionOutput() SubscriptionNetworkManagerConnectionOutput {
	return i.ToSubscriptionNetworkManagerConnectionOutputWithContext(context.Background())
}

func (i *SubscriptionNetworkManagerConnection) ToSubscriptionNetworkManagerConnectionOutputWithContext(ctx context.Context) SubscriptionNetworkManagerConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionNetworkManagerConnectionOutput)
}

func (i *SubscriptionNetworkManagerConnection) ToOutput(ctx context.Context) pulumix.Output[*SubscriptionNetworkManagerConnection] {
	return pulumix.Output[*SubscriptionNetworkManagerConnection]{
		OutputState: i.ToSubscriptionNetworkManagerConnectionOutputWithContext(ctx).OutputState,
	}
}

type SubscriptionNetworkManagerConnectionOutput struct{ *pulumi.OutputState }

func (SubscriptionNetworkManagerConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionNetworkManagerConnection)(nil)).Elem()
}

func (o SubscriptionNetworkManagerConnectionOutput) ToSubscriptionNetworkManagerConnectionOutput() SubscriptionNetworkManagerConnectionOutput {
	return o
}

func (o SubscriptionNetworkManagerConnectionOutput) ToSubscriptionNetworkManagerConnectionOutputWithContext(ctx context.Context) SubscriptionNetworkManagerConnectionOutput {
	return o
}

func (o SubscriptionNetworkManagerConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*SubscriptionNetworkManagerConnection] {
	return pulumix.Output[*SubscriptionNetworkManagerConnection]{
		OutputState: o.OutputState,
	}
}

// A description of the network manager connection.
func (o SubscriptionNetworkManagerConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o SubscriptionNetworkManagerConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource name.
func (o SubscriptionNetworkManagerConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network Manager Id.
func (o SubscriptionNetworkManagerConnectionOutput) NetworkManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringPtrOutput { return v.NetworkManagerId }).(pulumi.StringPtrOutput)
}

// The system metadata related to this resource.
func (o SubscriptionNetworkManagerConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o SubscriptionNetworkManagerConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionNetworkManagerConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionNetworkManagerConnectionOutput{})
}

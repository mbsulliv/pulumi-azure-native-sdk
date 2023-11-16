// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The routing intent child resource of a Virtual hub.
type RoutingIntent struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the RoutingIntent resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of routing policies.
	RoutingPolicies RoutingPolicyResponseArrayOutput `pulumi:"routingPolicies"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRoutingIntent registers a new resource with the given unique name, arguments, and options.
func NewRoutingIntent(ctx *pulumi.Context,
	name string, args *RoutingIntentArgs, opts ...pulumi.ResourceOption) (*RoutingIntent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:RoutingIntent"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:RoutingIntent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RoutingIntent
	err := ctx.RegisterResource("azure-native:network/v20230201:RoutingIntent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoutingIntent gets an existing RoutingIntent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoutingIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoutingIntentState, opts ...pulumi.ResourceOption) (*RoutingIntent, error) {
	var resource RoutingIntent
	err := ctx.ReadResource("azure-native:network/v20230201:RoutingIntent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoutingIntent resources.
type routingIntentState struct {
}

type RoutingIntentState struct {
}

func (RoutingIntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*routingIntentState)(nil)).Elem()
}

type routingIntentArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The resource group name of the RoutingIntent.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the per VirtualHub singleton Routing Intent resource.
	RoutingIntentName *string `pulumi:"routingIntentName"`
	// List of routing policies.
	RoutingPolicies []RoutingPolicy `pulumi:"routingPolicies"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The set of arguments for constructing a RoutingIntent resource.
type RoutingIntentArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The resource group name of the RoutingIntent.
	ResourceGroupName pulumi.StringInput
	// The name of the per VirtualHub singleton Routing Intent resource.
	RoutingIntentName pulumi.StringPtrInput
	// List of routing policies.
	RoutingPolicies RoutingPolicyArrayInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
}

func (RoutingIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routingIntentArgs)(nil)).Elem()
}

type RoutingIntentInput interface {
	pulumi.Input

	ToRoutingIntentOutput() RoutingIntentOutput
	ToRoutingIntentOutputWithContext(ctx context.Context) RoutingIntentOutput
}

func (*RoutingIntent) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingIntent)(nil)).Elem()
}

func (i *RoutingIntent) ToRoutingIntentOutput() RoutingIntentOutput {
	return i.ToRoutingIntentOutputWithContext(context.Background())
}

func (i *RoutingIntent) ToRoutingIntentOutputWithContext(ctx context.Context) RoutingIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingIntentOutput)
}

type RoutingIntentOutput struct{ *pulumi.OutputState }

func (RoutingIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingIntent)(nil)).Elem()
}

func (o RoutingIntentOutput) ToRoutingIntentOutput() RoutingIntentOutput {
	return o
}

func (o RoutingIntentOutput) ToRoutingIntentOutputWithContext(ctx context.Context) RoutingIntentOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o RoutingIntentOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingIntent) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o RoutingIntentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingIntent) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the RoutingIntent resource.
func (o RoutingIntentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingIntent) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of routing policies.
func (o RoutingIntentOutput) RoutingPolicies() RoutingPolicyResponseArrayOutput {
	return o.ApplyT(func(v *RoutingIntent) RoutingPolicyResponseArrayOutput { return v.RoutingPolicies }).(RoutingPolicyResponseArrayOutput)
}

// Resource type.
func (o RoutingIntentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RoutingIntent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RoutingIntentOutput{})
}

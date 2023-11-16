// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing Traffic Manager User Metrics.
// Azure REST API version: 2022-04-01. Prior API version in Azure Native 1.x: 2018-08-01.
//
// Other available API versions: 2017-09-01-preview, 2022-04-01-preview.
type TrafficManagerUserMetricsKey struct {
	pulumi.CustomResourceState

	// The key returned by the User Metrics operation.
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// The name of the resource
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewTrafficManagerUserMetricsKey registers a new resource with the given unique name, arguments, and options.
func NewTrafficManagerUserMetricsKey(ctx *pulumi.Context,
	name string, args *TrafficManagerUserMetricsKeyArgs, opts ...pulumi.ResourceOption) (*TrafficManagerUserMetricsKey, error) {
	if args == nil {
		args = &TrafficManagerUserMetricsKeyArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180401:TrafficManagerUserMetricsKey"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:TrafficManagerUserMetricsKey"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401:TrafficManagerUserMetricsKey"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:TrafficManagerUserMetricsKey"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TrafficManagerUserMetricsKey
	err := ctx.RegisterResource("azure-native:network:TrafficManagerUserMetricsKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficManagerUserMetricsKey gets an existing TrafficManagerUserMetricsKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficManagerUserMetricsKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficManagerUserMetricsKeyState, opts ...pulumi.ResourceOption) (*TrafficManagerUserMetricsKey, error) {
	var resource TrafficManagerUserMetricsKey
	err := ctx.ReadResource("azure-native:network:TrafficManagerUserMetricsKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficManagerUserMetricsKey resources.
type trafficManagerUserMetricsKeyState struct {
}

type TrafficManagerUserMetricsKeyState struct {
}

func (TrafficManagerUserMetricsKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerUserMetricsKeyState)(nil)).Elem()
}

type trafficManagerUserMetricsKeyArgs struct {
}

// The set of arguments for constructing a TrafficManagerUserMetricsKey resource.
type TrafficManagerUserMetricsKeyArgs struct {
}

func (TrafficManagerUserMetricsKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficManagerUserMetricsKeyArgs)(nil)).Elem()
}

type TrafficManagerUserMetricsKeyInput interface {
	pulumi.Input

	ToTrafficManagerUserMetricsKeyOutput() TrafficManagerUserMetricsKeyOutput
	ToTrafficManagerUserMetricsKeyOutputWithContext(ctx context.Context) TrafficManagerUserMetricsKeyOutput
}

func (*TrafficManagerUserMetricsKey) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficManagerUserMetricsKey)(nil)).Elem()
}

func (i *TrafficManagerUserMetricsKey) ToTrafficManagerUserMetricsKeyOutput() TrafficManagerUserMetricsKeyOutput {
	return i.ToTrafficManagerUserMetricsKeyOutputWithContext(context.Background())
}

func (i *TrafficManagerUserMetricsKey) ToTrafficManagerUserMetricsKeyOutputWithContext(ctx context.Context) TrafficManagerUserMetricsKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficManagerUserMetricsKeyOutput)
}

type TrafficManagerUserMetricsKeyOutput struct{ *pulumi.OutputState }

func (TrafficManagerUserMetricsKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficManagerUserMetricsKey)(nil)).Elem()
}

func (o TrafficManagerUserMetricsKeyOutput) ToTrafficManagerUserMetricsKeyOutput() TrafficManagerUserMetricsKeyOutput {
	return o
}

func (o TrafficManagerUserMetricsKeyOutput) ToTrafficManagerUserMetricsKeyOutputWithContext(ctx context.Context) TrafficManagerUserMetricsKeyOutput {
	return o
}

// The key returned by the User Metrics operation.
func (o TrafficManagerUserMetricsKeyOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficManagerUserMetricsKey) pulumi.StringPtrOutput { return v.Key }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o TrafficManagerUserMetricsKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficManagerUserMetricsKey) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
func (o TrafficManagerUserMetricsKeyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficManagerUserMetricsKey) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TrafficManagerUserMetricsKeyOutput{})
}

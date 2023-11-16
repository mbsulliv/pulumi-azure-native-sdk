// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package engagementfabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// EngagementFabric channel description
type ChannelTypeDescriptionResponse struct {
	// Text description for the channel
	ChannelDescription *string `pulumi:"channelDescription"`
	// All the available functions for the channel
	ChannelFunctions []string `pulumi:"channelFunctions"`
	// Channel type
	ChannelType *string `pulumi:"channelType"`
}

// EngagementFabric channel description
type ChannelTypeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ChannelTypeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelTypeDescriptionResponse)(nil)).Elem()
}

func (o ChannelTypeDescriptionResponseOutput) ToChannelTypeDescriptionResponseOutput() ChannelTypeDescriptionResponseOutput {
	return o
}

func (o ChannelTypeDescriptionResponseOutput) ToChannelTypeDescriptionResponseOutputWithContext(ctx context.Context) ChannelTypeDescriptionResponseOutput {
	return o
}

// Text description for the channel
func (o ChannelTypeDescriptionResponseOutput) ChannelDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelTypeDescriptionResponse) *string { return v.ChannelDescription }).(pulumi.StringPtrOutput)
}

// All the available functions for the channel
func (o ChannelTypeDescriptionResponseOutput) ChannelFunctions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ChannelTypeDescriptionResponse) []string { return v.ChannelFunctions }).(pulumi.StringArrayOutput)
}

// Channel type
func (o ChannelTypeDescriptionResponseOutput) ChannelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelTypeDescriptionResponse) *string { return v.ChannelType }).(pulumi.StringPtrOutput)
}

type ChannelTypeDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ChannelTypeDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelTypeDescriptionResponse)(nil)).Elem()
}

func (o ChannelTypeDescriptionResponseArrayOutput) ToChannelTypeDescriptionResponseArrayOutput() ChannelTypeDescriptionResponseArrayOutput {
	return o
}

func (o ChannelTypeDescriptionResponseArrayOutput) ToChannelTypeDescriptionResponseArrayOutputWithContext(ctx context.Context) ChannelTypeDescriptionResponseArrayOutput {
	return o
}

func (o ChannelTypeDescriptionResponseArrayOutput) Index(i pulumi.IntInput) ChannelTypeDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelTypeDescriptionResponse {
		return vs[0].([]ChannelTypeDescriptionResponse)[vs[1].(int)]
	}).(ChannelTypeDescriptionResponseOutput)
}

// The description of the EngagementFabric account key
type KeyDescriptionResponse struct {
	// The name of the key
	Name string `pulumi:"name"`
	// The rank of the key
	Rank string `pulumi:"rank"`
	// The value of the key
	Value string `pulumi:"value"`
}

// The description of the EngagementFabric account key
type KeyDescriptionResponseOutput struct{ *pulumi.OutputState }

func (KeyDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyDescriptionResponse)(nil)).Elem()
}

func (o KeyDescriptionResponseOutput) ToKeyDescriptionResponseOutput() KeyDescriptionResponseOutput {
	return o
}

func (o KeyDescriptionResponseOutput) ToKeyDescriptionResponseOutputWithContext(ctx context.Context) KeyDescriptionResponseOutput {
	return o
}

// The name of the key
func (o KeyDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KeyDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The rank of the key
func (o KeyDescriptionResponseOutput) Rank() pulumi.StringOutput {
	return o.ApplyT(func(v KeyDescriptionResponse) string { return v.Rank }).(pulumi.StringOutput)
}

// The value of the key
func (o KeyDescriptionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyDescriptionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type KeyDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyDescriptionResponse)(nil)).Elem()
}

func (o KeyDescriptionResponseArrayOutput) ToKeyDescriptionResponseArrayOutput() KeyDescriptionResponseArrayOutput {
	return o
}

func (o KeyDescriptionResponseArrayOutput) ToKeyDescriptionResponseArrayOutputWithContext(ctx context.Context) KeyDescriptionResponseArrayOutput {
	return o
}

func (o KeyDescriptionResponseArrayOutput) Index(i pulumi.IntInput) KeyDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyDescriptionResponse {
		return vs[0].([]KeyDescriptionResponse)[vs[1].(int)]
	}).(KeyDescriptionResponseOutput)
}

// The EngagementFabric SKU
type SKU struct {
	// The name of the SKU
	Name string `pulumi:"name"`
	// The price tier of the SKU
	Tier *string `pulumi:"tier"`
}

// SKUInput is an input type that accepts SKUArgs and SKUOutput values.
// You can construct a concrete instance of `SKUInput` via:
//
//	SKUArgs{...}
type SKUInput interface {
	pulumi.Input

	ToSKUOutput() SKUOutput
	ToSKUOutputWithContext(context.Context) SKUOutput
}

// The EngagementFabric SKU
type SKUArgs struct {
	// The name of the SKU
	Name pulumi.StringInput `pulumi:"name"`
	// The price tier of the SKU
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SKUArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SKU)(nil)).Elem()
}

func (i SKUArgs) ToSKUOutput() SKUOutput {
	return i.ToSKUOutputWithContext(context.Background())
}

func (i SKUArgs) ToSKUOutputWithContext(ctx context.Context) SKUOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SKUOutput)
}

// The EngagementFabric SKU
type SKUOutput struct{ *pulumi.OutputState }

func (SKUOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SKU)(nil)).Elem()
}

func (o SKUOutput) ToSKUOutput() SKUOutput {
	return o
}

func (o SKUOutput) ToSKUOutputWithContext(ctx context.Context) SKUOutput {
	return o
}

// The name of the SKU
func (o SKUOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SKU) string { return v.Name }).(pulumi.StringOutput)
}

// The price tier of the SKU
func (o SKUOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SKU) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

// The EngagementFabric SKU
type SKUResponse struct {
	// The name of the SKU
	Name string `pulumi:"name"`
	// The price tier of the SKU
	Tier *string `pulumi:"tier"`
}

// The EngagementFabric SKU
type SKUResponseOutput struct{ *pulumi.OutputState }

func (SKUResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SKUResponse)(nil)).Elem()
}

func (o SKUResponseOutput) ToSKUResponseOutput() SKUResponseOutput {
	return o
}

func (o SKUResponseOutput) ToSKUResponseOutputWithContext(ctx context.Context) SKUResponseOutput {
	return o
}

// The name of the SKU
func (o SKUResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SKUResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The price tier of the SKU
func (o SKUResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SKUResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelTypeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ChannelTypeDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyDescriptionResponseOutput{})
	pulumi.RegisterOutputType(KeyDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SKUOutput{})
	pulumi.RegisterOutputType(SKUResponseOutput{})
}

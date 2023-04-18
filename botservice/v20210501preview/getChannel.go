// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a BotService Channel registration specified by the parameters.
func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	var rv LookupChannelResult
	err := ctx.Invoke("azure-native:botservice/v20210501preview:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	// The name of the Bot resource.
	ChannelName string `pulumi:"channelName"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
}

// Bot channel resource definition
type LookupChannelResult struct {
	// Entity Tag.
	Etag *string `pulumi:"etag"`
	// Specifies the resource ID.
	Id string `pulumi:"id"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// The set of properties specific to bot channel resource
	Properties interface{} `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
	// Entity zones
	Zones []string `pulumi:"zones"`
}

func LookupChannelOutput(ctx *pulumi.Context, args LookupChannelOutputArgs, opts ...pulumi.InvokeOption) LookupChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelResult, error) {
			args := v.(LookupChannelArgs)
			r, err := LookupChannel(ctx, &args, opts...)
			var s LookupChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelResultOutput)
}

type LookupChannelOutputArgs struct {
	// The name of the Bot resource.
	ChannelName pulumi.StringInput `pulumi:"channelName"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelArgs)(nil)).Elem()
}

// Bot channel resource definition
type LookupChannelResultOutput struct{ *pulumi.OutputState }

func (LookupChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelResult)(nil)).Elem()
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutput() LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutputWithContext(ctx context.Context) LookupChannelResultOutput {
	return o
}

// Entity Tag.
func (o LookupChannelResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Specifies the resource ID.
func (o LookupChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

// Required. Gets or sets the Kind of the resource.
func (o LookupChannelResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Specifies the location of the resource.
func (o LookupChannelResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o LookupChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to bot channel resource
func (o LookupChannelResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupChannelResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Gets or sets the SKU of the resource.
func (o LookupChannelResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o LookupChannelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupChannelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o LookupChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

// Entity zones
func (o LookupChannelResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelResultOutput{})
}

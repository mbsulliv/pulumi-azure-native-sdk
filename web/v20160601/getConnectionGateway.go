// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a specific gateway under a subscription and in a specific resource group
func LookupConnectionGateway(ctx *pulumi.Context, args *LookupConnectionGatewayArgs, opts ...pulumi.InvokeOption) (*LookupConnectionGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectionGatewayResult
	err := ctx.Invoke("azure-native:web/v20160601:getConnectionGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionGatewayArgs struct {
	// The connection gateway name
	ConnectionGatewayName string `pulumi:"connectionGatewayName"`
	// The resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// The gateway definition
type LookupConnectionGatewayResult struct {
	// Resource ETag
	Etag *string `pulumi:"etag"`
	// Resource id
	Id string `pulumi:"id"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name       string                                        `pulumi:"name"`
	Properties ConnectionGatewayDefinitionResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupConnectionGatewayOutput(ctx *pulumi.Context, args LookupConnectionGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionGatewayResult, error) {
			args := v.(LookupConnectionGatewayArgs)
			r, err := LookupConnectionGateway(ctx, &args, opts...)
			var s LookupConnectionGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionGatewayResultOutput)
}

type LookupConnectionGatewayOutputArgs struct {
	// The connection gateway name
	ConnectionGatewayName pulumi.StringInput `pulumi:"connectionGatewayName"`
	// The resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupConnectionGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionGatewayArgs)(nil)).Elem()
}

// The gateway definition
type LookupConnectionGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionGatewayResult)(nil)).Elem()
}

func (o LookupConnectionGatewayResultOutput) ToLookupConnectionGatewayResultOutput() LookupConnectionGatewayResultOutput {
	return o
}

func (o LookupConnectionGatewayResultOutput) ToLookupConnectionGatewayResultOutputWithContext(ctx context.Context) LookupConnectionGatewayResultOutput {
	return o
}

func (o LookupConnectionGatewayResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConnectionGatewayResult] {
	return pulumix.Output[LookupConnectionGatewayResult]{
		OutputState: o.OutputState,
	}
}

// Resource ETag
func (o LookupConnectionGatewayResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource id
func (o LookupConnectionGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupConnectionGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupConnectionGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectionGatewayResultOutput) Properties() ConnectionGatewayDefinitionResponsePropertiesOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) ConnectionGatewayDefinitionResponseProperties {
		return v.Properties
	}).(ConnectionGatewayDefinitionResponsePropertiesOutput)
}

// Resource tags
func (o LookupConnectionGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupConnectionGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionGatewayResultOutput{})
}

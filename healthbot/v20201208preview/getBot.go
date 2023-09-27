// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201208preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Healthcare Bot.
func LookupBot(ctx *pulumi.Context, args *LookupBotArgs, opts ...pulumi.InvokeOption) (*LookupBotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBotResult
	err := ctx.Invoke("azure-native:healthbot/v20201208preview:getBot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBotArgs struct {
	// The name of the Bot resource.
	BotName string `pulumi:"botName"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// HealthBot resource definition
type LookupBotResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The set of properties specific to healthcare bot resource.
	Properties HealthBotPropertiesResponse `pulumi:"properties"`
	// SKU of the HealthBot.
	Sku *SkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupBotOutput(ctx *pulumi.Context, args LookupBotOutputArgs, opts ...pulumi.InvokeOption) LookupBotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBotResult, error) {
			args := v.(LookupBotArgs)
			r, err := LookupBot(ctx, &args, opts...)
			var s LookupBotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBotResultOutput)
}

type LookupBotOutputArgs struct {
	// The name of the Bot resource.
	BotName pulumi.StringInput `pulumi:"botName"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotArgs)(nil)).Elem()
}

// HealthBot resource definition
type LookupBotResultOutput struct{ *pulumi.OutputState }

func (LookupBotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotResult)(nil)).Elem()
}

func (o LookupBotResultOutput) ToLookupBotResultOutput() LookupBotResultOutput {
	return o
}

func (o LookupBotResultOutput) ToLookupBotResultOutputWithContext(ctx context.Context) LookupBotResultOutput {
	return o
}

func (o LookupBotResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBotResult] {
	return pulumix.Output[LookupBotResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupBotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupBotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to healthcare bot resource.
func (o LookupBotResultOutput) Properties() HealthBotPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBotResult) HealthBotPropertiesResponse { return v.Properties }).(HealthBotPropertiesResponseOutput)
}

// SKU of the HealthBot.
func (o LookupBotResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBotResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o LookupBotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupBotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupBotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBotResultOutput{})
}

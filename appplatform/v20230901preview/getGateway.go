// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the Spring Cloud Gateway and its properties.
func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayResult
	err := ctx.Invoke("azure-native:appplatform/v20230901preview:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGatewayArgs struct {
	// The name of Spring Cloud Gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Spring Cloud Gateway resource
type LookupGatewayResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Spring Cloud Gateway properties payload
	Properties GatewayPropertiesResponse `pulumi:"properties"`
	// Sku of the Spring Cloud Gateway resource
	Sku *SkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupGatewayResult
func (val *LookupGatewayResult) Defaults() *LookupGatewayResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayResult, error) {
			args := v.(LookupGatewayArgs)
			r, err := LookupGateway(ctx, &args, opts...)
			var s LookupGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayResultOutput)
}

type LookupGatewayOutputArgs struct {
	// The name of Spring Cloud Gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}

// Spring Cloud Gateway resource
type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGatewayResult] {
	return pulumix.Output[LookupGatewayResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// Spring Cloud Gateway properties payload
func (o LookupGatewayResultOutput) Properties() GatewayPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) GatewayPropertiesResponse { return v.Properties }).(GatewayPropertiesResponseOutput)
}

// Sku of the Spring Cloud Gateway resource
func (o LookupGatewayResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupGatewayResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}

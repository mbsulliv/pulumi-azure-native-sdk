// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the Spring Cloud Gateway route configs.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview.
func LookupGatewayRouteConfig(ctx *pulumi.Context, args *LookupGatewayRouteConfigArgs, opts ...pulumi.InvokeOption) (*LookupGatewayRouteConfigResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayRouteConfigResult
	err := ctx.Invoke("azure-native:appplatform:getGatewayRouteConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGatewayRouteConfigArgs struct {
	// The name of Spring Cloud Gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Spring Cloud Gateway route config.
	RouteConfigName string `pulumi:"routeConfigName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Spring Cloud Gateway route config resource
type LookupGatewayRouteConfigResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// API route config of the Spring Cloud Gateway
	Properties GatewayRouteConfigPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupGatewayRouteConfigResult
func (val *LookupGatewayRouteConfigResult) Defaults() *LookupGatewayRouteConfigResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupGatewayRouteConfigOutput(ctx *pulumi.Context, args LookupGatewayRouteConfigOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayRouteConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayRouteConfigResult, error) {
			args := v.(LookupGatewayRouteConfigArgs)
			r, err := LookupGatewayRouteConfig(ctx, &args, opts...)
			var s LookupGatewayRouteConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayRouteConfigResultOutput)
}

type LookupGatewayRouteConfigOutputArgs struct {
	// The name of Spring Cloud Gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Spring Cloud Gateway route config.
	RouteConfigName pulumi.StringInput `pulumi:"routeConfigName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayRouteConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayRouteConfigArgs)(nil)).Elem()
}

// Spring Cloud Gateway route config resource
type LookupGatewayRouteConfigResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayRouteConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayRouteConfigResult)(nil)).Elem()
}

func (o LookupGatewayRouteConfigResultOutput) ToLookupGatewayRouteConfigResultOutput() LookupGatewayRouteConfigResultOutput {
	return o
}

func (o LookupGatewayRouteConfigResultOutput) ToLookupGatewayRouteConfigResultOutputWithContext(ctx context.Context) LookupGatewayRouteConfigResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupGatewayRouteConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupGatewayRouteConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// API route config of the Spring Cloud Gateway
func (o LookupGatewayRouteConfigResultOutput) Properties() GatewayRouteConfigPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) GatewayRouteConfigPropertiesResponse { return v.Properties }).(GatewayRouteConfigPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupGatewayRouteConfigResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupGatewayRouteConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayRouteConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayRouteConfigResultOutput{})
}

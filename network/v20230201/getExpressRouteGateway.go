// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fetches the details of a ExpressRoute gateway in a resource group.
func LookupExpressRouteGateway(ctx *pulumi.Context, args *LookupExpressRouteGatewayArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExpressRouteGatewayResult
	err := ctx.Invoke("azure-native:network/v20230201:getExpressRouteGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteGatewayArgs struct {
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName string `pulumi:"expressRouteGatewayName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ExpressRoute gateway resource.
type LookupExpressRouteGatewayResult struct {
	// Configures this gateway to accept traffic from non Virtual WAN networks.
	AllowNonVirtualWanTraffic *bool `pulumi:"allowNonVirtualWanTraffic"`
	// Configuration for auto scaling.
	AutoScaleConfiguration *ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration `pulumi:"autoScaleConfiguration"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// List of ExpressRoute connections to the ExpressRoute gateway.
	ExpressRouteConnections []ExpressRouteConnectionResponse `pulumi:"expressRouteConnections"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the express route gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
	VirtualHub VirtualHubIdResponse `pulumi:"virtualHub"`
}

func LookupExpressRouteGatewayOutput(ctx *pulumi.Context, args LookupExpressRouteGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteGatewayResult, error) {
			args := v.(LookupExpressRouteGatewayArgs)
			r, err := LookupExpressRouteGateway(ctx, &args, opts...)
			var s LookupExpressRouteGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteGatewayResultOutput)
}

type LookupExpressRouteGatewayOutputArgs struct {
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName pulumi.StringInput `pulumi:"expressRouteGatewayName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteGatewayArgs)(nil)).Elem()
}

// ExpressRoute gateway resource.
type LookupExpressRouteGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteGatewayResult)(nil)).Elem()
}

func (o LookupExpressRouteGatewayResultOutput) ToLookupExpressRouteGatewayResultOutput() LookupExpressRouteGatewayResultOutput {
	return o
}

func (o LookupExpressRouteGatewayResultOutput) ToLookupExpressRouteGatewayResultOutputWithContext(ctx context.Context) LookupExpressRouteGatewayResultOutput {
	return o
}

// Configures this gateway to accept traffic from non Virtual WAN networks.
func (o LookupExpressRouteGatewayResultOutput) AllowNonVirtualWanTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) *bool { return v.AllowNonVirtualWanTraffic }).(pulumi.BoolPtrOutput)
}

// Configuration for auto scaling.
func (o LookupExpressRouteGatewayResultOutput) AutoScaleConfiguration() ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) *ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration {
		return v.AutoScaleConfiguration
	}).(ExpressRouteGatewayPropertiesResponseAutoScaleConfigurationPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupExpressRouteGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

// List of ExpressRoute connections to the ExpressRoute gateway.
func (o LookupExpressRouteGatewayResultOutput) ExpressRouteConnections() ExpressRouteConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) []ExpressRouteConnectionResponse {
		return v.ExpressRouteConnections
	}).(ExpressRouteConnectionResponseArrayOutput)
}

// Resource ID.
func (o LookupExpressRouteGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupExpressRouteGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupExpressRouteGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the express route gateway resource.
func (o LookupExpressRouteGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupExpressRouteGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupExpressRouteGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

// The Virtual Hub where the ExpressRoute gateway is or will be deployed.
func (o LookupExpressRouteGatewayResultOutput) VirtualHub() VirtualHubIdResponseOutput {
	return o.ApplyT(func(v LookupExpressRouteGatewayResult) VirtualHubIdResponse { return v.VirtualHub }).(VirtualHubIdResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteGatewayResultOutput{})
}

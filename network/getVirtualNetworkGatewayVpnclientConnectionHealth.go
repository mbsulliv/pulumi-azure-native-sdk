// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get VPN client connection health detail per P2S client connection of the virtual network gateway in the specified resource group.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2019-08-01, 2023-04-01, 2023-05-01.
func GetVirtualNetworkGatewayVpnclientConnectionHealth(ctx *pulumi.Context, args *GetVirtualNetworkGatewayVpnclientConnectionHealthArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayVpnclientConnectionHealthResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualNetworkGatewayVpnclientConnectionHealthResult
	err := ctx.Invoke("azure-native:network:getVirtualNetworkGatewayVpnclientConnectionHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayVpnclientConnectionHealthArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}

// List of virtual network gateway vpn client connection health.
type GetVirtualNetworkGatewayVpnclientConnectionHealthResult struct {
	// List of vpn client connection health.
	Value []VpnClientConnectionHealthDetailResponse `pulumi:"value"`
}

func GetVirtualNetworkGatewayVpnclientConnectionHealthOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayVpnclientConnectionHealthOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayVpnclientConnectionHealthResult, error) {
			args := v.(GetVirtualNetworkGatewayVpnclientConnectionHealthArgs)
			r, err := GetVirtualNetworkGatewayVpnclientConnectionHealth(ctx, &args, opts...)
			var s GetVirtualNetworkGatewayVpnclientConnectionHealthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput)
}

type GetVirtualNetworkGatewayVpnclientConnectionHealthOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayVpnclientConnectionHealthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnclientConnectionHealthArgs)(nil)).Elem()
}

// List of virtual network gateway vpn client connection health.
type GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnclientConnectionHealthResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) ToGetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput() GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) ToGetVirtualNetworkGatewayVpnclientConnectionHealthResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetVirtualNetworkGatewayVpnclientConnectionHealthResult] {
	return pulumix.Output[GetVirtualNetworkGatewayVpnclientConnectionHealthResult]{
		OutputState: o.OutputState,
	}
}

// List of vpn client connection health.
func (o GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) Value() VpnClientConnectionHealthDetailResponseArrayOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientConnectionHealthResult) []VpnClientConnectionHealthDetailResponse {
		return v.Value
	}).(VpnClientConnectionHealthDetailResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput{})
}

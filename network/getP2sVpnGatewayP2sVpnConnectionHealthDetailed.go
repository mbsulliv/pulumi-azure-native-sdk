// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01.
func GetP2sVpnGatewayP2sVpnConnectionHealthDetailed(ctx *pulumi.Context, args *GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs, opts ...pulumi.InvokeOption) (*GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult
	err := ctx.Invoke("azure-native:network:getP2sVpnGatewayP2sVpnConnectionHealthDetailed", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs struct {
	// The name of the P2SVpnGateway.
	GatewayName string `pulumi:"gatewayName"`
	// The sas-url to download the P2S Vpn connection health detail.
	OutputBlobSasUrl *string `pulumi:"outputBlobSasUrl"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of p2s vpn user names whose p2s vpn connection detailed health to retrieve for.
	VpnUserNamesFilter []string `pulumi:"vpnUserNamesFilter"`
}

// P2S Vpn connection detailed health written to sas url.
type GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult struct {
	// Returned sas url of the blob to which the p2s vpn connection detailed health will be written.
	SasUrl *string `pulumi:"sasUrl"`
}

func GetP2sVpnGatewayP2sVpnConnectionHealthDetailedOutput(ctx *pulumi.Context, args GetP2sVpnGatewayP2sVpnConnectionHealthDetailedOutputArgs, opts ...pulumi.InvokeOption) GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult, error) {
			args := v.(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs)
			r, err := GetP2sVpnGatewayP2sVpnConnectionHealthDetailed(ctx, &args, opts...)
			var s GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput)
}

type GetP2sVpnGatewayP2sVpnConnectionHealthDetailedOutputArgs struct {
	// The name of the P2SVpnGateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The sas-url to download the P2S Vpn connection health detail.
	OutputBlobSasUrl pulumi.StringPtrInput `pulumi:"outputBlobSasUrl"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The list of p2s vpn user names whose p2s vpn connection detailed health to retrieve for.
	VpnUserNamesFilter pulumi.StringArrayInput `pulumi:"vpnUserNamesFilter"`
}

func (GetP2sVpnGatewayP2sVpnConnectionHealthDetailedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs)(nil)).Elem()
}

// P2S Vpn connection detailed health written to sas url.
type GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput struct{ *pulumi.OutputState }

func (GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult)(nil)).Elem()
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput) ToGetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput() GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput {
	return o
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput) ToGetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutputWithContext(ctx context.Context) GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput {
	return o
}

// Returned sas url of the blob to which the p2s vpn connection detailed health will be written.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput{})
}

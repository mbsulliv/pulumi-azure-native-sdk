// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the details of a Virtual Hub Bgp Connection.
func LookupVirtualHubBgpConnection(ctx *pulumi.Context, args *LookupVirtualHubBgpConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubBgpConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualHubBgpConnectionResult
	err := ctx.Invoke("azure-native:network/v20230401:getVirtualHubBgpConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubBgpConnectionArgs struct {
	// The name of the connection.
	ConnectionName string `pulumi:"connectionName"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// Virtual Appliance Site resource.
type LookupVirtualHubBgpConnectionResult struct {
	// The current state of the VirtualHub to Peer.
	ConnectionState string `pulumi:"connectionState"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The reference to the HubVirtualNetworkConnection resource.
	HubVirtualNetworkConnection *SubResourceResponse `pulumi:"hubVirtualNetworkConnection"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the connection.
	Name *string `pulumi:"name"`
	// Peer ASN.
	PeerAsn *float64 `pulumi:"peerAsn"`
	// Peer IP.
	PeerIp *string `pulumi:"peerIp"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Connection type.
	Type string `pulumi:"type"`
}

func LookupVirtualHubBgpConnectionOutput(ctx *pulumi.Context, args LookupVirtualHubBgpConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubBgpConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubBgpConnectionResult, error) {
			args := v.(LookupVirtualHubBgpConnectionArgs)
			r, err := LookupVirtualHubBgpConnection(ctx, &args, opts...)
			var s LookupVirtualHubBgpConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubBgpConnectionResultOutput)
}

type LookupVirtualHubBgpConnectionOutputArgs struct {
	// The name of the connection.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubBgpConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubBgpConnectionArgs)(nil)).Elem()
}

// Virtual Appliance Site resource.
type LookupVirtualHubBgpConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubBgpConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubBgpConnectionResult)(nil)).Elem()
}

func (o LookupVirtualHubBgpConnectionResultOutput) ToLookupVirtualHubBgpConnectionResultOutput() LookupVirtualHubBgpConnectionResultOutput {
	return o
}

func (o LookupVirtualHubBgpConnectionResultOutput) ToLookupVirtualHubBgpConnectionResultOutputWithContext(ctx context.Context) LookupVirtualHubBgpConnectionResultOutput {
	return o
}

func (o LookupVirtualHubBgpConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualHubBgpConnectionResult] {
	return pulumix.Output[LookupVirtualHubBgpConnectionResult]{
		OutputState: o.OutputState,
	}
}

// The current state of the VirtualHub to Peer.
func (o LookupVirtualHubBgpConnectionResultOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) string { return v.ConnectionState }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualHubBgpConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The reference to the HubVirtualNetworkConnection resource.
func (o LookupVirtualHubBgpConnectionResultOutput) HubVirtualNetworkConnection() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *SubResourceResponse { return v.HubVirtualNetworkConnection }).(SubResourceResponsePtrOutput)
}

// Resource ID.
func (o LookupVirtualHubBgpConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the connection.
func (o LookupVirtualHubBgpConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Peer ASN.
func (o LookupVirtualHubBgpConnectionResultOutput) PeerAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *float64 { return v.PeerAsn }).(pulumi.Float64PtrOutput)
}

// Peer IP.
func (o LookupVirtualHubBgpConnectionResultOutput) PeerIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *string { return v.PeerIp }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o LookupVirtualHubBgpConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Connection type.
func (o LookupVirtualHubBgpConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubBgpConnectionResultOutput{})
}

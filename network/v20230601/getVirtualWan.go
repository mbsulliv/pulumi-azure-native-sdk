// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a VirtualWAN.
func LookupVirtualWan(ctx *pulumi.Context, args *LookupVirtualWanArgs, opts ...pulumi.InvokeOption) (*LookupVirtualWanResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualWanResult
	err := ctx.Invoke("azure-native:network/v20230601:getVirtualWan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualWanArgs struct {
	// The resource group name of the VirtualWan.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualWAN being retrieved.
	VirtualWANName string `pulumi:"virtualWANName"`
}

// VirtualWAN Resource.
type LookupVirtualWanResult struct {
	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic *bool `pulumi:"allowVnetToVnetTraffic"`
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The office local breakout category.
	Office365LocalBreakoutCategory string `pulumi:"office365LocalBreakoutCategory"`
	// The provisioning state of the virtual WAN resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// List of VirtualHubs in the VirtualWAN.
	VirtualHubs []SubResourceResponse `pulumi:"virtualHubs"`
	// List of VpnSites in the VirtualWAN.
	VpnSites []SubResourceResponse `pulumi:"vpnSites"`
}

func LookupVirtualWanOutput(ctx *pulumi.Context, args LookupVirtualWanOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualWanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualWanResult, error) {
			args := v.(LookupVirtualWanArgs)
			r, err := LookupVirtualWan(ctx, &args, opts...)
			var s LookupVirtualWanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualWanResultOutput)
}

type LookupVirtualWanOutputArgs struct {
	// The resource group name of the VirtualWan.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VirtualWAN being retrieved.
	VirtualWANName pulumi.StringInput `pulumi:"virtualWANName"`
}

func (LookupVirtualWanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWanArgs)(nil)).Elem()
}

// VirtualWAN Resource.
type LookupVirtualWanResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualWanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWanResult)(nil)).Elem()
}

func (o LookupVirtualWanResultOutput) ToLookupVirtualWanResultOutput() LookupVirtualWanResultOutput {
	return o
}

func (o LookupVirtualWanResultOutput) ToLookupVirtualWanResultOutputWithContext(ctx context.Context) LookupVirtualWanResultOutput {
	return o
}

// True if branch to branch traffic is allowed.
func (o LookupVirtualWanResultOutput) AllowBranchToBranchTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) *bool { return v.AllowBranchToBranchTraffic }).(pulumi.BoolPtrOutput)
}

// True if Vnet to Vnet traffic is allowed.
func (o LookupVirtualWanResultOutput) AllowVnetToVnetTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) *bool { return v.AllowVnetToVnetTraffic }).(pulumi.BoolPtrOutput)
}

// Vpn encryption to be disabled or not.
func (o LookupVirtualWanResultOutput) DisableVpnEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) *bool { return v.DisableVpnEncryption }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualWanResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVirtualWanResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupVirtualWanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupVirtualWanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Name }).(pulumi.StringOutput)
}

// The office local breakout category.
func (o LookupVirtualWanResultOutput) Office365LocalBreakoutCategory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Office365LocalBreakoutCategory }).(pulumi.StringOutput)
}

// The provisioning state of the virtual WAN resource.
func (o LookupVirtualWanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupVirtualWanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupVirtualWanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Type }).(pulumi.StringOutput)
}

// List of VirtualHubs in the VirtualWAN.
func (o LookupVirtualWanResultOutput) VirtualHubs() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) []SubResourceResponse { return v.VirtualHubs }).(SubResourceResponseArrayOutput)
}

// List of VpnSites in the VirtualWAN.
func (o LookupVirtualWanResultOutput) VpnSites() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) []SubResourceResponse { return v.VpnSites }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualWanResultOutput{})
}
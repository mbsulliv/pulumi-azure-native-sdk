// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Gets a virtual network the app (or deployment slot) is connected to by name.
func LookupWebAppVnetConnectionSlot(ctx *pulumi.Context, args *LookupWebAppVnetConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppVnetConnectionSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppVnetConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20220901:getWebAppVnetConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppVnetConnectionSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the named virtual network for the production slot.
	Slot string `pulumi:"slot"`
	// Name of the virtual network.
	VnetName string `pulumi:"vnetName"`
}

// Virtual Network information ARM resource.
type LookupWebAppVnetConnectionSlotResult struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob *string `pulumi:"certBlob"`
	// The client certificate thumbprint.
	CertThumbprint string `pulumi:"certThumbprint"`
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers *string `pulumi:"dnsServers"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Flag that is used to denote if this is VNET injection
	IsSwift *bool `pulumi:"isSwift"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// <code>true</code> if a resync is required; otherwise, <code>false</code>.
	ResyncRequired bool `pulumi:"resyncRequired"`
	// The routes that this Virtual Network connection uses.
	Routes []VnetRouteResponse `pulumi:"routes"`
	// Resource type.
	Type string `pulumi:"type"`
	// The Virtual Network's resource ID.
	VnetResourceId *string `pulumi:"vnetResourceId"`
}

func LookupWebAppVnetConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppVnetConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppVnetConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppVnetConnectionSlotResult, error) {
			args := v.(LookupWebAppVnetConnectionSlotArgs)
			r, err := LookupWebAppVnetConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppVnetConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppVnetConnectionSlotResultOutput)
}

type LookupWebAppVnetConnectionSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the named virtual network for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
	// Name of the virtual network.
	VnetName pulumi.StringInput `pulumi:"vnetName"`
}

func (LookupWebAppVnetConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppVnetConnectionSlotArgs)(nil)).Elem()
}

// Virtual Network information ARM resource.
type LookupWebAppVnetConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppVnetConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppVnetConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppVnetConnectionSlotResultOutput) ToLookupWebAppVnetConnectionSlotResultOutput() LookupWebAppVnetConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppVnetConnectionSlotResultOutput) ToLookupWebAppVnetConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppVnetConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppVnetConnectionSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppVnetConnectionSlotResult] {
	return pulumix.Output[LookupWebAppVnetConnectionSlotResult]{
		OutputState: o.OutputState,
	}
}

// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
// Point-To-Site VPN connection.
func (o LookupWebAppVnetConnectionSlotResultOutput) CertBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *string { return v.CertBlob }).(pulumi.StringPtrOutput)
}

// The client certificate thumbprint.
func (o LookupWebAppVnetConnectionSlotResultOutput) CertThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) string { return v.CertThumbprint }).(pulumi.StringOutput)
}

// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
func (o LookupWebAppVnetConnectionSlotResultOutput) DnsServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *string { return v.DnsServers }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppVnetConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Flag that is used to denote if this is VNET injection
func (o LookupWebAppVnetConnectionSlotResultOutput) IsSwift() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *bool { return v.IsSwift }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o LookupWebAppVnetConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppVnetConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// <code>true</code> if a resync is required; otherwise, <code>false</code>.
func (o LookupWebAppVnetConnectionSlotResultOutput) ResyncRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) bool { return v.ResyncRequired }).(pulumi.BoolOutput)
}

// The routes that this Virtual Network connection uses.
func (o LookupWebAppVnetConnectionSlotResultOutput) Routes() VnetRouteResponseArrayOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) []VnetRouteResponse { return v.Routes }).(VnetRouteResponseArrayOutput)
}

// Resource type.
func (o LookupWebAppVnetConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

// The Virtual Network's resource ID.
func (o LookupWebAppVnetConnectionSlotResultOutput) VnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *string { return v.VnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppVnetConnectionSlotResultOutput{})
}

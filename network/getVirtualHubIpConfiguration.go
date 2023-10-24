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

// Retrieves the details of a Virtual Hub Ip configuration.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01.
func LookupVirtualHubIpConfiguration(ctx *pulumi.Context, args *LookupVirtualHubIpConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubIpConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualHubIpConfigurationResult
	err := ctx.Invoke("azure-native:network:getVirtualHubIpConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualHubIpConfigurationArgs struct {
	// The name of the ipconfig.
	IpConfigName string `pulumi:"ipConfigName"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// IpConfigurations.
type LookupVirtualHubIpConfigurationResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the Ip Configuration.
	Name *string `pulumi:"name"`
	// The private IP address of the IP configuration.
	PrivateIPAddress *string `pulumi:"privateIPAddress"`
	// The private IP address allocation method.
	PrivateIPAllocationMethod *string `pulumi:"privateIPAllocationMethod"`
	// The provisioning state of the IP configuration resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The reference to the public IP resource.
	PublicIPAddress *PublicIPAddressResponse `pulumi:"publicIPAddress"`
	// The reference to the subnet resource.
	Subnet *SubnetResponse `pulumi:"subnet"`
	// Ipconfiguration type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupVirtualHubIpConfigurationResult
func (val *LookupVirtualHubIpConfigurationResult) Defaults() *LookupVirtualHubIpConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PublicIPAddress = tmp.PublicIPAddress.Defaults()

	tmp.Subnet = tmp.Subnet.Defaults()

	return &tmp
}

func LookupVirtualHubIpConfigurationOutput(ctx *pulumi.Context, args LookupVirtualHubIpConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubIpConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubIpConfigurationResult, error) {
			args := v.(LookupVirtualHubIpConfigurationArgs)
			r, err := LookupVirtualHubIpConfiguration(ctx, &args, opts...)
			var s LookupVirtualHubIpConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubIpConfigurationResultOutput)
}

type LookupVirtualHubIpConfigurationOutputArgs struct {
	// The name of the ipconfig.
	IpConfigName pulumi.StringInput `pulumi:"ipConfigName"`
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubIpConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubIpConfigurationArgs)(nil)).Elem()
}

// IpConfigurations.
type LookupVirtualHubIpConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubIpConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubIpConfigurationResult)(nil)).Elem()
}

func (o LookupVirtualHubIpConfigurationResultOutput) ToLookupVirtualHubIpConfigurationResultOutput() LookupVirtualHubIpConfigurationResultOutput {
	return o
}

func (o LookupVirtualHubIpConfigurationResultOutput) ToLookupVirtualHubIpConfigurationResultOutputWithContext(ctx context.Context) LookupVirtualHubIpConfigurationResultOutput {
	return o
}

func (o LookupVirtualHubIpConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualHubIpConfigurationResult] {
	return pulumix.Output[LookupVirtualHubIpConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualHubIpConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVirtualHubIpConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the Ip Configuration.
func (o LookupVirtualHubIpConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The private IP address of the IP configuration.
func (o LookupVirtualHubIpConfigurationResultOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

// The private IP address allocation method.
func (o LookupVirtualHubIpConfigurationResultOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

// The provisioning state of the IP configuration resource.
func (o LookupVirtualHubIpConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The reference to the public IP resource.
func (o LookupVirtualHubIpConfigurationResultOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *PublicIPAddressResponse { return v.PublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

// The reference to the subnet resource.
func (o LookupVirtualHubIpConfigurationResultOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

// Ipconfiguration type.
func (o LookupVirtualHubIpConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubIpConfigurationResultOutput{})
}

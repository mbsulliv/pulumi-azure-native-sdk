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

// Gets the specified public IP address in a specified resource group.
func LookupPublicIPAddress(ctx *pulumi.Context, args *LookupPublicIPAddressArgs, opts ...pulumi.InvokeOption) (*LookupPublicIPAddressResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPublicIPAddressResult
	err := ctx.Invoke("azure-native:network/v20230401:getPublicIPAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPublicIPAddressArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the public IP address.
	PublicIpAddressName string `pulumi:"publicIpAddressName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Public IP address resource.
type LookupPublicIPAddressResult struct {
	// The DDoS protection custom policy associated with the public IP address.
	DdosSettings *DdosSettingsResponse `pulumi:"ddosSettings"`
	// Specify what happens to the public IP address when the VM using it is deleted
	DeleteOption *string `pulumi:"deleteOption"`
	// The FQDN of the DNS record associated with the public IP address.
	DnsSettings *PublicIPAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The extended location of the public ip address.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The IP address associated with the public IP address resource.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP configuration associated with the public IP address.
	IpConfiguration IPConfigurationResponse `pulumi:"ipConfiguration"`
	// The list of tags associated with the public IP address.
	IpTags []IpTagResponse `pulumi:"ipTags"`
	// The linked public IP address of the public IP address resource.
	LinkedPublicIPAddress *PublicIPAddressResponse `pulumi:"linkedPublicIPAddress"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Migration phase of Public IP Address.
	MigrationPhase *string `pulumi:"migrationPhase"`
	// Resource name.
	Name string `pulumi:"name"`
	// The NatGateway for the Public IP address.
	NatGateway *NatGatewayResponse `pulumi:"natGateway"`
	// The provisioning state of the public IP address resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The public IP address version.
	PublicIPAddressVersion *string `pulumi:"publicIPAddressVersion"`
	// The public IP address allocation method.
	PublicIPAllocationMethod *string `pulumi:"publicIPAllocationMethod"`
	// The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix *SubResourceResponse `pulumi:"publicIPPrefix"`
	// The resource GUID property of the public IP address resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The service public IP address of the public IP address resource.
	ServicePublicIPAddress *PublicIPAddressResponse `pulumi:"servicePublicIPAddress"`
	// The public IP address SKU.
	Sku *PublicIPAddressSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupPublicIPAddressResult
func (val *LookupPublicIPAddressResult) Defaults() *LookupPublicIPAddressResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.IpConfiguration = *tmp.IpConfiguration.Defaults()

	tmp.LinkedPublicIPAddress = tmp.LinkedPublicIPAddress.Defaults()

	tmp.ServicePublicIPAddress = tmp.ServicePublicIPAddress.Defaults()

	return &tmp
}

func LookupPublicIPAddressOutput(ctx *pulumi.Context, args LookupPublicIPAddressOutputArgs, opts ...pulumi.InvokeOption) LookupPublicIPAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicIPAddressResult, error) {
			args := v.(LookupPublicIPAddressArgs)
			r, err := LookupPublicIPAddress(ctx, &args, opts...)
			var s LookupPublicIPAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicIPAddressResultOutput)
}

type LookupPublicIPAddressOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the public IP address.
	PublicIpAddressName pulumi.StringInput `pulumi:"publicIpAddressName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPublicIPAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIPAddressArgs)(nil)).Elem()
}

// Public IP address resource.
type LookupPublicIPAddressResultOutput struct{ *pulumi.OutputState }

func (LookupPublicIPAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIPAddressResult)(nil)).Elem()
}

func (o LookupPublicIPAddressResultOutput) ToLookupPublicIPAddressResultOutput() LookupPublicIPAddressResultOutput {
	return o
}

func (o LookupPublicIPAddressResultOutput) ToLookupPublicIPAddressResultOutputWithContext(ctx context.Context) LookupPublicIPAddressResultOutput {
	return o
}

func (o LookupPublicIPAddressResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPublicIPAddressResult] {
	return pulumix.Output[LookupPublicIPAddressResult]{
		OutputState: o.OutputState,
	}
}

// The DDoS protection custom policy associated with the public IP address.
func (o LookupPublicIPAddressResultOutput) DdosSettings() DdosSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *DdosSettingsResponse { return v.DdosSettings }).(DdosSettingsResponsePtrOutput)
}

// Specify what happens to the public IP address when the VM using it is deleted
func (o LookupPublicIPAddressResultOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

// The FQDN of the DNS record associated with the public IP address.
func (o LookupPublicIPAddressResultOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressDnsSettingsResponse { return v.DnsSettings }).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupPublicIPAddressResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the public ip address.
func (o LookupPublicIPAddressResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Resource ID.
func (o LookupPublicIPAddressResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The idle timeout of the public IP address.
func (o LookupPublicIPAddressResultOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// The IP address associated with the public IP address resource.
func (o LookupPublicIPAddressResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// The IP configuration associated with the public IP address.
func (o LookupPublicIPAddressResultOutput) IpConfiguration() IPConfigurationResponseOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) IPConfigurationResponse { return v.IpConfiguration }).(IPConfigurationResponseOutput)
}

// The list of tags associated with the public IP address.
func (o LookupPublicIPAddressResultOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) []IpTagResponse { return v.IpTags }).(IpTagResponseArrayOutput)
}

// The linked public IP address of the public IP address resource.
func (o LookupPublicIPAddressResultOutput) LinkedPublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressResponse { return v.LinkedPublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

// Resource location.
func (o LookupPublicIPAddressResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Migration phase of Public IP Address.
func (o LookupPublicIPAddressResultOutput) MigrationPhase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.MigrationPhase }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupPublicIPAddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.Name }).(pulumi.StringOutput)
}

// The NatGateway for the Public IP address.
func (o LookupPublicIPAddressResultOutput) NatGateway() NatGatewayResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *NatGatewayResponse { return v.NatGateway }).(NatGatewayResponsePtrOutput)
}

// The provisioning state of the public IP address resource.
func (o LookupPublicIPAddressResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The public IP address version.
func (o LookupPublicIPAddressResultOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

// The public IP address allocation method.
func (o LookupPublicIPAddressResultOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

// The Public IP Prefix this Public IP Address should be allocated from.
func (o LookupPublicIPAddressResultOutput) PublicIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *SubResourceResponse { return v.PublicIPPrefix }).(SubResourceResponsePtrOutput)
}

// The resource GUID property of the public IP address resource.
func (o LookupPublicIPAddressResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The service public IP address of the public IP address resource.
func (o LookupPublicIPAddressResultOutput) ServicePublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressResponse { return v.ServicePublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

// The public IP address SKU.
func (o LookupPublicIPAddressResultOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressSkuResponse { return v.Sku }).(PublicIPAddressSkuResponsePtrOutput)
}

// Resource tags.
func (o LookupPublicIPAddressResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupPublicIPAddressResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the IP allocated for the resource needs to come from.
func (o LookupPublicIPAddressResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicIPAddressResultOutput{})
}

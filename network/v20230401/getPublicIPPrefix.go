// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified public IP prefix in a specified resource group.
func LookupPublicIPPrefix(ctx *pulumi.Context, args *LookupPublicIPPrefixArgs, opts ...pulumi.InvokeOption) (*LookupPublicIPPrefixResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPublicIPPrefixResult
	err := ctx.Invoke("azure-native:network/v20230401:getPublicIPPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicIPPrefixArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the public IP prefix.
	PublicIpPrefixName string `pulumi:"publicIpPrefixName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Public IP prefix resource.
type LookupPublicIPPrefixResult struct {
	// The customIpPrefix that this prefix is associated with.
	CustomIPPrefix *SubResourceResponse `pulumi:"customIPPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The extended location of the public ip address.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The allocated Prefix.
	IpPrefix string `pulumi:"ipPrefix"`
	// The list of tags associated with the public IP prefix.
	IpTags []IpTagResponse `pulumi:"ipTags"`
	// The reference to load balancer frontend IP configuration associated with the public IP prefix.
	LoadBalancerFrontendIpConfiguration SubResourceResponse `pulumi:"loadBalancerFrontendIpConfiguration"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// NatGateway of Public IP Prefix.
	NatGateway *NatGatewayResponse `pulumi:"natGateway"`
	// The Length of the Public IP Prefix.
	PrefixLength *int `pulumi:"prefixLength"`
	// The provisioning state of the public IP prefix resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The public IP address version.
	PublicIPAddressVersion *string `pulumi:"publicIPAddressVersion"`
	// The list of all referenced PublicIPAddresses.
	PublicIPAddresses []ReferencedPublicIpAddressResponse `pulumi:"publicIPAddresses"`
	// The resource GUID property of the public IP prefix resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The public IP prefix SKU.
	Sku *PublicIPPrefixSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

func LookupPublicIPPrefixOutput(ctx *pulumi.Context, args LookupPublicIPPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupPublicIPPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicIPPrefixResult, error) {
			args := v.(LookupPublicIPPrefixArgs)
			r, err := LookupPublicIPPrefix(ctx, &args, opts...)
			var s LookupPublicIPPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicIPPrefixResultOutput)
}

type LookupPublicIPPrefixOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the public IP prefix.
	PublicIpPrefixName pulumi.StringInput `pulumi:"publicIpPrefixName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPublicIPPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIPPrefixArgs)(nil)).Elem()
}

// Public IP prefix resource.
type LookupPublicIPPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupPublicIPPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIPPrefixResult)(nil)).Elem()
}

func (o LookupPublicIPPrefixResultOutput) ToLookupPublicIPPrefixResultOutput() LookupPublicIPPrefixResultOutput {
	return o
}

func (o LookupPublicIPPrefixResultOutput) ToLookupPublicIPPrefixResultOutputWithContext(ctx context.Context) LookupPublicIPPrefixResultOutput {
	return o
}

// The customIpPrefix that this prefix is associated with.
func (o LookupPublicIPPrefixResultOutput) CustomIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *SubResourceResponse { return v.CustomIPPrefix }).(SubResourceResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupPublicIPPrefixResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the public ip address.
func (o LookupPublicIPPrefixResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Resource ID.
func (o LookupPublicIPPrefixResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The allocated Prefix.
func (o LookupPublicIPPrefixResultOutput) IpPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.IpPrefix }).(pulumi.StringOutput)
}

// The list of tags associated with the public IP prefix.
func (o LookupPublicIPPrefixResultOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) []IpTagResponse { return v.IpTags }).(IpTagResponseArrayOutput)
}

// The reference to load balancer frontend IP configuration associated with the public IP prefix.
func (o LookupPublicIPPrefixResultOutput) LoadBalancerFrontendIpConfiguration() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) SubResourceResponse { return v.LoadBalancerFrontendIpConfiguration }).(SubResourceResponseOutput)
}

// Resource location.
func (o LookupPublicIPPrefixResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupPublicIPPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

// NatGateway of Public IP Prefix.
func (o LookupPublicIPPrefixResultOutput) NatGateway() NatGatewayResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *NatGatewayResponse { return v.NatGateway }).(NatGatewayResponsePtrOutput)
}

// The Length of the Public IP Prefix.
func (o LookupPublicIPPrefixResultOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *int { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

// The provisioning state of the public IP prefix resource.
func (o LookupPublicIPPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The public IP address version.
func (o LookupPublicIPPrefixResultOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

// The list of all referenced PublicIPAddresses.
func (o LookupPublicIPPrefixResultOutput) PublicIPAddresses() ReferencedPublicIpAddressResponseArrayOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) []ReferencedPublicIpAddressResponse { return v.PublicIPAddresses }).(ReferencedPublicIpAddressResponseArrayOutput)
}

// The resource GUID property of the public IP prefix resource.
func (o LookupPublicIPPrefixResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The public IP prefix SKU.
func (o LookupPublicIPPrefixResultOutput) Sku() PublicIPPrefixSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *PublicIPPrefixSkuResponse { return v.Sku }).(PublicIPPrefixSkuResponsePtrOutput)
}

// Resource tags.
func (o LookupPublicIPPrefixResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupPublicIPPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the IP allocated for the resource needs to come from.
func (o LookupPublicIPPrefixResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicIPPrefixResultOutput{})
}

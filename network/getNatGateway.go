// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified nat gateway in a specified resource group.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2019-06-01, 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01.
func LookupNatGateway(ctx *pulumi.Context, args *LookupNatGatewayArgs, opts ...pulumi.InvokeOption) (*LookupNatGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNatGatewayResult
	err := ctx.Invoke("azure-native:network:getNatGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNatGatewayArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the nat gateway.
	NatGatewayName string `pulumi:"natGatewayName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Nat Gateway resource.
type LookupNatGatewayResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The idle timeout of the nat gateway.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the NAT gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// An array of public ip addresses associated with the nat gateway resource.
	PublicIpAddresses []SubResourceResponse `pulumi:"publicIpAddresses"`
	// An array of public ip prefixes associated with the nat gateway resource.
	PublicIpPrefixes []SubResourceResponse `pulumi:"publicIpPrefixes"`
	// The resource GUID property of the NAT gateway resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The nat gateway SKU.
	Sku *NatGatewaySkuResponse `pulumi:"sku"`
	// An array of references to the subnets using this nat gateway resource.
	Subnets []SubResourceResponse `pulumi:"subnets"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones []string `pulumi:"zones"`
}

func LookupNatGatewayOutput(ctx *pulumi.Context, args LookupNatGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupNatGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNatGatewayResult, error) {
			args := v.(LookupNatGatewayArgs)
			r, err := LookupNatGateway(ctx, &args, opts...)
			var s LookupNatGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNatGatewayResultOutput)
}

type LookupNatGatewayOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the nat gateway.
	NatGatewayName pulumi.StringInput `pulumi:"natGatewayName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNatGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatGatewayArgs)(nil)).Elem()
}

// Nat Gateway resource.
type LookupNatGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupNatGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatGatewayResult)(nil)).Elem()
}

func (o LookupNatGatewayResultOutput) ToLookupNatGatewayResultOutput() LookupNatGatewayResultOutput {
	return o
}

func (o LookupNatGatewayResultOutput) ToLookupNatGatewayResultOutputWithContext(ctx context.Context) LookupNatGatewayResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNatGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNatGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The idle timeout of the nat gateway.
func (o LookupNatGatewayResultOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// Resource location.
func (o LookupNatGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNatGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the NAT gateway resource.
func (o LookupNatGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// An array of public ip addresses associated with the nat gateway resource.
func (o LookupNatGatewayResultOutput) PublicIpAddresses() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []SubResourceResponse { return v.PublicIpAddresses }).(SubResourceResponseArrayOutput)
}

// An array of public ip prefixes associated with the nat gateway resource.
func (o LookupNatGatewayResultOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []SubResourceResponse { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

// The resource GUID property of the NAT gateway resource.
func (o LookupNatGatewayResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The nat gateway SKU.
func (o LookupNatGatewayResultOutput) Sku() NatGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) *NatGatewaySkuResponse { return v.Sku }).(NatGatewaySkuResponsePtrOutput)
}

// An array of references to the subnets using this nat gateway resource.
func (o LookupNatGatewayResultOutput) Subnets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []SubResourceResponse { return v.Subnets }).(SubResourceResponseArrayOutput)
}

// Resource tags.
func (o LookupNatGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNatGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
func (o LookupNatGatewayResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNatGatewayResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNatGatewayResultOutput{})
}

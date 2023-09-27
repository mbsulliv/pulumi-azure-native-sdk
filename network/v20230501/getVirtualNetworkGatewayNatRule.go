// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the details of a nat rule.
func LookupVirtualNetworkGatewayNatRule(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayNatRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualNetworkGatewayNatRuleResult
	err := ctx.Invoke("azure-native:network/v20230501:getVirtualNetworkGatewayNatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkGatewayNatRuleArgs struct {
	// The name of the nat rule.
	NatRuleName string `pulumi:"natRuleName"`
	// The resource group name of the Virtual Network Gateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the gateway.
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}

// VirtualNetworkGatewayNatRule Resource.
type LookupVirtualNetworkGatewayNatRuleResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The private IP address external mapping for NAT.
	ExternalMappings []VpnNatRuleMappingResponse `pulumi:"externalMappings"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The private IP address internal mapping for NAT.
	InternalMappings []VpnNatRuleMappingResponse `pulumi:"internalMappings"`
	// The IP Configuration ID this NAT rule applies to.
	IpConfigurationId *string `pulumi:"ipConfigurationId"`
	// The Source NAT direction of a VPN NAT.
	Mode *string `pulumi:"mode"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the NAT Rule resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupVirtualNetworkGatewayNatRuleOutput(ctx *pulumi.Context, args LookupVirtualNetworkGatewayNatRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkGatewayNatRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkGatewayNatRuleResult, error) {
			args := v.(LookupVirtualNetworkGatewayNatRuleArgs)
			r, err := LookupVirtualNetworkGatewayNatRule(ctx, &args, opts...)
			var s LookupVirtualNetworkGatewayNatRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkGatewayNatRuleResultOutput)
}

type LookupVirtualNetworkGatewayNatRuleOutputArgs struct {
	// The name of the nat rule.
	NatRuleName pulumi.StringInput `pulumi:"natRuleName"`
	// The resource group name of the Virtual Network Gateway.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the gateway.
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (LookupVirtualNetworkGatewayNatRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayNatRuleArgs)(nil)).Elem()
}

// VirtualNetworkGatewayNatRule Resource.
type LookupVirtualNetworkGatewayNatRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkGatewayNatRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayNatRuleResult)(nil)).Elem()
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ToLookupVirtualNetworkGatewayNatRuleResultOutput() LookupVirtualNetworkGatewayNatRuleResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ToLookupVirtualNetworkGatewayNatRuleResultOutputWithContext(ctx context.Context) LookupVirtualNetworkGatewayNatRuleResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualNetworkGatewayNatRuleResult] {
	return pulumix.Output[LookupVirtualNetworkGatewayNatRuleResult]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The private IP address external mapping for NAT.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ExternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) []VpnNatRuleMappingResponse {
		return v.ExternalMappings
	}).(VpnNatRuleMappingResponseArrayOutput)
}

// Resource ID.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The private IP address internal mapping for NAT.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) InternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) []VpnNatRuleMappingResponse {
		return v.InternalMappings
	}).(VpnNatRuleMappingResponseArrayOutput)
}

// The IP Configuration ID this NAT rule applies to.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) *string { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

// The Source NAT direction of a VPN NAT.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the NAT Rule resource.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkGatewayNatRuleResultOutput{})
}

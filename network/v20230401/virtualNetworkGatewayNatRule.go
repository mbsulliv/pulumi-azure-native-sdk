// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// VirtualNetworkGatewayNatRule Resource.
type VirtualNetworkGatewayNatRule struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The private IP address external mapping for NAT.
	ExternalMappings VpnNatRuleMappingResponseArrayOutput `pulumi:"externalMappings"`
	// The private IP address internal mapping for NAT.
	InternalMappings VpnNatRuleMappingResponseArrayOutput `pulumi:"internalMappings"`
	// The IP Configuration ID this NAT rule applies to.
	IpConfigurationId pulumi.StringPtrOutput `pulumi:"ipConfigurationId"`
	// The Source NAT direction of a VPN NAT.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the NAT Rule resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualNetworkGatewayNatRule registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkGatewayNatRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayNatRuleArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayNatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetworkGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkGatewayName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualNetworkGatewayNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VirtualNetworkGatewayNatRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualNetworkGatewayNatRule
	err := ctx.RegisterResource("azure-native:network/v20230401:VirtualNetworkGatewayNatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkGatewayNatRule gets an existing VirtualNetworkGatewayNatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkGatewayNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayNatRuleState, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayNatRule, error) {
	var resource VirtualNetworkGatewayNatRule
	err := ctx.ReadResource("azure-native:network/v20230401:VirtualNetworkGatewayNatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkGatewayNatRule resources.
type virtualNetworkGatewayNatRuleState struct {
}

type VirtualNetworkGatewayNatRuleState struct {
}

func (VirtualNetworkGatewayNatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayNatRuleState)(nil)).Elem()
}

type virtualNetworkGatewayNatRuleArgs struct {
	// The private IP address external mapping for NAT.
	ExternalMappings []VpnNatRuleMapping `pulumi:"externalMappings"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The private IP address internal mapping for NAT.
	InternalMappings []VpnNatRuleMapping `pulumi:"internalMappings"`
	// The IP Configuration ID this NAT rule applies to.
	IpConfigurationId *string `pulumi:"ipConfigurationId"`
	// The Source NAT direction of a VPN NAT.
	Mode *string `pulumi:"mode"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the nat rule.
	NatRuleName *string `pulumi:"natRuleName"`
	// The resource group name of the Virtual Network Gateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of NAT rule for VPN NAT.
	Type *string `pulumi:"type"`
	// The name of the gateway.
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}

// The set of arguments for constructing a VirtualNetworkGatewayNatRule resource.
type VirtualNetworkGatewayNatRuleArgs struct {
	// The private IP address external mapping for NAT.
	ExternalMappings VpnNatRuleMappingArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The private IP address internal mapping for NAT.
	InternalMappings VpnNatRuleMappingArrayInput
	// The IP Configuration ID this NAT rule applies to.
	IpConfigurationId pulumi.StringPtrInput
	// The Source NAT direction of a VPN NAT.
	Mode pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the nat rule.
	NatRuleName pulumi.StringPtrInput
	// The resource group name of the Virtual Network Gateway.
	ResourceGroupName pulumi.StringInput
	// The type of NAT rule for VPN NAT.
	Type pulumi.StringPtrInput
	// The name of the gateway.
	VirtualNetworkGatewayName pulumi.StringInput
}

func (VirtualNetworkGatewayNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayNatRuleArgs)(nil)).Elem()
}

type VirtualNetworkGatewayNatRuleInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayNatRuleOutput() VirtualNetworkGatewayNatRuleOutput
	ToVirtualNetworkGatewayNatRuleOutputWithContext(ctx context.Context) VirtualNetworkGatewayNatRuleOutput
}

func (*VirtualNetworkGatewayNatRule) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayNatRule)(nil)).Elem()
}

func (i *VirtualNetworkGatewayNatRule) ToVirtualNetworkGatewayNatRuleOutput() VirtualNetworkGatewayNatRuleOutput {
	return i.ToVirtualNetworkGatewayNatRuleOutputWithContext(context.Background())
}

func (i *VirtualNetworkGatewayNatRule) ToVirtualNetworkGatewayNatRuleOutputWithContext(ctx context.Context) VirtualNetworkGatewayNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayNatRuleOutput)
}

func (i *VirtualNetworkGatewayNatRule) ToOutput(ctx context.Context) pulumix.Output[*VirtualNetworkGatewayNatRule] {
	return pulumix.Output[*VirtualNetworkGatewayNatRule]{
		OutputState: i.ToVirtualNetworkGatewayNatRuleOutputWithContext(ctx).OutputState,
	}
}

type VirtualNetworkGatewayNatRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayNatRule)(nil)).Elem()
}

func (o VirtualNetworkGatewayNatRuleOutput) ToVirtualNetworkGatewayNatRuleOutput() VirtualNetworkGatewayNatRuleOutput {
	return o
}

func (o VirtualNetworkGatewayNatRuleOutput) ToVirtualNetworkGatewayNatRuleOutputWithContext(ctx context.Context) VirtualNetworkGatewayNatRuleOutput {
	return o
}

func (o VirtualNetworkGatewayNatRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualNetworkGatewayNatRule] {
	return pulumix.Output[*VirtualNetworkGatewayNatRule]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualNetworkGatewayNatRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayNatRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The private IP address external mapping for NAT.
func (o VirtualNetworkGatewayNatRuleOutput) ExternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayNatRule) VpnNatRuleMappingResponseArrayOutput { return v.ExternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

// The private IP address internal mapping for NAT.
func (o VirtualNetworkGatewayNatRuleOutput) InternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayNatRule) VpnNatRuleMappingResponseArrayOutput { return v.InternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

// The IP Configuration ID this NAT rule applies to.
func (o VirtualNetworkGatewayNatRuleOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayNatRule) pulumi.StringPtrOutput { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

// The Source NAT direction of a VPN NAT.
func (o VirtualNetworkGatewayNatRuleOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayNatRule) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o VirtualNetworkGatewayNatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayNatRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the NAT Rule resource.
func (o VirtualNetworkGatewayNatRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayNatRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o VirtualNetworkGatewayNatRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayNatRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkGatewayNatRuleOutput{})
}

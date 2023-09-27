// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// VpnGatewayNatRule Resource.
type NatRule struct {
	pulumi.CustomResourceState

	// List of egress VpnSiteLinkConnections.
	EgressVpnSiteLinkConnections SubResourceResponseArrayOutput `pulumi:"egressVpnSiteLinkConnections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The private IP address external mapping for NAT.
	ExternalMappings VpnNatRuleMappingResponseArrayOutput `pulumi:"externalMappings"`
	// List of ingress VpnSiteLinkConnections.
	IngressVpnSiteLinkConnections SubResourceResponseArrayOutput `pulumi:"ingressVpnSiteLinkConnections"`
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

// NewNatRule registers a new resource with the given unique name, arguments, and options.
func NewNatRule(ctx *pulumi.Context,
	name string, args *NatRuleArgs, opts ...pulumi.ResourceOption) (*NatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:NatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:NatRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NatRule
	err := ctx.RegisterResource("azure-native:network/v20230501:NatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatRule gets an existing NatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatRuleState, opts ...pulumi.ResourceOption) (*NatRule, error) {
	var resource NatRule
	err := ctx.ReadResource("azure-native:network/v20230501:NatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatRule resources.
type natRuleState struct {
}

type NatRuleState struct {
}

func (NatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*natRuleState)(nil)).Elem()
}

type natRuleArgs struct {
	// The private IP address external mapping for NAT.
	ExternalMappings []VpnNatRuleMapping `pulumi:"externalMappings"`
	// The name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
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
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of NAT rule for VPN NAT.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a NatRule resource.
type NatRuleArgs struct {
	// The private IP address external mapping for NAT.
	ExternalMappings VpnNatRuleMappingArrayInput
	// The name of the gateway.
	GatewayName pulumi.StringInput
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
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput
	// The type of NAT rule for VPN NAT.
	Type pulumi.StringPtrInput
}

func (NatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natRuleArgs)(nil)).Elem()
}

type NatRuleInput interface {
	pulumi.Input

	ToNatRuleOutput() NatRuleOutput
	ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput
}

func (*NatRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NatRule)(nil)).Elem()
}

func (i *NatRule) ToNatRuleOutput() NatRuleOutput {
	return i.ToNatRuleOutputWithContext(context.Background())
}

func (i *NatRule) ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatRuleOutput)
}

func (i *NatRule) ToOutput(ctx context.Context) pulumix.Output[*NatRule] {
	return pulumix.Output[*NatRule]{
		OutputState: i.ToNatRuleOutputWithContext(ctx).OutputState,
	}
}

type NatRuleOutput struct{ *pulumi.OutputState }

func (NatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatRule)(nil)).Elem()
}

func (o NatRuleOutput) ToNatRuleOutput() NatRuleOutput {
	return o
}

func (o NatRuleOutput) ToNatRuleOutputWithContext(ctx context.Context) NatRuleOutput {
	return o
}

func (o NatRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*NatRule] {
	return pulumix.Output[*NatRule]{
		OutputState: o.OutputState,
	}
}

// List of egress VpnSiteLinkConnections.
func (o NatRuleOutput) EgressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatRule) SubResourceResponseArrayOutput { return v.EgressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o NatRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The private IP address external mapping for NAT.
func (o NatRuleOutput) ExternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v *NatRule) VpnNatRuleMappingResponseArrayOutput { return v.ExternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

// List of ingress VpnSiteLinkConnections.
func (o NatRuleOutput) IngressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatRule) SubResourceResponseArrayOutput { return v.IngressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

// The private IP address internal mapping for NAT.
func (o NatRuleOutput) InternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v *NatRule) VpnNatRuleMappingResponseArrayOutput { return v.InternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

// The IP Configuration ID this NAT rule applies to.
func (o NatRuleOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringPtrOutput { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

// The Source NAT direction of a VPN NAT.
func (o NatRuleOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o NatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the NAT Rule resource.
func (o NatRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o NatRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NatRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NatRuleOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A virtual network rule.
type VirtualNetworkRule struct {
	pulumi.CustomResourceState

	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrOutput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Virtual Network Rule State
	State pulumi.StringOutput `pulumi:"state"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId pulumi.StringOutput `pulumi:"virtualNetworkSubnetId"`
}

// NewVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.VirtualNetworkSubnetId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkSubnetId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:VirtualNetworkRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:VirtualNetworkRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualNetworkRule
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:VirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkRule gets an existing VirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*VirtualNetworkRule, error) {
	var resource VirtualNetworkRule
	err := ctx.ReadResource("azure-native:sql/v20230201preview:VirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkRule resources.
type virtualNetworkRuleState struct {
}

type VirtualNetworkRuleState struct {
}

func (VirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleState)(nil)).Elem()
}

type virtualNetworkRuleArgs struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `pulumi:"ignoreMissingVnetServiceEndpoint"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the virtual network rule.
	VirtualNetworkRuleName *string `pulumi:"virtualNetworkRuleName"`
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId string `pulumi:"virtualNetworkSubnetId"`
}

// The set of arguments for constructing a VirtualNetworkRule resource.
type VirtualNetworkRuleArgs struct {
	// Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name of the virtual network rule.
	VirtualNetworkRuleName pulumi.StringPtrInput
	// The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId pulumi.StringInput
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRuleArgs)(nil)).Elem()
}

type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput
}

func (*VirtualNetworkRule) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRule)(nil)).Elem()
}

func (i *VirtualNetworkRule) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i *VirtualNetworkRule) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

// Create firewall rule before the virtual network has vnet service endpoint enabled.
func (o VirtualNetworkRuleOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.BoolPtrOutput { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

// Resource name.
func (o VirtualNetworkRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Virtual Network Rule State
func (o VirtualNetworkRuleOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Resource type.
func (o VirtualNetworkRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The ARM resource id of the virtual network subnet.
func (o VirtualNetworkRuleOutput) VirtualNetworkSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRule) pulumi.StringOutput { return v.VirtualNetworkSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
}

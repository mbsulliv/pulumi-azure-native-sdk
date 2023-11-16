// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single item in a List or Get VirtualNetworkRules operation
// Azure REST API version: 2018-01-01-preview. Prior API version in Azure Native 1.x: 2018-01-01-preview.
type NamespaceVirtualNetworkRule struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Resource ID of Virtual Network Subnet
	VirtualNetworkSubnetId pulumi.StringPtrOutput `pulumi:"virtualNetworkSubnetId"`
}

// NewNamespaceVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewNamespaceVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *NamespaceVirtualNetworkRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceVirtualNetworkRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:NamespaceVirtualNetworkRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NamespaceVirtualNetworkRule
	err := ctx.RegisterResource("azure-native:servicebus:NamespaceVirtualNetworkRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceVirtualNetworkRule gets an existing NamespaceVirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceVirtualNetworkRuleState, opts ...pulumi.ResourceOption) (*NamespaceVirtualNetworkRule, error) {
	var resource NamespaceVirtualNetworkRule
	err := ctx.ReadResource("azure-native:servicebus:NamespaceVirtualNetworkRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceVirtualNetworkRule resources.
type namespaceVirtualNetworkRuleState struct {
}

type NamespaceVirtualNetworkRuleState struct {
}

func (NamespaceVirtualNetworkRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceVirtualNetworkRuleState)(nil)).Elem()
}

type namespaceVirtualNetworkRuleArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Virtual Network Rule name.
	VirtualNetworkRuleName *string `pulumi:"virtualNetworkRuleName"`
	// Resource ID of Virtual Network Subnet
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}

// The set of arguments for constructing a NamespaceVirtualNetworkRule resource.
type NamespaceVirtualNetworkRuleArgs struct {
	// The namespace name
	NamespaceName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The Virtual Network Rule name.
	VirtualNetworkRuleName pulumi.StringPtrInput
	// Resource ID of Virtual Network Subnet
	VirtualNetworkSubnetId pulumi.StringPtrInput
}

func (NamespaceVirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceVirtualNetworkRuleArgs)(nil)).Elem()
}

type NamespaceVirtualNetworkRuleInput interface {
	pulumi.Input

	ToNamespaceVirtualNetworkRuleOutput() NamespaceVirtualNetworkRuleOutput
	ToNamespaceVirtualNetworkRuleOutputWithContext(ctx context.Context) NamespaceVirtualNetworkRuleOutput
}

func (*NamespaceVirtualNetworkRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceVirtualNetworkRule)(nil)).Elem()
}

func (i *NamespaceVirtualNetworkRule) ToNamespaceVirtualNetworkRuleOutput() NamespaceVirtualNetworkRuleOutput {
	return i.ToNamespaceVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i *NamespaceVirtualNetworkRule) ToNamespaceVirtualNetworkRuleOutputWithContext(ctx context.Context) NamespaceVirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceVirtualNetworkRuleOutput)
}

type NamespaceVirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (NamespaceVirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceVirtualNetworkRule)(nil)).Elem()
}

func (o NamespaceVirtualNetworkRuleOutput) ToNamespaceVirtualNetworkRuleOutput() NamespaceVirtualNetworkRuleOutput {
	return o
}

func (o NamespaceVirtualNetworkRuleOutput) ToNamespaceVirtualNetworkRuleOutputWithContext(ctx context.Context) NamespaceVirtualNetworkRuleOutput {
	return o
}

// Resource name
func (o NamespaceVirtualNetworkRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceVirtualNetworkRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o NamespaceVirtualNetworkRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceVirtualNetworkRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Resource ID of Virtual Network Subnet
func (o NamespaceVirtualNetworkRuleOutput) VirtualNetworkSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceVirtualNetworkRule) pulumi.StringPtrOutput { return v.VirtualNetworkSubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceVirtualNetworkRuleOutput{})
}

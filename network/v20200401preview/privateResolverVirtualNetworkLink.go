// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a virtual network link.
type PrivateResolverVirtualNetworkLink struct {
	pulumi.CustomResourceState

	// ETag of the virtual network link.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Metadata attached to the virtual network link.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The current provisioning state of the virtual network link. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork SubResourceResponsePtrOutput `pulumi:"virtualNetwork"`
}

// NewPrivateResolverVirtualNetworkLink registers a new resource with the given unique name, arguments, and options.
func NewPrivateResolverVirtualNetworkLink(ctx *pulumi.Context,
	name string, args *PrivateResolverVirtualNetworkLinkArgs, opts ...pulumi.ResourceOption) (*PrivateResolverVirtualNetworkLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsForwardingRulesetName == nil {
		return nil, errors.New("invalid value for required argument 'DnsForwardingRulesetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PrivateResolverVirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:PrivateResolverVirtualNetworkLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateResolverVirtualNetworkLink
	err := ctx.RegisterResource("azure-native:network/v20200401preview:PrivateResolverVirtualNetworkLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateResolverVirtualNetworkLink gets an existing PrivateResolverVirtualNetworkLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateResolverVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateResolverVirtualNetworkLinkState, opts ...pulumi.ResourceOption) (*PrivateResolverVirtualNetworkLink, error) {
	var resource PrivateResolverVirtualNetworkLink
	err := ctx.ReadResource("azure-native:network/v20200401preview:PrivateResolverVirtualNetworkLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateResolverVirtualNetworkLink resources.
type privateResolverVirtualNetworkLinkState struct {
}

type PrivateResolverVirtualNetworkLinkState struct {
}

func (PrivateResolverVirtualNetworkLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateResolverVirtualNetworkLinkState)(nil)).Elem()
}

type privateResolverVirtualNetworkLinkArgs struct {
	// The name of the DNS forwarding ruleset.
	DnsForwardingRulesetName string `pulumi:"dnsForwardingRulesetName"`
	// Metadata attached to the virtual network link.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork *SubResource `pulumi:"virtualNetwork"`
	// The name of the virtual network link.
	VirtualNetworkLinkName *string `pulumi:"virtualNetworkLinkName"`
}

// The set of arguments for constructing a PrivateResolverVirtualNetworkLink resource.
type PrivateResolverVirtualNetworkLinkArgs struct {
	// The name of the DNS forwarding ruleset.
	DnsForwardingRulesetName pulumi.StringInput
	// Metadata attached to the virtual network link.
	Metadata pulumi.StringMapInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork SubResourcePtrInput
	// The name of the virtual network link.
	VirtualNetworkLinkName pulumi.StringPtrInput
}

func (PrivateResolverVirtualNetworkLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateResolverVirtualNetworkLinkArgs)(nil)).Elem()
}

type PrivateResolverVirtualNetworkLinkInput interface {
	pulumi.Input

	ToPrivateResolverVirtualNetworkLinkOutput() PrivateResolverVirtualNetworkLinkOutput
	ToPrivateResolverVirtualNetworkLinkOutputWithContext(ctx context.Context) PrivateResolverVirtualNetworkLinkOutput
}

func (*PrivateResolverVirtualNetworkLink) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateResolverVirtualNetworkLink)(nil)).Elem()
}

func (i *PrivateResolverVirtualNetworkLink) ToPrivateResolverVirtualNetworkLinkOutput() PrivateResolverVirtualNetworkLinkOutput {
	return i.ToPrivateResolverVirtualNetworkLinkOutputWithContext(context.Background())
}

func (i *PrivateResolverVirtualNetworkLink) ToPrivateResolverVirtualNetworkLinkOutputWithContext(ctx context.Context) PrivateResolverVirtualNetworkLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateResolverVirtualNetworkLinkOutput)
}

type PrivateResolverVirtualNetworkLinkOutput struct{ *pulumi.OutputState }

func (PrivateResolverVirtualNetworkLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateResolverVirtualNetworkLink)(nil)).Elem()
}

func (o PrivateResolverVirtualNetworkLinkOutput) ToPrivateResolverVirtualNetworkLinkOutput() PrivateResolverVirtualNetworkLinkOutput {
	return o
}

func (o PrivateResolverVirtualNetworkLinkOutput) ToPrivateResolverVirtualNetworkLinkOutputWithContext(ctx context.Context) PrivateResolverVirtualNetworkLinkOutput {
	return o
}

// ETag of the virtual network link.
func (o PrivateResolverVirtualNetworkLinkOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateResolverVirtualNetworkLink) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Metadata attached to the virtual network link.
func (o PrivateResolverVirtualNetworkLinkOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateResolverVirtualNetworkLink) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o PrivateResolverVirtualNetworkLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateResolverVirtualNetworkLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the virtual network link. This is a read-only property and any attempt to set this value will be ignored.
func (o PrivateResolverVirtualNetworkLinkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateResolverVirtualNetworkLink) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o PrivateResolverVirtualNetworkLinkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateResolverVirtualNetworkLink) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateResolverVirtualNetworkLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateResolverVirtualNetworkLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The reference to the virtual network. This cannot be changed after creation.
func (o PrivateResolverVirtualNetworkLinkOutput) VirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *PrivateResolverVirtualNetworkLink) SubResourceResponsePtrOutput { return v.VirtualNetwork }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateResolverVirtualNetworkLinkOutput{})
}

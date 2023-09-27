// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Describes a DNS forwarding ruleset.
type DnsForwardingRuleset struct {
	pulumi.CustomResourceState

	// The reference to the DNS resolver outbound endpoints that are used to route DNS queries matching the forwarding rules in the ruleset to the target DNS servers.
	DnsResolverOutboundEndpoints SubResourceResponseArrayOutput `pulumi:"dnsResolverOutboundEndpoints"`
	// ETag of the DNS forwarding ruleset.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The current provisioning state of the DNS forwarding ruleset. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resourceGuid for the DNS forwarding ruleset.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDnsForwardingRuleset registers a new resource with the given unique name, arguments, and options.
func NewDnsForwardingRuleset(ctx *pulumi.Context,
	name string, args *DnsForwardingRulesetArgs, opts ...pulumi.ResourceOption) (*DnsForwardingRuleset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsResolverOutboundEndpoints == nil {
		return nil, errors.New("invalid value for required argument 'DnsResolverOutboundEndpoints'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:DnsForwardingRuleset"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401preview:DnsForwardingRuleset"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DnsForwardingRuleset
	err := ctx.RegisterResource("azure-native:network/v20220701:DnsForwardingRuleset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsForwardingRuleset gets an existing DnsForwardingRuleset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsForwardingRuleset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsForwardingRulesetState, opts ...pulumi.ResourceOption) (*DnsForwardingRuleset, error) {
	var resource DnsForwardingRuleset
	err := ctx.ReadResource("azure-native:network/v20220701:DnsForwardingRuleset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsForwardingRuleset resources.
type dnsForwardingRulesetState struct {
}

type DnsForwardingRulesetState struct {
}

func (DnsForwardingRulesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsForwardingRulesetState)(nil)).Elem()
}

type dnsForwardingRulesetArgs struct {
	// The name of the DNS forwarding ruleset.
	DnsForwardingRulesetName *string `pulumi:"dnsForwardingRulesetName"`
	// The reference to the DNS resolver outbound endpoints that are used to route DNS queries matching the forwarding rules in the ruleset to the target DNS servers.
	DnsResolverOutboundEndpoints []SubResource `pulumi:"dnsResolverOutboundEndpoints"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DnsForwardingRuleset resource.
type DnsForwardingRulesetArgs struct {
	// The name of the DNS forwarding ruleset.
	DnsForwardingRulesetName pulumi.StringPtrInput
	// The reference to the DNS resolver outbound endpoints that are used to route DNS queries matching the forwarding rules in the ruleset to the target DNS servers.
	DnsResolverOutboundEndpoints SubResourceArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DnsForwardingRulesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsForwardingRulesetArgs)(nil)).Elem()
}

type DnsForwardingRulesetInput interface {
	pulumi.Input

	ToDnsForwardingRulesetOutput() DnsForwardingRulesetOutput
	ToDnsForwardingRulesetOutputWithContext(ctx context.Context) DnsForwardingRulesetOutput
}

func (*DnsForwardingRuleset) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsForwardingRuleset)(nil)).Elem()
}

func (i *DnsForwardingRuleset) ToDnsForwardingRulesetOutput() DnsForwardingRulesetOutput {
	return i.ToDnsForwardingRulesetOutputWithContext(context.Background())
}

func (i *DnsForwardingRuleset) ToDnsForwardingRulesetOutputWithContext(ctx context.Context) DnsForwardingRulesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsForwardingRulesetOutput)
}

func (i *DnsForwardingRuleset) ToOutput(ctx context.Context) pulumix.Output[*DnsForwardingRuleset] {
	return pulumix.Output[*DnsForwardingRuleset]{
		OutputState: i.ToDnsForwardingRulesetOutputWithContext(ctx).OutputState,
	}
}

type DnsForwardingRulesetOutput struct{ *pulumi.OutputState }

func (DnsForwardingRulesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsForwardingRuleset)(nil)).Elem()
}

func (o DnsForwardingRulesetOutput) ToDnsForwardingRulesetOutput() DnsForwardingRulesetOutput {
	return o
}

func (o DnsForwardingRulesetOutput) ToDnsForwardingRulesetOutputWithContext(ctx context.Context) DnsForwardingRulesetOutput {
	return o
}

func (o DnsForwardingRulesetOutput) ToOutput(ctx context.Context) pulumix.Output[*DnsForwardingRuleset] {
	return pulumix.Output[*DnsForwardingRuleset]{
		OutputState: o.OutputState,
	}
}

// The reference to the DNS resolver outbound endpoints that are used to route DNS queries matching the forwarding rules in the ruleset to the target DNS servers.
func (o DnsForwardingRulesetOutput) DnsResolverOutboundEndpoints() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) SubResourceResponseArrayOutput { return v.DnsResolverOutboundEndpoints }).(SubResourceResponseArrayOutput)
}

// ETag of the DNS forwarding ruleset.
func (o DnsForwardingRulesetOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o DnsForwardingRulesetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DnsForwardingRulesetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the DNS forwarding ruleset. This is a read-only property and any attempt to set this value will be ignored.
func (o DnsForwardingRulesetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid for the DNS forwarding ruleset.
func (o DnsForwardingRulesetOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DnsForwardingRulesetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DnsForwardingRulesetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DnsForwardingRulesetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsForwardingRuleset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsForwardingRulesetOutput{})
}

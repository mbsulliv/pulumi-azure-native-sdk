// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Firewall resource.
type AzureFirewall struct {
	pulumi.CustomResourceState

	// The additional properties used to further config this azure firewall.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionResponseArrayOutput `pulumi:"applicationRuleCollections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy SubResourceResponsePtrOutput `pulumi:"firewallPolicy"`
	// IP addresses associated with AzureFirewall.
	HubIPAddresses HubIPAddressesResponsePtrOutput `pulumi:"hubIPAddresses"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// IpGroups associated with AzureFirewall.
	IpGroups AzureFirewallIpGroupsResponseArrayOutput `pulumi:"ipGroups"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// IP configuration of the Azure Firewall used for management traffic.
	ManagementIpConfiguration AzureFirewallIPConfigurationResponsePtrOutput `pulumi:"managementIpConfiguration"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionResponseArrayOutput `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionResponseArrayOutput `pulumi:"networkRuleCollections"`
	// The provisioning state of the Azure firewall resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The Azure Firewall Resource SKU.
	Sku AzureFirewallSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrOutput `pulumi:"threatIntelMode"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The virtualHub to which the firewall belongs.
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewAzureFirewall registers a new resource with the given unique name, arguments, and options.
func NewAzureFirewall(ctx *pulumi.Context,
	name string, args *AzureFirewallArgs, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:AzureFirewall"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AzureFirewall
	err := ctx.RegisterResource("azure-native:network/v20230201:AzureFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureFirewall gets an existing AzureFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureFirewallState, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	var resource AzureFirewall
	err := ctx.ReadResource("azure-native:network/v20230201:AzureFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureFirewall resources.
type azureFirewallState struct {
}

type AzureFirewallState struct {
}

func (AzureFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallState)(nil)).Elem()
}

type azureFirewallArgs struct {
	// The additional properties used to further config this azure firewall.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollection `pulumi:"applicationRuleCollections"`
	// The name of the Azure Firewall.
	AzureFirewallName *string `pulumi:"azureFirewallName"`
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy *SubResource `pulumi:"firewallPolicy"`
	// IP addresses associated with AzureFirewall.
	HubIPAddresses *HubIPAddresses `pulumi:"hubIPAddresses"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations []AzureFirewallIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// IP configuration of the Azure Firewall used for management traffic.
	ManagementIpConfiguration *AzureFirewallIPConfiguration `pulumi:"managementIpConfiguration"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections []AzureFirewallNatRuleCollection `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections []AzureFirewallNetworkRuleCollection `pulumi:"networkRuleCollections"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Azure Firewall Resource SKU.
	Sku *AzureFirewallSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// The virtualHub to which the firewall belongs.
	VirtualHub *SubResource `pulumi:"virtualHub"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a AzureFirewall resource.
type AzureFirewallArgs struct {
	// The additional properties used to further config this azure firewall.
	AdditionalProperties pulumi.StringMapInput
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionArrayInput
	// The name of the Azure Firewall.
	AzureFirewallName pulumi.StringPtrInput
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy SubResourcePtrInput
	// IP addresses associated with AzureFirewall.
	HubIPAddresses HubIPAddressesPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// IP configuration of the Azure Firewall used for management traffic.
	ManagementIpConfiguration AzureFirewallIPConfigurationPtrInput
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionArrayInput
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The Azure Firewall Resource SKU.
	Sku AzureFirewallSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
	// The virtualHub to which the firewall belongs.
	VirtualHub SubResourcePtrInput
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (AzureFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallArgs)(nil)).Elem()
}

type AzureFirewallInput interface {
	pulumi.Input

	ToAzureFirewallOutput() AzureFirewallOutput
	ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput
}

func (*AzureFirewall) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewall)(nil)).Elem()
}

func (i *AzureFirewall) ToAzureFirewallOutput() AzureFirewallOutput {
	return i.ToAzureFirewallOutputWithContext(context.Background())
}

func (i *AzureFirewall) ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFirewallOutput)
}

func (i *AzureFirewall) ToOutput(ctx context.Context) pulumix.Output[*AzureFirewall] {
	return pulumix.Output[*AzureFirewall]{
		OutputState: i.ToAzureFirewallOutputWithContext(ctx).OutputState,
	}
}

type AzureFirewallOutput struct{ *pulumi.OutputState }

func (AzureFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewall)(nil)).Elem()
}

func (o AzureFirewallOutput) ToAzureFirewallOutput() AzureFirewallOutput {
	return o
}

func (o AzureFirewallOutput) ToAzureFirewallOutputWithContext(ctx context.Context) AzureFirewallOutput {
	return o
}

func (o AzureFirewallOutput) ToOutput(ctx context.Context) pulumix.Output[*AzureFirewall] {
	return pulumix.Output[*AzureFirewall]{
		OutputState: o.OutputState,
	}
}

// The additional properties used to further config this azure firewall.
func (o AzureFirewallOutput) AdditionalProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringMapOutput { return v.AdditionalProperties }).(pulumi.StringMapOutput)
}

// Collection of application rule collections used by Azure Firewall.
func (o AzureFirewallOutput) ApplicationRuleCollections() AzureFirewallApplicationRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallApplicationRuleCollectionResponseArrayOutput {
		return v.ApplicationRuleCollections
	}).(AzureFirewallApplicationRuleCollectionResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o AzureFirewallOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The firewallPolicy associated with this azure firewall.
func (o AzureFirewallOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) SubResourceResponsePtrOutput { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
}

// IP addresses associated with AzureFirewall.
func (o AzureFirewallOutput) HubIPAddresses() HubIPAddressesResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) HubIPAddressesResponsePtrOutput { return v.HubIPAddresses }).(HubIPAddressesResponsePtrOutput)
}

// IP configuration of the Azure Firewall resource.
func (o AzureFirewallOutput) IpConfigurations() AzureFirewallIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallIPConfigurationResponseArrayOutput { return v.IpConfigurations }).(AzureFirewallIPConfigurationResponseArrayOutput)
}

// IpGroups associated with AzureFirewall.
func (o AzureFirewallOutput) IpGroups() AzureFirewallIpGroupsResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallIpGroupsResponseArrayOutput { return v.IpGroups }).(AzureFirewallIpGroupsResponseArrayOutput)
}

// Resource location.
func (o AzureFirewallOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// IP configuration of the Azure Firewall used for management traffic.
func (o AzureFirewallOutput) ManagementIpConfiguration() AzureFirewallIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallIPConfigurationResponsePtrOutput {
		return v.ManagementIpConfiguration
	}).(AzureFirewallIPConfigurationResponsePtrOutput)
}

// Resource name.
func (o AzureFirewallOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Collection of NAT rule collections used by Azure Firewall.
func (o AzureFirewallOutput) NatRuleCollections() AzureFirewallNatRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallNatRuleCollectionResponseArrayOutput { return v.NatRuleCollections }).(AzureFirewallNatRuleCollectionResponseArrayOutput)
}

// Collection of network rule collections used by Azure Firewall.
func (o AzureFirewallOutput) NetworkRuleCollections() AzureFirewallNetworkRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallNetworkRuleCollectionResponseArrayOutput {
		return v.NetworkRuleCollections
	}).(AzureFirewallNetworkRuleCollectionResponseArrayOutput)
}

// The provisioning state of the Azure firewall resource.
func (o AzureFirewallOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Azure Firewall Resource SKU.
func (o AzureFirewallOutput) Sku() AzureFirewallSkuResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) AzureFirewallSkuResponsePtrOutput { return v.Sku }).(AzureFirewallSkuResponsePtrOutput)
}

// Resource tags.
func (o AzureFirewallOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The operation mode for Threat Intelligence.
func (o AzureFirewallOutput) ThreatIntelMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringPtrOutput { return v.ThreatIntelMode }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o AzureFirewallOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The virtualHub to which the firewall belongs.
func (o AzureFirewallOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *AzureFirewall) SubResourceResponsePtrOutput { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

// A list of availability zones denoting where the resource needs to come from.
func (o AzureFirewallOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureFirewall) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureFirewallOutput{})
}

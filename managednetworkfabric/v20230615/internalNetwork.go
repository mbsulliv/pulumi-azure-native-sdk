// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the Internal Network resource.
type InternalNetwork struct {
	pulumi.CustomResourceState

	// Administrative state of the resource.
	AdministrativeState pulumi.StringOutput `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// BGP configuration properties.
	BgpConfiguration InternalNetworkPropertiesResponseBgpConfigurationPtrOutput `pulumi:"bgpConfiguration"`
	// Configuration state of the resource.
	ConfigurationState pulumi.StringOutput `pulumi:"configurationState"`
	// List of Connected IPv4 Subnets.
	ConnectedIPv4Subnets ConnectedSubnetResponseArrayOutput `pulumi:"connectedIPv4Subnets"`
	// List of connected IPv6 Subnets.
	ConnectedIPv6Subnets ConnectedSubnetResponseArrayOutput `pulumi:"connectedIPv6Subnets"`
	// Egress Acl. ARM resource ID of Access Control Lists.
	EgressAclId pulumi.StringPtrOutput `pulumi:"egressAclId"`
	// Export Route Policy either IPv4 or IPv6.
	ExportRoutePolicy ExportRoutePolicyResponsePtrOutput `pulumi:"exportRoutePolicy"`
	// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
	ExportRoutePolicyId pulumi.StringPtrOutput `pulumi:"exportRoutePolicyId"`
	// Extension. Example: NoExtension | NPB.
	Extension pulumi.StringPtrOutput `pulumi:"extension"`
	// Import Route Policy either IPv4 or IPv6.
	ImportRoutePolicy ImportRoutePolicyResponsePtrOutput `pulumi:"importRoutePolicy"`
	// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
	ImportRoutePolicyId pulumi.StringPtrOutput `pulumi:"importRoutePolicyId"`
	// Ingress Acl. ARM resource ID of Access Control Lists.
	IngressAclId pulumi.StringPtrOutput `pulumi:"ingressAclId"`
	// To check whether monitoring of internal network is enabled or not.
	IsMonitoringEnabled pulumi.StringPtrOutput `pulumi:"isMonitoringEnabled"`
	// Maximum transmission unit. Default value is 1500.
	Mtu pulumi.IntPtrOutput `pulumi:"mtu"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Static Route Configuration properties.
	StaticRouteConfiguration InternalNetworkPropertiesResponseStaticRouteConfigurationPtrOutput `pulumi:"staticRouteConfiguration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Vlan identifier. Example: 1001.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
}

// NewInternalNetwork registers a new resource with the given unique name, arguments, and options.
func NewInternalNetwork(ctx *pulumi.Context,
	name string, args *InternalNetworkArgs, opts ...pulumi.ResourceOption) (*InternalNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.L3IsolationDomainName == nil {
		return nil, errors.New("invalid value for required argument 'L3IsolationDomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	if args.BgpConfiguration != nil {
		args.BgpConfiguration = args.BgpConfiguration.ToInternalNetworkPropertiesBgpConfigurationPtrOutput().ApplyT(func(v *InternalNetworkPropertiesBgpConfiguration) *InternalNetworkPropertiesBgpConfiguration {
			return v.Defaults()
		}).(InternalNetworkPropertiesBgpConfigurationPtrOutput)
	}
	if args.Extension == nil {
		args.Extension = pulumi.StringPtr("NoExtension")
	}
	if args.IsMonitoringEnabled == nil {
		args.IsMonitoringEnabled = pulumi.StringPtr("False")
	}
	if args.Mtu == nil {
		args.Mtu = pulumi.IntPtr(1500)
	}
	if args.StaticRouteConfiguration != nil {
		args.StaticRouteConfiguration = args.StaticRouteConfiguration.ToInternalNetworkPropertiesStaticRouteConfigurationPtrOutput().ApplyT(func(v *InternalNetworkPropertiesStaticRouteConfiguration) *InternalNetworkPropertiesStaticRouteConfiguration {
			return v.Defaults()
		}).(InternalNetworkPropertiesStaticRouteConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:InternalNetwork"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230201preview:InternalNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InternalNetwork
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230615:InternalNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternalNetwork gets an existing InternalNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternalNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternalNetworkState, opts ...pulumi.ResourceOption) (*InternalNetwork, error) {
	var resource InternalNetwork
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230615:InternalNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InternalNetwork resources.
type internalNetworkState struct {
}

type InternalNetworkState struct {
}

func (InternalNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*internalNetworkState)(nil)).Elem()
}

type internalNetworkArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// BGP configuration properties.
	BgpConfiguration *InternalNetworkPropertiesBgpConfiguration `pulumi:"bgpConfiguration"`
	// List of Connected IPv4 Subnets.
	ConnectedIPv4Subnets []ConnectedSubnet `pulumi:"connectedIPv4Subnets"`
	// List of connected IPv6 Subnets.
	ConnectedIPv6Subnets []ConnectedSubnet `pulumi:"connectedIPv6Subnets"`
	// Egress Acl. ARM resource ID of Access Control Lists.
	EgressAclId *string `pulumi:"egressAclId"`
	// Export Route Policy either IPv4 or IPv6.
	ExportRoutePolicy *ExportRoutePolicy `pulumi:"exportRoutePolicy"`
	// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
	ExportRoutePolicyId *string `pulumi:"exportRoutePolicyId"`
	// Extension. Example: NoExtension | NPB.
	Extension *string `pulumi:"extension"`
	// Import Route Policy either IPv4 or IPv6.
	ImportRoutePolicy *ImportRoutePolicy `pulumi:"importRoutePolicy"`
	// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
	ImportRoutePolicyId *string `pulumi:"importRoutePolicyId"`
	// Ingress Acl. ARM resource ID of Access Control Lists.
	IngressAclId *string `pulumi:"ingressAclId"`
	// Name of the Internal Network.
	InternalNetworkName *string `pulumi:"internalNetworkName"`
	// To check whether monitoring of internal network is enabled or not.
	IsMonitoringEnabled *string `pulumi:"isMonitoringEnabled"`
	// Name of the L3 Isolation Domain.
	L3IsolationDomainName string `pulumi:"l3IsolationDomainName"`
	// Maximum transmission unit. Default value is 1500.
	Mtu *int `pulumi:"mtu"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Static Route Configuration properties.
	StaticRouteConfiguration *InternalNetworkPropertiesStaticRouteConfiguration `pulumi:"staticRouteConfiguration"`
	// Vlan identifier. Example: 1001.
	VlanId int `pulumi:"vlanId"`
}

// The set of arguments for constructing a InternalNetwork resource.
type InternalNetworkArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// BGP configuration properties.
	BgpConfiguration InternalNetworkPropertiesBgpConfigurationPtrInput
	// List of Connected IPv4 Subnets.
	ConnectedIPv4Subnets ConnectedSubnetArrayInput
	// List of connected IPv6 Subnets.
	ConnectedIPv6Subnets ConnectedSubnetArrayInput
	// Egress Acl. ARM resource ID of Access Control Lists.
	EgressAclId pulumi.StringPtrInput
	// Export Route Policy either IPv4 or IPv6.
	ExportRoutePolicy ExportRoutePolicyPtrInput
	// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
	ExportRoutePolicyId pulumi.StringPtrInput
	// Extension. Example: NoExtension | NPB.
	Extension pulumi.StringPtrInput
	// Import Route Policy either IPv4 or IPv6.
	ImportRoutePolicy ImportRoutePolicyPtrInput
	// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
	ImportRoutePolicyId pulumi.StringPtrInput
	// Ingress Acl. ARM resource ID of Access Control Lists.
	IngressAclId pulumi.StringPtrInput
	// Name of the Internal Network.
	InternalNetworkName pulumi.StringPtrInput
	// To check whether monitoring of internal network is enabled or not.
	IsMonitoringEnabled pulumi.StringPtrInput
	// Name of the L3 Isolation Domain.
	L3IsolationDomainName pulumi.StringInput
	// Maximum transmission unit. Default value is 1500.
	Mtu pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Static Route Configuration properties.
	StaticRouteConfiguration InternalNetworkPropertiesStaticRouteConfigurationPtrInput
	// Vlan identifier. Example: 1001.
	VlanId pulumi.IntInput
}

func (InternalNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internalNetworkArgs)(nil)).Elem()
}

type InternalNetworkInput interface {
	pulumi.Input

	ToInternalNetworkOutput() InternalNetworkOutput
	ToInternalNetworkOutputWithContext(ctx context.Context) InternalNetworkOutput
}

func (*InternalNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**InternalNetwork)(nil)).Elem()
}

func (i *InternalNetwork) ToInternalNetworkOutput() InternalNetworkOutput {
	return i.ToInternalNetworkOutputWithContext(context.Background())
}

func (i *InternalNetwork) ToInternalNetworkOutputWithContext(ctx context.Context) InternalNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternalNetworkOutput)
}

type InternalNetworkOutput struct{ *pulumi.OutputState }

func (InternalNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternalNetwork)(nil)).Elem()
}

func (o InternalNetworkOutput) ToInternalNetworkOutput() InternalNetworkOutput {
	return o
}

func (o InternalNetworkOutput) ToInternalNetworkOutputWithContext(ctx context.Context) InternalNetworkOutput {
	return o
}

// Administrative state of the resource.
func (o InternalNetworkOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringOutput { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o InternalNetworkOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// BGP configuration properties.
func (o InternalNetworkOutput) BgpConfiguration() InternalNetworkPropertiesResponseBgpConfigurationPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) InternalNetworkPropertiesResponseBgpConfigurationPtrOutput {
		return v.BgpConfiguration
	}).(InternalNetworkPropertiesResponseBgpConfigurationPtrOutput)
}

// Configuration state of the resource.
func (o InternalNetworkOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringOutput { return v.ConfigurationState }).(pulumi.StringOutput)
}

// List of Connected IPv4 Subnets.
func (o InternalNetworkOutput) ConnectedIPv4Subnets() ConnectedSubnetResponseArrayOutput {
	return o.ApplyT(func(v *InternalNetwork) ConnectedSubnetResponseArrayOutput { return v.ConnectedIPv4Subnets }).(ConnectedSubnetResponseArrayOutput)
}

// List of connected IPv6 Subnets.
func (o InternalNetworkOutput) ConnectedIPv6Subnets() ConnectedSubnetResponseArrayOutput {
	return o.ApplyT(func(v *InternalNetwork) ConnectedSubnetResponseArrayOutput { return v.ConnectedIPv6Subnets }).(ConnectedSubnetResponseArrayOutput)
}

// Egress Acl. ARM resource ID of Access Control Lists.
func (o InternalNetworkOutput) EgressAclId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringPtrOutput { return v.EgressAclId }).(pulumi.StringPtrOutput)
}

// Export Route Policy either IPv4 or IPv6.
func (o InternalNetworkOutput) ExportRoutePolicy() ExportRoutePolicyResponsePtrOutput {
	return o.ApplyT(func(v *InternalNetwork) ExportRoutePolicyResponsePtrOutput { return v.ExportRoutePolicy }).(ExportRoutePolicyResponsePtrOutput)
}

// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
func (o InternalNetworkOutput) ExportRoutePolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringPtrOutput { return v.ExportRoutePolicyId }).(pulumi.StringPtrOutput)
}

// Extension. Example: NoExtension | NPB.
func (o InternalNetworkOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringPtrOutput { return v.Extension }).(pulumi.StringPtrOutput)
}

// Import Route Policy either IPv4 or IPv6.
func (o InternalNetworkOutput) ImportRoutePolicy() ImportRoutePolicyResponsePtrOutput {
	return o.ApplyT(func(v *InternalNetwork) ImportRoutePolicyResponsePtrOutput { return v.ImportRoutePolicy }).(ImportRoutePolicyResponsePtrOutput)
}

// ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
func (o InternalNetworkOutput) ImportRoutePolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringPtrOutput { return v.ImportRoutePolicyId }).(pulumi.StringPtrOutput)
}

// Ingress Acl. ARM resource ID of Access Control Lists.
func (o InternalNetworkOutput) IngressAclId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringPtrOutput { return v.IngressAclId }).(pulumi.StringPtrOutput)
}

// To check whether monitoring of internal network is enabled or not.
func (o InternalNetworkOutput) IsMonitoringEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringPtrOutput { return v.IsMonitoringEnabled }).(pulumi.StringPtrOutput)
}

// Maximum transmission unit. Default value is 1500.
func (o InternalNetworkOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.IntPtrOutput { return v.Mtu }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o InternalNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o InternalNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Static Route Configuration properties.
func (o InternalNetworkOutput) StaticRouteConfiguration() InternalNetworkPropertiesResponseStaticRouteConfigurationPtrOutput {
	return o.ApplyT(func(v *InternalNetwork) InternalNetworkPropertiesResponseStaticRouteConfigurationPtrOutput {
		return v.StaticRouteConfiguration
	}).(InternalNetworkPropertiesResponseStaticRouteConfigurationPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o InternalNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InternalNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o InternalNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Vlan identifier. Example: 1001.
func (o InternalNetworkOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v *InternalNetwork) pulumi.IntOutput { return v.VlanId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(InternalNetworkOutput{})
}
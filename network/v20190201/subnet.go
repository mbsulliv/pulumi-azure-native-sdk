// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Subnet in a virtual network resource.
type Subnet struct {
	pulumi.CustomResourceState

	// The address prefix for the subnet.
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// List of  address prefixes for the subnet.
	AddressPrefixes pulumi.StringArrayOutput `pulumi:"addressPrefixes"`
	// Gets an array of references to the delegations on the subnet.
	Delegations DelegationResponseArrayOutput `pulumi:"delegations"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// An array of references to interface endpoints
	InterfaceEndpoints InterfaceEndpointResponseArrayOutput `pulumi:"interfaceEndpoints"`
	// Array of IP configuration profiles which reference this subnet.
	IpConfigurationProfiles IPConfigurationProfileResponseArrayOutput `pulumi:"ipConfigurationProfiles"`
	// Gets an array of references to the network interface IP configurations using subnet.
	IpConfigurations IPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Nat gateway associated with this subnet.
	NatGateway SubResourceResponsePtrOutput `pulumi:"natGateway"`
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupResponsePtrOutput `pulumi:"networkSecurityGroup"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// A read-only string identifying the intention of use for this subnet based on delegations and other user-defined properties.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// Gets an array of references to the external resources using subnet.
	ResourceNavigationLinks ResourceNavigationLinkResponseArrayOutput `pulumi:"resourceNavigationLinks"`
	// The reference of the RouteTable resource.
	RouteTable RouteTableResponsePtrOutput `pulumi:"routeTable"`
	// Gets an array of references to services injecting into this subnet.
	ServiceAssociationLinks ServiceAssociationLinkResponseArrayOutput `pulumi:"serviceAssociationLinks"`
	// An array of service endpoint policies.
	ServiceEndpointPolicies ServiceEndpointPolicyResponseArrayOutput `pulumi:"serviceEndpointPolicies"`
	// An array of service endpoints.
	ServiceEndpoints ServiceEndpointPropertiesFormatResponseArrayOutput `pulumi:"serviceEndpoints"`
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOption) (*Subnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:Subnet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Subnet
	err := ctx.RegisterResource("azure-native:network/v20190201:Subnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnet gets an existing Subnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetState, opts ...pulumi.ResourceOption) (*Subnet, error) {
	var resource Subnet
	err := ctx.ReadResource("azure-native:network/v20190201:Subnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnet resources.
type subnetState struct {
}

type SubnetState struct {
}

func (SubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetState)(nil)).Elem()
}

type subnetArgs struct {
	// The address prefix for the subnet.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// List of  address prefixes for the subnet.
	AddressPrefixes []string `pulumi:"addressPrefixes"`
	// Gets an array of references to the delegations on the subnet.
	Delegations []Delegation `pulumi:"delegations"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Nat gateway associated with this subnet.
	NatGateway *SubResource `pulumi:"natGateway"`
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroup `pulumi:"networkSecurityGroup"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets an array of references to the external resources using subnet.
	ResourceNavigationLinks []ResourceNavigationLink `pulumi:"resourceNavigationLinks"`
	// The reference of the RouteTable resource.
	RouteTable *RouteTable `pulumi:"routeTable"`
	// Gets an array of references to services injecting into this subnet.
	ServiceAssociationLinks []ServiceAssociationLink `pulumi:"serviceAssociationLinks"`
	// An array of service endpoint policies.
	ServiceEndpointPolicies []ServiceEndpointPolicy `pulumi:"serviceEndpointPolicies"`
	// An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat `pulumi:"serviceEndpoints"`
	// The name of the subnet.
	SubnetName *string `pulumi:"subnetName"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// The address prefix for the subnet.
	AddressPrefix pulumi.StringPtrInput
	// List of  address prefixes for the subnet.
	AddressPrefixes pulumi.StringArrayInput
	// Gets an array of references to the delegations on the subnet.
	Delegations DelegationArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Nat gateway associated with this subnet.
	NatGateway SubResourcePtrInput
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Gets an array of references to the external resources using subnet.
	ResourceNavigationLinks ResourceNavigationLinkArrayInput
	// The reference of the RouteTable resource.
	RouteTable RouteTablePtrInput
	// Gets an array of references to services injecting into this subnet.
	ServiceAssociationLinks ServiceAssociationLinkArrayInput
	// An array of service endpoint policies.
	ServiceEndpointPolicies ServiceEndpointPolicyArrayInput
	// An array of service endpoints.
	ServiceEndpoints ServiceEndpointPropertiesFormatArrayInput
	// The name of the subnet.
	SubnetName pulumi.StringPtrInput
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringInput
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetArgs)(nil)).Elem()
}

type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(ctx context.Context) SubnetOutput
}

func (*Subnet) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (i *Subnet) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i *Subnet) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

func (i *Subnet) ToOutput(ctx context.Context) pulumix.Output[*Subnet] {
	return pulumix.Output[*Subnet]{
		OutputState: i.ToSubnetOutputWithContext(ctx).OutputState,
	}
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) ToOutput(ctx context.Context) pulumix.Output[*Subnet] {
	return pulumix.Output[*Subnet]{
		OutputState: o.OutputState,
	}
}

// The address prefix for the subnet.
func (o SubnetOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// List of  address prefixes for the subnet.
func (o SubnetOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringArrayOutput { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

// Gets an array of references to the delegations on the subnet.
func (o SubnetOutput) Delegations() DelegationResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) DelegationResponseArrayOutput { return v.Delegations }).(DelegationResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o SubnetOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// An array of references to interface endpoints
func (o SubnetOutput) InterfaceEndpoints() InterfaceEndpointResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) InterfaceEndpointResponseArrayOutput { return v.InterfaceEndpoints }).(InterfaceEndpointResponseArrayOutput)
}

// Array of IP configuration profiles which reference this subnet.
func (o SubnetOutput) IpConfigurationProfiles() IPConfigurationProfileResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) IPConfigurationProfileResponseArrayOutput { return v.IpConfigurationProfiles }).(IPConfigurationProfileResponseArrayOutput)
}

// Gets an array of references to the network interface IP configurations using subnet.
func (o SubnetOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) IPConfigurationResponseArrayOutput { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o SubnetOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Nat gateway associated with this subnet.
func (o SubnetOutput) NatGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *Subnet) SubResourceResponsePtrOutput { return v.NatGateway }).(SubResourceResponsePtrOutput)
}

// The reference of the NetworkSecurityGroup resource.
func (o SubnetOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *Subnet) NetworkSecurityGroupResponsePtrOutput { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

// The provisioning state of the resource.
func (o SubnetOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// A read-only string identifying the intention of use for this subnet based on delegations and other user-defined properties.
func (o SubnetOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.Purpose }).(pulumi.StringOutput)
}

// Gets an array of references to the external resources using subnet.
func (o SubnetOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) ResourceNavigationLinkResponseArrayOutput { return v.ResourceNavigationLinks }).(ResourceNavigationLinkResponseArrayOutput)
}

// The reference of the RouteTable resource.
func (o SubnetOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v *Subnet) RouteTableResponsePtrOutput { return v.RouteTable }).(RouteTableResponsePtrOutput)
}

// Gets an array of references to services injecting into this subnet.
func (o SubnetOutput) ServiceAssociationLinks() ServiceAssociationLinkResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) ServiceAssociationLinkResponseArrayOutput { return v.ServiceAssociationLinks }).(ServiceAssociationLinkResponseArrayOutput)
}

// An array of service endpoint policies.
func (o SubnetOutput) ServiceEndpointPolicies() ServiceEndpointPolicyResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) ServiceEndpointPolicyResponseArrayOutput { return v.ServiceEndpointPolicies }).(ServiceEndpointPolicyResponseArrayOutput)
}

// An array of service endpoints.
func (o SubnetOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) ServiceEndpointPropertiesFormatResponseArrayOutput { return v.ServiceEndpoints }).(ServiceEndpointPropertiesFormatResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SubnetOutput{})
}

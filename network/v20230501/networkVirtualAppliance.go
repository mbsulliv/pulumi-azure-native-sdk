// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NetworkVirtualAppliance Resource.
type NetworkVirtualAppliance struct {
	pulumi.CustomResourceState

	// Details required for Additional Network Interface.
	AdditionalNics VirtualApplianceAdditionalNicPropertiesResponseArrayOutput `pulumi:"additionalNics"`
	// Address Prefix.
	AddressPrefix pulumi.StringOutput `pulumi:"addressPrefix"`
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs pulumi.StringArrayOutput `pulumi:"bootStrapConfigurationBlobs"`
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration pulumi.StringPtrOutput `pulumi:"cloudInitConfiguration"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs pulumi.StringArrayOutput `pulumi:"cloudInitConfigurationBlobs"`
	// The delegation for the Virtual Appliance
	Delegation DelegationPropertiesResponsePtrOutput `pulumi:"delegation"`
	// The deployment type. PartnerManaged for the SaaS NVA
	DeploymentType pulumi.StringOutput `pulumi:"deploymentType"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The service principal that has read access to cloud-init and config blob.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// List of references to InboundSecurityRules.
	InboundSecurityRules SubResourceResponseArrayOutput `pulumi:"inboundSecurityRules"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network Virtual Appliance SKU.
	NvaSku VirtualApplianceSkuPropertiesResponsePtrOutput `pulumi:"nvaSku"`
	// The delegation for the Virtual Appliance
	PartnerManagedResource PartnerManagedResourcePropertiesResponsePtrOutput `pulumi:"partnerManagedResource"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Public key for SSH login.
	SshPublicKey pulumi.StringPtrOutput `pulumi:"sshPublicKey"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// VirtualAppliance ASN. Microsoft private, public and IANA reserved ASN are not supported.
	VirtualApplianceAsn pulumi.Float64PtrOutput `pulumi:"virtualApplianceAsn"`
	// List of Virtual Appliance Network Interfaces.
	VirtualApplianceNics VirtualApplianceNicPropertiesResponseArrayOutput `pulumi:"virtualApplianceNics"`
	// List of references to VirtualApplianceSite.
	VirtualApplianceSites SubResourceResponseArrayOutput `pulumi:"virtualApplianceSites"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
}

// NewNetworkVirtualAppliance registers a new resource with the given unique name, arguments, and options.
func NewNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, args *NetworkVirtualApplianceArgs, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:NetworkVirtualAppliance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkVirtualAppliance
	err := ctx.RegisterResource("azure-native:network/v20230501:NetworkVirtualAppliance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkVirtualAppliance gets an existing NetworkVirtualAppliance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkVirtualApplianceState, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	var resource NetworkVirtualAppliance
	err := ctx.ReadResource("azure-native:network/v20230501:NetworkVirtualAppliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkVirtualAppliance resources.
type networkVirtualApplianceState struct {
}

type NetworkVirtualApplianceState struct {
}

func (NetworkVirtualApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceState)(nil)).Elem()
}

type networkVirtualApplianceArgs struct {
	// Details required for Additional Network Interface.
	AdditionalNics []VirtualApplianceAdditionalNicProperties `pulumi:"additionalNics"`
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs []string `pulumi:"bootStrapConfigurationBlobs"`
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration *string `pulumi:"cloudInitConfiguration"`
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs []string `pulumi:"cloudInitConfigurationBlobs"`
	// The delegation for the Virtual Appliance
	Delegation *DelegationProperties `pulumi:"delegation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The service principal that has read access to cloud-init and config blob.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of Network Virtual Appliance.
	NetworkVirtualApplianceName *string `pulumi:"networkVirtualApplianceName"`
	// Network Virtual Appliance SKU.
	NvaSku *VirtualApplianceSkuProperties `pulumi:"nvaSku"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Public key for SSH login.
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// VirtualAppliance ASN. Microsoft private, public and IANA reserved ASN are not supported.
	VirtualApplianceAsn *float64 `pulumi:"virtualApplianceAsn"`
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub *SubResource `pulumi:"virtualHub"`
}

// The set of arguments for constructing a NetworkVirtualAppliance resource.
type NetworkVirtualApplianceArgs struct {
	// Details required for Additional Network Interface.
	AdditionalNics VirtualApplianceAdditionalNicPropertiesArrayInput
	// BootStrapConfigurationBlobs storage URLs.
	BootStrapConfigurationBlobs pulumi.StringArrayInput
	// CloudInitConfiguration string in plain text.
	CloudInitConfiguration pulumi.StringPtrInput
	// CloudInitConfigurationBlob storage URLs.
	CloudInitConfigurationBlobs pulumi.StringArrayInput
	// The delegation for the Virtual Appliance
	Delegation DelegationPropertiesPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The service principal that has read access to cloud-init and config blob.
	Identity ManagedServiceIdentityPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of Network Virtual Appliance.
	NetworkVirtualApplianceName pulumi.StringPtrInput
	// Network Virtual Appliance SKU.
	NvaSku VirtualApplianceSkuPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Public key for SSH login.
	SshPublicKey pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// VirtualAppliance ASN. Microsoft private, public and IANA reserved ASN are not supported.
	VirtualApplianceAsn pulumi.Float64PtrInput
	// The Virtual Hub where Network Virtual Appliance is being deployed.
	VirtualHub SubResourcePtrInput
}

func (NetworkVirtualApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceArgs)(nil)).Elem()
}

type NetworkVirtualApplianceInput interface {
	pulumi.Input

	ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput
	ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput
}

func (*NetworkVirtualAppliance) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkVirtualAppliance)(nil)).Elem()
}

func (i *NetworkVirtualAppliance) ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput {
	return i.ToNetworkVirtualApplianceOutputWithContext(context.Background())
}

func (i *NetworkVirtualAppliance) ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkVirtualApplianceOutput)
}

type NetworkVirtualApplianceOutput struct{ *pulumi.OutputState }

func (NetworkVirtualApplianceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkVirtualAppliance)(nil)).Elem()
}

func (o NetworkVirtualApplianceOutput) ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput {
	return o
}

func (o NetworkVirtualApplianceOutput) ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput {
	return o
}

// Details required for Additional Network Interface.
func (o NetworkVirtualApplianceOutput) AdditionalNics() VirtualApplianceAdditionalNicPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) VirtualApplianceAdditionalNicPropertiesResponseArrayOutput {
		return v.AdditionalNics
	}).(VirtualApplianceAdditionalNicPropertiesResponseArrayOutput)
}

// Address Prefix.
func (o NetworkVirtualApplianceOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.AddressPrefix }).(pulumi.StringOutput)
}

// BootStrapConfigurationBlobs storage URLs.
func (o NetworkVirtualApplianceOutput) BootStrapConfigurationBlobs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringArrayOutput { return v.BootStrapConfigurationBlobs }).(pulumi.StringArrayOutput)
}

// CloudInitConfiguration string in plain text.
func (o NetworkVirtualApplianceOutput) CloudInitConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringPtrOutput { return v.CloudInitConfiguration }).(pulumi.StringPtrOutput)
}

// CloudInitConfigurationBlob storage URLs.
func (o NetworkVirtualApplianceOutput) CloudInitConfigurationBlobs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringArrayOutput { return v.CloudInitConfigurationBlobs }).(pulumi.StringArrayOutput)
}

// The delegation for the Virtual Appliance
func (o NetworkVirtualApplianceOutput) Delegation() DelegationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) DelegationPropertiesResponsePtrOutput { return v.Delegation }).(DelegationPropertiesResponsePtrOutput)
}

// The deployment type. PartnerManaged for the SaaS NVA
func (o NetworkVirtualApplianceOutput) DeploymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.DeploymentType }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o NetworkVirtualApplianceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The service principal that has read access to cloud-init and config blob.
func (o NetworkVirtualApplianceOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// List of references to InboundSecurityRules.
func (o NetworkVirtualApplianceOutput) InboundSecurityRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) SubResourceResponseArrayOutput { return v.InboundSecurityRules }).(SubResourceResponseArrayOutput)
}

// Resource location.
func (o NetworkVirtualApplianceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NetworkVirtualApplianceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network Virtual Appliance SKU.
func (o NetworkVirtualApplianceOutput) NvaSku() VirtualApplianceSkuPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) VirtualApplianceSkuPropertiesResponsePtrOutput { return v.NvaSku }).(VirtualApplianceSkuPropertiesResponsePtrOutput)
}

// The delegation for the Virtual Appliance
func (o NetworkVirtualApplianceOutput) PartnerManagedResource() PartnerManagedResourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) PartnerManagedResourcePropertiesResponsePtrOutput {
		return v.PartnerManagedResource
	}).(PartnerManagedResourcePropertiesResponsePtrOutput)
}

// The provisioning state of the resource.
func (o NetworkVirtualApplianceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Public key for SSH login.
func (o NetworkVirtualApplianceOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringPtrOutput { return v.SshPublicKey }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o NetworkVirtualApplianceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NetworkVirtualApplianceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// VirtualAppliance ASN. Microsoft private, public and IANA reserved ASN are not supported.
func (o NetworkVirtualApplianceOutput) VirtualApplianceAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) pulumi.Float64PtrOutput { return v.VirtualApplianceAsn }).(pulumi.Float64PtrOutput)
}

// List of Virtual Appliance Network Interfaces.
func (o NetworkVirtualApplianceOutput) VirtualApplianceNics() VirtualApplianceNicPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) VirtualApplianceNicPropertiesResponseArrayOutput {
		return v.VirtualApplianceNics
	}).(VirtualApplianceNicPropertiesResponseArrayOutput)
}

// List of references to VirtualApplianceSite.
func (o NetworkVirtualApplianceOutput) VirtualApplianceSites() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) SubResourceResponseArrayOutput { return v.VirtualApplianceSites }).(SubResourceResponseArrayOutput)
}

// The Virtual Hub where Network Virtual Appliance is being deployed.
func (o NetworkVirtualApplianceOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkVirtualAppliance) SubResourceResponsePtrOutput { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkVirtualApplianceOutput{})
}

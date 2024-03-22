// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A private cloud resource
type PrivateCloud struct {
	pulumi.CustomResourceState

	// Properties describing how the cloud is distributed across availability zones
	Availability AvailabilityPropertiesResponsePtrOutput `pulumi:"availability"`
	// An ExpressRoute Circuit
	Circuit CircuitResponsePtrOutput `pulumi:"circuit"`
	// Customer managed key encryption, can be enabled or disabled
	Encryption EncryptionResponsePtrOutput `pulumi:"encryption"`
	// The endpoints
	Endpoints EndpointsResponseOutput `pulumi:"endpoints"`
	// Array of cloud link IDs from other clouds that connect to this one
	ExternalCloudLinks pulumi.StringArrayOutput `pulumi:"externalCloudLinks"`
	// The identity of the private cloud, if configured.
	Identity PrivateCloudIdentityResponsePtrOutput `pulumi:"identity"`
	// vCenter Single Sign On Identity Sources
	IdentitySources IdentitySourceResponseArrayOutput `pulumi:"identitySources"`
	// Connectivity to internet is enabled or disabled
	Internet pulumi.StringPtrOutput `pulumi:"internet"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The default cluster used for management
	ManagementCluster ManagementClusterResponseOutput `pulumi:"managementCluster"`
	// Network used to access vCenter Server and NSX-T Manager
	ManagementNetwork pulumi.StringOutput `pulumi:"managementNetwork"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The block of addresses should be unique across VNet in your subscription as well as on-premise. Make sure the CIDR format is conformed to (A.B.C.D/X) where A,B,C,D are between 0 and 255, and X is between 0 and 22
	NetworkBlock pulumi.StringOutput `pulumi:"networkBlock"`
	// Flag to indicate whether the private cloud has the quota for provisioned NSX Public IP count raised from 64 to 1024
	NsxPublicIpQuotaRaised pulumi.StringOutput `pulumi:"nsxPublicIpQuotaRaised"`
	// Thumbprint of the NSX-T Manager SSL certificate
	NsxtCertificateThumbprint pulumi.StringOutput `pulumi:"nsxtCertificateThumbprint"`
	// Optionally, set the NSX-T Manager password when the private cloud is created
	NsxtPassword pulumi.StringPtrOutput `pulumi:"nsxtPassword"`
	// Used for virtual machine cold migration, cloning, and snapshot migration
	ProvisioningNetwork pulumi.StringOutput `pulumi:"provisioningNetwork"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// A secondary expressRoute circuit from a separate AZ. Only present in a stretched private cloud
	SecondaryCircuit CircuitResponsePtrOutput `pulumi:"secondaryCircuit"`
	// The private cloud SKU
	Sku SkuResponseOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Thumbprint of the vCenter Server SSL certificate
	VcenterCertificateThumbprint pulumi.StringOutput `pulumi:"vcenterCertificateThumbprint"`
	// Optionally, set the vCenter admin password when the private cloud is created
	VcenterPassword pulumi.StringPtrOutput `pulumi:"vcenterPassword"`
	// Used for live migration of virtual machines
	VmotionNetwork pulumi.StringOutput `pulumi:"vmotionNetwork"`
}

// NewPrivateCloud registers a new resource with the given unique name, arguments, and options.
func NewPrivateCloud(ctx *pulumi.Context,
	name string, args *PrivateCloudArgs, opts ...pulumi.ResourceOption) (*PrivateCloud, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementCluster == nil {
		return nil, errors.New("invalid value for required argument 'ManagementCluster'")
	}
	if args.NetworkBlock == nil {
		return nil, errors.New("invalid value for required argument 'NetworkBlock'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.Internet == nil {
		args.Internet = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200320:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:PrivateCloud"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230901:PrivateCloud"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateCloud
	err := ctx.RegisterResource("azure-native:avs/v20220501:PrivateCloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateCloud gets an existing PrivateCloud resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateCloudState, opts ...pulumi.ResourceOption) (*PrivateCloud, error) {
	var resource PrivateCloud
	err := ctx.ReadResource("azure-native:avs/v20220501:PrivateCloud", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateCloud resources.
type privateCloudState struct {
}

type PrivateCloudState struct {
}

func (PrivateCloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateCloudState)(nil)).Elem()
}

type privateCloudArgs struct {
	// Properties describing how the cloud is distributed across availability zones
	Availability *AvailabilityProperties `pulumi:"availability"`
	// Customer managed key encryption, can be enabled or disabled
	Encryption *Encryption `pulumi:"encryption"`
	// The identity of the private cloud, if configured.
	Identity *PrivateCloudIdentity `pulumi:"identity"`
	// vCenter Single Sign On Identity Sources
	IdentitySources []IdentitySource `pulumi:"identitySources"`
	// Connectivity to internet is enabled or disabled
	Internet *string `pulumi:"internet"`
	// Resource location
	Location *string `pulumi:"location"`
	// The default cluster used for management
	ManagementCluster ManagementCluster `pulumi:"managementCluster"`
	// The block of addresses should be unique across VNet in your subscription as well as on-premise. Make sure the CIDR format is conformed to (A.B.C.D/X) where A,B,C,D are between 0 and 255, and X is between 0 and 22
	NetworkBlock string `pulumi:"networkBlock"`
	// Optionally, set the NSX-T Manager password when the private cloud is created
	NsxtPassword *string `pulumi:"nsxtPassword"`
	// Name of the private cloud
	PrivateCloudName *string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The private cloud SKU
	Sku Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Optionally, set the vCenter admin password when the private cloud is created
	VcenterPassword *string `pulumi:"vcenterPassword"`
}

// The set of arguments for constructing a PrivateCloud resource.
type PrivateCloudArgs struct {
	// Properties describing how the cloud is distributed across availability zones
	Availability AvailabilityPropertiesPtrInput
	// Customer managed key encryption, can be enabled or disabled
	Encryption EncryptionPtrInput
	// The identity of the private cloud, if configured.
	Identity PrivateCloudIdentityPtrInput
	// vCenter Single Sign On Identity Sources
	IdentitySources IdentitySourceArrayInput
	// Connectivity to internet is enabled or disabled
	Internet pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The default cluster used for management
	ManagementCluster ManagementClusterInput
	// The block of addresses should be unique across VNet in your subscription as well as on-premise. Make sure the CIDR format is conformed to (A.B.C.D/X) where A,B,C,D are between 0 and 255, and X is between 0 and 22
	NetworkBlock pulumi.StringInput
	// Optionally, set the NSX-T Manager password when the private cloud is created
	NsxtPassword pulumi.StringPtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The private cloud SKU
	Sku SkuInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Optionally, set the vCenter admin password when the private cloud is created
	VcenterPassword pulumi.StringPtrInput
}

func (PrivateCloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateCloudArgs)(nil)).Elem()
}

type PrivateCloudInput interface {
	pulumi.Input

	ToPrivateCloudOutput() PrivateCloudOutput
	ToPrivateCloudOutputWithContext(ctx context.Context) PrivateCloudOutput
}

func (*PrivateCloud) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloud)(nil)).Elem()
}

func (i *PrivateCloud) ToPrivateCloudOutput() PrivateCloudOutput {
	return i.ToPrivateCloudOutputWithContext(context.Background())
}

func (i *PrivateCloud) ToPrivateCloudOutputWithContext(ctx context.Context) PrivateCloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudOutput)
}

type PrivateCloudOutput struct{ *pulumi.OutputState }

func (PrivateCloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloud)(nil)).Elem()
}

func (o PrivateCloudOutput) ToPrivateCloudOutput() PrivateCloudOutput {
	return o
}

func (o PrivateCloudOutput) ToPrivateCloudOutputWithContext(ctx context.Context) PrivateCloudOutput {
	return o
}

// Properties describing how the cloud is distributed across availability zones
func (o PrivateCloudOutput) Availability() AvailabilityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PrivateCloud) AvailabilityPropertiesResponsePtrOutput { return v.Availability }).(AvailabilityPropertiesResponsePtrOutput)
}

// An ExpressRoute Circuit
func (o PrivateCloudOutput) Circuit() CircuitResponsePtrOutput {
	return o.ApplyT(func(v *PrivateCloud) CircuitResponsePtrOutput { return v.Circuit }).(CircuitResponsePtrOutput)
}

// Customer managed key encryption, can be enabled or disabled
func (o PrivateCloudOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *PrivateCloud) EncryptionResponsePtrOutput { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// The endpoints
func (o PrivateCloudOutput) Endpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v *PrivateCloud) EndpointsResponseOutput { return v.Endpoints }).(EndpointsResponseOutput)
}

// Array of cloud link IDs from other clouds that connect to this one
func (o PrivateCloudOutput) ExternalCloudLinks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringArrayOutput { return v.ExternalCloudLinks }).(pulumi.StringArrayOutput)
}

// The identity of the private cloud, if configured.
func (o PrivateCloudOutput) Identity() PrivateCloudIdentityResponsePtrOutput {
	return o.ApplyT(func(v *PrivateCloud) PrivateCloudIdentityResponsePtrOutput { return v.Identity }).(PrivateCloudIdentityResponsePtrOutput)
}

// vCenter Single Sign On Identity Sources
func (o PrivateCloudOutput) IdentitySources() IdentitySourceResponseArrayOutput {
	return o.ApplyT(func(v *PrivateCloud) IdentitySourceResponseArrayOutput { return v.IdentitySources }).(IdentitySourceResponseArrayOutput)
}

// Connectivity to internet is enabled or disabled
func (o PrivateCloudOutput) Internet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringPtrOutput { return v.Internet }).(pulumi.StringPtrOutput)
}

// Resource location
func (o PrivateCloudOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The default cluster used for management
func (o PrivateCloudOutput) ManagementCluster() ManagementClusterResponseOutput {
	return o.ApplyT(func(v *PrivateCloud) ManagementClusterResponseOutput { return v.ManagementCluster }).(ManagementClusterResponseOutput)
}

// Network used to access vCenter Server and NSX-T Manager
func (o PrivateCloudOutput) ManagementNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.ManagementNetwork }).(pulumi.StringOutput)
}

// Resource name.
func (o PrivateCloudOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The block of addresses should be unique across VNet in your subscription as well as on-premise. Make sure the CIDR format is conformed to (A.B.C.D/X) where A,B,C,D are between 0 and 255, and X is between 0 and 22
func (o PrivateCloudOutput) NetworkBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.NetworkBlock }).(pulumi.StringOutput)
}

// Flag to indicate whether the private cloud has the quota for provisioned NSX Public IP count raised from 64 to 1024
func (o PrivateCloudOutput) NsxPublicIpQuotaRaised() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.NsxPublicIpQuotaRaised }).(pulumi.StringOutput)
}

// Thumbprint of the NSX-T Manager SSL certificate
func (o PrivateCloudOutput) NsxtCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.NsxtCertificateThumbprint }).(pulumi.StringOutput)
}

// Optionally, set the NSX-T Manager password when the private cloud is created
func (o PrivateCloudOutput) NsxtPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringPtrOutput { return v.NsxtPassword }).(pulumi.StringPtrOutput)
}

// Used for virtual machine cold migration, cloning, and snapshot migration
func (o PrivateCloudOutput) ProvisioningNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.ProvisioningNetwork }).(pulumi.StringOutput)
}

// The provisioning state
func (o PrivateCloudOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// A secondary expressRoute circuit from a separate AZ. Only present in a stretched private cloud
func (o PrivateCloudOutput) SecondaryCircuit() CircuitResponsePtrOutput {
	return o.ApplyT(func(v *PrivateCloud) CircuitResponsePtrOutput { return v.SecondaryCircuit }).(CircuitResponsePtrOutput)
}

// The private cloud SKU
func (o PrivateCloudOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *PrivateCloud) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Resource tags
func (o PrivateCloudOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o PrivateCloudOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Thumbprint of the vCenter Server SSL certificate
func (o PrivateCloudOutput) VcenterCertificateThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.VcenterCertificateThumbprint }).(pulumi.StringOutput)
}

// Optionally, set the vCenter admin password when the private cloud is created
func (o PrivateCloudOutput) VcenterPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringPtrOutput { return v.VcenterPassword }).(pulumi.StringPtrOutput)
}

// Used for live migration of virtual machines
func (o PrivateCloudOutput) VmotionNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateCloud) pulumi.StringOutput { return v.VmotionNetwork }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateCloudOutput{})
}

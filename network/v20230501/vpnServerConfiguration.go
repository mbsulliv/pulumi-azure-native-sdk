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

// VpnServerConfiguration Resource.
type VpnServerConfiguration struct {
	pulumi.CustomResourceState

	// The set of aad vpn authentication parameters.
	AadAuthenticationParameters AadAuthenticationParametersResponsePtrOutput `pulumi:"aadAuthenticationParameters"`
	// List of all VpnServerConfigurationPolicyGroups.
	ConfigurationPolicyGroups VpnServerConfigurationPolicyGroupResponseArrayOutput `pulumi:"configurationPolicyGroups"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of references to P2SVpnGateways.
	P2SVpnGateways P2SVpnGatewayResponseArrayOutput `pulumi:"p2SVpnGateways"`
	// The provisioning state of the VpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Radius client root certificate of VpnServerConfiguration.
	RadiusClientRootCertificates VpnServerConfigRadiusClientRootCertificateResponseArrayOutput `pulumi:"radiusClientRootCertificates"`
	// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress pulumi.StringPtrOutput `pulumi:"radiusServerAddress"`
	// Radius Server root certificate of VpnServerConfiguration.
	RadiusServerRootCertificates VpnServerConfigRadiusServerRootCertificateResponseArrayOutput `pulumi:"radiusServerRootCertificates"`
	// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret pulumi.StringPtrOutput `pulumi:"radiusServerSecret"`
	// Multiple Radius Server configuration for VpnServerConfiguration.
	RadiusServers RadiusServerResponseArrayOutput `pulumi:"radiusServers"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// VPN authentication types for the VpnServerConfiguration.
	VpnAuthenticationTypes pulumi.StringArrayOutput `pulumi:"vpnAuthenticationTypes"`
	// VpnClientIpsecPolicies for VpnServerConfiguration.
	VpnClientIpsecPolicies IpsecPolicyResponseArrayOutput `pulumi:"vpnClientIpsecPolicies"`
	// VPN client revoked certificate of VpnServerConfiguration.
	VpnClientRevokedCertificates VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput `pulumi:"vpnClientRevokedCertificates"`
	// VPN client root certificate of VpnServerConfiguration.
	VpnClientRootCertificates VpnServerConfigVpnClientRootCertificateResponseArrayOutput `pulumi:"vpnClientRootCertificates"`
	// VPN protocols for the VpnServerConfiguration.
	VpnProtocols pulumi.StringArrayOutput `pulumi:"vpnProtocols"`
}

// NewVpnServerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewVpnServerConfiguration(ctx *pulumi.Context,
	name string, args *VpnServerConfigurationArgs, opts ...pulumi.ResourceOption) (*VpnServerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:VpnServerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VpnServerConfiguration
	err := ctx.RegisterResource("azure-native:network/v20230501:VpnServerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnServerConfiguration gets an existing VpnServerConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnServerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnServerConfigurationState, opts ...pulumi.ResourceOption) (*VpnServerConfiguration, error) {
	var resource VpnServerConfiguration
	err := ctx.ReadResource("azure-native:network/v20230501:VpnServerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnServerConfiguration resources.
type vpnServerConfigurationState struct {
}

type VpnServerConfigurationState struct {
}

func (VpnServerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnServerConfigurationState)(nil)).Elem()
}

type vpnServerConfigurationArgs struct {
	// The set of aad vpn authentication parameters.
	AadAuthenticationParameters *AadAuthenticationParameters `pulumi:"aadAuthenticationParameters"`
	// List of all VpnServerConfigurationPolicyGroups.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	ConfigurationPolicyGroups []VpnServerConfigurationPolicyGroup `pulumi:"configurationPolicyGroups"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the VpnServerConfiguration that is unique within a resource group.
	Name *string `pulumi:"name"`
	// Radius client root certificate of VpnServerConfiguration.
	RadiusClientRootCertificates []VpnServerConfigRadiusClientRootCertificate `pulumi:"radiusClientRootCertificates"`
	// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress *string `pulumi:"radiusServerAddress"`
	// Radius Server root certificate of VpnServerConfiguration.
	RadiusServerRootCertificates []VpnServerConfigRadiusServerRootCertificate `pulumi:"radiusServerRootCertificates"`
	// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret *string `pulumi:"radiusServerSecret"`
	// Multiple Radius Server configuration for VpnServerConfiguration.
	RadiusServers []RadiusServer `pulumi:"radiusServers"`
	// The resource group name of the VpnServerConfiguration.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// VPN authentication types for the VpnServerConfiguration.
	VpnAuthenticationTypes []string `pulumi:"vpnAuthenticationTypes"`
	// VpnClientIpsecPolicies for VpnServerConfiguration.
	VpnClientIpsecPolicies []IpsecPolicy `pulumi:"vpnClientIpsecPolicies"`
	// VPN client revoked certificate of VpnServerConfiguration.
	VpnClientRevokedCertificates []VpnServerConfigVpnClientRevokedCertificate `pulumi:"vpnClientRevokedCertificates"`
	// VPN client root certificate of VpnServerConfiguration.
	VpnClientRootCertificates []VpnServerConfigVpnClientRootCertificate `pulumi:"vpnClientRootCertificates"`
	// VPN protocols for the VpnServerConfiguration.
	VpnProtocols []string `pulumi:"vpnProtocols"`
	// The name of the VpnServerConfiguration being created or updated.
	VpnServerConfigurationName *string `pulumi:"vpnServerConfigurationName"`
}

// The set of arguments for constructing a VpnServerConfiguration resource.
type VpnServerConfigurationArgs struct {
	// The set of aad vpn authentication parameters.
	AadAuthenticationParameters AadAuthenticationParametersPtrInput
	// List of all VpnServerConfigurationPolicyGroups.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	ConfigurationPolicyGroups VpnServerConfigurationPolicyGroupArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the VpnServerConfiguration that is unique within a resource group.
	Name pulumi.StringPtrInput
	// Radius client root certificate of VpnServerConfiguration.
	RadiusClientRootCertificates VpnServerConfigRadiusClientRootCertificateArrayInput
	// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress pulumi.StringPtrInput
	// Radius Server root certificate of VpnServerConfiguration.
	RadiusServerRootCertificates VpnServerConfigRadiusServerRootCertificateArrayInput
	// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret pulumi.StringPtrInput
	// Multiple Radius Server configuration for VpnServerConfiguration.
	RadiusServers RadiusServerArrayInput
	// The resource group name of the VpnServerConfiguration.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// VPN authentication types for the VpnServerConfiguration.
	VpnAuthenticationTypes pulumi.StringArrayInput
	// VpnClientIpsecPolicies for VpnServerConfiguration.
	VpnClientIpsecPolicies IpsecPolicyArrayInput
	// VPN client revoked certificate of VpnServerConfiguration.
	VpnClientRevokedCertificates VpnServerConfigVpnClientRevokedCertificateArrayInput
	// VPN client root certificate of VpnServerConfiguration.
	VpnClientRootCertificates VpnServerConfigVpnClientRootCertificateArrayInput
	// VPN protocols for the VpnServerConfiguration.
	VpnProtocols pulumi.StringArrayInput
	// The name of the VpnServerConfiguration being created or updated.
	VpnServerConfigurationName pulumi.StringPtrInput
}

func (VpnServerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnServerConfigurationArgs)(nil)).Elem()
}

type VpnServerConfigurationInput interface {
	pulumi.Input

	ToVpnServerConfigurationOutput() VpnServerConfigurationOutput
	ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput
}

func (*VpnServerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnServerConfiguration)(nil)).Elem()
}

func (i *VpnServerConfiguration) ToVpnServerConfigurationOutput() VpnServerConfigurationOutput {
	return i.ToVpnServerConfigurationOutputWithContext(context.Background())
}

func (i *VpnServerConfiguration) ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationOutput)
}

type VpnServerConfigurationOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnServerConfiguration)(nil)).Elem()
}

func (o VpnServerConfigurationOutput) ToVpnServerConfigurationOutput() VpnServerConfigurationOutput {
	return o
}

func (o VpnServerConfigurationOutput) ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput {
	return o
}

// The set of aad vpn authentication parameters.
func (o VpnServerConfigurationOutput) AadAuthenticationParameters() AadAuthenticationParametersResponsePtrOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) AadAuthenticationParametersResponsePtrOutput {
		return v.AadAuthenticationParameters
	}).(AadAuthenticationParametersResponsePtrOutput)
}

// List of all VpnServerConfigurationPolicyGroups.
func (o VpnServerConfigurationOutput) ConfigurationPolicyGroups() VpnServerConfigurationPolicyGroupResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigurationPolicyGroupResponseArrayOutput {
		return v.ConfigurationPolicyGroups
	}).(VpnServerConfigurationPolicyGroupResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VpnServerConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o VpnServerConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o VpnServerConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of references to P2SVpnGateways.
func (o VpnServerConfigurationOutput) P2SVpnGateways() P2SVpnGatewayResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) P2SVpnGatewayResponseArrayOutput { return v.P2SVpnGateways }).(P2SVpnGatewayResponseArrayOutput)
}

// The provisioning state of the VpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o VpnServerConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Radius client root certificate of VpnServerConfiguration.
func (o VpnServerConfigurationOutput) RadiusClientRootCertificates() VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
		return v.RadiusClientRootCertificates
	}).(VpnServerConfigRadiusClientRootCertificateResponseArrayOutput)
}

// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
func (o VpnServerConfigurationOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringPtrOutput { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

// Radius Server root certificate of VpnServerConfiguration.
func (o VpnServerConfigurationOutput) RadiusServerRootCertificates() VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
		return v.RadiusServerRootCertificates
	}).(VpnServerConfigRadiusServerRootCertificateResponseArrayOutput)
}

// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
func (o VpnServerConfigurationOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringPtrOutput { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

// Multiple Radius Server configuration for VpnServerConfiguration.
func (o VpnServerConfigurationOutput) RadiusServers() RadiusServerResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) RadiusServerResponseArrayOutput { return v.RadiusServers }).(RadiusServerResponseArrayOutput)
}

// Resource tags.
func (o VpnServerConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VpnServerConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// VPN authentication types for the VpnServerConfiguration.
func (o VpnServerConfigurationOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringArrayOutput { return v.VpnAuthenticationTypes }).(pulumi.StringArrayOutput)
}

// VpnClientIpsecPolicies for VpnServerConfiguration.
func (o VpnServerConfigurationOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) IpsecPolicyResponseArrayOutput { return v.VpnClientIpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

// VPN client revoked certificate of VpnServerConfiguration.
func (o VpnServerConfigurationOutput) VpnClientRevokedCertificates() VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
		return v.VpnClientRevokedCertificates
	}).(VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput)
}

// VPN client root certificate of VpnServerConfiguration.
func (o VpnServerConfigurationOutput) VpnClientRootCertificates() VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
		return v.VpnClientRootCertificates
	}).(VpnServerConfigVpnClientRootCertificateResponseArrayOutput)
}

// VPN protocols for the VpnServerConfiguration.
func (o VpnServerConfigurationOutput) VpnProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringArrayOutput { return v.VpnProtocols }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnServerConfigurationOutput{})
}

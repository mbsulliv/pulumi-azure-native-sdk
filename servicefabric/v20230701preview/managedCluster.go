// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The managed cluster resource
type ManagedCluster struct {
	pulumi.CustomResourceState

	// List of add-on features to enable on the cluster.
	AddonFeatures pulumi.StringArrayOutput `pulumi:"addonFeatures"`
	// VM admin user password.
	AdminPassword pulumi.StringPtrOutput `pulumi:"adminPassword"`
	// VM admin user name.
	AdminUserName pulumi.StringOutput `pulumi:"adminUserName"`
	// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to Internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
	AllowRdpAccess pulumi.BoolPtrOutput `pulumi:"allowRdpAccess"`
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy ApplicationTypeVersionsCleanupPolicyResponsePtrOutput `pulumi:"applicationTypeVersionsCleanupPolicy"`
	// Auxiliary subnets for the cluster.
	AuxiliarySubnets SubnetResponseArrayOutput `pulumi:"auxiliarySubnets"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory AzureActiveDirectoryResponsePtrOutput `pulumi:"azureActiveDirectory"`
	// The port used for client connections to the cluster.
	ClientConnectionPort pulumi.IntPtrOutput `pulumi:"clientConnectionPort"`
	// Client certificates that are allowed to manage the cluster.
	Clients ClientCertificateResponseArrayOutput `pulumi:"clients"`
	// List of thumbprints of the cluster certificates.
	ClusterCertificateThumbprints pulumi.StringArrayOutput `pulumi:"clusterCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property is required when **clusterUpgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion pulumi.StringPtrOutput `pulumi:"clusterCodeVersion"`
	// A service generated unique identifier for the cluster resource.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The current state of the cluster.
	ClusterState pulumi.StringOutput `pulumi:"clusterState"`
	// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
	ClusterUpgradeCadence pulumi.StringPtrOutput `pulumi:"clusterUpgradeCadence"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	ClusterUpgradeMode pulumi.StringPtrOutput `pulumi:"clusterUpgradeMode"`
	// Specify the resource id of a DDoS network protection plan that will be associated with the virtual network of the cluster.
	DdosProtectionPlanId pulumi.StringPtrOutput `pulumi:"ddosProtectionPlanId"`
	// The cluster dns name.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
	EnableAutoOSUpgrade pulumi.BoolPtrOutput `pulumi:"enableAutoOSUpgrade"`
	// Setting this to true creates IPv6 address space for the default VNet used by the cluster. This setting cannot be changed once the cluster is created. The default value for this setting is false.
	EnableIpv6 pulumi.BoolPtrOutput `pulumi:"enableIpv6"`
	// Setting this to true will link the IPv4 address as the ServicePublicIP of the IPv6 address. It can only be set to True if IPv6 is enabled on the cluster.
	EnableServicePublicIP pulumi.BoolPtrOutput `pulumi:"enableServicePublicIP"`
	// Azure resource etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings SettingsSectionDescriptionResponseArrayOutput `pulumi:"fabricSettings"`
	// The fully qualified domain name associated with the public load balancer of the cluster.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The port used for HTTP connections to the cluster.
	HttpGatewayConnectionPort pulumi.IntPtrOutput `pulumi:"httpGatewayConnectionPort"`
	// The list of IP tags associated with the default public IP address of the cluster.
	IpTags IPTagResponseArrayOutput `pulumi:"ipTags"`
	// The IPv4 address associated with the public load balancer of the cluster.
	Ipv4Address pulumi.StringOutput `pulumi:"ipv4Address"`
	// IPv6 address for the cluster if IPv6 is enabled.
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// Load balancing rules that are applied to the public load balancer of the cluster.
	LoadBalancingRules LoadBalancingRuleResponseArrayOutput `pulumi:"loadBalancingRules"`
	// Azure resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Custom Network Security Rules that are applied to the Virtual Network of the cluster.
	NetworkSecurityRules NetworkSecurityRuleResponseArrayOutput `pulumi:"networkSecurityRules"`
	// The provisioning state of the managed cluster resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specify the resource id of a public IP prefix that the load balancer will allocate a public IP address from. Only supports IPv4.
	PublicIPPrefixId pulumi.StringPtrOutput `pulumi:"publicIPPrefixId"`
	// Service endpoints for subnets in the cluster.
	ServiceEndpoints ServiceEndpointResponseArrayOutput `pulumi:"serviceEndpoints"`
	// The sku of the managed cluster
	Sku SkuResponseOutput `pulumi:"sku"`
	// If specified, the node types for the cluster are created in this subnet instead of the default VNet. The **networkSecurityRules** specified for the cluster are also applied to this subnet. This setting cannot be changed once the cluster is created.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// For new clusters, this parameter indicates that it uses Bring your own VNet, but the subnet is specified at node type level; and for such clusters, the subnetId property is required for node types.
	UseCustomVnet pulumi.BoolPtrOutput `pulumi:"useCustomVnet"`
	// Indicates if the cluster has zone resiliency.
	ZonalResiliency pulumi.BoolPtrOutput `pulumi:"zonalResiliency"`
	// Indicates the update mode for Cross Az clusters.
	ZonalUpdateMode pulumi.StringPtrOutput `pulumi:"zonalUpdateMode"`
}

// NewManagedCluster registers a new resource with the given unique name, arguments, and options.
func NewManagedCluster(ctx *pulumi.Context,
	name string, args *ManagedClusterArgs, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminUserName == nil {
		return nil, errors.New("invalid value for required argument 'AdminUserName'")
	}
	if args.DnsName == nil {
		return nil, errors.New("invalid value for required argument 'DnsName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.ClientConnectionPort == nil {
		args.ClientConnectionPort = pulumi.IntPtr(19000)
	}
	if args.HttpGatewayConnectionPort == nil {
		args.HttpGatewayConnectionPort = pulumi.IntPtr(19080)
	}
	if args.ZonalResiliency == nil {
		args.ZonalResiliency = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210701preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220801preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20221001preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230201preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230301preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230901preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20231101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20231201preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20240201preview:ManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedCluster
	err := ctx.RegisterResource("azure-native:servicefabric/v20230701preview:ManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedCluster gets an existing ManagedCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterState, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	var resource ManagedCluster
	err := ctx.ReadResource("azure-native:servicefabric/v20230701preview:ManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedCluster resources.
type managedClusterState struct {
}

type ManagedClusterState struct {
}

func (ManagedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterState)(nil)).Elem()
}

type managedClusterArgs struct {
	// List of add-on features to enable on the cluster.
	AddonFeatures []string `pulumi:"addonFeatures"`
	// VM admin user password.
	AdminPassword *string `pulumi:"adminPassword"`
	// VM admin user name.
	AdminUserName string `pulumi:"adminUserName"`
	// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to Internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
	AllowRdpAccess *bool `pulumi:"allowRdpAccess"`
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicy `pulumi:"applicationTypeVersionsCleanupPolicy"`
	// Auxiliary subnets for the cluster.
	AuxiliarySubnets []Subnet `pulumi:"auxiliarySubnets"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory *AzureActiveDirectory `pulumi:"azureActiveDirectory"`
	// The port used for client connections to the cluster.
	ClientConnectionPort *int `pulumi:"clientConnectionPort"`
	// Client certificates that are allowed to manage the cluster.
	Clients []ClientCertificate `pulumi:"clients"`
	// The Service Fabric runtime version of the cluster. This property is required when **clusterUpgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// The name of the cluster resource.
	ClusterName *string `pulumi:"clusterName"`
	// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
	ClusterUpgradeCadence *string `pulumi:"clusterUpgradeCadence"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	ClusterUpgradeMode *string `pulumi:"clusterUpgradeMode"`
	// Specify the resource id of a DDoS network protection plan that will be associated with the virtual network of the cluster.
	DdosProtectionPlanId *string `pulumi:"ddosProtectionPlanId"`
	// The cluster dns name.
	DnsName string `pulumi:"dnsName"`
	// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
	EnableAutoOSUpgrade *bool `pulumi:"enableAutoOSUpgrade"`
	// Setting this to true creates IPv6 address space for the default VNet used by the cluster. This setting cannot be changed once the cluster is created. The default value for this setting is false.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// Setting this to true will link the IPv4 address as the ServicePublicIP of the IPv6 address. It can only be set to True if IPv6 is enabled on the cluster.
	EnableServicePublicIP *bool `pulumi:"enableServicePublicIP"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings []SettingsSectionDescription `pulumi:"fabricSettings"`
	// The port used for HTTP connections to the cluster.
	HttpGatewayConnectionPort *int `pulumi:"httpGatewayConnectionPort"`
	// The list of IP tags associated with the default public IP address of the cluster.
	IpTags []IPTag `pulumi:"ipTags"`
	// Load balancing rules that are applied to the public load balancer of the cluster.
	LoadBalancingRules []LoadBalancingRule `pulumi:"loadBalancingRules"`
	// Azure resource location.
	Location *string `pulumi:"location"`
	// Custom Network Security Rules that are applied to the Virtual Network of the cluster.
	NetworkSecurityRules []NetworkSecurityRule `pulumi:"networkSecurityRules"`
	// Specify the resource id of a public IP prefix that the load balancer will allocate a public IP address from. Only supports IPv4.
	PublicIPPrefixId *string `pulumi:"publicIPPrefixId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Service endpoints for subnets in the cluster.
	ServiceEndpoints []ServiceEndpoint `pulumi:"serviceEndpoints"`
	// The sku of the managed cluster
	Sku Sku `pulumi:"sku"`
	// If specified, the node types for the cluster are created in this subnet instead of the default VNet. The **networkSecurityRules** specified for the cluster are also applied to this subnet. This setting cannot be changed once the cluster is created.
	SubnetId *string `pulumi:"subnetId"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// For new clusters, this parameter indicates that it uses Bring your own VNet, but the subnet is specified at node type level; and for such clusters, the subnetId property is required for node types.
	UseCustomVnet *bool `pulumi:"useCustomVnet"`
	// Indicates if the cluster has zone resiliency.
	ZonalResiliency *bool `pulumi:"zonalResiliency"`
	// Indicates the update mode for Cross Az clusters.
	ZonalUpdateMode *string `pulumi:"zonalUpdateMode"`
}

// The set of arguments for constructing a ManagedCluster resource.
type ManagedClusterArgs struct {
	// List of add-on features to enable on the cluster.
	AddonFeatures pulumi.StringArrayInput
	// VM admin user password.
	AdminPassword pulumi.StringPtrInput
	// VM admin user name.
	AdminUserName pulumi.StringInput
	// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to Internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
	AllowRdpAccess pulumi.BoolPtrInput
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy ApplicationTypeVersionsCleanupPolicyPtrInput
	// Auxiliary subnets for the cluster.
	AuxiliarySubnets SubnetArrayInput
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory AzureActiveDirectoryPtrInput
	// The port used for client connections to the cluster.
	ClientConnectionPort pulumi.IntPtrInput
	// Client certificates that are allowed to manage the cluster.
	Clients ClientCertificateArrayInput
	// The Service Fabric runtime version of the cluster. This property is required when **clusterUpgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion pulumi.StringPtrInput
	// The name of the cluster resource.
	ClusterName pulumi.StringPtrInput
	// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
	ClusterUpgradeCadence pulumi.StringPtrInput
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	ClusterUpgradeMode pulumi.StringPtrInput
	// Specify the resource id of a DDoS network protection plan that will be associated with the virtual network of the cluster.
	DdosProtectionPlanId pulumi.StringPtrInput
	// The cluster dns name.
	DnsName pulumi.StringInput
	// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
	EnableAutoOSUpgrade pulumi.BoolPtrInput
	// Setting this to true creates IPv6 address space for the default VNet used by the cluster. This setting cannot be changed once the cluster is created. The default value for this setting is false.
	EnableIpv6 pulumi.BoolPtrInput
	// Setting this to true will link the IPv4 address as the ServicePublicIP of the IPv6 address. It can only be set to True if IPv6 is enabled on the cluster.
	EnableServicePublicIP pulumi.BoolPtrInput
	// The list of custom fabric settings to configure the cluster.
	FabricSettings SettingsSectionDescriptionArrayInput
	// The port used for HTTP connections to the cluster.
	HttpGatewayConnectionPort pulumi.IntPtrInput
	// The list of IP tags associated with the default public IP address of the cluster.
	IpTags IPTagArrayInput
	// Load balancing rules that are applied to the public load balancer of the cluster.
	LoadBalancingRules LoadBalancingRuleArrayInput
	// Azure resource location.
	Location pulumi.StringPtrInput
	// Custom Network Security Rules that are applied to the Virtual Network of the cluster.
	NetworkSecurityRules NetworkSecurityRuleArrayInput
	// Specify the resource id of a public IP prefix that the load balancer will allocate a public IP address from. Only supports IPv4.
	PublicIPPrefixId pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Service endpoints for subnets in the cluster.
	ServiceEndpoints ServiceEndpointArrayInput
	// The sku of the managed cluster
	Sku SkuInput
	// If specified, the node types for the cluster are created in this subnet instead of the default VNet. The **networkSecurityRules** specified for the cluster are also applied to this subnet. This setting cannot be changed once the cluster is created.
	SubnetId pulumi.StringPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// For new clusters, this parameter indicates that it uses Bring your own VNet, but the subnet is specified at node type level; and for such clusters, the subnetId property is required for node types.
	UseCustomVnet pulumi.BoolPtrInput
	// Indicates if the cluster has zone resiliency.
	ZonalResiliency pulumi.BoolPtrInput
	// Indicates the update mode for Cross Az clusters.
	ZonalUpdateMode pulumi.StringPtrInput
}

func (ManagedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterArgs)(nil)).Elem()
}

type ManagedClusterInput interface {
	pulumi.Input

	ToManagedClusterOutput() ManagedClusterOutput
	ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput
}

func (*ManagedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCluster)(nil)).Elem()
}

func (i *ManagedCluster) ToManagedClusterOutput() ManagedClusterOutput {
	return i.ToManagedClusterOutputWithContext(context.Background())
}

func (i *ManagedCluster) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterOutput)
}

type ManagedClusterOutput struct{ *pulumi.OutputState }

func (ManagedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCluster)(nil)).Elem()
}

func (o ManagedClusterOutput) ToManagedClusterOutput() ManagedClusterOutput {
	return o
}

func (o ManagedClusterOutput) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return o
}

// List of add-on features to enable on the cluster.
func (o ManagedClusterOutput) AddonFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringArrayOutput { return v.AddonFeatures }).(pulumi.StringArrayOutput)
}

// VM admin user password.
func (o ManagedClusterOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

// VM admin user name.
func (o ManagedClusterOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.AdminUserName }).(pulumi.StringOutput)
}

// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to Internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
func (o ManagedClusterOutput) AllowRdpAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.AllowRdpAccess }).(pulumi.BoolPtrOutput)
}

// The policy used to clean up unused versions.
func (o ManagedClusterOutput) ApplicationTypeVersionsCleanupPolicy() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
		return v.ApplicationTypeVersionsCleanupPolicy
	}).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

// Auxiliary subnets for the cluster.
func (o ManagedClusterOutput) AuxiliarySubnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) SubnetResponseArrayOutput { return v.AuxiliarySubnets }).(SubnetResponseArrayOutput)
}

// The AAD authentication settings of the cluster.
func (o ManagedClusterOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) AzureActiveDirectoryResponsePtrOutput { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

// The port used for client connections to the cluster.
func (o ManagedClusterOutput) ClientConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.IntPtrOutput { return v.ClientConnectionPort }).(pulumi.IntPtrOutput)
}

// Client certificates that are allowed to manage the cluster.
func (o ManagedClusterOutput) Clients() ClientCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) ClientCertificateResponseArrayOutput { return v.Clients }).(ClientCertificateResponseArrayOutput)
}

// List of thumbprints of the cluster certificates.
func (o ManagedClusterOutput) ClusterCertificateThumbprints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringArrayOutput { return v.ClusterCertificateThumbprints }).(pulumi.StringArrayOutput)
}

// The Service Fabric runtime version of the cluster. This property is required when **clusterUpgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
func (o ManagedClusterOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

// A service generated unique identifier for the cluster resource.
func (o ManagedClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The current state of the cluster.
func (o ManagedClusterOutput) ClusterState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ClusterState }).(pulumi.StringOutput)
}

// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
func (o ManagedClusterOutput) ClusterUpgradeCadence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.ClusterUpgradeCadence }).(pulumi.StringPtrOutput)
}

// The upgrade mode of the cluster when new Service Fabric runtime version is available.
func (o ManagedClusterOutput) ClusterUpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.ClusterUpgradeMode }).(pulumi.StringPtrOutput)
}

// Specify the resource id of a DDoS network protection plan that will be associated with the virtual network of the cluster.
func (o ManagedClusterOutput) DdosProtectionPlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.DdosProtectionPlanId }).(pulumi.StringPtrOutput)
}

// The cluster dns name.
func (o ManagedClusterOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
func (o ManagedClusterOutput) EnableAutoOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.EnableAutoOSUpgrade }).(pulumi.BoolPtrOutput)
}

// Setting this to true creates IPv6 address space for the default VNet used by the cluster. This setting cannot be changed once the cluster is created. The default value for this setting is false.
func (o ManagedClusterOutput) EnableIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.EnableIpv6 }).(pulumi.BoolPtrOutput)
}

// Setting this to true will link the IPv4 address as the ServicePublicIP of the IPv6 address. It can only be set to True if IPv6 is enabled on the cluster.
func (o ManagedClusterOutput) EnableServicePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.EnableServicePublicIP }).(pulumi.BoolPtrOutput)
}

// Azure resource etag.
func (o ManagedClusterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The list of custom fabric settings to configure the cluster.
func (o ManagedClusterOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) SettingsSectionDescriptionResponseArrayOutput { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

// The fully qualified domain name associated with the public load balancer of the cluster.
func (o ManagedClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// The port used for HTTP connections to the cluster.
func (o ManagedClusterOutput) HttpGatewayConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.IntPtrOutput { return v.HttpGatewayConnectionPort }).(pulumi.IntPtrOutput)
}

// The list of IP tags associated with the default public IP address of the cluster.
func (o ManagedClusterOutput) IpTags() IPTagResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) IPTagResponseArrayOutput { return v.IpTags }).(IPTagResponseArrayOutput)
}

// The IPv4 address associated with the public load balancer of the cluster.
func (o ManagedClusterOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Ipv4Address }).(pulumi.StringOutput)
}

// IPv6 address for the cluster if IPv6 is enabled.
func (o ManagedClusterOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Ipv6Address }).(pulumi.StringOutput)
}

// Load balancing rules that are applied to the public load balancer of the cluster.
func (o ManagedClusterOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) LoadBalancingRuleResponseArrayOutput { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
}

// Azure resource location.
func (o ManagedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name.
func (o ManagedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Custom Network Security Rules that are applied to the Virtual Network of the cluster.
func (o ManagedClusterOutput) NetworkSecurityRules() NetworkSecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) NetworkSecurityRuleResponseArrayOutput { return v.NetworkSecurityRules }).(NetworkSecurityRuleResponseArrayOutput)
}

// The provisioning state of the managed cluster resource.
func (o ManagedClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specify the resource id of a public IP prefix that the load balancer will allocate a public IP address from. Only supports IPv4.
func (o ManagedClusterOutput) PublicIPPrefixId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.PublicIPPrefixId }).(pulumi.StringPtrOutput)
}

// Service endpoints for subnets in the cluster.
func (o ManagedClusterOutput) ServiceEndpoints() ServiceEndpointResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) ServiceEndpointResponseArrayOutput { return v.ServiceEndpoints }).(ServiceEndpointResponseArrayOutput)
}

// The sku of the managed cluster
func (o ManagedClusterOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *ManagedCluster) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// If specified, the node types for the cluster are created in this subnet instead of the default VNet. The **networkSecurityRules** specified for the cluster are also applied to this subnet. This setting cannot be changed once the cluster is created.
func (o ManagedClusterOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ManagedClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o ManagedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o ManagedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// For new clusters, this parameter indicates that it uses Bring your own VNet, but the subnet is specified at node type level; and for such clusters, the subnetId property is required for node types.
func (o ManagedClusterOutput) UseCustomVnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.UseCustomVnet }).(pulumi.BoolPtrOutput)
}

// Indicates if the cluster has zone resiliency.
func (o ManagedClusterOutput) ZonalResiliency() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.ZonalResiliency }).(pulumi.BoolPtrOutput)
}

// Indicates the update mode for Cross Az clusters.
func (o ManagedClusterOutput) ZonalUpdateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.ZonalUpdateMode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterOutput{})
}

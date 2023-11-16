// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Service Fabric managed cluster resource created or in the process of being created in the specified resource group.
func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:servicefabric/v20221001preview:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedClusterArgs struct {
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The manged cluster resource
type LookupManagedClusterResult struct {
	// List of add-on features to enable on the cluster.
	AddonFeatures []string `pulumi:"addonFeatures"`
	// VM admin user password.
	AdminPassword *string `pulumi:"adminPassword"`
	// VM admin user name.
	AdminUserName string `pulumi:"adminUserName"`
	// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to Internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
	AllowRdpAccess *bool `pulumi:"allowRdpAccess"`
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicyResponse `pulumi:"applicationTypeVersionsCleanupPolicy"`
	// Auxiliary subnets for the cluster.
	AuxiliarySubnets []SubnetResponse `pulumi:"auxiliarySubnets"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory *AzureActiveDirectoryResponse `pulumi:"azureActiveDirectory"`
	// The port used for client connections to the cluster.
	ClientConnectionPort *int `pulumi:"clientConnectionPort"`
	// Client certificates that are allowed to manage the cluster.
	Clients []ClientCertificateResponse `pulumi:"clients"`
	// List of thumbprints of the cluster certificates.
	ClusterCertificateThumbprints []string `pulumi:"clusterCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property is required when **clusterUpgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// A service generated unique identifier for the cluster resource.
	ClusterId string `pulumi:"clusterId"`
	// The current state of the cluster.
	ClusterState string `pulumi:"clusterState"`
	// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
	ClusterUpgradeCadence *string `pulumi:"clusterUpgradeCadence"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	ClusterUpgradeMode *string `pulumi:"clusterUpgradeMode"`
	// The cluster dns name.
	DnsName string `pulumi:"dnsName"`
	// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
	EnableAutoOSUpgrade *bool `pulumi:"enableAutoOSUpgrade"`
	// Setting this to true creates IPv6 address space for the default VNet used by the cluster. This setting cannot be changed once the cluster is created. The default value for this setting is false.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// Setting this to true will link the IPv4 address as the ServicePublicIP of the IPv6 address. It can only be set to True if IPv6 is enabled on the cluster.
	EnableServicePublicIP *bool `pulumi:"enableServicePublicIP"`
	// Azure resource etag.
	Etag string `pulumi:"etag"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings []SettingsSectionDescriptionResponse `pulumi:"fabricSettings"`
	// The fully qualified domain name associated with the public load balancer of the cluster.
	Fqdn string `pulumi:"fqdn"`
	// The port used for HTTP connections to the cluster.
	HttpGatewayConnectionPort *int `pulumi:"httpGatewayConnectionPort"`
	// Azure resource identifier.
	Id string `pulumi:"id"`
	// The list of IP tags associated with the default public IP address of the cluster.
	IpTags []IPTagResponse `pulumi:"ipTags"`
	// The IPv4 address associated with the public load balancer of the cluster.
	Ipv4Address string `pulumi:"ipv4Address"`
	// IPv6 address for the cluster if IPv6 is enabled.
	Ipv6Address string `pulumi:"ipv6Address"`
	// Load balancing rules that are applied to the public load balancer of the cluster.
	LoadBalancingRules []LoadBalancingRuleResponse `pulumi:"loadBalancingRules"`
	// Azure resource location.
	Location string `pulumi:"location"`
	// Azure resource name.
	Name string `pulumi:"name"`
	// Custom Network Security Rules that are applied to the Virtual Network of the cluster.
	NetworkSecurityRules []NetworkSecurityRuleResponse `pulumi:"networkSecurityRules"`
	// The provisioning state of the managed cluster resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Service endpoints for subnets in the cluster.
	ServiceEndpoints []ServiceEndpointResponse `pulumi:"serviceEndpoints"`
	// The sku of the managed cluster
	Sku *SkuResponse `pulumi:"sku"`
	// If specified, the node types for the cluster are created in this subnet instead of the default VNet. The **networkSecurityRules** specified for the cluster are also applied to this subnet. This setting cannot be changed once the cluster is created.
	SubnetId *string `pulumi:"subnetId"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type string `pulumi:"type"`
	// For new clusters, this parameter indicates that it uses Bring your own VNet, but the subnet is specified at node type level; and for such clusters, the subnetId property is required for node types.
	UseCustomVnet *bool `pulumi:"useCustomVnet"`
	// Indicates if the cluster has zone resiliency.
	ZonalResiliency *bool `pulumi:"zonalResiliency"`
	// Indicates the update mode for Cross Az clusters.
	ZonalUpdateMode *string `pulumi:"zonalUpdateMode"`
}

// Defaults sets the appropriate defaults for LookupManagedClusterResult
func (val *LookupManagedClusterResult) Defaults() *LookupManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.ClientConnectionPort == nil {
		clientConnectionPort_ := 19000
		tmp.ClientConnectionPort = &clientConnectionPort_
	}
	if tmp.HttpGatewayConnectionPort == nil {
		httpGatewayConnectionPort_ := 19080
		tmp.HttpGatewayConnectionPort = &httpGatewayConnectionPort_
	}
	if tmp.ZonalResiliency == nil {
		zonalResiliency_ := false
		tmp.ZonalResiliency = &zonalResiliency_
	}
	return &tmp
}

func LookupManagedClusterOutput(ctx *pulumi.Context, args LookupManagedClusterOutputArgs, opts ...pulumi.InvokeOption) LookupManagedClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedClusterResult, error) {
			args := v.(LookupManagedClusterArgs)
			r, err := LookupManagedCluster(ctx, &args, opts...)
			var s LookupManagedClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedClusterResultOutput)
}

type LookupManagedClusterOutputArgs struct {
	// The name of the cluster resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterArgs)(nil)).Elem()
}

// The manged cluster resource
type LookupManagedClusterResultOutput struct{ *pulumi.OutputState }

func (LookupManagedClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterResult)(nil)).Elem()
}

func (o LookupManagedClusterResultOutput) ToLookupManagedClusterResultOutput() LookupManagedClusterResultOutput {
	return o
}

func (o LookupManagedClusterResultOutput) ToLookupManagedClusterResultOutputWithContext(ctx context.Context) LookupManagedClusterResultOutput {
	return o
}

// List of add-on features to enable on the cluster.
func (o LookupManagedClusterResultOutput) AddonFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []string { return v.AddonFeatures }).(pulumi.StringArrayOutput)
}

// VM admin user password.
func (o LookupManagedClusterResultOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

// VM admin user name.
func (o LookupManagedClusterResultOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.AdminUserName }).(pulumi.StringOutput)
}

// Setting this to true enables RDP access to the VM. The default NSG rule opens RDP port to Internet which can be overridden with custom Network Security Rules. The default value for this setting is false.
func (o LookupManagedClusterResultOutput) AllowRdpAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.AllowRdpAccess }).(pulumi.BoolPtrOutput)
}

// The policy used to clean up unused versions.
func (o LookupManagedClusterResultOutput) ApplicationTypeVersionsCleanupPolicy() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ApplicationTypeVersionsCleanupPolicyResponse {
		return v.ApplicationTypeVersionsCleanupPolicy
	}).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

// Auxiliary subnets for the cluster.
func (o LookupManagedClusterResultOutput) AuxiliarySubnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []SubnetResponse { return v.AuxiliarySubnets }).(SubnetResponseArrayOutput)
}

// The AAD authentication settings of the cluster.
func (o LookupManagedClusterResultOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *AzureActiveDirectoryResponse { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

// The port used for client connections to the cluster.
func (o LookupManagedClusterResultOutput) ClientConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *int { return v.ClientConnectionPort }).(pulumi.IntPtrOutput)
}

// Client certificates that are allowed to manage the cluster.
func (o LookupManagedClusterResultOutput) Clients() ClientCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []ClientCertificateResponse { return v.Clients }).(ClientCertificateResponseArrayOutput)
}

// List of thumbprints of the cluster certificates.
func (o LookupManagedClusterResultOutput) ClusterCertificateThumbprints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []string { return v.ClusterCertificateThumbprints }).(pulumi.StringArrayOutput)
}

// The Service Fabric runtime version of the cluster. This property is required when **clusterUpgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
func (o LookupManagedClusterResultOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

// A service generated unique identifier for the cluster resource.
func (o LookupManagedClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The current state of the cluster.
func (o LookupManagedClusterResultOutput) ClusterState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ClusterState }).(pulumi.StringOutput)
}

// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **clusterUpgradeMode** is set to 'Automatic'.
func (o LookupManagedClusterResultOutput) ClusterUpgradeCadence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.ClusterUpgradeCadence }).(pulumi.StringPtrOutput)
}

// The upgrade mode of the cluster when new Service Fabric runtime version is available.
func (o LookupManagedClusterResultOutput) ClusterUpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.ClusterUpgradeMode }).(pulumi.StringPtrOutput)
}

// The cluster dns name.
func (o LookupManagedClusterResultOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.DnsName }).(pulumi.StringOutput)
}

// Setting this to true enables automatic OS upgrade for the node types that are created using any platform OS image with version 'latest'. The default value for this setting is false.
func (o LookupManagedClusterResultOutput) EnableAutoOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnableAutoOSUpgrade }).(pulumi.BoolPtrOutput)
}

// Setting this to true creates IPv6 address space for the default VNet used by the cluster. This setting cannot be changed once the cluster is created. The default value for this setting is false.
func (o LookupManagedClusterResultOutput) EnableIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnableIpv6 }).(pulumi.BoolPtrOutput)
}

// Setting this to true will link the IPv4 address as the ServicePublicIP of the IPv6 address. It can only be set to True if IPv6 is enabled on the cluster.
func (o LookupManagedClusterResultOutput) EnableServicePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnableServicePublicIP }).(pulumi.BoolPtrOutput)
}

// Azure resource etag.
func (o LookupManagedClusterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The list of custom fabric settings to configure the cluster.
func (o LookupManagedClusterResultOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []SettingsSectionDescriptionResponse { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

// The fully qualified domain name associated with the public load balancer of the cluster.
func (o LookupManagedClusterResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// The port used for HTTP connections to the cluster.
func (o LookupManagedClusterResultOutput) HttpGatewayConnectionPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *int { return v.HttpGatewayConnectionPort }).(pulumi.IntPtrOutput)
}

// Azure resource identifier.
func (o LookupManagedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of IP tags associated with the default public IP address of the cluster.
func (o LookupManagedClusterResultOutput) IpTags() IPTagResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []IPTagResponse { return v.IpTags }).(IPTagResponseArrayOutput)
}

// The IPv4 address associated with the public load balancer of the cluster.
func (o LookupManagedClusterResultOutput) Ipv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Ipv4Address }).(pulumi.StringOutput)
}

// IPv6 address for the cluster if IPv6 is enabled.
func (o LookupManagedClusterResultOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Ipv6Address }).(pulumi.StringOutput)
}

// Load balancing rules that are applied to the public load balancer of the cluster.
func (o LookupManagedClusterResultOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []LoadBalancingRuleResponse { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
}

// Azure resource location.
func (o LookupManagedClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name.
func (o LookupManagedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Custom Network Security Rules that are applied to the Virtual Network of the cluster.
func (o LookupManagedClusterResultOutput) NetworkSecurityRules() NetworkSecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []NetworkSecurityRuleResponse { return v.NetworkSecurityRules }).(NetworkSecurityRuleResponseArrayOutput)
}

// The provisioning state of the managed cluster resource.
func (o LookupManagedClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Service endpoints for subnets in the cluster.
func (o LookupManagedClusterResultOutput) ServiceEndpoints() ServiceEndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []ServiceEndpointResponse { return v.ServiceEndpoints }).(ServiceEndpointResponseArrayOutput)
}

// The sku of the managed cluster
func (o LookupManagedClusterResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// If specified, the node types for the cluster are created in this subnet instead of the default VNet. The **networkSecurityRules** specified for the cluster are also applied to this subnet. This setting cannot be changed once the cluster is created.
func (o LookupManagedClusterResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupManagedClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o LookupManagedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o LookupManagedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// For new clusters, this parameter indicates that it uses Bring your own VNet, but the subnet is specified at node type level; and for such clusters, the subnetId property is required for node types.
func (o LookupManagedClusterResultOutput) UseCustomVnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.UseCustomVnet }).(pulumi.BoolPtrOutput)
}

// Indicates if the cluster has zone resiliency.
func (o LookupManagedClusterResultOutput) ZonalResiliency() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.ZonalResiliency }).(pulumi.BoolPtrOutput)
}

// Indicates the update mode for Cross Az clusters.
func (o LookupManagedClusterResultOutput) ZonalUpdateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.ZonalUpdateMode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedClusterResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Managed cluster.
type ManagedCluster struct {
	pulumi.CustomResourceState

	// The Azure Active Directory configuration.
	AadProfile ManagedClusterAADProfileResponsePtrOutput `pulumi:"aadProfile"`
	// The profile of managed cluster add-on.
	AddonProfiles ManagedClusterAddonProfileResponseMapOutput `pulumi:"addonProfiles"`
	// The agent pool properties.
	AgentPoolProfiles ManagedClusterAgentPoolProfileResponseArrayOutput `pulumi:"agentPoolProfiles"`
	// The access profile for managed cluster API server.
	ApiServerAccessProfile ManagedClusterAPIServerAccessProfileResponsePtrOutput `pulumi:"apiServerAccessProfile"`
	// Parameters to be applied to the cluster-autoscaler when enabled
	AutoScalerProfile ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput `pulumi:"autoScalerProfile"`
	// The auto upgrade configuration.
	AutoUpgradeProfile ManagedClusterAutoUpgradeProfileResponsePtrOutput `pulumi:"autoUpgradeProfile"`
	// The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.
	AzurePortalFQDN pulumi.StringOutput `pulumi:"azurePortalFQDN"`
	// If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).
	DisableLocalAccounts pulumi.BoolPtrOutput `pulumi:"disableLocalAccounts"`
	// This is of the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'
	DiskEncryptionSetID pulumi.StringPtrOutput `pulumi:"diskEncryptionSetID"`
	// This cannot be updated once the Managed Cluster has been created.
	DnsPrefix pulumi.StringPtrOutput `pulumi:"dnsPrefix"`
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy pulumi.BoolPtrOutput `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC pulumi.BoolPtrOutput `pulumi:"enableRBAC"`
	// The extended location of the Virtual Machine.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The FQDN of the master pool.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// This cannot be updated once the Managed Cluster has been created.
	FqdnSubdomain pulumi.StringPtrOutput `pulumi:"fqdnSubdomain"`
	// Configurations for provisioning the cluster with HTTP proxy servers.
	HttpProxyConfig ManagedClusterHTTPProxyConfigResponsePtrOutput `pulumi:"httpProxyConfig"`
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityResponsePtrOutput `pulumi:"identity"`
	// Identities associated with the cluster.
	IdentityProfile UserAssignedIdentityResponseMapOutput `pulumi:"identityProfile"`
	// When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details.
	KubernetesVersion pulumi.StringPtrOutput `pulumi:"kubernetesVersion"`
	// The profile for Linux VMs in the Managed Cluster.
	LinuxProfile ContainerServiceLinuxProfileResponsePtrOutput `pulumi:"linuxProfile"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The max number of agent pools for the managed cluster.
	MaxAgentPools pulumi.IntOutput `pulumi:"maxAgentPools"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The network configuration profile.
	NetworkProfile ContainerServiceNetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// The name of the resource group containing agent pool nodes.
	NodeResourceGroup pulumi.StringPtrOutput `pulumi:"nodeResourceGroup"`
	// See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.
	PodIdentityProfile ManagedClusterPodIdentityProfileResponsePtrOutput `pulumi:"podIdentityProfile"`
	// The Power State of the cluster.
	PowerState PowerStateResponseOutput `pulumi:"powerState"`
	// The FQDN of private cluster.
	PrivateFQDN pulumi.StringOutput `pulumi:"privateFQDN"`
	// Private link resources associated with the cluster.
	PrivateLinkResources PrivateLinkResourceResponseArrayOutput `pulumi:"privateLinkResources"`
	// The current provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Allow or deny public network access for AKS
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Security profile for the managed cluster.
	SecurityProfile ManagedClusterSecurityProfileResponsePtrOutput `pulumi:"securityProfile"`
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile ManagedClusterServicePrincipalProfileResponsePtrOutput `pulumi:"servicePrincipalProfile"`
	// The managed cluster SKU.
	Sku ManagedClusterSKUResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The profile for Windows VMs in the Managed Cluster.
	WindowsProfile ManagedClusterWindowsProfileResponsePtrOutput `pulumi:"windowsProfile"`
}

// NewManagedCluster registers a new resource with the given unique name, arguments, and options.
func NewManagedCluster(ctx *pulumi.Context,
	name string, args *ManagedClusterArgs, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.NetworkProfile != nil {
		args.NetworkProfile = args.NetworkProfile.ToContainerServiceNetworkProfilePtrOutput().ApplyT(func(v *ContainerServiceNetworkProfile) *ContainerServiceNetworkProfile { return v.Defaults() }).(ContainerServiceNetworkProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20170831:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20180331:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20180801preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221002preview:ManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedCluster
	err := ctx.RegisterResource("azure-native:containerservice/v20211001:ManagedCluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:containerservice/v20211001:ManagedCluster", name, id, state, &resource, opts...)
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
	// The Azure Active Directory configuration.
	AadProfile *ManagedClusterAADProfile `pulumi:"aadProfile"`
	// The profile of managed cluster add-on.
	AddonProfiles map[string]ManagedClusterAddonProfile `pulumi:"addonProfiles"`
	// The agent pool properties.
	AgentPoolProfiles []ManagedClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	// The access profile for managed cluster API server.
	ApiServerAccessProfile *ManagedClusterAPIServerAccessProfile `pulumi:"apiServerAccessProfile"`
	// Parameters to be applied to the cluster-autoscaler when enabled
	AutoScalerProfile *ManagedClusterPropertiesAutoScalerProfile `pulumi:"autoScalerProfile"`
	// The auto upgrade configuration.
	AutoUpgradeProfile *ManagedClusterAutoUpgradeProfile `pulumi:"autoUpgradeProfile"`
	// If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).
	DisableLocalAccounts *bool `pulumi:"disableLocalAccounts"`
	// This is of the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'
	DiskEncryptionSetID *string `pulumi:"diskEncryptionSetID"`
	// This cannot be updated once the Managed Cluster has been created.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC *bool `pulumi:"enableRBAC"`
	// The extended location of the Virtual Machine.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// This cannot be updated once the Managed Cluster has been created.
	FqdnSubdomain *string `pulumi:"fqdnSubdomain"`
	// Configurations for provisioning the cluster with HTTP proxy servers.
	HttpProxyConfig *ManagedClusterHTTPProxyConfig `pulumi:"httpProxyConfig"`
	// The identity of the managed cluster, if configured.
	Identity *ManagedClusterIdentity `pulumi:"identity"`
	// Identities associated with the cluster.
	IdentityProfile map[string]UserAssignedIdentity `pulumi:"identityProfile"`
	// When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// The profile for Linux VMs in the Managed Cluster.
	LinuxProfile *ContainerServiceLinuxProfile `pulumi:"linuxProfile"`
	// Resource location
	Location *string `pulumi:"location"`
	// The network configuration profile.
	NetworkProfile *ContainerServiceNetworkProfile `pulumi:"networkProfile"`
	// The name of the resource group containing agent pool nodes.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.
	PodIdentityProfile *ManagedClusterPodIdentityProfile `pulumi:"podIdentityProfile"`
	// Private link resources associated with the cluster.
	PrivateLinkResources []PrivateLinkResource `pulumi:"privateLinkResources"`
	// Allow or deny public network access for AKS
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName *string `pulumi:"resourceName"`
	// Security profile for the managed cluster.
	SecurityProfile *ManagedClusterSecurityProfile `pulumi:"securityProfile"`
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfile `pulumi:"servicePrincipalProfile"`
	// The managed cluster SKU.
	Sku *ManagedClusterSKU `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The profile for Windows VMs in the Managed Cluster.
	WindowsProfile *ManagedClusterWindowsProfile `pulumi:"windowsProfile"`
}

// The set of arguments for constructing a ManagedCluster resource.
type ManagedClusterArgs struct {
	// The Azure Active Directory configuration.
	AadProfile ManagedClusterAADProfilePtrInput
	// The profile of managed cluster add-on.
	AddonProfiles ManagedClusterAddonProfileMapInput
	// The agent pool properties.
	AgentPoolProfiles ManagedClusterAgentPoolProfileArrayInput
	// The access profile for managed cluster API server.
	ApiServerAccessProfile ManagedClusterAPIServerAccessProfilePtrInput
	// Parameters to be applied to the cluster-autoscaler when enabled
	AutoScalerProfile ManagedClusterPropertiesAutoScalerProfilePtrInput
	// The auto upgrade configuration.
	AutoUpgradeProfile ManagedClusterAutoUpgradeProfilePtrInput
	// If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).
	DisableLocalAccounts pulumi.BoolPtrInput
	// This is of the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'
	DiskEncryptionSetID pulumi.StringPtrInput
	// This cannot be updated once the Managed Cluster has been created.
	DnsPrefix pulumi.StringPtrInput
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC pulumi.BoolPtrInput
	// The extended location of the Virtual Machine.
	ExtendedLocation ExtendedLocationPtrInput
	// This cannot be updated once the Managed Cluster has been created.
	FqdnSubdomain pulumi.StringPtrInput
	// Configurations for provisioning the cluster with HTTP proxy servers.
	HttpProxyConfig ManagedClusterHTTPProxyConfigPtrInput
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityPtrInput
	// Identities associated with the cluster.
	IdentityProfile UserAssignedIdentityMapInput
	// When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details.
	KubernetesVersion pulumi.StringPtrInput
	// The profile for Linux VMs in the Managed Cluster.
	LinuxProfile ContainerServiceLinuxProfilePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The network configuration profile.
	NetworkProfile ContainerServiceNetworkProfilePtrInput
	// The name of the resource group containing agent pool nodes.
	NodeResourceGroup pulumi.StringPtrInput
	// See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.
	PodIdentityProfile ManagedClusterPodIdentityProfilePtrInput
	// Private link resources associated with the cluster.
	PrivateLinkResources PrivateLinkResourceArrayInput
	// Allow or deny public network access for AKS
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringPtrInput
	// Security profile for the managed cluster.
	SecurityProfile ManagedClusterSecurityProfilePtrInput
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile ManagedClusterServicePrincipalProfilePtrInput
	// The managed cluster SKU.
	Sku ManagedClusterSKUPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The profile for Windows VMs in the Managed Cluster.
	WindowsProfile ManagedClusterWindowsProfilePtrInput
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

// The Azure Active Directory configuration.
func (o ManagedClusterOutput) AadProfile() ManagedClusterAADProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAADProfileResponsePtrOutput { return v.AadProfile }).(ManagedClusterAADProfileResponsePtrOutput)
}

// The profile of managed cluster add-on.
func (o ManagedClusterOutput) AddonProfiles() ManagedClusterAddonProfileResponseMapOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAddonProfileResponseMapOutput { return v.AddonProfiles }).(ManagedClusterAddonProfileResponseMapOutput)
}

// The agent pool properties.
func (o ManagedClusterOutput) AgentPoolProfiles() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAgentPoolProfileResponseArrayOutput { return v.AgentPoolProfiles }).(ManagedClusterAgentPoolProfileResponseArrayOutput)
}

// The access profile for managed cluster API server.
func (o ManagedClusterOutput) ApiServerAccessProfile() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
		return v.ApiServerAccessProfile
	}).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

// Parameters to be applied to the cluster-autoscaler when enabled
func (o ManagedClusterOutput) AutoScalerProfile() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
		return v.AutoScalerProfile
	}).(ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput)
}

// The auto upgrade configuration.
func (o ManagedClusterOutput) AutoUpgradeProfile() ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAutoUpgradeProfileResponsePtrOutput { return v.AutoUpgradeProfile }).(ManagedClusterAutoUpgradeProfileResponsePtrOutput)
}

// The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.
func (o ManagedClusterOutput) AzurePortalFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.AzurePortalFQDN }).(pulumi.StringOutput)
}

// If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).
func (o ManagedClusterOutput) DisableLocalAccounts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.DisableLocalAccounts }).(pulumi.BoolPtrOutput)
}

// This is of the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'
func (o ManagedClusterOutput) DiskEncryptionSetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.DiskEncryptionSetID }).(pulumi.StringPtrOutput)
}

// This cannot be updated once the Managed Cluster has been created.
func (o ManagedClusterOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
func (o ManagedClusterOutput) EnablePodSecurityPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.EnablePodSecurityPolicy }).(pulumi.BoolPtrOutput)
}

// Whether to enable Kubernetes Role-Based Access Control.
func (o ManagedClusterOutput) EnableRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.EnableRBAC }).(pulumi.BoolPtrOutput)
}

// The extended location of the Virtual Machine.
func (o ManagedClusterOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The FQDN of the master pool.
func (o ManagedClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// This cannot be updated once the Managed Cluster has been created.
func (o ManagedClusterOutput) FqdnSubdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.FqdnSubdomain }).(pulumi.StringPtrOutput)
}

// Configurations for provisioning the cluster with HTTP proxy servers.
func (o ManagedClusterOutput) HttpProxyConfig() ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterHTTPProxyConfigResponsePtrOutput { return v.HttpProxyConfig }).(ManagedClusterHTTPProxyConfigResponsePtrOutput)
}

// The identity of the managed cluster, if configured.
func (o ManagedClusterOutput) Identity() ManagedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterIdentityResponsePtrOutput { return v.Identity }).(ManagedClusterIdentityResponsePtrOutput)
}

// Identities associated with the cluster.
func (o ManagedClusterOutput) IdentityProfile() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedCluster) UserAssignedIdentityResponseMapOutput { return v.IdentityProfile }).(UserAssignedIdentityResponseMapOutput)
}

// When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details.
func (o ManagedClusterOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

// The profile for Linux VMs in the Managed Cluster.
func (o ManagedClusterOutput) LinuxProfile() ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ContainerServiceLinuxProfileResponsePtrOutput { return v.LinuxProfile }).(ContainerServiceLinuxProfileResponsePtrOutput)
}

// Resource location
func (o ManagedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The max number of agent pools for the managed cluster.
func (o ManagedClusterOutput) MaxAgentPools() pulumi.IntOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.IntOutput { return v.MaxAgentPools }).(pulumi.IntOutput)
}

// Resource name
func (o ManagedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network configuration profile.
func (o ManagedClusterOutput) NetworkProfile() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ContainerServiceNetworkProfileResponsePtrOutput { return v.NetworkProfile }).(ContainerServiceNetworkProfileResponsePtrOutput)
}

// The name of the resource group containing agent pool nodes.
func (o ManagedClusterOutput) NodeResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.NodeResourceGroup }).(pulumi.StringPtrOutput)
}

// See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.
func (o ManagedClusterOutput) PodIdentityProfile() ManagedClusterPodIdentityProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterPodIdentityProfileResponsePtrOutput { return v.PodIdentityProfile }).(ManagedClusterPodIdentityProfileResponsePtrOutput)
}

// The Power State of the cluster.
func (o ManagedClusterOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v *ManagedCluster) PowerStateResponseOutput { return v.PowerState }).(PowerStateResponseOutput)
}

// The FQDN of private cluster.
func (o ManagedClusterOutput) PrivateFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.PrivateFQDN }).(pulumi.StringOutput)
}

// Private link resources associated with the cluster.
func (o ManagedClusterOutput) PrivateLinkResources() PrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) PrivateLinkResourceResponseArrayOutput { return v.PrivateLinkResources }).(PrivateLinkResourceResponseArrayOutput)
}

// The current provisioning state.
func (o ManagedClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Allow or deny public network access for AKS
func (o ManagedClusterOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Security profile for the managed cluster.
func (o ManagedClusterOutput) SecurityProfile() ManagedClusterSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterSecurityProfileResponsePtrOutput { return v.SecurityProfile }).(ManagedClusterSecurityProfileResponsePtrOutput)
}

// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
func (o ManagedClusterOutput) ServicePrincipalProfile() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterServicePrincipalProfileResponsePtrOutput {
		return v.ServicePrincipalProfile
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

// The managed cluster SKU.
func (o ManagedClusterOutput) Sku() ManagedClusterSKUResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterSKUResponsePtrOutput { return v.Sku }).(ManagedClusterSKUResponsePtrOutput)
}

// Resource tags
func (o ManagedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ManagedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The profile for Windows VMs in the Managed Cluster.
func (o ManagedClusterOutput) WindowsProfile() ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterWindowsProfileResponsePtrOutput { return v.WindowsProfile }).(ManagedClusterWindowsProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterOutput{})
}

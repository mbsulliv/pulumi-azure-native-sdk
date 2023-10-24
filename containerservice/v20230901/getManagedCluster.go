// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Managed cluster.
func LookupManagedCluster(ctx *pulumi.Context, args *LookupManagedClusterArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedClusterResult
	err := ctx.Invoke("azure-native:containerservice/v20230901:getManagedCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedClusterArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// Managed cluster.
type LookupManagedClusterResult struct {
	// The Azure Active Directory configuration.
	AadProfile *ManagedClusterAADProfileResponse `pulumi:"aadProfile"`
	// The profile of managed cluster add-on.
	AddonProfiles map[string]ManagedClusterAddonProfileResponse `pulumi:"addonProfiles"`
	// The agent pool properties.
	AgentPoolProfiles []ManagedClusterAgentPoolProfileResponse `pulumi:"agentPoolProfiles"`
	// The access profile for managed cluster API server.
	ApiServerAccessProfile *ManagedClusterAPIServerAccessProfileResponse `pulumi:"apiServerAccessProfile"`
	// Parameters to be applied to the cluster-autoscaler when enabled
	AutoScalerProfile *ManagedClusterPropertiesResponseAutoScalerProfile `pulumi:"autoScalerProfile"`
	// The auto upgrade configuration.
	AutoUpgradeProfile *ManagedClusterAutoUpgradeProfileResponse `pulumi:"autoUpgradeProfile"`
	// Azure Monitor addon profiles for monitoring the managed cluster.
	AzureMonitorProfile *ManagedClusterAzureMonitorProfileResponse `pulumi:"azureMonitorProfile"`
	// The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.
	AzurePortalFQDN string `pulumi:"azurePortalFQDN"`
	// If kubernetesVersion was a fully specified version <major.minor.patch>, this field will be exactly equal to it. If kubernetesVersion was <major.minor>, this field will contain the full <major.minor.patch> version being used.
	CurrentKubernetesVersion string `pulumi:"currentKubernetesVersion"`
	// If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).
	DisableLocalAccounts *bool `pulumi:"disableLocalAccounts"`
	// This is of the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'
	DiskEncryptionSetID *string `pulumi:"diskEncryptionSetID"`
	// This cannot be updated once the Managed Cluster has been created.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at https://aka.ms/k8s/psp and https://aka.ms/aks/psp.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC *bool `pulumi:"enableRBAC"`
	// The extended location of the Virtual Machine.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The FQDN of the master pool.
	Fqdn string `pulumi:"fqdn"`
	// This cannot be updated once the Managed Cluster has been created.
	FqdnSubdomain *string `pulumi:"fqdnSubdomain"`
	// Configurations for provisioning the cluster with HTTP proxy servers.
	HttpProxyConfig *ManagedClusterHTTPProxyConfigResponse `pulumi:"httpProxyConfig"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The identity of the managed cluster, if configured.
	Identity *ManagedClusterIdentityResponse `pulumi:"identity"`
	// Identities associated with the cluster.
	IdentityProfile map[string]UserAssignedIdentityResponse `pulumi:"identityProfile"`
	// Both patch version <major.minor.patch> (e.g. 1.20.13) and <major.minor> (e.g. 1.20) are supported. When <major.minor> is specified, the latest supported GA patch version is chosen automatically. Updating the cluster with the same <major.minor> once it has been created (e.g. 1.14.x -> 1.14) will not trigger an upgrade, even if a newer patch version is available. When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// The profile for Linux VMs in the Managed Cluster.
	LinuxProfile *ContainerServiceLinuxProfileResponse `pulumi:"linuxProfile"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The max number of agent pools for the managed cluster.
	MaxAgentPools int `pulumi:"maxAgentPools"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The network configuration profile.
	NetworkProfile *ContainerServiceNetworkProfileResponse `pulumi:"networkProfile"`
	// The name of the resource group containing agent pool nodes.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// The OIDC issuer profile of the Managed Cluster.
	OidcIssuerProfile *ManagedClusterOIDCIssuerProfileResponse `pulumi:"oidcIssuerProfile"`
	// See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.
	PodIdentityProfile *ManagedClusterPodIdentityProfileResponse `pulumi:"podIdentityProfile"`
	// The Power State of the cluster.
	PowerState PowerStateResponse `pulumi:"powerState"`
	// The FQDN of private cluster.
	PrivateFQDN string `pulumi:"privateFQDN"`
	// Private link resources associated with the cluster.
	PrivateLinkResources []PrivateLinkResourceResponse `pulumi:"privateLinkResources"`
	// The current provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Allow or deny public network access for AKS
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The resourceUID uniquely identifies ManagedClusters that reuse ARM ResourceIds (i.e: create, delete, create sequence)
	ResourceUID string `pulumi:"resourceUID"`
	// Security profile for the managed cluster.
	SecurityProfile *ManagedClusterSecurityProfileResponse `pulumi:"securityProfile"`
	// Service mesh profile for a managed cluster.
	ServiceMeshProfile *ServiceMeshProfileResponse `pulumi:"serviceMeshProfile"`
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	// The managed cluster SKU.
	Sku *ManagedClusterSKUResponse `pulumi:"sku"`
	// Storage profile for the managed cluster.
	StorageProfile *ManagedClusterStorageProfileResponse `pulumi:"storageProfile"`
	// The support plan for the Managed Cluster. If unspecified, the default is 'KubernetesOfficial'.
	SupportPlan *string `pulumi:"supportPlan"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Settings for upgrading a cluster.
	UpgradeSettings *ClusterUpgradeSettingsResponse `pulumi:"upgradeSettings"`
	// The profile for Windows VMs in the Managed Cluster.
	WindowsProfile *ManagedClusterWindowsProfileResponse `pulumi:"windowsProfile"`
	// Workload Auto-scaler profile for the managed cluster.
	WorkloadAutoScalerProfile *ManagedClusterWorkloadAutoScalerProfileResponse `pulumi:"workloadAutoScalerProfile"`
}

// Defaults sets the appropriate defaults for LookupManagedClusterResult
func (val *LookupManagedClusterResult) Defaults() *LookupManagedClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkProfile = tmp.NetworkProfile.Defaults()

	tmp.SecurityProfile = tmp.SecurityProfile.Defaults()

	tmp.WorkloadAutoScalerProfile = tmp.WorkloadAutoScalerProfile.Defaults()

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
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupManagedClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterArgs)(nil)).Elem()
}

// Managed cluster.
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

func (o LookupManagedClusterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagedClusterResult] {
	return pulumix.Output[LookupManagedClusterResult]{
		OutputState: o.OutputState,
	}
}

// The Azure Active Directory configuration.
func (o LookupManagedClusterResultOutput) AadProfile() ManagedClusterAADProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterAADProfileResponse { return v.AadProfile }).(ManagedClusterAADProfileResponsePtrOutput)
}

// The profile of managed cluster add-on.
func (o LookupManagedClusterResultOutput) AddonProfiles() ManagedClusterAddonProfileResponseMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]ManagedClusterAddonProfileResponse {
		return v.AddonProfiles
	}).(ManagedClusterAddonProfileResponseMapOutput)
}

// The agent pool properties.
func (o LookupManagedClusterResultOutput) AgentPoolProfiles() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []ManagedClusterAgentPoolProfileResponse {
		return v.AgentPoolProfiles
	}).(ManagedClusterAgentPoolProfileResponseArrayOutput)
}

// The access profile for managed cluster API server.
func (o LookupManagedClusterResultOutput) ApiServerAccessProfile() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterAPIServerAccessProfileResponse {
		return v.ApiServerAccessProfile
	}).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

// Parameters to be applied to the cluster-autoscaler when enabled
func (o LookupManagedClusterResultOutput) AutoScalerProfile() ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterPropertiesResponseAutoScalerProfile {
		return v.AutoScalerProfile
	}).(ManagedClusterPropertiesResponseAutoScalerProfilePtrOutput)
}

// The auto upgrade configuration.
func (o LookupManagedClusterResultOutput) AutoUpgradeProfile() ManagedClusterAutoUpgradeProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterAutoUpgradeProfileResponse {
		return v.AutoUpgradeProfile
	}).(ManagedClusterAutoUpgradeProfileResponsePtrOutput)
}

// Azure Monitor addon profiles for monitoring the managed cluster.
func (o LookupManagedClusterResultOutput) AzureMonitorProfile() ManagedClusterAzureMonitorProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterAzureMonitorProfileResponse {
		return v.AzureMonitorProfile
	}).(ManagedClusterAzureMonitorProfileResponsePtrOutput)
}

// The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.
func (o LookupManagedClusterResultOutput) AzurePortalFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.AzurePortalFQDN }).(pulumi.StringOutput)
}

// If kubernetesVersion was a fully specified version <major.minor.patch>, this field will be exactly equal to it. If kubernetesVersion was <major.minor>, this field will contain the full <major.minor.patch> version being used.
func (o LookupManagedClusterResultOutput) CurrentKubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.CurrentKubernetesVersion }).(pulumi.StringOutput)
}

// If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).
func (o LookupManagedClusterResultOutput) DisableLocalAccounts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.DisableLocalAccounts }).(pulumi.BoolPtrOutput)
}

// This is of the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'
func (o LookupManagedClusterResultOutput) DiskEncryptionSetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.DiskEncryptionSetID }).(pulumi.StringPtrOutput)
}

// This cannot be updated once the Managed Cluster has been created.
func (o LookupManagedClusterResultOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

// (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at https://aka.ms/k8s/psp and https://aka.ms/aks/psp.
func (o LookupManagedClusterResultOutput) EnablePodSecurityPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnablePodSecurityPolicy }).(pulumi.BoolPtrOutput)
}

// Whether to enable Kubernetes Role-Based Access Control.
func (o LookupManagedClusterResultOutput) EnableRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *bool { return v.EnableRBAC }).(pulumi.BoolPtrOutput)
}

// The extended location of the Virtual Machine.
func (o LookupManagedClusterResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The FQDN of the master pool.
func (o LookupManagedClusterResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// This cannot be updated once the Managed Cluster has been created.
func (o LookupManagedClusterResultOutput) FqdnSubdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.FqdnSubdomain }).(pulumi.StringPtrOutput)
}

// Configurations for provisioning the cluster with HTTP proxy servers.
func (o LookupManagedClusterResultOutput) HttpProxyConfig() ManagedClusterHTTPProxyConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterHTTPProxyConfigResponse { return v.HttpProxyConfig }).(ManagedClusterHTTPProxyConfigResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupManagedClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the managed cluster, if configured.
func (o LookupManagedClusterResultOutput) Identity() ManagedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterIdentityResponse { return v.Identity }).(ManagedClusterIdentityResponsePtrOutput)
}

// Identities associated with the cluster.
func (o LookupManagedClusterResultOutput) IdentityProfile() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]UserAssignedIdentityResponse { return v.IdentityProfile }).(UserAssignedIdentityResponseMapOutput)
}

// Both patch version <major.minor.patch> (e.g. 1.20.13) and <major.minor> (e.g. 1.20) are supported. When <major.minor> is specified, the latest supported GA patch version is chosen automatically. Updating the cluster with the same <major.minor> once it has been created (e.g. 1.14.x -> 1.14) will not trigger an upgrade, even if a newer patch version is available. When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details.
func (o LookupManagedClusterResultOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

// The profile for Linux VMs in the Managed Cluster.
func (o LookupManagedClusterResultOutput) LinuxProfile() ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ContainerServiceLinuxProfileResponse { return v.LinuxProfile }).(ContainerServiceLinuxProfileResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupManagedClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The max number of agent pools for the managed cluster.
func (o LookupManagedClusterResultOutput) MaxAgentPools() pulumi.IntOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) int { return v.MaxAgentPools }).(pulumi.IntOutput)
}

// The name of the resource
func (o LookupManagedClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The network configuration profile.
func (o LookupManagedClusterResultOutput) NetworkProfile() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ContainerServiceNetworkProfileResponse { return v.NetworkProfile }).(ContainerServiceNetworkProfileResponsePtrOutput)
}

// The name of the resource group containing agent pool nodes.
func (o LookupManagedClusterResultOutput) NodeResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.NodeResourceGroup }).(pulumi.StringPtrOutput)
}

// The OIDC issuer profile of the Managed Cluster.
func (o LookupManagedClusterResultOutput) OidcIssuerProfile() ManagedClusterOIDCIssuerProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterOIDCIssuerProfileResponse {
		return v.OidcIssuerProfile
	}).(ManagedClusterOIDCIssuerProfileResponsePtrOutput)
}

// See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.
func (o LookupManagedClusterResultOutput) PodIdentityProfile() ManagedClusterPodIdentityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterPodIdentityProfileResponse {
		return v.PodIdentityProfile
	}).(ManagedClusterPodIdentityProfileResponsePtrOutput)
}

// The Power State of the cluster.
func (o LookupManagedClusterResultOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) PowerStateResponse { return v.PowerState }).(PowerStateResponseOutput)
}

// The FQDN of private cluster.
func (o LookupManagedClusterResultOutput) PrivateFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.PrivateFQDN }).(pulumi.StringOutput)
}

// Private link resources associated with the cluster.
func (o LookupManagedClusterResultOutput) PrivateLinkResources() PrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) []PrivateLinkResourceResponse { return v.PrivateLinkResources }).(PrivateLinkResourceResponseArrayOutput)
}

// The current provisioning state.
func (o LookupManagedClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Allow or deny public network access for AKS
func (o LookupManagedClusterResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The resourceUID uniquely identifies ManagedClusters that reuse ARM ResourceIds (i.e: create, delete, create sequence)
func (o LookupManagedClusterResultOutput) ResourceUID() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.ResourceUID }).(pulumi.StringOutput)
}

// Security profile for the managed cluster.
func (o LookupManagedClusterResultOutput) SecurityProfile() ManagedClusterSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterSecurityProfileResponse { return v.SecurityProfile }).(ManagedClusterSecurityProfileResponsePtrOutput)
}

// Service mesh profile for a managed cluster.
func (o LookupManagedClusterResultOutput) ServiceMeshProfile() ServiceMeshProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ServiceMeshProfileResponse { return v.ServiceMeshProfile }).(ServiceMeshProfileResponsePtrOutput)
}

// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
func (o LookupManagedClusterResultOutput) ServicePrincipalProfile() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterServicePrincipalProfileResponse {
		return v.ServicePrincipalProfile
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

// The managed cluster SKU.
func (o LookupManagedClusterResultOutput) Sku() ManagedClusterSKUResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterSKUResponse { return v.Sku }).(ManagedClusterSKUResponsePtrOutput)
}

// Storage profile for the managed cluster.
func (o LookupManagedClusterResultOutput) StorageProfile() ManagedClusterStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterStorageProfileResponse { return v.StorageProfile }).(ManagedClusterStorageProfileResponsePtrOutput)
}

// The support plan for the Managed Cluster. If unspecified, the default is 'KubernetesOfficial'.
func (o LookupManagedClusterResultOutput) SupportPlan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *string { return v.SupportPlan }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupManagedClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupManagedClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagedClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// Settings for upgrading a cluster.
func (o LookupManagedClusterResultOutput) UpgradeSettings() ClusterUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ClusterUpgradeSettingsResponse { return v.UpgradeSettings }).(ClusterUpgradeSettingsResponsePtrOutput)
}

// The profile for Windows VMs in the Managed Cluster.
func (o LookupManagedClusterResultOutput) WindowsProfile() ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterWindowsProfileResponse { return v.WindowsProfile }).(ManagedClusterWindowsProfileResponsePtrOutput)
}

// Workload Auto-scaler profile for the managed cluster.
func (o LookupManagedClusterResultOutput) WorkloadAutoScalerProfile() ManagedClusterWorkloadAutoScalerProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedClusterResult) *ManagedClusterWorkloadAutoScalerProfileResponse {
		return v.WorkloadAutoScalerProfile
	}).(ManagedClusterWorkloadAutoScalerProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedClusterResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Managed cluster.
type ManagedCluster struct {
	pulumi.CustomResourceState

	// Profile of Azure Active Directory configuration.
	AadProfile ManagedClusterAADProfileResponsePtrOutput `pulumi:"aadProfile"`
	// Profile of managed cluster add-on.
	AddonProfiles ManagedClusterAddonProfileResponseMapOutput `pulumi:"addonProfiles"`
	// Properties of the agent pool.
	AgentPoolProfiles ManagedClusterAgentPoolProfileResponseArrayOutput `pulumi:"agentPoolProfiles"`
	// Access profile for managed cluster API server.
	ApiServerAccessProfile ManagedClusterAPIServerAccessProfileResponsePtrOutput `pulumi:"apiServerAccessProfile"`
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix pulumi.StringPtrOutput `pulumi:"dnsPrefix"`
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy pulumi.BoolPtrOutput `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC pulumi.BoolPtrOutput `pulumi:"enableRBAC"`
	// FQDN for the master pool.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityResponsePtrOutput `pulumi:"identity"`
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion pulumi.StringPtrOutput `pulumi:"kubernetesVersion"`
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile ContainerServiceLinuxProfileResponsePtrOutput `pulumi:"linuxProfile"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The max number of agent pools for the managed cluster.
	MaxAgentPools pulumi.IntOutput `pulumi:"maxAgentPools"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Profile of network configuration.
	NetworkProfile ContainerServiceNetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup pulumi.StringPtrOutput `pulumi:"nodeResourceGroup"`
	// FQDN of private cluster.
	PrivateFQDN pulumi.StringOutput `pulumi:"privateFQDN"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile ManagedClusterServicePrincipalProfileResponsePtrOutput `pulumi:"servicePrincipalProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Profile for Windows VMs in the container service cluster.
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
			Type: pulumi.String("azure-native:containerservice/v20211001:ManagedCluster"),
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
	err := ctx.RegisterResource("azure-native:containerservice/v20191001:ManagedCluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:containerservice/v20191001:ManagedCluster", name, id, state, &resource, opts...)
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
	// Profile of Azure Active Directory configuration.
	AadProfile *ManagedClusterAADProfile `pulumi:"aadProfile"`
	// Profile of managed cluster add-on.
	AddonProfiles map[string]ManagedClusterAddonProfile `pulumi:"addonProfiles"`
	// Properties of the agent pool.
	AgentPoolProfiles []ManagedClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	// Access profile for managed cluster API server.
	ApiServerAccessProfile *ManagedClusterAPIServerAccessProfile `pulumi:"apiServerAccessProfile"`
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC *bool `pulumi:"enableRBAC"`
	// The identity of the managed cluster, if configured.
	Identity *ManagedClusterIdentity `pulumi:"identity"`
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile *ContainerServiceLinuxProfile `pulumi:"linuxProfile"`
	// Resource location
	Location *string `pulumi:"location"`
	// Profile of network configuration.
	NetworkProfile *ContainerServiceNetworkProfile `pulumi:"networkProfile"`
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName *string `pulumi:"resourceName"`
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfile `pulumi:"servicePrincipalProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Profile for Windows VMs in the container service cluster.
	WindowsProfile *ManagedClusterWindowsProfile `pulumi:"windowsProfile"`
}

// The set of arguments for constructing a ManagedCluster resource.
type ManagedClusterArgs struct {
	// Profile of Azure Active Directory configuration.
	AadProfile ManagedClusterAADProfilePtrInput
	// Profile of managed cluster add-on.
	AddonProfiles ManagedClusterAddonProfileMapInput
	// Properties of the agent pool.
	AgentPoolProfiles ManagedClusterAgentPoolProfileArrayInput
	// Access profile for managed cluster API server.
	ApiServerAccessProfile ManagedClusterAPIServerAccessProfilePtrInput
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix pulumi.StringPtrInput
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC pulumi.BoolPtrInput
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityPtrInput
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion pulumi.StringPtrInput
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile ContainerServiceLinuxProfilePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Profile of network configuration.
	NetworkProfile ContainerServiceNetworkProfilePtrInput
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringPtrInput
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile ManagedClusterServicePrincipalProfilePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Profile for Windows VMs in the container service cluster.
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

// Profile of Azure Active Directory configuration.
func (o ManagedClusterOutput) AadProfile() ManagedClusterAADProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAADProfileResponsePtrOutput { return v.AadProfile }).(ManagedClusterAADProfileResponsePtrOutput)
}

// Profile of managed cluster add-on.
func (o ManagedClusterOutput) AddonProfiles() ManagedClusterAddonProfileResponseMapOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAddonProfileResponseMapOutput { return v.AddonProfiles }).(ManagedClusterAddonProfileResponseMapOutput)
}

// Properties of the agent pool.
func (o ManagedClusterOutput) AgentPoolProfiles() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAgentPoolProfileResponseArrayOutput { return v.AgentPoolProfiles }).(ManagedClusterAgentPoolProfileResponseArrayOutput)
}

// Access profile for managed cluster API server.
func (o ManagedClusterOutput) ApiServerAccessProfile() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
		return v.ApiServerAccessProfile
	}).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

// DNS prefix specified when creating the managed cluster.
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

// FQDN for the master pool.
func (o ManagedClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// The identity of the managed cluster, if configured.
func (o ManagedClusterOutput) Identity() ManagedClusterIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterIdentityResponsePtrOutput { return v.Identity }).(ManagedClusterIdentityResponsePtrOutput)
}

// Version of Kubernetes specified when creating the managed cluster.
func (o ManagedClusterOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

// Profile for Linux VMs in the container service cluster.
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

// Profile of network configuration.
func (o ManagedClusterOutput) NetworkProfile() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ContainerServiceNetworkProfileResponsePtrOutput { return v.NetworkProfile }).(ContainerServiceNetworkProfileResponsePtrOutput)
}

// Name of the resource group containing agent pool nodes.
func (o ManagedClusterOutput) NodeResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.NodeResourceGroup }).(pulumi.StringPtrOutput)
}

// FQDN of private cluster.
func (o ManagedClusterOutput) PrivateFQDN() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.PrivateFQDN }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o ManagedClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
func (o ManagedClusterOutput) ServicePrincipalProfile() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterServicePrincipalProfileResponsePtrOutput {
		return v.ServicePrincipalProfile
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

// Resource tags
func (o ManagedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ManagedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Profile for Windows VMs in the container service cluster.
func (o ManagedClusterOutput) WindowsProfile() ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterWindowsProfileResponsePtrOutput { return v.WindowsProfile }).(ManagedClusterWindowsProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterOutput{})
}

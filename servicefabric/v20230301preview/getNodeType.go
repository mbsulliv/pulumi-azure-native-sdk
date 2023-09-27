// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Service Fabric node type of a given managed cluster.
func LookupNodeType(ctx *pulumi.Context, args *LookupNodeTypeArgs, opts ...pulumi.InvokeOption) (*LookupNodeTypeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNodeTypeResult
	err := ctx.Invoke("azure-native:servicefabric/v20230301preview:getNodeType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNodeTypeArgs struct {
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The name of the node type.
	NodeTypeName string `pulumi:"nodeTypeName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a node type in the cluster, each node type represents sub set of nodes in the cluster.
type LookupNodeTypeResult struct {
	// Additional managed data disks.
	AdditionalDataDisks []VmssDataDiskResponse `pulumi:"additionalDataDisks"`
	// The range of ports from which cluster assigned port to Service Fabric applications.
	ApplicationPorts *EndpointRangeDescriptionResponse `pulumi:"applicationPorts"`
	// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
	Capacities map[string]string `pulumi:"capacities"`
	// Managed data disk letter. It can not use the reserved letter C or D and it can not change after created.
	DataDiskLetter *string `pulumi:"dataDiskLetter"`
	// Disk size for the managed disk attached to the vms on the node type in GBs.
	DataDiskSizeGB *int `pulumi:"dataDiskSizeGB"`
	// Managed data disk type. Specifies the storage account type for the managed disk
	DataDiskType *string `pulumi:"dataDiskType"`
	// Specifies whether the network interface is accelerated networking-enabled.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Enable or disable the Host Encryption for the virtual machines on the node type. This will enable the encryption for all the disks including Resource/Temp disk at host itself. Default: The Encryption at host will be disabled unless this property is set to true for the resource.
	EnableEncryptionAtHost *bool `pulumi:"enableEncryptionAtHost"`
	// Specifies whether each node is allocated its own public IP address. This is only supported on secondary node types with custom Load Balancers.
	EnableNodePublicIP *bool `pulumi:"enableNodePublicIP"`
	// Specifies whether the node type should be overprovisioned. It is only allowed for stateless node types.
	EnableOverProvisioning *bool `pulumi:"enableOverProvisioning"`
	// The range of ephemeral ports that nodes in this node type should be configured with.
	EphemeralPorts *EndpointRangeDescriptionResponse `pulumi:"ephemeralPorts"`
	// Specifies the eviction policy for virtual machines in a SPOT node type. Default is Delete.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Indicates the node type uses its own frontend configurations instead of the default one for the cluster. This setting can only be specified for non-primary node types and can not be added or removed after the node type is created.
	FrontendConfigurations []FrontendConfigurationResponse `pulumi:"frontendConfigurations"`
	// Specifies the full host group resource Id. This property is used for deploying on azure dedicated hosts.
	HostGroupId *string `pulumi:"hostGroupId"`
	// Azure resource identifier.
	Id string `pulumi:"id"`
	// Indicates the Service Fabric system services for the cluster will run on this node type. This setting cannot be changed once the node type is created.
	IsPrimary bool `pulumi:"isPrimary"`
	// Indicates whether the node type will be Spot Virtual Machines. Azure will allocate the VMs if there is capacity available and the VMs can be evicted at any time.
	IsSpotVM *bool `pulumi:"isSpotVM"`
	// Indicates if the node type can only host Stateless workloads.
	IsStateless *bool `pulumi:"isStateless"`
	// Indicates if scale set associated with the node type can be composed of multiple placement groups.
	MultiplePlacementGroups *bool `pulumi:"multiplePlacementGroups"`
	// Azure resource name.
	Name string `pulumi:"name"`
	// Specifies the resource id of a NAT Gateway to attach to the subnet of this node type. Node type must use custom load balancer.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// The Network Security Rules for this node type. This setting can only be specified for node types that are configured with frontend configurations.
	NetworkSecurityRules []NetworkSecurityRuleResponse `pulumi:"networkSecurityRules"`
	// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
	PlacementProperties map[string]string `pulumi:"placementProperties"`
	// The provisioning state of the node type resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Specifies whether secure boot should be enabled on the nodeType. Can only be used with TrustedLaunch SecurityType
	SecureBootEnabled *bool `pulumi:"secureBootEnabled"`
	// Specifies the security type of the nodeType. Only TrustedLaunch is currently supported
	SecurityType *string `pulumi:"securityType"`
	// The node type sku.
	Sku *NodeTypeSkuResponse `pulumi:"sku"`
	// Indicates the time duration after which the platform will not try to restore the VMSS SPOT instances specified as ISO 8601.
	SpotRestoreTimeout *string `pulumi:"spotRestoreTimeout"`
	// Indicates the resource id of the subnet for the node type.
	SubnetId *string `pulumi:"subnetId"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type string `pulumi:"type"`
	// Specifies whether the use public load balancer. If not specified and the node type doesn't have its own frontend configuration, it will be attached to the default load balancer. If the node type uses its own Load balancer and useDefaultPublicLoadBalancer is true, then the frontend has to be an Internal Load Balancer. If the node type uses its own Load balancer and useDefaultPublicLoadBalancer is false or not set, then the custom load balancer must include a public load balancer to provide outbound connectivity.
	UseDefaultPublicLoadBalancer *bool `pulumi:"useDefaultPublicLoadBalancer"`
	// Indicates whether to use ephemeral os disk. The sku selected on the vmSize property needs to support this feature.
	UseEphemeralOSDisk *bool `pulumi:"useEphemeralOSDisk"`
	// Specifies whether to use the temporary disk for the service fabric data root, in which case no managed data disk will be attached and the temporary disk will be used. It is only allowed for stateless node types.
	UseTempDataDisk *bool `pulumi:"useTempDataDisk"`
	// Set of extensions that should be installed onto the virtual machines.
	VmExtensions []VMSSExtensionResponse `pulumi:"vmExtensions"`
	// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
	VmImageOffer *string `pulumi:"vmImageOffer"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use. In the Azure portal, find the marketplace image that you want to use and then click Want to deploy programmatically, Get Started ->. Enter any required information and then click Save.
	VmImagePlan *VmImagePlanResponse `pulumi:"vmImagePlan"`
	// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
	VmImagePublisher *string `pulumi:"vmImagePublisher"`
	// Indicates the resource id of the vm image. This parameter is used for custom vm image.
	VmImageResourceId *string `pulumi:"vmImageResourceId"`
	// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
	VmImageSku *string `pulumi:"vmImageSku"`
	// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
	VmImageVersion *string `pulumi:"vmImageVersion"`
	// The number of nodes in the node type. <br /><br />**Values:** <br />-1 - Use when auto scale rules are configured or sku.capacity is defined <br /> 0 - Not supported <br /> >0 - Use for manual scale.
	VmInstanceCount int `pulumi:"vmInstanceCount"`
	// Identities to assign to the virtual machine scale set under the node type.
	VmManagedIdentity *VmManagedIdentityResponse `pulumi:"vmManagedIdentity"`
	// The secrets to install in the virtual machines.
	VmSecrets []VaultSecretGroupResponse `pulumi:"vmSecrets"`
	// Specifies the actions to be performed on the vms before bootstrapping the service fabric runtime.
	VmSetupActions []string `pulumi:"vmSetupActions"`
	// Indicates the resource id of the vm shared galleries image. This parameter is used for custom vm image.
	VmSharedGalleryImageId *string `pulumi:"vmSharedGalleryImageId"`
	// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
	VmSize *string `pulumi:"vmSize"`
	// Specifies the availability zones where the node type would span across. If the cluster is not spanning across availability zones, initiates az migration for the cluster.
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupNodeTypeResult
func (val *LookupNodeTypeResult) Defaults() *LookupNodeTypeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EnableEncryptionAtHost == nil {
		enableEncryptionAtHost_ := false
		tmp.EnableEncryptionAtHost = &enableEncryptionAtHost_
	}
	if tmp.IsStateless == nil {
		isStateless_ := false
		tmp.IsStateless = &isStateless_
	}
	if tmp.MultiplePlacementGroups == nil {
		multiplePlacementGroups_ := false
		tmp.MultiplePlacementGroups = &multiplePlacementGroups_
	}
	return &tmp
}

func LookupNodeTypeOutput(ctx *pulumi.Context, args LookupNodeTypeOutputArgs, opts ...pulumi.InvokeOption) LookupNodeTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNodeTypeResult, error) {
			args := v.(LookupNodeTypeArgs)
			r, err := LookupNodeType(ctx, &args, opts...)
			var s LookupNodeTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNodeTypeResultOutput)
}

type LookupNodeTypeOutputArgs struct {
	// The name of the cluster resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the node type.
	NodeTypeName pulumi.StringInput `pulumi:"nodeTypeName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNodeTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeTypeArgs)(nil)).Elem()
}

// Describes a node type in the cluster, each node type represents sub set of nodes in the cluster.
type LookupNodeTypeResultOutput struct{ *pulumi.OutputState }

func (LookupNodeTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeTypeResult)(nil)).Elem()
}

func (o LookupNodeTypeResultOutput) ToLookupNodeTypeResultOutput() LookupNodeTypeResultOutput {
	return o
}

func (o LookupNodeTypeResultOutput) ToLookupNodeTypeResultOutputWithContext(ctx context.Context) LookupNodeTypeResultOutput {
	return o
}

func (o LookupNodeTypeResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNodeTypeResult] {
	return pulumix.Output[LookupNodeTypeResult]{
		OutputState: o.OutputState,
	}
}

// Additional managed data disks.
func (o LookupNodeTypeResultOutput) AdditionalDataDisks() VmssDataDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []VmssDataDiskResponse { return v.AdditionalDataDisks }).(VmssDataDiskResponseArrayOutput)
}

// The range of ports from which cluster assigned port to Service Fabric applications.
func (o LookupNodeTypeResultOutput) ApplicationPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *EndpointRangeDescriptionResponse { return v.ApplicationPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
func (o LookupNodeTypeResultOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) map[string]string { return v.Capacities }).(pulumi.StringMapOutput)
}

// Managed data disk letter. It can not use the reserved letter C or D and it can not change after created.
func (o LookupNodeTypeResultOutput) DataDiskLetter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.DataDiskLetter }).(pulumi.StringPtrOutput)
}

// Disk size for the managed disk attached to the vms on the node type in GBs.
func (o LookupNodeTypeResultOutput) DataDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *int { return v.DataDiskSizeGB }).(pulumi.IntPtrOutput)
}

// Managed data disk type. Specifies the storage account type for the managed disk
func (o LookupNodeTypeResultOutput) DataDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.DataDiskType }).(pulumi.StringPtrOutput)
}

// Specifies whether the network interface is accelerated networking-enabled.
func (o LookupNodeTypeResultOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

// Enable or disable the Host Encryption for the virtual machines on the node type. This will enable the encryption for all the disks including Resource/Temp disk at host itself. Default: The Encryption at host will be disabled unless this property is set to true for the resource.
func (o LookupNodeTypeResultOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

// Specifies whether each node is allocated its own public IP address. This is only supported on secondary node types with custom Load Balancers.
func (o LookupNodeTypeResultOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

// Specifies whether the node type should be overprovisioned. It is only allowed for stateless node types.
func (o LookupNodeTypeResultOutput) EnableOverProvisioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.EnableOverProvisioning }).(pulumi.BoolPtrOutput)
}

// The range of ephemeral ports that nodes in this node type should be configured with.
func (o LookupNodeTypeResultOutput) EphemeralPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *EndpointRangeDescriptionResponse { return v.EphemeralPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

// Specifies the eviction policy for virtual machines in a SPOT node type. Default is Delete.
func (o LookupNodeTypeResultOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

// Indicates the node type uses its own frontend configurations instead of the default one for the cluster. This setting can only be specified for non-primary node types and can not be added or removed after the node type is created.
func (o LookupNodeTypeResultOutput) FrontendConfigurations() FrontendConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []FrontendConfigurationResponse { return v.FrontendConfigurations }).(FrontendConfigurationResponseArrayOutput)
}

// Specifies the full host group resource Id. This property is used for deploying on azure dedicated hosts.
func (o LookupNodeTypeResultOutput) HostGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.HostGroupId }).(pulumi.StringPtrOutput)
}

// Azure resource identifier.
func (o LookupNodeTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates the Service Fabric system services for the cluster will run on this node type. This setting cannot be changed once the node type is created.
func (o LookupNodeTypeResultOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

// Indicates whether the node type will be Spot Virtual Machines. Azure will allocate the VMs if there is capacity available and the VMs can be evicted at any time.
func (o LookupNodeTypeResultOutput) IsSpotVM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.IsSpotVM }).(pulumi.BoolPtrOutput)
}

// Indicates if the node type can only host Stateless workloads.
func (o LookupNodeTypeResultOutput) IsStateless() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.IsStateless }).(pulumi.BoolPtrOutput)
}

// Indicates if scale set associated with the node type can be composed of multiple placement groups.
func (o LookupNodeTypeResultOutput) MultiplePlacementGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.MultiplePlacementGroups }).(pulumi.BoolPtrOutput)
}

// Azure resource name.
func (o LookupNodeTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the resource id of a NAT Gateway to attach to the subnet of this node type. Node type must use custom load balancer.
func (o LookupNodeTypeResultOutput) NatGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.NatGatewayId }).(pulumi.StringPtrOutput)
}

// The Network Security Rules for this node type. This setting can only be specified for node types that are configured with frontend configurations.
func (o LookupNodeTypeResultOutput) NetworkSecurityRules() NetworkSecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []NetworkSecurityRuleResponse { return v.NetworkSecurityRules }).(NetworkSecurityRuleResponseArrayOutput)
}

// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
func (o LookupNodeTypeResultOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) map[string]string { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

// The provisioning state of the node type resource.
func (o LookupNodeTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies whether secure boot should be enabled on the nodeType. Can only be used with TrustedLaunch SecurityType
func (o LookupNodeTypeResultOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.SecureBootEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies the security type of the nodeType. Only TrustedLaunch is currently supported
func (o LookupNodeTypeResultOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.SecurityType }).(pulumi.StringPtrOutput)
}

// The node type sku.
func (o LookupNodeTypeResultOutput) Sku() NodeTypeSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *NodeTypeSkuResponse { return v.Sku }).(NodeTypeSkuResponsePtrOutput)
}

// Indicates the time duration after which the platform will not try to restore the VMSS SPOT instances specified as ISO 8601.
func (o LookupNodeTypeResultOutput) SpotRestoreTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.SpotRestoreTimeout }).(pulumi.StringPtrOutput)
}

// Indicates the resource id of the subnet for the node type.
func (o LookupNodeTypeResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupNodeTypeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o LookupNodeTypeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o LookupNodeTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

// Specifies whether the use public load balancer. If not specified and the node type doesn't have its own frontend configuration, it will be attached to the default load balancer. If the node type uses its own Load balancer and useDefaultPublicLoadBalancer is true, then the frontend has to be an Internal Load Balancer. If the node type uses its own Load balancer and useDefaultPublicLoadBalancer is false or not set, then the custom load balancer must include a public load balancer to provide outbound connectivity.
func (o LookupNodeTypeResultOutput) UseDefaultPublicLoadBalancer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.UseDefaultPublicLoadBalancer }).(pulumi.BoolPtrOutput)
}

// Indicates whether to use ephemeral os disk. The sku selected on the vmSize property needs to support this feature.
func (o LookupNodeTypeResultOutput) UseEphemeralOSDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.UseEphemeralOSDisk }).(pulumi.BoolPtrOutput)
}

// Specifies whether to use the temporary disk for the service fabric data root, in which case no managed data disk will be attached and the temporary disk will be used. It is only allowed for stateless node types.
func (o LookupNodeTypeResultOutput) UseTempDataDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.UseTempDataDisk }).(pulumi.BoolPtrOutput)
}

// Set of extensions that should be installed onto the virtual machines.
func (o LookupNodeTypeResultOutput) VmExtensions() VMSSExtensionResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []VMSSExtensionResponse { return v.VmExtensions }).(VMSSExtensionResponseArrayOutput)
}

// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
func (o LookupNodeTypeResultOutput) VmImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageOffer }).(pulumi.StringPtrOutput)
}

// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use. In the Azure portal, find the marketplace image that you want to use and then click Want to deploy programmatically, Get Started ->. Enter any required information and then click Save.
func (o LookupNodeTypeResultOutput) VmImagePlan() VmImagePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *VmImagePlanResponse { return v.VmImagePlan }).(VmImagePlanResponsePtrOutput)
}

// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
func (o LookupNodeTypeResultOutput) VmImagePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImagePublisher }).(pulumi.StringPtrOutput)
}

// Indicates the resource id of the vm image. This parameter is used for custom vm image.
func (o LookupNodeTypeResultOutput) VmImageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageResourceId }).(pulumi.StringPtrOutput)
}

// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
func (o LookupNodeTypeResultOutput) VmImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageSku }).(pulumi.StringPtrOutput)
}

// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
func (o LookupNodeTypeResultOutput) VmImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageVersion }).(pulumi.StringPtrOutput)
}

// The number of nodes in the node type. <br /><br />**Values:** <br />-1 - Use when auto scale rules are configured or sku.capacity is defined <br /> 0 - Not supported <br /> >0 - Use for manual scale.
func (o LookupNodeTypeResultOutput) VmInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) int { return v.VmInstanceCount }).(pulumi.IntOutput)
}

// Identities to assign to the virtual machine scale set under the node type.
func (o LookupNodeTypeResultOutput) VmManagedIdentity() VmManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *VmManagedIdentityResponse { return v.VmManagedIdentity }).(VmManagedIdentityResponsePtrOutput)
}

// The secrets to install in the virtual machines.
func (o LookupNodeTypeResultOutput) VmSecrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []VaultSecretGroupResponse { return v.VmSecrets }).(VaultSecretGroupResponseArrayOutput)
}

// Specifies the actions to be performed on the vms before bootstrapping the service fabric runtime.
func (o LookupNodeTypeResultOutput) VmSetupActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []string { return v.VmSetupActions }).(pulumi.StringArrayOutput)
}

// Indicates the resource id of the vm shared galleries image. This parameter is used for custom vm image.
func (o LookupNodeTypeResultOutput) VmSharedGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmSharedGalleryImageId }).(pulumi.StringPtrOutput)
}

// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
func (o LookupNodeTypeResultOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

// Specifies the availability zones where the node type would span across. If the cluster is not spanning across availability zones, initiates az migration for the cluster.
func (o LookupNodeTypeResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNodeTypeResultOutput{})
}

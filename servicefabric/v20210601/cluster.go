// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The cluster resource
type Cluster struct {
	pulumi.CustomResourceState

	// The list of add-on features to enable in the cluster.
	AddOnFeatures pulumi.StringArrayOutput `pulumi:"addOnFeatures"`
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy ApplicationTypeVersionsCleanupPolicyResponsePtrOutput `pulumi:"applicationTypeVersionsCleanupPolicy"`
	// The Service Fabric runtime versions available for this cluster.
	AvailableClusterVersions ClusterVersionDetailsResponseArrayOutput `pulumi:"availableClusterVersions"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory AzureActiveDirectoryResponsePtrOutput `pulumi:"azureActiveDirectory"`
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate CertificateDescriptionResponsePtrOutput `pulumi:"certificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames ServerCertificateCommonNamesResponsePtrOutput `pulumi:"certificateCommonNames"`
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames ClientCertificateCommonNameResponseArrayOutput `pulumi:"clientCertificateCommonNames"`
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints ClientCertificateThumbprintResponseArrayOutput `pulumi:"clientCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](https://learn.microsoft.com/rest/api/servicefabric/cluster-versions/list). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion pulumi.StringPtrOutput `pulumi:"clusterCodeVersion"`
	// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
	ClusterEndpoint pulumi.StringOutput `pulumi:"clusterEndpoint"`
	// A service generated unique identifier for the cluster resource.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The current state of the cluster.
	//
	//   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
	//   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
	//   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
	//   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
	//   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
	//   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
	//   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
	//   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
	//   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
	//   - Ready - Indicates that the cluster is in a stable state.
	ClusterState pulumi.StringOutput `pulumi:"clusterState"`
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig DiagnosticsStorageAccountConfigResponsePtrOutput `pulumi:"diagnosticsStorageAccountConfig"`
	// Azure resource etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled pulumi.BoolPtrOutput `pulumi:"eventStoreServiceEnabled"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings SettingsSectionDescriptionResponseArrayOutput `pulumi:"fabricSettings"`
	// Indicates if infrastructure service manager is enabled.
	InfrastructureServiceManager pulumi.BoolPtrOutput `pulumi:"infrastructureServiceManager"`
	// Azure resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The http management endpoint of the cluster.
	ManagementEndpoint pulumi.StringOutput `pulumi:"managementEndpoint"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of node types in the cluster.
	NodeTypes NodeTypeDescriptionResponseArrayOutput `pulumi:"nodeTypes"`
	// Indicates a list of notification channels for cluster events.
	Notifications NotificationResponseArrayOutput `pulumi:"notifications"`
	// The provisioning state of the cluster resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel pulumi.StringPtrOutput `pulumi:"reliabilityLevel"`
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate CertificateDescriptionResponsePtrOutput `pulumi:"reverseProxyCertificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames ServerCertificateCommonNamesResponsePtrOutput `pulumi:"reverseProxyCertificateCommonNames"`
	// This property controls the logical grouping of VMs in upgrade domains (UDs). This property can't be modified if a node type with multiple Availability Zones is already present in the cluster.
	SfZonalUpgradeMode pulumi.StringPtrOutput `pulumi:"sfZonalUpgradeMode"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The policy to use when upgrading the cluster.
	UpgradeDescription ClusterUpgradePolicyResponsePtrOutput `pulumi:"upgradeDescription"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	UpgradeMode pulumi.StringPtrOutput `pulumi:"upgradeMode"`
	// Indicates the end date and time to pause automatic runtime version upgrades on the cluster for an specific period of time on the cluster (UTC).
	UpgradePauseEndTimestampUtc pulumi.StringPtrOutput `pulumi:"upgradePauseEndTimestampUtc"`
	// Indicates the start date and time to pause automatic runtime version upgrades on the cluster for an specific period of time on the cluster (UTC).
	UpgradePauseStartTimestampUtc pulumi.StringPtrOutput `pulumi:"upgradePauseStartTimestampUtc"`
	// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **upgradeMode** is set to 'Automatic'.
	UpgradeWave pulumi.StringPtrOutput `pulumi:"upgradeWave"`
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage pulumi.StringPtrOutput `pulumi:"vmImage"`
	// This property defines the upgrade mode for the virtual machine scale set, it is mandatory if a node type with multiple Availability Zones is added.
	VmssZonalUpgradeMode pulumi.StringPtrOutput `pulumi:"vmssZonalUpgradeMode"`
	// Boolean to pause automatic runtime version upgrades to the cluster.
	WaveUpgradePaused pulumi.BoolPtrOutput `pulumi:"waveUpgradePaused"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ManagementEndpoint'")
	}
	if args.NodeTypes == nil {
		return nil, errors.New("invalid value for required argument 'NodeTypes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UpgradeDescription != nil {
		args.UpgradeDescription = args.UpgradeDescription.ToClusterUpgradePolicyPtrOutput().ApplyT(func(v *ClusterUpgradePolicy) *ClusterUpgradePolicy { return v.Defaults() }).(ClusterUpgradePolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20160901:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20170701preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20180201:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190601preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20191101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200301:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20201201preview:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:servicefabric/v20210601:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:servicefabric/v20210601:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The list of add-on features to enable in the cluster.
	AddOnFeatures []string `pulumi:"addOnFeatures"`
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy *ApplicationTypeVersionsCleanupPolicy `pulumi:"applicationTypeVersionsCleanupPolicy"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory *AzureActiveDirectory `pulumi:"azureActiveDirectory"`
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate *CertificateDescription `pulumi:"certificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames *ServerCertificateCommonNames `pulumi:"certificateCommonNames"`
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames []ClientCertificateCommonName `pulumi:"clientCertificateCommonNames"`
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints []ClientCertificateThumbprint `pulumi:"clientCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](https://learn.microsoft.com/rest/api/servicefabric/cluster-versions/list). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// The name of the cluster resource.
	ClusterName *string `pulumi:"clusterName"`
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig *DiagnosticsStorageAccountConfig `pulumi:"diagnosticsStorageAccountConfig"`
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled *bool `pulumi:"eventStoreServiceEnabled"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings []SettingsSectionDescription `pulumi:"fabricSettings"`
	// Indicates if infrastructure service manager is enabled.
	InfrastructureServiceManager *bool `pulumi:"infrastructureServiceManager"`
	// Azure resource location.
	Location *string `pulumi:"location"`
	// The http management endpoint of the cluster.
	ManagementEndpoint string `pulumi:"managementEndpoint"`
	// The list of node types in the cluster.
	NodeTypes []NodeTypeDescription `pulumi:"nodeTypes"`
	// Indicates a list of notification channels for cluster events.
	Notifications []Notification `pulumi:"notifications"`
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel *string `pulumi:"reliabilityLevel"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate *CertificateDescription `pulumi:"reverseProxyCertificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames *ServerCertificateCommonNames `pulumi:"reverseProxyCertificateCommonNames"`
	// This property controls the logical grouping of VMs in upgrade domains (UDs). This property can't be modified if a node type with multiple Availability Zones is already present in the cluster.
	SfZonalUpgradeMode *string `pulumi:"sfZonalUpgradeMode"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The policy to use when upgrading the cluster.
	UpgradeDescription *ClusterUpgradePolicy `pulumi:"upgradeDescription"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	UpgradeMode *string `pulumi:"upgradeMode"`
	// Indicates the end date and time to pause automatic runtime version upgrades on the cluster for an specific period of time on the cluster (UTC).
	UpgradePauseEndTimestampUtc *string `pulumi:"upgradePauseEndTimestampUtc"`
	// Indicates the start date and time to pause automatic runtime version upgrades on the cluster for an specific period of time on the cluster (UTC).
	UpgradePauseStartTimestampUtc *string `pulumi:"upgradePauseStartTimestampUtc"`
	// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **upgradeMode** is set to 'Automatic'.
	UpgradeWave *string `pulumi:"upgradeWave"`
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage *string `pulumi:"vmImage"`
	// This property defines the upgrade mode for the virtual machine scale set, it is mandatory if a node type with multiple Availability Zones is added.
	VmssZonalUpgradeMode *string `pulumi:"vmssZonalUpgradeMode"`
	// Boolean to pause automatic runtime version upgrades to the cluster.
	WaveUpgradePaused *bool `pulumi:"waveUpgradePaused"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The list of add-on features to enable in the cluster.
	AddOnFeatures pulumi.StringArrayInput
	// The policy used to clean up unused versions.
	ApplicationTypeVersionsCleanupPolicy ApplicationTypeVersionsCleanupPolicyPtrInput
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory AzureActiveDirectoryPtrInput
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate CertificateDescriptionPtrInput
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames ServerCertificateCommonNamesPtrInput
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames ClientCertificateCommonNameArrayInput
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints ClientCertificateThumbprintArrayInput
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](https://learn.microsoft.com/rest/api/servicefabric/cluster-versions/list). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion pulumi.StringPtrInput
	// The name of the cluster resource.
	ClusterName pulumi.StringPtrInput
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig DiagnosticsStorageAccountConfigPtrInput
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled pulumi.BoolPtrInput
	// The list of custom fabric settings to configure the cluster.
	FabricSettings SettingsSectionDescriptionArrayInput
	// Indicates if infrastructure service manager is enabled.
	InfrastructureServiceManager pulumi.BoolPtrInput
	// Azure resource location.
	Location pulumi.StringPtrInput
	// The http management endpoint of the cluster.
	ManagementEndpoint pulumi.StringInput
	// The list of node types in the cluster.
	NodeTypes NodeTypeDescriptionArrayInput
	// Indicates a list of notification channels for cluster events.
	Notifications NotificationArrayInput
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate CertificateDescriptionPtrInput
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames ServerCertificateCommonNamesPtrInput
	// This property controls the logical grouping of VMs in upgrade domains (UDs). This property can't be modified if a node type with multiple Availability Zones is already present in the cluster.
	SfZonalUpgradeMode pulumi.StringPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// The policy to use when upgrading the cluster.
	UpgradeDescription ClusterUpgradePolicyPtrInput
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	UpgradeMode pulumi.StringPtrInput
	// Indicates the end date and time to pause automatic runtime version upgrades on the cluster for an specific period of time on the cluster (UTC).
	UpgradePauseEndTimestampUtc pulumi.StringPtrInput
	// Indicates the start date and time to pause automatic runtime version upgrades on the cluster for an specific period of time on the cluster (UTC).
	UpgradePauseStartTimestampUtc pulumi.StringPtrInput
	// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **upgradeMode** is set to 'Automatic'.
	UpgradeWave pulumi.StringPtrInput
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage pulumi.StringPtrInput
	// This property defines the upgrade mode for the virtual machine scale set, it is mandatory if a node type with multiple Availability Zones is added.
	VmssZonalUpgradeMode pulumi.StringPtrInput
	// Boolean to pause automatic runtime version upgrades to the cluster.
	WaveUpgradePaused pulumi.BoolPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The list of add-on features to enable in the cluster.
func (o ClusterOutput) AddOnFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.AddOnFeatures }).(pulumi.StringArrayOutput)
}

// The policy used to clean up unused versions.
func (o ClusterOutput) ApplicationTypeVersionsCleanupPolicy() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
		return v.ApplicationTypeVersionsCleanupPolicy
	}).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

// The Service Fabric runtime versions available for this cluster.
func (o ClusterOutput) AvailableClusterVersions() ClusterVersionDetailsResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterVersionDetailsResponseArrayOutput { return v.AvailableClusterVersions }).(ClusterVersionDetailsResponseArrayOutput)
}

// The AAD authentication settings of the cluster.
func (o ClusterOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) AzureActiveDirectoryResponsePtrOutput { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
func (o ClusterOutput) Certificate() CertificateDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) CertificateDescriptionResponsePtrOutput { return v.Certificate }).(CertificateDescriptionResponsePtrOutput)
}

// Describes a list of server certificates referenced by common name that are used to secure the cluster.
func (o ClusterOutput) CertificateCommonNames() ServerCertificateCommonNamesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ServerCertificateCommonNamesResponsePtrOutput { return v.CertificateCommonNames }).(ServerCertificateCommonNamesResponsePtrOutput)
}

// The list of client certificates referenced by common name that are allowed to manage the cluster.
func (o ClusterOutput) ClientCertificateCommonNames() ClientCertificateCommonNameResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ClientCertificateCommonNameResponseArrayOutput { return v.ClientCertificateCommonNames }).(ClientCertificateCommonNameResponseArrayOutput)
}

// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
func (o ClusterOutput) ClientCertificateThumbprints() ClientCertificateThumbprintResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) ClientCertificateThumbprintResponseArrayOutput { return v.ClientCertificateThumbprints }).(ClientCertificateThumbprintResponseArrayOutput)
}

// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](https://learn.microsoft.com/rest/api/servicefabric/cluster-versions/list). To get the list of available version for existing clusters use **availableClusterVersions**.
func (o ClusterOutput) ClusterCodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ClusterCodeVersion }).(pulumi.StringPtrOutput)
}

// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
func (o ClusterOutput) ClusterEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterEndpoint }).(pulumi.StringOutput)
}

// A service generated unique identifier for the cluster resource.
func (o ClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The current state of the cluster.
//
//   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
//   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
//   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
//   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
//   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
//   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
//   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
//   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
//   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
//   - Ready - Indicates that the cluster is in a stable state.
func (o ClusterOutput) ClusterState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterState }).(pulumi.StringOutput)
}

// The storage account information for storing Service Fabric diagnostic logs.
func (o ClusterOutput) DiagnosticsStorageAccountConfig() DiagnosticsStorageAccountConfigResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) DiagnosticsStorageAccountConfigResponsePtrOutput {
		return v.DiagnosticsStorageAccountConfig
	}).(DiagnosticsStorageAccountConfigResponsePtrOutput)
}

// Azure resource etag.
func (o ClusterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Indicates if the event store service is enabled.
func (o ClusterOutput) EventStoreServiceEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EventStoreServiceEnabled }).(pulumi.BoolPtrOutput)
}

// The list of custom fabric settings to configure the cluster.
func (o ClusterOutput) FabricSettings() SettingsSectionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) SettingsSectionDescriptionResponseArrayOutput { return v.FabricSettings }).(SettingsSectionDescriptionResponseArrayOutput)
}

// Indicates if infrastructure service manager is enabled.
func (o ClusterOutput) InfrastructureServiceManager() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.InfrastructureServiceManager }).(pulumi.BoolPtrOutput)
}

// Azure resource location.
func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The http management endpoint of the cluster.
func (o ClusterOutput) ManagementEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ManagementEndpoint }).(pulumi.StringOutput)
}

// Azure resource name.
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of node types in the cluster.
func (o ClusterOutput) NodeTypes() NodeTypeDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) NodeTypeDescriptionResponseArrayOutput { return v.NodeTypes }).(NodeTypeDescriptionResponseArrayOutput)
}

// Indicates a list of notification channels for cluster events.
func (o ClusterOutput) Notifications() NotificationResponseArrayOutput {
	return o.ApplyT(func(v *Cluster) NotificationResponseArrayOutput { return v.Notifications }).(NotificationResponseArrayOutput)
}

// The provisioning state of the cluster resource.
func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
//
//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
//   - Silver - Run the System services with a target replica set count of 5.
//   - Gold - Run the System services with a target replica set count of 7.
//   - Platinum - Run the System services with a target replica set count of 9.
func (o ClusterOutput) ReliabilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ReliabilityLevel }).(pulumi.StringPtrOutput)
}

// The server certificate used by reverse proxy.
func (o ClusterOutput) ReverseProxyCertificate() CertificateDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) CertificateDescriptionResponsePtrOutput { return v.ReverseProxyCertificate }).(CertificateDescriptionResponsePtrOutput)
}

// Describes a list of server certificates referenced by common name that are used to secure the cluster.
func (o ClusterOutput) ReverseProxyCertificateCommonNames() ServerCertificateCommonNamesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ServerCertificateCommonNamesResponsePtrOutput {
		return v.ReverseProxyCertificateCommonNames
	}).(ServerCertificateCommonNamesResponsePtrOutput)
}

// This property controls the logical grouping of VMs in upgrade domains (UDs). This property can't be modified if a node type with multiple Availability Zones is already present in the cluster.
func (o ClusterOutput) SfZonalUpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SfZonalUpgradeMode }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The policy to use when upgrading the cluster.
func (o ClusterOutput) UpgradeDescription() ClusterUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterUpgradePolicyResponsePtrOutput { return v.UpgradeDescription }).(ClusterUpgradePolicyResponsePtrOutput)
}

// The upgrade mode of the cluster when new Service Fabric runtime version is available.
func (o ClusterOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

// Indicates the end date and time to pause automatic runtime version upgrades on the cluster for an specific period of time on the cluster (UTC).
func (o ClusterOutput) UpgradePauseEndTimestampUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.UpgradePauseEndTimestampUtc }).(pulumi.StringPtrOutput)
}

// Indicates the start date and time to pause automatic runtime version upgrades on the cluster for an specific period of time on the cluster (UTC).
func (o ClusterOutput) UpgradePauseStartTimestampUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.UpgradePauseStartTimestampUtc }).(pulumi.StringPtrOutput)
}

// Indicates when new cluster runtime version upgrades will be applied after they are released. By default is Wave0. Only applies when **upgradeMode** is set to 'Automatic'.
func (o ClusterOutput) UpgradeWave() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.UpgradeWave }).(pulumi.StringPtrOutput)
}

// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
func (o ClusterOutput) VmImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.VmImage }).(pulumi.StringPtrOutput)
}

// This property defines the upgrade mode for the virtual machine scale set, it is mandatory if a node type with multiple Availability Zones is added.
func (o ClusterOutput) VmssZonalUpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.VmssZonalUpgradeMode }).(pulumi.StringPtrOutput)
}

// Boolean to pause automatic runtime version upgrades to the cluster.
func (o ClusterOutput) WaveUpgradePaused() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.WaveUpgradePaused }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}

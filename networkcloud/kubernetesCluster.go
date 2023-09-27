// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2023-05-01-preview.
type KubernetesCluster struct {
	pulumi.CustomResourceState

	// The Azure Active Directory Integration properties.
	AadConfiguration AadConfigurationResponsePtrOutput `pulumi:"aadConfiguration"`
	// The administrative credentials that will be applied to the control plane and agent pool nodes that do not specify their own values.
	AdministratorConfiguration AdministratorConfigurationResponsePtrOutput `pulumi:"administratorConfiguration"`
	// The full list of network resource IDs that are attached to this cluster, including those attached only to specific agent pools.
	AttachedNetworkIds pulumi.StringArrayOutput `pulumi:"attachedNetworkIds"`
	// The list of versions that this Kubernetes cluster can be upgraded to.
	AvailableUpgrades AvailableUpgradeResponseArrayOutput `pulumi:"availableUpgrades"`
	// The resource ID of the Network Cloud cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The resource ID of the connected cluster set up when this Kubernetes cluster is created.
	ConnectedClusterId pulumi.StringOutput `pulumi:"connectedClusterId"`
	// The current running version of Kubernetes on the control plane.
	ControlPlaneKubernetesVersion pulumi.StringOutput `pulumi:"controlPlaneKubernetesVersion"`
	// The defining characteristics of the control plane for this Kubernetes Cluster.
	ControlPlaneNodeConfiguration ControlPlaneNodeConfigurationResponseOutput `pulumi:"controlPlaneNodeConfiguration"`
	// The current status of the Kubernetes cluster.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The current feature settings.
	FeatureStatuses FeatureStatusResponseArrayOutput `pulumi:"featureStatuses"`
	// The agent pools that are created with this Kubernetes cluster for running critical system services and workloads. This data in this field is only used during creation, and the field will be empty following the creation of the Kubernetes Cluster. After creation, the management of agent pools is done using the agentPools sub-resource.
	InitialAgentPoolConfigurations InitialAgentPoolConfigurationResponseArrayOutput `pulumi:"initialAgentPoolConfigurations"`
	// The Kubernetes version for this cluster. Accepts n.n, n.n.n, and n.n.n-n format. The interpreted version used will be resolved into this field after creation or update.
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration ManagedResourceGroupConfigurationResponsePtrOutput `pulumi:"managedResourceGroupConfiguration"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The configuration of the Kubernetes cluster networking, including the attachment of networks that span the cluster.
	NetworkConfiguration NetworkConfigurationResponseOutput `pulumi:"networkConfiguration"`
	// The details of the nodes in this cluster.
	Nodes KubernetesClusterNodeResponseArrayOutput `pulumi:"nodes"`
	// The provisioning state of the Kubernetes cluster resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKubernetesCluster registers a new resource with the given unique name, arguments, and options.
func NewKubernetesCluster(ctx *pulumi.Context,
	name string, args *KubernetesClusterArgs, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneNodeConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneNodeConfiguration'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.InitialAgentPoolConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'InitialAgentPoolConfigurations'")
	}
	if args.KubernetesVersion == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesVersion'")
	}
	if args.NetworkConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConfiguration'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.NetworkConfiguration = args.NetworkConfiguration.ToNetworkConfigurationOutput().ApplyT(func(v NetworkConfiguration) NetworkConfiguration { return *v.Defaults() }).(NetworkConfigurationOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud/v20230501preview:KubernetesCluster"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230701:KubernetesCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource KubernetesCluster
	err := ctx.RegisterResource("azure-native:networkcloud:KubernetesCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesCluster gets an existing KubernetesCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesClusterState, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	var resource KubernetesCluster
	err := ctx.ReadResource("azure-native:networkcloud:KubernetesCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesCluster resources.
type kubernetesClusterState struct {
}

type KubernetesClusterState struct {
}

func (KubernetesClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterState)(nil)).Elem()
}

type kubernetesClusterArgs struct {
	// The Azure Active Directory Integration properties.
	AadConfiguration *AadConfiguration `pulumi:"aadConfiguration"`
	// The administrative credentials that will be applied to the control plane and agent pool nodes that do not specify their own values.
	AdministratorConfiguration *AdministratorConfiguration `pulumi:"administratorConfiguration"`
	// The defining characteristics of the control plane for this Kubernetes Cluster.
	ControlPlaneNodeConfiguration ControlPlaneNodeConfiguration `pulumi:"controlPlaneNodeConfiguration"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The agent pools that are created with this Kubernetes cluster for running critical system services and workloads. This data in this field is only used during creation, and the field will be empty following the creation of the Kubernetes Cluster. After creation, the management of agent pools is done using the agentPools sub-resource.
	InitialAgentPoolConfigurations []InitialAgentPoolConfiguration `pulumi:"initialAgentPoolConfigurations"`
	// The name of the Kubernetes cluster.
	KubernetesClusterName *string `pulumi:"kubernetesClusterName"`
	// The Kubernetes version for this cluster. Accepts n.n, n.n.n, and n.n.n-n format. The interpreted version used will be resolved into this field after creation or update.
	KubernetesVersion string `pulumi:"kubernetesVersion"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfiguration `pulumi:"managedResourceGroupConfiguration"`
	// The configuration of the Kubernetes cluster networking, including the attachment of networks that span the cluster.
	NetworkConfiguration NetworkConfiguration `pulumi:"networkConfiguration"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a KubernetesCluster resource.
type KubernetesClusterArgs struct {
	// The Azure Active Directory Integration properties.
	AadConfiguration AadConfigurationPtrInput
	// The administrative credentials that will be applied to the control plane and agent pool nodes that do not specify their own values.
	AdministratorConfiguration AdministratorConfigurationPtrInput
	// The defining characteristics of the control plane for this Kubernetes Cluster.
	ControlPlaneNodeConfiguration ControlPlaneNodeConfigurationInput
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// The agent pools that are created with this Kubernetes cluster for running critical system services and workloads. This data in this field is only used during creation, and the field will be empty following the creation of the Kubernetes Cluster. After creation, the management of agent pools is done using the agentPools sub-resource.
	InitialAgentPoolConfigurations InitialAgentPoolConfigurationArrayInput
	// The name of the Kubernetes cluster.
	KubernetesClusterName pulumi.StringPtrInput
	// The Kubernetes version for this cluster. Accepts n.n, n.n.n, and n.n.n-n format. The interpreted version used will be resolved into this field after creation or update.
	KubernetesVersion pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The configuration of the managed resource group associated with the resource.
	ManagedResourceGroupConfiguration ManagedResourceGroupConfigurationPtrInput
	// The configuration of the Kubernetes cluster networking, including the attachment of networks that span the cluster.
	NetworkConfiguration NetworkConfigurationInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (KubernetesClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterArgs)(nil)).Elem()
}

type KubernetesClusterInput interface {
	pulumi.Input

	ToKubernetesClusterOutput() KubernetesClusterOutput
	ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput
}

func (*KubernetesCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesCluster)(nil)).Elem()
}

func (i *KubernetesCluster) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return i.ToKubernetesClusterOutputWithContext(context.Background())
}

func (i *KubernetesCluster) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterOutput)
}

func (i *KubernetesCluster) ToOutput(ctx context.Context) pulumix.Output[*KubernetesCluster] {
	return pulumix.Output[*KubernetesCluster]{
		OutputState: i.ToKubernetesClusterOutputWithContext(ctx).OutputState,
	}
}

type KubernetesClusterOutput struct{ *pulumi.OutputState }

func (KubernetesClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesCluster)(nil)).Elem()
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return o
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return o
}

func (o KubernetesClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*KubernetesCluster] {
	return pulumix.Output[*KubernetesCluster]{
		OutputState: o.OutputState,
	}
}

// The Azure Active Directory Integration properties.
func (o KubernetesClusterOutput) AadConfiguration() AadConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) AadConfigurationResponsePtrOutput { return v.AadConfiguration }).(AadConfigurationResponsePtrOutput)
}

// The administrative credentials that will be applied to the control plane and agent pool nodes that do not specify their own values.
func (o KubernetesClusterOutput) AdministratorConfiguration() AdministratorConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) AdministratorConfigurationResponsePtrOutput {
		return v.AdministratorConfiguration
	}).(AdministratorConfigurationResponsePtrOutput)
}

// The full list of network resource IDs that are attached to this cluster, including those attached only to specific agent pools.
func (o KubernetesClusterOutput) AttachedNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringArrayOutput { return v.AttachedNetworkIds }).(pulumi.StringArrayOutput)
}

// The list of versions that this Kubernetes cluster can be upgraded to.
func (o KubernetesClusterOutput) AvailableUpgrades() AvailableUpgradeResponseArrayOutput {
	return o.ApplyT(func(v *KubernetesCluster) AvailableUpgradeResponseArrayOutput { return v.AvailableUpgrades }).(AvailableUpgradeResponseArrayOutput)
}

// The resource ID of the Network Cloud cluster.
func (o KubernetesClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The resource ID of the connected cluster set up when this Kubernetes cluster is created.
func (o KubernetesClusterOutput) ConnectedClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.ConnectedClusterId }).(pulumi.StringOutput)
}

// The current running version of Kubernetes on the control plane.
func (o KubernetesClusterOutput) ControlPlaneKubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.ControlPlaneKubernetesVersion }).(pulumi.StringOutput)
}

// The defining characteristics of the control plane for this Kubernetes Cluster.
func (o KubernetesClusterOutput) ControlPlaneNodeConfiguration() ControlPlaneNodeConfigurationResponseOutput {
	return o.ApplyT(func(v *KubernetesCluster) ControlPlaneNodeConfigurationResponseOutput {
		return v.ControlPlaneNodeConfiguration
	}).(ControlPlaneNodeConfigurationResponseOutput)
}

// The current status of the Kubernetes cluster.
func (o KubernetesClusterOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o KubernetesClusterOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o KubernetesClusterOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *KubernetesCluster) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The current feature settings.
func (o KubernetesClusterOutput) FeatureStatuses() FeatureStatusResponseArrayOutput {
	return o.ApplyT(func(v *KubernetesCluster) FeatureStatusResponseArrayOutput { return v.FeatureStatuses }).(FeatureStatusResponseArrayOutput)
}

// The agent pools that are created with this Kubernetes cluster for running critical system services and workloads. This data in this field is only used during creation, and the field will be empty following the creation of the Kubernetes Cluster. After creation, the management of agent pools is done using the agentPools sub-resource.
func (o KubernetesClusterOutput) InitialAgentPoolConfigurations() InitialAgentPoolConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *KubernetesCluster) InitialAgentPoolConfigurationResponseArrayOutput {
		return v.InitialAgentPoolConfigurations
	}).(InitialAgentPoolConfigurationResponseArrayOutput)
}

// The Kubernetes version for this cluster. Accepts n.n, n.n.n, and n.n.n-n format. The interpreted version used will be resolved into this field after creation or update.
func (o KubernetesClusterOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o KubernetesClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The configuration of the managed resource group associated with the resource.
func (o KubernetesClusterOutput) ManagedResourceGroupConfiguration() ManagedResourceGroupConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesCluster) ManagedResourceGroupConfigurationResponsePtrOutput {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedResourceGroupConfigurationResponsePtrOutput)
}

// The name of the resource
func (o KubernetesClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The configuration of the Kubernetes cluster networking, including the attachment of networks that span the cluster.
func (o KubernetesClusterOutput) NetworkConfiguration() NetworkConfigurationResponseOutput {
	return o.ApplyT(func(v *KubernetesCluster) NetworkConfigurationResponseOutput { return v.NetworkConfiguration }).(NetworkConfigurationResponseOutput)
}

// The details of the nodes in this cluster.
func (o KubernetesClusterOutput) Nodes() KubernetesClusterNodeResponseArrayOutput {
	return o.ApplyT(func(v *KubernetesCluster) KubernetesClusterNodeResponseArrayOutput { return v.Nodes }).(KubernetesClusterNodeResponseArrayOutput)
}

// The provisioning state of the Kubernetes cluster resource.
func (o KubernetesClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o KubernetesClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KubernetesCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o KubernetesClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o KubernetesClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KubernetesClusterOutput{})
}

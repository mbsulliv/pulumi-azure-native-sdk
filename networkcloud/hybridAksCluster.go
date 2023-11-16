// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The details are specific to the Network Cloud use of the Hybrid AKS cluster.
// Azure REST API version: 2022-12-12-preview. Prior API version in Azure Native 1.x: 2022-12-12-preview.
type HybridAksCluster struct {
	pulumi.CustomResourceState

	// The list of resource IDs for the workload networks associated with the Hybrid AKS cluster. It can be any of l2Networks, l3Networks, or trunkedNetworks resources. This field will also contain one cloudServicesNetwork and one defaultCniNetwork.
	AssociatedNetworkIds pulumi.StringArrayOutput `pulumi:"associatedNetworkIds"`
	// The resource ID of the associated cloud services network.
	CloudServicesNetworkId pulumi.StringOutput `pulumi:"cloudServicesNetworkId"`
	// The resource ID of the Network Cloud cluster hosting the Hybrid AKS cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The number of control plane node VMs.
	ControlPlaneCount pulumi.Float64Output `pulumi:"controlPlaneCount"`
	// The list of node configurations detailing associated VMs that are part of the control plane nodes of this Hybrid AKS cluster.
	ControlPlaneNodes NodeConfigurationResponseArrayOutput `pulumi:"controlPlaneNodes"`
	// The resource ID of the associated default CNI network.
	DefaultCniNetworkId pulumi.StringOutput `pulumi:"defaultCniNetworkId"`
	// The more detailed status of this Hybrid AKS cluster.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The resource ID of the Hybrid AKS cluster that this additional information is for.
	HybridAksProvisionedClusterId pulumi.StringOutput `pulumi:"hybridAksProvisionedClusterId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the Hybrid AKS cluster resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The resource IDs of volumes that are attached to the Hybrid AKS cluster.
	Volumes pulumi.StringArrayOutput `pulumi:"volumes"`
	// The number of worker node VMs.
	WorkerCount pulumi.Float64Output `pulumi:"workerCount"`
	// The list of node configurations detailing associated VMs that are part of the worker nodes of this Hybrid AKS cluster.
	WorkerNodes NodeConfigurationResponseArrayOutput `pulumi:"workerNodes"`
}

// NewHybridAksCluster registers a new resource with the given unique name, arguments, and options.
func NewHybridAksCluster(ctx *pulumi.Context,
	name string, args *HybridAksClusterArgs, opts ...pulumi.ResourceOption) (*HybridAksCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssociatedNetworkIds == nil {
		return nil, errors.New("invalid value for required argument 'AssociatedNetworkIds'")
	}
	if args.ControlPlaneCount == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneCount'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.HybridAksProvisionedClusterId == nil {
		return nil, errors.New("invalid value for required argument 'HybridAksProvisionedClusterId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkerCount == nil {
		return nil, errors.New("invalid value for required argument 'WorkerCount'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:HybridAksCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HybridAksCluster
	err := ctx.RegisterResource("azure-native:networkcloud:HybridAksCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridAksCluster gets an existing HybridAksCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridAksCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridAksClusterState, opts ...pulumi.ResourceOption) (*HybridAksCluster, error) {
	var resource HybridAksCluster
	err := ctx.ReadResource("azure-native:networkcloud:HybridAksCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridAksCluster resources.
type hybridAksClusterState struct {
}

type HybridAksClusterState struct {
}

func (HybridAksClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridAksClusterState)(nil)).Elem()
}

type hybridAksClusterArgs struct {
	// The list of resource IDs for the workload networks associated with the Hybrid AKS cluster. It can be any of l2Networks, l3Networks, or trunkedNetworks resources. This field will also contain one cloudServicesNetwork and one defaultCniNetwork.
	AssociatedNetworkIds []string `pulumi:"associatedNetworkIds"`
	// The number of control plane node VMs.
	ControlPlaneCount float64 `pulumi:"controlPlaneCount"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The name of the Hybrid AKS cluster.
	HybridAksClusterName *string `pulumi:"hybridAksClusterName"`
	// The resource ID of the Hybrid AKS cluster that this additional information is for.
	HybridAksProvisionedClusterId string `pulumi:"hybridAksProvisionedClusterId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The number of worker node VMs.
	WorkerCount float64 `pulumi:"workerCount"`
}

// The set of arguments for constructing a HybridAksCluster resource.
type HybridAksClusterArgs struct {
	// The list of resource IDs for the workload networks associated with the Hybrid AKS cluster. It can be any of l2Networks, l3Networks, or trunkedNetworks resources. This field will also contain one cloudServicesNetwork and one defaultCniNetwork.
	AssociatedNetworkIds pulumi.StringArrayInput
	// The number of control plane node VMs.
	ControlPlaneCount pulumi.Float64Input
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// The name of the Hybrid AKS cluster.
	HybridAksClusterName pulumi.StringPtrInput
	// The resource ID of the Hybrid AKS cluster that this additional information is for.
	HybridAksProvisionedClusterId pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The number of worker node VMs.
	WorkerCount pulumi.Float64Input
}

func (HybridAksClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridAksClusterArgs)(nil)).Elem()
}

type HybridAksClusterInput interface {
	pulumi.Input

	ToHybridAksClusterOutput() HybridAksClusterOutput
	ToHybridAksClusterOutputWithContext(ctx context.Context) HybridAksClusterOutput
}

func (*HybridAksCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridAksCluster)(nil)).Elem()
}

func (i *HybridAksCluster) ToHybridAksClusterOutput() HybridAksClusterOutput {
	return i.ToHybridAksClusterOutputWithContext(context.Background())
}

func (i *HybridAksCluster) ToHybridAksClusterOutputWithContext(ctx context.Context) HybridAksClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridAksClusterOutput)
}

type HybridAksClusterOutput struct{ *pulumi.OutputState }

func (HybridAksClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridAksCluster)(nil)).Elem()
}

func (o HybridAksClusterOutput) ToHybridAksClusterOutput() HybridAksClusterOutput {
	return o
}

func (o HybridAksClusterOutput) ToHybridAksClusterOutputWithContext(ctx context.Context) HybridAksClusterOutput {
	return o
}

// The list of resource IDs for the workload networks associated with the Hybrid AKS cluster. It can be any of l2Networks, l3Networks, or trunkedNetworks resources. This field will also contain one cloudServicesNetwork and one defaultCniNetwork.
func (o HybridAksClusterOutput) AssociatedNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringArrayOutput { return v.AssociatedNetworkIds }).(pulumi.StringArrayOutput)
}

// The resource ID of the associated cloud services network.
func (o HybridAksClusterOutput) CloudServicesNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.CloudServicesNetworkId }).(pulumi.StringOutput)
}

// The resource ID of the Network Cloud cluster hosting the Hybrid AKS cluster.
func (o HybridAksClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The number of control plane node VMs.
func (o HybridAksClusterOutput) ControlPlaneCount() pulumi.Float64Output {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.Float64Output { return v.ControlPlaneCount }).(pulumi.Float64Output)
}

// The list of node configurations detailing associated VMs that are part of the control plane nodes of this Hybrid AKS cluster.
func (o HybridAksClusterOutput) ControlPlaneNodes() NodeConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *HybridAksCluster) NodeConfigurationResponseArrayOutput { return v.ControlPlaneNodes }).(NodeConfigurationResponseArrayOutput)
}

// The resource ID of the associated default CNI network.
func (o HybridAksClusterOutput) DefaultCniNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.DefaultCniNetworkId }).(pulumi.StringOutput)
}

// The more detailed status of this Hybrid AKS cluster.
func (o HybridAksClusterOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o HybridAksClusterOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o HybridAksClusterOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *HybridAksCluster) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The resource ID of the Hybrid AKS cluster that this additional information is for.
func (o HybridAksClusterOutput) HybridAksProvisionedClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.HybridAksProvisionedClusterId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o HybridAksClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o HybridAksClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the Hybrid AKS cluster resource.
func (o HybridAksClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o HybridAksClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridAksCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o HybridAksClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HybridAksClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The resource IDs of volumes that are attached to the Hybrid AKS cluster.
func (o HybridAksClusterOutput) Volumes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.StringArrayOutput { return v.Volumes }).(pulumi.StringArrayOutput)
}

// The number of worker node VMs.
func (o HybridAksClusterOutput) WorkerCount() pulumi.Float64Output {
	return o.ApplyT(func(v *HybridAksCluster) pulumi.Float64Output { return v.WorkerCount }).(pulumi.Float64Output)
}

// The list of node configurations detailing associated VMs that are part of the worker nodes of this Hybrid AKS cluster.
func (o HybridAksClusterOutput) WorkerNodes() NodeConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *HybridAksCluster) NodeConfigurationResponseArrayOutput { return v.WorkerNodes }).(NodeConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridAksClusterOutput{})
}

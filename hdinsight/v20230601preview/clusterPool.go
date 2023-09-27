// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Cluster pool.
type ClusterPool struct {
	pulumi.CustomResourceState

	// Properties of underlying AKS cluster.
	AksClusterProfile ClusterPoolResourcePropertiesResponseAksClusterProfileOutput `pulumi:"aksClusterProfile"`
	// A resource group created by AKS, to hold the infrastructure resources created by AKS on-behalf of customers. It is generated by cluster pool name and managed resource group name by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}
	AksManagedResourceGroupName pulumi.StringOutput `pulumi:"aksManagedResourceGroupName"`
	// CLuster pool profile.
	ClusterPoolProfile ClusterPoolResourcePropertiesResponseClusterPoolProfilePtrOutput `pulumi:"clusterPoolProfile"`
	// CLuster pool compute profile.
	ComputeProfile ClusterPoolResourcePropertiesResponseComputeProfileOutput `pulumi:"computeProfile"`
	// A unique id generated by the RP to identify the resource.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Cluster pool log analytics profile to enable OMS agent for AKS cluster.
	LogAnalyticsProfile ClusterPoolResourcePropertiesResponseLogAnalyticsProfilePtrOutput `pulumi:"logAnalyticsProfile"`
	// A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
	ManagedResourceGroupName pulumi.StringPtrOutput `pulumi:"managedResourceGroupName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Cluster pool network profile.
	NetworkProfile ClusterPoolResourcePropertiesResponseNetworkProfilePtrOutput `pulumi:"networkProfile"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Business status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewClusterPool registers a new resource with the given unique name, arguments, and options.
func NewClusterPool(ctx *pulumi.Context,
	name string, args *ClusterPoolArgs, opts ...pulumi.ResourceOption) (*ClusterPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeProfile == nil {
		return nil, errors.New("invalid value for required argument 'ComputeProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hdinsight:ClusterPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ClusterPool
	err := ctx.RegisterResource("azure-native:hdinsight/v20230601preview:ClusterPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterPool gets an existing ClusterPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterPoolState, opts ...pulumi.ResourceOption) (*ClusterPool, error) {
	var resource ClusterPool
	err := ctx.ReadResource("azure-native:hdinsight/v20230601preview:ClusterPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterPool resources.
type clusterPoolState struct {
}

type ClusterPoolState struct {
}

func (ClusterPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPoolState)(nil)).Elem()
}

type clusterPoolArgs struct {
	// The name of the cluster pool.
	ClusterPoolName *string `pulumi:"clusterPoolName"`
	// CLuster pool profile.
	ClusterPoolProfile *ClusterPoolResourcePropertiesClusterPoolProfile `pulumi:"clusterPoolProfile"`
	// CLuster pool compute profile.
	ComputeProfile ClusterPoolResourcePropertiesComputeProfile `pulumi:"computeProfile"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Cluster pool log analytics profile to enable OMS agent for AKS cluster.
	LogAnalyticsProfile *ClusterPoolResourcePropertiesLogAnalyticsProfile `pulumi:"logAnalyticsProfile"`
	// A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Cluster pool network profile.
	NetworkProfile *ClusterPoolResourcePropertiesNetworkProfile `pulumi:"networkProfile"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterPool resource.
type ClusterPoolArgs struct {
	// The name of the cluster pool.
	ClusterPoolName pulumi.StringPtrInput
	// CLuster pool profile.
	ClusterPoolProfile ClusterPoolResourcePropertiesClusterPoolProfilePtrInput
	// CLuster pool compute profile.
	ComputeProfile ClusterPoolResourcePropertiesComputeProfileInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Cluster pool log analytics profile to enable OMS agent for AKS cluster.
	LogAnalyticsProfile ClusterPoolResourcePropertiesLogAnalyticsProfilePtrInput
	// A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
	ManagedResourceGroupName pulumi.StringPtrInput
	// Cluster pool network profile.
	NetworkProfile ClusterPoolResourcePropertiesNetworkProfilePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ClusterPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPoolArgs)(nil)).Elem()
}

type ClusterPoolInput interface {
	pulumi.Input

	ToClusterPoolOutput() ClusterPoolOutput
	ToClusterPoolOutputWithContext(ctx context.Context) ClusterPoolOutput
}

func (*ClusterPool) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPool)(nil)).Elem()
}

func (i *ClusterPool) ToClusterPoolOutput() ClusterPoolOutput {
	return i.ToClusterPoolOutputWithContext(context.Background())
}

func (i *ClusterPool) ToClusterPoolOutputWithContext(ctx context.Context) ClusterPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPoolOutput)
}

func (i *ClusterPool) ToOutput(ctx context.Context) pulumix.Output[*ClusterPool] {
	return pulumix.Output[*ClusterPool]{
		OutputState: i.ToClusterPoolOutputWithContext(ctx).OutputState,
	}
}

type ClusterPoolOutput struct{ *pulumi.OutputState }

func (ClusterPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPool)(nil)).Elem()
}

func (o ClusterPoolOutput) ToClusterPoolOutput() ClusterPoolOutput {
	return o
}

func (o ClusterPoolOutput) ToClusterPoolOutputWithContext(ctx context.Context) ClusterPoolOutput {
	return o
}

func (o ClusterPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*ClusterPool] {
	return pulumix.Output[*ClusterPool]{
		OutputState: o.OutputState,
	}
}

// Properties of underlying AKS cluster.
func (o ClusterPoolOutput) AksClusterProfile() ClusterPoolResourcePropertiesResponseAksClusterProfileOutput {
	return o.ApplyT(func(v *ClusterPool) ClusterPoolResourcePropertiesResponseAksClusterProfileOutput {
		return v.AksClusterProfile
	}).(ClusterPoolResourcePropertiesResponseAksClusterProfileOutput)
}

// A resource group created by AKS, to hold the infrastructure resources created by AKS on-behalf of customers. It is generated by cluster pool name and managed resource group name by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}
func (o ClusterPoolOutput) AksManagedResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringOutput { return v.AksManagedResourceGroupName }).(pulumi.StringOutput)
}

// CLuster pool profile.
func (o ClusterPoolOutput) ClusterPoolProfile() ClusterPoolResourcePropertiesResponseClusterPoolProfilePtrOutput {
	return o.ApplyT(func(v *ClusterPool) ClusterPoolResourcePropertiesResponseClusterPoolProfilePtrOutput {
		return v.ClusterPoolProfile
	}).(ClusterPoolResourcePropertiesResponseClusterPoolProfilePtrOutput)
}

// CLuster pool compute profile.
func (o ClusterPoolOutput) ComputeProfile() ClusterPoolResourcePropertiesResponseComputeProfileOutput {
	return o.ApplyT(func(v *ClusterPool) ClusterPoolResourcePropertiesResponseComputeProfileOutput {
		return v.ComputeProfile
	}).(ClusterPoolResourcePropertiesResponseComputeProfileOutput)
}

// A unique id generated by the RP to identify the resource.
func (o ClusterPoolOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ClusterPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Cluster pool log analytics profile to enable OMS agent for AKS cluster.
func (o ClusterPoolOutput) LogAnalyticsProfile() ClusterPoolResourcePropertiesResponseLogAnalyticsProfilePtrOutput {
	return o.ApplyT(func(v *ClusterPool) ClusterPoolResourcePropertiesResponseLogAnalyticsProfilePtrOutput {
		return v.LogAnalyticsProfile
	}).(ClusterPoolResourcePropertiesResponseLogAnalyticsProfilePtrOutput)
}

// A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
func (o ClusterPoolOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringPtrOutput { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ClusterPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Cluster pool network profile.
func (o ClusterPoolOutput) NetworkProfile() ClusterPoolResourcePropertiesResponseNetworkProfilePtrOutput {
	return o.ApplyT(func(v *ClusterPool) ClusterPoolResourcePropertiesResponseNetworkProfilePtrOutput {
		return v.NetworkProfile
	}).(ClusterPoolResourcePropertiesResponseNetworkProfilePtrOutput)
}

// Provisioning state of the resource.
func (o ClusterPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Business status of the resource.
func (o ClusterPoolOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ClusterPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ClusterPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ClusterPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterPoolOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a cluster pool.
func LookupClusterPool(ctx *pulumi.Context, args *LookupClusterPoolArgs, opts ...pulumi.InvokeOption) (*LookupClusterPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterPoolResult
	err := ctx.Invoke("azure-native:hdinsight/v20230601preview:getClusterPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterPoolArgs struct {
	// The name of the cluster pool.
	ClusterPoolName string `pulumi:"clusterPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Cluster pool.
type LookupClusterPoolResult struct {
	// Properties of underlying AKS cluster.
	AksClusterProfile ClusterPoolResourcePropertiesResponseAksClusterProfile `pulumi:"aksClusterProfile"`
	// A resource group created by AKS, to hold the infrastructure resources created by AKS on-behalf of customers. It is generated by cluster pool name and managed resource group name by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}
	AksManagedResourceGroupName string `pulumi:"aksManagedResourceGroupName"`
	// CLuster pool profile.
	ClusterPoolProfile *ClusterPoolResourcePropertiesResponseClusterPoolProfile `pulumi:"clusterPoolProfile"`
	// CLuster pool compute profile.
	ComputeProfile ClusterPoolResourcePropertiesResponseComputeProfile `pulumi:"computeProfile"`
	// A unique id generated by the RP to identify the resource.
	DeploymentId string `pulumi:"deploymentId"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Cluster pool log analytics profile to enable OMS agent for AKS cluster.
	LogAnalyticsProfile *ClusterPoolResourcePropertiesResponseLogAnalyticsProfile `pulumi:"logAnalyticsProfile"`
	// A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Cluster pool network profile.
	NetworkProfile *ClusterPoolResourcePropertiesResponseNetworkProfile `pulumi:"networkProfile"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Business status of the resource.
	Status string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupClusterPoolOutput(ctx *pulumi.Context, args LookupClusterPoolOutputArgs, opts ...pulumi.InvokeOption) LookupClusterPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterPoolResult, error) {
			args := v.(LookupClusterPoolArgs)
			r, err := LookupClusterPool(ctx, &args, opts...)
			var s LookupClusterPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterPoolResultOutput)
}

type LookupClusterPoolOutputArgs struct {
	// The name of the cluster pool.
	ClusterPoolName pulumi.StringInput `pulumi:"clusterPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterPoolArgs)(nil)).Elem()
}

// Cluster pool.
type LookupClusterPoolResultOutput struct{ *pulumi.OutputState }

func (LookupClusterPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterPoolResult)(nil)).Elem()
}

func (o LookupClusterPoolResultOutput) ToLookupClusterPoolResultOutput() LookupClusterPoolResultOutput {
	return o
}

func (o LookupClusterPoolResultOutput) ToLookupClusterPoolResultOutputWithContext(ctx context.Context) LookupClusterPoolResultOutput {
	return o
}

func (o LookupClusterPoolResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupClusterPoolResult] {
	return pulumix.Output[LookupClusterPoolResult]{
		OutputState: o.OutputState,
	}
}

// Properties of underlying AKS cluster.
func (o LookupClusterPoolResultOutput) AksClusterProfile() ClusterPoolResourcePropertiesResponseAksClusterProfileOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) ClusterPoolResourcePropertiesResponseAksClusterProfile {
		return v.AksClusterProfile
	}).(ClusterPoolResourcePropertiesResponseAksClusterProfileOutput)
}

// A resource group created by AKS, to hold the infrastructure resources created by AKS on-behalf of customers. It is generated by cluster pool name and managed resource group name by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}
func (o LookupClusterPoolResultOutput) AksManagedResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) string { return v.AksManagedResourceGroupName }).(pulumi.StringOutput)
}

// CLuster pool profile.
func (o LookupClusterPoolResultOutput) ClusterPoolProfile() ClusterPoolResourcePropertiesResponseClusterPoolProfilePtrOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) *ClusterPoolResourcePropertiesResponseClusterPoolProfile {
		return v.ClusterPoolProfile
	}).(ClusterPoolResourcePropertiesResponseClusterPoolProfilePtrOutput)
}

// CLuster pool compute profile.
func (o LookupClusterPoolResultOutput) ComputeProfile() ClusterPoolResourcePropertiesResponseComputeProfileOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) ClusterPoolResourcePropertiesResponseComputeProfile {
		return v.ComputeProfile
	}).(ClusterPoolResourcePropertiesResponseComputeProfileOutput)
}

// A unique id generated by the RP to identify the resource.
func (o LookupClusterPoolResultOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) string { return v.DeploymentId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupClusterPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupClusterPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// Cluster pool log analytics profile to enable OMS agent for AKS cluster.
func (o LookupClusterPoolResultOutput) LogAnalyticsProfile() ClusterPoolResourcePropertiesResponseLogAnalyticsProfilePtrOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) *ClusterPoolResourcePropertiesResponseLogAnalyticsProfile {
		return v.LogAnalyticsProfile
	}).(ClusterPoolResourcePropertiesResponseLogAnalyticsProfilePtrOutput)
}

// A resource group created by RP, to hold the resources created by RP on-behalf of customers. It will also be used to generate aksManagedResourceGroupName by pattern: MC_{managedResourceGroupName}_{clusterPoolName}_{region}. Please make sure it meets resource group name restriction.
func (o LookupClusterPoolResultOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) *string { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupClusterPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Cluster pool network profile.
func (o LookupClusterPoolResultOutput) NetworkProfile() ClusterPoolResourcePropertiesResponseNetworkProfilePtrOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) *ClusterPoolResourcePropertiesResponseNetworkProfile {
		return v.NetworkProfile
	}).(ClusterPoolResourcePropertiesResponseNetworkProfilePtrOutput)
}

// Provisioning state of the resource.
func (o LookupClusterPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Business status of the resource.
func (o LookupClusterPoolResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) string { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupClusterPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupClusterPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterPoolResultOutput{})
}

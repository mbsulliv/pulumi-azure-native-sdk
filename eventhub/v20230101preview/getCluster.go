// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the resource description of the specified Event Hubs Cluster.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:eventhub/v20230101preview:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// The name of the Event Hubs Cluster.
	ClusterName string `pulumi:"clusterName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Single Event Hubs Cluster resource in List or Get operations.
type LookupClusterResult struct {
	// The UTC time when the Event Hubs Cluster was created.
	CreatedAt string `pulumi:"createdAt"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
	MetricId string `pulumi:"metricId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the Cluster.
	ProvisioningState string `pulumi:"provisioningState"`
	// Properties of the cluster SKU.
	Sku *ClusterSkuResponse `pulumi:"sku"`
	// Status of the Cluster resource
	Status string `pulumi:"status"`
	// A value that indicates whether Scaling is Supported.
	SupportsScaling *bool `pulumi:"supportsScaling"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The UTC time when the Event Hubs Cluster was last updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	// The name of the Event Hubs Cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// Single Event Hubs Cluster resource in List or Get operations.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupClusterResult] {
	return pulumix.Output[LookupClusterResult]{
		OutputState: o.OutputState,
	}
}

// The UTC time when the Event Hubs Cluster was created.
func (o LookupClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupClusterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
func (o LookupClusterResultOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.MetricId }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Cluster.
func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Properties of the cluster SKU.
func (o LookupClusterResultOutput) Sku() ClusterSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterSkuResponse { return v.Sku }).(ClusterSkuResponsePtrOutput)
}

// Status of the Cluster resource
func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// A value that indicates whether Scaling is Supported.
func (o LookupClusterResultOutput) SupportsScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.SupportsScaling }).(pulumi.BoolPtrOutput)
}

// The system meta data relating to this resource.
func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// The UTC time when the Event Hubs Cluster was last updated.
func (o LookupClusterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}

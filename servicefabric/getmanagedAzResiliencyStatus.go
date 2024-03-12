// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Action to get Az Resiliency Status of all the Base resources constituting Service Fabric Managed Clusters.
// Azure REST API version: 2023-03-01-preview.
//
// Other available API versions: 2022-10-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview.
func GetmanagedAzResiliencyStatus(ctx *pulumi.Context, args *GetmanagedAzResiliencyStatusArgs, opts ...pulumi.InvokeOption) (*GetmanagedAzResiliencyStatusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetmanagedAzResiliencyStatusResult
	err := ctx.Invoke("azure-native:servicefabric:getmanagedAzResiliencyStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetmanagedAzResiliencyStatusArgs struct {
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes the result of the request to list Managed VM Sizes for Service Fabric Managed Clusters.
type GetmanagedAzResiliencyStatusResult struct {
	// List of Managed VM Sizes for Service Fabric Managed Clusters.
	BaseResourceStatus []ResourceAzStatusResponse `pulumi:"baseResourceStatus"`
	// URL to get the next set of Managed VM Sizes if there are any.
	IsClusterZoneResilient bool `pulumi:"isClusterZoneResilient"`
}

func GetmanagedAzResiliencyStatusOutput(ctx *pulumi.Context, args GetmanagedAzResiliencyStatusOutputArgs, opts ...pulumi.InvokeOption) GetmanagedAzResiliencyStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetmanagedAzResiliencyStatusResult, error) {
			args := v.(GetmanagedAzResiliencyStatusArgs)
			r, err := GetmanagedAzResiliencyStatus(ctx, &args, opts...)
			var s GetmanagedAzResiliencyStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetmanagedAzResiliencyStatusResultOutput)
}

type GetmanagedAzResiliencyStatusOutputArgs struct {
	// The name of the cluster resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetmanagedAzResiliencyStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmanagedAzResiliencyStatusArgs)(nil)).Elem()
}

// Describes the result of the request to list Managed VM Sizes for Service Fabric Managed Clusters.
type GetmanagedAzResiliencyStatusResultOutput struct{ *pulumi.OutputState }

func (GetmanagedAzResiliencyStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmanagedAzResiliencyStatusResult)(nil)).Elem()
}

func (o GetmanagedAzResiliencyStatusResultOutput) ToGetmanagedAzResiliencyStatusResultOutput() GetmanagedAzResiliencyStatusResultOutput {
	return o
}

func (o GetmanagedAzResiliencyStatusResultOutput) ToGetmanagedAzResiliencyStatusResultOutputWithContext(ctx context.Context) GetmanagedAzResiliencyStatusResultOutput {
	return o
}

// List of Managed VM Sizes for Service Fabric Managed Clusters.
func (o GetmanagedAzResiliencyStatusResultOutput) BaseResourceStatus() ResourceAzStatusResponseArrayOutput {
	return o.ApplyT(func(v GetmanagedAzResiliencyStatusResult) []ResourceAzStatusResponse { return v.BaseResourceStatus }).(ResourceAzStatusResponseArrayOutput)
}

// URL to get the next set of Managed VM Sizes if there are any.
func (o GetmanagedAzResiliencyStatusResultOutput) IsClusterZoneResilient() pulumi.BoolOutput {
	return o.ApplyT(func(v GetmanagedAzResiliencyStatusResult) bool { return v.IsClusterZoneResilient }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetmanagedAzResiliencyStatusResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A cluster resource
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:avs/v20230301:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// Name of the cluster in the private cloud
	ClusterName string `pulumi:"clusterName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A cluster resource
type LookupClusterResult struct {
	// The identity
	ClusterId int `pulumi:"clusterId"`
	// The cluster size
	ClusterSize *int `pulumi:"clusterSize"`
	// The hosts
	Hosts []string `pulumi:"hosts"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The state of the cluster provisioning
	ProvisioningState string `pulumi:"provisioningState"`
	// The cluster SKU
	Sku SkuResponse `pulumi:"sku"`
	// Resource type.
	Type string `pulumi:"type"`
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
	// Name of the cluster in the private cloud
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A cluster resource
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

// The identity
func (o LookupClusterResultOutput) ClusterId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.ClusterId }).(pulumi.IntOutput)
}

// The cluster size
func (o LookupClusterResultOutput) ClusterSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.ClusterSize }).(pulumi.IntPtrOutput)
}

// The hosts
func (o LookupClusterResultOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the cluster provisioning
func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The cluster SKU
func (o LookupClusterResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Resource type.
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of all zones and associated hosts for a cluster
// Azure REST API version: 2022-05-01.
//
// Other available API versions: 2023-03-01.
func ListClusterZones(ctx *pulumi.Context, args *ListClusterZonesArgs, opts ...pulumi.InvokeOption) (*ListClusterZonesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListClusterZonesResult
	err := ctx.Invoke("azure-native:avs:listClusterZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterZonesArgs struct {
	// Name of the cluster in the private cloud
	ClusterName string `pulumi:"clusterName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of all zones and associated hosts for a cluster
type ListClusterZonesResult struct {
	// Zone and associated hosts info
	Zones []ClusterZoneResponse `pulumi:"zones"`
}

func ListClusterZonesOutput(ctx *pulumi.Context, args ListClusterZonesOutputArgs, opts ...pulumi.InvokeOption) ListClusterZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListClusterZonesResult, error) {
			args := v.(ListClusterZonesArgs)
			r, err := ListClusterZones(ctx, &args, opts...)
			var s ListClusterZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListClusterZonesResultOutput)
}

type ListClusterZonesOutputArgs struct {
	// Name of the cluster in the private cloud
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListClusterZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterZonesArgs)(nil)).Elem()
}

// List of all zones and associated hosts for a cluster
type ListClusterZonesResultOutput struct{ *pulumi.OutputState }

func (ListClusterZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterZonesResult)(nil)).Elem()
}

func (o ListClusterZonesResultOutput) ToListClusterZonesResultOutput() ListClusterZonesResultOutput {
	return o
}

func (o ListClusterZonesResultOutput) ToListClusterZonesResultOutputWithContext(ctx context.Context) ListClusterZonesResultOutput {
	return o
}

// Zone and associated hosts info
func (o ListClusterZonesResultOutput) Zones() ClusterZoneResponseArrayOutput {
	return o.ApplyT(func(v ListClusterZonesResult) []ClusterZoneResponse { return v.Zones }).(ClusterZoneResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterZonesResultOutput{})
}

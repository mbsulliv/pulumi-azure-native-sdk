// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A vSphere Distributed Resource Scheduler (DRS) placement policy
// Azure REST API version: 2022-05-01.
//
// Other available API versions: 2023-03-01, 2023-09-01.
func LookupPlacementPolicy(ctx *pulumi.Context, args *LookupPlacementPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPlacementPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPlacementPolicyResult
	err := ctx.Invoke("azure-native:avs:getPlacementPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlacementPolicyArgs struct {
	// Name of the cluster in the private cloud
	ClusterName string `pulumi:"clusterName"`
	// Name of the VMware vSphere Distributed Resource Scheduler (DRS) placement policy
	PlacementPolicyName string `pulumi:"placementPolicyName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A vSphere Distributed Resource Scheduler (DRS) placement policy
type LookupPlacementPolicyResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// placement policy properties
	Properties interface{} `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupPlacementPolicyOutput(ctx *pulumi.Context, args LookupPlacementPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPlacementPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlacementPolicyResult, error) {
			args := v.(LookupPlacementPolicyArgs)
			r, err := LookupPlacementPolicy(ctx, &args, opts...)
			var s LookupPlacementPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPlacementPolicyResultOutput)
}

type LookupPlacementPolicyOutputArgs struct {
	// Name of the cluster in the private cloud
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Name of the VMware vSphere Distributed Resource Scheduler (DRS) placement policy
	PlacementPolicyName pulumi.StringInput `pulumi:"placementPolicyName"`
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPlacementPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementPolicyArgs)(nil)).Elem()
}

// A vSphere Distributed Resource Scheduler (DRS) placement policy
type LookupPlacementPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPlacementPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlacementPolicyResult)(nil)).Elem()
}

func (o LookupPlacementPolicyResultOutput) ToLookupPlacementPolicyResultOutput() LookupPlacementPolicyResultOutput {
	return o
}

func (o LookupPlacementPolicyResultOutput) ToLookupPlacementPolicyResultOutputWithContext(ctx context.Context) LookupPlacementPolicyResultOutput {
	return o
}

// Resource ID.
func (o LookupPlacementPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupPlacementPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// placement policy properties
func (o LookupPlacementPolicyResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPlacementPolicyResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Resource type.
func (o LookupPlacementPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlacementPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlacementPolicyResultOutput{})
}

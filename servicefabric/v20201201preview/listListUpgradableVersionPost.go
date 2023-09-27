// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// If a target is not provided, it will get the minimum and maximum versions available from the current cluster version. If a target is given, it will provide the required path to get from the current cluster version to the target version.
func ListListUpgradableVersionPost(ctx *pulumi.Context, args *ListListUpgradableVersionPostArgs, opts ...pulumi.InvokeOption) (*ListListUpgradableVersionPostResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListListUpgradableVersionPostResult
	err := ctx.Invoke("azure-native:servicefabric/v20201201preview:listListUpgradableVersionPost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListUpgradableVersionPostArgs struct {
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The target code version.
	TargetVersion string `pulumi:"targetVersion"`
}

// The list of intermediate cluster code versions for an upgrade or downgrade. Or minimum and maximum upgradable version if no target was given
type ListListUpgradableVersionPostResult struct {
	SupportedPath []string `pulumi:"supportedPath"`
}

func ListListUpgradableVersionPostOutput(ctx *pulumi.Context, args ListListUpgradableVersionPostOutputArgs, opts ...pulumi.InvokeOption) ListListUpgradableVersionPostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListListUpgradableVersionPostResult, error) {
			args := v.(ListListUpgradableVersionPostArgs)
			r, err := ListListUpgradableVersionPost(ctx, &args, opts...)
			var s ListListUpgradableVersionPostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListListUpgradableVersionPostResultOutput)
}

type ListListUpgradableVersionPostOutputArgs struct {
	// The name of the cluster resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The target code version.
	TargetVersion pulumi.StringInput `pulumi:"targetVersion"`
}

func (ListListUpgradableVersionPostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListUpgradableVersionPostArgs)(nil)).Elem()
}

// The list of intermediate cluster code versions for an upgrade or downgrade. Or minimum and maximum upgradable version if no target was given
type ListListUpgradableVersionPostResultOutput struct{ *pulumi.OutputState }

func (ListListUpgradableVersionPostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListUpgradableVersionPostResult)(nil)).Elem()
}

func (o ListListUpgradableVersionPostResultOutput) ToListListUpgradableVersionPostResultOutput() ListListUpgradableVersionPostResultOutput {
	return o
}

func (o ListListUpgradableVersionPostResultOutput) ToListListUpgradableVersionPostResultOutputWithContext(ctx context.Context) ListListUpgradableVersionPostResultOutput {
	return o
}

func (o ListListUpgradableVersionPostResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListListUpgradableVersionPostResult] {
	return pulumix.Output[ListListUpgradableVersionPostResult]{
		OutputState: o.OutputState,
	}
}

func (o ListListUpgradableVersionPostResultOutput) SupportedPath() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListListUpgradableVersionPostResult) []string { return v.SupportedPath }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListUpgradableVersionPostResultOutput{})
}

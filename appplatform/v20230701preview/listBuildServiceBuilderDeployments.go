// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List deployments that are using the builder.
func ListBuildServiceBuilderDeployments(ctx *pulumi.Context, args *ListBuildServiceBuilderDeploymentsArgs, opts ...pulumi.InvokeOption) (*ListBuildServiceBuilderDeploymentsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListBuildServiceBuilderDeploymentsResult
	err := ctx.Invoke("azure-native:appplatform/v20230701preview:listBuildServiceBuilderDeployments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBuildServiceBuilderDeploymentsArgs struct {
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// The name of the builder resource.
	BuilderName string `pulumi:"builderName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// A list of deployments resource ids.
type ListBuildServiceBuilderDeploymentsResult struct {
	// A list of deployment resource ids.
	Deployments []string `pulumi:"deployments"`
}

func ListBuildServiceBuilderDeploymentsOutput(ctx *pulumi.Context, args ListBuildServiceBuilderDeploymentsOutputArgs, opts ...pulumi.InvokeOption) ListBuildServiceBuilderDeploymentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBuildServiceBuilderDeploymentsResult, error) {
			args := v.(ListBuildServiceBuilderDeploymentsArgs)
			r, err := ListBuildServiceBuilderDeployments(ctx, &args, opts...)
			var s ListBuildServiceBuilderDeploymentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBuildServiceBuilderDeploymentsResultOutput)
}

type ListBuildServiceBuilderDeploymentsOutputArgs struct {
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput `pulumi:"buildServiceName"`
	// The name of the builder resource.
	BuilderName pulumi.StringInput `pulumi:"builderName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ListBuildServiceBuilderDeploymentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildServiceBuilderDeploymentsArgs)(nil)).Elem()
}

// A list of deployments resource ids.
type ListBuildServiceBuilderDeploymentsResultOutput struct{ *pulumi.OutputState }

func (ListBuildServiceBuilderDeploymentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildServiceBuilderDeploymentsResult)(nil)).Elem()
}

func (o ListBuildServiceBuilderDeploymentsResultOutput) ToListBuildServiceBuilderDeploymentsResultOutput() ListBuildServiceBuilderDeploymentsResultOutput {
	return o
}

func (o ListBuildServiceBuilderDeploymentsResultOutput) ToListBuildServiceBuilderDeploymentsResultOutputWithContext(ctx context.Context) ListBuildServiceBuilderDeploymentsResultOutput {
	return o
}

// A list of deployment resource ids.
func (o ListBuildServiceBuilderDeploymentsResultOutput) Deployments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListBuildServiceBuilderDeploymentsResult) []string { return v.Deployments }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBuildServiceBuilderDeploymentsResultOutput{})
}

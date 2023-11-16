// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The properties of deployment in Elastic cloud corresponding to the Elastic monitor resource.
func ListDeploymentInfo(ctx *pulumi.Context, args *ListDeploymentInfoArgs, opts ...pulumi.InvokeOption) (*ListDeploymentInfoResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDeploymentInfoResult
	err := ctx.Invoke("azure-native:elastic/v20230701preview:listDeploymentInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeploymentInfoArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of deployment in Elastic cloud corresponding to the Elastic monitor resource.
type ListDeploymentInfoResult struct {
	// Deployment URL of the elasticsearch in Elastic cloud deployment.
	DeploymentUrl string `pulumi:"deploymentUrl"`
	// Disk capacity of the elasticsearch in Elastic cloud deployment.
	DiskCapacity string `pulumi:"diskCapacity"`
	// Marketplace SaaS Info of the resource.
	MarketplaceSaasInfo MarketplaceSaaSInfoResponse `pulumi:"marketplaceSaasInfo"`
	// RAM capacity of the elasticsearch in Elastic cloud deployment.
	MemoryCapacity string `pulumi:"memoryCapacity"`
	// The Elastic deployment status.
	Status string `pulumi:"status"`
	// Version of the elasticsearch in Elastic cloud deployment.
	Version string `pulumi:"version"`
}

func ListDeploymentInfoOutput(ctx *pulumi.Context, args ListDeploymentInfoOutputArgs, opts ...pulumi.InvokeOption) ListDeploymentInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDeploymentInfoResult, error) {
			args := v.(ListDeploymentInfoArgs)
			r, err := ListDeploymentInfo(ctx, &args, opts...)
			var s ListDeploymentInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDeploymentInfoResultOutput)
}

type ListDeploymentInfoOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDeploymentInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeploymentInfoArgs)(nil)).Elem()
}

// The properties of deployment in Elastic cloud corresponding to the Elastic monitor resource.
type ListDeploymentInfoResultOutput struct{ *pulumi.OutputState }

func (ListDeploymentInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeploymentInfoResult)(nil)).Elem()
}

func (o ListDeploymentInfoResultOutput) ToListDeploymentInfoResultOutput() ListDeploymentInfoResultOutput {
	return o
}

func (o ListDeploymentInfoResultOutput) ToListDeploymentInfoResultOutputWithContext(ctx context.Context) ListDeploymentInfoResultOutput {
	return o
}

// Deployment URL of the elasticsearch in Elastic cloud deployment.
func (o ListDeploymentInfoResultOutput) DeploymentUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.DeploymentUrl }).(pulumi.StringOutput)
}

// Disk capacity of the elasticsearch in Elastic cloud deployment.
func (o ListDeploymentInfoResultOutput) DiskCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.DiskCapacity }).(pulumi.StringOutput)
}

// Marketplace SaaS Info of the resource.
func (o ListDeploymentInfoResultOutput) MarketplaceSaasInfo() MarketplaceSaaSInfoResponseOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) MarketplaceSaaSInfoResponse { return v.MarketplaceSaasInfo }).(MarketplaceSaaSInfoResponseOutput)
}

// RAM capacity of the elasticsearch in Elastic cloud deployment.
func (o ListDeploymentInfoResultOutput) MemoryCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.MemoryCapacity }).(pulumi.StringOutput)
}

// The Elastic deployment status.
func (o ListDeploymentInfoResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.Status }).(pulumi.StringOutput)
}

// Version of the elasticsearch in Elastic cloud deployment.
func (o ListDeploymentInfoResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ListDeploymentInfoResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDeploymentInfoResultOutput{})
}

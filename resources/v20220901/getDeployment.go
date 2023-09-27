// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a deployment.
func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDeploymentResult
	err := ctx.Invoke("azure-native:resources/v20220901:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentArgs struct {
	// The name of the deployment.
	DeploymentName string `pulumi:"deploymentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Deployment information.
type LookupDeploymentResult struct {
	// The ID of the deployment.
	Id string `pulumi:"id"`
	// the location of the deployment.
	Location *string `pulumi:"location"`
	// The name of the deployment.
	Name string `pulumi:"name"`
	// Deployment properties.
	Properties DeploymentPropertiesExtendedResponse `pulumi:"properties"`
	// Deployment tags
	Tags map[string]string `pulumi:"tags"`
	// The type of the deployment.
	Type string `pulumi:"type"`
}

func LookupDeploymentOutput(ctx *pulumi.Context, args LookupDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentResult, error) {
			args := v.(LookupDeploymentArgs)
			r, err := LookupDeployment(ctx, &args, opts...)
			var s LookupDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentResultOutput)
}

type LookupDeploymentOutputArgs struct {
	// The name of the deployment.
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentArgs)(nil)).Elem()
}

// Deployment information.
type LookupDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentResult)(nil)).Elem()
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutput() LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) ToLookupDeploymentResultOutputWithContext(ctx context.Context) LookupDeploymentResultOutput {
	return o
}

func (o LookupDeploymentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDeploymentResult] {
	return pulumix.Output[LookupDeploymentResult]{
		OutputState: o.OutputState,
	}
}

// The ID of the deployment.
func (o LookupDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

// the location of the deployment.
func (o LookupDeploymentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the deployment.
func (o LookupDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Deployment properties.
func (o LookupDeploymentResultOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v LookupDeploymentResult) DeploymentPropertiesExtendedResponse { return v.Properties }).(DeploymentPropertiesExtendedResponseOutput)
}

// Deployment tags
func (o LookupDeploymentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the deployment.
func (o LookupDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentResultOutput{})
}

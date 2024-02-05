// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Deployment and its properties.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-01-01-preview.
func LookupDeployment(ctx *pulumi.Context, args *LookupDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDeploymentResult
	err := ctx.Invoke("azure-native:appplatform:getDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDeploymentArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the Deployment resource.
	DeploymentName string `pulumi:"deploymentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Deployment resource payload
type LookupDeploymentResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the Deployment resource
	Properties DeploymentResourcePropertiesResponse `pulumi:"properties"`
	// Sku of the Deployment resource
	Sku *SkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDeploymentResult
func (val *LookupDeploymentResult) Defaults() *LookupDeploymentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
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
	// The name of the App resource.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The name of the Deployment resource.
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentArgs)(nil)).Elem()
}

// Deployment resource payload
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

// Fully qualified resource Id for the resource.
func (o LookupDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the Deployment resource
func (o LookupDeploymentResultOutput) Properties() DeploymentResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupDeploymentResult) DeploymentResourcePropertiesResponse { return v.Properties }).(DeploymentResourcePropertiesResponseOutput)
}

// Sku of the Deployment resource
func (o LookupDeploymentResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDeploymentResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDeploymentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeploymentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentResultOutput{})
}

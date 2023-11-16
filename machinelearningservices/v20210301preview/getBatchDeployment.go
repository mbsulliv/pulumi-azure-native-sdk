// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBatchDeployment(ctx *pulumi.Context, args *LookupBatchDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupBatchDeploymentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBatchDeploymentResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getBatchDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchDeploymentArgs struct {
	// The identifier for the Batch deployments.
	DeploymentName string `pulumi:"deploymentName"`
	// Endpoint name
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupBatchDeploymentResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Service identity associated with a resource.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties BatchDeploymentResponse `pulumi:"properties"`
	// System data associated with resource provider
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBatchDeploymentOutput(ctx *pulumi.Context, args LookupBatchDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupBatchDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBatchDeploymentResult, error) {
			args := v.(LookupBatchDeploymentArgs)
			r, err := LookupBatchDeployment(ctx, &args, opts...)
			var s LookupBatchDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBatchDeploymentResultOutput)
}

type LookupBatchDeploymentOutputArgs struct {
	// The identifier for the Batch deployments.
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	// Endpoint name
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBatchDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchDeploymentArgs)(nil)).Elem()
}

type LookupBatchDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupBatchDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchDeploymentResult)(nil)).Elem()
}

func (o LookupBatchDeploymentResultOutput) ToLookupBatchDeploymentResultOutput() LookupBatchDeploymentResultOutput {
	return o
}

func (o LookupBatchDeploymentResultOutput) ToLookupBatchDeploymentResultOutputWithContext(ctx context.Context) LookupBatchDeploymentResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBatchDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Service identity associated with a resource.
func (o LookupBatchDeploymentResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o LookupBatchDeploymentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupBatchDeploymentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBatchDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o LookupBatchDeploymentResultOutput) Properties() BatchDeploymentResponseOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) BatchDeploymentResponse { return v.Properties }).(BatchDeploymentResponseOutput)
}

// System data associated with resource provider
func (o LookupBatchDeploymentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupBatchDeploymentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBatchDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBatchDeploymentResultOutput{})
}

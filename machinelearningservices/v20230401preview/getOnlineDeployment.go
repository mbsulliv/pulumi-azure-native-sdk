// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupOnlineDeployment(ctx *pulumi.Context, args *LookupOnlineDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupOnlineDeploymentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOnlineDeploymentResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getOnlineDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOnlineDeploymentArgs struct {
	// Inference Endpoint Deployment name.
	DeploymentName string `pulumi:"deploymentName"`
	// Inference endpoint name.
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupOnlineDeploymentResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	OnlineDeploymentProperties interface{} `pulumi:"onlineDeploymentProperties"`
	// Sku details required for ARM contract for Autoscaling.
	Sku *SkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupOnlineDeploymentOutput(ctx *pulumi.Context, args LookupOnlineDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupOnlineDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOnlineDeploymentResult, error) {
			args := v.(LookupOnlineDeploymentArgs)
			r, err := LookupOnlineDeployment(ctx, &args, opts...)
			var s LookupOnlineDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOnlineDeploymentResultOutput)
}

type LookupOnlineDeploymentOutputArgs struct {
	// Inference Endpoint Deployment name.
	DeploymentName pulumi.StringInput `pulumi:"deploymentName"`
	// Inference endpoint name.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOnlineDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineDeploymentArgs)(nil)).Elem()
}

type LookupOnlineDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupOnlineDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineDeploymentResult)(nil)).Elem()
}

func (o LookupOnlineDeploymentResultOutput) ToLookupOnlineDeploymentResultOutput() LookupOnlineDeploymentResultOutput {
	return o
}

func (o LookupOnlineDeploymentResultOutput) ToLookupOnlineDeploymentResultOutputWithContext(ctx context.Context) LookupOnlineDeploymentResultOutput {
	return o
}

func (o LookupOnlineDeploymentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOnlineDeploymentResult] {
	return pulumix.Output[LookupOnlineDeploymentResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOnlineDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o LookupOnlineDeploymentResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o LookupOnlineDeploymentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupOnlineDeploymentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOnlineDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o LookupOnlineDeploymentResultOutput) OnlineDeploymentProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) interface{} { return v.OnlineDeploymentProperties }).(pulumi.AnyOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o LookupOnlineDeploymentResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupOnlineDeploymentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupOnlineDeploymentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOnlineDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOnlineDeploymentResultOutput{})
}

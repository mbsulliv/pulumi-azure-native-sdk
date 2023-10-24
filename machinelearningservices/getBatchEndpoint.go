// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01.
func LookupBatchEndpoint(ctx *pulumi.Context, args *LookupBatchEndpointArgs, opts ...pulumi.InvokeOption) (*LookupBatchEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBatchEndpointResult
	err := ctx.Invoke("azure-native:machinelearningservices:getBatchEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchEndpointArgs struct {
	// Name for the Batch Endpoint.
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupBatchEndpointResult struct {
	// [Required] Additional attributes of the entity.
	BatchEndpointProperties BatchEndpointResponse `pulumi:"batchEndpointProperties"`
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
	// Sku details required for ARM contract for Autoscaling.
	Sku *SkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBatchEndpointOutput(ctx *pulumi.Context, args LookupBatchEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupBatchEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBatchEndpointResult, error) {
			args := v.(LookupBatchEndpointArgs)
			r, err := LookupBatchEndpoint(ctx, &args, opts...)
			var s LookupBatchEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBatchEndpointResultOutput)
}

type LookupBatchEndpointOutputArgs struct {
	// Name for the Batch Endpoint.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBatchEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchEndpointArgs)(nil)).Elem()
}

type LookupBatchEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupBatchEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchEndpointResult)(nil)).Elem()
}

func (o LookupBatchEndpointResultOutput) ToLookupBatchEndpointResultOutput() LookupBatchEndpointResultOutput {
	return o
}

func (o LookupBatchEndpointResultOutput) ToLookupBatchEndpointResultOutputWithContext(ctx context.Context) LookupBatchEndpointResultOutput {
	return o
}

func (o LookupBatchEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBatchEndpointResult] {
	return pulumix.Output[LookupBatchEndpointResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupBatchEndpointResultOutput) BatchEndpointProperties() BatchEndpointResponseOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) BatchEndpointResponse { return v.BatchEndpointProperties }).(BatchEndpointResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBatchEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o LookupBatchEndpointResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o LookupBatchEndpointResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupBatchEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBatchEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o LookupBatchEndpointResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBatchEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupBatchEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBatchEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBatchEndpointResultOutput{})
}

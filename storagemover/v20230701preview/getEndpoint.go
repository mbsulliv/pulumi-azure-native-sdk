// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an Endpoint resource.
func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:storagemover/v20230701preview:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	// The name of the Endpoint resource.
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName string `pulumi:"storageMoverName"`
}

// The Endpoint resource, which contains information about file sources and targets.
type LookupEndpointResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource specific properties for the Storage Mover resource.
	Properties interface{} `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

type LookupEndpointOutputArgs struct {
	// The name of the Endpoint resource.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName pulumi.StringInput `pulumi:"storageMoverName"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}

// The Endpoint resource, which contains information about file sources and targets.
type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource specific properties for the Storage Mover resource.
func (o LookupEndpointResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupEndpointResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}

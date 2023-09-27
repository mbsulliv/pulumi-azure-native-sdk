// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the Cache specified by its identifier.
// Azure REST API version: 2022-08-01.
func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCacheResult
	err := ctx.Invoke("azure-native:apimanagement:getCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCacheArgs struct {
	// Identifier of the Cache entity. Cache identifier (should be either 'default' or valid Azure region identifier).
	CacheId string `pulumi:"cacheId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Cache details.
type LookupCacheResult struct {
	// Runtime connection string to cache
	ConnectionString string `pulumi:"connectionString"`
	// Cache description
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Original uri of entity in external system cache points to
	ResourceId *string `pulumi:"resourceId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Location identifier to use cache from (should be either 'default' or valid Azure region identifier)
	UseFromLocation string `pulumi:"useFromLocation"`
}

func LookupCacheOutput(ctx *pulumi.Context, args LookupCacheOutputArgs, opts ...pulumi.InvokeOption) LookupCacheResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCacheResult, error) {
			args := v.(LookupCacheArgs)
			r, err := LookupCache(ctx, &args, opts...)
			var s LookupCacheResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCacheResultOutput)
}

type LookupCacheOutputArgs struct {
	// Identifier of the Cache entity. Cache identifier (should be either 'default' or valid Azure region identifier).
	CacheId pulumi.StringInput `pulumi:"cacheId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCacheOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheArgs)(nil)).Elem()
}

// Cache details.
type LookupCacheResultOutput struct{ *pulumi.OutputState }

func (LookupCacheResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheResult)(nil)).Elem()
}

func (o LookupCacheResultOutput) ToLookupCacheResultOutput() LookupCacheResultOutput {
	return o
}

func (o LookupCacheResultOutput) ToLookupCacheResultOutputWithContext(ctx context.Context) LookupCacheResultOutput {
	return o
}

func (o LookupCacheResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCacheResult] {
	return pulumix.Output[LookupCacheResult]{
		OutputState: o.OutputState,
	}
}

// Runtime connection string to cache
func (o LookupCacheResultOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.ConnectionString }).(pulumi.StringOutput)
}

// Cache description
func (o LookupCacheResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCacheResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCacheResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Name }).(pulumi.StringOutput)
}

// Original uri of entity in external system cache points to
func (o LookupCacheResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCacheResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Type }).(pulumi.StringOutput)
}

// Location identifier to use cache from (should be either 'default' or valid Azure region identifier)
func (o LookupCacheResultOutput) UseFromLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.UseFromLocation }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCacheResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the developer portal's content type. Content types describe content items' properties, validation rules, and constraints.
func LookupContentType(ctx *pulumi.Context, args *LookupContentTypeArgs, opts ...pulumi.InvokeOption) (*LookupContentTypeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupContentTypeResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getContentType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentTypeArgs struct {
	// Content type identifier.
	ContentTypeId string `pulumi:"contentTypeId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Content type contract details.
type LookupContentTypeResult struct {
	// Content type description.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Content type schema.
	Schema interface{} `pulumi:"schema"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Content type version.
	Version *string `pulumi:"version"`
}

func LookupContentTypeOutput(ctx *pulumi.Context, args LookupContentTypeOutputArgs, opts ...pulumi.InvokeOption) LookupContentTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentTypeResult, error) {
			args := v.(LookupContentTypeArgs)
			r, err := LookupContentType(ctx, &args, opts...)
			var s LookupContentTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentTypeResultOutput)
}

type LookupContentTypeOutputArgs struct {
	// Content type identifier.
	ContentTypeId pulumi.StringInput `pulumi:"contentTypeId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupContentTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentTypeArgs)(nil)).Elem()
}

// Content type contract details.
type LookupContentTypeResultOutput struct{ *pulumi.OutputState }

func (LookupContentTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentTypeResult)(nil)).Elem()
}

func (o LookupContentTypeResultOutput) ToLookupContentTypeResultOutput() LookupContentTypeResultOutput {
	return o
}

func (o LookupContentTypeResultOutput) ToLookupContentTypeResultOutputWithContext(ctx context.Context) LookupContentTypeResultOutput {
	return o
}

// Content type description.
func (o LookupContentTypeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTypeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupContentTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupContentTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Content type schema.
func (o LookupContentTypeResultOutput) Schema() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupContentTypeResult) interface{} { return v.Schema }).(pulumi.AnyOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupContentTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

// Content type version.
func (o LookupContentTypeResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTypeResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentTypeResultOutput{})
}

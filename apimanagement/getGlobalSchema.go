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

// Gets the details of the Schema specified by its identifier.
// Azure REST API version: 2022-08-01.
func LookupGlobalSchema(ctx *pulumi.Context, args *LookupGlobalSchemaArgs, opts ...pulumi.InvokeOption) (*LookupGlobalSchemaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGlobalSchemaResult
	err := ctx.Invoke("azure-native:apimanagement:getGlobalSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalSchemaArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId string `pulumi:"schemaId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Global Schema Contract details.
type LookupGlobalSchemaResult struct {
	// Free-form schema entity description.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Schema Type. Immutable.
	SchemaType string `pulumi:"schemaType"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Json-encoded string for non json-based schema.
	Value interface{} `pulumi:"value"`
}

func LookupGlobalSchemaOutput(ctx *pulumi.Context, args LookupGlobalSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalSchemaResult, error) {
			args := v.(LookupGlobalSchemaArgs)
			r, err := LookupGlobalSchema(ctx, &args, opts...)
			var s LookupGlobalSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalSchemaResultOutput)
}

type LookupGlobalSchemaOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId pulumi.StringInput `pulumi:"schemaId"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGlobalSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalSchemaArgs)(nil)).Elem()
}

// Global Schema Contract details.
type LookupGlobalSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalSchemaResult)(nil)).Elem()
}

func (o LookupGlobalSchemaResultOutput) ToLookupGlobalSchemaResultOutput() LookupGlobalSchemaResultOutput {
	return o
}

func (o LookupGlobalSchemaResultOutput) ToLookupGlobalSchemaResultOutputWithContext(ctx context.Context) LookupGlobalSchemaResultOutput {
	return o
}

func (o LookupGlobalSchemaResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGlobalSchemaResult] {
	return pulumix.Output[LookupGlobalSchemaResult]{
		OutputState: o.OutputState,
	}
}

// Free-form schema entity description.
func (o LookupGlobalSchemaResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGlobalSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGlobalSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Schema Type. Immutable.
func (o LookupGlobalSchemaResultOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) string { return v.SchemaType }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGlobalSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

// Json-encoded string for non json-based schema.
func (o LookupGlobalSchemaResultOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupGlobalSchemaResult) interface{} { return v.Value }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalSchemaResultOutput{})
}

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
// Azure REST API version: 2022-09-01-preview.
func LookupWorkspaceGlobalSchema(ctx *pulumi.Context, args *LookupWorkspaceGlobalSchemaArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceGlobalSchemaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceGlobalSchemaResult
	err := ctx.Invoke("azure-native:apimanagement:getWorkspaceGlobalSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceGlobalSchemaArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId string `pulumi:"schemaId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Global Schema Contract details.
type LookupWorkspaceGlobalSchemaResult struct {
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

func LookupWorkspaceGlobalSchemaOutput(ctx *pulumi.Context, args LookupWorkspaceGlobalSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceGlobalSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceGlobalSchemaResult, error) {
			args := v.(LookupWorkspaceGlobalSchemaArgs)
			r, err := LookupWorkspaceGlobalSchema(ctx, &args, opts...)
			var s LookupWorkspaceGlobalSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceGlobalSchemaResultOutput)
}

type LookupWorkspaceGlobalSchemaOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId pulumi.StringInput `pulumi:"schemaId"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceGlobalSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceGlobalSchemaArgs)(nil)).Elem()
}

// Global Schema Contract details.
type LookupWorkspaceGlobalSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceGlobalSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceGlobalSchemaResult)(nil)).Elem()
}

func (o LookupWorkspaceGlobalSchemaResultOutput) ToLookupWorkspaceGlobalSchemaResultOutput() LookupWorkspaceGlobalSchemaResultOutput {
	return o
}

func (o LookupWorkspaceGlobalSchemaResultOutput) ToLookupWorkspaceGlobalSchemaResultOutputWithContext(ctx context.Context) LookupWorkspaceGlobalSchemaResultOutput {
	return o
}

func (o LookupWorkspaceGlobalSchemaResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceGlobalSchemaResult] {
	return pulumix.Output[LookupWorkspaceGlobalSchemaResult]{
		OutputState: o.OutputState,
	}
}

// Free-form schema entity description.
func (o LookupWorkspaceGlobalSchemaResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceGlobalSchemaResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceGlobalSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceGlobalSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceGlobalSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceGlobalSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Schema Type. Immutable.
func (o LookupWorkspaceGlobalSchemaResultOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceGlobalSchemaResult) string { return v.SchemaType }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceGlobalSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceGlobalSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

// Json-encoded string for non json-based schema.
func (o LookupWorkspaceGlobalSchemaResultOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWorkspaceGlobalSchemaResult) interface{} { return v.Value }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceGlobalSchemaResultOutput{})
}

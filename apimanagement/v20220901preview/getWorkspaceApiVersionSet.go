// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the Api Version Set specified by its identifier.
func LookupWorkspaceApiVersionSet(ctx *pulumi.Context, args *LookupWorkspaceApiVersionSetArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceApiVersionSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceApiVersionSetResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getWorkspaceApiVersionSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceApiVersionSetArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Api Version Set identifier. Must be unique in the current API Management service instance.
	VersionSetId string `pulumi:"versionSetId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// API Version Set Contract details.
type LookupWorkspaceApiVersionSetResult struct {
	// Description of API Version Set.
	Description *string `pulumi:"description"`
	// Name of API Version Set
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	// Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
	VersionQueryName *string `pulumi:"versionQueryName"`
	// An value that determines where the API Version identifier will be located in a HTTP request.
	VersioningScheme string `pulumi:"versioningScheme"`
}

func LookupWorkspaceApiVersionSetOutput(ctx *pulumi.Context, args LookupWorkspaceApiVersionSetOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceApiVersionSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceApiVersionSetResult, error) {
			args := v.(LookupWorkspaceApiVersionSetArgs)
			r, err := LookupWorkspaceApiVersionSet(ctx, &args, opts...)
			var s LookupWorkspaceApiVersionSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceApiVersionSetResultOutput)
}

type LookupWorkspaceApiVersionSetOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Api Version Set identifier. Must be unique in the current API Management service instance.
	VersionSetId pulumi.StringInput `pulumi:"versionSetId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupWorkspaceApiVersionSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceApiVersionSetArgs)(nil)).Elem()
}

// API Version Set Contract details.
type LookupWorkspaceApiVersionSetResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceApiVersionSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceApiVersionSetResult)(nil)).Elem()
}

func (o LookupWorkspaceApiVersionSetResultOutput) ToLookupWorkspaceApiVersionSetResultOutput() LookupWorkspaceApiVersionSetResultOutput {
	return o
}

func (o LookupWorkspaceApiVersionSetResultOutput) ToLookupWorkspaceApiVersionSetResultOutputWithContext(ctx context.Context) LookupWorkspaceApiVersionSetResultOutput {
	return o
}

// Description of API Version Set.
func (o LookupWorkspaceApiVersionSetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiVersionSetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of API Version Set
func (o LookupWorkspaceApiVersionSetResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiVersionSetResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceApiVersionSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiVersionSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceApiVersionSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiVersionSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceApiVersionSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiVersionSetResult) string { return v.Type }).(pulumi.StringOutput)
}

// Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
func (o LookupWorkspaceApiVersionSetResultOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiVersionSetResult) *string { return v.VersionHeaderName }).(pulumi.StringPtrOutput)
}

// Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
func (o LookupWorkspaceApiVersionSetResultOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceApiVersionSetResult) *string { return v.VersionQueryName }).(pulumi.StringPtrOutput)
}

// An value that determines where the API Version identifier will be located in a HTTP request.
func (o LookupWorkspaceApiVersionSetResultOutput) VersioningScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceApiVersionSetResult) string { return v.VersioningScheme }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceApiVersionSetResultOutput{})
}
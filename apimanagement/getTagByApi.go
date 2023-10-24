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

// Get tag associated with the API.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
func LookupTagByApi(ctx *pulumi.Context, args *LookupTagByApiArgs, opts ...pulumi.InvokeOption) (*LookupTagByApiResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTagByApiResult
	err := ctx.Invoke("azure-native:apimanagement:getTagByApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagByApiArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// Tag Contract details.
type LookupTagByApiResult struct {
	// Tag name.
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTagByApiOutput(ctx *pulumi.Context, args LookupTagByApiOutputArgs, opts ...pulumi.InvokeOption) LookupTagByApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagByApiResult, error) {
			args := v.(LookupTagByApiArgs)
			r, err := LookupTagByApi(ctx, &args, opts...)
			var s LookupTagByApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagByApiResultOutput)
}

type LookupTagByApiOutputArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput `pulumi:"tagId"`
}

func (LookupTagByApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagByApiArgs)(nil)).Elem()
}

// Tag Contract details.
type LookupTagByApiResultOutput struct{ *pulumi.OutputState }

func (LookupTagByApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagByApiResult)(nil)).Elem()
}

func (o LookupTagByApiResultOutput) ToLookupTagByApiResultOutput() LookupTagByApiResultOutput {
	return o
}

func (o LookupTagByApiResultOutput) ToLookupTagByApiResultOutputWithContext(ctx context.Context) LookupTagByApiResultOutput {
	return o
}

func (o LookupTagByApiResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTagByApiResult] {
	return pulumix.Output[LookupTagByApiResult]{
		OutputState: o.OutputState,
	}
}

// Tag name.
func (o LookupTagByApiResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByApiResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTagByApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByApiResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTagByApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTagByApiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByApiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagByApiResultOutput{})
}

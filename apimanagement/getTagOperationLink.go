// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the operation link for the tag.
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
func LookupTagOperationLink(ctx *pulumi.Context, args *LookupTagOperationLinkArgs, opts ...pulumi.InvokeOption) (*LookupTagOperationLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTagOperationLinkResult
	err := ctx.Invoke("azure-native:apimanagement:getTagOperationLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagOperationLinkArgs struct {
	// Tag-operation link identifier. Must be unique in the current API Management service instance.
	OperationLinkId string `pulumi:"operationLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// Tag-operation link details.
type LookupTagOperationLinkResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Full resource Id of an API operation.
	OperationId string `pulumi:"operationId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTagOperationLinkOutput(ctx *pulumi.Context, args LookupTagOperationLinkOutputArgs, opts ...pulumi.InvokeOption) LookupTagOperationLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagOperationLinkResult, error) {
			args := v.(LookupTagOperationLinkArgs)
			r, err := LookupTagOperationLink(ctx, &args, opts...)
			var s LookupTagOperationLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagOperationLinkResultOutput)
}

type LookupTagOperationLinkOutputArgs struct {
	// Tag-operation link identifier. Must be unique in the current API Management service instance.
	OperationLinkId pulumi.StringInput `pulumi:"operationLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput `pulumi:"tagId"`
}

func (LookupTagOperationLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagOperationLinkArgs)(nil)).Elem()
}

// Tag-operation link details.
type LookupTagOperationLinkResultOutput struct{ *pulumi.OutputState }

func (LookupTagOperationLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagOperationLinkResult)(nil)).Elem()
}

func (o LookupTagOperationLinkResultOutput) ToLookupTagOperationLinkResultOutput() LookupTagOperationLinkResultOutput {
	return o
}

func (o LookupTagOperationLinkResultOutput) ToLookupTagOperationLinkResultOutputWithContext(ctx context.Context) LookupTagOperationLinkResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTagOperationLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagOperationLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTagOperationLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagOperationLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Full resource Id of an API operation.
func (o LookupTagOperationLinkResultOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagOperationLinkResult) string { return v.OperationId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTagOperationLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagOperationLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagOperationLinkResultOutput{})
}

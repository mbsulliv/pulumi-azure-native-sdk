// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists policy resources that reference the policy fragment.
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
func ListWorkspacePolicyFragmentReferences(ctx *pulumi.Context, args *ListWorkspacePolicyFragmentReferencesArgs, opts ...pulumi.InvokeOption) (*ListWorkspacePolicyFragmentReferencesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkspacePolicyFragmentReferencesResult
	err := ctx.Invoke("azure-native:apimanagement:listWorkspacePolicyFragmentReferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspacePolicyFragmentReferencesArgs struct {
	// A resource identifier.
	Id string `pulumi:"id"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Number of records to skip.
	Skip *int `pulumi:"skip"`
	// Number of records to return.
	Top *int `pulumi:"top"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// A collection of resources.
type ListWorkspacePolicyFragmentReferencesResult struct {
	// Total record count number.
	Count *float64 `pulumi:"count"`
	// Next page link if any.
	NextLink *string `pulumi:"nextLink"`
	// A collection of resources.
	Value []ResourceCollectionResponseValue `pulumi:"value"`
}

func ListWorkspacePolicyFragmentReferencesOutput(ctx *pulumi.Context, args ListWorkspacePolicyFragmentReferencesOutputArgs, opts ...pulumi.InvokeOption) ListWorkspacePolicyFragmentReferencesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspacePolicyFragmentReferencesResult, error) {
			args := v.(ListWorkspacePolicyFragmentReferencesArgs)
			r, err := ListWorkspacePolicyFragmentReferences(ctx, &args, opts...)
			var s ListWorkspacePolicyFragmentReferencesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspacePolicyFragmentReferencesResultOutput)
}

type ListWorkspacePolicyFragmentReferencesOutputArgs struct {
	// A resource identifier.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Number of records to skip.
	Skip pulumi.IntPtrInput `pulumi:"skip"`
	// Number of records to return.
	Top pulumi.IntPtrInput `pulumi:"top"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (ListWorkspacePolicyFragmentReferencesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspacePolicyFragmentReferencesArgs)(nil)).Elem()
}

// A collection of resources.
type ListWorkspacePolicyFragmentReferencesResultOutput struct{ *pulumi.OutputState }

func (ListWorkspacePolicyFragmentReferencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspacePolicyFragmentReferencesResult)(nil)).Elem()
}

func (o ListWorkspacePolicyFragmentReferencesResultOutput) ToListWorkspacePolicyFragmentReferencesResultOutput() ListWorkspacePolicyFragmentReferencesResultOutput {
	return o
}

func (o ListWorkspacePolicyFragmentReferencesResultOutput) ToListWorkspacePolicyFragmentReferencesResultOutputWithContext(ctx context.Context) ListWorkspacePolicyFragmentReferencesResultOutput {
	return o
}

// Total record count number.
func (o ListWorkspacePolicyFragmentReferencesResultOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListWorkspacePolicyFragmentReferencesResult) *float64 { return v.Count }).(pulumi.Float64PtrOutput)
}

// Next page link if any.
func (o ListWorkspacePolicyFragmentReferencesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspacePolicyFragmentReferencesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// A collection of resources.
func (o ListWorkspacePolicyFragmentReferencesResultOutput) Value() ResourceCollectionResponseValueArrayOutput {
	return o.ApplyT(func(v ListWorkspacePolicyFragmentReferencesResult) []ResourceCollectionResponseValue { return v.Value }).(ResourceCollectionResponseValueArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspacePolicyFragmentReferencesResultOutput{})
}

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

// Lists policy resources that reference the policy fragment.
// Azure REST API version: 2022-08-01.
func ListPolicyFragmentReferences(ctx *pulumi.Context, args *ListPolicyFragmentReferencesArgs, opts ...pulumi.InvokeOption) (*ListPolicyFragmentReferencesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListPolicyFragmentReferencesResult
	err := ctx.Invoke("azure-native:apimanagement:listPolicyFragmentReferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPolicyFragmentReferencesArgs struct {
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
}

// A collection of resources.
type ListPolicyFragmentReferencesResult struct {
	// Total record count number.
	Count *float64 `pulumi:"count"`
	// Next page link if any.
	NextLink *string `pulumi:"nextLink"`
	// A collection of resources.
	Value []ResourceCollectionResponseValue `pulumi:"value"`
}

func ListPolicyFragmentReferencesOutput(ctx *pulumi.Context, args ListPolicyFragmentReferencesOutputArgs, opts ...pulumi.InvokeOption) ListPolicyFragmentReferencesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPolicyFragmentReferencesResult, error) {
			args := v.(ListPolicyFragmentReferencesArgs)
			r, err := ListPolicyFragmentReferences(ctx, &args, opts...)
			var s ListPolicyFragmentReferencesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPolicyFragmentReferencesResultOutput)
}

type ListPolicyFragmentReferencesOutputArgs struct {
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
}

func (ListPolicyFragmentReferencesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicyFragmentReferencesArgs)(nil)).Elem()
}

// A collection of resources.
type ListPolicyFragmentReferencesResultOutput struct{ *pulumi.OutputState }

func (ListPolicyFragmentReferencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicyFragmentReferencesResult)(nil)).Elem()
}

func (o ListPolicyFragmentReferencesResultOutput) ToListPolicyFragmentReferencesResultOutput() ListPolicyFragmentReferencesResultOutput {
	return o
}

func (o ListPolicyFragmentReferencesResultOutput) ToListPolicyFragmentReferencesResultOutputWithContext(ctx context.Context) ListPolicyFragmentReferencesResultOutput {
	return o
}

func (o ListPolicyFragmentReferencesResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListPolicyFragmentReferencesResult] {
	return pulumix.Output[ListPolicyFragmentReferencesResult]{
		OutputState: o.OutputState,
	}
}

// Total record count number.
func (o ListPolicyFragmentReferencesResultOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListPolicyFragmentReferencesResult) *float64 { return v.Count }).(pulumi.Float64PtrOutput)
}

// Next page link if any.
func (o ListPolicyFragmentReferencesResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListPolicyFragmentReferencesResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// A collection of resources.
func (o ListPolicyFragmentReferencesResultOutput) Value() ResourceCollectionResponseValueArrayOutput {
	return o.ApplyT(func(v ListPolicyFragmentReferencesResult) []ResourceCollectionResponseValue { return v.Value }).(ResourceCollectionResponseValueArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPolicyFragmentReferencesResultOutput{})
}

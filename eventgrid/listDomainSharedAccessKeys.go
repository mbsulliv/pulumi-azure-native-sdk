// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List the two keys used to publish to a domain.
// Azure REST API version: 2022-06-15.
//
// Other available API versions: 2020-04-01-preview, 2023-06-01-preview.
func ListDomainSharedAccessKeys(ctx *pulumi.Context, args *ListDomainSharedAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListDomainSharedAccessKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDomainSharedAccessKeysResult
	err := ctx.Invoke("azure-native:eventgrid:listDomainSharedAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDomainSharedAccessKeysArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Shared access keys of the Domain.
type ListDomainSharedAccessKeysResult struct {
	// Shared access key1 for the domain.
	Key1 *string `pulumi:"key1"`
	// Shared access key2 for the domain.
	Key2 *string `pulumi:"key2"`
}

func ListDomainSharedAccessKeysOutput(ctx *pulumi.Context, args ListDomainSharedAccessKeysOutputArgs, opts ...pulumi.InvokeOption) ListDomainSharedAccessKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDomainSharedAccessKeysResult, error) {
			args := v.(ListDomainSharedAccessKeysArgs)
			r, err := ListDomainSharedAccessKeys(ctx, &args, opts...)
			var s ListDomainSharedAccessKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDomainSharedAccessKeysResultOutput)
}

type ListDomainSharedAccessKeysOutputArgs struct {
	// Name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDomainSharedAccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainSharedAccessKeysArgs)(nil)).Elem()
}

// Shared access keys of the Domain.
type ListDomainSharedAccessKeysResultOutput struct{ *pulumi.OutputState }

func (ListDomainSharedAccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDomainSharedAccessKeysResult)(nil)).Elem()
}

func (o ListDomainSharedAccessKeysResultOutput) ToListDomainSharedAccessKeysResultOutput() ListDomainSharedAccessKeysResultOutput {
	return o
}

func (o ListDomainSharedAccessKeysResultOutput) ToListDomainSharedAccessKeysResultOutputWithContext(ctx context.Context) ListDomainSharedAccessKeysResultOutput {
	return o
}

func (o ListDomainSharedAccessKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListDomainSharedAccessKeysResult] {
	return pulumix.Output[ListDomainSharedAccessKeysResult]{
		OutputState: o.OutputState,
	}
}

// Shared access key1 for the domain.
func (o ListDomainSharedAccessKeysResultOutput) Key1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListDomainSharedAccessKeysResult) *string { return v.Key1 }).(pulumi.StringPtrOutput)
}

// Shared access key2 for the domain.
func (o ListDomainSharedAccessKeysResultOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListDomainSharedAccessKeysResult) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDomainSharedAccessKeysResultOutput{})
}

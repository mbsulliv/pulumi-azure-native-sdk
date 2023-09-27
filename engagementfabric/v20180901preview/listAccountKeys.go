// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The list of the EngagementFabric account keys
func ListAccountKeys(ctx *pulumi.Context, args *ListAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListAccountKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAccountKeysResult
	err := ctx.Invoke("azure-native:engagementfabric/v20180901preview:listAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountKeysArgs struct {
	// Account Name
	AccountName string `pulumi:"accountName"`
	// Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The list of the EngagementFabric account keys
type ListAccountKeysResult struct {
	// Account keys
	Value []KeyDescriptionResponse `pulumi:"value"`
}

func ListAccountKeysOutput(ctx *pulumi.Context, args ListAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAccountKeysResult, error) {
			args := v.(ListAccountKeysArgs)
			r, err := ListAccountKeys(ctx, &args, opts...)
			var s ListAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAccountKeysResultOutput)
}

type ListAccountKeysOutputArgs struct {
	// Account Name
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Resource Group Name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountKeysArgs)(nil)).Elem()
}

// The list of the EngagementFabric account keys
type ListAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountKeysResult)(nil)).Elem()
}

func (o ListAccountKeysResultOutput) ToListAccountKeysResultOutput() ListAccountKeysResultOutput {
	return o
}

func (o ListAccountKeysResultOutput) ToListAccountKeysResultOutputWithContext(ctx context.Context) ListAccountKeysResultOutput {
	return o
}

func (o ListAccountKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListAccountKeysResult] {
	return pulumix.Output[ListAccountKeysResult]{
		OutputState: o.OutputState,
	}
}

// Account keys
func (o ListAccountKeysResultOutput) Value() KeyDescriptionResponseArrayOutput {
	return o.ApplyT(func(v ListAccountKeysResult) []KeyDescriptionResponse { return v.Value }).(KeyDescriptionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountKeysResultOutput{})
}

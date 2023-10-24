// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mixedreality

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List Both of the 2 Keys of a Remote Rendering Account
// Azure REST API version: 2021-01-01.
//
// Other available API versions: 2021-03-01-preview.
func ListRemoteRenderingAccountKeys(ctx *pulumi.Context, args *ListRemoteRenderingAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListRemoteRenderingAccountKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListRemoteRenderingAccountKeysResult
	err := ctx.Invoke("azure-native:mixedreality:listRemoteRenderingAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemoteRenderingAccountKeysArgs struct {
	// Name of an Mixed Reality Account.
	AccountName string `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Developer Keys of account
type ListRemoteRenderingAccountKeysResult struct {
	// value of primary key.
	PrimaryKey string `pulumi:"primaryKey"`
	// value of secondary key.
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListRemoteRenderingAccountKeysOutput(ctx *pulumi.Context, args ListRemoteRenderingAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListRemoteRenderingAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRemoteRenderingAccountKeysResult, error) {
			args := v.(ListRemoteRenderingAccountKeysArgs)
			r, err := ListRemoteRenderingAccountKeys(ctx, &args, opts...)
			var s ListRemoteRenderingAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRemoteRenderingAccountKeysResultOutput)
}

type ListRemoteRenderingAccountKeysOutputArgs struct {
	// Name of an Mixed Reality Account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListRemoteRenderingAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemoteRenderingAccountKeysArgs)(nil)).Elem()
}

// Developer Keys of account
type ListRemoteRenderingAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListRemoteRenderingAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemoteRenderingAccountKeysResult)(nil)).Elem()
}

func (o ListRemoteRenderingAccountKeysResultOutput) ToListRemoteRenderingAccountKeysResultOutput() ListRemoteRenderingAccountKeysResultOutput {
	return o
}

func (o ListRemoteRenderingAccountKeysResultOutput) ToListRemoteRenderingAccountKeysResultOutputWithContext(ctx context.Context) ListRemoteRenderingAccountKeysResultOutput {
	return o
}

func (o ListRemoteRenderingAccountKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListRemoteRenderingAccountKeysResult] {
	return pulumix.Output[ListRemoteRenderingAccountKeysResult]{
		OutputState: o.OutputState,
	}
}

// value of primary key.
func (o ListRemoteRenderingAccountKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemoteRenderingAccountKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// value of secondary key.
func (o ListRemoteRenderingAccountKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemoteRenderingAccountKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRemoteRenderingAccountKeysResultOutput{})
}

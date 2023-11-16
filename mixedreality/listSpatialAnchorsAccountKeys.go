// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mixedreality

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List Both of the 2 Keys of a Spatial Anchors Account
// Azure REST API version: 2021-01-01.
//
// Other available API versions: 2021-03-01-preview.
func ListSpatialAnchorsAccountKeys(ctx *pulumi.Context, args *ListSpatialAnchorsAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListSpatialAnchorsAccountKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSpatialAnchorsAccountKeysResult
	err := ctx.Invoke("azure-native:mixedreality:listSpatialAnchorsAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSpatialAnchorsAccountKeysArgs struct {
	// Name of an Mixed Reality Account.
	AccountName string `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Developer Keys of account
type ListSpatialAnchorsAccountKeysResult struct {
	// value of primary key.
	PrimaryKey string `pulumi:"primaryKey"`
	// value of secondary key.
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListSpatialAnchorsAccountKeysOutput(ctx *pulumi.Context, args ListSpatialAnchorsAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListSpatialAnchorsAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSpatialAnchorsAccountKeysResult, error) {
			args := v.(ListSpatialAnchorsAccountKeysArgs)
			r, err := ListSpatialAnchorsAccountKeys(ctx, &args, opts...)
			var s ListSpatialAnchorsAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSpatialAnchorsAccountKeysResultOutput)
}

type ListSpatialAnchorsAccountKeysOutputArgs struct {
	// Name of an Mixed Reality Account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSpatialAnchorsAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSpatialAnchorsAccountKeysArgs)(nil)).Elem()
}

// Developer Keys of account
type ListSpatialAnchorsAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListSpatialAnchorsAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSpatialAnchorsAccountKeysResult)(nil)).Elem()
}

func (o ListSpatialAnchorsAccountKeysResultOutput) ToListSpatialAnchorsAccountKeysResultOutput() ListSpatialAnchorsAccountKeysResultOutput {
	return o
}

func (o ListSpatialAnchorsAccountKeysResultOutput) ToListSpatialAnchorsAccountKeysResultOutputWithContext(ctx context.Context) ListSpatialAnchorsAccountKeysResultOutput {
	return o
}

// value of primary key.
func (o ListSpatialAnchorsAccountKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListSpatialAnchorsAccountKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// value of secondary key.
func (o ListSpatialAnchorsAccountKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListSpatialAnchorsAccountKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSpatialAnchorsAccountKeysResultOutput{})
}

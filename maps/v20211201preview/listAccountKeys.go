// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the keys to use with the Maps APIs. A key is used to authenticate and authorize access to the Maps REST APIs. Only one key is needed at a time; two are given to provide seamless key regeneration.
func ListAccountKeys(ctx *pulumi.Context, args *ListAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListAccountKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAccountKeysResult
	err := ctx.Invoke("azure-native:maps/v20211201preview:listAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountKeysArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of keys which can be used to access the Maps REST APIs. Two keys are provided for key rotation without interruption.
type ListAccountKeysResult struct {
	// The primary key for accessing the Maps REST APIs.
	PrimaryKey string `pulumi:"primaryKey"`
	// The last updated date and time of the primary key.
	PrimaryKeyLastUpdated string `pulumi:"primaryKeyLastUpdated"`
	// The secondary key for accessing the Maps REST APIs.
	SecondaryKey string `pulumi:"secondaryKey"`
	// The last updated date and time of the secondary key.
	SecondaryKeyLastUpdated string `pulumi:"secondaryKeyLastUpdated"`
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
	// The name of the Maps Account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountKeysArgs)(nil)).Elem()
}

// The set of keys which can be used to access the Maps REST APIs. Two keys are provided for key rotation without interruption.
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

// The primary key for accessing the Maps REST APIs.
func (o ListAccountKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The last updated date and time of the primary key.
func (o ListAccountKeysResultOutput) PrimaryKeyLastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountKeysResult) string { return v.PrimaryKeyLastUpdated }).(pulumi.StringOutput)
}

// The secondary key for accessing the Maps REST APIs.
func (o ListAccountKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

// The last updated date and time of the secondary key.
func (o ListAccountKeysResultOutput) SecondaryKeyLastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountKeysResult) string { return v.SecondaryKeyLastUpdated }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountKeysResultOutput{})
}

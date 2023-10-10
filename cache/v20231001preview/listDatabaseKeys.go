// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves the access keys for the RedisEnterprise database.
func ListDatabaseKeys(ctx *pulumi.Context, args *ListDatabaseKeysArgs, opts ...pulumi.InvokeOption) (*ListDatabaseKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDatabaseKeysResult
	err := ctx.Invoke("azure-native:cache/v20231001preview:listDatabaseKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseKeysArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The secret access keys used for authenticating connections to redis
type ListDatabaseKeysResult struct {
	// The current primary key that clients can use to authenticate
	PrimaryKey string `pulumi:"primaryKey"`
	// The current secondary key that clients can use to authenticate
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListDatabaseKeysOutput(ctx *pulumi.Context, args ListDatabaseKeysOutputArgs, opts ...pulumi.InvokeOption) ListDatabaseKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabaseKeysResult, error) {
			args := v.(ListDatabaseKeysArgs)
			r, err := ListDatabaseKeys(ctx, &args, opts...)
			var s ListDatabaseKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatabaseKeysResultOutput)
}

type ListDatabaseKeysOutputArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabaseKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseKeysArgs)(nil)).Elem()
}

// The secret access keys used for authenticating connections to redis
type ListDatabaseKeysResultOutput struct{ *pulumi.OutputState }

func (ListDatabaseKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseKeysResult)(nil)).Elem()
}

func (o ListDatabaseKeysResultOutput) ToListDatabaseKeysResultOutput() ListDatabaseKeysResultOutput {
	return o
}

func (o ListDatabaseKeysResultOutput) ToListDatabaseKeysResultOutputWithContext(ctx context.Context) ListDatabaseKeysResultOutput {
	return o
}

func (o ListDatabaseKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListDatabaseKeysResult] {
	return pulumix.Output[ListDatabaseKeysResult]{
		OutputState: o.OutputState,
	}
}

// The current primary key that clients can use to authenticate
func (o ListDatabaseKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// The current secondary key that clients can use to authenticate
func (o ListDatabaseKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDatabaseKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabaseKeysResultOutput{})
}

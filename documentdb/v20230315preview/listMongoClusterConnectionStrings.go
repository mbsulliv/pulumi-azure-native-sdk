// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List mongo cluster connection strings. This includes the default connection string using SCRAM-SHA-256, as well as other connection strings supported by the cluster.
func ListMongoClusterConnectionStrings(ctx *pulumi.Context, args *ListMongoClusterConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListMongoClusterConnectionStringsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListMongoClusterConnectionStringsResult
	err := ctx.Invoke("azure-native:documentdb/v20230315preview:listMongoClusterConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMongoClusterConnectionStringsArgs struct {
	// The name of the mongo cluster.
	MongoClusterName string `pulumi:"mongoClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connection strings for the given mongo cluster.
type ListMongoClusterConnectionStringsResult struct {
	// An array that contains the connection strings for a mongo cluster.
	ConnectionStrings []ConnectionStringResponse `pulumi:"connectionStrings"`
}

func ListMongoClusterConnectionStringsOutput(ctx *pulumi.Context, args ListMongoClusterConnectionStringsOutputArgs, opts ...pulumi.InvokeOption) ListMongoClusterConnectionStringsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMongoClusterConnectionStringsResult, error) {
			args := v.(ListMongoClusterConnectionStringsArgs)
			r, err := ListMongoClusterConnectionStrings(ctx, &args, opts...)
			var s ListMongoClusterConnectionStringsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMongoClusterConnectionStringsResultOutput)
}

type ListMongoClusterConnectionStringsOutputArgs struct {
	// The name of the mongo cluster.
	MongoClusterName pulumi.StringInput `pulumi:"mongoClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMongoClusterConnectionStringsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMongoClusterConnectionStringsArgs)(nil)).Elem()
}

// The connection strings for the given mongo cluster.
type ListMongoClusterConnectionStringsResultOutput struct{ *pulumi.OutputState }

func (ListMongoClusterConnectionStringsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMongoClusterConnectionStringsResult)(nil)).Elem()
}

func (o ListMongoClusterConnectionStringsResultOutput) ToListMongoClusterConnectionStringsResultOutput() ListMongoClusterConnectionStringsResultOutput {
	return o
}

func (o ListMongoClusterConnectionStringsResultOutput) ToListMongoClusterConnectionStringsResultOutputWithContext(ctx context.Context) ListMongoClusterConnectionStringsResultOutput {
	return o
}

func (o ListMongoClusterConnectionStringsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListMongoClusterConnectionStringsResult] {
	return pulumix.Output[ListMongoClusterConnectionStringsResult]{
		OutputState: o.OutputState,
	}
}

// An array that contains the connection strings for a mongo cluster.
func (o ListMongoClusterConnectionStringsResultOutput) ConnectionStrings() ConnectionStringResponseArrayOutput {
	return o.ApplyT(func(v ListMongoClusterConnectionStringsResult) []ConnectionStringResponse { return v.ConnectionStrings }).(ConnectionStringResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMongoClusterConnectionStringsResultOutput{})
}

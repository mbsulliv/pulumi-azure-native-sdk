// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the connection strings for the specified Azure Cosmos DB database account.
func ListDatabaseAccountConnectionStrings(ctx *pulumi.Context, args *ListDatabaseAccountConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListDatabaseAccountConnectionStringsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDatabaseAccountConnectionStringsResult
	err := ctx.Invoke("azure-native:documentdb/v20230415:listDatabaseAccountConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseAccountConnectionStringsArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connection strings for the given database account.
type ListDatabaseAccountConnectionStringsResult struct {
	// An array that contains the connection strings for the Cosmos DB account.
	ConnectionStrings []DatabaseAccountConnectionStringResponse `pulumi:"connectionStrings"`
}

func ListDatabaseAccountConnectionStringsOutput(ctx *pulumi.Context, args ListDatabaseAccountConnectionStringsOutputArgs, opts ...pulumi.InvokeOption) ListDatabaseAccountConnectionStringsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDatabaseAccountConnectionStringsResult, error) {
			args := v.(ListDatabaseAccountConnectionStringsArgs)
			r, err := ListDatabaseAccountConnectionStrings(ctx, &args, opts...)
			var s ListDatabaseAccountConnectionStringsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDatabaseAccountConnectionStringsResultOutput)
}

type ListDatabaseAccountConnectionStringsOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabaseAccountConnectionStringsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountConnectionStringsArgs)(nil)).Elem()
}

// The connection strings for the given database account.
type ListDatabaseAccountConnectionStringsResultOutput struct{ *pulumi.OutputState }

func (ListDatabaseAccountConnectionStringsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabaseAccountConnectionStringsResult)(nil)).Elem()
}

func (o ListDatabaseAccountConnectionStringsResultOutput) ToListDatabaseAccountConnectionStringsResultOutput() ListDatabaseAccountConnectionStringsResultOutput {
	return o
}

func (o ListDatabaseAccountConnectionStringsResultOutput) ToListDatabaseAccountConnectionStringsResultOutputWithContext(ctx context.Context) ListDatabaseAccountConnectionStringsResultOutput {
	return o
}

// An array that contains the connection strings for the Cosmos DB account.
func (o ListDatabaseAccountConnectionStringsResultOutput) ConnectionStrings() DatabaseAccountConnectionStringResponseArrayOutput {
	return o.ApplyT(func(v ListDatabaseAccountConnectionStringsResult) []DatabaseAccountConnectionStringResponse {
		return v.ConnectionStrings
	}).(DatabaseAccountConnectionStringResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabaseAccountConnectionStringsResultOutput{})
}

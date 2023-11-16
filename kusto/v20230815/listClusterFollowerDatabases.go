// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230815

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a list of databases that are owned by this cluster and were followed by another cluster.
func ListClusterFollowerDatabases(ctx *pulumi.Context, args *ListClusterFollowerDatabasesArgs, opts ...pulumi.InvokeOption) (*ListClusterFollowerDatabasesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListClusterFollowerDatabasesResult
	err := ctx.Invoke("azure-native:kusto/v20230815:listClusterFollowerDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterFollowerDatabasesArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The list Kusto database principals operation response.
type ListClusterFollowerDatabasesResult struct {
	// The list of follower database result.
	Value []FollowerDatabaseDefinitionResponse `pulumi:"value"`
}

func ListClusterFollowerDatabasesOutput(ctx *pulumi.Context, args ListClusterFollowerDatabasesOutputArgs, opts ...pulumi.InvokeOption) ListClusterFollowerDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListClusterFollowerDatabasesResult, error) {
			args := v.(ListClusterFollowerDatabasesArgs)
			r, err := ListClusterFollowerDatabases(ctx, &args, opts...)
			var s ListClusterFollowerDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListClusterFollowerDatabasesResultOutput)
}

type ListClusterFollowerDatabasesOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListClusterFollowerDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterFollowerDatabasesArgs)(nil)).Elem()
}

// The list Kusto database principals operation response.
type ListClusterFollowerDatabasesResultOutput struct{ *pulumi.OutputState }

func (ListClusterFollowerDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListClusterFollowerDatabasesResult)(nil)).Elem()
}

func (o ListClusterFollowerDatabasesResultOutput) ToListClusterFollowerDatabasesResultOutput() ListClusterFollowerDatabasesResultOutput {
	return o
}

func (o ListClusterFollowerDatabasesResultOutput) ToListClusterFollowerDatabasesResultOutputWithContext(ctx context.Context) ListClusterFollowerDatabasesResultOutput {
	return o
}

// The list of follower database result.
func (o ListClusterFollowerDatabasesResultOutput) Value() FollowerDatabaseDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ListClusterFollowerDatabasesResult) []FollowerDatabaseDefinitionResponse { return v.Value }).(FollowerDatabaseDefinitionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListClusterFollowerDatabasesResultOutput{})
}

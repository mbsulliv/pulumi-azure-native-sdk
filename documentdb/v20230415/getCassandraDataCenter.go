// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230415

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the properties of a managed Cassandra data center.
func LookupCassandraDataCenter(ctx *pulumi.Context, args *LookupCassandraDataCenterArgs, opts ...pulumi.InvokeOption) (*LookupCassandraDataCenterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCassandraDataCenterResult
	err := ctx.Invoke("azure-native:documentdb/v20230415:getCassandraDataCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraDataCenterArgs struct {
	// Managed Cassandra cluster name.
	ClusterName string `pulumi:"clusterName"`
	// Data center name in a managed Cassandra cluster.
	DataCenterName string `pulumi:"dataCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A managed Cassandra data center.
type LookupCassandraDataCenterResult struct {
	// The unique resource identifier of the database account.
	Id string `pulumi:"id"`
	// The name of the database account.
	Name string `pulumi:"name"`
	// Properties of a managed Cassandra data center.
	Properties DataCenterResourceResponseProperties `pulumi:"properties"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupCassandraDataCenterOutput(ctx *pulumi.Context, args LookupCassandraDataCenterOutputArgs, opts ...pulumi.InvokeOption) LookupCassandraDataCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCassandraDataCenterResult, error) {
			args := v.(LookupCassandraDataCenterArgs)
			r, err := LookupCassandraDataCenter(ctx, &args, opts...)
			var s LookupCassandraDataCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCassandraDataCenterResultOutput)
}

type LookupCassandraDataCenterOutputArgs struct {
	// Managed Cassandra cluster name.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Data center name in a managed Cassandra cluster.
	DataCenterName pulumi.StringInput `pulumi:"dataCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCassandraDataCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraDataCenterArgs)(nil)).Elem()
}

// A managed Cassandra data center.
type LookupCassandraDataCenterResultOutput struct{ *pulumi.OutputState }

func (LookupCassandraDataCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCassandraDataCenterResult)(nil)).Elem()
}

func (o LookupCassandraDataCenterResultOutput) ToLookupCassandraDataCenterResultOutput() LookupCassandraDataCenterResultOutput {
	return o
}

func (o LookupCassandraDataCenterResultOutput) ToLookupCassandraDataCenterResultOutputWithContext(ctx context.Context) LookupCassandraDataCenterResultOutput {
	return o
}

// The unique resource identifier of the database account.
func (o LookupCassandraDataCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraDataCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the database account.
func (o LookupCassandraDataCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraDataCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of a managed Cassandra data center.
func (o LookupCassandraDataCenterResultOutput) Properties() DataCenterResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupCassandraDataCenterResult) DataCenterResourceResponseProperties { return v.Properties }).(DataCenterResourceResponsePropertiesOutput)
}

// The type of Azure resource.
func (o LookupCassandraDataCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCassandraDataCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCassandraDataCenterResultOutput{})
}

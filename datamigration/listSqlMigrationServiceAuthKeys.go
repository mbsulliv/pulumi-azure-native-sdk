// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datamigration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the List of Authentication Keys for Self Hosted Integration Runtime.
// Azure REST API version: 2022-03-30-preview.
func ListSqlMigrationServiceAuthKeys(ctx *pulumi.Context, args *ListSqlMigrationServiceAuthKeysArgs, opts ...pulumi.InvokeOption) (*ListSqlMigrationServiceAuthKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSqlMigrationServiceAuthKeysResult
	err := ctx.Invoke("azure-native:datamigration:listSqlMigrationServiceAuthKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSqlMigrationServiceAuthKeysArgs struct {
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL Migration Service.
	SqlMigrationServiceName string `pulumi:"sqlMigrationServiceName"`
}

// An authentication key.
type ListSqlMigrationServiceAuthKeysResult struct {
	// The first authentication key.
	AuthKey1 *string `pulumi:"authKey1"`
	// The second authentication key.
	AuthKey2 *string `pulumi:"authKey2"`
}

func ListSqlMigrationServiceAuthKeysOutput(ctx *pulumi.Context, args ListSqlMigrationServiceAuthKeysOutputArgs, opts ...pulumi.InvokeOption) ListSqlMigrationServiceAuthKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSqlMigrationServiceAuthKeysResult, error) {
			args := v.(ListSqlMigrationServiceAuthKeysArgs)
			r, err := ListSqlMigrationServiceAuthKeys(ctx, &args, opts...)
			var s ListSqlMigrationServiceAuthKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSqlMigrationServiceAuthKeysResultOutput)
}

type ListSqlMigrationServiceAuthKeysOutputArgs struct {
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the SQL Migration Service.
	SqlMigrationServiceName pulumi.StringInput `pulumi:"sqlMigrationServiceName"`
}

func (ListSqlMigrationServiceAuthKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSqlMigrationServiceAuthKeysArgs)(nil)).Elem()
}

// An authentication key.
type ListSqlMigrationServiceAuthKeysResultOutput struct{ *pulumi.OutputState }

func (ListSqlMigrationServiceAuthKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSqlMigrationServiceAuthKeysResult)(nil)).Elem()
}

func (o ListSqlMigrationServiceAuthKeysResultOutput) ToListSqlMigrationServiceAuthKeysResultOutput() ListSqlMigrationServiceAuthKeysResultOutput {
	return o
}

func (o ListSqlMigrationServiceAuthKeysResultOutput) ToListSqlMigrationServiceAuthKeysResultOutputWithContext(ctx context.Context) ListSqlMigrationServiceAuthKeysResultOutput {
	return o
}

// The first authentication key.
func (o ListSqlMigrationServiceAuthKeysResultOutput) AuthKey1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSqlMigrationServiceAuthKeysResult) *string { return v.AuthKey1 }).(pulumi.StringPtrOutput)
}

// The second authentication key.
func (o ListSqlMigrationServiceAuthKeysResultOutput) AuthKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSqlMigrationServiceAuthKeysResult) *string { return v.AuthKey2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSqlMigrationServiceAuthKeysResultOutput{})
}

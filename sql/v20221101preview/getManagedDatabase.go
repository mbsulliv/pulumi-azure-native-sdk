// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a managed database.
func LookupManagedDatabase(ctx *pulumi.Context, args *LookupManagedDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupManagedDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedDatabaseResult
	err := ctx.Invoke("azure-native:sql/v20221101preview:getManagedDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedDatabaseArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A managed database resource.
type LookupManagedDatabaseResult struct {
	// Collation of the metadata catalog.
	CatalogCollation *string `pulumi:"catalogCollation"`
	// Collation of the managed database.
	Collation *string `pulumi:"collation"`
	// Creation date of the database.
	CreationDate string `pulumi:"creationDate"`
	// Geo paired region.
	DefaultSecondaryLocation string `pulumi:"defaultSecondaryLocation"`
	// Earliest restore point in time for point in time restore.
	EarliestRestorePoint string `pulumi:"earliestRestorePoint"`
	// Instance Failover Group resource identifier that this managed database belongs to.
	FailoverGroupId string `pulumi:"failoverGroupId"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn *bool `pulumi:"isLedgerOn"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Status of the database.
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupManagedDatabaseOutput(ctx *pulumi.Context, args LookupManagedDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupManagedDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedDatabaseResult, error) {
			args := v.(LookupManagedDatabaseArgs)
			r, err := LookupManagedDatabase(ctx, &args, opts...)
			var s LookupManagedDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedDatabaseResultOutput)
}

type LookupManagedDatabaseOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedDatabaseArgs)(nil)).Elem()
}

// A managed database resource.
type LookupManagedDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupManagedDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedDatabaseResult)(nil)).Elem()
}

func (o LookupManagedDatabaseResultOutput) ToLookupManagedDatabaseResultOutput() LookupManagedDatabaseResultOutput {
	return o
}

func (o LookupManagedDatabaseResultOutput) ToLookupManagedDatabaseResultOutputWithContext(ctx context.Context) LookupManagedDatabaseResultOutput {
	return o
}

// Collation of the metadata catalog.
func (o LookupManagedDatabaseResultOutput) CatalogCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) *string { return v.CatalogCollation }).(pulumi.StringPtrOutput)
}

// Collation of the managed database.
func (o LookupManagedDatabaseResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

// Creation date of the database.
func (o LookupManagedDatabaseResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Geo paired region.
func (o LookupManagedDatabaseResultOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

// Earliest restore point in time for point in time restore.
func (o LookupManagedDatabaseResultOutput) EarliestRestorePoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.EarliestRestorePoint }).(pulumi.StringOutput)
}

// Instance Failover Group resource identifier that this managed database belongs to.
func (o LookupManagedDatabaseResultOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.FailoverGroupId }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupManagedDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
func (o LookupManagedDatabaseResultOutput) IsLedgerOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) *bool { return v.IsLedgerOn }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o LookupManagedDatabaseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupManagedDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Status of the database.
func (o LookupManagedDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupManagedDatabaseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupManagedDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedDatabaseResultOutput{})
}

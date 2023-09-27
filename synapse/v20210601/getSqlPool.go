// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get SQL pool properties
func LookupSqlPool(ctx *pulumi.Context, args *LookupSqlPoolArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlPoolResult
	err := ctx.Invoke("azure-native:synapse/v20210601:getSqlPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSqlPoolArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// A SQL Analytics pool
type LookupSqlPoolResult struct {
	// Collation mode
	Collation *string `pulumi:"collation"`
	// Date the SQL pool was created
	CreationDate string `pulumi:"creationDate"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Maximum size in bytes
	MaxSizeBytes *float64 `pulumi:"maxSizeBytes"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Resource state
	ProvisioningState *string `pulumi:"provisioningState"`
	// Backup database to restore from
	RecoverableDatabaseId *string `pulumi:"recoverableDatabaseId"`
	// Snapshot time to restore
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// SQL pool SKU
	Sku *SkuResponse `pulumi:"sku"`
	// Specifies the time that the sql pool was deleted
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// Resource status
	Status string `pulumi:"status"`
	// The storage account type used to store backups for this sql pool.
	StorageAccountType *string `pulumi:"storageAccountType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupSqlPoolResult
func (val *LookupSqlPoolResult) Defaults() *LookupSqlPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Collation == nil {
		collation_ := ""
		tmp.Collation = &collation_
	}
	if tmp.StorageAccountType == nil {
		storageAccountType_ := "GRS"
		tmp.StorageAccountType = &storageAccountType_
	}
	return &tmp
}

func LookupSqlPoolOutput(ctx *pulumi.Context, args LookupSqlPoolOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolResult, error) {
			args := v.(LookupSqlPoolArgs)
			r, err := LookupSqlPool(ctx, &args, opts...)
			var s LookupSqlPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolResultOutput)
}

type LookupSqlPoolOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName pulumi.StringInput `pulumi:"sqlPoolName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolArgs)(nil)).Elem()
}

// A SQL Analytics pool
type LookupSqlPoolResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolResult)(nil)).Elem()
}

func (o LookupSqlPoolResultOutput) ToLookupSqlPoolResultOutput() LookupSqlPoolResultOutput {
	return o
}

func (o LookupSqlPoolResultOutput) ToLookupSqlPoolResultOutputWithContext(ctx context.Context) LookupSqlPoolResultOutput {
	return o
}

func (o LookupSqlPoolResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlPoolResult] {
	return pulumix.Output[LookupSqlPoolResult]{
		OutputState: o.OutputState,
	}
}

// Collation mode
func (o LookupSqlPoolResultOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

// Date the SQL pool was created
func (o LookupSqlPoolResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSqlPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maximum size in bytes
func (o LookupSqlPoolResultOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *float64 { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

// The name of the resource
func (o LookupSqlPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource state
func (o LookupSqlPoolResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Backup database to restore from
func (o LookupSqlPoolResultOutput) RecoverableDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.RecoverableDatabaseId }).(pulumi.StringPtrOutput)
}

// Snapshot time to restore
func (o LookupSqlPoolResultOutput) RestorePointInTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.RestorePointInTime }).(pulumi.StringPtrOutput)
}

// SQL pool SKU
func (o LookupSqlPoolResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Specifies the time that the sql pool was deleted
func (o LookupSqlPoolResultOutput) SourceDatabaseDeletionDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.SourceDatabaseDeletionDate }).(pulumi.StringPtrOutput)
}

// Resource status
func (o LookupSqlPoolResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Status }).(pulumi.StringOutput)
}

// The storage account type used to store backups for this sql pool.
func (o LookupSqlPoolResultOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupSqlPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolResultOutput{})
}

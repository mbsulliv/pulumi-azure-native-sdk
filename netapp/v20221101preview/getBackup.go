// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the specified Backup under Backup Vault.
func LookupBackup(ctx *pulumi.Context, args *LookupBackupArgs, opts ...pulumi.InvokeOption) (*LookupBackupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupResult
	err := ctx.Invoke("azure-native:netapp/v20221101preview:getBackup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBackupArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the backup
	BackupName string `pulumi:"backupName"`
	// The name of the Backup Vault
	BackupVaultName string `pulumi:"backupVaultName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Backup under a Backup Vault
type LookupBackupResult struct {
	// UUID v4 used to identify the Backup
	BackupId string `pulumi:"backupId"`
	// Type of backup Manual or Scheduled
	BackupType string `pulumi:"backupType"`
	// The creation date of the backup
	CreationDate string `pulumi:"creationDate"`
	// Failure reason
	FailureReason string `pulumi:"failureReason"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Label for backup
	Label *string `pulumi:"label"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// Size of backup in bytes
	Size float64 `pulumi:"size"`
	// The name of the snapshot
	SnapshotName *string `pulumi:"snapshotName"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
	UseExistingSnapshot *bool `pulumi:"useExistingSnapshot"`
	// ResourceId used to identify the Volume
	VolumeResourceId string `pulumi:"volumeResourceId"`
}

// Defaults sets the appropriate defaults for LookupBackupResult
func (val *LookupBackupResult) Defaults() *LookupBackupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.UseExistingSnapshot == nil {
		useExistingSnapshot_ := false
		tmp.UseExistingSnapshot = &useExistingSnapshot_
	}
	return &tmp
}

func LookupBackupOutput(ctx *pulumi.Context, args LookupBackupOutputArgs, opts ...pulumi.InvokeOption) LookupBackupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupResult, error) {
			args := v.(LookupBackupArgs)
			r, err := LookupBackup(ctx, &args, opts...)
			var s LookupBackupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupResultOutput)
}

type LookupBackupOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the backup
	BackupName pulumi.StringInput `pulumi:"backupName"`
	// The name of the Backup Vault
	BackupVaultName pulumi.StringInput `pulumi:"backupVaultName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBackupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupArgs)(nil)).Elem()
}

// Backup under a Backup Vault
type LookupBackupResultOutput struct{ *pulumi.OutputState }

func (LookupBackupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupResult)(nil)).Elem()
}

func (o LookupBackupResultOutput) ToLookupBackupResultOutput() LookupBackupResultOutput {
	return o
}

func (o LookupBackupResultOutput) ToLookupBackupResultOutputWithContext(ctx context.Context) LookupBackupResultOutput {
	return o
}

// UUID v4 used to identify the Backup
func (o LookupBackupResultOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.BackupId }).(pulumi.StringOutput)
}

// Type of backup Manual or Scheduled
func (o LookupBackupResultOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.BackupType }).(pulumi.StringOutput)
}

// The creation date of the backup
func (o LookupBackupResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Failure reason
func (o LookupBackupResultOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.FailureReason }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBackupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Label for backup
func (o LookupBackupResultOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupResult) *string { return v.Label }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupBackupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupBackupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Size of backup in bytes
func (o LookupBackupResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBackupResult) float64 { return v.Size }).(pulumi.Float64Output)
}

// The name of the snapshot
func (o LookupBackupResultOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupResult) *string { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBackupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBackupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBackupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.Type }).(pulumi.StringOutput)
}

// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
func (o LookupBackupResultOutput) UseExistingSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBackupResult) *bool { return v.UseExistingSnapshot }).(pulumi.BoolPtrOutput)
}

// ResourceId used to identify the Volume
func (o LookupBackupResultOutput) VolumeResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupResult) string { return v.VolumeResourceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupResultOutput{})
}

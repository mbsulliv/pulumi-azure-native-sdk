// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets status of a web app backup that may be in progress, including secrets associated with the backup, such as the Azure Storage SAS URL. Also can be used to update the SAS URL for the backup if a new URL is passed in the request body.
func ListWebAppBackupStatusSecrets(ctx *pulumi.Context, args *ListWebAppBackupStatusSecretsArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupStatusSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppBackupStatusSecretsResult
	err := ctx.Invoke("azure-native:web/v20181101:listWebAppBackupStatusSecrets", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppBackupStatusSecretsArgs struct {
	// ID of backup.
	BackupId string `pulumi:"backupId"`
	// Name of the backup.
	BackupName *string `pulumi:"backupName"`
	// Schedule for the backup if it is executed periodically.
	BackupSchedule *BackupSchedule `pulumi:"backupSchedule"`
	// Databases included in the backup.
	Databases []DatabaseBackupSetting `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
	Enabled *bool `pulumi:"enabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of web app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SAS URL to the container.
	StorageAccountUrl string `pulumi:"storageAccountUrl"`
}

// Defaults sets the appropriate defaults for ListWebAppBackupStatusSecretsArgs
func (val *ListWebAppBackupStatusSecretsArgs) Defaults() *ListWebAppBackupStatusSecretsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackupSchedule = tmp.BackupSchedule.Defaults()

	return &tmp
}

// Backup description.
type ListWebAppBackupStatusSecretsResult struct {
	// Id of the backup.
	BackupId int `pulumi:"backupId"`
	// Name of the blob which contains data for this backup.
	BlobName string `pulumi:"blobName"`
	// Unique correlation identifier. Please use this along with the timestamp while communicating with Azure support.
	CorrelationId string `pulumi:"correlationId"`
	// Timestamp of the backup creation.
	Created string `pulumi:"created"`
	// List of databases included in the backup.
	Databases []DatabaseBackupSettingResponse `pulumi:"databases"`
	// Timestamp when this backup finished.
	FinishedTimeStamp string `pulumi:"finishedTimeStamp"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Timestamp of a last restore operation which used this backup.
	LastRestoreTimeStamp string `pulumi:"lastRestoreTimeStamp"`
	// Details regarding this backup. Might contain an error message.
	Log string `pulumi:"log"`
	// Resource Name.
	Name string `pulumi:"name"`
	// True if this backup has been created due to a schedule being triggered.
	Scheduled bool `pulumi:"scheduled"`
	// Size of the backup in bytes.
	SizeInBytes float64 `pulumi:"sizeInBytes"`
	// Backup status.
	Status string `pulumi:"status"`
	// SAS URL for the storage account container which contains this backup.
	StorageAccountUrl string `pulumi:"storageAccountUrl"`
	// Resource type.
	Type string `pulumi:"type"`
	// Size of the original web app which has been backed up.
	WebsiteSizeInBytes float64 `pulumi:"websiteSizeInBytes"`
}

func ListWebAppBackupStatusSecretsOutput(ctx *pulumi.Context, args ListWebAppBackupStatusSecretsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppBackupStatusSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppBackupStatusSecretsResult, error) {
			args := v.(ListWebAppBackupStatusSecretsArgs)
			r, err := ListWebAppBackupStatusSecrets(ctx, &args, opts...)
			var s ListWebAppBackupStatusSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppBackupStatusSecretsResultOutput)
}

type ListWebAppBackupStatusSecretsOutputArgs struct {
	// ID of backup.
	BackupId pulumi.StringInput `pulumi:"backupId"`
	// Name of the backup.
	BackupName pulumi.StringPtrInput `pulumi:"backupName"`
	// Schedule for the backup if it is executed periodically.
	BackupSchedule BackupSchedulePtrInput `pulumi:"backupSchedule"`
	// Databases included in the backup.
	Databases DatabaseBackupSettingArrayInput `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// Kind of resource.
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Name of web app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// SAS URL to the container.
	StorageAccountUrl pulumi.StringInput `pulumi:"storageAccountUrl"`
}

func (ListWebAppBackupStatusSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupStatusSecretsArgs)(nil)).Elem()
}

// Backup description.
type ListWebAppBackupStatusSecretsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppBackupStatusSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupStatusSecretsResult)(nil)).Elem()
}

func (o ListWebAppBackupStatusSecretsResultOutput) ToListWebAppBackupStatusSecretsResultOutput() ListWebAppBackupStatusSecretsResultOutput {
	return o
}

func (o ListWebAppBackupStatusSecretsResultOutput) ToListWebAppBackupStatusSecretsResultOutputWithContext(ctx context.Context) ListWebAppBackupStatusSecretsResultOutput {
	return o
}

func (o ListWebAppBackupStatusSecretsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppBackupStatusSecretsResult] {
	return pulumix.Output[ListWebAppBackupStatusSecretsResult]{
		OutputState: o.OutputState,
	}
}

// Id of the backup.
func (o ListWebAppBackupStatusSecretsResultOutput) BackupId() pulumi.IntOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) int { return v.BackupId }).(pulumi.IntOutput)
}

// Name of the blob which contains data for this backup.
func (o ListWebAppBackupStatusSecretsResultOutput) BlobName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.BlobName }).(pulumi.StringOutput)
}

// Unique correlation identifier. Please use this along with the timestamp while communicating with Azure support.
func (o ListWebAppBackupStatusSecretsResultOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.CorrelationId }).(pulumi.StringOutput)
}

// Timestamp of the backup creation.
func (o ListWebAppBackupStatusSecretsResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Created }).(pulumi.StringOutput)
}

// List of databases included in the backup.
func (o ListWebAppBackupStatusSecretsResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

// Timestamp when this backup finished.
func (o ListWebAppBackupStatusSecretsResultOutput) FinishedTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.FinishedTimeStamp }).(pulumi.StringOutput)
}

// Resource Id.
func (o ListWebAppBackupStatusSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppBackupStatusSecretsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Timestamp of a last restore operation which used this backup.
func (o ListWebAppBackupStatusSecretsResultOutput) LastRestoreTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.LastRestoreTimeStamp }).(pulumi.StringOutput)
}

// Details regarding this backup. Might contain an error message.
func (o ListWebAppBackupStatusSecretsResultOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Log }).(pulumi.StringOutput)
}

// Resource Name.
func (o ListWebAppBackupStatusSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

// True if this backup has been created due to a schedule being triggered.
func (o ListWebAppBackupStatusSecretsResultOutput) Scheduled() pulumi.BoolOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) bool { return v.Scheduled }).(pulumi.BoolOutput)
}

// Size of the backup in bytes.
func (o ListWebAppBackupStatusSecretsResultOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

// Backup status.
func (o ListWebAppBackupStatusSecretsResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Status }).(pulumi.StringOutput)
}

// SAS URL for the storage account container which contains this backup.
func (o ListWebAppBackupStatusSecretsResultOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

// Resource type.
func (o ListWebAppBackupStatusSecretsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) string { return v.Type }).(pulumi.StringOutput)
}

// Size of the original web app which has been backed up.
func (o ListWebAppBackupStatusSecretsResultOutput) WebsiteSizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v ListWebAppBackupStatusSecretsResult) float64 { return v.WebsiteSizeInBytes }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppBackupStatusSecretsResultOutput{})
}

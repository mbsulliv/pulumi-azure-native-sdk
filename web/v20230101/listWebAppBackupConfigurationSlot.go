// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the backup configuration of an app.
func ListWebAppBackupConfigurationSlot(ctx *pulumi.Context, args *ListWebAppBackupConfigurationSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupConfigurationSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppBackupConfigurationSlotResult
	err := ctx.Invoke("azure-native:web/v20230101:listWebAppBackupConfigurationSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListWebAppBackupConfigurationSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the backup configuration for the production slot.
	Slot string `pulumi:"slot"`
}

// Description of a backup which will be performed.
type ListWebAppBackupConfigurationSlotResult struct {
	// Name of the backup.
	BackupName *string `pulumi:"backupName"`
	// Schedule for the backup if it is executed periodically.
	BackupSchedule *BackupScheduleResponse `pulumi:"backupSchedule"`
	// Databases included in the backup.
	Databases []DatabaseBackupSettingResponse `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
	Enabled *bool `pulumi:"enabled"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// SAS URL to the container.
	StorageAccountUrl string `pulumi:"storageAccountUrl"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for ListWebAppBackupConfigurationSlotResult
func (val *ListWebAppBackupConfigurationSlotResult) Defaults() *ListWebAppBackupConfigurationSlotResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackupSchedule = tmp.BackupSchedule.Defaults()

	return &tmp
}

func ListWebAppBackupConfigurationSlotOutput(ctx *pulumi.Context, args ListWebAppBackupConfigurationSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppBackupConfigurationSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppBackupConfigurationSlotResult, error) {
			args := v.(ListWebAppBackupConfigurationSlotArgs)
			r, err := ListWebAppBackupConfigurationSlot(ctx, &args, opts...)
			var s ListWebAppBackupConfigurationSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppBackupConfigurationSlotResultOutput)
}

type ListWebAppBackupConfigurationSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the backup configuration for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppBackupConfigurationSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupConfigurationSlotArgs)(nil)).Elem()
}

// Description of a backup which will be performed.
type ListWebAppBackupConfigurationSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppBackupConfigurationSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupConfigurationSlotResult)(nil)).Elem()
}

func (o ListWebAppBackupConfigurationSlotResultOutput) ToListWebAppBackupConfigurationSlotResultOutput() ListWebAppBackupConfigurationSlotResultOutput {
	return o
}

func (o ListWebAppBackupConfigurationSlotResultOutput) ToListWebAppBackupConfigurationSlotResultOutputWithContext(ctx context.Context) ListWebAppBackupConfigurationSlotResultOutput {
	return o
}

// Name of the backup.
func (o ListWebAppBackupConfigurationSlotResultOutput) BackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) *string { return v.BackupName }).(pulumi.StringPtrOutput)
}

// Schedule for the backup if it is executed periodically.
func (o ListWebAppBackupConfigurationSlotResultOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) *BackupScheduleResponse { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

// Databases included in the backup.
func (o ListWebAppBackupConfigurationSlotResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
func (o ListWebAppBackupConfigurationSlotResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Resource Id.
func (o ListWebAppBackupConfigurationSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppBackupConfigurationSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppBackupConfigurationSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// SAS URL to the container.
func (o ListWebAppBackupConfigurationSlotResultOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

// Resource type.
func (o ListWebAppBackupConfigurationSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppBackupConfigurationSlotResultOutput{})
}

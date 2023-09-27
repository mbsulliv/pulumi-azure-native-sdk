// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Gets the backup configuration of an app.
func ListWebAppBackupConfiguration(ctx *pulumi.Context, args *ListWebAppBackupConfigurationArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppBackupConfigurationResult
	err := ctx.Invoke("azure-native:web/v20220901:listWebAppBackupConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListWebAppBackupConfigurationArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Description of a backup which will be performed.
type ListWebAppBackupConfigurationResult struct {
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

// Defaults sets the appropriate defaults for ListWebAppBackupConfigurationResult
func (val *ListWebAppBackupConfigurationResult) Defaults() *ListWebAppBackupConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackupSchedule = tmp.BackupSchedule.Defaults()

	return &tmp
}

func ListWebAppBackupConfigurationOutput(ctx *pulumi.Context, args ListWebAppBackupConfigurationOutputArgs, opts ...pulumi.InvokeOption) ListWebAppBackupConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppBackupConfigurationResult, error) {
			args := v.(ListWebAppBackupConfigurationArgs)
			r, err := ListWebAppBackupConfiguration(ctx, &args, opts...)
			var s ListWebAppBackupConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppBackupConfigurationResultOutput)
}

type ListWebAppBackupConfigurationOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppBackupConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupConfigurationArgs)(nil)).Elem()
}

// Description of a backup which will be performed.
type ListWebAppBackupConfigurationResultOutput struct{ *pulumi.OutputState }

func (ListWebAppBackupConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppBackupConfigurationResult)(nil)).Elem()
}

func (o ListWebAppBackupConfigurationResultOutput) ToListWebAppBackupConfigurationResultOutput() ListWebAppBackupConfigurationResultOutput {
	return o
}

func (o ListWebAppBackupConfigurationResultOutput) ToListWebAppBackupConfigurationResultOutputWithContext(ctx context.Context) ListWebAppBackupConfigurationResultOutput {
	return o
}

func (o ListWebAppBackupConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppBackupConfigurationResult] {
	return pulumix.Output[ListWebAppBackupConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Name of the backup.
func (o ListWebAppBackupConfigurationResultOutput) BackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) *string { return v.BackupName }).(pulumi.StringPtrOutput)
}

// Schedule for the backup if it is executed periodically.
func (o ListWebAppBackupConfigurationResultOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) *BackupScheduleResponse { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

// Databases included in the backup.
func (o ListWebAppBackupConfigurationResultOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
func (o ListWebAppBackupConfigurationResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Resource Id.
func (o ListWebAppBackupConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppBackupConfigurationResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppBackupConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// SAS URL to the container.
func (o ListWebAppBackupConfigurationResultOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

// Resource type.
func (o ListWebAppBackupConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppBackupConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppBackupConfigurationResultOutput{})
}

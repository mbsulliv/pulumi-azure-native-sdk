// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a particular backup Policy
// Azure REST API version: 2022-11-01.
//
// Other available API versions: 2021-04-01, 2021-04-01-preview, 2022-11-01-preview, 2023-05-01.
func LookupBackupPolicy(ctx *pulumi.Context, args *LookupBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupPolicyResult
	err := ctx.Invoke("azure-native:netapp:getBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPolicyArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Backup policy Name which uniquely identify backup policy.
	BackupPolicyName string `pulumi:"backupPolicyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Backup policy information
type LookupBackupPolicyResult struct {
	// Backup Policy Resource ID
	BackupPolicyId string `pulumi:"backupPolicyId"`
	// Daily backups count to keep
	DailyBackupsToKeep *int `pulumi:"dailyBackupsToKeep"`
	// The property to decide policy is enabled or not
	Enabled *bool `pulumi:"enabled"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Monthly backups count to keep
	MonthlyBackupsToKeep *int `pulumi:"monthlyBackupsToKeep"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// A list of volumes assigned to this policy
	VolumeBackups []VolumeBackupsResponse `pulumi:"volumeBackups"`
	// Volumes using current backup policy
	VolumesAssigned int `pulumi:"volumesAssigned"`
	// Weekly backups count to keep
	WeeklyBackupsToKeep *int `pulumi:"weeklyBackupsToKeep"`
}

func LookupBackupPolicyOutput(ctx *pulumi.Context, args LookupBackupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBackupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupPolicyResult, error) {
			args := v.(LookupBackupPolicyArgs)
			r, err := LookupBackupPolicy(ctx, &args, opts...)
			var s LookupBackupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupPolicyResultOutput)
}

type LookupBackupPolicyOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Backup policy Name which uniquely identify backup policy.
	BackupPolicyName pulumi.StringInput `pulumi:"backupPolicyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBackupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPolicyArgs)(nil)).Elem()
}

// Backup policy information
type LookupBackupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBackupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPolicyResult)(nil)).Elem()
}

func (o LookupBackupPolicyResultOutput) ToLookupBackupPolicyResultOutput() LookupBackupPolicyResultOutput {
	return o
}

func (o LookupBackupPolicyResultOutput) ToLookupBackupPolicyResultOutputWithContext(ctx context.Context) LookupBackupPolicyResultOutput {
	return o
}

// Backup Policy Resource ID
func (o LookupBackupPolicyResultOutput) BackupPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.BackupPolicyId }).(pulumi.StringOutput)
}

// Daily backups count to keep
func (o LookupBackupPolicyResultOutput) DailyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *int { return v.DailyBackupsToKeep }).(pulumi.IntPtrOutput)
}

// The property to decide policy is enabled or not
func (o LookupBackupPolicyResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupBackupPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBackupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupBackupPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// Monthly backups count to keep
func (o LookupBackupPolicyResultOutput) MonthlyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *int { return v.MonthlyBackupsToKeep }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o LookupBackupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupBackupPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBackupPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupBackupPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBackupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of volumes assigned to this policy
func (o LookupBackupPolicyResultOutput) VolumeBackups() VolumeBackupsResponseArrayOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) []VolumeBackupsResponse { return v.VolumeBackups }).(VolumeBackupsResponseArrayOutput)
}

// Volumes using current backup policy
func (o LookupBackupPolicyResultOutput) VolumesAssigned() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) int { return v.VolumesAssigned }).(pulumi.IntOutput)
}

// Weekly backups count to keep
func (o LookupBackupPolicyResultOutput) WeeklyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) *int { return v.WeeklyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupPolicyResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Backup policy information
type BackupPolicy struct {
	pulumi.CustomResourceState

	// Backup Policy Resource ID
	BackupPolicyId pulumi.StringOutput `pulumi:"backupPolicyId"`
	// Daily backups count to keep
	DailyBackupsToKeep pulumi.IntPtrOutput `pulumi:"dailyBackupsToKeep"`
	// The property to decide policy is enabled or not
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Monthly backups count to keep
	MonthlyBackupsToKeep pulumi.IntPtrOutput `pulumi:"monthlyBackupsToKeep"`
	// Name of backup policy
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of volumes assigned to this policy
	VolumeBackups VolumeBackupsResponseArrayOutput `pulumi:"volumeBackups"`
	// Volumes using current backup policy
	VolumesAssigned pulumi.IntPtrOutput `pulumi:"volumesAssigned"`
	// Weekly backups count to keep
	WeeklyBackupsToKeep pulumi.IntPtrOutput `pulumi:"weeklyBackupsToKeep"`
	// Yearly backups count to keep
	YearlyBackupsToKeep pulumi.IntPtrOutput `pulumi:"yearlyBackupsToKeep"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230501:BackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BackupPolicy
	err := ctx.RegisterResource("azure-native:netapp/v20210401:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPolicy gets an existing BackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("azure-native:netapp/v20210401:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
}

type BackupPolicyState struct {
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Backup policy Name which uniquely identify backup policy.
	BackupPolicyName *string `pulumi:"backupPolicyName"`
	// Daily backups count to keep
	DailyBackupsToKeep *int `pulumi:"dailyBackupsToKeep"`
	// The property to decide policy is enabled or not
	Enabled *bool `pulumi:"enabled"`
	// Resource location
	Location *string `pulumi:"location"`
	// Monthly backups count to keep
	MonthlyBackupsToKeep *int `pulumi:"monthlyBackupsToKeep"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// A list of volumes assigned to this policy
	VolumeBackups []VolumeBackups `pulumi:"volumeBackups"`
	// Volumes using current backup policy
	VolumesAssigned *int `pulumi:"volumesAssigned"`
	// Weekly backups count to keep
	WeeklyBackupsToKeep *int `pulumi:"weeklyBackupsToKeep"`
	// Yearly backups count to keep
	YearlyBackupsToKeep *int `pulumi:"yearlyBackupsToKeep"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// Backup policy Name which uniquely identify backup policy.
	BackupPolicyName pulumi.StringPtrInput
	// Daily backups count to keep
	DailyBackupsToKeep pulumi.IntPtrInput
	// The property to decide policy is enabled or not
	Enabled pulumi.BoolPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Monthly backups count to keep
	MonthlyBackupsToKeep pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// A list of volumes assigned to this policy
	VolumeBackups VolumeBackupsArrayInput
	// Volumes using current backup policy
	VolumesAssigned pulumi.IntPtrInput
	// Weekly backups count to keep
	WeeklyBackupsToKeep pulumi.IntPtrInput
	// Yearly backups count to keep
	YearlyBackupsToKeep pulumi.IntPtrInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}

type BackupPolicyInput interface {
	pulumi.Input

	ToBackupPolicyOutput() BackupPolicyOutput
	ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput
}

func (*BackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

type BackupPolicyOutput struct{ *pulumi.OutputState }

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

// Backup Policy Resource ID
func (o BackupPolicyOutput) BackupPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.BackupPolicyId }).(pulumi.StringOutput)
}

// Daily backups count to keep
func (o BackupPolicyOutput) DailyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.DailyBackupsToKeep }).(pulumi.IntPtrOutput)
}

// The property to decide policy is enabled or not
func (o BackupPolicyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o BackupPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location
func (o BackupPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Monthly backups count to keep
func (o BackupPolicyOutput) MonthlyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.MonthlyBackupsToKeep }).(pulumi.IntPtrOutput)
}

// Name of backup policy
func (o BackupPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o BackupPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o BackupPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BackupPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o BackupPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o BackupPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of volumes assigned to this policy
func (o BackupPolicyOutput) VolumeBackups() VolumeBackupsResponseArrayOutput {
	return o.ApplyT(func(v *BackupPolicy) VolumeBackupsResponseArrayOutput { return v.VolumeBackups }).(VolumeBackupsResponseArrayOutput)
}

// Volumes using current backup policy
func (o BackupPolicyOutput) VolumesAssigned() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.VolumesAssigned }).(pulumi.IntPtrOutput)
}

// Weekly backups count to keep
func (o BackupPolicyOutput) WeeklyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.WeeklyBackupsToKeep }).(pulumi.IntPtrOutput)
}

// Yearly backups count to keep
func (o BackupPolicyOutput) YearlyBackupsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.YearlyBackupsToKeep }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupPolicyOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The backup policy.
// Azure REST API version: 2017-06-01. Prior API version in Azure Native 1.x: 2017-06-01
type BackupPolicy struct {
	pulumi.CustomResourceState

	// The backup policy creation type. Indicates whether this was created through SaaS or through StorSimple Snapshot Manager.
	BackupPolicyCreationType pulumi.StringOutput `pulumi:"backupPolicyCreationType"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The time of the last backup for the backup policy.
	LastBackupTime pulumi.StringOutput `pulumi:"lastBackupTime"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The time of the next backup for the backup policy.
	NextBackupTime pulumi.StringOutput `pulumi:"nextBackupTime"`
	// Indicates whether at least one of the schedules in the backup policy is active or not.
	ScheduledBackupStatus pulumi.StringOutput `pulumi:"scheduledBackupStatus"`
	// The count of schedules the backup policy contains.
	SchedulesCount pulumi.Float64Output `pulumi:"schedulesCount"`
	// If the backup policy was created by StorSimple Snapshot Manager, then this field indicates the hostname of the StorSimple Snapshot Manager.
	SsmHostName pulumi.StringOutput `pulumi:"ssmHostName"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// The path IDs of the volumes which are part of the backup policy.
	VolumeIds pulumi.StringArrayOutput `pulumi:"volumeIds"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeIds == nil {
		return nil, errors.New("invalid value for required argument 'VolumeIds'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:BackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BackupPolicy
	err := ctx.RegisterResource("azure-native:storsimple:BackupPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:storsimple:BackupPolicy", name, id, state, &resource, opts...)
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
	// The name of the backup policy to be created/updated.
	BackupPolicyName *string `pulumi:"backupPolicyName"`
	// The device name
	DeviceName string `pulumi:"deviceName"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *Kind `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The path IDs of the volumes which are part of the backup policy.
	VolumeIds []string `pulumi:"volumeIds"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// The name of the backup policy to be created/updated.
	BackupPolicyName pulumi.StringPtrInput
	// The device name
	DeviceName pulumi.StringInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind KindPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The path IDs of the volumes which are part of the backup policy.
	VolumeIds pulumi.StringArrayInput
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

func (i *BackupPolicy) ToOutput(ctx context.Context) pulumix.Output[*BackupPolicy] {
	return pulumix.Output[*BackupPolicy]{
		OutputState: i.ToBackupPolicyOutputWithContext(ctx).OutputState,
	}
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

func (o BackupPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*BackupPolicy] {
	return pulumix.Output[*BackupPolicy]{
		OutputState: o.OutputState,
	}
}

// The backup policy creation type. Indicates whether this was created through SaaS or through StorSimple Snapshot Manager.
func (o BackupPolicyOutput) BackupPolicyCreationType() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.BackupPolicyCreationType }).(pulumi.StringOutput)
}

// The Kind of the object. Currently only Series8000 is supported
func (o BackupPolicyOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The time of the last backup for the backup policy.
func (o BackupPolicyOutput) LastBackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.LastBackupTime }).(pulumi.StringOutput)
}

// The name of the object.
func (o BackupPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The time of the next backup for the backup policy.
func (o BackupPolicyOutput) NextBackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.NextBackupTime }).(pulumi.StringOutput)
}

// Indicates whether at least one of the schedules in the backup policy is active or not.
func (o BackupPolicyOutput) ScheduledBackupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.ScheduledBackupStatus }).(pulumi.StringOutput)
}

// The count of schedules the backup policy contains.
func (o BackupPolicyOutput) SchedulesCount() pulumi.Float64Output {
	return o.ApplyT(func(v *BackupPolicy) pulumi.Float64Output { return v.SchedulesCount }).(pulumi.Float64Output)
}

// If the backup policy was created by StorSimple Snapshot Manager, then this field indicates the hostname of the StorSimple Snapshot Manager.
func (o BackupPolicyOutput) SsmHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.SsmHostName }).(pulumi.StringOutput)
}

// The hierarchical type of the object.
func (o BackupPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The path IDs of the volumes which are part of the backup policy.
func (o BackupPolicyOutput) VolumeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringArrayOutput { return v.VolumeIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupPolicyOutput{})
}

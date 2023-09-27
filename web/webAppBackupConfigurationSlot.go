// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description of a backup which will be performed.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01
type WebAppBackupConfigurationSlot struct {
	pulumi.CustomResourceState

	// Name of the backup.
	BackupName pulumi.StringPtrOutput `pulumi:"backupName"`
	// Schedule for the backup if it is executed periodically.
	BackupSchedule BackupScheduleResponsePtrOutput `pulumi:"backupSchedule"`
	// Databases included in the backup.
	Databases DatabaseBackupSettingResponseArrayOutput `pulumi:"databases"`
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SAS URL to the container.
	StorageAccountUrl pulumi.StringOutput `pulumi:"storageAccountUrl"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppBackupConfigurationSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppBackupConfigurationSlot(ctx *pulumi.Context,
	name string, args *WebAppBackupConfigurationSlotArgs, opts ...pulumi.ResourceOption) (*WebAppBackupConfigurationSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	if args.StorageAccountUrl == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountUrl'")
	}
	if args.BackupSchedule != nil {
		args.BackupSchedule = args.BackupSchedule.ToBackupSchedulePtrOutput().ApplyT(func(v *BackupSchedule) *BackupSchedule { return v.Defaults() }).(BackupSchedulePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppBackupConfigurationSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppBackupConfigurationSlot
	err := ctx.RegisterResource("azure-native:web:WebAppBackupConfigurationSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppBackupConfigurationSlot gets an existing WebAppBackupConfigurationSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppBackupConfigurationSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppBackupConfigurationSlotState, opts ...pulumi.ResourceOption) (*WebAppBackupConfigurationSlot, error) {
	var resource WebAppBackupConfigurationSlot
	err := ctx.ReadResource("azure-native:web:WebAppBackupConfigurationSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppBackupConfigurationSlot resources.
type webAppBackupConfigurationSlotState struct {
}

type WebAppBackupConfigurationSlotState struct {
}

func (WebAppBackupConfigurationSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppBackupConfigurationSlotState)(nil)).Elem()
}

type webAppBackupConfigurationSlotArgs struct {
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
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the backup configuration for the production slot.
	Slot string `pulumi:"slot"`
	// SAS URL to the container.
	StorageAccountUrl string `pulumi:"storageAccountUrl"`
}

// The set of arguments for constructing a WebAppBackupConfigurationSlot resource.
type WebAppBackupConfigurationSlotArgs struct {
	// Name of the backup.
	BackupName pulumi.StringPtrInput
	// Schedule for the backup if it is executed periodically.
	BackupSchedule BackupSchedulePtrInput
	// Databases included in the backup.
	Databases DatabaseBackupSettingArrayInput
	// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
	Enabled pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will update the backup configuration for the production slot.
	Slot pulumi.StringInput
	// SAS URL to the container.
	StorageAccountUrl pulumi.StringInput
}

func (WebAppBackupConfigurationSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppBackupConfigurationSlotArgs)(nil)).Elem()
}

type WebAppBackupConfigurationSlotInput interface {
	pulumi.Input

	ToWebAppBackupConfigurationSlotOutput() WebAppBackupConfigurationSlotOutput
	ToWebAppBackupConfigurationSlotOutputWithContext(ctx context.Context) WebAppBackupConfigurationSlotOutput
}

func (*WebAppBackupConfigurationSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppBackupConfigurationSlot)(nil)).Elem()
}

func (i *WebAppBackupConfigurationSlot) ToWebAppBackupConfigurationSlotOutput() WebAppBackupConfigurationSlotOutput {
	return i.ToWebAppBackupConfigurationSlotOutputWithContext(context.Background())
}

func (i *WebAppBackupConfigurationSlot) ToWebAppBackupConfigurationSlotOutputWithContext(ctx context.Context) WebAppBackupConfigurationSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppBackupConfigurationSlotOutput)
}

func (i *WebAppBackupConfigurationSlot) ToOutput(ctx context.Context) pulumix.Output[*WebAppBackupConfigurationSlot] {
	return pulumix.Output[*WebAppBackupConfigurationSlot]{
		OutputState: i.ToWebAppBackupConfigurationSlotOutputWithContext(ctx).OutputState,
	}
}

type WebAppBackupConfigurationSlotOutput struct{ *pulumi.OutputState }

func (WebAppBackupConfigurationSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppBackupConfigurationSlot)(nil)).Elem()
}

func (o WebAppBackupConfigurationSlotOutput) ToWebAppBackupConfigurationSlotOutput() WebAppBackupConfigurationSlotOutput {
	return o
}

func (o WebAppBackupConfigurationSlotOutput) ToWebAppBackupConfigurationSlotOutputWithContext(ctx context.Context) WebAppBackupConfigurationSlotOutput {
	return o
}

func (o WebAppBackupConfigurationSlotOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppBackupConfigurationSlot] {
	return pulumix.Output[*WebAppBackupConfigurationSlot]{
		OutputState: o.OutputState,
	}
}

// Name of the backup.
func (o WebAppBackupConfigurationSlotOutput) BackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringPtrOutput { return v.BackupName }).(pulumi.StringPtrOutput)
}

// Schedule for the backup if it is executed periodically.
func (o WebAppBackupConfigurationSlotOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) BackupScheduleResponsePtrOutput { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

// Databases included in the backup.
func (o WebAppBackupConfigurationSlotOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) DatabaseBackupSettingResponseArrayOutput { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
func (o WebAppBackupConfigurationSlotOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o WebAppBackupConfigurationSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppBackupConfigurationSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SAS URL to the container.
func (o WebAppBackupConfigurationSlotOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringOutput { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

// Resource type.
func (o WebAppBackupConfigurationSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppBackupConfigurationSlotOutput{})
}

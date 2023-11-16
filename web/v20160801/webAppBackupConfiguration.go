// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of a backup which will be performed.
type WebAppBackupConfiguration struct {
	pulumi.CustomResourceState

	// Name of the backup.
	BackupRequestName pulumi.StringOutput `pulumi:"backupRequestName"`
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

// NewWebAppBackupConfiguration registers a new resource with the given unique name, arguments, and options.
func NewWebAppBackupConfiguration(ctx *pulumi.Context,
	name string, args *WebAppBackupConfigurationArgs, opts ...pulumi.ResourceOption) (*WebAppBackupConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupRequestName == nil {
		return nil, errors.New("invalid value for required argument 'BackupRequestName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountUrl == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountUrl'")
	}
	if args.BackupSchedule != nil {
		args.BackupSchedule = args.BackupSchedule.ToBackupSchedulePtrOutput().ApplyT(func(v *BackupSchedule) *BackupSchedule { return v.Defaults() }).(BackupSchedulePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppBackupConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppBackupConfiguration
	err := ctx.RegisterResource("azure-native:web/v20160801:WebAppBackupConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppBackupConfiguration gets an existing WebAppBackupConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppBackupConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppBackupConfigurationState, opts ...pulumi.ResourceOption) (*WebAppBackupConfiguration, error) {
	var resource WebAppBackupConfiguration
	err := ctx.ReadResource("azure-native:web/v20160801:WebAppBackupConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppBackupConfiguration resources.
type webAppBackupConfigurationState struct {
}

type WebAppBackupConfigurationState struct {
}

func (WebAppBackupConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppBackupConfigurationState)(nil)).Elem()
}

type webAppBackupConfigurationArgs struct {
	// Name of the backup.
	BackupRequestName string `pulumi:"backupRequestName"`
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
	// SAS URL to the container.
	StorageAccountUrl string `pulumi:"storageAccountUrl"`
	// Type of the backup.
	Type *BackupRestoreOperationType `pulumi:"type"`
}

// The set of arguments for constructing a WebAppBackupConfiguration resource.
type WebAppBackupConfigurationArgs struct {
	// Name of the backup.
	BackupRequestName pulumi.StringInput
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
	// SAS URL to the container.
	StorageAccountUrl pulumi.StringInput
	// Type of the backup.
	Type BackupRestoreOperationTypePtrInput
}

func (WebAppBackupConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppBackupConfigurationArgs)(nil)).Elem()
}

type WebAppBackupConfigurationInput interface {
	pulumi.Input

	ToWebAppBackupConfigurationOutput() WebAppBackupConfigurationOutput
	ToWebAppBackupConfigurationOutputWithContext(ctx context.Context) WebAppBackupConfigurationOutput
}

func (*WebAppBackupConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppBackupConfiguration)(nil)).Elem()
}

func (i *WebAppBackupConfiguration) ToWebAppBackupConfigurationOutput() WebAppBackupConfigurationOutput {
	return i.ToWebAppBackupConfigurationOutputWithContext(context.Background())
}

func (i *WebAppBackupConfiguration) ToWebAppBackupConfigurationOutputWithContext(ctx context.Context) WebAppBackupConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppBackupConfigurationOutput)
}

type WebAppBackupConfigurationOutput struct{ *pulumi.OutputState }

func (WebAppBackupConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppBackupConfiguration)(nil)).Elem()
}

func (o WebAppBackupConfigurationOutput) ToWebAppBackupConfigurationOutput() WebAppBackupConfigurationOutput {
	return o
}

func (o WebAppBackupConfigurationOutput) ToWebAppBackupConfigurationOutputWithContext(ctx context.Context) WebAppBackupConfigurationOutput {
	return o
}

// Name of the backup.
func (o WebAppBackupConfigurationOutput) BackupRequestName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringOutput { return v.BackupRequestName }).(pulumi.StringOutput)
}

// Schedule for the backup if it is executed periodically.
func (o WebAppBackupConfigurationOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) BackupScheduleResponsePtrOutput { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

// Databases included in the backup.
func (o WebAppBackupConfigurationOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) DatabaseBackupSettingResponseArrayOutput { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
func (o WebAppBackupConfigurationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o WebAppBackupConfigurationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppBackupConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SAS URL to the container.
func (o WebAppBackupConfigurationOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringOutput { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

// Resource type.
func (o WebAppBackupConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppBackupConfigurationOutput{})
}

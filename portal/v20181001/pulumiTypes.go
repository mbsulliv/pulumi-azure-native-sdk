// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Cloud shell properties for creating a console.
type ConsoleCreateProperties struct {
	// The operating system type of the cloud shell.
	OsType string `pulumi:"osType"`
	// Provisioning state of the console.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Uri of the console.
	Uri *string `pulumi:"uri"`
}

// ConsoleCreatePropertiesInput is an input type that accepts ConsoleCreatePropertiesArgs and ConsoleCreatePropertiesOutput values.
// You can construct a concrete instance of `ConsoleCreatePropertiesInput` via:
//
//	ConsoleCreatePropertiesArgs{...}
type ConsoleCreatePropertiesInput interface {
	pulumi.Input

	ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput
	ToConsoleCreatePropertiesOutputWithContext(context.Context) ConsoleCreatePropertiesOutput
}

// Cloud shell properties for creating a console.
type ConsoleCreatePropertiesArgs struct {
	// The operating system type of the cloud shell.
	OsType pulumi.StringInput `pulumi:"osType"`
	// Provisioning state of the console.
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	// Uri of the console.
	Uri pulumi.StringPtrInput `pulumi:"uri"`
}

func (ConsoleCreatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleCreateProperties)(nil)).Elem()
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput {
	return i.ToConsoleCreatePropertiesOutputWithContext(context.Background())
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesOutputWithContext(ctx context.Context) ConsoleCreatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleCreatePropertiesOutput)
}

// Cloud shell properties for creating a console.
type ConsoleCreatePropertiesOutput struct{ *pulumi.OutputState }

func (ConsoleCreatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleCreateProperties)(nil)).Elem()
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput {
	return o
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesOutputWithContext(ctx context.Context) ConsoleCreatePropertiesOutput {
	return o
}

// The operating system type of the cloud shell.
func (o ConsoleCreatePropertiesOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) string { return v.OsType }).(pulumi.StringOutput)
}

// Provisioning state of the console.
func (o ConsoleCreatePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Uri of the console.
func (o ConsoleCreatePropertiesOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

// Cloud shell console properties.
type ConsolePropertiesResponse struct {
	// The operating system type of the cloud shell.
	OsType string `pulumi:"osType"`
	// Provisioning state of the console.
	ProvisioningState string `pulumi:"provisioningState"`
	// Uri of the console.
	Uri string `pulumi:"uri"`
}

// Cloud shell console properties.
type ConsolePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConsolePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsolePropertiesResponse)(nil)).Elem()
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponseOutput() ConsolePropertiesResponseOutput {
	return o
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponseOutputWithContext(ctx context.Context) ConsolePropertiesResponseOutput {
	return o
}

// The operating system type of the cloud shell.
func (o ConsolePropertiesResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.OsType }).(pulumi.StringOutput)
}

// Provisioning state of the console.
func (o ConsolePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Uri of the console.
func (o ConsolePropertiesResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.Uri }).(pulumi.StringOutput)
}

// The storage profile of the user settings.
type StorageProfile struct {
	// Size of file share
	DiskSizeInGB *int `pulumi:"diskSizeInGB"`
	// Name of the mounted file share. 63 characters or less, lowercase alphabet, numbers, and -
	FileShareName *string `pulumi:"fileShareName"`
	// Full resource ID of storage account.
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}

// StorageProfileInput is an input type that accepts StorageProfileArgs and StorageProfileOutput values.
// You can construct a concrete instance of `StorageProfileInput` via:
//
//	StorageProfileArgs{...}
type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

// The storage profile of the user settings.
type StorageProfileArgs struct {
	// Size of file share
	DiskSizeInGB pulumi.IntPtrInput `pulumi:"diskSizeInGB"`
	// Name of the mounted file share. 63 characters or less, lowercase alphabet, numbers, and -
	FileShareName pulumi.StringPtrInput `pulumi:"fileShareName"`
	// Full resource ID of storage account.
	StorageAccountResourceId pulumi.StringPtrInput `pulumi:"storageAccountResourceId"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

// The storage profile of the user settings.
type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

// Size of file share
func (o StorageProfileOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.DiskSizeInGB }).(pulumi.IntPtrOutput)
}

// Name of the mounted file share. 63 characters or less, lowercase alphabet, numbers, and -
func (o StorageProfileOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.FileShareName }).(pulumi.StringPtrOutput)
}

// Full resource ID of storage account.
func (o StorageProfileOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

// The storage profile of the user settings.
type StorageProfileResponse struct {
	// Size of file share
	DiskSizeInGB *int `pulumi:"diskSizeInGB"`
	// Name of the mounted file share. 63 characters or less, lowercase alphabet, numbers, and -
	FileShareName *string `pulumi:"fileShareName"`
	// Full resource ID of storage account.
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}

// The storage profile of the user settings.
type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

// Size of file share
func (o StorageProfileResponseOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.DiskSizeInGB }).(pulumi.IntPtrOutput)
}

// Name of the mounted file share. 63 characters or less, lowercase alphabet, numbers, and -
func (o StorageProfileResponseOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.FileShareName }).(pulumi.StringPtrOutput)
}

// Full resource ID of storage account.
func (o StorageProfileResponseOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

// Settings for terminal appearance.
type TerminalSettings struct {
	// Size of terminal font.
	FontSize *string `pulumi:"fontSize"`
	// Style of terminal font.
	FontStyle *string `pulumi:"fontStyle"`
}

// TerminalSettingsInput is an input type that accepts TerminalSettingsArgs and TerminalSettingsOutput values.
// You can construct a concrete instance of `TerminalSettingsInput` via:
//
//	TerminalSettingsArgs{...}
type TerminalSettingsInput interface {
	pulumi.Input

	ToTerminalSettingsOutput() TerminalSettingsOutput
	ToTerminalSettingsOutputWithContext(context.Context) TerminalSettingsOutput
}

// Settings for terminal appearance.
type TerminalSettingsArgs struct {
	// Size of terminal font.
	FontSize pulumi.StringPtrInput `pulumi:"fontSize"`
	// Style of terminal font.
	FontStyle pulumi.StringPtrInput `pulumi:"fontStyle"`
}

func (TerminalSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettings)(nil)).Elem()
}

func (i TerminalSettingsArgs) ToTerminalSettingsOutput() TerminalSettingsOutput {
	return i.ToTerminalSettingsOutputWithContext(context.Background())
}

func (i TerminalSettingsArgs) ToTerminalSettingsOutputWithContext(ctx context.Context) TerminalSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsOutput)
}

// Settings for terminal appearance.
type TerminalSettingsOutput struct{ *pulumi.OutputState }

func (TerminalSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettings)(nil)).Elem()
}

func (o TerminalSettingsOutput) ToTerminalSettingsOutput() TerminalSettingsOutput {
	return o
}

func (o TerminalSettingsOutput) ToTerminalSettingsOutputWithContext(ctx context.Context) TerminalSettingsOutput {
	return o
}

// Size of terminal font.
func (o TerminalSettingsOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettings) *string { return v.FontSize }).(pulumi.StringPtrOutput)
}

// Style of terminal font.
func (o TerminalSettingsOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettings) *string { return v.FontStyle }).(pulumi.StringPtrOutput)
}

// Settings for terminal appearance.
type TerminalSettingsResponse struct {
	// Size of terminal font.
	FontSize *string `pulumi:"fontSize"`
	// Style of terminal font.
	FontStyle *string `pulumi:"fontStyle"`
}

// Settings for terminal appearance.
type TerminalSettingsResponseOutput struct{ *pulumi.OutputState }

func (TerminalSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettingsResponse)(nil)).Elem()
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponseOutput() TerminalSettingsResponseOutput {
	return o
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponseOutputWithContext(ctx context.Context) TerminalSettingsResponseOutput {
	return o
}

// Size of terminal font.
func (o TerminalSettingsResponseOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettingsResponse) *string { return v.FontSize }).(pulumi.StringPtrOutput)
}

// Style of terminal font.
func (o TerminalSettingsResponseOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettingsResponse) *string { return v.FontStyle }).(pulumi.StringPtrOutput)
}

// The cloud shell user settings properties.
type UserProperties struct {
	// The preferred location of the cloud shell.
	PreferredLocation string `pulumi:"preferredLocation"`
	// The operating system type of the cloud shell. Deprecated, use preferredShellType.
	PreferredOsType string `pulumi:"preferredOsType"`
	// The shell type of the cloud shell.
	PreferredShellType string `pulumi:"preferredShellType"`
	// The storage profile of the user settings.
	StorageProfile StorageProfile `pulumi:"storageProfile"`
	// Settings for terminal appearance.
	TerminalSettings TerminalSettings `pulumi:"terminalSettings"`
}

// UserPropertiesInput is an input type that accepts UserPropertiesArgs and UserPropertiesOutput values.
// You can construct a concrete instance of `UserPropertiesInput` via:
//
//	UserPropertiesArgs{...}
type UserPropertiesInput interface {
	pulumi.Input

	ToUserPropertiesOutput() UserPropertiesOutput
	ToUserPropertiesOutputWithContext(context.Context) UserPropertiesOutput
}

// The cloud shell user settings properties.
type UserPropertiesArgs struct {
	// The preferred location of the cloud shell.
	PreferredLocation pulumi.StringInput `pulumi:"preferredLocation"`
	// The operating system type of the cloud shell. Deprecated, use preferredShellType.
	PreferredOsType pulumi.StringInput `pulumi:"preferredOsType"`
	// The shell type of the cloud shell.
	PreferredShellType pulumi.StringInput `pulumi:"preferredShellType"`
	// The storage profile of the user settings.
	StorageProfile StorageProfileInput `pulumi:"storageProfile"`
	// Settings for terminal appearance.
	TerminalSettings TerminalSettingsInput `pulumi:"terminalSettings"`
}

func (UserPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProperties)(nil)).Elem()
}

func (i UserPropertiesArgs) ToUserPropertiesOutput() UserPropertiesOutput {
	return i.ToUserPropertiesOutputWithContext(context.Background())
}

func (i UserPropertiesArgs) ToUserPropertiesOutputWithContext(ctx context.Context) UserPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesOutput)
}

// The cloud shell user settings properties.
type UserPropertiesOutput struct{ *pulumi.OutputState }

func (UserPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProperties)(nil)).Elem()
}

func (o UserPropertiesOutput) ToUserPropertiesOutput() UserPropertiesOutput {
	return o
}

func (o UserPropertiesOutput) ToUserPropertiesOutputWithContext(ctx context.Context) UserPropertiesOutput {
	return o
}

// The preferred location of the cloud shell.
func (o UserPropertiesOutput) PreferredLocation() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredLocation }).(pulumi.StringOutput)
}

// The operating system type of the cloud shell. Deprecated, use preferredShellType.
func (o UserPropertiesOutput) PreferredOsType() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredOsType }).(pulumi.StringOutput)
}

// The shell type of the cloud shell.
func (o UserPropertiesOutput) PreferredShellType() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredShellType }).(pulumi.StringOutput)
}

// The storage profile of the user settings.
func (o UserPropertiesOutput) StorageProfile() StorageProfileOutput {
	return o.ApplyT(func(v UserProperties) StorageProfile { return v.StorageProfile }).(StorageProfileOutput)
}

// Settings for terminal appearance.
func (o UserPropertiesOutput) TerminalSettings() TerminalSettingsOutput {
	return o.ApplyT(func(v UserProperties) TerminalSettings { return v.TerminalSettings }).(TerminalSettingsOutput)
}

// The cloud shell user settings properties.
type UserPropertiesResponse struct {
	// The preferred location of the cloud shell.
	PreferredLocation string `pulumi:"preferredLocation"`
	// The operating system type of the cloud shell. Deprecated, use preferredShellType.
	PreferredOsType string `pulumi:"preferredOsType"`
	// The shell type of the cloud shell.
	PreferredShellType string `pulumi:"preferredShellType"`
	// The storage profile of the user settings.
	StorageProfile StorageProfileResponse `pulumi:"storageProfile"`
	// Settings for terminal appearance.
	TerminalSettings TerminalSettingsResponse `pulumi:"terminalSettings"`
}

// The cloud shell user settings properties.
type UserPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPropertiesResponse)(nil)).Elem()
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponseOutput() UserPropertiesResponseOutput {
	return o
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponseOutputWithContext(ctx context.Context) UserPropertiesResponseOutput {
	return o
}

// The preferred location of the cloud shell.
func (o UserPropertiesResponseOutput) PreferredLocation() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredLocation }).(pulumi.StringOutput)
}

// The operating system type of the cloud shell. Deprecated, use preferredShellType.
func (o UserPropertiesResponseOutput) PreferredOsType() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredOsType }).(pulumi.StringOutput)
}

// The shell type of the cloud shell.
func (o UserPropertiesResponseOutput) PreferredShellType() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredShellType }).(pulumi.StringOutput)
}

// The storage profile of the user settings.
func (o UserPropertiesResponseOutput) StorageProfile() StorageProfileResponseOutput {
	return o.ApplyT(func(v UserPropertiesResponse) StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponseOutput)
}

// Settings for terminal appearance.
func (o UserPropertiesResponseOutput) TerminalSettings() TerminalSettingsResponseOutput {
	return o.ApplyT(func(v UserPropertiesResponse) TerminalSettingsResponse { return v.TerminalSettings }).(TerminalSettingsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ConsoleCreatePropertiesOutput{})
	pulumi.RegisterOutputType(ConsolePropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(TerminalSettingsOutput{})
	pulumi.RegisterOutputType(TerminalSettingsResponseOutput{})
	pulumi.RegisterOutputType(UserPropertiesOutput{})
	pulumi.RegisterOutputType(UserPropertiesResponseOutput{})
}

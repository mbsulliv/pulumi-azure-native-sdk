// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package intune

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Android Policy entity for Intune MAM.
// Azure REST API version: 2015-01-14-preview. Prior API version in Azure Native 1.x: 2015-01-14-preview.
//
// Other available API versions: 2015-01-14-privatepreview.
type AndroidMAMPolicyByName struct {
	pulumi.CustomResourceState

	AccessRecheckOfflineTimeout pulumi.StringPtrOutput `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  pulumi.StringPtrOutput `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         pulumi.StringPtrOutput `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           pulumi.StringPtrOutput `pulumi:"appSharingToLevel"`
	Authentication              pulumi.StringPtrOutput `pulumi:"authentication"`
	ClipboardSharingLevel       pulumi.StringPtrOutput `pulumi:"clipboardSharingLevel"`
	DataBackup                  pulumi.StringPtrOutput `pulumi:"dataBackup"`
	Description                 pulumi.StringPtrOutput `pulumi:"description"`
	DeviceCompliance            pulumi.StringPtrOutput `pulumi:"deviceCompliance"`
	FileEncryption              pulumi.StringPtrOutput `pulumi:"fileEncryption"`
	FileSharingSaveAs           pulumi.StringPtrOutput `pulumi:"fileSharingSaveAs"`
	FriendlyName                pulumi.StringOutput    `pulumi:"friendlyName"`
	GroupStatus                 pulumi.StringOutput    `pulumi:"groupStatus"`
	LastModifiedTime            pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	// Resource Location
	Location       pulumi.StringPtrOutput `pulumi:"location"`
	ManagedBrowser pulumi.StringPtrOutput `pulumi:"managedBrowser"`
	// Resource name
	Name               pulumi.StringOutput    `pulumi:"name"`
	NumOfApps          pulumi.IntOutput       `pulumi:"numOfApps"`
	OfflineWipeTimeout pulumi.StringPtrOutput `pulumi:"offlineWipeTimeout"`
	Pin                pulumi.StringPtrOutput `pulumi:"pin"`
	PinNumRetry        pulumi.IntPtrOutput    `pulumi:"pinNumRetry"`
	ScreenCapture      pulumi.StringPtrOutput `pulumi:"screenCapture"`
	// Resource Tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAndroidMAMPolicyByName registers a new resource with the given unique name, arguments, and options.
func NewAndroidMAMPolicyByName(ctx *pulumi.Context,
	name string, args *AndroidMAMPolicyByNameArgs, opts ...pulumi.ResourceOption) (*AndroidMAMPolicyByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FriendlyName == nil {
		return nil, errors.New("invalid value for required argument 'FriendlyName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.AppSharingFromLevel == nil {
		args.AppSharingFromLevel = pulumi.StringPtr("none")
	}
	if args.AppSharingToLevel == nil {
		args.AppSharingToLevel = pulumi.StringPtr("none")
	}
	if args.Authentication == nil {
		args.Authentication = pulumi.StringPtr("required")
	}
	if args.ClipboardSharingLevel == nil {
		args.ClipboardSharingLevel = pulumi.StringPtr("blocked")
	}
	if args.DataBackup == nil {
		args.DataBackup = pulumi.StringPtr("allow")
	}
	if args.DeviceCompliance == nil {
		args.DeviceCompliance = pulumi.StringPtr("enable")
	}
	if args.FileEncryption == nil {
		args.FileEncryption = pulumi.StringPtr("required")
	}
	if args.FileSharingSaveAs == nil {
		args.FileSharingSaveAs = pulumi.StringPtr("allow")
	}
	if args.ManagedBrowser == nil {
		args.ManagedBrowser = pulumi.StringPtr("required")
	}
	if args.Pin == nil {
		args.Pin = pulumi.StringPtr("required")
	}
	if args.ScreenCapture == nil {
		args.ScreenCapture = pulumi.StringPtr("allow")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:intune/v20150114preview:AndroidMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-native:intune/v20150114privatepreview:AndroidMAMPolicyByName"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AndroidMAMPolicyByName
	err := ctx.RegisterResource("azure-native:intune:AndroidMAMPolicyByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAndroidMAMPolicyByName gets an existing AndroidMAMPolicyByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAndroidMAMPolicyByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AndroidMAMPolicyByNameState, opts ...pulumi.ResourceOption) (*AndroidMAMPolicyByName, error) {
	var resource AndroidMAMPolicyByName
	err := ctx.ReadResource("azure-native:intune:AndroidMAMPolicyByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AndroidMAMPolicyByName resources.
type androidMAMPolicyByNameState struct {
}

type AndroidMAMPolicyByNameState struct {
}

func (AndroidMAMPolicyByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*androidMAMPolicyByNameState)(nil)).Elem()
}

type androidMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout *string `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string `pulumi:"appSharingToLevel"`
	Authentication              *string `pulumi:"authentication"`
	ClipboardSharingLevel       *string `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string `pulumi:"dataBackup"`
	Description                 *string `pulumi:"description"`
	DeviceCompliance            *string `pulumi:"deviceCompliance"`
	FileEncryption              *string `pulumi:"fileEncryption"`
	FileSharingSaveAs           *string `pulumi:"fileSharingSaveAs"`
	FriendlyName                string  `pulumi:"friendlyName"`
	// Location hostName for the tenant
	HostName string `pulumi:"hostName"`
	// Resource Location
	Location           *string `pulumi:"location"`
	ManagedBrowser     *string `pulumi:"managedBrowser"`
	OfflineWipeTimeout *string `pulumi:"offlineWipeTimeout"`
	Pin                *string `pulumi:"pin"`
	PinNumRetry        *int    `pulumi:"pinNumRetry"`
	// Unique name for the policy
	PolicyName    *string `pulumi:"policyName"`
	ScreenCapture *string `pulumi:"screenCapture"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AndroidMAMPolicyByName resource.
type AndroidMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout pulumi.StringPtrInput
	AccessRecheckOnlineTimeout  pulumi.StringPtrInput
	AppSharingFromLevel         pulumi.StringPtrInput
	AppSharingToLevel           pulumi.StringPtrInput
	Authentication              pulumi.StringPtrInput
	ClipboardSharingLevel       pulumi.StringPtrInput
	DataBackup                  pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	DeviceCompliance            pulumi.StringPtrInput
	FileEncryption              pulumi.StringPtrInput
	FileSharingSaveAs           pulumi.StringPtrInput
	FriendlyName                pulumi.StringInput
	// Location hostName for the tenant
	HostName pulumi.StringInput
	// Resource Location
	Location           pulumi.StringPtrInput
	ManagedBrowser     pulumi.StringPtrInput
	OfflineWipeTimeout pulumi.StringPtrInput
	Pin                pulumi.StringPtrInput
	PinNumRetry        pulumi.IntPtrInput
	// Unique name for the policy
	PolicyName    pulumi.StringPtrInput
	ScreenCapture pulumi.StringPtrInput
	// Resource Tags
	Tags pulumi.StringMapInput
}

func (AndroidMAMPolicyByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*androidMAMPolicyByNameArgs)(nil)).Elem()
}

type AndroidMAMPolicyByNameInput interface {
	pulumi.Input

	ToAndroidMAMPolicyByNameOutput() AndroidMAMPolicyByNameOutput
	ToAndroidMAMPolicyByNameOutputWithContext(ctx context.Context) AndroidMAMPolicyByNameOutput
}

func (*AndroidMAMPolicyByName) ElementType() reflect.Type {
	return reflect.TypeOf((**AndroidMAMPolicyByName)(nil)).Elem()
}

func (i *AndroidMAMPolicyByName) ToAndroidMAMPolicyByNameOutput() AndroidMAMPolicyByNameOutput {
	return i.ToAndroidMAMPolicyByNameOutputWithContext(context.Background())
}

func (i *AndroidMAMPolicyByName) ToAndroidMAMPolicyByNameOutputWithContext(ctx context.Context) AndroidMAMPolicyByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AndroidMAMPolicyByNameOutput)
}

func (i *AndroidMAMPolicyByName) ToOutput(ctx context.Context) pulumix.Output[*AndroidMAMPolicyByName] {
	return pulumix.Output[*AndroidMAMPolicyByName]{
		OutputState: i.ToAndroidMAMPolicyByNameOutputWithContext(ctx).OutputState,
	}
}

type AndroidMAMPolicyByNameOutput struct{ *pulumi.OutputState }

func (AndroidMAMPolicyByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AndroidMAMPolicyByName)(nil)).Elem()
}

func (o AndroidMAMPolicyByNameOutput) ToAndroidMAMPolicyByNameOutput() AndroidMAMPolicyByNameOutput {
	return o
}

func (o AndroidMAMPolicyByNameOutput) ToAndroidMAMPolicyByNameOutputWithContext(ctx context.Context) AndroidMAMPolicyByNameOutput {
	return o
}

func (o AndroidMAMPolicyByNameOutput) ToOutput(ctx context.Context) pulumix.Output[*AndroidMAMPolicyByName] {
	return pulumix.Output[*AndroidMAMPolicyByName]{
		OutputState: o.OutputState,
	}
}

func (o AndroidMAMPolicyByNameOutput) AccessRecheckOfflineTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.AccessRecheckOfflineTimeout }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) AccessRecheckOnlineTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.AccessRecheckOnlineTimeout }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) AppSharingFromLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.AppSharingFromLevel }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) AppSharingToLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.AppSharingToLevel }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.Authentication }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) ClipboardSharingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.ClipboardSharingLevel }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) DataBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.DataBackup }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) DeviceCompliance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.DeviceCompliance }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) FileEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.FileEncryption }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) FileSharingSaveAs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.FileSharingSaveAs }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o AndroidMAMPolicyByNameOutput) GroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringOutput { return v.GroupStatus }).(pulumi.StringOutput)
}

func (o AndroidMAMPolicyByNameOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// Resource Location
func (o AndroidMAMPolicyByNameOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) ManagedBrowser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.ManagedBrowser }).(pulumi.StringPtrOutput)
}

// Resource name
func (o AndroidMAMPolicyByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AndroidMAMPolicyByNameOutput) NumOfApps() pulumi.IntOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.IntOutput { return v.NumOfApps }).(pulumi.IntOutput)
}

func (o AndroidMAMPolicyByNameOutput) OfflineWipeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.OfflineWipeTimeout }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) Pin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.Pin }).(pulumi.StringPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) PinNumRetry() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.IntPtrOutput { return v.PinNumRetry }).(pulumi.IntPtrOutput)
}

func (o AndroidMAMPolicyByNameOutput) ScreenCapture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringPtrOutput { return v.ScreenCapture }).(pulumi.StringPtrOutput)
}

// Resource Tags
func (o AndroidMAMPolicyByNameOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o AndroidMAMPolicyByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidMAMPolicyByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AndroidMAMPolicyByNameOutput{})
}

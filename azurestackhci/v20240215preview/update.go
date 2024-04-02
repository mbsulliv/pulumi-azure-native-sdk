// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240215preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Update details
type Update struct {
	pulumi.CustomResourceState

	// Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
	AdditionalProperties pulumi.StringPtrOutput `pulumi:"additionalProperties"`
	// Indicates the way the update content can be downloaded.
	AvailabilityType pulumi.StringPtrOutput `pulumi:"availabilityType"`
	// Description of the update.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display name of the Update
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Last time the package-specific checks were run.
	HealthCheckDate pulumi.StringPtrOutput `pulumi:"healthCheckDate"`
	// Date that the update was installed.
	InstalledDate pulumi.StringPtrOutput `pulumi:"installedDate"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Brief message with instructions for updates of AvailabilityType Notify.
	NotifyMessage pulumi.StringPtrOutput `pulumi:"notifyMessage"`
	// Path where the update package is available.
	PackagePath pulumi.StringPtrOutput `pulumi:"packagePath"`
	// Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
	PackageSizeInMb pulumi.Float64PtrOutput `pulumi:"packageSizeInMb"`
	// Customer-visible type of the update.
	PackageType pulumi.StringPtrOutput `pulumi:"packageType"`
	// If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
	Prerequisites UpdatePrerequisiteResponseArrayOutput `pulumi:"prerequisites"`
	// Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
	ProgressPercentage pulumi.Float64PtrOutput `pulumi:"progressPercentage"`
	// Provisioning state of the Updates proxy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Publisher of the update package.
	Publisher pulumi.StringPtrOutput `pulumi:"publisher"`
	// Link to release notes for the update.
	ReleaseLink pulumi.StringPtrOutput `pulumi:"releaseLink"`
	// State of the update as it relates to this stamp.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the update.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewUpdate registers a new resource with the given unique name, arguments, and options.
func NewUpdate(ctx *pulumi.Context,
	name string, args *UpdateArgs, opts ...pulumi.ResourceOption) (*Update, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221201:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230201:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230301:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230601:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801preview:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20231101preview:Update"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240101:Update"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Update
	err := ctx.RegisterResource("azure-native:azurestackhci/v20240215preview:Update", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpdate gets an existing Update resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpdate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpdateState, opts ...pulumi.ResourceOption) (*Update, error) {
	var resource Update
	err := ctx.ReadResource("azure-native:azurestackhci/v20240215preview:Update", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Update resources.
type updateState struct {
}

type UpdateState struct {
}

func (UpdateState) ElementType() reflect.Type {
	return reflect.TypeOf((*updateState)(nil)).Elem()
}

type updateArgs struct {
	// Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
	AdditionalProperties *string `pulumi:"additionalProperties"`
	// Indicates the way the update content can be downloaded.
	AvailabilityType *string `pulumi:"availabilityType"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// Description of the update.
	Description *string `pulumi:"description"`
	// Display name of the Update
	DisplayName *string `pulumi:"displayName"`
	// Last time the package-specific checks were run.
	HealthCheckDate *string `pulumi:"healthCheckDate"`
	// Date that the update was installed.
	InstalledDate *string `pulumi:"installedDate"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Brief message with instructions for updates of AvailabilityType Notify.
	NotifyMessage *string `pulumi:"notifyMessage"`
	// Path where the update package is available.
	PackagePath *string `pulumi:"packagePath"`
	// Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
	PackageSizeInMb *float64 `pulumi:"packageSizeInMb"`
	// Customer-visible type of the update.
	PackageType *string `pulumi:"packageType"`
	// If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
	Prerequisites []UpdatePrerequisite `pulumi:"prerequisites"`
	// Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
	ProgressPercentage *float64 `pulumi:"progressPercentage"`
	// Publisher of the update package.
	Publisher *string `pulumi:"publisher"`
	// Link to release notes for the update.
	ReleaseLink *string `pulumi:"releaseLink"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// State of the update as it relates to this stamp.
	State *string `pulumi:"state"`
	// The name of the Update
	UpdateName *string `pulumi:"updateName"`
	// Version of the update.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Update resource.
type UpdateArgs struct {
	// Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
	AdditionalProperties pulumi.StringPtrInput
	// Indicates the way the update content can be downloaded.
	AvailabilityType pulumi.StringPtrInput
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// Description of the update.
	Description pulumi.StringPtrInput
	// Display name of the Update
	DisplayName pulumi.StringPtrInput
	// Last time the package-specific checks were run.
	HealthCheckDate pulumi.StringPtrInput
	// Date that the update was installed.
	InstalledDate pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Brief message with instructions for updates of AvailabilityType Notify.
	NotifyMessage pulumi.StringPtrInput
	// Path where the update package is available.
	PackagePath pulumi.StringPtrInput
	// Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
	PackageSizeInMb pulumi.Float64PtrInput
	// Customer-visible type of the update.
	PackageType pulumi.StringPtrInput
	// If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
	Prerequisites UpdatePrerequisiteArrayInput
	// Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
	ProgressPercentage pulumi.Float64PtrInput
	// Publisher of the update package.
	Publisher pulumi.StringPtrInput
	// Link to release notes for the update.
	ReleaseLink pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// State of the update as it relates to this stamp.
	State pulumi.StringPtrInput
	// The name of the Update
	UpdateName pulumi.StringPtrInput
	// Version of the update.
	Version pulumi.StringPtrInput
}

func (UpdateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*updateArgs)(nil)).Elem()
}

type UpdateInput interface {
	pulumi.Input

	ToUpdateOutput() UpdateOutput
	ToUpdateOutputWithContext(ctx context.Context) UpdateOutput
}

func (*Update) ElementType() reflect.Type {
	return reflect.TypeOf((**Update)(nil)).Elem()
}

func (i *Update) ToUpdateOutput() UpdateOutput {
	return i.ToUpdateOutputWithContext(context.Background())
}

func (i *Update) ToUpdateOutputWithContext(ctx context.Context) UpdateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateOutput)
}

type UpdateOutput struct{ *pulumi.OutputState }

func (UpdateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Update)(nil)).Elem()
}

func (o UpdateOutput) ToUpdateOutput() UpdateOutput {
	return o
}

func (o UpdateOutput) ToUpdateOutputWithContext(ctx context.Context) UpdateOutput {
	return o
}

// Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
func (o UpdateOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.AdditionalProperties }).(pulumi.StringPtrOutput)
}

// Indicates the way the update content can be downloaded.
func (o UpdateOutput) AvailabilityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.AvailabilityType }).(pulumi.StringPtrOutput)
}

// Description of the update.
func (o UpdateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name of the Update
func (o UpdateOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Last time the package-specific checks were run.
func (o UpdateOutput) HealthCheckDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.HealthCheckDate }).(pulumi.StringPtrOutput)
}

// Date that the update was installed.
func (o UpdateOutput) InstalledDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.InstalledDate }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o UpdateOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o UpdateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Update) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Brief message with instructions for updates of AvailabilityType Notify.
func (o UpdateOutput) NotifyMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.NotifyMessage }).(pulumi.StringPtrOutput)
}

// Path where the update package is available.
func (o UpdateOutput) PackagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.PackagePath }).(pulumi.StringPtrOutput)
}

// Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
func (o UpdateOutput) PackageSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Update) pulumi.Float64PtrOutput { return v.PackageSizeInMb }).(pulumi.Float64PtrOutput)
}

// Customer-visible type of the update.
func (o UpdateOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.PackageType }).(pulumi.StringPtrOutput)
}

// If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
func (o UpdateOutput) Prerequisites() UpdatePrerequisiteResponseArrayOutput {
	return o.ApplyT(func(v *Update) UpdatePrerequisiteResponseArrayOutput { return v.Prerequisites }).(UpdatePrerequisiteResponseArrayOutput)
}

// Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
func (o UpdateOutput) ProgressPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Update) pulumi.Float64PtrOutput { return v.ProgressPercentage }).(pulumi.Float64PtrOutput)
}

// Provisioning state of the Updates proxy resource.
func (o UpdateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Update) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Publisher of the update package.
func (o UpdateOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

// Link to release notes for the update.
func (o UpdateOutput) ReleaseLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.ReleaseLink }).(pulumi.StringPtrOutput)
}

// State of the update as it relates to this stamp.
func (o UpdateOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o UpdateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Update) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o UpdateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Update) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the update.
func (o UpdateOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(UpdateOutput{})
}

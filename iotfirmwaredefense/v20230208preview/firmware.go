// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230208preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Firmware definition
type Firmware struct {
	pulumi.CustomResourceState

	// User-specified description of the firmware.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// File name for a firmware that user uploaded.
	FileName pulumi.StringPtrOutput `pulumi:"fileName"`
	// File size of the uploaded firmware image.
	FileSize pulumi.Float64PtrOutput `pulumi:"fileSize"`
	// Firmware model.
	Model pulumi.StringPtrOutput `pulumi:"model"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The status of firmware scan.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// A list of errors or other messages generated during firmware analysis
	StatusMessages pulumi.ArrayOutput `pulumi:"statusMessages"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Firmware vendor.
	Vendor pulumi.StringPtrOutput `pulumi:"vendor"`
	// Firmware version.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewFirmware registers a new resource with the given unique name, arguments, and options.
func NewFirmware(ctx *pulumi.Context,
	name string, args *FirmwareArgs, opts ...pulumi.ResourceOption) (*Firmware, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	if args.Status == nil {
		args.Status = pulumi.StringPtr("Pending")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotfirmwaredefense:Firmware"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Firmware
	err := ctx.RegisterResource("azure-native:iotfirmwaredefense/v20230208preview:Firmware", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirmware gets an existing Firmware resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirmware(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirmwareState, opts ...pulumi.ResourceOption) (*Firmware, error) {
	var resource Firmware
	err := ctx.ReadResource("azure-native:iotfirmwaredefense/v20230208preview:Firmware", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firmware resources.
type firmwareState struct {
}

type FirmwareState struct {
}

func (FirmwareState) ElementType() reflect.Type {
	return reflect.TypeOf((*firmwareState)(nil)).Elem()
}

type firmwareArgs struct {
	// User-specified description of the firmware.
	Description *string `pulumi:"description"`
	// File name for a firmware that user uploaded.
	FileName *string `pulumi:"fileName"`
	// File size of the uploaded firmware image.
	FileSize *float64 `pulumi:"fileSize"`
	// The id of the firmware.
	FirmwareId *string `pulumi:"firmwareId"`
	// Firmware model.
	Model *string `pulumi:"model"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of firmware scan.
	Status *string `pulumi:"status"`
	// A list of errors or other messages generated during firmware analysis
	StatusMessages []interface{} `pulumi:"statusMessages"`
	// Firmware vendor.
	Vendor *string `pulumi:"vendor"`
	// Firmware version.
	Version *string `pulumi:"version"`
	// The name of the firmware analysis workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Firmware resource.
type FirmwareArgs struct {
	// User-specified description of the firmware.
	Description pulumi.StringPtrInput
	// File name for a firmware that user uploaded.
	FileName pulumi.StringPtrInput
	// File size of the uploaded firmware image.
	FileSize pulumi.Float64PtrInput
	// The id of the firmware.
	FirmwareId pulumi.StringPtrInput
	// Firmware model.
	Model pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The status of firmware scan.
	Status pulumi.StringPtrInput
	// A list of errors or other messages generated during firmware analysis
	StatusMessages pulumi.ArrayInput
	// Firmware vendor.
	Vendor pulumi.StringPtrInput
	// Firmware version.
	Version pulumi.StringPtrInput
	// The name of the firmware analysis workspace.
	WorkspaceName pulumi.StringInput
}

func (FirmwareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firmwareArgs)(nil)).Elem()
}

type FirmwareInput interface {
	pulumi.Input

	ToFirmwareOutput() FirmwareOutput
	ToFirmwareOutputWithContext(ctx context.Context) FirmwareOutput
}

func (*Firmware) ElementType() reflect.Type {
	return reflect.TypeOf((**Firmware)(nil)).Elem()
}

func (i *Firmware) ToFirmwareOutput() FirmwareOutput {
	return i.ToFirmwareOutputWithContext(context.Background())
}

func (i *Firmware) ToFirmwareOutputWithContext(ctx context.Context) FirmwareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirmwareOutput)
}

func (i *Firmware) ToOutput(ctx context.Context) pulumix.Output[*Firmware] {
	return pulumix.Output[*Firmware]{
		OutputState: i.ToFirmwareOutputWithContext(ctx).OutputState,
	}
}

type FirmwareOutput struct{ *pulumi.OutputState }

func (FirmwareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firmware)(nil)).Elem()
}

func (o FirmwareOutput) ToFirmwareOutput() FirmwareOutput {
	return o
}

func (o FirmwareOutput) ToFirmwareOutputWithContext(ctx context.Context) FirmwareOutput {
	return o
}

func (o FirmwareOutput) ToOutput(ctx context.Context) pulumix.Output[*Firmware] {
	return pulumix.Output[*Firmware]{
		OutputState: o.OutputState,
	}
}

// User-specified description of the firmware.
func (o FirmwareOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// File name for a firmware that user uploaded.
func (o FirmwareOutput) FileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringPtrOutput { return v.FileName }).(pulumi.StringPtrOutput)
}

// File size of the uploaded firmware image.
func (o FirmwareOutput) FileSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Firmware) pulumi.Float64PtrOutput { return v.FileSize }).(pulumi.Float64PtrOutput)
}

// Firmware model.
func (o FirmwareOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringPtrOutput { return v.Model }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o FirmwareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o FirmwareOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The status of firmware scan.
func (o FirmwareOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// A list of errors or other messages generated during firmware analysis
func (o FirmwareOutput) StatusMessages() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Firmware) pulumi.ArrayOutput { return v.StatusMessages }).(pulumi.ArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FirmwareOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Firmware) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FirmwareOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Firmware vendor.
func (o FirmwareOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringPtrOutput { return v.Vendor }).(pulumi.StringPtrOutput)
}

// Firmware version.
func (o FirmwareOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firmware) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FirmwareOutput{})
}

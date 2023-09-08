// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
type CloudInitDataSource string

const (
	CloudInitDataSourceNoCloud = CloudInitDataSource("NoCloud")
	CloudInitDataSourceAzure   = CloudInitDataSource("Azure")
)

// The format of the actual VHD file [vhd, vhdx]
type DiskFileFormat string

const (
	DiskFileFormatVhdx = DiskFileFormat("vhdx")
	DiskFileFormatVhd  = DiskFileFormat("vhd")
)

// The type of the extended location.
type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation = ExtendedLocationTypes("CustomLocation")
)

// The hypervisor generation of the Virtual Machine [V1, V2]
type HyperVGeneration string

const (
	HyperVGenerationV1 = HyperVGeneration("V1")
	HyperVGenerationV2 = HyperVGeneration("V2")
)

// IPAllocationMethod - The IP address allocation method. Possible values include: 'Static', 'Dynamic'
type IpAllocationMethodEnum string

const (
	IpAllocationMethodEnumDynamic = IpAllocationMethodEnum("Dynamic")
	IpAllocationMethodEnumStatic  = IpAllocationMethodEnum("Static")
)

// Type of the network
type NetworkTypeEnum string

const (
	NetworkTypeEnumNAT         = NetworkTypeEnum("NAT")
	NetworkTypeEnumTransparent = NetworkTypeEnum("Transparent")
	NetworkTypeEnumL2Bridge    = NetworkTypeEnum("L2Bridge")
	NetworkTypeEnumL2Tunnel    = NetworkTypeEnum("L2Tunnel")
	NetworkTypeEnumICS         = NetworkTypeEnum("ICS")
	NetworkTypeEnumPrivate     = NetworkTypeEnum("Private")
	NetworkTypeEnumOverlay     = NetworkTypeEnum("Overlay")
	NetworkTypeEnumInternal    = NetworkTypeEnum("Internal")
	NetworkTypeEnumMirrored    = NetworkTypeEnum("Mirrored")
)

// This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD. Possible values are: **Windows,** **Linux.**
type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
)

func (OperatingSystemTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemTypes)(nil)).Elem()
}

func (e OperatingSystemTypes) ToOperatingSystemTypesOutput() OperatingSystemTypesOutput {
	return pulumi.ToOutput(e).(OperatingSystemTypesOutput)
}

func (e OperatingSystemTypes) ToOperatingSystemTypesOutputWithContext(ctx context.Context) OperatingSystemTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemTypesOutput)
}

func (e OperatingSystemTypes) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return e.ToOperatingSystemTypesPtrOutputWithContext(context.Background())
}

func (e OperatingSystemTypes) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return OperatingSystemTypes(e).ToOperatingSystemTypesOutputWithContext(ctx).ToOperatingSystemTypesPtrOutputWithContext(ctx)
}

func (e OperatingSystemTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystemTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystemTypesOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemTypes)(nil)).Elem()
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesOutput() OperatingSystemTypesOutput {
	return o
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesOutputWithContext(ctx context.Context) OperatingSystemTypesOutput {
	return o
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return o.ToOperatingSystemTypesPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystemTypes) *OperatingSystemTypes {
		return &v
	}).(OperatingSystemTypesPtrOutput)
}

func (o OperatingSystemTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemTypesPtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystemTypes)(nil)).Elem()
}

func (o OperatingSystemTypesPtrOutput) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return o
}

func (o OperatingSystemTypesPtrOutput) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return o
}

func (o OperatingSystemTypesPtrOutput) Elem() OperatingSystemTypesOutput {
	return o.ApplyT(func(v *OperatingSystemTypes) OperatingSystemTypes {
		if v != nil {
			return *v
		}
		var ret OperatingSystemTypes
		return ret
	}).(OperatingSystemTypesOutput)
}

func (o OperatingSystemTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystemTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OperatingSystemTypesInput is an input type that accepts OperatingSystemTypesArgs and OperatingSystemTypesOutput values.
// You can construct a concrete instance of `OperatingSystemTypesInput` via:
//
//	OperatingSystemTypesArgs{...}
type OperatingSystemTypesInput interface {
	pulumi.Input

	ToOperatingSystemTypesOutput() OperatingSystemTypesOutput
	ToOperatingSystemTypesOutputWithContext(context.Context) OperatingSystemTypesOutput
}

var operatingSystemTypesPtrType = reflect.TypeOf((**OperatingSystemTypes)(nil)).Elem()

type OperatingSystemTypesPtrInput interface {
	pulumi.Input

	ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput
	ToOperatingSystemTypesPtrOutputWithContext(context.Context) OperatingSystemTypesPtrOutput
}

type operatingSystemTypesPtr string

func OperatingSystemTypesPtr(v string) OperatingSystemTypesPtrInput {
	return (*operatingSystemTypesPtr)(&v)
}

func (*operatingSystemTypesPtr) ElementType() reflect.Type {
	return operatingSystemTypesPtrType
}

func (in *operatingSystemTypesPtr) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemTypesPtrOutput)
}

func (in *operatingSystemTypesPtr) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemTypesPtrOutput)
}

// The guest agent provisioning action.
type ProvisioningAction string

const (
	ProvisioningActionInstall   = ProvisioningAction("install")
	ProvisioningActionUninstall = ProvisioningAction("uninstall")
	ProvisioningActionRepair    = ProvisioningAction("repair")
)

// The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//	ResourceIdentityTypeArgs{...}
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

// Specifies the SecurityType of the virtual machine. EnableTPM and SecureBootEnabled must be set to true for SecurityType to function.
type SecurityTypes string

const (
	SecurityTypesTrustedLaunch  = SecurityTypes("TrustedLaunch")
	SecurityTypesConfidentialVM = SecurityTypes("ConfidentialVM")
)

type VmSizeEnum string

const (
	VmSizeEnumDefault           = VmSizeEnum("Default")
	VmSizeEnum_Standard_A2_v2   = VmSizeEnum("Standard_A2_v2")
	VmSizeEnum_Standard_A4_v2   = VmSizeEnum("Standard_A4_v2")
	VmSizeEnum_Standard_D2s_v3  = VmSizeEnum("Standard_D2s_v3")
	VmSizeEnum_Standard_D4s_v3  = VmSizeEnum("Standard_D4s_v3")
	VmSizeEnum_Standard_D8s_v3  = VmSizeEnum("Standard_D8s_v3")
	VmSizeEnum_Standard_D16s_v3 = VmSizeEnum("Standard_D16s_v3")
	VmSizeEnum_Standard_D32s_v3 = VmSizeEnum("Standard_D32s_v3")
	VmSizeEnum_Standard_DS2_v2  = VmSizeEnum("Standard_DS2_v2")
	VmSizeEnum_Standard_DS3_v2  = VmSizeEnum("Standard_DS3_v2")
	VmSizeEnum_Standard_DS4_v2  = VmSizeEnum("Standard_DS4_v2")
	VmSizeEnum_Standard_DS5_v2  = VmSizeEnum("Standard_DS5_v2")
	VmSizeEnum_Standard_DS13_v2 = VmSizeEnum("Standard_DS13_v2")
	VmSizeEnum_Standard_K8S_v1  = VmSizeEnum("Standard_K8S_v1")
	VmSizeEnum_Standard_K8S2_v1 = VmSizeEnum("Standard_K8S2_v1")
	VmSizeEnum_Standard_K8S3_v1 = VmSizeEnum("Standard_K8S3_v1")
	VmSizeEnum_Standard_K8S4_v1 = VmSizeEnum("Standard_K8S4_v1")
	VmSizeEnum_Standard_NK6     = VmSizeEnum("Standard_NK6")
	VmSizeEnum_Standard_NK12    = VmSizeEnum("Standard_NK12")
	VmSizeEnum_Standard_NV6     = VmSizeEnum("Standard_NV6")
	VmSizeEnum_Standard_NV12    = VmSizeEnum("Standard_NV12")
	VmSizeEnum_Standard_K8S5_v1 = VmSizeEnum("Standard_K8S5_v1")
	VmSizeEnumCustom            = VmSizeEnum("Custom")
)

func init() {
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
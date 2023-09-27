// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Packet core control plane resource.
type PacketCoreControlPlane struct {
	pulumi.CustomResourceState

	// The control plane interface on the access network. In 5G networks this is called as N2 interface whereas in 4G networks this is called as S1-MME interface.
	ControlPlaneAccessInterface InterfacePropertiesResponseOutput `pulumi:"controlPlaneAccessInterface"`
	// The core network technology generation.
	CoreNetworkTechnology pulumi.StringPtrOutput `pulumi:"coreNetworkTechnology"`
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrOutput `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrOutput `pulumi:"createdByType"`
	// Azure ARC custom location where the packet core is deployed.
	CustomLocation CustomLocationResourceIdResponsePtrOutput `pulumi:"customLocation"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrOutput `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrOutput `pulumi:"lastModifiedByType"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Mobile network that this packet core control plane belongs to
	MobileNetwork MobileNetworkResourceIdResponseOutput `pulumi:"mobileNetwork"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the packet core control plane resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The version of the packet core software that is deployed.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewPacketCoreControlPlane registers a new resource with the given unique name, arguments, and options.
func NewPacketCoreControlPlane(ctx *pulumi.Context,
	name string, args *PacketCoreControlPlaneArgs, opts ...pulumi.ResourceOption) (*PacketCoreControlPlane, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPlaneAccessInterface == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneAccessInterface'")
	}
	if args.MobileNetwork == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetwork'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:PacketCoreControlPlane"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:PacketCoreControlPlane"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:PacketCoreControlPlane"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20230601:PacketCoreControlPlane"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PacketCoreControlPlane
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220301preview:PacketCoreControlPlane", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPacketCoreControlPlane gets an existing PacketCoreControlPlane resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPacketCoreControlPlane(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketCoreControlPlaneState, opts ...pulumi.ResourceOption) (*PacketCoreControlPlane, error) {
	var resource PacketCoreControlPlane
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220301preview:PacketCoreControlPlane", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PacketCoreControlPlane resources.
type packetCoreControlPlaneState struct {
}

type PacketCoreControlPlaneState struct {
}

func (PacketCoreControlPlaneState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCoreControlPlaneState)(nil)).Elem()
}

type packetCoreControlPlaneArgs struct {
	// The control plane interface on the access network. In 5G networks this is called as N2 interface whereas in 4G networks this is called as S1-MME interface.
	ControlPlaneAccessInterface InterfaceProperties `pulumi:"controlPlaneAccessInterface"`
	// The core network technology generation.
	CoreNetworkTechnology *string `pulumi:"coreNetworkTechnology"`
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// Azure ARC custom location where the packet core is deployed.
	CustomLocation *CustomLocationResourceId `pulumi:"customLocation"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Mobile network that this packet core control plane belongs to
	MobileNetwork MobileNetworkResourceId `pulumi:"mobileNetwork"`
	// The name of the packet core control plane.
	PacketCoreControlPlaneName *string `pulumi:"packetCoreControlPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The version of the packet core software that is deployed.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a PacketCoreControlPlane resource.
type PacketCoreControlPlaneArgs struct {
	// The control plane interface on the access network. In 5G networks this is called as N2 interface whereas in 4G networks this is called as S1-MME interface.
	ControlPlaneAccessInterface InterfacePropertiesInput
	// The core network technology generation.
	CoreNetworkTechnology pulumi.StringPtrInput
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrInput
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrInput
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrInput
	// Azure ARC custom location where the packet core is deployed.
	CustomLocation CustomLocationResourceIdPtrInput
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrInput
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrInput
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Mobile network that this packet core control plane belongs to
	MobileNetwork MobileNetworkResourceIdInput
	// The name of the packet core control plane.
	PacketCoreControlPlaneName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The version of the packet core software that is deployed.
	Version pulumi.StringPtrInput
}

func (PacketCoreControlPlaneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCoreControlPlaneArgs)(nil)).Elem()
}

type PacketCoreControlPlaneInput interface {
	pulumi.Input

	ToPacketCoreControlPlaneOutput() PacketCoreControlPlaneOutput
	ToPacketCoreControlPlaneOutputWithContext(ctx context.Context) PacketCoreControlPlaneOutput
}

func (*PacketCoreControlPlane) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCoreControlPlane)(nil)).Elem()
}

func (i *PacketCoreControlPlane) ToPacketCoreControlPlaneOutput() PacketCoreControlPlaneOutput {
	return i.ToPacketCoreControlPlaneOutputWithContext(context.Background())
}

func (i *PacketCoreControlPlane) ToPacketCoreControlPlaneOutputWithContext(ctx context.Context) PacketCoreControlPlaneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCoreControlPlaneOutput)
}

func (i *PacketCoreControlPlane) ToOutput(ctx context.Context) pulumix.Output[*PacketCoreControlPlane] {
	return pulumix.Output[*PacketCoreControlPlane]{
		OutputState: i.ToPacketCoreControlPlaneOutputWithContext(ctx).OutputState,
	}
}

type PacketCoreControlPlaneOutput struct{ *pulumi.OutputState }

func (PacketCoreControlPlaneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketCoreControlPlane)(nil)).Elem()
}

func (o PacketCoreControlPlaneOutput) ToPacketCoreControlPlaneOutput() PacketCoreControlPlaneOutput {
	return o
}

func (o PacketCoreControlPlaneOutput) ToPacketCoreControlPlaneOutputWithContext(ctx context.Context) PacketCoreControlPlaneOutput {
	return o
}

func (o PacketCoreControlPlaneOutput) ToOutput(ctx context.Context) pulumix.Output[*PacketCoreControlPlane] {
	return pulumix.Output[*PacketCoreControlPlane]{
		OutputState: o.OutputState,
	}
}

// The control plane interface on the access network. In 5G networks this is called as N2 interface whereas in 4G networks this is called as S1-MME interface.
func (o PacketCoreControlPlaneOutput) ControlPlaneAccessInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) InterfacePropertiesResponseOutput {
		return v.ControlPlaneAccessInterface
	}).(InterfacePropertiesResponseOutput)
}

// The core network technology generation.
func (o PacketCoreControlPlaneOutput) CoreNetworkTechnology() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.CoreNetworkTechnology }).(pulumi.StringPtrOutput)
}

// The timestamp of resource creation (UTC).
func (o PacketCoreControlPlaneOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o PacketCoreControlPlaneOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o PacketCoreControlPlaneOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// Azure ARC custom location where the packet core is deployed.
func (o PacketCoreControlPlaneOutput) CustomLocation() CustomLocationResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) CustomLocationResourceIdResponsePtrOutput { return v.CustomLocation }).(CustomLocationResourceIdResponsePtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o PacketCoreControlPlaneOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o PacketCoreControlPlaneOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o PacketCoreControlPlaneOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o PacketCoreControlPlaneOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Mobile network that this packet core control plane belongs to
func (o PacketCoreControlPlaneOutput) MobileNetwork() MobileNetworkResourceIdResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) MobileNetworkResourceIdResponseOutput { return v.MobileNetwork }).(MobileNetworkResourceIdResponseOutput)
}

// The name of the resource
func (o PacketCoreControlPlaneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the packet core control plane resource.
func (o PacketCoreControlPlaneOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PacketCoreControlPlaneOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PacketCoreControlPlaneOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PacketCoreControlPlaneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The version of the packet core software that is deployed.
func (o PacketCoreControlPlaneOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PacketCoreControlPlane) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PacketCoreControlPlaneOutput{})
}

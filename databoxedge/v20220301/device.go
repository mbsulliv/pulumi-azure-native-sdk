// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Data Box Edge/Gateway device.
type Device struct {
	pulumi.CustomResourceState

	// Type of compute roles configured.
	ConfiguredRoleTypes pulumi.StringArrayOutput `pulumi:"configuredRoleTypes"`
	// The Data Box Edge/Gateway device culture.
	Culture pulumi.StringOutput `pulumi:"culture"`
	// The status of the Data Box Edge/Gateway device.
	DataBoxEdgeDeviceStatus pulumi.StringOutput `pulumi:"dataBoxEdgeDeviceStatus"`
	// The details of data-residency related properties for this resource
	DataResidency DataResidencyResponsePtrOutput `pulumi:"dataResidency"`
	// The Description of the Data Box Edge/Gateway device.
	Description pulumi.StringOutput `pulumi:"description"`
	// The device software version number of the device (eg: 1.2.18105.6).
	DeviceHcsVersion pulumi.StringOutput `pulumi:"deviceHcsVersion"`
	// The Data Box Edge/Gateway device local capacity in MB.
	DeviceLocalCapacity pulumi.Float64Output `pulumi:"deviceLocalCapacity"`
	// The Data Box Edge/Gateway device model.
	DeviceModel pulumi.StringOutput `pulumi:"deviceModel"`
	// The Data Box Edge/Gateway device software version.
	DeviceSoftwareVersion pulumi.StringOutput `pulumi:"deviceSoftwareVersion"`
	// The type of the Data Box Edge/Gateway device.
	DeviceType pulumi.StringOutput `pulumi:"deviceType"`
	// The details of Edge Profile for this resource
	EdgeProfile EdgeProfileResponseOutput `pulumi:"edgeProfile"`
	// The etag for the devices.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The Data Box Edge/Gateway device name.
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// Msi identity of the resource
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// The kind of the device.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location pulumi.StringOutput `pulumi:"location"`
	// The description of the Data Box Edge/Gateway device model.
	ModelDescription pulumi.StringOutput `pulumi:"modelDescription"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes in the cluster.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// The details of the move operation on this resource.
	ResourceMoveDetails ResourceMoveDetailsResponseOutput `pulumi:"resourceMoveDetails"`
	// The Serial Number of Data Box Edge/Gateway device.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// The SKU type.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// DataBoxEdge Resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Data Box Edge/Gateway device timezone.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230701:Device"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Device
	err := ctx.RegisterResource("azure-native:databoxedge/v20220301:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("azure-native:databoxedge/v20220301:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
}

type DeviceState struct {
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// The details of data-residency related properties for this resource
	DataResidency *DataResidency `pulumi:"dataResidency"`
	// The device name.
	DeviceName *string `pulumi:"deviceName"`
	// Msi identity of the resource
	Identity *ResourceIdentity `pulumi:"identity"`
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU type.
	Sku *Sku `pulumi:"sku"`
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// The details of data-residency related properties for this resource
	DataResidency DataResidencyPtrInput
	// The device name.
	DeviceName pulumi.StringPtrInput
	// Msi identity of the resource
	Identity ResourceIdentityPtrInput
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The SKU type.
	Sku SkuPtrInput
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags pulumi.StringMapInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

func (i *Device) ToOutput(ctx context.Context) pulumix.Output[*Device] {
	return pulumix.Output[*Device]{
		OutputState: i.ToDeviceOutputWithContext(ctx).OutputState,
	}
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

func (o DeviceOutput) ToOutput(ctx context.Context) pulumix.Output[*Device] {
	return pulumix.Output[*Device]{
		OutputState: o.OutputState,
	}
}

// Type of compute roles configured.
func (o DeviceOutput) ConfiguredRoleTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Device) pulumi.StringArrayOutput { return v.ConfiguredRoleTypes }).(pulumi.StringArrayOutput)
}

// The Data Box Edge/Gateway device culture.
func (o DeviceOutput) Culture() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Culture }).(pulumi.StringOutput)
}

// The status of the Data Box Edge/Gateway device.
func (o DeviceOutput) DataBoxEdgeDeviceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DataBoxEdgeDeviceStatus }).(pulumi.StringOutput)
}

// The details of data-residency related properties for this resource
func (o DeviceOutput) DataResidency() DataResidencyResponsePtrOutput {
	return o.ApplyT(func(v *Device) DataResidencyResponsePtrOutput { return v.DataResidency }).(DataResidencyResponsePtrOutput)
}

// The Description of the Data Box Edge/Gateway device.
func (o DeviceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The device software version number of the device (eg: 1.2.18105.6).
func (o DeviceOutput) DeviceHcsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceHcsVersion }).(pulumi.StringOutput)
}

// The Data Box Edge/Gateway device local capacity in MB.
func (o DeviceOutput) DeviceLocalCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v *Device) pulumi.Float64Output { return v.DeviceLocalCapacity }).(pulumi.Float64Output)
}

// The Data Box Edge/Gateway device model.
func (o DeviceOutput) DeviceModel() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceModel }).(pulumi.StringOutput)
}

// The Data Box Edge/Gateway device software version.
func (o DeviceOutput) DeviceSoftwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceSoftwareVersion }).(pulumi.StringOutput)
}

// The type of the Data Box Edge/Gateway device.
func (o DeviceOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceType }).(pulumi.StringOutput)
}

// The details of Edge Profile for this resource
func (o DeviceOutput) EdgeProfile() EdgeProfileResponseOutput {
	return o.ApplyT(func(v *Device) EdgeProfileResponseOutput { return v.EdgeProfile }).(EdgeProfileResponseOutput)
}

// The etag for the devices.
func (o DeviceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The Data Box Edge/Gateway device name.
func (o DeviceOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

// Msi identity of the resource
func (o DeviceOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Device) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// The kind of the device.
func (o DeviceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
func (o DeviceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The description of the Data Box Edge/Gateway device model.
func (o DeviceOutput) ModelDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.ModelDescription }).(pulumi.StringOutput)
}

// The object name.
func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of nodes in the cluster.
func (o DeviceOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Device) pulumi.IntOutput { return v.NodeCount }).(pulumi.IntOutput)
}

// The details of the move operation on this resource.
func (o DeviceOutput) ResourceMoveDetails() ResourceMoveDetailsResponseOutput {
	return o.ApplyT(func(v *Device) ResourceMoveDetailsResponseOutput { return v.ResourceMoveDetails }).(ResourceMoveDetailsResponseOutput)
}

// The Serial Number of Data Box Edge/Gateway device.
func (o DeviceOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// The SKU type.
func (o DeviceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Device) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// DataBoxEdge Resource
func (o DeviceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Device) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
func (o DeviceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Device) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The Data Box Edge/Gateway device timezone.
func (o DeviceOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

// The hierarchical type of the object.
func (o DeviceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeviceOutput{})
}

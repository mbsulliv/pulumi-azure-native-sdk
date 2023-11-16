// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Device resource.
// Azure REST API version: 2022-01-01-preview. Prior API version in Azure Native 1.x: 2020-01-01-preview.
type Device struct {
	pulumi.CustomResourceState

	// The type of the device.
	DeviceType pulumi.StringOutput `pulumi:"deviceType"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of network functions deployed on the device.
	NetworkFunctions SubResourceResponseArrayOutput `pulumi:"networkFunctions"`
	// The provisioning state of the device resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The current device status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceType == nil {
		return nil, errors.New("invalid value for required argument 'DeviceType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20200101preview:Device"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20210501:Device"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20220101preview:Device"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Device
	err := ctx.RegisterResource("azure-native:hybridnetwork:Device", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:hybridnetwork:Device", name, id, state, &resource, opts...)
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
	// Resource name for the device resource.
	DeviceName *string `pulumi:"deviceName"`
	// The type of the device.
	DeviceType string `pulumi:"deviceType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// Resource name for the device resource.
	DeviceName pulumi.StringPtrInput
	// The type of the device.
	DeviceType pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
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

// The type of the device.
func (o DeviceOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceType }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o DeviceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of network functions deployed on the device.
func (o DeviceOutput) NetworkFunctions() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *Device) SubResourceResponseArrayOutput { return v.NetworkFunctions }).(SubResourceResponseArrayOutput)
}

// The provisioning state of the device resource.
func (o DeviceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The current device status.
func (o DeviceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o DeviceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Device) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DeviceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Device) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DeviceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeviceOutput{})
}

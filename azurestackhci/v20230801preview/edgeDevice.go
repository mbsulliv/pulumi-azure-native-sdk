// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Edge device resource
type EdgeDevice struct {
	pulumi.CustomResourceState

	// Device Configuration
	DeviceConfiguration DeviceConfigurationResponseOutput `pulumi:"deviceConfiguration"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of edgeDevice resource
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEdgeDevice registers a new resource with the given unique name, arguments, and options.
func NewEdgeDevice(ctx *pulumi.Context,
	name string, args *EdgeDeviceArgs, opts ...pulumi.ResourceOption) (*EdgeDevice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DeviceConfiguration'")
	}
	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:EdgeDevice"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20231101preview:EdgeDevice"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240101:EdgeDevice"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240215preview:EdgeDevice"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EdgeDevice
	err := ctx.RegisterResource("azure-native:azurestackhci/v20230801preview:EdgeDevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeDevice gets an existing EdgeDevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeDeviceState, opts ...pulumi.ResourceOption) (*EdgeDevice, error) {
	var resource EdgeDevice
	err := ctx.ReadResource("azure-native:azurestackhci/v20230801preview:EdgeDevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeDevice resources.
type edgeDeviceState struct {
}

type EdgeDeviceState struct {
}

func (EdgeDeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeDeviceState)(nil)).Elem()
}

type edgeDeviceArgs struct {
	// Device Configuration
	DeviceConfiguration DeviceConfiguration `pulumi:"deviceConfiguration"`
	// Name of Device
	EdgeDeviceName *string `pulumi:"edgeDeviceName"`
	// The fully qualified Azure Resource manager identifier of the resource.
	ResourceUri string `pulumi:"resourceUri"`
}

// The set of arguments for constructing a EdgeDevice resource.
type EdgeDeviceArgs struct {
	// Device Configuration
	DeviceConfiguration DeviceConfigurationInput
	// Name of Device
	EdgeDeviceName pulumi.StringPtrInput
	// The fully qualified Azure Resource manager identifier of the resource.
	ResourceUri pulumi.StringInput
}

func (EdgeDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeDeviceArgs)(nil)).Elem()
}

type EdgeDeviceInput interface {
	pulumi.Input

	ToEdgeDeviceOutput() EdgeDeviceOutput
	ToEdgeDeviceOutputWithContext(ctx context.Context) EdgeDeviceOutput
}

func (*EdgeDevice) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeDevice)(nil)).Elem()
}

func (i *EdgeDevice) ToEdgeDeviceOutput() EdgeDeviceOutput {
	return i.ToEdgeDeviceOutputWithContext(context.Background())
}

func (i *EdgeDevice) ToEdgeDeviceOutputWithContext(ctx context.Context) EdgeDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeDeviceOutput)
}

type EdgeDeviceOutput struct{ *pulumi.OutputState }

func (EdgeDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeDevice)(nil)).Elem()
}

func (o EdgeDeviceOutput) ToEdgeDeviceOutput() EdgeDeviceOutput {
	return o
}

func (o EdgeDeviceOutput) ToEdgeDeviceOutputWithContext(ctx context.Context) EdgeDeviceOutput {
	return o
}

// Device Configuration
func (o EdgeDeviceOutput) DeviceConfiguration() DeviceConfigurationResponseOutput {
	return o.ApplyT(func(v *EdgeDevice) DeviceConfigurationResponseOutput { return v.DeviceConfiguration }).(DeviceConfigurationResponseOutput)
}

// The name of the resource
func (o EdgeDeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeDevice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of edgeDevice resource
func (o EdgeDeviceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeDevice) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EdgeDeviceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EdgeDevice) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EdgeDeviceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeDevice) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EdgeDeviceOutput{})
}

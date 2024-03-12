// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Asset definition.
// Azure REST API version: 2023-11-01-preview.
type Asset struct {
	pulumi.CustomResourceState

	// A reference to the asset endpoint profile (connection information) used by brokers to connect to an endpoint that provides data points for this asset. Must have the format <ModuleCR.metadata.namespace>/<ModuleCR.metadata.name>.
	AssetEndpointProfileUri pulumi.StringOutput `pulumi:"assetEndpointProfileUri"`
	// Resource path to asset type (model) definition.
	AssetType pulumi.StringPtrOutput `pulumi:"assetType"`
	// A set of key-value pairs that contain custom attributes set by the customer.
	Attributes pulumi.AnyOutput `pulumi:"attributes"`
	// Array of data points that are part of the asset. Each data point can reference an asset type capability and have per-data point configuration. See below for more details for the definition of the dataPoints element.
	DataPoints DataPointResponseArrayOutput `pulumi:"dataPoints"`
	// Protocol-specific default configuration for all data points. Each data point can have its own configuration that overrides the default settings here. This assumes that each asset instance has one protocol.
	DefaultDataPointsConfiguration pulumi.StringPtrOutput `pulumi:"defaultDataPointsConfiguration"`
	// Protocol-specific default configuration for all events. Each event can have its own configuration that overrides the default settings here. This assumes that each asset instance has one protocol.
	DefaultEventsConfiguration pulumi.StringPtrOutput `pulumi:"defaultEventsConfiguration"`
	// Human-readable description of the asset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Human-readable display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Reference to the documentation.
	DocumentationUri pulumi.StringPtrOutput `pulumi:"documentationUri"`
	// Enabled/Disabled status of the asset.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Array of events that are part of the asset. Each event can reference an asset type capability and have per-event configuration. See below for more details about the definition of the events element.
	Events EventResponseArrayOutput `pulumi:"events"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Asset id provided by the customer.
	ExternalAssetId pulumi.StringPtrOutput `pulumi:"externalAssetId"`
	// Revision number of the hardware.
	HardwareRevision pulumi.StringPtrOutput `pulumi:"hardwareRevision"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Asset manufacturer name.
	Manufacturer pulumi.StringPtrOutput `pulumi:"manufacturer"`
	// Asset manufacturer URI.
	ManufacturerUri pulumi.StringPtrOutput `pulumi:"manufacturerUri"`
	// Asset model name.
	Model pulumi.StringPtrOutput `pulumi:"model"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Asset product code.
	ProductCode pulumi.StringPtrOutput `pulumi:"productCode"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Asset serial number.
	SerialNumber pulumi.StringPtrOutput `pulumi:"serialNumber"`
	// Revision number of the software.
	SoftwareRevision pulumi.StringPtrOutput `pulumi:"softwareRevision"`
	// Read only object to reflect changes that have occurred on the Edge. Similar to Kubernetes status property for custom resources.
	Status AssetStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Globally unique, immutable, non-reusable id.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// An integer that is incremented each time the resource is modified.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewAsset registers a new resource with the given unique name, arguments, and options.
func NewAsset(ctx *pulumi.Context,
	name string, args *AssetArgs, opts ...pulumi.ResourceOption) (*Asset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssetEndpointProfileUri == nil {
		return nil, errors.New("invalid value for required argument 'AssetEndpointProfileUri'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deviceregistry/v20231101preview:Asset"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Asset
	err := ctx.RegisterResource("azure-native:deviceregistry:Asset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAsset gets an existing Asset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetState, opts ...pulumi.ResourceOption) (*Asset, error) {
	var resource Asset
	err := ctx.ReadResource("azure-native:deviceregistry:Asset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Asset resources.
type assetState struct {
}

type AssetState struct {
}

func (AssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetState)(nil)).Elem()
}

type assetArgs struct {
	// A reference to the asset endpoint profile (connection information) used by brokers to connect to an endpoint that provides data points for this asset. Must have the format <ModuleCR.metadata.namespace>/<ModuleCR.metadata.name>.
	AssetEndpointProfileUri string `pulumi:"assetEndpointProfileUri"`
	// Asset name parameter.
	AssetName *string `pulumi:"assetName"`
	// Resource path to asset type (model) definition.
	AssetType *string `pulumi:"assetType"`
	// A set of key-value pairs that contain custom attributes set by the customer.
	Attributes interface{} `pulumi:"attributes"`
	// Array of data points that are part of the asset. Each data point can reference an asset type capability and have per-data point configuration. See below for more details for the definition of the dataPoints element.
	DataPoints []DataPoint `pulumi:"dataPoints"`
	// Protocol-specific default configuration for all data points. Each data point can have its own configuration that overrides the default settings here. This assumes that each asset instance has one protocol.
	DefaultDataPointsConfiguration *string `pulumi:"defaultDataPointsConfiguration"`
	// Protocol-specific default configuration for all events. Each event can have its own configuration that overrides the default settings here. This assumes that each asset instance has one protocol.
	DefaultEventsConfiguration *string `pulumi:"defaultEventsConfiguration"`
	// Human-readable description of the asset.
	Description *string `pulumi:"description"`
	// Human-readable display name.
	DisplayName *string `pulumi:"displayName"`
	// Reference to the documentation.
	DocumentationUri *string `pulumi:"documentationUri"`
	// Enabled/Disabled status of the asset.
	Enabled *bool `pulumi:"enabled"`
	// Array of events that are part of the asset. Each event can reference an asset type capability and have per-event configuration. See below for more details about the definition of the events element.
	Events []Event `pulumi:"events"`
	// The extended location.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Asset id provided by the customer.
	ExternalAssetId *string `pulumi:"externalAssetId"`
	// Revision number of the hardware.
	HardwareRevision *string `pulumi:"hardwareRevision"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Asset manufacturer name.
	Manufacturer *string `pulumi:"manufacturer"`
	// Asset manufacturer URI.
	ManufacturerUri *string `pulumi:"manufacturerUri"`
	// Asset model name.
	Model *string `pulumi:"model"`
	// Asset product code.
	ProductCode *string `pulumi:"productCode"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Asset serial number.
	SerialNumber *string `pulumi:"serialNumber"`
	// Revision number of the software.
	SoftwareRevision *string `pulumi:"softwareRevision"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Asset resource.
type AssetArgs struct {
	// A reference to the asset endpoint profile (connection information) used by brokers to connect to an endpoint that provides data points for this asset. Must have the format <ModuleCR.metadata.namespace>/<ModuleCR.metadata.name>.
	AssetEndpointProfileUri pulumi.StringInput
	// Asset name parameter.
	AssetName pulumi.StringPtrInput
	// Resource path to asset type (model) definition.
	AssetType pulumi.StringPtrInput
	// A set of key-value pairs that contain custom attributes set by the customer.
	Attributes pulumi.Input
	// Array of data points that are part of the asset. Each data point can reference an asset type capability and have per-data point configuration. See below for more details for the definition of the dataPoints element.
	DataPoints DataPointArrayInput
	// Protocol-specific default configuration for all data points. Each data point can have its own configuration that overrides the default settings here. This assumes that each asset instance has one protocol.
	DefaultDataPointsConfiguration pulumi.StringPtrInput
	// Protocol-specific default configuration for all events. Each event can have its own configuration that overrides the default settings here. This assumes that each asset instance has one protocol.
	DefaultEventsConfiguration pulumi.StringPtrInput
	// Human-readable description of the asset.
	Description pulumi.StringPtrInput
	// Human-readable display name.
	DisplayName pulumi.StringPtrInput
	// Reference to the documentation.
	DocumentationUri pulumi.StringPtrInput
	// Enabled/Disabled status of the asset.
	Enabled pulumi.BoolPtrInput
	// Array of events that are part of the asset. Each event can reference an asset type capability and have per-event configuration. See below for more details about the definition of the events element.
	Events EventArrayInput
	// The extended location.
	ExtendedLocation ExtendedLocationInput
	// Asset id provided by the customer.
	ExternalAssetId pulumi.StringPtrInput
	// Revision number of the hardware.
	HardwareRevision pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Asset manufacturer name.
	Manufacturer pulumi.StringPtrInput
	// Asset manufacturer URI.
	ManufacturerUri pulumi.StringPtrInput
	// Asset model name.
	Model pulumi.StringPtrInput
	// Asset product code.
	ProductCode pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Asset serial number.
	SerialNumber pulumi.StringPtrInput
	// Revision number of the software.
	SoftwareRevision pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetArgs)(nil)).Elem()
}

type AssetInput interface {
	pulumi.Input

	ToAssetOutput() AssetOutput
	ToAssetOutputWithContext(ctx context.Context) AssetOutput
}

func (*Asset) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (i *Asset) ToAssetOutput() AssetOutput {
	return i.ToAssetOutputWithContext(context.Background())
}

func (i *Asset) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetOutput)
}

type AssetOutput struct{ *pulumi.OutputState }

func (AssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (o AssetOutput) ToAssetOutput() AssetOutput {
	return o
}

func (o AssetOutput) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return o
}

// A reference to the asset endpoint profile (connection information) used by brokers to connect to an endpoint that provides data points for this asset. Must have the format <ModuleCR.metadata.namespace>/<ModuleCR.metadata.name>.
func (o AssetOutput) AssetEndpointProfileUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.AssetEndpointProfileUri }).(pulumi.StringOutput)
}

// Resource path to asset type (model) definition.
func (o AssetOutput) AssetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.AssetType }).(pulumi.StringPtrOutput)
}

// A set of key-value pairs that contain custom attributes set by the customer.
func (o AssetOutput) Attributes() pulumi.AnyOutput {
	return o.ApplyT(func(v *Asset) pulumi.AnyOutput { return v.Attributes }).(pulumi.AnyOutput)
}

// Array of data points that are part of the asset. Each data point can reference an asset type capability and have per-data point configuration. See below for more details for the definition of the dataPoints element.
func (o AssetOutput) DataPoints() DataPointResponseArrayOutput {
	return o.ApplyT(func(v *Asset) DataPointResponseArrayOutput { return v.DataPoints }).(DataPointResponseArrayOutput)
}

// Protocol-specific default configuration for all data points. Each data point can have its own configuration that overrides the default settings here. This assumes that each asset instance has one protocol.
func (o AssetOutput) DefaultDataPointsConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.DefaultDataPointsConfiguration }).(pulumi.StringPtrOutput)
}

// Protocol-specific default configuration for all events. Each event can have its own configuration that overrides the default settings here. This assumes that each asset instance has one protocol.
func (o AssetOutput) DefaultEventsConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.DefaultEventsConfiguration }).(pulumi.StringPtrOutput)
}

// Human-readable description of the asset.
func (o AssetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Human-readable display name.
func (o AssetOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Reference to the documentation.
func (o AssetOutput) DocumentationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.DocumentationUri }).(pulumi.StringPtrOutput)
}

// Enabled/Disabled status of the asset.
func (o AssetOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Array of events that are part of the asset. Each event can reference an asset type capability and have per-event configuration. See below for more details about the definition of the events element.
func (o AssetOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v *Asset) EventResponseArrayOutput { return v.Events }).(EventResponseArrayOutput)
}

// The extended location.
func (o AssetOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Asset) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Asset id provided by the customer.
func (o AssetOutput) ExternalAssetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.ExternalAssetId }).(pulumi.StringPtrOutput)
}

// Revision number of the hardware.
func (o AssetOutput) HardwareRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.HardwareRevision }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o AssetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Asset manufacturer name.
func (o AssetOutput) Manufacturer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.Manufacturer }).(pulumi.StringPtrOutput)
}

// Asset manufacturer URI.
func (o AssetOutput) ManufacturerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.ManufacturerUri }).(pulumi.StringPtrOutput)
}

// Asset model name.
func (o AssetOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.Model }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o AssetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Asset product code.
func (o AssetOutput) ProductCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.ProductCode }).(pulumi.StringPtrOutput)
}

// Provisioning state of the resource.
func (o AssetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Asset serial number.
func (o AssetOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.SerialNumber }).(pulumi.StringPtrOutput)
}

// Revision number of the software.
func (o AssetOutput) SoftwareRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringPtrOutput { return v.SoftwareRevision }).(pulumi.StringPtrOutput)
}

// Read only object to reflect changes that have occurred on the Edge. Similar to Kubernetes status property for custom resources.
func (o AssetOutput) Status() AssetStatusResponseOutput {
	return o.ApplyT(func(v *Asset) AssetStatusResponseOutput { return v.Status }).(AssetStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AssetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Asset) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AssetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AssetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Globally unique, immutable, non-reusable id.
func (o AssetOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// An integer that is incremented each time the resource is modified.
func (o AssetOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *Asset) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(AssetOutput{})
}

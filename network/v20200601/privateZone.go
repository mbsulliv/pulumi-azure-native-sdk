// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Private DNS zone.
type PrivateZone struct {
	pulumi.CustomResourceState

	// The ETag of the zone.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Private zone internal Id
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The maximum number of record sets that can be created in this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets pulumi.Float64Output `pulumi:"maxNumberOfRecordSets"`
	// The maximum number of virtual networks that can be linked to this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfVirtualNetworkLinks pulumi.Float64Output `pulumi:"maxNumberOfVirtualNetworkLinks"`
	// The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfVirtualNetworkLinksWithRegistration pulumi.Float64Output `pulumi:"maxNumberOfVirtualNetworkLinksWithRegistration"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The current number of record sets in this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets pulumi.Float64Output `pulumi:"numberOfRecordSets"`
	// The current number of virtual networks that are linked to this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NumberOfVirtualNetworkLinks pulumi.Float64Output `pulumi:"numberOfVirtualNetworkLinks"`
	// The current number of virtual networks that are linked to this Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
	NumberOfVirtualNetworkLinksWithRegistration pulumi.Float64Output `pulumi:"numberOfVirtualNetworkLinksWithRegistration"`
	// The provisioning state of the resource. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateZone registers a new resource with the given unique name, arguments, and options.
func NewPrivateZone(ctx *pulumi.Context,
	name string, args *PrivateZoneArgs, opts ...pulumi.ResourceOption) (*PrivateZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PrivateZone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180901:PrivateZone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200101:PrivateZone"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateZone
	err := ctx.RegisterResource("azure-native:network/v20200601:PrivateZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateZone gets an existing PrivateZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateZoneState, opts ...pulumi.ResourceOption) (*PrivateZone, error) {
	var resource PrivateZone
	err := ctx.ReadResource("azure-native:network/v20200601:PrivateZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateZone resources.
type privateZoneState struct {
}

type PrivateZoneState struct {
}

func (PrivateZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateZoneState)(nil)).Elem()
}

type privateZoneArgs struct {
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the Private DNS zone (without a terminating dot).
	PrivateZoneName *string `pulumi:"privateZoneName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateZone resource.
type PrivateZoneArgs struct {
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the Private DNS zone (without a terminating dot).
	PrivateZoneName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateZoneArgs)(nil)).Elem()
}

type PrivateZoneInput interface {
	pulumi.Input

	ToPrivateZoneOutput() PrivateZoneOutput
	ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput
}

func (*PrivateZone) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateZone)(nil)).Elem()
}

func (i *PrivateZone) ToPrivateZoneOutput() PrivateZoneOutput {
	return i.ToPrivateZoneOutputWithContext(context.Background())
}

func (i *PrivateZone) ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateZoneOutput)
}

type PrivateZoneOutput struct{ *pulumi.OutputState }

func (PrivateZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateZone)(nil)).Elem()
}

func (o PrivateZoneOutput) ToPrivateZoneOutput() PrivateZoneOutput {
	return o
}

func (o PrivateZoneOutput) ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput {
	return o
}

// The ETag of the zone.
func (o PrivateZoneOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Private zone internal Id
func (o PrivateZoneOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringOutput { return v.InternalId }).(pulumi.StringOutput)
}

// The Azure Region where the resource lives
func (o PrivateZoneOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The maximum number of record sets that can be created in this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
func (o PrivateZoneOutput) MaxNumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.MaxNumberOfRecordSets }).(pulumi.Float64Output)
}

// The maximum number of virtual networks that can be linked to this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
func (o PrivateZoneOutput) MaxNumberOfVirtualNetworkLinks() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.MaxNumberOfVirtualNetworkLinks }).(pulumi.Float64Output)
}

// The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
func (o PrivateZoneOutput) MaxNumberOfVirtualNetworkLinksWithRegistration() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.MaxNumberOfVirtualNetworkLinksWithRegistration }).(pulumi.Float64Output)
}

// The name of the resource
func (o PrivateZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current number of record sets in this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
func (o PrivateZoneOutput) NumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.NumberOfRecordSets }).(pulumi.Float64Output)
}

// The current number of virtual networks that are linked to this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
func (o PrivateZoneOutput) NumberOfVirtualNetworkLinks() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.NumberOfVirtualNetworkLinks }).(pulumi.Float64Output)
}

// The current number of virtual networks that are linked to this Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
func (o PrivateZoneOutput) NumberOfVirtualNetworkLinksWithRegistration() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateZone) pulumi.Float64Output { return v.NumberOfVirtualNetworkLinksWithRegistration }).(pulumi.Float64Output)
}

// The provisioning state of the resource. This is a read-only property and any attempt to set this value will be ignored.
func (o PrivateZoneOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o PrivateZoneOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
func (o PrivateZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateZone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateZoneOutput{})
}

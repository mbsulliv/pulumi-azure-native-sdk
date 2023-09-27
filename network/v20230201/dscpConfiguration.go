// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Differentiated Services Code Point configuration for any given network interface
type DscpConfiguration struct {
	pulumi.CustomResourceState

	// Associated Network Interfaces to the DSCP Configuration.
	AssociatedNetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"associatedNetworkInterfaces"`
	// Destination IP ranges.
	DestinationIpRanges QosIpRangeResponseArrayOutput `pulumi:"destinationIpRanges"`
	// Destination port ranges.
	DestinationPortRanges QosPortRangeResponseArrayOutput `pulumi:"destinationPortRanges"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// List of markings to be used in the configuration.
	Markings pulumi.IntArrayOutput `pulumi:"markings"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// RNM supported protocol types.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The provisioning state of the DSCP Configuration resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Qos Collection ID generated by RNM.
	QosCollectionId pulumi.StringOutput `pulumi:"qosCollectionId"`
	// QoS object definitions
	QosDefinitionCollection QosDefinitionResponseArrayOutput `pulumi:"qosDefinitionCollection"`
	// The resource GUID property of the DSCP Configuration resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Source IP ranges.
	SourceIpRanges QosIpRangeResponseArrayOutput `pulumi:"sourceIpRanges"`
	// Sources port ranges.
	SourcePortRanges QosPortRangeResponseArrayOutput `pulumi:"sourcePortRanges"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDscpConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDscpConfiguration(ctx *pulumi.Context,
	name string, args *DscpConfigurationArgs, opts ...pulumi.ResourceOption) (*DscpConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:DscpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:DscpConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DscpConfiguration
	err := ctx.RegisterResource("azure-native:network/v20230201:DscpConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDscpConfiguration gets an existing DscpConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDscpConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscpConfigurationState, opts ...pulumi.ResourceOption) (*DscpConfiguration, error) {
	var resource DscpConfiguration
	err := ctx.ReadResource("azure-native:network/v20230201:DscpConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DscpConfiguration resources.
type dscpConfigurationState struct {
}

type DscpConfigurationState struct {
}

func (DscpConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscpConfigurationState)(nil)).Elem()
}

type dscpConfigurationArgs struct {
	// Destination IP ranges.
	DestinationIpRanges []QosIpRange `pulumi:"destinationIpRanges"`
	// Destination port ranges.
	DestinationPortRanges []QosPortRange `pulumi:"destinationPortRanges"`
	// The name of the resource.
	DscpConfigurationName *string `pulumi:"dscpConfigurationName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// List of markings to be used in the configuration.
	Markings []int `pulumi:"markings"`
	// RNM supported protocol types.
	Protocol *string `pulumi:"protocol"`
	// QoS object definitions
	QosDefinitionCollection []QosDefinition `pulumi:"qosDefinitionCollection"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Source IP ranges.
	SourceIpRanges []QosIpRange `pulumi:"sourceIpRanges"`
	// Sources port ranges.
	SourcePortRanges []QosPortRange `pulumi:"sourcePortRanges"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DscpConfiguration resource.
type DscpConfigurationArgs struct {
	// Destination IP ranges.
	DestinationIpRanges QosIpRangeArrayInput
	// Destination port ranges.
	DestinationPortRanges QosPortRangeArrayInput
	// The name of the resource.
	DscpConfigurationName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// List of markings to be used in the configuration.
	Markings pulumi.IntArrayInput
	// RNM supported protocol types.
	Protocol pulumi.StringPtrInput
	// QoS object definitions
	QosDefinitionCollection QosDefinitionArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Source IP ranges.
	SourceIpRanges QosIpRangeArrayInput
	// Sources port ranges.
	SourcePortRanges QosPortRangeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DscpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscpConfigurationArgs)(nil)).Elem()
}

type DscpConfigurationInput interface {
	pulumi.Input

	ToDscpConfigurationOutput() DscpConfigurationOutput
	ToDscpConfigurationOutputWithContext(ctx context.Context) DscpConfigurationOutput
}

func (*DscpConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**DscpConfiguration)(nil)).Elem()
}

func (i *DscpConfiguration) ToDscpConfigurationOutput() DscpConfigurationOutput {
	return i.ToDscpConfigurationOutputWithContext(context.Background())
}

func (i *DscpConfiguration) ToDscpConfigurationOutputWithContext(ctx context.Context) DscpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscpConfigurationOutput)
}

func (i *DscpConfiguration) ToOutput(ctx context.Context) pulumix.Output[*DscpConfiguration] {
	return pulumix.Output[*DscpConfiguration]{
		OutputState: i.ToDscpConfigurationOutputWithContext(ctx).OutputState,
	}
}

type DscpConfigurationOutput struct{ *pulumi.OutputState }

func (DscpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscpConfiguration)(nil)).Elem()
}

func (o DscpConfigurationOutput) ToDscpConfigurationOutput() DscpConfigurationOutput {
	return o
}

func (o DscpConfigurationOutput) ToDscpConfigurationOutputWithContext(ctx context.Context) DscpConfigurationOutput {
	return o
}

func (o DscpConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[*DscpConfiguration] {
	return pulumix.Output[*DscpConfiguration]{
		OutputState: o.OutputState,
	}
}

// Associated Network Interfaces to the DSCP Configuration.
func (o DscpConfigurationOutput) AssociatedNetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) NetworkInterfaceResponseArrayOutput { return v.AssociatedNetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// Destination IP ranges.
func (o DscpConfigurationOutput) DestinationIpRanges() QosIpRangeResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosIpRangeResponseArrayOutput { return v.DestinationIpRanges }).(QosIpRangeResponseArrayOutput)
}

// Destination port ranges.
func (o DscpConfigurationOutput) DestinationPortRanges() QosPortRangeResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosPortRangeResponseArrayOutput { return v.DestinationPortRanges }).(QosPortRangeResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o DscpConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o DscpConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// List of markings to be used in the configuration.
func (o DscpConfigurationOutput) Markings() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.IntArrayOutput { return v.Markings }).(pulumi.IntArrayOutput)
}

// Resource name.
func (o DscpConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// RNM supported protocol types.
func (o DscpConfigurationOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The provisioning state of the DSCP Configuration resource.
func (o DscpConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Qos Collection ID generated by RNM.
func (o DscpConfigurationOutput) QosCollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.QosCollectionId }).(pulumi.StringOutput)
}

// QoS object definitions
func (o DscpConfigurationOutput) QosDefinitionCollection() QosDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosDefinitionResponseArrayOutput { return v.QosDefinitionCollection }).(QosDefinitionResponseArrayOutput)
}

// The resource GUID property of the DSCP Configuration resource.
func (o DscpConfigurationOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Source IP ranges.
func (o DscpConfigurationOutput) SourceIpRanges() QosIpRangeResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosIpRangeResponseArrayOutput { return v.SourceIpRanges }).(QosIpRangeResponseArrayOutput)
}

// Sources port ranges.
func (o DscpConfigurationOutput) SourcePortRanges() QosPortRangeResponseArrayOutput {
	return o.ApplyT(func(v *DscpConfiguration) QosPortRangeResponseArrayOutput { return v.SourcePortRanges }).(QosPortRangeResponseArrayOutput)
}

// Resource tags.
func (o DscpConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o DscpConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DscpConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DscpConfigurationOutput{})
}

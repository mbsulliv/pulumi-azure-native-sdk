// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies information about the dedicated host group that the dedicated hosts should be assigned to. Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.
type DedicatedHostGroup struct {
	pulumi.CustomResourceState

	// Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
	AdditionalCapabilities DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput `pulumi:"additionalCapabilities"`
	// A list of references to all dedicated hosts in the dedicated host group.
	Hosts SubResourceReadOnlyResponseArrayOutput `pulumi:"hosts"`
	// The dedicated host group instance view, which has the list of instance view of the dedicated hosts under the dedicated host group.
	InstanceView DedicatedHostGroupInstanceViewResponseOutput `pulumi:"instanceView"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount pulumi.IntOutput `pulumi:"platformFaultDomainCount"`
	// Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
	SupportAutomaticPlacement pulumi.BoolPtrOutput `pulumi:"supportAutomaticPlacement"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewDedicatedHostGroup registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHostGroup(ctx *pulumi.Context,
	name string, args *DedicatedHostGroupArgs, opts ...pulumi.ResourceOption) (*DedicatedHostGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PlatformFaultDomainCount == nil {
		return nil, errors.New("invalid value for required argument 'PlatformFaultDomainCount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:DedicatedHostGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:DedicatedHostGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource DedicatedHostGroup
	err := ctx.RegisterResource("azure-native:compute/v20230701:DedicatedHostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHostGroup gets an existing DedicatedHostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostGroupState, opts ...pulumi.ResourceOption) (*DedicatedHostGroup, error) {
	var resource DedicatedHostGroup
	err := ctx.ReadResource("azure-native:compute/v20230701:DedicatedHostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHostGroup resources.
type dedicatedHostGroupState struct {
}

type DedicatedHostGroupState struct {
}

func (DedicatedHostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostGroupState)(nil)).Elem()
}

type dedicatedHostGroupArgs struct {
	// Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
	AdditionalCapabilities *DedicatedHostGroupPropertiesAdditionalCapabilities `pulumi:"additionalCapabilities"`
	// The name of the dedicated host group.
	HostGroupName *string `pulumi:"hostGroupName"`
	// Resource location
	Location *string `pulumi:"location"`
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount int `pulumi:"platformFaultDomainCount"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
	SupportAutomaticPlacement *bool `pulumi:"supportAutomaticPlacement"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a DedicatedHostGroup resource.
type DedicatedHostGroupArgs struct {
	// Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
	AdditionalCapabilities DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrInput
	// The name of the dedicated host group.
	HostGroupName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Number of fault domains that the host group can span.
	PlatformFaultDomainCount pulumi.IntInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
	SupportAutomaticPlacement pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
	Zones pulumi.StringArrayInput
}

func (DedicatedHostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostGroupArgs)(nil)).Elem()
}

type DedicatedHostGroupInput interface {
	pulumi.Input

	ToDedicatedHostGroupOutput() DedicatedHostGroupOutput
	ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput
}

func (*DedicatedHostGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostGroup)(nil)).Elem()
}

func (i *DedicatedHostGroup) ToDedicatedHostGroupOutput() DedicatedHostGroupOutput {
	return i.ToDedicatedHostGroupOutputWithContext(context.Background())
}

func (i *DedicatedHostGroup) ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostGroupOutput)
}

type DedicatedHostGroupOutput struct{ *pulumi.OutputState }

func (DedicatedHostGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostGroup)(nil)).Elem()
}

func (o DedicatedHostGroupOutput) ToDedicatedHostGroupOutput() DedicatedHostGroupOutput {
	return o
}

func (o DedicatedHostGroupOutput) ToDedicatedHostGroupOutputWithContext(ctx context.Context) DedicatedHostGroupOutput {
	return o
}

// Enables or disables a capability on the dedicated host group. Minimum api-version: 2022-03-01.
func (o DedicatedHostGroupOutput) AdditionalCapabilities() DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput {
		return v.AdditionalCapabilities
	}).(DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput)
}

// A list of references to all dedicated hosts in the dedicated host group.
func (o DedicatedHostGroupOutput) Hosts() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) SubResourceReadOnlyResponseArrayOutput { return v.Hosts }).(SubResourceReadOnlyResponseArrayOutput)
}

// The dedicated host group instance view, which has the list of instance view of the dedicated hosts under the dedicated host group.
func (o DedicatedHostGroupOutput) InstanceView() DedicatedHostGroupInstanceViewResponseOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) DedicatedHostGroupInstanceViewResponseOutput { return v.InstanceView }).(DedicatedHostGroupInstanceViewResponseOutput)
}

// Resource location
func (o DedicatedHostGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o DedicatedHostGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of fault domains that the host group can span.
func (o DedicatedHostGroupOutput) PlatformFaultDomainCount() pulumi.IntOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) pulumi.IntOutput { return v.PlatformFaultDomainCount }).(pulumi.IntOutput)
}

// Specifies whether virtual machines or virtual machine scale sets can be placed automatically on the dedicated host group. Automatic placement means resources are allocated on dedicated hosts, that are chosen by Azure, under the dedicated host group. The value is defaulted to 'false' when not provided. Minimum api-version: 2020-06-01.
func (o DedicatedHostGroupOutput) SupportAutomaticPlacement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) pulumi.BoolPtrOutput { return v.SupportAutomaticPlacement }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o DedicatedHostGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o DedicatedHostGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
func (o DedicatedHostGroupOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DedicatedHostGroup) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedHostGroupOutput{})
}

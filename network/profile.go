// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Class representing a Traffic Manager profile.
// Azure REST API version: 2022-04-01. Prior API version in Azure Native 1.x: 2018-08-01
type Profile struct {
	pulumi.CustomResourceState

	// The list of allowed endpoint record types.
	AllowedEndpointRecordTypes pulumi.StringArrayOutput `pulumi:"allowedEndpointRecordTypes"`
	// The DNS settings of the Traffic Manager profile.
	DnsConfig DnsConfigResponsePtrOutput `pulumi:"dnsConfig"`
	// The list of endpoints in the Traffic Manager profile.
	Endpoints EndpointResponseArrayOutput `pulumi:"endpoints"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Maximum number of endpoints to be returned for MultiValue routing type.
	MaxReturn pulumi.Float64PtrOutput `pulumi:"maxReturn"`
	// The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig MonitorConfigResponsePtrOutput `pulumi:"monitorConfig"`
	// The name of the resource
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The status of the Traffic Manager profile.
	ProfileStatus pulumi.StringPtrOutput `pulumi:"profileStatus"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod pulumi.StringPtrOutput `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus pulumi.StringPtrOutput `pulumi:"trafficViewEnrollmentStatus"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20151101:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170501:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401:Profile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:Profile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Profile
	err := ctx.RegisterResource("azure-native:network:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure-native:network:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
}

type ProfileState struct {
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// The list of allowed endpoint record types.
	AllowedEndpointRecordTypes []string `pulumi:"allowedEndpointRecordTypes"`
	// The DNS settings of the Traffic Manager profile.
	DnsConfig *DnsConfig `pulumi:"dnsConfig"`
	// The list of endpoints in the Traffic Manager profile.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Endpoints []EndpointType `pulumi:"endpoints"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id *string `pulumi:"id"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// Maximum number of endpoints to be returned for MultiValue routing type.
	MaxReturn *float64 `pulumi:"maxReturn"`
	// The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfig `pulumi:"monitorConfig"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The name of the Traffic Manager profile.
	ProfileName *string `pulumi:"profileName"`
	// The status of the Traffic Manager profile.
	ProfileStatus *string `pulumi:"profileStatus"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod *string `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus *string `pulumi:"trafficViewEnrollmentStatus"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// The list of allowed endpoint record types.
	AllowedEndpointRecordTypes pulumi.StringArrayInput
	// The DNS settings of the Traffic Manager profile.
	DnsConfig DnsConfigPtrInput
	// The list of endpoints in the Traffic Manager profile.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	Endpoints EndpointTypeArrayInput
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// Maximum number of endpoints to be returned for MultiValue routing type.
	MaxReturn pulumi.Float64PtrInput
	// The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig MonitorConfigPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The name of the Traffic Manager profile.
	ProfileName pulumi.StringPtrInput
	// The status of the Traffic Manager profile.
	ProfileStatus pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod pulumi.StringPtrInput
	// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

func (i *Profile) ToOutput(ctx context.Context) pulumix.Output[*Profile] {
	return pulumix.Output[*Profile]{
		OutputState: i.ToProfileOutputWithContext(ctx).OutputState,
	}
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

func (o ProfileOutput) ToOutput(ctx context.Context) pulumix.Output[*Profile] {
	return pulumix.Output[*Profile]{
		OutputState: o.OutputState,
	}
}

// The list of allowed endpoint record types.
func (o ProfileOutput) AllowedEndpointRecordTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringArrayOutput { return v.AllowedEndpointRecordTypes }).(pulumi.StringArrayOutput)
}

// The DNS settings of the Traffic Manager profile.
func (o ProfileOutput) DnsConfig() DnsConfigResponsePtrOutput {
	return o.ApplyT(func(v *Profile) DnsConfigResponsePtrOutput { return v.DnsConfig }).(DnsConfigResponsePtrOutput)
}

// The list of endpoints in the Traffic Manager profile.
func (o ProfileOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v *Profile) EndpointResponseArrayOutput { return v.Endpoints }).(EndpointResponseArrayOutput)
}

// The Azure Region where the resource lives
func (o ProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Maximum number of endpoints to be returned for MultiValue routing type.
func (o ProfileOutput) MaxReturn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.Float64PtrOutput { return v.MaxReturn }).(pulumi.Float64PtrOutput)
}

// The endpoint monitoring settings of the Traffic Manager profile.
func (o ProfileOutput) MonitorConfig() MonitorConfigResponsePtrOutput {
	return o.ApplyT(func(v *Profile) MonitorConfigResponsePtrOutput { return v.MonitorConfig }).(MonitorConfigResponsePtrOutput)
}

// The name of the resource
func (o ProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The status of the Traffic Manager profile.
func (o ProfileOutput) ProfileStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.ProfileStatus }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o ProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The traffic routing method of the Traffic Manager profile.
func (o ProfileOutput) TrafficRoutingMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.TrafficRoutingMethod }).(pulumi.StringPtrOutput)
}

// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
func (o ProfileOutput) TrafficViewEnrollmentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.TrafficViewEnrollmentStatus }).(pulumi.StringPtrOutput)
}

// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
func (o ProfileOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfileOutput{})
}

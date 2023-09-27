// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
type OriginGroup struct {
	pulumi.CustomResourceState

	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings HealthProbeParametersResponsePtrOutput `pulumi:"healthProbeSettings"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The source of the content being delivered via CDN within given origin group.
	Origins ResourceReferenceResponseArrayOutput `pulumi:"origins"`
	// Provisioning status of the origin group.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource status of the origin group.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
	ResponseBasedOriginErrorDetectionSettings ResponseBasedOriginErrorDetectionParametersResponsePtrOutput `pulumi:"responseBasedOriginErrorDetectionSettings"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrOutput `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOriginGroup registers a new resource with the given unique name, arguments, and options.
func NewOriginGroup(ctx *pulumi.Context,
	name string, args *OriginGroupArgs, opts ...pulumi.ResourceOption) (*OriginGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.Origins == nil {
		return nil, errors.New("invalid value for required argument 'Origins'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230501:OriginGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OriginGroup
	err := ctx.RegisterResource("azure-native:cdn/v20230701preview:OriginGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOriginGroup gets an existing OriginGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOriginGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginGroupState, opts ...pulumi.ResourceOption) (*OriginGroup, error) {
	var resource OriginGroup
	err := ctx.ReadResource("azure-native:cdn/v20230701preview:OriginGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OriginGroup resources.
type originGroupState struct {
}

type OriginGroupState struct {
}

func (OriginGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*originGroupState)(nil)).Elem()
}

type originGroupArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings *HealthProbeParameters `pulumi:"healthProbeSettings"`
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName *string `pulumi:"originGroupName"`
	// The source of the content being delivered via CDN within given origin group.
	Origins []ResourceReference `pulumi:"origins"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
	ResponseBasedOriginErrorDetectionSettings *ResponseBasedOriginErrorDetectionParameters `pulumi:"responseBasedOriginErrorDetectionSettings"`
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
}

// The set of arguments for constructing a OriginGroup resource.
type OriginGroupArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput
	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings HealthProbeParametersPtrInput
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName pulumi.StringPtrInput
	// The source of the content being delivered via CDN within given origin group.
	Origins ResourceReferenceArrayInput
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
	ResponseBasedOriginErrorDetectionSettings ResponseBasedOriginErrorDetectionParametersPtrInput
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrInput
}

func (OriginGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originGroupArgs)(nil)).Elem()
}

type OriginGroupInput interface {
	pulumi.Input

	ToOriginGroupOutput() OriginGroupOutput
	ToOriginGroupOutputWithContext(ctx context.Context) OriginGroupOutput
}

func (*OriginGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginGroup)(nil)).Elem()
}

func (i *OriginGroup) ToOriginGroupOutput() OriginGroupOutput {
	return i.ToOriginGroupOutputWithContext(context.Background())
}

func (i *OriginGroup) ToOriginGroupOutputWithContext(ctx context.Context) OriginGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginGroupOutput)
}

func (i *OriginGroup) ToOutput(ctx context.Context) pulumix.Output[*OriginGroup] {
	return pulumix.Output[*OriginGroup]{
		OutputState: i.ToOriginGroupOutputWithContext(ctx).OutputState,
	}
}

type OriginGroupOutput struct{ *pulumi.OutputState }

func (OriginGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginGroup)(nil)).Elem()
}

func (o OriginGroupOutput) ToOriginGroupOutput() OriginGroupOutput {
	return o
}

func (o OriginGroupOutput) ToOriginGroupOutputWithContext(ctx context.Context) OriginGroupOutput {
	return o
}

func (o OriginGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*OriginGroup] {
	return pulumix.Output[*OriginGroup]{
		OutputState: o.OutputState,
	}
}

// Health probe settings to the origin that is used to determine the health of the origin.
func (o OriginGroupOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v *OriginGroup) HealthProbeParametersResponsePtrOutput { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

// Resource name.
func (o OriginGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The source of the content being delivered via CDN within given origin group.
func (o OriginGroupOutput) Origins() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *OriginGroup) ResourceReferenceResponseArrayOutput { return v.Origins }).(ResourceReferenceResponseArrayOutput)
}

// Provisioning status of the origin group.
func (o OriginGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource status of the origin group.
func (o OriginGroupOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginGroup) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
func (o OriginGroupOutput) ResponseBasedOriginErrorDetectionSettings() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o.ApplyT(func(v *OriginGroup) ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
		return v.ResponseBasedOriginErrorDetectionSettings
	}).(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput)
}

// Read only system data
func (o OriginGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OriginGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
func (o OriginGroupOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OriginGroup) pulumi.IntPtrOutput {
		return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes
	}).(pulumi.IntPtrOutput)
}

// Resource type.
func (o OriginGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OriginGroupOutput{})
}

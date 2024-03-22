// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AFDOrigin group comprising of origins is used for load balancing to origins when the content cannot be served from Azure Front Door.
// Azure REST API version: 2023-05-01. Prior API version in Azure Native 1.x: 2020-09-01.
//
// Other available API versions: 2020-09-01, 2023-07-01-preview, 2024-02-01.
type AFDOriginGroup struct {
	pulumi.CustomResourceState

	DeploymentStatus pulumi.StringOutput `pulumi:"deploymentStatus"`
	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings HealthProbeParametersResponsePtrOutput `pulumi:"healthProbeSettings"`
	// Load balancing settings for a backend pool
	LoadBalancingSettings LoadBalancingSettingsParametersResponsePtrOutput `pulumi:"loadBalancingSettings"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the profile which holds the origin group.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// Provisioning status
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
	SessionAffinityState pulumi.StringPtrOutput `pulumi:"sessionAffinityState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrOutput `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAFDOriginGroup registers a new resource with the given unique name, arguments, and options.
func NewAFDOriginGroup(ctx *pulumi.Context,
	name string, args *AFDOriginGroupArgs, opts ...pulumi.ResourceOption) (*AFDOriginGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn/v20200901:AFDOriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:AFDOriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:AFDOriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:AFDOriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230501:AFDOriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230701preview:AFDOriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20240201:AFDOriginGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AFDOriginGroup
	err := ctx.RegisterResource("azure-native:cdn:AFDOriginGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAFDOriginGroup gets an existing AFDOriginGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAFDOriginGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDOriginGroupState, opts ...pulumi.ResourceOption) (*AFDOriginGroup, error) {
	var resource AFDOriginGroup
	err := ctx.ReadResource("azure-native:cdn:AFDOriginGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AFDOriginGroup resources.
type afdoriginGroupState struct {
}

type AFDOriginGroupState struct {
}

func (AFDOriginGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginGroupState)(nil)).Elem()
}

type afdoriginGroupArgs struct {
	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings *HealthProbeParameters `pulumi:"healthProbeSettings"`
	// Load balancing settings for a backend pool
	LoadBalancingSettings *LoadBalancingSettingsParameters `pulumi:"loadBalancingSettings"`
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName *string `pulumi:"originGroupName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
	SessionAffinityState *string `pulumi:"sessionAffinityState"`
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
}

// The set of arguments for constructing a AFDOriginGroup resource.
type AFDOriginGroupArgs struct {
	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings HealthProbeParametersPtrInput
	// Load balancing settings for a backend pool
	LoadBalancingSettings LoadBalancingSettingsParametersPtrInput
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName pulumi.StringPtrInput
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
	SessionAffinityState pulumi.StringPtrInput
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrInput
}

func (AFDOriginGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginGroupArgs)(nil)).Elem()
}

type AFDOriginGroupInput interface {
	pulumi.Input

	ToAFDOriginGroupOutput() AFDOriginGroupOutput
	ToAFDOriginGroupOutputWithContext(ctx context.Context) AFDOriginGroupOutput
}

func (*AFDOriginGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDOriginGroup)(nil)).Elem()
}

func (i *AFDOriginGroup) ToAFDOriginGroupOutput() AFDOriginGroupOutput {
	return i.ToAFDOriginGroupOutputWithContext(context.Background())
}

func (i *AFDOriginGroup) ToAFDOriginGroupOutputWithContext(ctx context.Context) AFDOriginGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDOriginGroupOutput)
}

type AFDOriginGroupOutput struct{ *pulumi.OutputState }

func (AFDOriginGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDOriginGroup)(nil)).Elem()
}

func (o AFDOriginGroupOutput) ToAFDOriginGroupOutput() AFDOriginGroupOutput {
	return o
}

func (o AFDOriginGroupOutput) ToAFDOriginGroupOutputWithContext(ctx context.Context) AFDOriginGroupOutput {
	return o
}

func (o AFDOriginGroupOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Health probe settings to the origin that is used to determine the health of the origin.
func (o AFDOriginGroupOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) HealthProbeParametersResponsePtrOutput { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

// Load balancing settings for a backend pool
func (o AFDOriginGroupOutput) LoadBalancingSettings() LoadBalancingSettingsParametersResponsePtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) LoadBalancingSettingsParametersResponsePtrOutput {
		return v.LoadBalancingSettings
	}).(LoadBalancingSettingsParametersResponsePtrOutput)
}

// Resource name.
func (o AFDOriginGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the profile which holds the origin group.
func (o AFDOriginGroupOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

// Provisioning status
func (o AFDOriginGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
func (o AFDOriginGroupOutput) SessionAffinityState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringPtrOutput { return v.SessionAffinityState }).(pulumi.StringPtrOutput)
}

// Read only system data
func (o AFDOriginGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AFDOriginGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
func (o AFDOriginGroupOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.IntPtrOutput {
		return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes
	}).(pulumi.IntPtrOutput)
}

// Resource type.
func (o AFDOriginGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AFDOriginGroupOutput{})
}

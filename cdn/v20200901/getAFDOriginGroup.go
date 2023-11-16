// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing origin group within a profile.
func LookupAFDOriginGroup(ctx *pulumi.Context, args *LookupAFDOriginGroupArgs, opts ...pulumi.InvokeOption) (*LookupAFDOriginGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAFDOriginGroupResult
	err := ctx.Invoke("azure-native:cdn/v20200901:getAFDOriginGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDOriginGroupArgs struct {
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName string `pulumi:"originGroupName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// AFDOrigin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
type LookupAFDOriginGroupResult struct {
	DeploymentStatus string `pulumi:"deploymentStatus"`
	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings *HealthProbeParametersResponse `pulumi:"healthProbeSettings"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Load balancing settings for a backend pool
	LoadBalancingSettings *LoadBalancingSettingsParametersResponse `pulumi:"loadBalancingSettings"`
	// Resource name.
	Name string `pulumi:"name"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
	ResponseBasedAfdOriginErrorDetectionSettings *ResponseBasedOriginErrorDetectionParametersResponse `pulumi:"responseBasedAfdOriginErrorDetectionSettings"`
	// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
	SessionAffinityState *string `pulumi:"sessionAffinityState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupAFDOriginGroupOutput(ctx *pulumi.Context, args LookupAFDOriginGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAFDOriginGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAFDOriginGroupResult, error) {
			args := v.(LookupAFDOriginGroupArgs)
			r, err := LookupAFDOriginGroup(ctx, &args, opts...)
			var s LookupAFDOriginGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAFDOriginGroupResultOutput)
}

type LookupAFDOriginGroupOutputArgs struct {
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName pulumi.StringInput `pulumi:"originGroupName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAFDOriginGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDOriginGroupArgs)(nil)).Elem()
}

// AFDOrigin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
type LookupAFDOriginGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAFDOriginGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDOriginGroupResult)(nil)).Elem()
}

func (o LookupAFDOriginGroupResultOutput) ToLookupAFDOriginGroupResultOutput() LookupAFDOriginGroupResultOutput {
	return o
}

func (o LookupAFDOriginGroupResultOutput) ToLookupAFDOriginGroupResultOutputWithContext(ctx context.Context) LookupAFDOriginGroupResultOutput {
	return o
}

func (o LookupAFDOriginGroupResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Health probe settings to the origin that is used to determine the health of the origin.
func (o LookupAFDOriginGroupResultOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *HealthProbeParametersResponse { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

// Resource ID.
func (o LookupAFDOriginGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Load balancing settings for a backend pool
func (o LookupAFDOriginGroupResultOutput) LoadBalancingSettings() LoadBalancingSettingsParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *LoadBalancingSettingsParametersResponse {
		return v.LoadBalancingSettings
	}).(LoadBalancingSettingsParametersResponsePtrOutput)
}

// Resource name.
func (o LookupAFDOriginGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning status
func (o LookupAFDOriginGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
func (o LookupAFDOriginGroupResultOutput) ResponseBasedAfdOriginErrorDetectionSettings() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *ResponseBasedOriginErrorDetectionParametersResponse {
		return v.ResponseBasedAfdOriginErrorDetectionSettings
	}).(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput)
}

// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
func (o LookupAFDOriginGroupResultOutput) SessionAffinityState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *string { return v.SessionAffinityState }).(pulumi.StringPtrOutput)
}

// Read only system data
func (o LookupAFDOriginGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
func (o LookupAFDOriginGroupResultOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *int {
		return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes
	}).(pulumi.IntPtrOutput)
}

// Resource type.
func (o LookupAFDOriginGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAFDOriginGroupResultOutput{})
}

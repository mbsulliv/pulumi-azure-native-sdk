// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing origin group within an endpoint.
func LookupOriginGroup(ctx *pulumi.Context, args *LookupOriginGroupArgs, opts ...pulumi.InvokeOption) (*LookupOriginGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOriginGroupResult
	err := ctx.Invoke("azure-native:cdn/v20240201:getOriginGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOriginGroupArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName string `pulumi:"originGroupName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
type LookupOriginGroupResult struct {
	// Health probe settings to the origin that is used to determine the health of the origin.
	HealthProbeSettings *HealthProbeParametersResponse `pulumi:"healthProbeSettings"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The source of the content being delivered via CDN within given origin group.
	Origins []ResourceReferenceResponse `pulumi:"origins"`
	// Provisioning status of the origin group.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource status of the origin group.
	ResourceState string `pulumi:"resourceState"`
	// The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
	ResponseBasedOriginErrorDetectionSettings *ResponseBasedOriginErrorDetectionParametersResponse `pulumi:"responseBasedOriginErrorDetectionSettings"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupOriginGroupOutput(ctx *pulumi.Context, args LookupOriginGroupOutputArgs, opts ...pulumi.InvokeOption) LookupOriginGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOriginGroupResult, error) {
			args := v.(LookupOriginGroupArgs)
			r, err := LookupOriginGroup(ctx, &args, opts...)
			var s LookupOriginGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOriginGroupResultOutput)
}

type LookupOriginGroupOutputArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// Name of the origin group which is unique within the endpoint.
	OriginGroupName pulumi.StringInput `pulumi:"originGroupName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOriginGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOriginGroupArgs)(nil)).Elem()
}

// Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
type LookupOriginGroupResultOutput struct{ *pulumi.OutputState }

func (LookupOriginGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOriginGroupResult)(nil)).Elem()
}

func (o LookupOriginGroupResultOutput) ToLookupOriginGroupResultOutput() LookupOriginGroupResultOutput {
	return o
}

func (o LookupOriginGroupResultOutput) ToLookupOriginGroupResultOutputWithContext(ctx context.Context) LookupOriginGroupResultOutput {
	return o
}

// Health probe settings to the origin that is used to determine the health of the origin.
func (o LookupOriginGroupResultOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) *HealthProbeParametersResponse { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

// Resource ID.
func (o LookupOriginGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupOriginGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The source of the content being delivered via CDN within given origin group.
func (o LookupOriginGroupResultOutput) Origins() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) []ResourceReferenceResponse { return v.Origins }).(ResourceReferenceResponseArrayOutput)
}

// Provisioning status of the origin group.
func (o LookupOriginGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource status of the origin group.
func (o LookupOriginGroupResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
func (o LookupOriginGroupResultOutput) ResponseBasedOriginErrorDetectionSettings() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) *ResponseBasedOriginErrorDetectionParametersResponse {
		return v.ResponseBasedOriginErrorDetectionSettings
	}).(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput)
}

// Read only system data
func (o LookupOriginGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
func (o LookupOriginGroupResultOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) *int { return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o LookupOriginGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOriginGroupResultOutput{})
}

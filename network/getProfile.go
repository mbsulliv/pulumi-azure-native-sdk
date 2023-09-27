// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Traffic Manager profile.
// Azure REST API version: 2022-04-01.
func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProfileResult
	err := ctx.Invoke("azure-native:network:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	// The name of the Traffic Manager profile.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a Traffic Manager profile.
type LookupProfileResult struct {
	// The list of allowed endpoint record types.
	AllowedEndpointRecordTypes []string `pulumi:"allowedEndpointRecordTypes"`
	// The DNS settings of the Traffic Manager profile.
	DnsConfig *DnsConfigResponse `pulumi:"dnsConfig"`
	// The list of endpoints in the Traffic Manager profile.
	Endpoints []EndpointResponse `pulumi:"endpoints"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id *string `pulumi:"id"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// Maximum number of endpoints to be returned for MultiValue routing type.
	MaxReturn *float64 `pulumi:"maxReturn"`
	// The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfigResponse `pulumi:"monitorConfig"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The status of the Traffic Manager profile.
	ProfileStatus *string `pulumi:"profileStatus"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod *string `pulumi:"trafficRoutingMethod"`
	// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus *string `pulumi:"trafficViewEnrollmentStatus"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResult, error) {
			args := v.(LookupProfileArgs)
			r, err := LookupProfile(ctx, &args, opts...)
			var s LookupProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfileResultOutput)
}

type LookupProfileOutputArgs struct {
	// The name of the Traffic Manager profile.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}

// Class representing a Traffic Manager profile.
type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupProfileResult] {
	return pulumix.Output[LookupProfileResult]{
		OutputState: o.OutputState,
	}
}

// The list of allowed endpoint record types.
func (o LookupProfileResultOutput) AllowedEndpointRecordTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []string { return v.AllowedEndpointRecordTypes }).(pulumi.StringArrayOutput)
}

// The DNS settings of the Traffic Manager profile.
func (o LookupProfileResultOutput) DnsConfig() DnsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *DnsConfigResponse { return v.DnsConfig }).(DnsConfigResponsePtrOutput)
}

// The list of endpoints in the Traffic Manager profile.
func (o LookupProfileResultOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []EndpointResponse { return v.Endpoints }).(EndpointResponseArrayOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
func (o LookupProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The Azure Region where the resource lives
func (o LookupProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Maximum number of endpoints to be returned for MultiValue routing type.
func (o LookupProfileResultOutput) MaxReturn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *float64 { return v.MaxReturn }).(pulumi.Float64PtrOutput)
}

// The endpoint monitoring settings of the Traffic Manager profile.
func (o LookupProfileResultOutput) MonitorConfig() MonitorConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *MonitorConfigResponse { return v.MonitorConfig }).(MonitorConfigResponsePtrOutput)
}

// The name of the resource
func (o LookupProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The status of the Traffic Manager profile.
func (o LookupProfileResultOutput) ProfileStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.ProfileStatus }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The traffic routing method of the Traffic Manager profile.
func (o LookupProfileResultOutput) TrafficRoutingMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.TrafficRoutingMethod }).(pulumi.StringPtrOutput)
}

// Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
func (o LookupProfileResultOutput) TrafficViewEnrollmentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.TrafficViewEnrollmentStatus }).(pulumi.StringPtrOutput)
}

// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
func (o LookupProfileResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}

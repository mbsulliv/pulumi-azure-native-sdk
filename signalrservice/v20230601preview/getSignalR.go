// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the resource and its properties.
func LookupSignalR(ctx *pulumi.Context, args *LookupSignalRArgs, opts ...pulumi.InvokeOption) (*LookupSignalRResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSignalRResult
	err := ctx.Invoke("azure-native:signalrservice/v20230601preview:getSignalR", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSignalRArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A class represent a resource.
type LookupSignalRResult struct {
	// Cross-Origin Resource Sharing (CORS) settings.
	Cors *SignalRCorsSettingsResponse `pulumi:"cors"`
	// DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth *bool `pulumi:"disableAadAuth"`
	// DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// The publicly accessible IP of the resource.
	ExternalIP string `pulumi:"externalIP"`
	// List of the featureFlags.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, its globally default value will be used
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features []SignalRFeatureResponse `pulumi:"features"`
	// FQDN of the service instance.
	HostName string `pulumi:"hostName"`
	// Deprecated.
	HostNamePrefix string `pulumi:"hostNamePrefix"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// A class represent managed identities used for request and response
	Identity *ManagedIdentityResponse `pulumi:"identity"`
	// The kind of the service
	Kind *string `pulumi:"kind"`
	// Live trace configuration of a Microsoft.SignalRService resource.
	LiveTraceConfiguration *LiveTraceConfigurationResponse `pulumi:"liveTraceConfiguration"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Network ACLs for the resource
	NetworkACLs *SignalRNetworkACLsResponse `pulumi:"networkACLs"`
	// Private endpoint connections to the resource.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Enable or disable public network access. Default to "Enabled".
	// When it's Enabled, network ACLs still apply.
	// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The publicly accessible port of the resource which is designed for browser/client side usage.
	PublicPort int `pulumi:"publicPort"`
	// Resource log configuration of a Microsoft.SignalRService resource.
	ResourceLogConfiguration *ResourceLogConfigurationResponse `pulumi:"resourceLogConfiguration"`
	// The publicly accessible port of the resource which is designed for customer server side usage.
	ServerPort int `pulumi:"serverPort"`
	// Serverless settings.
	Serverless *ServerlessSettingsResponse `pulumi:"serverless"`
	// The list of shared private link resources.
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	// The billing information of the resource.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// TLS settings for the resource
	Tls *SignalRTlsSettingsResponse `pulumi:"tls"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The settings for the Upstream when the service is in server-less mode.
	Upstream *ServerlessUpstreamSettingsResponse `pulumi:"upstream"`
	// Version of the resource. Probably you need the same or higher version of client SDKs.
	Version string `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupSignalRResult
func (val *LookupSignalRResult) Defaults() *LookupSignalRResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DisableAadAuth == nil {
		disableAadAuth_ := false
		tmp.DisableAadAuth = &disableAadAuth_
	}
	if tmp.DisableLocalAuth == nil {
		disableLocalAuth_ := false
		tmp.DisableLocalAuth = &disableLocalAuth_
	}
	tmp.LiveTraceConfiguration = tmp.LiveTraceConfiguration.Defaults()

	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	tmp.Serverless = tmp.Serverless.Defaults()

	tmp.Tls = tmp.Tls.Defaults()

	return &tmp
}

func LookupSignalROutput(ctx *pulumi.Context, args LookupSignalROutputArgs, opts ...pulumi.InvokeOption) LookupSignalRResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRResult, error) {
			args := v.(LookupSignalRArgs)
			r, err := LookupSignalR(ctx, &args, opts...)
			var s LookupSignalRResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRResultOutput)
}

type LookupSignalROutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalROutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRArgs)(nil)).Elem()
}

// A class represent a resource.
type LookupSignalRResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRResult)(nil)).Elem()
}

func (o LookupSignalRResultOutput) ToLookupSignalRResultOutput() LookupSignalRResultOutput {
	return o
}

func (o LookupSignalRResultOutput) ToLookupSignalRResultOutputWithContext(ctx context.Context) LookupSignalRResultOutput {
	return o
}

// Cross-Origin Resource Sharing (CORS) settings.
func (o LookupSignalRResultOutput) Cors() SignalRCorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *SignalRCorsSettingsResponse { return v.Cors }).(SignalRCorsSettingsResponsePtrOutput)
}

// DisableLocalAuth
// Enable or disable aad auth
// When set as true, connection with AuthType=aad won't work.
func (o LookupSignalRResultOutput) DisableAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *bool { return v.DisableAadAuth }).(pulumi.BoolPtrOutput)
}

// DisableLocalAuth
// Enable or disable local auth with AccessKey
// When set as true, connection with AccessKey=xxx won't work.
func (o LookupSignalRResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// The publicly accessible IP of the resource.
func (o LookupSignalRResultOutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.ExternalIP }).(pulumi.StringOutput)
}

// List of the featureFlags.
//
// FeatureFlags that are not included in the parameters for the update operation will not be modified.
// And the response will only include featureFlags that are explicitly set.
// When a featureFlag is not explicitly set, its globally default value will be used
// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
func (o LookupSignalRResultOutput) Features() SignalRFeatureResponseArrayOutput {
	return o.ApplyT(func(v LookupSignalRResult) []SignalRFeatureResponse { return v.Features }).(SignalRFeatureResponseArrayOutput)
}

// FQDN of the service instance.
func (o LookupSignalRResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Deprecated.
func (o LookupSignalRResultOutput) HostNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.HostNamePrefix }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSignalRResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Id }).(pulumi.StringOutput)
}

// A class represent managed identities used for request and response
func (o LookupSignalRResultOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ManagedIdentityResponse { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// The kind of the service
func (o LookupSignalRResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Live trace configuration of a Microsoft.SignalRService resource.
func (o LookupSignalRResultOutput) LiveTraceConfiguration() LiveTraceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *LiveTraceConfigurationResponse { return v.LiveTraceConfiguration }).(LiveTraceConfigurationResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupSignalRResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSignalRResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network ACLs for the resource
func (o LookupSignalRResultOutput) NetworkACLs() SignalRNetworkACLsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *SignalRNetworkACLsResponse { return v.NetworkACLs }).(SignalRNetworkACLsResponsePtrOutput)
}

// Private endpoint connections to the resource.
func (o LookupSignalRResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupSignalRResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the resource.
func (o LookupSignalRResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Enable or disable public network access. Default to "Enabled".
// When it's Enabled, network ACLs still apply.
// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
func (o LookupSignalRResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The publicly accessible port of the resource which is designed for browser/client side usage.
func (o LookupSignalRResultOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSignalRResult) int { return v.PublicPort }).(pulumi.IntOutput)
}

// Resource log configuration of a Microsoft.SignalRService resource.
func (o LookupSignalRResultOutput) ResourceLogConfiguration() ResourceLogConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ResourceLogConfigurationResponse { return v.ResourceLogConfiguration }).(ResourceLogConfigurationResponsePtrOutput)
}

// The publicly accessible port of the resource which is designed for customer server side usage.
func (o LookupSignalRResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSignalRResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

// Serverless settings.
func (o LookupSignalRResultOutput) Serverless() ServerlessSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ServerlessSettingsResponse { return v.Serverless }).(ServerlessSettingsResponsePtrOutput)
}

// The list of shared private link resources.
func (o LookupSignalRResultOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupSignalRResult) []SharedPrivateLinkResourceResponse { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

// The billing information of the resource.
func (o LookupSignalRResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSignalRResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSignalRResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSignalRResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSignalRResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// TLS settings for the resource
func (o LookupSignalRResultOutput) Tls() SignalRTlsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *SignalRTlsSettingsResponse { return v.Tls }).(SignalRTlsSettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSignalRResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Type }).(pulumi.StringOutput)
}

// The settings for the Upstream when the service is in server-less mode.
func (o LookupSignalRResultOutput) Upstream() ServerlessUpstreamSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ServerlessUpstreamSettingsResponse { return v.Upstream }).(ServerlessUpstreamSettingsResponsePtrOutput)
}

// Version of the resource. Probably you need the same or higher version of client SDKs.
func (o LookupSignalRResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRResultOutput{})
}

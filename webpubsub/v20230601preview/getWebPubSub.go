// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the resource and its properties.
func LookupWebPubSub(ctx *pulumi.Context, args *LookupWebPubSubArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebPubSubResult
	err := ctx.Invoke("azure-native:webpubsub/v20230601preview:getWebPubSub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebPubSubArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A class represent a resource.
type LookupWebPubSubResult struct {
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
	NetworkACLs *WebPubSubNetworkACLsResponse `pulumi:"networkACLs"`
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
	// The list of shared private link resources.
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	// The billing information of the resource.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// TLS settings for the resource
	Tls *WebPubSubTlsSettingsResponse `pulumi:"tls"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Version of the resource. Probably you need the same or higher version of client SDKs.
	Version string `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupWebPubSubResult
func (val *LookupWebPubSubResult) Defaults() *LookupWebPubSubResult {
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
	tmp.Tls = tmp.Tls.Defaults()

	return &tmp
}

func LookupWebPubSubOutput(ctx *pulumi.Context, args LookupWebPubSubOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubResult, error) {
			args := v.(LookupWebPubSubArgs)
			r, err := LookupWebPubSub(ctx, &args, opts...)
			var s LookupWebPubSubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubResultOutput)
}

type LookupWebPubSubOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubArgs)(nil)).Elem()
}

// A class represent a resource.
type LookupWebPubSubResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubResult)(nil)).Elem()
}

func (o LookupWebPubSubResultOutput) ToLookupWebPubSubResultOutput() LookupWebPubSubResultOutput {
	return o
}

func (o LookupWebPubSubResultOutput) ToLookupWebPubSubResultOutputWithContext(ctx context.Context) LookupWebPubSubResultOutput {
	return o
}

func (o LookupWebPubSubResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebPubSubResult] {
	return pulumix.Output[LookupWebPubSubResult]{
		OutputState: o.OutputState,
	}
}

// DisableLocalAuth
// Enable or disable aad auth
// When set as true, connection with AuthType=aad won't work.
func (o LookupWebPubSubResultOutput) DisableAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *bool { return v.DisableAadAuth }).(pulumi.BoolPtrOutput)
}

// DisableLocalAuth
// Enable or disable local auth with AccessKey
// When set as true, connection with AccessKey=xxx won't work.
func (o LookupWebPubSubResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// The publicly accessible IP of the resource.
func (o LookupWebPubSubResultOutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.ExternalIP }).(pulumi.StringOutput)
}

// FQDN of the service instance.
func (o LookupWebPubSubResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Deprecated.
func (o LookupWebPubSubResultOutput) HostNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.HostNamePrefix }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupWebPubSubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Id }).(pulumi.StringOutput)
}

// A class represent managed identities used for request and response
func (o LookupWebPubSubResultOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *ManagedIdentityResponse { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// The kind of the service
func (o LookupWebPubSubResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Live trace configuration of a Microsoft.SignalRService resource.
func (o LookupWebPubSubResultOutput) LiveTraceConfiguration() LiveTraceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *LiveTraceConfigurationResponse { return v.LiveTraceConfiguration }).(LiveTraceConfigurationResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupWebPubSubResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWebPubSubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network ACLs for the resource
func (o LookupWebPubSubResultOutput) NetworkACLs() WebPubSubNetworkACLsResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *WebPubSubNetworkACLsResponse { return v.NetworkACLs }).(WebPubSubNetworkACLsResponsePtrOutput)
}

// Private endpoint connections to the resource.
func (o LookupWebPubSubResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the resource.
func (o LookupWebPubSubResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Enable or disable public network access. Default to "Enabled".
// When it's Enabled, network ACLs still apply.
// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
func (o LookupWebPubSubResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The publicly accessible port of the resource which is designed for browser/client side usage.
func (o LookupWebPubSubResultOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) int { return v.PublicPort }).(pulumi.IntOutput)
}

// Resource log configuration of a Microsoft.SignalRService resource.
func (o LookupWebPubSubResultOutput) ResourceLogConfiguration() ResourceLogConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *ResourceLogConfigurationResponse { return v.ResourceLogConfiguration }).(ResourceLogConfigurationResponsePtrOutput)
}

// The publicly accessible port of the resource which is designed for customer server side usage.
func (o LookupWebPubSubResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

// The list of shared private link resources.
func (o LookupWebPubSubResultOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) []SharedPrivateLinkResourceResponse { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

// The billing information of the resource.
func (o LookupWebPubSubResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWebPubSubResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupWebPubSubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// TLS settings for the resource
func (o LookupWebPubSubResultOutput) Tls() WebPubSubTlsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) *WebPubSubTlsSettingsResponse { return v.Tls }).(WebPubSubTlsSettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWebPubSubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version of the resource. Probably you need the same or higher version of client SDKs.
func (o LookupWebPubSubResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubResultOutput{})
}

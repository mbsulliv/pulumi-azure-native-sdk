// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A class represent a resource.
type WebPubSub struct {
	pulumi.CustomResourceState

	// DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth pulumi.BoolPtrOutput `pulumi:"disableAadAuth"`
	// DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// The publicly accessible IP of the resource.
	ExternalIP pulumi.StringOutput `pulumi:"externalIP"`
	// FQDN of the service instance.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Deprecated.
	HostNamePrefix pulumi.StringOutput `pulumi:"hostNamePrefix"`
	// A class represent managed identities used for request and response
	Identity ManagedIdentityResponsePtrOutput `pulumi:"identity"`
	// The kind of the service
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Live trace configuration of a Microsoft.SignalRService resource.
	LiveTraceConfiguration LiveTraceConfigurationResponsePtrOutput `pulumi:"liveTraceConfiguration"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Network ACLs for the resource
	NetworkACLs WebPubSubNetworkACLsResponsePtrOutput `pulumi:"networkACLs"`
	// Private endpoint connections to the resource.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Enable or disable public network access. Default to "Enabled".
	// When it's Enabled, network ACLs still apply.
	// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The publicly accessible port of the resource which is designed for browser/client side usage.
	PublicPort pulumi.IntOutput `pulumi:"publicPort"`
	// Enable or disable the regional endpoint. Default to "Enabled".
	// When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
	// This property is replica specific. Disable the regional endpoint without replica is not allowed.
	RegionEndpointEnabled pulumi.StringPtrOutput `pulumi:"regionEndpointEnabled"`
	// Resource log configuration of a Microsoft.SignalRService resource.
	ResourceLogConfiguration ResourceLogConfigurationResponsePtrOutput `pulumi:"resourceLogConfiguration"`
	// Stop or start the resource.  Default to "false".
	// When it's true, the data plane of the resource is shutdown.
	// When it's false, the data plane of the resource is started.
	ResourceStopped pulumi.StringPtrOutput `pulumi:"resourceStopped"`
	// The publicly accessible port of the resource which is designed for customer server side usage.
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// The list of shared private link resources.
	SharedPrivateLinkResources SharedPrivateLinkResourceResponseArrayOutput `pulumi:"sharedPrivateLinkResources"`
	// The billing information of the resource.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// TLS settings for the resource
	Tls WebPubSubTlsSettingsResponsePtrOutput `pulumi:"tls"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the resource. Probably you need the same or higher version of client SDKs.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewWebPubSub registers a new resource with the given unique name, arguments, and options.
func NewWebPubSub(ctx *pulumi.Context,
	name string, args *WebPubSubArgs, opts ...pulumi.ResourceOption) (*WebPubSub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DisableAadAuth == nil {
		args.DisableAadAuth = pulumi.BoolPtr(false)
	}
	if args.DisableLocalAuth == nil {
		args.DisableLocalAuth = pulumi.BoolPtr(false)
	}
	if args.LiveTraceConfiguration != nil {
		args.LiveTraceConfiguration = args.LiveTraceConfiguration.ToLiveTraceConfigurationPtrOutput().ApplyT(func(v *LiveTraceConfiguration) *LiveTraceConfiguration { return v.Defaults() }).(LiveTraceConfigurationPtrOutput)
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	if args.RegionEndpointEnabled == nil {
		args.RegionEndpointEnabled = pulumi.StringPtr("Enabled")
	}
	if args.ResourceStopped == nil {
		args.ResourceStopped = pulumi.StringPtr("false")
	}
	if args.Tls != nil {
		args.Tls = args.Tls.ToWebPubSubTlsSettingsPtrOutput().ApplyT(func(v *WebPubSubTlsSettings) *WebPubSubTlsSettings { return v.Defaults() }).(WebPubSubTlsSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:webpubsub:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210401preview:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210601preview:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210901preview:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20211001:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20220801preview:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230201:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230301preview:WebPubSub"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230601preview:WebPubSub"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebPubSub
	err := ctx.RegisterResource("azure-native:webpubsub/v20230801preview:WebPubSub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebPubSub gets an existing WebPubSub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebPubSub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubState, opts ...pulumi.ResourceOption) (*WebPubSub, error) {
	var resource WebPubSub
	err := ctx.ReadResource("azure-native:webpubsub/v20230801preview:WebPubSub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebPubSub resources.
type webPubSubState struct {
}

type WebPubSubState struct {
}

func (WebPubSubState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubState)(nil)).Elem()
}

type webPubSubArgs struct {
	// DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth *bool `pulumi:"disableAadAuth"`
	// DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// A class represent managed identities used for request and response
	Identity *ManagedIdentity `pulumi:"identity"`
	// The kind of the service
	Kind *string `pulumi:"kind"`
	// Live trace configuration of a Microsoft.SignalRService resource.
	LiveTraceConfiguration *LiveTraceConfiguration `pulumi:"liveTraceConfiguration"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Network ACLs for the resource
	NetworkACLs *WebPubSubNetworkACLs `pulumi:"networkACLs"`
	// Enable or disable public network access. Default to "Enabled".
	// When it's Enabled, network ACLs still apply.
	// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Enable or disable the regional endpoint. Default to "Enabled".
	// When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
	// This property is replica specific. Disable the regional endpoint without replica is not allowed.
	RegionEndpointEnabled *string `pulumi:"regionEndpointEnabled"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource log configuration of a Microsoft.SignalRService resource.
	ResourceLogConfiguration *ResourceLogConfiguration `pulumi:"resourceLogConfiguration"`
	// The name of the resource.
	ResourceName *string `pulumi:"resourceName"`
	// Stop or start the resource.  Default to "false".
	// When it's true, the data plane of the resource is shutdown.
	// When it's false, the data plane of the resource is started.
	ResourceStopped *string `pulumi:"resourceStopped"`
	// The billing information of the resource.
	Sku *ResourceSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// TLS settings for the resource
	Tls *WebPubSubTlsSettings `pulumi:"tls"`
}

// The set of arguments for constructing a WebPubSub resource.
type WebPubSubArgs struct {
	// DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth pulumi.BoolPtrInput
	// DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth pulumi.BoolPtrInput
	// A class represent managed identities used for request and response
	Identity ManagedIdentityPtrInput
	// The kind of the service
	Kind pulumi.StringPtrInput
	// Live trace configuration of a Microsoft.SignalRService resource.
	LiveTraceConfiguration LiveTraceConfigurationPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Network ACLs for the resource
	NetworkACLs WebPubSubNetworkACLsPtrInput
	// Enable or disable public network access. Default to "Enabled".
	// When it's Enabled, network ACLs still apply.
	// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
	PublicNetworkAccess pulumi.StringPtrInput
	// Enable or disable the regional endpoint. Default to "Enabled".
	// When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
	// This property is replica specific. Disable the regional endpoint without replica is not allowed.
	RegionEndpointEnabled pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource log configuration of a Microsoft.SignalRService resource.
	ResourceLogConfiguration ResourceLogConfigurationPtrInput
	// The name of the resource.
	ResourceName pulumi.StringPtrInput
	// Stop or start the resource.  Default to "false".
	// When it's true, the data plane of the resource is shutdown.
	// When it's false, the data plane of the resource is started.
	ResourceStopped pulumi.StringPtrInput
	// The billing information of the resource.
	Sku ResourceSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// TLS settings for the resource
	Tls WebPubSubTlsSettingsPtrInput
}

func (WebPubSubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubArgs)(nil)).Elem()
}

type WebPubSubInput interface {
	pulumi.Input

	ToWebPubSubOutput() WebPubSubOutput
	ToWebPubSubOutputWithContext(ctx context.Context) WebPubSubOutput
}

func (*WebPubSub) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSub)(nil)).Elem()
}

func (i *WebPubSub) ToWebPubSubOutput() WebPubSubOutput {
	return i.ToWebPubSubOutputWithContext(context.Background())
}

func (i *WebPubSub) ToWebPubSubOutputWithContext(ctx context.Context) WebPubSubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubOutput)
}

func (i *WebPubSub) ToOutput(ctx context.Context) pulumix.Output[*WebPubSub] {
	return pulumix.Output[*WebPubSub]{
		OutputState: i.ToWebPubSubOutputWithContext(ctx).OutputState,
	}
}

type WebPubSubOutput struct{ *pulumi.OutputState }

func (WebPubSubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSub)(nil)).Elem()
}

func (o WebPubSubOutput) ToWebPubSubOutput() WebPubSubOutput {
	return o
}

func (o WebPubSubOutput) ToWebPubSubOutputWithContext(ctx context.Context) WebPubSubOutput {
	return o
}

func (o WebPubSubOutput) ToOutput(ctx context.Context) pulumix.Output[*WebPubSub] {
	return pulumix.Output[*WebPubSub]{
		OutputState: o.OutputState,
	}
}

// DisableLocalAuth
// Enable or disable aad auth
// When set as true, connection with AuthType=aad won't work.
func (o WebPubSubOutput) DisableAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.BoolPtrOutput { return v.DisableAadAuth }).(pulumi.BoolPtrOutput)
}

// DisableLocalAuth
// Enable or disable local auth with AccessKey
// When set as true, connection with AccessKey=xxx won't work.
func (o WebPubSubOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// The publicly accessible IP of the resource.
func (o WebPubSubOutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringOutput { return v.ExternalIP }).(pulumi.StringOutput)
}

// FQDN of the service instance.
func (o WebPubSubOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Deprecated.
func (o WebPubSubOutput) HostNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringOutput { return v.HostNamePrefix }).(pulumi.StringOutput)
}

// A class represent managed identities used for request and response
func (o WebPubSubOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSub) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// The kind of the service
func (o WebPubSubOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Live trace configuration of a Microsoft.SignalRService resource.
func (o WebPubSubOutput) LiveTraceConfiguration() LiveTraceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSub) LiveTraceConfigurationResponsePtrOutput { return v.LiveTraceConfiguration }).(LiveTraceConfigurationResponsePtrOutput)
}

// The geo-location where the resource lives
func (o WebPubSubOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o WebPubSubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network ACLs for the resource
func (o WebPubSubOutput) NetworkACLs() WebPubSubNetworkACLsResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSub) WebPubSubNetworkACLsResponsePtrOutput { return v.NetworkACLs }).(WebPubSubNetworkACLsResponsePtrOutput)
}

// Private endpoint connections to the resource.
func (o WebPubSubOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *WebPubSub) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the resource.
func (o WebPubSubOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Enable or disable public network access. Default to "Enabled".
// When it's Enabled, network ACLs still apply.
// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
func (o WebPubSubOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The publicly accessible port of the resource which is designed for browser/client side usage.
func (o WebPubSubOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.IntOutput { return v.PublicPort }).(pulumi.IntOutput)
}

// Enable or disable the regional endpoint. Default to "Enabled".
// When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
// This property is replica specific. Disable the regional endpoint without replica is not allowed.
func (o WebPubSubOutput) RegionEndpointEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringPtrOutput { return v.RegionEndpointEnabled }).(pulumi.StringPtrOutput)
}

// Resource log configuration of a Microsoft.SignalRService resource.
func (o WebPubSubOutput) ResourceLogConfiguration() ResourceLogConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSub) ResourceLogConfigurationResponsePtrOutput { return v.ResourceLogConfiguration }).(ResourceLogConfigurationResponsePtrOutput)
}

// Stop or start the resource.  Default to "false".
// When it's true, the data plane of the resource is shutdown.
// When it's false, the data plane of the resource is started.
func (o WebPubSubOutput) ResourceStopped() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringPtrOutput { return v.ResourceStopped }).(pulumi.StringPtrOutput)
}

// The publicly accessible port of the resource which is designed for customer server side usage.
func (o WebPubSubOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.IntOutput { return v.ServerPort }).(pulumi.IntOutput)
}

// The list of shared private link resources.
func (o WebPubSubOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v *WebPubSub) SharedPrivateLinkResourceResponseArrayOutput { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

// The billing information of the resource.
func (o WebPubSubOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSub) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WebPubSubOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSub) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o WebPubSubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// TLS settings for the resource
func (o WebPubSubOutput) Tls() WebPubSubTlsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSub) WebPubSubTlsSettingsResponsePtrOutput { return v.Tls }).(WebPubSubTlsSettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WebPubSubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the resource. Probably you need the same or higher version of client SDKs.
func (o WebPubSubOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSub) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubOutput{})
}

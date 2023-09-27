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
type SignalR struct {
	pulumi.CustomResourceState

	// Cross-Origin Resource Sharing (CORS) settings.
	Cors SignalRCorsSettingsResponsePtrOutput `pulumi:"cors"`
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
	// List of the featureFlags.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, its globally default value will be used
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features SignalRFeatureResponseArrayOutput `pulumi:"features"`
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
	NetworkACLs SignalRNetworkACLsResponsePtrOutput `pulumi:"networkACLs"`
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
	// Stop or start the resource.  Default to "False".
	// When it's true, the data plane of the resource is shutdown.
	// When it's false, the data plane of the resource is started.
	ResourceStopped pulumi.StringPtrOutput `pulumi:"resourceStopped"`
	// The publicly accessible port of the resource which is designed for customer server side usage.
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// Serverless settings.
	Serverless ServerlessSettingsResponsePtrOutput `pulumi:"serverless"`
	// The list of shared private link resources.
	SharedPrivateLinkResources SharedPrivateLinkResourceResponseArrayOutput `pulumi:"sharedPrivateLinkResources"`
	// The billing information of the resource.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// TLS settings for the resource
	Tls SignalRTlsSettingsResponsePtrOutput `pulumi:"tls"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The settings for the Upstream when the service is in server-less mode.
	Upstream ServerlessUpstreamSettingsResponsePtrOutput `pulumi:"upstream"`
	// Version of the resource. Probably you need the same or higher version of client SDKs.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSignalR registers a new resource with the given unique name, arguments, and options.
func NewSignalR(ctx *pulumi.Context,
	name string, args *SignalRArgs, opts ...pulumi.ResourceOption) (*SignalR, error) {
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
	if args.Serverless != nil {
		args.Serverless = args.Serverless.ToServerlessSettingsPtrOutput().ApplyT(func(v *ServerlessSettings) *ServerlessSettings { return v.Defaults() }).(ServerlessSettingsPtrOutput)
	}
	if args.Tls != nil {
		args.Tls = args.Tls.ToSignalRTlsSettingsPtrOutput().ApplyT(func(v *SignalRTlsSettings) *SignalRTlsSettings { return v.Defaults() }).(SignalRTlsSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20180301preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20181001:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20200501:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20200701preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210401preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210601preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210901preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20211001:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230201:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230301preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230601preview:SignalR"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SignalR
	err := ctx.RegisterResource("azure-native:signalrservice/v20230801preview:SignalR", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignalR gets an existing SignalR resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignalR(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRState, opts ...pulumi.ResourceOption) (*SignalR, error) {
	var resource SignalR
	err := ctx.ReadResource("azure-native:signalrservice/v20230801preview:SignalR", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SignalR resources.
type signalRState struct {
}

type SignalRState struct {
}

func (SignalRState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRState)(nil)).Elem()
}

type signalRArgs struct {
	// Cross-Origin Resource Sharing (CORS) settings.
	Cors *SignalRCorsSettings `pulumi:"cors"`
	// DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth *bool `pulumi:"disableAadAuth"`
	// DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// List of the featureFlags.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, its globally default value will be used
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features []SignalRFeature `pulumi:"features"`
	// A class represent managed identities used for request and response
	Identity *ManagedIdentity `pulumi:"identity"`
	// The kind of the service
	Kind *string `pulumi:"kind"`
	// Live trace configuration of a Microsoft.SignalRService resource.
	LiveTraceConfiguration *LiveTraceConfiguration `pulumi:"liveTraceConfiguration"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Network ACLs for the resource
	NetworkACLs *SignalRNetworkACLs `pulumi:"networkACLs"`
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
	// Stop or start the resource.  Default to "False".
	// When it's true, the data plane of the resource is shutdown.
	// When it's false, the data plane of the resource is started.
	ResourceStopped *string `pulumi:"resourceStopped"`
	// Serverless settings.
	Serverless *ServerlessSettings `pulumi:"serverless"`
	// The billing information of the resource.
	Sku *ResourceSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// TLS settings for the resource
	Tls *SignalRTlsSettings `pulumi:"tls"`
	// The settings for the Upstream when the service is in server-less mode.
	Upstream *ServerlessUpstreamSettings `pulumi:"upstream"`
}

// The set of arguments for constructing a SignalR resource.
type SignalRArgs struct {
	// Cross-Origin Resource Sharing (CORS) settings.
	Cors SignalRCorsSettingsPtrInput
	// DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth pulumi.BoolPtrInput
	// DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth pulumi.BoolPtrInput
	// List of the featureFlags.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, its globally default value will be used
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features SignalRFeatureArrayInput
	// A class represent managed identities used for request and response
	Identity ManagedIdentityPtrInput
	// The kind of the service
	Kind pulumi.StringPtrInput
	// Live trace configuration of a Microsoft.SignalRService resource.
	LiveTraceConfiguration LiveTraceConfigurationPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Network ACLs for the resource
	NetworkACLs SignalRNetworkACLsPtrInput
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
	// Stop or start the resource.  Default to "False".
	// When it's true, the data plane of the resource is shutdown.
	// When it's false, the data plane of the resource is started.
	ResourceStopped pulumi.StringPtrInput
	// Serverless settings.
	Serverless ServerlessSettingsPtrInput
	// The billing information of the resource.
	Sku ResourceSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// TLS settings for the resource
	Tls SignalRTlsSettingsPtrInput
	// The settings for the Upstream when the service is in server-less mode.
	Upstream ServerlessUpstreamSettingsPtrInput
}

func (SignalRArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRArgs)(nil)).Elem()
}

type SignalRInput interface {
	pulumi.Input

	ToSignalROutput() SignalROutput
	ToSignalROutputWithContext(ctx context.Context) SignalROutput
}

func (*SignalR) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalR)(nil)).Elem()
}

func (i *SignalR) ToSignalROutput() SignalROutput {
	return i.ToSignalROutputWithContext(context.Background())
}

func (i *SignalR) ToSignalROutputWithContext(ctx context.Context) SignalROutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalROutput)
}

func (i *SignalR) ToOutput(ctx context.Context) pulumix.Output[*SignalR] {
	return pulumix.Output[*SignalR]{
		OutputState: i.ToSignalROutputWithContext(ctx).OutputState,
	}
}

type SignalROutput struct{ *pulumi.OutputState }

func (SignalROutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalR)(nil)).Elem()
}

func (o SignalROutput) ToSignalROutput() SignalROutput {
	return o
}

func (o SignalROutput) ToSignalROutputWithContext(ctx context.Context) SignalROutput {
	return o
}

func (o SignalROutput) ToOutput(ctx context.Context) pulumix.Output[*SignalR] {
	return pulumix.Output[*SignalR]{
		OutputState: o.OutputState,
	}
}

// Cross-Origin Resource Sharing (CORS) settings.
func (o SignalROutput) Cors() SignalRCorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) SignalRCorsSettingsResponsePtrOutput { return v.Cors }).(SignalRCorsSettingsResponsePtrOutput)
}

// DisableLocalAuth
// Enable or disable aad auth
// When set as true, connection with AuthType=aad won't work.
func (o SignalROutput) DisableAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.BoolPtrOutput { return v.DisableAadAuth }).(pulumi.BoolPtrOutput)
}

// DisableLocalAuth
// Enable or disable local auth with AccessKey
// When set as true, connection with AccessKey=xxx won't work.
func (o SignalROutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// The publicly accessible IP of the resource.
func (o SignalROutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.ExternalIP }).(pulumi.StringOutput)
}

// List of the featureFlags.
//
// FeatureFlags that are not included in the parameters for the update operation will not be modified.
// And the response will only include featureFlags that are explicitly set.
// When a featureFlag is not explicitly set, its globally default value will be used
// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
func (o SignalROutput) Features() SignalRFeatureResponseArrayOutput {
	return o.ApplyT(func(v *SignalR) SignalRFeatureResponseArrayOutput { return v.Features }).(SignalRFeatureResponseArrayOutput)
}

// FQDN of the service instance.
func (o SignalROutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Deprecated.
func (o SignalROutput) HostNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.HostNamePrefix }).(pulumi.StringOutput)
}

// A class represent managed identities used for request and response
func (o SignalROutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// The kind of the service
func (o SignalROutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Live trace configuration of a Microsoft.SignalRService resource.
func (o SignalROutput) LiveTraceConfiguration() LiveTraceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) LiveTraceConfigurationResponsePtrOutput { return v.LiveTraceConfiguration }).(LiveTraceConfigurationResponsePtrOutput)
}

// The geo-location where the resource lives
func (o SignalROutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SignalROutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network ACLs for the resource
func (o SignalROutput) NetworkACLs() SignalRNetworkACLsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) SignalRNetworkACLsResponsePtrOutput { return v.NetworkACLs }).(SignalRNetworkACLsResponsePtrOutput)
}

// Private endpoint connections to the resource.
func (o SignalROutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *SignalR) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the resource.
func (o SignalROutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Enable or disable public network access. Default to "Enabled".
// When it's Enabled, network ACLs still apply.
// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
func (o SignalROutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The publicly accessible port of the resource which is designed for browser/client side usage.
func (o SignalROutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SignalR) pulumi.IntOutput { return v.PublicPort }).(pulumi.IntOutput)
}

// Enable or disable the regional endpoint. Default to "Enabled".
// When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
// This property is replica specific. Disable the regional endpoint without replica is not allowed.
func (o SignalROutput) RegionEndpointEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.RegionEndpointEnabled }).(pulumi.StringPtrOutput)
}

// Resource log configuration of a Microsoft.SignalRService resource.
func (o SignalROutput) ResourceLogConfiguration() ResourceLogConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ResourceLogConfigurationResponsePtrOutput { return v.ResourceLogConfiguration }).(ResourceLogConfigurationResponsePtrOutput)
}

// Stop or start the resource.  Default to "False".
// When it's true, the data plane of the resource is shutdown.
// When it's false, the data plane of the resource is started.
func (o SignalROutput) ResourceStopped() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.ResourceStopped }).(pulumi.StringPtrOutput)
}

// The publicly accessible port of the resource which is designed for customer server side usage.
func (o SignalROutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SignalR) pulumi.IntOutput { return v.ServerPort }).(pulumi.IntOutput)
}

// Serverless settings.
func (o SignalROutput) Serverless() ServerlessSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ServerlessSettingsResponsePtrOutput { return v.Serverless }).(ServerlessSettingsResponsePtrOutput)
}

// The list of shared private link resources.
func (o SignalROutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v *SignalR) SharedPrivateLinkResourceResponseArrayOutput { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

// The billing information of the resource.
func (o SignalROutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SignalROutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalR) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SignalROutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// TLS settings for the resource
func (o SignalROutput) Tls() SignalRTlsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) SignalRTlsSettingsResponsePtrOutput { return v.Tls }).(SignalRTlsSettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SignalROutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The settings for the Upstream when the service is in server-less mode.
func (o SignalROutput) Upstream() ServerlessUpstreamSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ServerlessUpstreamSettingsResponsePtrOutput { return v.Upstream }).(ServerlessUpstreamSettingsResponsePtrOutput)
}

// Version of the resource. Probably you need the same or higher version of client SDKs.
func (o SignalROutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalROutput{})
}

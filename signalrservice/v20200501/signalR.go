// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A class represent a SignalR service resource.
type SignalR struct {
	pulumi.CustomResourceState

	// Cross-Origin Resource Sharing (CORS) settings.
	Cors SignalRCorsSettingsResponsePtrOutput `pulumi:"cors"`
	// The publicly accessible IP of the SignalR service.
	ExternalIP pulumi.StringOutput `pulumi:"externalIP"`
	// List of SignalR featureFlags. e.g. ServiceMode.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, SignalR service will use its globally default value.
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features SignalRFeatureResponseArrayOutput `pulumi:"features"`
	// FQDN of the SignalR service instance. Format: xxx.service.signalr.net
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Prefix for the hostName of the SignalR service. Retained for future use.
	// The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
	HostNamePrefix pulumi.StringOutput `pulumi:"hostNamePrefix"`
	// The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network ACLs
	NetworkACLs SignalRNetworkACLsResponsePtrOutput `pulumi:"networkACLs"`
	// Private endpoint connections to the SignalR resource.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The publicly accessible port of the SignalR service which is designed for browser/client side usage.
	PublicPort pulumi.IntOutput `pulumi:"publicPort"`
	// The publicly accessible port of the SignalR service which is designed for customer server side usage.
	ServerPort pulumi.IntOutput `pulumi:"serverPort"`
	// The billing information of the resource.(e.g. Free, Standard)
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type pulumi.StringOutput `pulumi:"type"`
	// Upstream settings when the Azure SignalR is in server-less mode.
	Upstream ServerlessUpstreamSettingsResponsePtrOutput `pulumi:"upstream"`
	// Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
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
	if args.NetworkACLs != nil {
		args.NetworkACLs = args.NetworkACLs.ToSignalRNetworkACLsPtrOutput().ApplyT(func(v *SignalRNetworkACLs) *SignalRNetworkACLs { return v.Defaults() }).(SignalRNetworkACLsPtrOutput)
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
	})
	opts = append(opts, aliases)
	var resource SignalR
	err := ctx.RegisterResource("azure-native:signalrservice/v20200501:SignalR", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:signalrservice/v20200501:SignalR", name, id, state, &resource, opts...)
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
	// List of SignalR featureFlags. e.g. ServiceMode.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, SignalR service will use its globally default value.
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features []SignalRFeature `pulumi:"features"`
	// The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
	Kind *string `pulumi:"kind"`
	// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
	Location *string `pulumi:"location"`
	// Network ACLs
	NetworkACLs *SignalRNetworkACLs `pulumi:"networkACLs"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SignalR resource.
	ResourceName *string `pulumi:"resourceName"`
	// The billing information of the resource.(e.g. Free, Standard)
	Sku *ResourceSku `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
	// Upstream settings when the Azure SignalR is in server-less mode.
	Upstream *ServerlessUpstreamSettings `pulumi:"upstream"`
}

// The set of arguments for constructing a SignalR resource.
type SignalRArgs struct {
	// Cross-Origin Resource Sharing (CORS) settings.
	Cors SignalRCorsSettingsPtrInput
	// List of SignalR featureFlags. e.g. ServiceMode.
	//
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, SignalR service will use its globally default value.
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features SignalRFeatureArrayInput
	// The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
	Kind pulumi.StringPtrInput
	// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
	Location pulumi.StringPtrInput
	// Network ACLs
	NetworkACLs SignalRNetworkACLsPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the SignalR resource.
	ResourceName pulumi.StringPtrInput
	// The billing information of the resource.(e.g. Free, Standard)
	Sku ResourceSkuPtrInput
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags pulumi.StringMapInput
	// Upstream settings when the Azure SignalR is in server-less mode.
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

// Cross-Origin Resource Sharing (CORS) settings.
func (o SignalROutput) Cors() SignalRCorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) SignalRCorsSettingsResponsePtrOutput { return v.Cors }).(SignalRCorsSettingsResponsePtrOutput)
}

// The publicly accessible IP of the SignalR service.
func (o SignalROutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.ExternalIP }).(pulumi.StringOutput)
}

// List of SignalR featureFlags. e.g. ServiceMode.
//
// FeatureFlags that are not included in the parameters for the update operation will not be modified.
// And the response will only include featureFlags that are explicitly set.
// When a featureFlag is not explicitly set, SignalR service will use its globally default value.
// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
func (o SignalROutput) Features() SignalRFeatureResponseArrayOutput {
	return o.ApplyT(func(v *SignalR) SignalRFeatureResponseArrayOutput { return v.Features }).(SignalRFeatureResponseArrayOutput)
}

// FQDN of the SignalR service instance. Format: xxx.service.signalr.net
func (o SignalROutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Prefix for the hostName of the SignalR service. Retained for future use.
// The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
func (o SignalROutput) HostNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.HostNamePrefix }).(pulumi.StringOutput)
}

// The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
func (o SignalROutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
func (o SignalROutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o SignalROutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network ACLs
func (o SignalROutput) NetworkACLs() SignalRNetworkACLsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) SignalRNetworkACLsResponsePtrOutput { return v.NetworkACLs }).(SignalRNetworkACLsResponsePtrOutput)
}

// Private endpoint connections to the SignalR resource.
func (o SignalROutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *SignalR) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the resource.
func (o SignalROutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The publicly accessible port of the SignalR service which is designed for browser/client side usage.
func (o SignalROutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SignalR) pulumi.IntOutput { return v.PublicPort }).(pulumi.IntOutput)
}

// The publicly accessible port of the SignalR service which is designed for customer server side usage.
func (o SignalROutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SignalR) pulumi.IntOutput { return v.ServerPort }).(pulumi.IntOutput)
}

// The billing information of the resource.(e.g. Free, Standard)
func (o SignalROutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Tags of the service which is a list of key value pairs that describe the resource.
func (o SignalROutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
func (o SignalROutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Upstream settings when the Azure SignalR is in server-less mode.
func (o SignalROutput) Upstream() ServerlessUpstreamSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ServerlessUpstreamSettingsResponsePtrOutput { return v.Upstream }).(ServerlessUpstreamSettingsResponsePtrOutput)
}

// Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
func (o SignalROutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalROutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webpubsub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A private endpoint connection to an azure resource
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2021-04-01-preview
type WebPubSubPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Group IDs
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// Connection state of the private endpoint connection
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebPubSubPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewWebPubSubPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *WebPubSubPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*WebPubSubPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:webpubsub/v20210401preview:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210601preview:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20210901preview:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20211001:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20220801preview:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230201:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230301preview:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230601preview:WebPubSubPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230801preview:WebPubSubPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebPubSubPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:webpubsub:WebPubSubPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebPubSubPrivateEndpointConnection gets an existing WebPubSubPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebPubSubPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*WebPubSubPrivateEndpointConnection, error) {
	var resource WebPubSubPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:webpubsub:WebPubSubPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebPubSubPrivateEndpointConnection resources.
type webPubSubPrivateEndpointConnectionState struct {
}

type WebPubSubPrivateEndpointConnectionState struct {
}

func (WebPubSubPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubPrivateEndpointConnectionState)(nil)).Elem()
}

type webPubSubPrivateEndpointConnectionArgs struct {
	// Private endpoint
	PrivateEndpoint *PrivateEndpoint `pulumi:"privateEndpoint"`
	// The name of the private endpoint connection
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// Connection state of the private endpoint connection
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a WebPubSubPrivateEndpointConnection resource.
type WebPubSubPrivateEndpointConnectionArgs struct {
	// Private endpoint
	PrivateEndpoint PrivateEndpointPtrInput
	// The name of the private endpoint connection
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// Connection state of the private endpoint connection
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringInput
}

func (WebPubSubPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubPrivateEndpointConnectionArgs)(nil)).Elem()
}

type WebPubSubPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToWebPubSubPrivateEndpointConnectionOutput() WebPubSubPrivateEndpointConnectionOutput
	ToWebPubSubPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebPubSubPrivateEndpointConnectionOutput
}

func (*WebPubSubPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubPrivateEndpointConnection)(nil)).Elem()
}

func (i *WebPubSubPrivateEndpointConnection) ToWebPubSubPrivateEndpointConnectionOutput() WebPubSubPrivateEndpointConnectionOutput {
	return i.ToWebPubSubPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *WebPubSubPrivateEndpointConnection) ToWebPubSubPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebPubSubPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubPrivateEndpointConnectionOutput)
}

func (i *WebPubSubPrivateEndpointConnection) ToOutput(ctx context.Context) pulumix.Output[*WebPubSubPrivateEndpointConnection] {
	return pulumix.Output[*WebPubSubPrivateEndpointConnection]{
		OutputState: i.ToWebPubSubPrivateEndpointConnectionOutputWithContext(ctx).OutputState,
	}
}

type WebPubSubPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (WebPubSubPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubPrivateEndpointConnection)(nil)).Elem()
}

func (o WebPubSubPrivateEndpointConnectionOutput) ToWebPubSubPrivateEndpointConnectionOutput() WebPubSubPrivateEndpointConnectionOutput {
	return o
}

func (o WebPubSubPrivateEndpointConnectionOutput) ToWebPubSubPrivateEndpointConnectionOutputWithContext(ctx context.Context) WebPubSubPrivateEndpointConnectionOutput {
	return o
}

func (o WebPubSubPrivateEndpointConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*WebPubSubPrivateEndpointConnection] {
	return pulumix.Output[*WebPubSubPrivateEndpointConnection]{
		OutputState: o.OutputState,
	}
}

// Group IDs
func (o WebPubSubPrivateEndpointConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The name of the resource.
func (o WebPubSubPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint
func (o WebPubSubPrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// Connection state of the private endpoint connection
func (o WebPubSubPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) PrivateLinkServiceConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// Provisioning state of the resource.
func (o WebPubSubPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o WebPubSubPrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
func (o WebPubSubPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubPrivateEndpointConnectionOutput{})
}

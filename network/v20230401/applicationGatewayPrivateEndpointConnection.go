// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private Endpoint connection on an application gateway.
type ApplicationGatewayPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The consumer link id.
	LinkIdentifier pulumi.StringOutput `pulumi:"linkIdentifier"`
	// Name of the private endpoint connection on an application gateway.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponseOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the application gateway private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationGatewayPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewApplicationGatewayPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *ApplicationGatewayPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*ApplicationGatewayPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:ApplicationGatewayPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationGatewayPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:network/v20230401:ApplicationGatewayPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGatewayPrivateEndpointConnection gets an existing ApplicationGatewayPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGatewayPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGatewayPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*ApplicationGatewayPrivateEndpointConnection, error) {
	var resource ApplicationGatewayPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:network/v20230401:ApplicationGatewayPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGatewayPrivateEndpointConnection resources.
type applicationGatewayPrivateEndpointConnectionState struct {
}

type ApplicationGatewayPrivateEndpointConnectionState struct {
}

func (ApplicationGatewayPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayPrivateEndpointConnectionState)(nil)).Elem()
}

type applicationGatewayPrivateEndpointConnectionArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	// The name of the application gateway private endpoint connection.
	ConnectionName *string `pulumi:"connectionName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the private endpoint connection on an application gateway.
	Name *string `pulumi:"name"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ApplicationGatewayPrivateEndpointConnection resource.
type ApplicationGatewayPrivateEndpointConnectionArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName pulumi.StringInput
	// The name of the application gateway private endpoint connection.
	ConnectionName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Name of the private endpoint connection on an application gateway.
	Name pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ApplicationGatewayPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayPrivateEndpointConnectionArgs)(nil)).Elem()
}

type ApplicationGatewayPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToApplicationGatewayPrivateEndpointConnectionOutput() ApplicationGatewayPrivateEndpointConnectionOutput
	ToApplicationGatewayPrivateEndpointConnectionOutputWithContext(ctx context.Context) ApplicationGatewayPrivateEndpointConnectionOutput
}

func (*ApplicationGatewayPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayPrivateEndpointConnection)(nil)).Elem()
}

func (i *ApplicationGatewayPrivateEndpointConnection) ToApplicationGatewayPrivateEndpointConnectionOutput() ApplicationGatewayPrivateEndpointConnectionOutput {
	return i.ToApplicationGatewayPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *ApplicationGatewayPrivateEndpointConnection) ToApplicationGatewayPrivateEndpointConnectionOutputWithContext(ctx context.Context) ApplicationGatewayPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPrivateEndpointConnectionOutput)
}

type ApplicationGatewayPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayPrivateEndpointConnection)(nil)).Elem()
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) ToApplicationGatewayPrivateEndpointConnectionOutput() ApplicationGatewayPrivateEndpointConnectionOutput {
	return o
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) ToApplicationGatewayPrivateEndpointConnectionOutputWithContext(ctx context.Context) ApplicationGatewayPrivateEndpointConnectionOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o ApplicationGatewayPrivateEndpointConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The consumer link id.
func (o ApplicationGatewayPrivateEndpointConnectionOutput) LinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringOutput { return v.LinkIdentifier }).(pulumi.StringOutput)
}

// Name of the private endpoint connection on an application gateway.
func (o ApplicationGatewayPrivateEndpointConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The resource of private end point.
func (o ApplicationGatewayPrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) PrivateEndpointResponseOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponseOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o ApplicationGatewayPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) PrivateLinkServiceConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// The provisioning state of the application gateway private endpoint connection resource.
func (o ApplicationGatewayPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o ApplicationGatewayPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGatewayPrivateEndpointConnectionOutput{})
}

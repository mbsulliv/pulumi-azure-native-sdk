// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalrservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A private endpoint connection to an azure resource
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-05-01.
//
// Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview.
type SignalRPrivateEndpointConnection struct {
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

// NewSignalRPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewSignalRPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *SignalRPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*SignalRPrivateEndpointConnection, error) {
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
			Type: pulumi.String("azure-native:signalrservice/v20200501:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20200701preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210401preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210601preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210901preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20211001:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230201:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230301preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230601preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230801preview:SignalRPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SignalRPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:signalrservice:SignalRPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignalRPrivateEndpointConnection gets an existing SignalRPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignalRPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*SignalRPrivateEndpointConnection, error) {
	var resource SignalRPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:signalrservice:SignalRPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SignalRPrivateEndpointConnection resources.
type signalRPrivateEndpointConnectionState struct {
}

type SignalRPrivateEndpointConnectionState struct {
}

func (SignalRPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRPrivateEndpointConnectionState)(nil)).Elem()
}

type signalRPrivateEndpointConnectionArgs struct {
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

// The set of arguments for constructing a SignalRPrivateEndpointConnection resource.
type SignalRPrivateEndpointConnectionArgs struct {
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

func (SignalRPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRPrivateEndpointConnectionArgs)(nil)).Elem()
}

type SignalRPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToSignalRPrivateEndpointConnectionOutput() SignalRPrivateEndpointConnectionOutput
	ToSignalRPrivateEndpointConnectionOutputWithContext(ctx context.Context) SignalRPrivateEndpointConnectionOutput
}

func (*SignalRPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRPrivateEndpointConnection)(nil)).Elem()
}

func (i *SignalRPrivateEndpointConnection) ToSignalRPrivateEndpointConnectionOutput() SignalRPrivateEndpointConnectionOutput {
	return i.ToSignalRPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *SignalRPrivateEndpointConnection) ToSignalRPrivateEndpointConnectionOutputWithContext(ctx context.Context) SignalRPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRPrivateEndpointConnectionOutput)
}

type SignalRPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (SignalRPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRPrivateEndpointConnection)(nil)).Elem()
}

func (o SignalRPrivateEndpointConnectionOutput) ToSignalRPrivateEndpointConnectionOutput() SignalRPrivateEndpointConnectionOutput {
	return o
}

func (o SignalRPrivateEndpointConnectionOutput) ToSignalRPrivateEndpointConnectionOutputWithContext(ctx context.Context) SignalRPrivateEndpointConnectionOutput {
	return o
}

// Group IDs
func (o SignalRPrivateEndpointConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The name of the resource.
func (o SignalRPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint
func (o SignalRPrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// Connection state of the private endpoint connection
func (o SignalRPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) PrivateLinkServiceConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// Provisioning state of the resource.
func (o SignalRPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o SignalRPrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
func (o SignalRPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalRPrivateEndpointConnectionOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A private endpoint connection
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Group IDs.
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint which the connection belongs to.
	PrivateEndpoint PrivateEndpointPropertyResponsePtrOutput `pulumi:"privateEndpoint"`
	// Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// State of the private endpoint connection.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnection gets an existing PrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:sql/v20230201preview:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnection resources.
type privateEndpointConnectionState struct {
}

type PrivateEndpointConnectionState struct {
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	// Private endpoint which the connection belongs to.
	PrivateEndpoint               *PrivateEndpointProperty `pulumi:"privateEndpoint"`
	PrivateEndpointConnectionName *string                  `pulumi:"privateEndpointConnectionName"`
	// Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// Private endpoint which the connection belongs to.
	PrivateEndpoint               PrivateEndpointPropertyPtrInput
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}

type PrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput
}

func (*PrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

// Group IDs.
func (o PrivateEndpointConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint which the connection belongs to.
func (o PrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointPropertyResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

// Connection state of the private endpoint connection.
func (o PrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

// State of the private endpoint connection.
func (o PrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}

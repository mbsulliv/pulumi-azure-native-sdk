// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// GroupIds from the private link service resource.
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Private Endpoint resource for this Connection.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// Details about the state of the connection.
	PrivateLinkServiceConnectionState ConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentName == nil {
		return nil, errors.New("invalid value for required argument 'ParentName'")
	}
	if args.ParentType == nil {
		return nil, errors.New("invalid value for required argument 'ParentType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20230601preview:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:eventgrid/v20220615:PrivateEndpointConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:eventgrid/v20220615:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	// GroupIds from the private link service resource.
	GroupIds []string `pulumi:"groupIds"`
	// The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
	ParentName string `pulumi:"parentName"`
	// The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\'.
	ParentType string `pulumi:"parentType"`
	// The Private Endpoint resource for this Connection.
	PrivateEndpoint *PrivateEndpoint `pulumi:"privateEndpoint"`
	// The name of the private endpoint connection connection.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// Details about the state of the connection.
	PrivateLinkServiceConnectionState *ConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// GroupIds from the private link service resource.
	GroupIds pulumi.StringArrayInput
	// The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
	ParentName pulumi.StringInput
	// The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\'.
	ParentType pulumi.StringInput
	// The Private Endpoint resource for this Connection.
	PrivateEndpoint PrivateEndpointPtrInput
	// The name of the private endpoint connection connection.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// Details about the state of the connection.
	PrivateLinkServiceConnectionState ConnectionStatePtrInput
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
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

func (i *PrivateEndpointConnection) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnection] {
	return pulumix.Output[*PrivateEndpointConnection]{
		OutputState: i.ToPrivateEndpointConnectionOutputWithContext(ctx).OutputState,
	}
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

func (o PrivateEndpointConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnection] {
	return pulumix.Output[*PrivateEndpointConnection]{
		OutputState: o.OutputState,
	}
}

// GroupIds from the private link service resource.
func (o PrivateEndpointConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// Name of the resource.
func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Private Endpoint resource for this Connection.
func (o PrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// Details about the state of the connection.
func (o PrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() ConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) ConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionStateResponsePtrOutput)
}

// Provisioning state of the Private Endpoint Connection.
func (o PrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Type of the resource.
func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}

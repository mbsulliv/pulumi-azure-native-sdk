// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A private endpoint connection
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Group id of the private endpoint.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint which the connection belongs to.
	PrivateEndpoint PrivateEndpointPropertyResponsePtrOutput `pulumi:"privateEndpoint"`
	// Connection State of the Private Endpoint Connection.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the private endpoint.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:documentdb/v20230915preview:PrivateEndpointConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:documentdb/v20230915preview:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Group id of the private endpoint.
	GroupId *string `pulumi:"groupId"`
	// Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty `pulumi:"privateEndpoint"`
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// Connection State of the Private Endpoint Connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the private endpoint.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Group id of the private endpoint.
	GroupId pulumi.StringPtrInput
	// Private endpoint which the connection belongs to.
	PrivateEndpoint PrivateEndpointPropertyPtrInput
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// Connection State of the Private Endpoint Connection.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyPtrInput
	// Provisioning state of the private endpoint.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
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

// Group id of the private endpoint.
func (o PrivateEndpointConnectionOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint which the connection belongs to.
func (o PrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointPropertyResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

// Connection State of the Private Endpoint Connection.
func (o PrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

// Provisioning state of the private endpoint.
func (o PrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}

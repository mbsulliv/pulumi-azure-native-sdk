// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A private endpoint connection
type ManagedInstancePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint which the connection belongs to.
	PrivateEndpoint ManagedInstancePrivateEndpointPropertyResponsePtrOutput `pulumi:"privateEndpoint"`
	// Connection State of the Private Endpoint Connection.
	PrivateLinkServiceConnectionState ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// State of the Private Endpoint Connection.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedInstancePrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewManagedInstancePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *ManagedInstancePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*ManagedInstancePrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:ManagedInstancePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedInstancePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:sql/v20211101:ManagedInstancePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedInstancePrivateEndpointConnection gets an existing ManagedInstancePrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedInstancePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstancePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*ManagedInstancePrivateEndpointConnection, error) {
	var resource ManagedInstancePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:sql/v20211101:ManagedInstancePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedInstancePrivateEndpointConnection resources.
type managedInstancePrivateEndpointConnectionState struct {
}

type ManagedInstancePrivateEndpointConnectionState struct {
}

func (ManagedInstancePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstancePrivateEndpointConnectionState)(nil)).Elem()
}

type managedInstancePrivateEndpointConnectionArgs struct {
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// Private endpoint which the connection belongs to.
	PrivateEndpoint               *ManagedInstancePrivateEndpointProperty `pulumi:"privateEndpoint"`
	PrivateEndpointConnectionName *string                                 `pulumi:"privateEndpointConnectionName"`
	// Connection State of the Private Endpoint Connection.
	PrivateLinkServiceConnectionState *ManagedInstancePrivateLinkServiceConnectionStateProperty `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedInstancePrivateEndpointConnection resource.
type ManagedInstancePrivateEndpointConnectionArgs struct {
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// Private endpoint which the connection belongs to.
	PrivateEndpoint               ManagedInstancePrivateEndpointPropertyPtrInput
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// Connection State of the Private Endpoint Connection.
	PrivateLinkServiceConnectionState ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
}

func (ManagedInstancePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstancePrivateEndpointConnectionArgs)(nil)).Elem()
}

type ManagedInstancePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToManagedInstancePrivateEndpointConnectionOutput() ManagedInstancePrivateEndpointConnectionOutput
	ToManagedInstancePrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointConnectionOutput
}

func (*ManagedInstancePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateEndpointConnection)(nil)).Elem()
}

func (i *ManagedInstancePrivateEndpointConnection) ToManagedInstancePrivateEndpointConnectionOutput() ManagedInstancePrivateEndpointConnectionOutput {
	return i.ToManagedInstancePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *ManagedInstancePrivateEndpointConnection) ToManagedInstancePrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointConnectionOutput)
}

type ManagedInstancePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateEndpointConnection)(nil)).Elem()
}

func (o ManagedInstancePrivateEndpointConnectionOutput) ToManagedInstancePrivateEndpointConnectionOutput() ManagedInstancePrivateEndpointConnectionOutput {
	return o
}

func (o ManagedInstancePrivateEndpointConnectionOutput) ToManagedInstancePrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointConnectionOutput {
	return o
}

// Resource name.
func (o ManagedInstancePrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint which the connection belongs to.
func (o ManagedInstancePrivateEndpointConnectionOutput) PrivateEndpoint() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
		return v.PrivateEndpoint
	}).(ManagedInstancePrivateEndpointPropertyResponsePtrOutput)
}

// Connection State of the Private Endpoint Connection.
func (o ManagedInstancePrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

// State of the Private Endpoint Connection.
func (o ManagedInstancePrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource type.
func (o ManagedInstancePrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstancePrivateEndpointConnectionOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230707preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
type PrivateEndpointConnectionByHostPool struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnectionByHostPool registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionByHostPool(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionByHostPoolArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByHostPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostPoolName == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolName'")
	}
	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:PrivateEndpointConnectionByHostPool"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionByHostPool
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20230707preview:PrivateEndpointConnectionByHostPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionByHostPool gets an existing PrivateEndpointConnectionByHostPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionByHostPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionByHostPoolState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByHostPool, error) {
	var resource PrivateEndpointConnectionByHostPool
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20230707preview:PrivateEndpointConnectionByHostPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionByHostPool resources.
type privateEndpointConnectionByHostPoolState struct {
}

type PrivateEndpointConnectionByHostPoolState struct {
}

func (PrivateEndpointConnectionByHostPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByHostPoolState)(nil)).Elem()
}

type privateEndpointConnectionByHostPoolArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName string `pulumi:"hostPoolName"`
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionByHostPool resource.
type PrivateEndpointConnectionByHostPoolArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName pulumi.StringInput
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (PrivateEndpointConnectionByHostPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByHostPoolArgs)(nil)).Elem()
}

type PrivateEndpointConnectionByHostPoolInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionByHostPoolOutput() PrivateEndpointConnectionByHostPoolOutput
	ToPrivateEndpointConnectionByHostPoolOutputWithContext(ctx context.Context) PrivateEndpointConnectionByHostPoolOutput
}

func (*PrivateEndpointConnectionByHostPool) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByHostPool)(nil)).Elem()
}

func (i *PrivateEndpointConnectionByHostPool) ToPrivateEndpointConnectionByHostPoolOutput() PrivateEndpointConnectionByHostPoolOutput {
	return i.ToPrivateEndpointConnectionByHostPoolOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionByHostPool) ToPrivateEndpointConnectionByHostPoolOutputWithContext(ctx context.Context) PrivateEndpointConnectionByHostPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionByHostPoolOutput)
}

type PrivateEndpointConnectionByHostPoolOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionByHostPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByHostPool)(nil)).Elem()
}

func (o PrivateEndpointConnectionByHostPoolOutput) ToPrivateEndpointConnectionByHostPoolOutput() PrivateEndpointConnectionByHostPoolOutput {
	return o
}

func (o PrivateEndpointConnectionByHostPoolOutput) ToPrivateEndpointConnectionByHostPoolOutputWithContext(ctx context.Context) PrivateEndpointConnectionByHostPoolOutput {
	return o
}

// The name of the resource
func (o PrivateEndpointConnectionByHostPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o PrivateEndpointConnectionByHostPoolOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) PrivateEndpointResponsePtrOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o PrivateEndpointConnectionByHostPoolOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o PrivateEndpointConnectionByHostPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o PrivateEndpointConnectionByHostPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionByHostPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionByHostPoolOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Private Endpoint Connection resource.
type PrivateEndpointConnectionByName struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnectionByName registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionByName(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionByNameArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:PrivateEndpointConnectionByName"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnectionByName
	err := ctx.RegisterResource("azure-native:apimanagement/v20230301preview:PrivateEndpointConnectionByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionByName gets an existing PrivateEndpointConnectionByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionByNameState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByName, error) {
	var resource PrivateEndpointConnectionByName
	err := ctx.ReadResource("azure-native:apimanagement/v20230301preview:PrivateEndpointConnectionByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionByName resources.
type privateEndpointConnectionByNameState struct {
}

type PrivateEndpointConnectionByNameState struct {
}

func (PrivateEndpointConnectionByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByNameState)(nil)).Elem()
}

type privateEndpointConnectionByNameArgs struct {
	// Private Endpoint Connection Resource Id.
	Id *string `pulumi:"id"`
	// Name of the private endpoint connection.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// The connection state of the private endpoint connection.
	Properties *PrivateEndpointConnectionRequestProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionByName resource.
type PrivateEndpointConnectionByNameArgs struct {
	// Private Endpoint Connection Resource Id.
	Id pulumi.StringPtrInput
	// Name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// The connection state of the private endpoint connection.
	Properties PrivateEndpointConnectionRequestPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (PrivateEndpointConnectionByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByNameArgs)(nil)).Elem()
}

type PrivateEndpointConnectionByNameInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionByNameOutput() PrivateEndpointConnectionByNameOutput
	ToPrivateEndpointConnectionByNameOutputWithContext(ctx context.Context) PrivateEndpointConnectionByNameOutput
}

func (*PrivateEndpointConnectionByName) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByName)(nil)).Elem()
}

func (i *PrivateEndpointConnectionByName) ToPrivateEndpointConnectionByNameOutput() PrivateEndpointConnectionByNameOutput {
	return i.ToPrivateEndpointConnectionByNameOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionByName) ToPrivateEndpointConnectionByNameOutputWithContext(ctx context.Context) PrivateEndpointConnectionByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionByNameOutput)
}

func (i *PrivateEndpointConnectionByName) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnectionByName] {
	return pulumix.Output[*PrivateEndpointConnectionByName]{
		OutputState: i.ToPrivateEndpointConnectionByNameOutputWithContext(ctx).OutputState,
	}
}

type PrivateEndpointConnectionByNameOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByName)(nil)).Elem()
}

func (o PrivateEndpointConnectionByNameOutput) ToPrivateEndpointConnectionByNameOutput() PrivateEndpointConnectionByNameOutput {
	return o
}

func (o PrivateEndpointConnectionByNameOutput) ToPrivateEndpointConnectionByNameOutputWithContext(ctx context.Context) PrivateEndpointConnectionByNameOutput {
	return o
}

func (o PrivateEndpointConnectionByNameOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnectionByName] {
	return pulumix.Output[*PrivateEndpointConnectionByName]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o PrivateEndpointConnectionByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o PrivateEndpointConnectionByNameOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o PrivateEndpointConnectionByNameOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o PrivateEndpointConnectionByNameOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionByNameOutput{})
}

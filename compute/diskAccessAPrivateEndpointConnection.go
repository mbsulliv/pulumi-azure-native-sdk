// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Private Endpoint Connection resource.
// Azure REST API version: 2022-07-02. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2023-01-02, 2023-04-02.
type DiskAccessAPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// private endpoint connection name
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponseOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// private endpoint connection type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDiskAccessAPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *DiskAccessAPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*DiskAccessAPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskAccessName == nil {
		return nil, errors.New("invalid value for required argument 'DiskAccessName'")
	}
	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20200930:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211201:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220702:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230102:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230402:DiskAccessAPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DiskAccessAPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:compute:DiskAccessAPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskAccessAPrivateEndpointConnection gets an existing DiskAccessAPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAccessAPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*DiskAccessAPrivateEndpointConnection, error) {
	var resource DiskAccessAPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:compute:DiskAccessAPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskAccessAPrivateEndpointConnection resources.
type diskAccessAPrivateEndpointConnectionState struct {
}

type DiskAccessAPrivateEndpointConnectionState struct {
}

func (DiskAccessAPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessAPrivateEndpointConnectionState)(nil)).Elem()
}

type diskAccessAPrivateEndpointConnectionArgs struct {
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskAccessName string `pulumi:"diskAccessName"`
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DiskAccessAPrivateEndpointConnection resource.
type DiskAccessAPrivateEndpointConnectionArgs struct {
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskAccessName pulumi.StringInput
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (DiskAccessAPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessAPrivateEndpointConnectionArgs)(nil)).Elem()
}

type DiskAccessAPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput
	ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput
}

func (*DiskAccessAPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccessAPrivateEndpointConnection)(nil)).Elem()
}

func (i *DiskAccessAPrivateEndpointConnection) ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput {
	return i.ToDiskAccessAPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *DiskAccessAPrivateEndpointConnection) ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAccessAPrivateEndpointConnectionOutput)
}

func (i *DiskAccessAPrivateEndpointConnection) ToOutput(ctx context.Context) pulumix.Output[*DiskAccessAPrivateEndpointConnection] {
	return pulumix.Output[*DiskAccessAPrivateEndpointConnection]{
		OutputState: i.ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx).OutputState,
	}
}

type DiskAccessAPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (DiskAccessAPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccessAPrivateEndpointConnection)(nil)).Elem()
}

func (o DiskAccessAPrivateEndpointConnectionOutput) ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput {
	return o
}

func (o DiskAccessAPrivateEndpointConnectionOutput) ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput {
	return o
}

func (o DiskAccessAPrivateEndpointConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*DiskAccessAPrivateEndpointConnection] {
	return pulumix.Output[*DiskAccessAPrivateEndpointConnection]{
		OutputState: o.OutputState,
	}
}

// private endpoint connection name
func (o DiskAccessAPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccessAPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o DiskAccessAPrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *DiskAccessAPrivateEndpointConnection) PrivateEndpointResponseOutput { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
}

// A collection of information about the state of the connection between DiskAccess and Virtual Network.
func (o DiskAccessAPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *DiskAccessAPrivateEndpointConnection) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o DiskAccessAPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccessAPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// private endpoint connection type
func (o DiskAccessAPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccessAPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskAccessAPrivateEndpointConnectionOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230630preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The private endpoint connection of an IotHub
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devices:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200615:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200710preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210201preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210303preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210331:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20220430preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20221115preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20230630:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:devices/v20230630preview:PrivateEndpointConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:devices/v20230630preview:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	// The name of the private endpoint connection
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionProperties `pulumi:"properties"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesInput
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput
	// The name of the IoT hub.
	ResourceName pulumi.StringInput
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

// The resource name.
func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of a private endpoint connection
func (o PrivateEndpointConnectionOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointConnectionPropertiesResponseOutput {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

// The resource type.
func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}

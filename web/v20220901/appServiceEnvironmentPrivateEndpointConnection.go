// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Remote Private Endpoint Connection ARM resource.
type AppServiceEnvironmentPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Private IPAddresses mapped to the remote private endpoint
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// PrivateEndpoint of a remote private endpoint connection
	PrivateEndpoint ArmIdWrapperResponsePtrOutput `pulumi:"privateEndpoint"`
	// The state of a private link connection
	PrivateLinkServiceConnectionState PrivateLinkConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                         `pulumi:"provisioningState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppServiceEnvironmentPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewAppServiceEnvironmentPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *AppServiceEnvironmentPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*AppServiceEnvironmentPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:AppServiceEnvironmentPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AppServiceEnvironmentPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:web/v20220901:AppServiceEnvironmentPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServiceEnvironmentPrivateEndpointConnection gets an existing AppServiceEnvironmentPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServiceEnvironmentPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceEnvironmentPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*AppServiceEnvironmentPrivateEndpointConnection, error) {
	var resource AppServiceEnvironmentPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:web/v20220901:AppServiceEnvironmentPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServiceEnvironmentPrivateEndpointConnection resources.
type appServiceEnvironmentPrivateEndpointConnectionState struct {
}

type AppServiceEnvironmentPrivateEndpointConnectionState struct {
}

func (AppServiceEnvironmentPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentPrivateEndpointConnectionState)(nil)).Elem()
}

type appServiceEnvironmentPrivateEndpointConnectionArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the App Service Environment.
	Name                          string  `pulumi:"name"`
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// The state of a private link connection
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AppServiceEnvironmentPrivateEndpointConnection resource.
type AppServiceEnvironmentPrivateEndpointConnectionArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the App Service Environment.
	Name                          pulumi.StringInput
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// The state of a private link connection
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (AppServiceEnvironmentPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentPrivateEndpointConnectionArgs)(nil)).Elem()
}

type AppServiceEnvironmentPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToAppServiceEnvironmentPrivateEndpointConnectionOutput() AppServiceEnvironmentPrivateEndpointConnectionOutput
	ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) AppServiceEnvironmentPrivateEndpointConnectionOutput
}

func (*AppServiceEnvironmentPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironmentPrivateEndpointConnection)(nil)).Elem()
}

func (i *AppServiceEnvironmentPrivateEndpointConnection) ToAppServiceEnvironmentPrivateEndpointConnectionOutput() AppServiceEnvironmentPrivateEndpointConnectionOutput {
	return i.ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *AppServiceEnvironmentPrivateEndpointConnection) ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) AppServiceEnvironmentPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceEnvironmentPrivateEndpointConnectionOutput)
}

func (i *AppServiceEnvironmentPrivateEndpointConnection) ToOutput(ctx context.Context) pulumix.Output[*AppServiceEnvironmentPrivateEndpointConnection] {
	return pulumix.Output[*AppServiceEnvironmentPrivateEndpointConnection]{
		OutputState: i.ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(ctx).OutputState,
	}
}

type AppServiceEnvironmentPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (AppServiceEnvironmentPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceEnvironmentPrivateEndpointConnection)(nil)).Elem()
}

func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) ToAppServiceEnvironmentPrivateEndpointConnectionOutput() AppServiceEnvironmentPrivateEndpointConnectionOutput {
	return o
}

func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) AppServiceEnvironmentPrivateEndpointConnectionOutput {
	return o
}

func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*AppServiceEnvironmentPrivateEndpointConnection] {
	return pulumix.Output[*AppServiceEnvironmentPrivateEndpointConnection]{
		OutputState: o.OutputState,
	}
}

// Private IPAddresses mapped to the remote private endpoint
func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentPrivateEndpointConnection) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Kind of resource.
func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentPrivateEndpointConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// PrivateEndpoint of a remote private endpoint connection
func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentPrivateEndpointConnection) ArmIdWrapperResponsePtrOutput {
		return v.PrivateEndpoint
	}).(ArmIdWrapperResponsePtrOutput)
}

// The state of a private link connection
func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentPrivateEndpointConnection) PrivateLinkConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentPrivateEndpointConnection) pulumi.StringOutput {
		return v.ProvisioningState
	}).(pulumi.StringOutput)
}

// Resource type.
func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceEnvironmentPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServiceEnvironmentPrivateEndpointConnectionOutput{})
}

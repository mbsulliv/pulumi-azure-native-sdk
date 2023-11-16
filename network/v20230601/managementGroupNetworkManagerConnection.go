// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Network Manager Connection resource
type ManagementGroupNetworkManagerConnection struct {
	pulumi.CustomResourceState

	// A description of the network manager connection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network Manager Id.
	NetworkManagerId pulumi.StringPtrOutput `pulumi:"networkManagerId"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementGroupNetworkManagerConnection registers a new resource with the given unique name, arguments, and options.
func NewManagementGroupNetworkManagerConnection(ctx *pulumi.Context,
	name string, args *ManagementGroupNetworkManagerConnectionArgs, opts ...pulumi.ResourceOption) (*ManagementGroupNetworkManagerConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ManagementGroupNetworkManagerConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ManagementGroupNetworkManagerConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagementGroupNetworkManagerConnection
	err := ctx.RegisterResource("azure-native:network/v20230601:ManagementGroupNetworkManagerConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementGroupNetworkManagerConnection gets an existing ManagementGroupNetworkManagerConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementGroupNetworkManagerConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupNetworkManagerConnectionState, opts ...pulumi.ResourceOption) (*ManagementGroupNetworkManagerConnection, error) {
	var resource ManagementGroupNetworkManagerConnection
	err := ctx.ReadResource("azure-native:network/v20230601:ManagementGroupNetworkManagerConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementGroupNetworkManagerConnection resources.
type managementGroupNetworkManagerConnectionState struct {
}

type ManagementGroupNetworkManagerConnectionState struct {
}

func (ManagementGroupNetworkManagerConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupNetworkManagerConnectionState)(nil)).Elem()
}

type managementGroupNetworkManagerConnectionArgs struct {
	// A description of the network manager connection.
	Description *string `pulumi:"description"`
	// The management group Id which uniquely identify the Microsoft Azure management group.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// Name for the network manager connection.
	NetworkManagerConnectionName *string `pulumi:"networkManagerConnectionName"`
	// Network Manager Id.
	NetworkManagerId *string `pulumi:"networkManagerId"`
}

// The set of arguments for constructing a ManagementGroupNetworkManagerConnection resource.
type ManagementGroupNetworkManagerConnectionArgs struct {
	// A description of the network manager connection.
	Description pulumi.StringPtrInput
	// The management group Id which uniquely identify the Microsoft Azure management group.
	ManagementGroupId pulumi.StringInput
	// Name for the network manager connection.
	NetworkManagerConnectionName pulumi.StringPtrInput
	// Network Manager Id.
	NetworkManagerId pulumi.StringPtrInput
}

func (ManagementGroupNetworkManagerConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupNetworkManagerConnectionArgs)(nil)).Elem()
}

type ManagementGroupNetworkManagerConnectionInput interface {
	pulumi.Input

	ToManagementGroupNetworkManagerConnectionOutput() ManagementGroupNetworkManagerConnectionOutput
	ToManagementGroupNetworkManagerConnectionOutputWithContext(ctx context.Context) ManagementGroupNetworkManagerConnectionOutput
}

func (*ManagementGroupNetworkManagerConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupNetworkManagerConnection)(nil)).Elem()
}

func (i *ManagementGroupNetworkManagerConnection) ToManagementGroupNetworkManagerConnectionOutput() ManagementGroupNetworkManagerConnectionOutput {
	return i.ToManagementGroupNetworkManagerConnectionOutputWithContext(context.Background())
}

func (i *ManagementGroupNetworkManagerConnection) ToManagementGroupNetworkManagerConnectionOutputWithContext(ctx context.Context) ManagementGroupNetworkManagerConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupNetworkManagerConnectionOutput)
}

type ManagementGroupNetworkManagerConnectionOutput struct{ *pulumi.OutputState }

func (ManagementGroupNetworkManagerConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupNetworkManagerConnection)(nil)).Elem()
}

func (o ManagementGroupNetworkManagerConnectionOutput) ToManagementGroupNetworkManagerConnectionOutput() ManagementGroupNetworkManagerConnectionOutput {
	return o
}

func (o ManagementGroupNetworkManagerConnectionOutput) ToManagementGroupNetworkManagerConnectionOutputWithContext(ctx context.Context) ManagementGroupNetworkManagerConnectionOutput {
	return o
}

// A description of the network manager connection.
func (o ManagementGroupNetworkManagerConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ManagementGroupNetworkManagerConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource name.
func (o ManagementGroupNetworkManagerConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network Manager Id.
func (o ManagementGroupNetworkManagerConnectionOutput) NetworkManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringPtrOutput { return v.NetworkManagerId }).(pulumi.StringPtrOutput)
}

// The system metadata related to this resource.
func (o ManagementGroupNetworkManagerConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o ManagementGroupNetworkManagerConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupNetworkManagerConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupNetworkManagerConnectionOutput{})
}

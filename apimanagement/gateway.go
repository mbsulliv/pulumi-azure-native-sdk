// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gateway details.
// Azure REST API version: 2022-08-01. Prior API version in Azure Native 1.x: 2020-12-01
type Gateway struct {
	pulumi.CustomResourceState

	// Gateway description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gateway location.
	LocationData ResourceLocationDataContractResponsePtrOutput `pulumi:"locationData"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
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
			Type: pulumi.String("azure-native:apimanagement/v20191201:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:Gateway"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Gateway
	err := ctx.RegisterResource("azure-native:apimanagement:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("azure-native:apimanagement:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
}

type GatewayState struct {
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// Gateway description
	Description *string `pulumi:"description"`
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId *string `pulumi:"gatewayId"`
	// Gateway location.
	LocationData *ResourceLocationDataContract `pulumi:"locationData"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// Gateway description
	Description pulumi.StringPtrInput
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId pulumi.StringPtrInput
	// Gateway location.
	LocationData ResourceLocationDataContractPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

func (i *Gateway) ToOutput(ctx context.Context) pulumix.Output[*Gateway] {
	return pulumix.Output[*Gateway]{
		OutputState: i.ToGatewayOutputWithContext(ctx).OutputState,
	}
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func (o GatewayOutput) ToOutput(ctx context.Context) pulumix.Output[*Gateway] {
	return pulumix.Output[*Gateway]{
		OutputState: o.OutputState,
	}
}

// Gateway description
func (o GatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gateway location.
func (o GatewayOutput) LocationData() ResourceLocationDataContractResponsePtrOutput {
	return o.ApplyT(func(v *Gateway) ResourceLocationDataContractResponsePtrOutput { return v.LocationData }).(ResourceLocationDataContractResponsePtrOutput)
}

// The name of the resource
func (o GatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayOutput{})
}

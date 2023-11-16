// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagemover

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Endpoint resource, which contains information about file sources and targets.
// Azure REST API version: 2023-03-01. Prior API version in Azure Native 1.x: 2022-07-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-10-01.
type Endpoint struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource specific properties for the Storage Mover resource.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource system metadata.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageMoverName == nil {
		return nil, errors.New("invalid value for required argument 'StorageMoverName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagemover/v20220701preview:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagemover/v20230301:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagemover/v20230701preview:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagemover/v20231001:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Endpoint
	err := ctx.RegisterResource("azure-native:storagemover:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure-native:storagemover:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// The name of the Endpoint resource.
	EndpointName *string `pulumi:"endpointName"`
	// The resource specific properties for the Storage Mover resource.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName string `pulumi:"storageMoverName"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// The name of the Endpoint resource.
	EndpointName pulumi.StringPtrInput
	// The resource specific properties for the Storage Mover resource.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Storage Mover resource.
	StorageMoverName pulumi.StringInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

// The name of the resource
func (o EndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource specific properties for the Storage Mover resource.
func (o EndpointOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource system metadata.
func (o EndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Endpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
}

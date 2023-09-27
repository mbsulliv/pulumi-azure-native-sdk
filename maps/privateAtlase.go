// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package maps

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Azure resource which represents which will provision the ability to create private location data.
// Azure REST API version: 2020-02-01-preview. Prior API version in Azure Native 1.x: 2020-02-01-preview
type PrivateAtlase struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Private Atlas resource properties.
	Properties PrivateAtlasPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateAtlase registers a new resource with the given unique name, arguments, and options.
func NewPrivateAtlase(ctx *pulumi.Context,
	name string, args *PrivateAtlaseArgs, opts ...pulumi.ResourceOption) (*PrivateAtlase, error) {
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
			Type: pulumi.String("azure-native:maps/v20200201preview:PrivateAtlase"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateAtlase
	err := ctx.RegisterResource("azure-native:maps:PrivateAtlase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateAtlase gets an existing PrivateAtlase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateAtlase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateAtlaseState, opts ...pulumi.ResourceOption) (*PrivateAtlase, error) {
	var resource PrivateAtlase
	err := ctx.ReadResource("azure-native:maps:PrivateAtlase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateAtlase resources.
type privateAtlaseState struct {
}

type PrivateAtlaseState struct {
}

func (PrivateAtlaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateAtlaseState)(nil)).Elem()
}

type privateAtlaseArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the Private Atlas instance.
	PrivateAtlasName *string `pulumi:"privateAtlasName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateAtlase resource.
type PrivateAtlaseArgs struct {
	// The name of the Maps Account.
	AccountName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the Private Atlas instance.
	PrivateAtlasName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (PrivateAtlaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateAtlaseArgs)(nil)).Elem()
}

type PrivateAtlaseInput interface {
	pulumi.Input

	ToPrivateAtlaseOutput() PrivateAtlaseOutput
	ToPrivateAtlaseOutputWithContext(ctx context.Context) PrivateAtlaseOutput
}

func (*PrivateAtlase) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateAtlase)(nil)).Elem()
}

func (i *PrivateAtlase) ToPrivateAtlaseOutput() PrivateAtlaseOutput {
	return i.ToPrivateAtlaseOutputWithContext(context.Background())
}

func (i *PrivateAtlase) ToPrivateAtlaseOutputWithContext(ctx context.Context) PrivateAtlaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateAtlaseOutput)
}

func (i *PrivateAtlase) ToOutput(ctx context.Context) pulumix.Output[*PrivateAtlase] {
	return pulumix.Output[*PrivateAtlase]{
		OutputState: i.ToPrivateAtlaseOutputWithContext(ctx).OutputState,
	}
}

type PrivateAtlaseOutput struct{ *pulumi.OutputState }

func (PrivateAtlaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateAtlase)(nil)).Elem()
}

func (o PrivateAtlaseOutput) ToPrivateAtlaseOutput() PrivateAtlaseOutput {
	return o
}

func (o PrivateAtlaseOutput) ToPrivateAtlaseOutputWithContext(ctx context.Context) PrivateAtlaseOutput {
	return o
}

func (o PrivateAtlaseOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateAtlase] {
	return pulumix.Output[*PrivateAtlase]{
		OutputState: o.OutputState,
	}
}

// The geo-location where the resource lives
func (o PrivateAtlaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAtlase) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrivateAtlaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAtlase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Private Atlas resource properties.
func (o PrivateAtlaseOutput) Properties() PrivateAtlasPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateAtlase) PrivateAtlasPropertiesResponseOutput { return v.Properties }).(PrivateAtlasPropertiesResponseOutput)
}

// Resource tags.
func (o PrivateAtlaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateAtlase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateAtlaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateAtlase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateAtlaseOutput{})
}

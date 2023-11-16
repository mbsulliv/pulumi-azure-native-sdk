// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230402

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// disk access resource.
type DiskAccess struct {
	pulumi.CustomResourceState

	// The extended location where the disk access will be created. Extended location cannot be changed.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// A readonly collection of private endpoint connections created on the disk. Currently only one endpoint connection is supported.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The disk access resource provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The time when the disk access was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDiskAccess registers a new resource with the given unique name, arguments, and options.
func NewDiskAccess(ctx *pulumi.Context,
	name string, args *DiskAccessArgs, opts ...pulumi.ResourceOption) (*DiskAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211201:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220702:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230102:DiskAccess"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DiskAccess
	err := ctx.RegisterResource("azure-native:compute/v20230402:DiskAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskAccess gets an existing DiskAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAccessState, opts ...pulumi.ResourceOption) (*DiskAccess, error) {
	var resource DiskAccess
	err := ctx.ReadResource("azure-native:compute/v20230402:DiskAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskAccess resources.
type diskAccessState struct {
}

type DiskAccessState struct {
}

func (DiskAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessState)(nil)).Elem()
}

type diskAccessArgs struct {
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskAccessName *string `pulumi:"diskAccessName"`
	// The extended location where the disk access will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DiskAccess resource.
type DiskAccessArgs struct {
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskAccessName pulumi.StringPtrInput
	// The extended location where the disk access will be created. Extended location cannot be changed.
	ExtendedLocation ExtendedLocationPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (DiskAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessArgs)(nil)).Elem()
}

type DiskAccessInput interface {
	pulumi.Input

	ToDiskAccessOutput() DiskAccessOutput
	ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput
}

func (*DiskAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccess)(nil)).Elem()
}

func (i *DiskAccess) ToDiskAccessOutput() DiskAccessOutput {
	return i.ToDiskAccessOutputWithContext(context.Background())
}

func (i *DiskAccess) ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAccessOutput)
}

type DiskAccessOutput struct{ *pulumi.OutputState }

func (DiskAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccess)(nil)).Elem()
}

func (o DiskAccessOutput) ToDiskAccessOutput() DiskAccessOutput {
	return o
}

func (o DiskAccessOutput) ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput {
	return o
}

// The extended location where the disk access will be created. Extended location cannot be changed.
func (o DiskAccessOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *DiskAccess) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Resource location
func (o DiskAccessOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o DiskAccessOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A readonly collection of private endpoint connections created on the disk. Currently only one endpoint connection is supported.
func (o DiskAccessOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *DiskAccess) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The disk access resource provisioning state.
func (o DiskAccessOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags
func (o DiskAccessOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The time when the disk access was created.
func (o DiskAccessOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o DiskAccessOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskAccessOutput{})
}

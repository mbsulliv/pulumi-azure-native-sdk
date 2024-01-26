// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a container on the  Data Box Edge/Gateway device.
// Azure REST API version: 2022-03-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2023-01-01-preview, 2023-07-01, 2023-12-01.
type Container struct {
	pulumi.CustomResourceState

	// Current status of the container.
	ContainerStatus pulumi.StringOutput `pulumi:"containerStatus"`
	// The UTC time when container got created.
	CreatedDateTime pulumi.StringOutput `pulumi:"createdDateTime"`
	// DataFormat for Container
	DataFormat pulumi.StringOutput `pulumi:"dataFormat"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Details of the refresh job on this container.
	RefreshDetails RefreshDetailsResponseOutput `pulumi:"refreshDetails"`
	// Metadata pertaining to creation and last modification of Container
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFormat == nil {
		return nil, errors.New("invalid value for required argument 'DataFormat'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230701:Container"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20231201:Container"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Container
	err := ctx.RegisterResource("azure-native:databoxedge:Container", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerState, opts ...pulumi.ResourceOption) (*Container, error) {
	var resource Container
	err := ctx.ReadResource("azure-native:databoxedge:Container", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Container resources.
type containerState struct {
}

type ContainerState struct {
}

func (ContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerState)(nil)).Elem()
}

type containerArgs struct {
	// The container name.
	ContainerName *string `pulumi:"containerName"`
	// DataFormat for Container
	DataFormat string `pulumi:"dataFormat"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Storage Account Name
	StorageAccountName string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// The container name.
	ContainerName pulumi.StringPtrInput
	// DataFormat for Container
	DataFormat pulumi.StringInput
	// The device name.
	DeviceName pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The Storage Account Name
	StorageAccountName pulumi.StringInput
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerArgs)(nil)).Elem()
}

type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(ctx context.Context) ContainerOutput
}

func (*Container) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (i *Container) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i *Container) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

// Current status of the container.
func (o ContainerOutput) ContainerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.ContainerStatus }).(pulumi.StringOutput)
}

// The UTC time when container got created.
func (o ContainerOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

// DataFormat for Container
func (o ContainerOutput) DataFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.DataFormat }).(pulumi.StringOutput)
}

// The object name.
func (o ContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Details of the refresh job on this container.
func (o ContainerOutput) RefreshDetails() RefreshDetailsResponseOutput {
	return o.ApplyT(func(v *Container) RefreshDetailsResponseOutput { return v.RefreshDetails }).(RefreshDetailsResponseOutput)
}

// Metadata pertaining to creation and last modification of Container
func (o ContainerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Container) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o ContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerOutput{})
}

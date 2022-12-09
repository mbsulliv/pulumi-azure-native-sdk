// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Base class for container with backup items. Containers with specific workloads are derived from this class.
type ProtectionContainer struct {
	pulumi.CustomResourceState

	// Optional ETag.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ProtectionContainerResource properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProtectionContainer registers a new resource with the given unique name, arguments, and options.
func NewProtectionContainer(ctx *pulumi.Context,
	name string, args *ProtectionContainerArgs, opts ...pulumi.ResourceOption) (*ProtectionContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20161201:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201201:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210101:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220901preview:ProtectionContainer"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ProtectionContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource ProtectionContainer
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210601:ProtectionContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectionContainer gets an existing ProtectionContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectionContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionContainerState, opts ...pulumi.ResourceOption) (*ProtectionContainer, error) {
	var resource ProtectionContainer
	err := ctx.ReadResource("azure-native:recoveryservices/v20210601:ProtectionContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectionContainer resources.
type protectionContainerState struct {
}

type ProtectionContainerState struct {
}

func (ProtectionContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionContainerState)(nil)).Elem()
}

type protectionContainerArgs struct {
	// Name of the container to be registered.
	ContainerName *string `pulumi:"containerName"`
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Fabric name associated with the container.
	FabricName string `pulumi:"fabricName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// ProtectionContainerResource properties
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a ProtectionContainer resource.
type ProtectionContainerArgs struct {
	// Name of the container to be registered.
	ContainerName pulumi.StringPtrInput
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Fabric name associated with the container.
	FabricName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// ProtectionContainerResource properties
	Properties pulumi.Input
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the recovery services vault.
	VaultName pulumi.StringInput
}

func (ProtectionContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionContainerArgs)(nil)).Elem()
}

type ProtectionContainerInput interface {
	pulumi.Input

	ToProtectionContainerOutput() ProtectionContainerOutput
	ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput
}

func (*ProtectionContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainer)(nil)).Elem()
}

func (i *ProtectionContainer) ToProtectionContainerOutput() ProtectionContainerOutput {
	return i.ToProtectionContainerOutputWithContext(context.Background())
}

func (i *ProtectionContainer) ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerOutput)
}

type ProtectionContainerOutput struct{ *pulumi.OutputState }

func (ProtectionContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainer)(nil)).Elem()
}

func (o ProtectionContainerOutput) ToProtectionContainerOutput() ProtectionContainerOutput {
	return o
}

func (o ProtectionContainerOutput) ToProtectionContainerOutputWithContext(ctx context.Context) ProtectionContainerOutput {
	return o
}

// Optional ETag.
func (o ProtectionContainerOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainer) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ProtectionContainerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainer) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o ProtectionContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ProtectionContainerResource properties
func (o ProtectionContainerOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ProtectionContainer) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource tags.
func (o ProtectionContainerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProtectionContainer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o ProtectionContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectionContainerOutput{})
}

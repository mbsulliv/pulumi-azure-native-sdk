// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Base class for backup items.
type ProtectedItem struct {
	pulumi.CustomResourceState

	// Optional ETag.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ProtectedItemResource properties
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProtectedItem registers a new resource with the given unique name, arguments, and options.
func NewProtectedItem(ctx *pulumi.Context,
	name string, args *ProtectedItemArgs, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
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
			Type: pulumi.String("azure-native:recoveryservices:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160601:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20190513:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20190615:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201201:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210101:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220901preview:ProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ProtectedItem"),
		},
	})
	opts = append(opts, aliases)
	var resource ProtectedItem
	err := ctx.RegisterResource("azure-native:recoveryservices/v20220401:ProtectedItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectedItem gets an existing ProtectedItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectedItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectedItemState, opts ...pulumi.ResourceOption) (*ProtectedItem, error) {
	var resource ProtectedItem
	err := ctx.ReadResource("azure-native:recoveryservices/v20220401:ProtectedItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectedItem resources.
type protectedItemState struct {
}

type ProtectedItemState struct {
}

func (ProtectedItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedItemState)(nil)).Elem()
}

type protectedItemArgs struct {
	// Container name associated with the backup item.
	ContainerName string `pulumi:"containerName"`
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Fabric name associated with the backup item.
	FabricName string `pulumi:"fabricName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// ProtectedItemResource properties
	Properties interface{} `pulumi:"properties"`
	// Item name to be backed up.
	ProtectedItemName *string `pulumi:"protectedItemName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a ProtectedItem resource.
type ProtectedItemArgs struct {
	// Container name associated with the backup item.
	ContainerName pulumi.StringInput
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Fabric name associated with the backup item.
	FabricName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// ProtectedItemResource properties
	Properties pulumi.Input
	// Item name to be backed up.
	ProtectedItemName pulumi.StringPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the recovery services vault.
	VaultName pulumi.StringInput
}

func (ProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectedItemArgs)(nil)).Elem()
}

type ProtectedItemInput interface {
	pulumi.Input

	ToProtectedItemOutput() ProtectedItemOutput
	ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput
}

func (*ProtectedItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedItem)(nil)).Elem()
}

func (i *ProtectedItem) ToProtectedItemOutput() ProtectedItemOutput {
	return i.ToProtectedItemOutputWithContext(context.Background())
}

func (i *ProtectedItem) ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectedItemOutput)
}

type ProtectedItemOutput struct{ *pulumi.OutputState }

func (ProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedItem)(nil)).Elem()
}

func (o ProtectedItemOutput) ToProtectedItemOutput() ProtectedItemOutput {
	return o
}

func (o ProtectedItemOutput) ToProtectedItemOutputWithContext(ctx context.Context) ProtectedItemOutput {
	return o
}

// Optional ETag.
func (o ProtectedItemOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ProtectedItemOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o ProtectedItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ProtectedItemResource properties
func (o ProtectedItemOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource tags.
func (o ProtectedItemOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o ProtectedItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectedItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectedItemOutput{})
}

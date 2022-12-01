// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Storage mapping object.
type ReplicationStorageClassificationMapping struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the storage mapping object.
	Properties StorageClassificationMappingPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationStorageClassificationMapping registers a new resource with the given unique name, arguments, and options.
func NewReplicationStorageClassificationMapping(ctx *pulumi.Context,
	name string, args *ReplicationStorageClassificationMappingArgs, opts ...pulumi.ResourceOption) (*ReplicationStorageClassificationMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.StorageClassificationName == nil {
		return nil, errors.New("invalid value for required argument 'StorageClassificationName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationStorageClassificationMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationStorageClassificationMapping
	err := ctx.RegisterResource("azure-native:recoveryservices/v20220201:ReplicationStorageClassificationMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationStorageClassificationMapping gets an existing ReplicationStorageClassificationMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationStorageClassificationMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationStorageClassificationMappingState, opts ...pulumi.ResourceOption) (*ReplicationStorageClassificationMapping, error) {
	var resource ReplicationStorageClassificationMapping
	err := ctx.ReadResource("azure-native:recoveryservices/v20220201:ReplicationStorageClassificationMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationStorageClassificationMapping resources.
type replicationStorageClassificationMappingState struct {
}

type ReplicationStorageClassificationMappingState struct {
}

func (ReplicationStorageClassificationMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationStorageClassificationMappingState)(nil)).Elem()
}

type replicationStorageClassificationMappingArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// Storage mapping input properties.
	Properties *StorageMappingInputProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
	// Storage classification mapping name.
	StorageClassificationMappingName *string `pulumi:"storageClassificationMappingName"`
	// Storage classification name.
	StorageClassificationName string `pulumi:"storageClassificationName"`
}

// The set of arguments for constructing a ReplicationStorageClassificationMapping resource.
type ReplicationStorageClassificationMappingArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput
	// Storage mapping input properties.
	Properties StorageMappingInputPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
	// Storage classification mapping name.
	StorageClassificationMappingName pulumi.StringPtrInput
	// Storage classification name.
	StorageClassificationName pulumi.StringInput
}

func (ReplicationStorageClassificationMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationStorageClassificationMappingArgs)(nil)).Elem()
}

type ReplicationStorageClassificationMappingInput interface {
	pulumi.Input

	ToReplicationStorageClassificationMappingOutput() ReplicationStorageClassificationMappingOutput
	ToReplicationStorageClassificationMappingOutputWithContext(ctx context.Context) ReplicationStorageClassificationMappingOutput
}

func (*ReplicationStorageClassificationMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationStorageClassificationMapping)(nil)).Elem()
}

func (i *ReplicationStorageClassificationMapping) ToReplicationStorageClassificationMappingOutput() ReplicationStorageClassificationMappingOutput {
	return i.ToReplicationStorageClassificationMappingOutputWithContext(context.Background())
}

func (i *ReplicationStorageClassificationMapping) ToReplicationStorageClassificationMappingOutputWithContext(ctx context.Context) ReplicationStorageClassificationMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationStorageClassificationMappingOutput)
}

type ReplicationStorageClassificationMappingOutput struct{ *pulumi.OutputState }

func (ReplicationStorageClassificationMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationStorageClassificationMapping)(nil)).Elem()
}

func (o ReplicationStorageClassificationMappingOutput) ToReplicationStorageClassificationMappingOutput() ReplicationStorageClassificationMappingOutput {
	return o
}

func (o ReplicationStorageClassificationMappingOutput) ToReplicationStorageClassificationMappingOutputWithContext(ctx context.Context) ReplicationStorageClassificationMappingOutput {
	return o
}

// Resource Location
func (o ReplicationStorageClassificationMappingOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationStorageClassificationMapping) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationStorageClassificationMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationStorageClassificationMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the storage mapping object.
func (o ReplicationStorageClassificationMappingOutput) Properties() StorageClassificationMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationStorageClassificationMapping) StorageClassificationMappingPropertiesResponseOutput {
		return v.Properties
	}).(StorageClassificationMappingPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationStorageClassificationMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationStorageClassificationMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationStorageClassificationMappingOutput{})
}

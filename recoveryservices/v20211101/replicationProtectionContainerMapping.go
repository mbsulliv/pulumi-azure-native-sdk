// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Protection container mapping object.
type ReplicationProtectionContainerMapping struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The custom data.
	Properties ProtectionContainerMappingPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationProtectionContainerMapping registers a new resource with the given unique name, arguments, and options.
func NewReplicationProtectionContainerMapping(ctx *pulumi.Context,
	name string, args *ReplicationProtectionContainerMappingArgs, opts ...pulumi.ResourceOption) (*ReplicationProtectionContainerMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ProtectionContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ProtectionContainerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationProtectionContainerMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationProtectionContainerMapping
	err := ctx.RegisterResource("azure-native:recoveryservices/v20211101:ReplicationProtectionContainerMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationProtectionContainerMapping gets an existing ReplicationProtectionContainerMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationProtectionContainerMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationProtectionContainerMappingState, opts ...pulumi.ResourceOption) (*ReplicationProtectionContainerMapping, error) {
	var resource ReplicationProtectionContainerMapping
	err := ctx.ReadResource("azure-native:recoveryservices/v20211101:ReplicationProtectionContainerMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationProtectionContainerMapping resources.
type replicationProtectionContainerMappingState struct {
}

type ReplicationProtectionContainerMappingState struct {
}

func (ReplicationProtectionContainerMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectionContainerMappingState)(nil)).Elem()
}

type replicationProtectionContainerMappingArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// Protection container mapping name.
	MappingName *string `pulumi:"mappingName"`
	// Configure protection input properties.
	Properties *CreateProtectionContainerMappingInputProperties `pulumi:"properties"`
	// Protection container name.
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationProtectionContainerMapping resource.
type ReplicationProtectionContainerMappingArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput
	// Protection container mapping name.
	MappingName pulumi.StringPtrInput
	// Configure protection input properties.
	Properties CreateProtectionContainerMappingInputPropertiesPtrInput
	// Protection container name.
	ProtectionContainerName pulumi.StringInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationProtectionContainerMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectionContainerMappingArgs)(nil)).Elem()
}

type ReplicationProtectionContainerMappingInput interface {
	pulumi.Input

	ToReplicationProtectionContainerMappingOutput() ReplicationProtectionContainerMappingOutput
	ToReplicationProtectionContainerMappingOutputWithContext(ctx context.Context) ReplicationProtectionContainerMappingOutput
}

func (*ReplicationProtectionContainerMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectionContainerMapping)(nil)).Elem()
}

func (i *ReplicationProtectionContainerMapping) ToReplicationProtectionContainerMappingOutput() ReplicationProtectionContainerMappingOutput {
	return i.ToReplicationProtectionContainerMappingOutputWithContext(context.Background())
}

func (i *ReplicationProtectionContainerMapping) ToReplicationProtectionContainerMappingOutputWithContext(ctx context.Context) ReplicationProtectionContainerMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProtectionContainerMappingOutput)
}

type ReplicationProtectionContainerMappingOutput struct{ *pulumi.OutputState }

func (ReplicationProtectionContainerMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectionContainerMapping)(nil)).Elem()
}

func (o ReplicationProtectionContainerMappingOutput) ToReplicationProtectionContainerMappingOutput() ReplicationProtectionContainerMappingOutput {
	return o
}

func (o ReplicationProtectionContainerMappingOutput) ToReplicationProtectionContainerMappingOutputWithContext(ctx context.Context) ReplicationProtectionContainerMappingOutput {
	return o
}

// Resource Location
func (o ReplicationProtectionContainerMappingOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectionContainerMapping) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationProtectionContainerMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationProtectionContainerMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The custom data.
func (o ReplicationProtectionContainerMappingOutput) Properties() ProtectionContainerMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationProtectionContainerMapping) ProtectionContainerMappingPropertiesResponseOutput {
		return v.Properties
	}).(ProtectionContainerMappingPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationProtectionContainerMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationProtectionContainerMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationProtectionContainerMappingOutput{})
}

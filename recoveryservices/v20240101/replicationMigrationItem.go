// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Migration item.
type ReplicationMigrationItem struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The migration item properties.
	Properties MigrationItemPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationMigrationItem registers a new resource with the given unique name, arguments, and options.
func NewReplicationMigrationItem(ctx *pulumi.Context,
	name string, args *ReplicationMigrationItemArgs, opts ...pulumi.ResourceOption) (*ReplicationMigrationItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
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
			Type: pulumi.String("azure-native:recoveryservices:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230101:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230201:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230401:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230601:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230801:ReplicationMigrationItem"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReplicationMigrationItem
	err := ctx.RegisterResource("azure-native:recoveryservices/v20240101:ReplicationMigrationItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationMigrationItem gets an existing ReplicationMigrationItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationMigrationItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationMigrationItemState, opts ...pulumi.ResourceOption) (*ReplicationMigrationItem, error) {
	var resource ReplicationMigrationItem
	err := ctx.ReadResource("azure-native:recoveryservices/v20240101:ReplicationMigrationItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationMigrationItem resources.
type replicationMigrationItemState struct {
}

type ReplicationMigrationItemState struct {
}

func (ReplicationMigrationItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationMigrationItemState)(nil)).Elem()
}

type replicationMigrationItemArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// Migration item name.
	MigrationItemName *string `pulumi:"migrationItemName"`
	// Enable migration input properties.
	Properties EnableMigrationInputProperties `pulumi:"properties"`
	// Protection container name.
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationMigrationItem resource.
type ReplicationMigrationItemArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput
	// Migration item name.
	MigrationItemName pulumi.StringPtrInput
	// Enable migration input properties.
	Properties EnableMigrationInputPropertiesInput
	// Protection container name.
	ProtectionContainerName pulumi.StringInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationMigrationItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationMigrationItemArgs)(nil)).Elem()
}

type ReplicationMigrationItemInput interface {
	pulumi.Input

	ToReplicationMigrationItemOutput() ReplicationMigrationItemOutput
	ToReplicationMigrationItemOutputWithContext(ctx context.Context) ReplicationMigrationItemOutput
}

func (*ReplicationMigrationItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationMigrationItem)(nil)).Elem()
}

func (i *ReplicationMigrationItem) ToReplicationMigrationItemOutput() ReplicationMigrationItemOutput {
	return i.ToReplicationMigrationItemOutputWithContext(context.Background())
}

func (i *ReplicationMigrationItem) ToReplicationMigrationItemOutputWithContext(ctx context.Context) ReplicationMigrationItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationMigrationItemOutput)
}

type ReplicationMigrationItemOutput struct{ *pulumi.OutputState }

func (ReplicationMigrationItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationMigrationItem)(nil)).Elem()
}

func (o ReplicationMigrationItemOutput) ToReplicationMigrationItemOutput() ReplicationMigrationItemOutput {
	return o
}

func (o ReplicationMigrationItemOutput) ToReplicationMigrationItemOutputWithContext(ctx context.Context) ReplicationMigrationItemOutput {
	return o
}

// Resource Location
func (o ReplicationMigrationItemOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationMigrationItem) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationMigrationItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationMigrationItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The migration item properties.
func (o ReplicationMigrationItemOutput) Properties() MigrationItemPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationMigrationItem) MigrationItemPropertiesResponseOutput { return v.Properties }).(MigrationItemPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationMigrationItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationMigrationItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationMigrationItemOutput{})
}

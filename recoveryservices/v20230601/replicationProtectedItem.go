// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replication protected item.
type ReplicationProtectedItem struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The custom data.
	Properties ReplicationProtectedItemPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationProtectedItem registers a new resource with the given unique name, arguments, and options.
func NewReplicationProtectedItem(ctx *pulumi.Context,
	name string, args *ReplicationProtectedItemArgs, opts ...pulumi.ResourceOption) (*ReplicationProtectedItem, error) {
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
			Type: pulumi.String("azure-native:recoveryservices:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230101:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230201:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230401:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20230801:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20240101:ReplicationProtectedItem"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReplicationProtectedItem
	err := ctx.RegisterResource("azure-native:recoveryservices/v20230601:ReplicationProtectedItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationProtectedItem gets an existing ReplicationProtectedItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationProtectedItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationProtectedItemState, opts ...pulumi.ResourceOption) (*ReplicationProtectedItem, error) {
	var resource ReplicationProtectedItem
	err := ctx.ReadResource("azure-native:recoveryservices/v20230601:ReplicationProtectedItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationProtectedItem resources.
type replicationProtectedItemState struct {
}

type ReplicationProtectedItemState struct {
}

func (ReplicationProtectedItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectedItemState)(nil)).Elem()
}

type replicationProtectedItemArgs struct {
	// Name of the fabric.
	FabricName string `pulumi:"fabricName"`
	// Enable protection input properties.
	Properties *EnableProtectionInputProperties `pulumi:"properties"`
	// Protection container name.
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	// A name for the replication protected item.
	ReplicatedProtectedItemName *string `pulumi:"replicatedProtectedItemName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationProtectedItem resource.
type ReplicationProtectedItemArgs struct {
	// Name of the fabric.
	FabricName pulumi.StringInput
	// Enable protection input properties.
	Properties EnableProtectionInputPropertiesPtrInput
	// Protection container name.
	ProtectionContainerName pulumi.StringInput
	// A name for the replication protected item.
	ReplicatedProtectedItemName pulumi.StringPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectedItemArgs)(nil)).Elem()
}

type ReplicationProtectedItemInput interface {
	pulumi.Input

	ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput
	ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput
}

func (*ReplicationProtectedItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectedItem)(nil)).Elem()
}

func (i *ReplicationProtectedItem) ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput {
	return i.ToReplicationProtectedItemOutputWithContext(context.Background())
}

func (i *ReplicationProtectedItem) ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProtectedItemOutput)
}

type ReplicationProtectedItemOutput struct{ *pulumi.OutputState }

func (ReplicationProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectedItem)(nil)).Elem()
}

func (o ReplicationProtectedItemOutput) ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput {
	return o
}

func (o ReplicationProtectedItemOutput) ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput {
	return o
}

// Resource Location
func (o ReplicationProtectedItemOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItem) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationProtectedItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationProtectedItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The custom data.
func (o ReplicationProtectedItemOutput) Properties() ReplicationProtectedItemPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationProtectedItem) ReplicationProtectedItemPropertiesResponseOutput {
		return v.Properties
	}).(ReplicationProtectedItemPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationProtectedItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationProtectedItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationProtectedItemOutput{})
}

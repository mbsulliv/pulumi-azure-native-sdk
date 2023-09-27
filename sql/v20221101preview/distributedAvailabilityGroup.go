// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Distributed availability group between box and Sql Managed Instance.
type DistributedAvailabilityGroup struct {
	pulumi.CustomResourceState

	// The distributed availability group id
	DistributedAvailabilityGroupId pulumi.StringOutput `pulumi:"distributedAvailabilityGroupId"`
	// Role of managed instance
	InstanceRole pulumi.StringOutput `pulumi:"instanceRole"`
	// The last hardened lsn
	LastHardenedLsn pulumi.StringOutput `pulumi:"lastHardenedLsn"`
	// The link state
	LinkState pulumi.StringOutput `pulumi:"linkState"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary availability group name
	PrimaryAvailabilityGroupName pulumi.StringPtrOutput `pulumi:"primaryAvailabilityGroupName"`
	// The replication mode of a distributed availability group. Parameter will be ignored during link creation.
	ReplicationMode pulumi.StringPtrOutput `pulumi:"replicationMode"`
	// The secondary availability group name
	SecondaryAvailabilityGroupName pulumi.StringPtrOutput `pulumi:"secondaryAvailabilityGroupName"`
	// The source endpoint
	SourceEndpoint pulumi.StringPtrOutput `pulumi:"sourceEndpoint"`
	// The source replica id
	SourceReplicaId pulumi.StringOutput `pulumi:"sourceReplicaId"`
	// The name of the target database
	TargetDatabase pulumi.StringPtrOutput `pulumi:"targetDatabase"`
	// The target replica id
	TargetReplicaId pulumi.StringOutput `pulumi:"targetReplicaId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDistributedAvailabilityGroup registers a new resource with the given unique name, arguments, and options.
func NewDistributedAvailabilityGroup(ctx *pulumi.Context,
	name string, args *DistributedAvailabilityGroupArgs, opts ...pulumi.ResourceOption) (*DistributedAvailabilityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:DistributedAvailabilityGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DistributedAvailabilityGroup
	err := ctx.RegisterResource("azure-native:sql/v20221101preview:DistributedAvailabilityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDistributedAvailabilityGroup gets an existing DistributedAvailabilityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDistributedAvailabilityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DistributedAvailabilityGroupState, opts ...pulumi.ResourceOption) (*DistributedAvailabilityGroup, error) {
	var resource DistributedAvailabilityGroup
	err := ctx.ReadResource("azure-native:sql/v20221101preview:DistributedAvailabilityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DistributedAvailabilityGroup resources.
type distributedAvailabilityGroupState struct {
}

type DistributedAvailabilityGroupState struct {
}

func (DistributedAvailabilityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*distributedAvailabilityGroupState)(nil)).Elem()
}

type distributedAvailabilityGroupArgs struct {
	// The distributed availability group name.
	DistributedAvailabilityGroupName *string `pulumi:"distributedAvailabilityGroupName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The primary availability group name
	PrimaryAvailabilityGroupName *string `pulumi:"primaryAvailabilityGroupName"`
	// The replication mode of a distributed availability group. Parameter will be ignored during link creation.
	ReplicationMode *string `pulumi:"replicationMode"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The secondary availability group name
	SecondaryAvailabilityGroupName *string `pulumi:"secondaryAvailabilityGroupName"`
	// The source endpoint
	SourceEndpoint *string `pulumi:"sourceEndpoint"`
	// The name of the target database
	TargetDatabase *string `pulumi:"targetDatabase"`
}

// The set of arguments for constructing a DistributedAvailabilityGroup resource.
type DistributedAvailabilityGroupArgs struct {
	// The distributed availability group name.
	DistributedAvailabilityGroupName pulumi.StringPtrInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The primary availability group name
	PrimaryAvailabilityGroupName pulumi.StringPtrInput
	// The replication mode of a distributed availability group. Parameter will be ignored during link creation.
	ReplicationMode pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The secondary availability group name
	SecondaryAvailabilityGroupName pulumi.StringPtrInput
	// The source endpoint
	SourceEndpoint pulumi.StringPtrInput
	// The name of the target database
	TargetDatabase pulumi.StringPtrInput
}

func (DistributedAvailabilityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*distributedAvailabilityGroupArgs)(nil)).Elem()
}

type DistributedAvailabilityGroupInput interface {
	pulumi.Input

	ToDistributedAvailabilityGroupOutput() DistributedAvailabilityGroupOutput
	ToDistributedAvailabilityGroupOutputWithContext(ctx context.Context) DistributedAvailabilityGroupOutput
}

func (*DistributedAvailabilityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributedAvailabilityGroup)(nil)).Elem()
}

func (i *DistributedAvailabilityGroup) ToDistributedAvailabilityGroupOutput() DistributedAvailabilityGroupOutput {
	return i.ToDistributedAvailabilityGroupOutputWithContext(context.Background())
}

func (i *DistributedAvailabilityGroup) ToDistributedAvailabilityGroupOutputWithContext(ctx context.Context) DistributedAvailabilityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributedAvailabilityGroupOutput)
}

func (i *DistributedAvailabilityGroup) ToOutput(ctx context.Context) pulumix.Output[*DistributedAvailabilityGroup] {
	return pulumix.Output[*DistributedAvailabilityGroup]{
		OutputState: i.ToDistributedAvailabilityGroupOutputWithContext(ctx).OutputState,
	}
}

type DistributedAvailabilityGroupOutput struct{ *pulumi.OutputState }

func (DistributedAvailabilityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributedAvailabilityGroup)(nil)).Elem()
}

func (o DistributedAvailabilityGroupOutput) ToDistributedAvailabilityGroupOutput() DistributedAvailabilityGroupOutput {
	return o
}

func (o DistributedAvailabilityGroupOutput) ToDistributedAvailabilityGroupOutputWithContext(ctx context.Context) DistributedAvailabilityGroupOutput {
	return o
}

func (o DistributedAvailabilityGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*DistributedAvailabilityGroup] {
	return pulumix.Output[*DistributedAvailabilityGroup]{
		OutputState: o.OutputState,
	}
}

// The distributed availability group id
func (o DistributedAvailabilityGroupOutput) DistributedAvailabilityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.DistributedAvailabilityGroupId }).(pulumi.StringOutput)
}

// Role of managed instance
func (o DistributedAvailabilityGroupOutput) InstanceRole() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.InstanceRole }).(pulumi.StringOutput)
}

// The last hardened lsn
func (o DistributedAvailabilityGroupOutput) LastHardenedLsn() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.LastHardenedLsn }).(pulumi.StringOutput)
}

// The link state
func (o DistributedAvailabilityGroupOutput) LinkState() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.LinkState }).(pulumi.StringOutput)
}

// Resource name.
func (o DistributedAvailabilityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The primary availability group name
func (o DistributedAvailabilityGroupOutput) PrimaryAvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.PrimaryAvailabilityGroupName }).(pulumi.StringPtrOutput)
}

// The replication mode of a distributed availability group. Parameter will be ignored during link creation.
func (o DistributedAvailabilityGroupOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.ReplicationMode }).(pulumi.StringPtrOutput)
}

// The secondary availability group name
func (o DistributedAvailabilityGroupOutput) SecondaryAvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.SecondaryAvailabilityGroupName }).(pulumi.StringPtrOutput)
}

// The source endpoint
func (o DistributedAvailabilityGroupOutput) SourceEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.SourceEndpoint }).(pulumi.StringPtrOutput)
}

// The source replica id
func (o DistributedAvailabilityGroupOutput) SourceReplicaId() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.SourceReplicaId }).(pulumi.StringOutput)
}

// The name of the target database
func (o DistributedAvailabilityGroupOutput) TargetDatabase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.TargetDatabase }).(pulumi.StringPtrOutput)
}

// The target replica id
func (o DistributedAvailabilityGroupOutput) TargetReplicaId() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.TargetReplicaId }).(pulumi.StringOutput)
}

// Resource type.
func (o DistributedAvailabilityGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DistributedAvailabilityGroupOutput{})
}

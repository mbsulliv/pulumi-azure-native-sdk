// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Distributed availability group between box and Sql Managed Instance.
type DistributedAvailabilityGroup struct {
	pulumi.CustomResourceState

	// Databases in the distributed availability group
	Databases DistributedAvailabilityGroupDatabaseResponseArrayOutput `pulumi:"databases"`
	// ID of the distributed availability group
	DistributedAvailabilityGroupId pulumi.StringOutput `pulumi:"distributedAvailabilityGroupId"`
	// Name of the distributed availability group
	DistributedAvailabilityGroupName pulumi.StringOutput `pulumi:"distributedAvailabilityGroupName"`
	// The link failover mode - can be Manual if intended to be used for two-way failover with a supported SQL Server, or None for one-way failover to Azure.
	FailoverMode pulumi.StringPtrOutput `pulumi:"failoverMode"`
	// Managed instance side availability group name
	InstanceAvailabilityGroupName pulumi.StringPtrOutput `pulumi:"instanceAvailabilityGroupName"`
	// Managed instance side link role
	InstanceLinkRole pulumi.StringPtrOutput `pulumi:"instanceLinkRole"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SQL server side availability group name
	PartnerAvailabilityGroupName pulumi.StringPtrOutput `pulumi:"partnerAvailabilityGroupName"`
	// SQL server side endpoint - IP or DNS resolvable name
	PartnerEndpoint pulumi.StringPtrOutput `pulumi:"partnerEndpoint"`
	// SQL server side link role
	PartnerLinkRole pulumi.StringOutput `pulumi:"partnerLinkRole"`
	// Replication mode of the link
	ReplicationMode pulumi.StringPtrOutput `pulumi:"replicationMode"`
	// Database seeding mode – can be Automatic (default), or Manual for supported scenarios.
	SeedingMode pulumi.StringPtrOutput `pulumi:"seedingMode"`
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
			Type: pulumi.String("azure-native:sql/v20221101preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:DistributedAvailabilityGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DistributedAvailabilityGroup
	err := ctx.RegisterResource("azure-native:sql/v20230801preview:DistributedAvailabilityGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:sql/v20230801preview:DistributedAvailabilityGroup", name, id, state, &resource, opts...)
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
	// Databases in the distributed availability group
	Databases []DistributedAvailabilityGroupDatabase `pulumi:"databases"`
	// The distributed availability group name.
	DistributedAvailabilityGroupName *string `pulumi:"distributedAvailabilityGroupName"`
	// The link failover mode - can be Manual if intended to be used for two-way failover with a supported SQL Server, or None for one-way failover to Azure.
	FailoverMode *string `pulumi:"failoverMode"`
	// Managed instance side availability group name
	InstanceAvailabilityGroupName *string `pulumi:"instanceAvailabilityGroupName"`
	// Managed instance side link role
	InstanceLinkRole *string `pulumi:"instanceLinkRole"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// SQL server side availability group name
	PartnerAvailabilityGroupName *string `pulumi:"partnerAvailabilityGroupName"`
	// SQL server side endpoint - IP or DNS resolvable name
	PartnerEndpoint *string `pulumi:"partnerEndpoint"`
	// Replication mode of the link
	ReplicationMode *string `pulumi:"replicationMode"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Database seeding mode – can be Automatic (default), or Manual for supported scenarios.
	SeedingMode *string `pulumi:"seedingMode"`
}

// The set of arguments for constructing a DistributedAvailabilityGroup resource.
type DistributedAvailabilityGroupArgs struct {
	// Databases in the distributed availability group
	Databases DistributedAvailabilityGroupDatabaseArrayInput
	// The distributed availability group name.
	DistributedAvailabilityGroupName pulumi.StringPtrInput
	// The link failover mode - can be Manual if intended to be used for two-way failover with a supported SQL Server, or None for one-way failover to Azure.
	FailoverMode pulumi.StringPtrInput
	// Managed instance side availability group name
	InstanceAvailabilityGroupName pulumi.StringPtrInput
	// Managed instance side link role
	InstanceLinkRole pulumi.StringPtrInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// SQL server side availability group name
	PartnerAvailabilityGroupName pulumi.StringPtrInput
	// SQL server side endpoint - IP or DNS resolvable name
	PartnerEndpoint pulumi.StringPtrInput
	// Replication mode of the link
	ReplicationMode pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Database seeding mode – can be Automatic (default), or Manual for supported scenarios.
	SeedingMode pulumi.StringPtrInput
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

// Databases in the distributed availability group
func (o DistributedAvailabilityGroupOutput) Databases() DistributedAvailabilityGroupDatabaseResponseArrayOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) DistributedAvailabilityGroupDatabaseResponseArrayOutput {
		return v.Databases
	}).(DistributedAvailabilityGroupDatabaseResponseArrayOutput)
}

// ID of the distributed availability group
func (o DistributedAvailabilityGroupOutput) DistributedAvailabilityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.DistributedAvailabilityGroupId }).(pulumi.StringOutput)
}

// Name of the distributed availability group
func (o DistributedAvailabilityGroupOutput) DistributedAvailabilityGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.DistributedAvailabilityGroupName }).(pulumi.StringOutput)
}

// The link failover mode - can be Manual if intended to be used for two-way failover with a supported SQL Server, or None for one-way failover to Azure.
func (o DistributedAvailabilityGroupOutput) FailoverMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.FailoverMode }).(pulumi.StringPtrOutput)
}

// Managed instance side availability group name
func (o DistributedAvailabilityGroupOutput) InstanceAvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.InstanceAvailabilityGroupName }).(pulumi.StringPtrOutput)
}

// Managed instance side link role
func (o DistributedAvailabilityGroupOutput) InstanceLinkRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.InstanceLinkRole }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o DistributedAvailabilityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SQL server side availability group name
func (o DistributedAvailabilityGroupOutput) PartnerAvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.PartnerAvailabilityGroupName }).(pulumi.StringPtrOutput)
}

// SQL server side endpoint - IP or DNS resolvable name
func (o DistributedAvailabilityGroupOutput) PartnerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.PartnerEndpoint }).(pulumi.StringPtrOutput)
}

// SQL server side link role
func (o DistributedAvailabilityGroupOutput) PartnerLinkRole() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.PartnerLinkRole }).(pulumi.StringOutput)
}

// Replication mode of the link
func (o DistributedAvailabilityGroupOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.ReplicationMode }).(pulumi.StringPtrOutput)
}

// Database seeding mode – can be Automatic (default), or Manual for supported scenarios.
func (o DistributedAvailabilityGroupOutput) SeedingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.SeedingMode }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o DistributedAvailabilityGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DistributedAvailabilityGroupOutput{})
}

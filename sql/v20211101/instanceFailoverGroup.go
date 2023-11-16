// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An instance failover group.
type InstanceFailoverGroup struct {
	pulumi.CustomResourceState

	// List of managed instance pairs in the failover group.
	ManagedInstancePairs ManagedInstancePairInfoResponseArrayOutput `pulumi:"managedInstancePairs"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Partner region information for the failover group.
	PartnerRegions PartnerRegionInfoResponseArrayOutput `pulumi:"partnerRegions"`
	// Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput `pulumi:"readOnlyEndpoint"`
	// Read-write endpoint of the failover group instance.
	ReadWriteEndpoint InstanceFailoverGroupReadWriteEndpointResponseOutput `pulumi:"readWriteEndpoint"`
	// Local replication role of the failover group instance.
	ReplicationRole pulumi.StringOutput `pulumi:"replicationRole"`
	// Replication state of the failover group instance.
	ReplicationState pulumi.StringOutput `pulumi:"replicationState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInstanceFailoverGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceFailoverGroup(ctx *pulumi.Context,
	name string, args *InstanceFailoverGroupArgs, opts ...pulumi.ResourceOption) (*InstanceFailoverGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocationName == nil {
		return nil, errors.New("invalid value for required argument 'LocationName'")
	}
	if args.ManagedInstancePairs == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstancePairs'")
	}
	if args.PartnerRegions == nil {
		return nil, errors.New("invalid value for required argument 'PartnerRegions'")
	}
	if args.ReadWriteEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ReadWriteEndpoint'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:InstanceFailoverGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:InstanceFailoverGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InstanceFailoverGroup
	err := ctx.RegisterResource("azure-native:sql/v20211101:InstanceFailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceFailoverGroup gets an existing InstanceFailoverGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceFailoverGroupState, opts ...pulumi.ResourceOption) (*InstanceFailoverGroup, error) {
	var resource InstanceFailoverGroup
	err := ctx.ReadResource("azure-native:sql/v20211101:InstanceFailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceFailoverGroup resources.
type instanceFailoverGroupState struct {
}

type InstanceFailoverGroupState struct {
}

func (InstanceFailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFailoverGroupState)(nil)).Elem()
}

type instanceFailoverGroupArgs struct {
	// The name of the failover group.
	FailoverGroupName *string `pulumi:"failoverGroupName"`
	// The name of the region where the resource is located.
	LocationName string `pulumi:"locationName"`
	// List of managed instance pairs in the failover group.
	ManagedInstancePairs []ManagedInstancePairInfo `pulumi:"managedInstancePairs"`
	// Partner region information for the failover group.
	PartnerRegions []PartnerRegionInfo `pulumi:"partnerRegions"`
	// Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint *InstanceFailoverGroupReadOnlyEndpoint `pulumi:"readOnlyEndpoint"`
	// Read-write endpoint of the failover group instance.
	ReadWriteEndpoint InstanceFailoverGroupReadWriteEndpoint `pulumi:"readWriteEndpoint"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a InstanceFailoverGroup resource.
type InstanceFailoverGroupArgs struct {
	// The name of the failover group.
	FailoverGroupName pulumi.StringPtrInput
	// The name of the region where the resource is located.
	LocationName pulumi.StringInput
	// List of managed instance pairs in the failover group.
	ManagedInstancePairs ManagedInstancePairInfoArrayInput
	// Partner region information for the failover group.
	PartnerRegions PartnerRegionInfoArrayInput
	// Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint InstanceFailoverGroupReadOnlyEndpointPtrInput
	// Read-write endpoint of the failover group instance.
	ReadWriteEndpoint InstanceFailoverGroupReadWriteEndpointInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
}

func (InstanceFailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFailoverGroupArgs)(nil)).Elem()
}

type InstanceFailoverGroupInput interface {
	pulumi.Input

	ToInstanceFailoverGroupOutput() InstanceFailoverGroupOutput
	ToInstanceFailoverGroupOutputWithContext(ctx context.Context) InstanceFailoverGroupOutput
}

func (*InstanceFailoverGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroup)(nil)).Elem()
}

func (i *InstanceFailoverGroup) ToInstanceFailoverGroupOutput() InstanceFailoverGroupOutput {
	return i.ToInstanceFailoverGroupOutputWithContext(context.Background())
}

func (i *InstanceFailoverGroup) ToInstanceFailoverGroupOutputWithContext(ctx context.Context) InstanceFailoverGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceFailoverGroupOutput)
}

type InstanceFailoverGroupOutput struct{ *pulumi.OutputState }

func (InstanceFailoverGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceFailoverGroup)(nil)).Elem()
}

func (o InstanceFailoverGroupOutput) ToInstanceFailoverGroupOutput() InstanceFailoverGroupOutput {
	return o
}

func (o InstanceFailoverGroupOutput) ToInstanceFailoverGroupOutputWithContext(ctx context.Context) InstanceFailoverGroupOutput {
	return o
}

// List of managed instance pairs in the failover group.
func (o InstanceFailoverGroupOutput) ManagedInstancePairs() ManagedInstancePairInfoResponseArrayOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) ManagedInstancePairInfoResponseArrayOutput {
		return v.ManagedInstancePairs
	}).(ManagedInstancePairInfoResponseArrayOutput)
}

// Resource name.
func (o InstanceFailoverGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Partner region information for the failover group.
func (o InstanceFailoverGroupOutput) PartnerRegions() PartnerRegionInfoResponseArrayOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) PartnerRegionInfoResponseArrayOutput { return v.PartnerRegions }).(PartnerRegionInfoResponseArrayOutput)
}

// Read-only endpoint of the failover group instance.
func (o InstanceFailoverGroupOutput) ReadOnlyEndpoint() InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput {
		return v.ReadOnlyEndpoint
	}).(InstanceFailoverGroupReadOnlyEndpointResponsePtrOutput)
}

// Read-write endpoint of the failover group instance.
func (o InstanceFailoverGroupOutput) ReadWriteEndpoint() InstanceFailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) InstanceFailoverGroupReadWriteEndpointResponseOutput {
		return v.ReadWriteEndpoint
	}).(InstanceFailoverGroupReadWriteEndpointResponseOutput)
}

// Local replication role of the failover group instance.
func (o InstanceFailoverGroupOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) pulumi.StringOutput { return v.ReplicationRole }).(pulumi.StringOutput)
}

// Replication state of the failover group instance.
func (o InstanceFailoverGroupOutput) ReplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) pulumi.StringOutput { return v.ReplicationState }).(pulumi.StringOutput)
}

// Resource type.
func (o InstanceFailoverGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceFailoverGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceFailoverGroupOutput{})
}

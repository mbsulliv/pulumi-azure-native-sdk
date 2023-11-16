// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190907

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a read only following database.
type ReadOnlyFollowingDatabase struct {
	pulumi.CustomResourceState

	// The name of the attached database configuration cluster
	AttachedDatabaseConfigurationName pulumi.StringOutput `pulumi:"attachedDatabaseConfigurationName"`
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod pulumi.StringPtrOutput `pulumi:"hotCachePeriod"`
	// Kind of the database
	// Expected value is 'ReadOnlyFollowing'.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the leader cluster
	LeaderClusterResourceId pulumi.StringOutput `pulumi:"leaderClusterResourceId"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The principals modification kind of the database
	PrincipalsModificationKind pulumi.StringOutput `pulumi:"principalsModificationKind"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod pulumi.StringOutput `pulumi:"softDeletePeriod"`
	// The statistics of the database.
	Statistics DatabaseStatisticsResponseOutput `pulumi:"statistics"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReadOnlyFollowingDatabase registers a new resource with the given unique name, arguments, and options.
func NewReadOnlyFollowingDatabase(ctx *pulumi.Context,
	name string, args *ReadOnlyFollowingDatabaseArgs, opts ...pulumi.ResourceOption) (*ReadOnlyFollowingDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Kind = pulumi.StringPtr("ReadOnlyFollowing")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221229:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230815:ReadOnlyFollowingDatabase"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReadOnlyFollowingDatabase
	err := ctx.RegisterResource("azure-native:kusto/v20190907:ReadOnlyFollowingDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadOnlyFollowingDatabase gets an existing ReadOnlyFollowingDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadOnlyFollowingDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadOnlyFollowingDatabaseState, opts ...pulumi.ResourceOption) (*ReadOnlyFollowingDatabase, error) {
	var resource ReadOnlyFollowingDatabase
	err := ctx.ReadResource("azure-native:kusto/v20190907:ReadOnlyFollowingDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadOnlyFollowingDatabase resources.
type readOnlyFollowingDatabaseState struct {
}

type ReadOnlyFollowingDatabaseState struct {
}

func (ReadOnlyFollowingDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*readOnlyFollowingDatabaseState)(nil)).Elem()
}

type readOnlyFollowingDatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName *string `pulumi:"databaseName"`
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod *string `pulumi:"hotCachePeriod"`
	// Kind of the database
	// Expected value is 'ReadOnlyFollowing'.
	Kind *string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ReadOnlyFollowingDatabase resource.
type ReadOnlyFollowingDatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringPtrInput
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod pulumi.StringPtrInput
	// Kind of the database
	// Expected value is 'ReadOnlyFollowing'.
	Kind pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
}

func (ReadOnlyFollowingDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readOnlyFollowingDatabaseArgs)(nil)).Elem()
}

type ReadOnlyFollowingDatabaseInput interface {
	pulumi.Input

	ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput
	ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput
}

func (*ReadOnlyFollowingDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadOnlyFollowingDatabase)(nil)).Elem()
}

func (i *ReadOnlyFollowingDatabase) ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput {
	return i.ToReadOnlyFollowingDatabaseOutputWithContext(context.Background())
}

func (i *ReadOnlyFollowingDatabase) ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadOnlyFollowingDatabaseOutput)
}

type ReadOnlyFollowingDatabaseOutput struct{ *pulumi.OutputState }

func (ReadOnlyFollowingDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadOnlyFollowingDatabase)(nil)).Elem()
}

func (o ReadOnlyFollowingDatabaseOutput) ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput {
	return o
}

func (o ReadOnlyFollowingDatabaseOutput) ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput {
	return o
}

// The name of the attached database configuration cluster
func (o ReadOnlyFollowingDatabaseOutput) AttachedDatabaseConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.AttachedDatabaseConfigurationName }).(pulumi.StringOutput)
}

// The time the data should be kept in cache for fast queries in TimeSpan.
func (o ReadOnlyFollowingDatabaseOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringPtrOutput { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

// Kind of the database
// Expected value is 'ReadOnlyFollowing'.
func (o ReadOnlyFollowingDatabaseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the leader cluster
func (o ReadOnlyFollowingDatabaseOutput) LeaderClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.LeaderClusterResourceId }).(pulumi.StringOutput)
}

// Resource location.
func (o ReadOnlyFollowingDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ReadOnlyFollowingDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The principals modification kind of the database
func (o ReadOnlyFollowingDatabaseOutput) PrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.PrincipalsModificationKind }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o ReadOnlyFollowingDatabaseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The time the data should be kept before it stops being accessible to queries in TimeSpan.
func (o ReadOnlyFollowingDatabaseOutput) SoftDeletePeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.SoftDeletePeriod }).(pulumi.StringOutput)
}

// The statistics of the database.
func (o ReadOnlyFollowingDatabaseOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) DatabaseStatisticsResponseOutput { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ReadOnlyFollowingDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReadOnlyFollowingDatabaseOutput{})
}

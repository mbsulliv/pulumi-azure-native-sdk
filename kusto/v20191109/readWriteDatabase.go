// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191109

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a read write database.
type ReadWriteDatabase struct {
	pulumi.CustomResourceState

	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod pulumi.StringPtrOutput `pulumi:"hotCachePeriod"`
	// Indicates whether the database is followed.
	IsFollowed pulumi.StringOutput `pulumi:"isFollowed"`
	// Kind of the database
	// Expected value is 'ReadWrite'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod pulumi.StringPtrOutput `pulumi:"softDeletePeriod"`
	// The statistics of the database.
	Statistics DatabaseStatisticsResponseOutput `pulumi:"statistics"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReadWriteDatabase registers a new resource with the given unique name, arguments, and options.
func NewReadWriteDatabase(ctx *pulumi.Context,
	name string, args *ReadWriteDatabaseArgs, opts ...pulumi.ResourceOption) (*ReadWriteDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Kind = pulumi.String("ReadWrite")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221229:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230815:ReadWriteDatabase"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReadWriteDatabase
	err := ctx.RegisterResource("azure-native:kusto/v20191109:ReadWriteDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadWriteDatabase gets an existing ReadWriteDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadWriteDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadWriteDatabaseState, opts ...pulumi.ResourceOption) (*ReadWriteDatabase, error) {
	var resource ReadWriteDatabase
	err := ctx.ReadResource("azure-native:kusto/v20191109:ReadWriteDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadWriteDatabase resources.
type readWriteDatabaseState struct {
}

type ReadWriteDatabaseState struct {
}

func (ReadWriteDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*readWriteDatabaseState)(nil)).Elem()
}

type readWriteDatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName *string `pulumi:"databaseName"`
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod *string `pulumi:"hotCachePeriod"`
	// Kind of the database
	// Expected value is 'ReadWrite'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod *string `pulumi:"softDeletePeriod"`
}

// The set of arguments for constructing a ReadWriteDatabase resource.
type ReadWriteDatabaseArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringPtrInput
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod pulumi.StringPtrInput
	// Kind of the database
	// Expected value is 'ReadWrite'.
	Kind pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod pulumi.StringPtrInput
}

func (ReadWriteDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readWriteDatabaseArgs)(nil)).Elem()
}

type ReadWriteDatabaseInput interface {
	pulumi.Input

	ToReadWriteDatabaseOutput() ReadWriteDatabaseOutput
	ToReadWriteDatabaseOutputWithContext(ctx context.Context) ReadWriteDatabaseOutput
}

func (*ReadWriteDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadWriteDatabase)(nil)).Elem()
}

func (i *ReadWriteDatabase) ToReadWriteDatabaseOutput() ReadWriteDatabaseOutput {
	return i.ToReadWriteDatabaseOutputWithContext(context.Background())
}

func (i *ReadWriteDatabase) ToReadWriteDatabaseOutputWithContext(ctx context.Context) ReadWriteDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadWriteDatabaseOutput)
}

type ReadWriteDatabaseOutput struct{ *pulumi.OutputState }

func (ReadWriteDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadWriteDatabase)(nil)).Elem()
}

func (o ReadWriteDatabaseOutput) ToReadWriteDatabaseOutput() ReadWriteDatabaseOutput {
	return o
}

func (o ReadWriteDatabaseOutput) ToReadWriteDatabaseOutputWithContext(ctx context.Context) ReadWriteDatabaseOutput {
	return o
}

// The time the data should be kept in cache for fast queries in TimeSpan.
func (o ReadWriteDatabaseOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringPtrOutput { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

// Indicates whether the database is followed.
func (o ReadWriteDatabaseOutput) IsFollowed() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.IsFollowed }).(pulumi.StringOutput)
}

// Kind of the database
// Expected value is 'ReadWrite'.
func (o ReadWriteDatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o ReadWriteDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ReadWriteDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o ReadWriteDatabaseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The time the data should be kept before it stops being accessible to queries in TimeSpan.
func (o ReadWriteDatabaseOutput) SoftDeletePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringPtrOutput { return v.SoftDeletePeriod }).(pulumi.StringPtrOutput)
}

// The statistics of the database.
func (o ReadWriteDatabaseOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) DatabaseStatisticsResponseOutput { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ReadWriteDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReadWriteDatabaseOutput{})
}

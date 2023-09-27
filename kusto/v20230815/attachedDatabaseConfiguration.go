// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230815

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Class representing an attached database configuration.
type AttachedDatabaseConfiguration struct {
	pulumi.CustomResourceState

	// The list of databases from the clusterResourceId which are currently attached to the cluster.
	AttachedDatabaseNames pulumi.StringArrayOutput `pulumi:"attachedDatabaseNames"`
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId pulumi.StringOutput `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Overrides the original database name. Relevant only when attaching to a specific database.
	DatabaseNameOverride pulumi.StringPtrOutput `pulumi:"databaseNameOverride"`
	// Adds a prefix to the attached databases name. When following an entire cluster, that prefix would be added to all of the databases original names from leader cluster.
	DatabaseNamePrefix pulumi.StringPtrOutput `pulumi:"databaseNamePrefix"`
	// The default principals modification kind
	DefaultPrincipalsModificationKind pulumi.StringOutput `pulumi:"defaultPrincipalsModificationKind"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Table level sharing specifications
	TableLevelSharingProperties TableLevelSharingPropertiesResponsePtrOutput `pulumi:"tableLevelSharingProperties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAttachedDatabaseConfiguration registers a new resource with the given unique name, arguments, and options.
func NewAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, args *AttachedDatabaseConfigurationArgs, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ClusterResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterResourceId'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.DefaultPrincipalsModificationKind == nil {
		return nil, errors.New("invalid value for required argument 'DefaultPrincipalsModificationKind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221229:AttachedDatabaseConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:AttachedDatabaseConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AttachedDatabaseConfiguration
	err := ctx.RegisterResource("azure-native:kusto/v20230815:AttachedDatabaseConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedDatabaseConfiguration gets an existing AttachedDatabaseConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDatabaseConfigurationState, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	var resource AttachedDatabaseConfiguration
	err := ctx.ReadResource("azure-native:kusto/v20230815:AttachedDatabaseConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedDatabaseConfiguration resources.
type attachedDatabaseConfigurationState struct {
}

type AttachedDatabaseConfigurationState struct {
}

func (AttachedDatabaseConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationState)(nil)).Elem()
}

type attachedDatabaseConfigurationArgs struct {
	// The name of the attached database configuration.
	AttachedDatabaseConfigurationName *string `pulumi:"attachedDatabaseConfigurationName"`
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId string `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName string `pulumi:"databaseName"`
	// Overrides the original database name. Relevant only when attaching to a specific database.
	DatabaseNameOverride *string `pulumi:"databaseNameOverride"`
	// Adds a prefix to the attached databases name. When following an entire cluster, that prefix would be added to all of the databases original names from leader cluster.
	DatabaseNamePrefix *string `pulumi:"databaseNamePrefix"`
	// The default principals modification kind
	DefaultPrincipalsModificationKind string `pulumi:"defaultPrincipalsModificationKind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Table level sharing specifications
	TableLevelSharingProperties *TableLevelSharingProperties `pulumi:"tableLevelSharingProperties"`
}

// The set of arguments for constructing a AttachedDatabaseConfiguration resource.
type AttachedDatabaseConfigurationArgs struct {
	// The name of the attached database configuration.
	AttachedDatabaseConfigurationName pulumi.StringPtrInput
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId pulumi.StringInput
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName pulumi.StringInput
	// Overrides the original database name. Relevant only when attaching to a specific database.
	DatabaseNameOverride pulumi.StringPtrInput
	// Adds a prefix to the attached databases name. When following an entire cluster, that prefix would be added to all of the databases original names from leader cluster.
	DatabaseNamePrefix pulumi.StringPtrInput
	// The default principals modification kind
	DefaultPrincipalsModificationKind pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Table level sharing specifications
	TableLevelSharingProperties TableLevelSharingPropertiesPtrInput
}

func (AttachedDatabaseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationArgs)(nil)).Elem()
}

type AttachedDatabaseConfigurationInput interface {
	pulumi.Input

	ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput
	ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput
}

func (*AttachedDatabaseConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDatabaseConfiguration)(nil)).Elem()
}

func (i *AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput {
	return i.ToAttachedDatabaseConfigurationOutputWithContext(context.Background())
}

func (i *AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDatabaseConfigurationOutput)
}

func (i *AttachedDatabaseConfiguration) ToOutput(ctx context.Context) pulumix.Output[*AttachedDatabaseConfiguration] {
	return pulumix.Output[*AttachedDatabaseConfiguration]{
		OutputState: i.ToAttachedDatabaseConfigurationOutputWithContext(ctx).OutputState,
	}
}

type AttachedDatabaseConfigurationOutput struct{ *pulumi.OutputState }

func (AttachedDatabaseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDatabaseConfiguration)(nil)).Elem()
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput {
	return o
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput {
	return o
}

func (o AttachedDatabaseConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[*AttachedDatabaseConfiguration] {
	return pulumix.Output[*AttachedDatabaseConfiguration]{
		OutputState: o.OutputState,
	}
}

// The list of databases from the clusterResourceId which are currently attached to the cluster.
func (o AttachedDatabaseConfigurationOutput) AttachedDatabaseNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringArrayOutput { return v.AttachedDatabaseNames }).(pulumi.StringArrayOutput)
}

// The resource id of the cluster where the databases you would like to attach reside.
func (o AttachedDatabaseConfigurationOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.ClusterResourceId }).(pulumi.StringOutput)
}

// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
func (o AttachedDatabaseConfigurationOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// Overrides the original database name. Relevant only when attaching to a specific database.
func (o AttachedDatabaseConfigurationOutput) DatabaseNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringPtrOutput { return v.DatabaseNameOverride }).(pulumi.StringPtrOutput)
}

// Adds a prefix to the attached databases name. When following an entire cluster, that prefix would be added to all of the databases original names from leader cluster.
func (o AttachedDatabaseConfigurationOutput) DatabaseNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringPtrOutput { return v.DatabaseNamePrefix }).(pulumi.StringPtrOutput)
}

// The default principals modification kind
func (o AttachedDatabaseConfigurationOutput) DefaultPrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.DefaultPrincipalsModificationKind }).(pulumi.StringOutput)
}

// Resource location.
func (o AttachedDatabaseConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o AttachedDatabaseConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o AttachedDatabaseConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Table level sharing specifications
func (o AttachedDatabaseConfigurationOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) TableLevelSharingPropertiesResponsePtrOutput {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AttachedDatabaseConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedDatabaseConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttachedDatabaseConfigurationOutput{})
}

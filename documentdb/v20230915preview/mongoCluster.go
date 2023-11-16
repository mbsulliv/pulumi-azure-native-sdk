// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a mongo cluster resource.
type MongoCluster struct {
	pulumi.CustomResourceState

	// The administrator's login for the mongo cluster.
	AdministratorLogin pulumi.StringPtrOutput `pulumi:"administratorLogin"`
	// A status of the mongo cluster.
	ClusterStatus pulumi.StringOutput `pulumi:"clusterStatus"`
	// The default mongo connection string for the cluster.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Earliest restore timestamp in UTC ISO8601 format.
	EarliestRestoreTime pulumi.StringOutput `pulumi:"earliestRestoreTime"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of node group specs in the cluster.
	NodeGroupSpecs NodeGroupSpecResponseArrayOutput `pulumi:"nodeGroupSpecs"`
	// A provisioning state of the mongo cluster.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The Mongo DB server version. Defaults to the latest available version if not specified.
	ServerVersion pulumi.StringPtrOutput `pulumi:"serverVersion"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMongoCluster registers a new resource with the given unique name, arguments, and options.
func NewMongoCluster(ctx *pulumi.Context,
	name string, args *MongoClusterArgs, opts ...pulumi.ResourceOption) (*MongoCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.CreateMode == nil {
		args.CreateMode = pulumi.StringPtr("Default")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:MongoCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:MongoCluster"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:MongoCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MongoCluster
	err := ctx.RegisterResource("azure-native:documentdb/v20230915preview:MongoCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoCluster gets an existing MongoCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoClusterState, opts ...pulumi.ResourceOption) (*MongoCluster, error) {
	var resource MongoCluster
	err := ctx.ReadResource("azure-native:documentdb/v20230915preview:MongoCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoCluster resources.
type mongoClusterState struct {
}

type MongoClusterState struct {
}

func (MongoClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoClusterState)(nil)).Elem()
}

type mongoClusterArgs struct {
	// The administrator's login for the mongo cluster.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// The password of the administrator login.
	AdministratorLoginPassword *string `pulumi:"administratorLoginPassword"`
	// The mode to create a mongo cluster.
	CreateMode *string `pulumi:"createMode"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the mongo cluster.
	MongoClusterName *string `pulumi:"mongoClusterName"`
	// The list of node group specs in the cluster.
	NodeGroupSpecs []NodeGroupSpec `pulumi:"nodeGroupSpecs"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Parameters used for restore operations
	RestoreParameters *MongoClusterRestoreParameters `pulumi:"restoreParameters"`
	// The Mongo DB server version. Defaults to the latest available version if not specified.
	ServerVersion *string `pulumi:"serverVersion"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MongoCluster resource.
type MongoClusterArgs struct {
	// The administrator's login for the mongo cluster.
	AdministratorLogin pulumi.StringPtrInput
	// The password of the administrator login.
	AdministratorLoginPassword pulumi.StringPtrInput
	// The mode to create a mongo cluster.
	CreateMode pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the mongo cluster.
	MongoClusterName pulumi.StringPtrInput
	// The list of node group specs in the cluster.
	NodeGroupSpecs NodeGroupSpecArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Parameters used for restore operations
	RestoreParameters MongoClusterRestoreParametersPtrInput
	// The Mongo DB server version. Defaults to the latest available version if not specified.
	ServerVersion pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MongoClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoClusterArgs)(nil)).Elem()
}

type MongoClusterInput interface {
	pulumi.Input

	ToMongoClusterOutput() MongoClusterOutput
	ToMongoClusterOutputWithContext(ctx context.Context) MongoClusterOutput
}

func (*MongoCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoCluster)(nil)).Elem()
}

func (i *MongoCluster) ToMongoClusterOutput() MongoClusterOutput {
	return i.ToMongoClusterOutputWithContext(context.Background())
}

func (i *MongoCluster) ToMongoClusterOutputWithContext(ctx context.Context) MongoClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoClusterOutput)
}

type MongoClusterOutput struct{ *pulumi.OutputState }

func (MongoClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoCluster)(nil)).Elem()
}

func (o MongoClusterOutput) ToMongoClusterOutput() MongoClusterOutput {
	return o
}

func (o MongoClusterOutput) ToMongoClusterOutputWithContext(ctx context.Context) MongoClusterOutput {
	return o
}

// The administrator's login for the mongo cluster.
func (o MongoClusterOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// A status of the mongo cluster.
func (o MongoClusterOutput) ClusterStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.ClusterStatus }).(pulumi.StringOutput)
}

// The default mongo connection string for the cluster.
func (o MongoClusterOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// Earliest restore timestamp in UTC ISO8601 format.
func (o MongoClusterOutput) EarliestRestoreTime() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.EarliestRestoreTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o MongoClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MongoClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of node group specs in the cluster.
func (o MongoClusterOutput) NodeGroupSpecs() NodeGroupSpecResponseArrayOutput {
	return o.ApplyT(func(v *MongoCluster) NodeGroupSpecResponseArrayOutput { return v.NodeGroupSpecs }).(NodeGroupSpecResponseArrayOutput)
}

// A provisioning state of the mongo cluster.
func (o MongoClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Mongo DB server version. Defaults to the latest available version if not specified.
func (o MongoClusterOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringPtrOutput { return v.ServerVersion }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MongoClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MongoCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MongoClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MongoClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MongoClusterOutput{})
}

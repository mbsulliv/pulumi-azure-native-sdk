// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221229

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a CosmosDb data connection.
type CosmosDbDataConnection struct {
	pulumi.CustomResourceState

	// The resource ID of the Cosmos DB account used to create the data connection.
	CosmosDbAccountResourceId pulumi.StringOutput `pulumi:"cosmosDbAccountResourceId"`
	// The name of an existing container in the Cosmos DB database.
	CosmosDbContainer pulumi.StringOutput `pulumi:"cosmosDbContainer"`
	// The name of an existing database in the Cosmos DB account.
	CosmosDbDatabase pulumi.StringOutput `pulumi:"cosmosDbDatabase"`
	// Kind of the endpoint for the data connection
	// Expected value is 'CosmosDb'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The object ID of the managed identity resource.
	ManagedIdentityObjectId pulumi.StringOutput `pulumi:"managedIdentityObjectId"`
	// The resource ID of a managed system or user-assigned identity. The identity is used to authenticate with Cosmos DB.
	ManagedIdentityResourceId pulumi.StringOutput `pulumi:"managedIdentityResourceId"`
	// The name of an existing mapping rule to use when ingesting the retrieved data.
	MappingRuleName pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Optional. If defined, the data connection retrieves Cosmos DB documents created or updated after the specified retrieval start date.
	RetrievalStartDate pulumi.StringPtrOutput `pulumi:"retrievalStartDate"`
	// The case-sensitive name of the existing target table in your cluster. Retrieved data is ingested into this table.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCosmosDbDataConnection registers a new resource with the given unique name, arguments, and options.
func NewCosmosDbDataConnection(ctx *pulumi.Context,
	name string, args *CosmosDbDataConnectionArgs, opts ...pulumi.ResourceOption) (*CosmosDbDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.CosmosDbAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'CosmosDbAccountResourceId'")
	}
	if args.CosmosDbContainer == nil {
		return nil, errors.New("invalid value for required argument 'CosmosDbContainer'")
	}
	if args.CosmosDbDatabase == nil {
		return nil, errors.New("invalid value for required argument 'CosmosDbDatabase'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ManagedIdentityResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ManagedIdentityResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	args.Kind = pulumi.String("CosmosDb")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:CosmosDbDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230815:CosmosDbDataConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CosmosDbDataConnection
	err := ctx.RegisterResource("azure-native:kusto/v20221229:CosmosDbDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCosmosDbDataConnection gets an existing CosmosDbDataConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCosmosDbDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CosmosDbDataConnectionState, opts ...pulumi.ResourceOption) (*CosmosDbDataConnection, error) {
	var resource CosmosDbDataConnection
	err := ctx.ReadResource("azure-native:kusto/v20221229:CosmosDbDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CosmosDbDataConnection resources.
type cosmosDbDataConnectionState struct {
}

type CosmosDbDataConnectionState struct {
}

func (CosmosDbDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*cosmosDbDataConnectionState)(nil)).Elem()
}

type cosmosDbDataConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The resource ID of the Cosmos DB account used to create the data connection.
	CosmosDbAccountResourceId string `pulumi:"cosmosDbAccountResourceId"`
	// The name of an existing container in the Cosmos DB database.
	CosmosDbContainer string `pulumi:"cosmosDbContainer"`
	// The name of an existing database in the Cosmos DB account.
	CosmosDbDatabase string `pulumi:"cosmosDbDatabase"`
	// The name of the data connection.
	DataConnectionName *string `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// Kind of the endpoint for the data connection
	// Expected value is 'CosmosDb'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The resource ID of a managed system or user-assigned identity. The identity is used to authenticate with Cosmos DB.
	ManagedIdentityResourceId string `pulumi:"managedIdentityResourceId"`
	// The name of an existing mapping rule to use when ingesting the retrieved data.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Optional. If defined, the data connection retrieves Cosmos DB documents created or updated after the specified retrieval start date.
	RetrievalStartDate *string `pulumi:"retrievalStartDate"`
	// The case-sensitive name of the existing target table in your cluster. Retrieved data is ingested into this table.
	TableName string `pulumi:"tableName"`
}

// The set of arguments for constructing a CosmosDbDataConnection resource.
type CosmosDbDataConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The resource ID of the Cosmos DB account used to create the data connection.
	CosmosDbAccountResourceId pulumi.StringInput
	// The name of an existing container in the Cosmos DB database.
	CosmosDbContainer pulumi.StringInput
	// The name of an existing database in the Cosmos DB account.
	CosmosDbDatabase pulumi.StringInput
	// The name of the data connection.
	DataConnectionName pulumi.StringPtrInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput
	// Kind of the endpoint for the data connection
	// Expected value is 'CosmosDb'.
	Kind pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The resource ID of a managed system or user-assigned identity. The identity is used to authenticate with Cosmos DB.
	ManagedIdentityResourceId pulumi.StringInput
	// The name of an existing mapping rule to use when ingesting the retrieved data.
	MappingRuleName pulumi.StringPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// Optional. If defined, the data connection retrieves Cosmos DB documents created or updated after the specified retrieval start date.
	RetrievalStartDate pulumi.StringPtrInput
	// The case-sensitive name of the existing target table in your cluster. Retrieved data is ingested into this table.
	TableName pulumi.StringInput
}

func (CosmosDbDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cosmosDbDataConnectionArgs)(nil)).Elem()
}

type CosmosDbDataConnectionInput interface {
	pulumi.Input

	ToCosmosDbDataConnectionOutput() CosmosDbDataConnectionOutput
	ToCosmosDbDataConnectionOutputWithContext(ctx context.Context) CosmosDbDataConnectionOutput
}

func (*CosmosDbDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbDataConnection)(nil)).Elem()
}

func (i *CosmosDbDataConnection) ToCosmosDbDataConnectionOutput() CosmosDbDataConnectionOutput {
	return i.ToCosmosDbDataConnectionOutputWithContext(context.Background())
}

func (i *CosmosDbDataConnection) ToCosmosDbDataConnectionOutputWithContext(ctx context.Context) CosmosDbDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbDataConnectionOutput)
}

type CosmosDbDataConnectionOutput struct{ *pulumi.OutputState }

func (CosmosDbDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbDataConnection)(nil)).Elem()
}

func (o CosmosDbDataConnectionOutput) ToCosmosDbDataConnectionOutput() CosmosDbDataConnectionOutput {
	return o
}

func (o CosmosDbDataConnectionOutput) ToCosmosDbDataConnectionOutputWithContext(ctx context.Context) CosmosDbDataConnectionOutput {
	return o
}

// The resource ID of the Cosmos DB account used to create the data connection.
func (o CosmosDbDataConnectionOutput) CosmosDbAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.CosmosDbAccountResourceId }).(pulumi.StringOutput)
}

// The name of an existing container in the Cosmos DB database.
func (o CosmosDbDataConnectionOutput) CosmosDbContainer() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.CosmosDbContainer }).(pulumi.StringOutput)
}

// The name of an existing database in the Cosmos DB account.
func (o CosmosDbDataConnectionOutput) CosmosDbDatabase() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.CosmosDbDatabase }).(pulumi.StringOutput)
}

// Kind of the endpoint for the data connection
// Expected value is 'CosmosDb'.
func (o CosmosDbDataConnectionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o CosmosDbDataConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The object ID of the managed identity resource.
func (o CosmosDbDataConnectionOutput) ManagedIdentityObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.ManagedIdentityObjectId }).(pulumi.StringOutput)
}

// The resource ID of a managed system or user-assigned identity. The identity is used to authenticate with Cosmos DB.
func (o CosmosDbDataConnectionOutput) ManagedIdentityResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.ManagedIdentityResourceId }).(pulumi.StringOutput)
}

// The name of an existing mapping rule to use when ingesting the retrieved data.
func (o CosmosDbDataConnectionOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringPtrOutput { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o CosmosDbDataConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o CosmosDbDataConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Optional. If defined, the data connection retrieves Cosmos DB documents created or updated after the specified retrieval start date.
func (o CosmosDbDataConnectionOutput) RetrievalStartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringPtrOutput { return v.RetrievalStartDate }).(pulumi.StringPtrOutput)
}

// The case-sensitive name of the existing target table in your cluster. Retrieved data is ingested into this table.
func (o CosmosDbDataConnectionOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CosmosDbDataConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CosmosDbDataConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CosmosDbDataConnectionOutput{})
}

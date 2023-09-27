// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about a database in a RedisEnterprise cluster.
func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:cache/v20230301preview:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a database on the RedisEnterprise cluster
type LookupDatabaseResult struct {
	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted.
	ClientProtocol *string `pulumi:"clientProtocol"`
	// Clustering policy - default is OSSCluster. Specified at create time.
	ClusteringPolicy *string `pulumi:"clusteringPolicy"`
	// Redis eviction policy - default is VolatileLRU
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Optional set of properties to configure geo replication for this database.
	GeoReplication *DatabasePropertiesResponseGeoReplication `pulumi:"geoReplication"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Optional set of redis modules to enable in this database - modules can only be added at creation time.
	Modules []ModuleResponse `pulumi:"modules"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Persistence settings
	Persistence *PersistenceResponse `pulumi:"persistence"`
	// TCP port of the database endpoint. Specified at create time. Defaults to an available port.
	Port *int `pulumi:"port"`
	// Current provisioning status of the database
	ProvisioningState string `pulumi:"provisioningState"`
	// Current resource status of the database
	ResourceState string `pulumi:"resourceState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			var s LookupDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseResultOutput)
}

type LookupDatabaseOutputArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}

// Describes a database on the RedisEnterprise cluster
type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDatabaseResult] {
	return pulumix.Output[LookupDatabaseResult]{
		OutputState: o.OutputState,
	}
}

// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted.
func (o LookupDatabaseResultOutput) ClientProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ClientProtocol }).(pulumi.StringPtrOutput)
}

// Clustering policy - default is OSSCluster. Specified at create time.
func (o LookupDatabaseResultOutput) ClusteringPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.ClusteringPolicy }).(pulumi.StringPtrOutput)
}

// Redis eviction policy - default is VolatileLRU
func (o LookupDatabaseResultOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

// Optional set of properties to configure geo replication for this database.
func (o LookupDatabaseResultOutput) GeoReplication() DatabasePropertiesResponseGeoReplicationPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *DatabasePropertiesResponseGeoReplication { return v.GeoReplication }).(DatabasePropertiesResponseGeoReplicationPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// Optional set of redis modules to enable in this database - modules can only be added at creation time.
func (o LookupDatabaseResultOutput) Modules() ModuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDatabaseResult) []ModuleResponse { return v.Modules }).(ModuleResponseArrayOutput)
}

// The name of the resource
func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Persistence settings
func (o LookupDatabaseResultOutput) Persistence() PersistenceResponsePtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *PersistenceResponse { return v.Persistence }).(PersistenceResponsePtrOutput)
}

// TCP port of the database endpoint. Specified at create time. Defaults to an available port.
func (o LookupDatabaseResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// Current provisioning status of the database
func (o LookupDatabaseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Current resource status of the database
func (o LookupDatabaseResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDatabaseResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDatabaseResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}

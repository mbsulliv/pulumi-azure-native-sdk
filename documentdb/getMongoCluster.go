// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about a mongo cluster.
// Azure REST API version: 2023-03-15-preview.
func LookupMongoCluster(ctx *pulumi.Context, args *LookupMongoClusterArgs, opts ...pulumi.InvokeOption) (*LookupMongoClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMongoClusterResult
	err := ctx.Invoke("azure-native:documentdb:getMongoCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoClusterArgs struct {
	// The name of the mongo cluster.
	MongoClusterName string `pulumi:"mongoClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a mongo cluster resource.
type LookupMongoClusterResult struct {
	// The administrator's login for the mongo cluster.
	AdministratorLogin *string `pulumi:"administratorLogin"`
	// A status of the mongo cluster.
	ClusterStatus string `pulumi:"clusterStatus"`
	// The default mongo connection string for the cluster.
	ConnectionString string `pulumi:"connectionString"`
	// Earliest restore timestamp in UTC ISO8601 format.
	EarliestRestoreTime string `pulumi:"earliestRestoreTime"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The list of node group specs in the cluster.
	NodeGroupSpecs []NodeGroupSpecResponse `pulumi:"nodeGroupSpecs"`
	// A provisioning state of the mongo cluster.
	ProvisioningState string `pulumi:"provisioningState"`
	// The Mongo DB server version. Defaults to the latest available version if not specified.
	ServerVersion *string `pulumi:"serverVersion"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMongoClusterOutput(ctx *pulumi.Context, args LookupMongoClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMongoClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoClusterResult, error) {
			args := v.(LookupMongoClusterArgs)
			r, err := LookupMongoCluster(ctx, &args, opts...)
			var s LookupMongoClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoClusterResultOutput)
}

type LookupMongoClusterOutputArgs struct {
	// The name of the mongo cluster.
	MongoClusterName pulumi.StringInput `pulumi:"mongoClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoClusterArgs)(nil)).Elem()
}

// Represents a mongo cluster resource.
type LookupMongoClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMongoClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoClusterResult)(nil)).Elem()
}

func (o LookupMongoClusterResultOutput) ToLookupMongoClusterResultOutput() LookupMongoClusterResultOutput {
	return o
}

func (o LookupMongoClusterResultOutput) ToLookupMongoClusterResultOutputWithContext(ctx context.Context) LookupMongoClusterResultOutput {
	return o
}

func (o LookupMongoClusterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMongoClusterResult] {
	return pulumix.Output[LookupMongoClusterResult]{
		OutputState: o.OutputState,
	}
}

// The administrator's login for the mongo cluster.
func (o LookupMongoClusterResultOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) *string { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// A status of the mongo cluster.
func (o LookupMongoClusterResultOutput) ClusterStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.ClusterStatus }).(pulumi.StringOutput)
}

// The default mongo connection string for the cluster.
func (o LookupMongoClusterResultOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.ConnectionString }).(pulumi.StringOutput)
}

// Earliest restore timestamp in UTC ISO8601 format.
func (o LookupMongoClusterResultOutput) EarliestRestoreTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.EarliestRestoreTime }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupMongoClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupMongoClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMongoClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of node group specs in the cluster.
func (o LookupMongoClusterResultOutput) NodeGroupSpecs() NodeGroupSpecResponseArrayOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) []NodeGroupSpecResponse { return v.NodeGroupSpecs }).(NodeGroupSpecResponseArrayOutput)
}

// A provisioning state of the mongo cluster.
func (o LookupMongoClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Mongo DB server version. Defaults to the latest available version if not specified.
func (o LookupMongoClusterResultOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) *string { return v.ServerVersion }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMongoClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupMongoClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMongoClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoClusterResultOutput{})
}

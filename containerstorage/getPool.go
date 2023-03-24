// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Pool
// API Version: 2023-03-01-preview.
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:containerstorage:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPoolArgs struct {
	// Pool Object
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Pool resource
type LookupPoolResult struct {
	// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups. For local and standard this must be a single reference. For portable there can be many.
	Assignments []string `pulumi:"assignments"`
	// Disk Pool Properties
	DiskPoolProperties *DiskPoolPropertiesResponse `pulumi:"diskPoolProperties"`
	// Elastic San Pool Properties
	ElasticSanPoolProperties ElasticSanPoolPropertiesResponse `pulumi:"elasticSanPoolProperties"`
	// Ephemeral Pool Properties
	EphemeralPoolProperties *EphemeralPoolPropertiesResponse `pulumi:"ephemeralPoolProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Initial capacity of the pool in GiB.
	PoolCapacityGiB float64 `pulumi:"poolCapacityGiB"`
	// Type of the Pool: ephemeral, disk, managed, or elasticsan.
	PoolType float64 `pulumi:"poolType"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// List of availability zones that resources can be created in.
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupPoolResult
func (val *LookupPoolResult) Defaults() *LookupPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DiskPoolProperties = tmp.DiskPoolProperties.Defaults()

	tmp.EphemeralPoolProperties = tmp.EphemeralPoolProperties.Defaults()

	return &tmp
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPoolResult, error) {
			args := v.(LookupPoolArgs)
			r, err := LookupPool(ctx, &args, opts...)
			var s LookupPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPoolResultOutput)
}

type LookupPoolOutputArgs struct {
	// Pool Object
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}

// Pool resource
type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups. For local and standard this must be a single reference. For portable there can be many.
func (o LookupPoolResultOutput) Assignments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []string { return v.Assignments }).(pulumi.StringArrayOutput)
}

// Disk Pool Properties
func (o LookupPoolResultOutput) DiskPoolProperties() DiskPoolPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *DiskPoolPropertiesResponse { return v.DiskPoolProperties }).(DiskPoolPropertiesResponsePtrOutput)
}

// Elastic San Pool Properties
func (o LookupPoolResultOutput) ElasticSanPoolProperties() ElasticSanPoolPropertiesResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) ElasticSanPoolPropertiesResponse { return v.ElasticSanPoolProperties }).(ElasticSanPoolPropertiesResponseOutput)
}

// Ephemeral Pool Properties
func (o LookupPoolResultOutput) EphemeralPoolProperties() EphemeralPoolPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *EphemeralPoolPropertiesResponse { return v.EphemeralPoolProperties }).(EphemeralPoolPropertiesResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Initial capacity of the pool in GiB.
func (o LookupPoolResultOutput) PoolCapacityGiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPoolResult) float64 { return v.PoolCapacityGiB }).(pulumi.Float64Output)
}

// Type of the Pool: ephemeral, disk, managed, or elasticsan.
func (o LookupPoolResultOutput) PoolType() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPoolResult) float64 { return v.PoolType }).(pulumi.Float64Output)
}

// The status of the last operation.
func (o LookupPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// List of availability zones that resources can be created in.
func (o LookupPoolResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Big Data pool.
func LookupBigDataPool(ctx *pulumi.Context, args *LookupBigDataPoolArgs, opts ...pulumi.InvokeOption) (*LookupBigDataPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBigDataPoolResult
	err := ctx.Invoke("azure-native:synapse/v20210601:getBigDataPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBigDataPoolArgs struct {
	// Big Data pool name
	BigDataPoolName string `pulumi:"bigDataPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// A Big Data pool
type LookupBigDataPoolResult struct {
	// Auto-pausing properties
	AutoPause *AutoPausePropertiesResponse `pulumi:"autoPause"`
	// Auto-scaling properties
	AutoScale *AutoScalePropertiesResponse `pulumi:"autoScale"`
	// The cache size
	CacheSize int `pulumi:"cacheSize"`
	// The time when the Big Data pool was created.
	CreationDate string `pulumi:"creationDate"`
	// List of custom libraries/packages associated with the spark pool.
	CustomLibraries []LibraryInfoResponse `pulumi:"customLibraries"`
	// The default folder where Spark logs will be written.
	DefaultSparkLogFolder *string `pulumi:"defaultSparkLogFolder"`
	// Dynamic Executor Allocation
	DynamicExecutorAllocation *DynamicExecutorAllocationResponse `pulumi:"dynamicExecutorAllocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Whether autotune is required or not.
	IsAutotuneEnabled *bool `pulumi:"isAutotuneEnabled"`
	// Whether compute isolation is required or not.
	IsComputeIsolationEnabled *bool `pulumi:"isComputeIsolationEnabled"`
	// The time when the Big Data pool was updated successfully.
	LastSucceededTimestamp string `pulumi:"lastSucceededTimestamp"`
	// Library version requirements
	LibraryRequirements *LibraryRequirementsResponse `pulumi:"libraryRequirements"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The number of nodes in the Big Data pool.
	NodeCount *int `pulumi:"nodeCount"`
	// The level of compute power that each node in the Big Data pool has.
	NodeSize *string `pulumi:"nodeSize"`
	// The kind of nodes that the Big Data pool provides.
	NodeSizeFamily *string `pulumi:"nodeSizeFamily"`
	// The state of the Big Data pool.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Whether session level packages enabled.
	SessionLevelPackagesEnabled *bool `pulumi:"sessionLevelPackagesEnabled"`
	// Spark configuration file to specify additional properties
	SparkConfigProperties *SparkConfigPropertiesResponse `pulumi:"sparkConfigProperties"`
	// The Spark events folder
	SparkEventsFolder *string `pulumi:"sparkEventsFolder"`
	// The Apache Spark version.
	SparkVersion *string `pulumi:"sparkVersion"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBigDataPoolOutput(ctx *pulumi.Context, args LookupBigDataPoolOutputArgs, opts ...pulumi.InvokeOption) LookupBigDataPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBigDataPoolResult, error) {
			args := v.(LookupBigDataPoolArgs)
			r, err := LookupBigDataPool(ctx, &args, opts...)
			var s LookupBigDataPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBigDataPoolResultOutput)
}

type LookupBigDataPoolOutputArgs struct {
	// Big Data pool name
	BigDataPoolName pulumi.StringInput `pulumi:"bigDataPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBigDataPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBigDataPoolArgs)(nil)).Elem()
}

// A Big Data pool
type LookupBigDataPoolResultOutput struct{ *pulumi.OutputState }

func (LookupBigDataPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBigDataPoolResult)(nil)).Elem()
}

func (o LookupBigDataPoolResultOutput) ToLookupBigDataPoolResultOutput() LookupBigDataPoolResultOutput {
	return o
}

func (o LookupBigDataPoolResultOutput) ToLookupBigDataPoolResultOutputWithContext(ctx context.Context) LookupBigDataPoolResultOutput {
	return o
}

func (o LookupBigDataPoolResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBigDataPoolResult] {
	return pulumix.Output[LookupBigDataPoolResult]{
		OutputState: o.OutputState,
	}
}

// Auto-pausing properties
func (o LookupBigDataPoolResultOutput) AutoPause() AutoPausePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *AutoPausePropertiesResponse { return v.AutoPause }).(AutoPausePropertiesResponsePtrOutput)
}

// Auto-scaling properties
func (o LookupBigDataPoolResultOutput) AutoScale() AutoScalePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *AutoScalePropertiesResponse { return v.AutoScale }).(AutoScalePropertiesResponsePtrOutput)
}

// The cache size
func (o LookupBigDataPoolResultOutput) CacheSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) int { return v.CacheSize }).(pulumi.IntOutput)
}

// The time when the Big Data pool was created.
func (o LookupBigDataPoolResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// List of custom libraries/packages associated with the spark pool.
func (o LookupBigDataPoolResultOutput) CustomLibraries() LibraryInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) []LibraryInfoResponse { return v.CustomLibraries }).(LibraryInfoResponseArrayOutput)
}

// The default folder where Spark logs will be written.
func (o LookupBigDataPoolResultOutput) DefaultSparkLogFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.DefaultSparkLogFolder }).(pulumi.StringPtrOutput)
}

// Dynamic Executor Allocation
func (o LookupBigDataPoolResultOutput) DynamicExecutorAllocation() DynamicExecutorAllocationResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *DynamicExecutorAllocationResponse { return v.DynamicExecutorAllocation }).(DynamicExecutorAllocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBigDataPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether autotune is required or not.
func (o LookupBigDataPoolResultOutput) IsAutotuneEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *bool { return v.IsAutotuneEnabled }).(pulumi.BoolPtrOutput)
}

// Whether compute isolation is required or not.
func (o LookupBigDataPoolResultOutput) IsComputeIsolationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *bool { return v.IsComputeIsolationEnabled }).(pulumi.BoolPtrOutput)
}

// The time when the Big Data pool was updated successfully.
func (o LookupBigDataPoolResultOutput) LastSucceededTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.LastSucceededTimestamp }).(pulumi.StringOutput)
}

// Library version requirements
func (o LookupBigDataPoolResultOutput) LibraryRequirements() LibraryRequirementsResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *LibraryRequirementsResponse { return v.LibraryRequirements }).(LibraryRequirementsResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupBigDataPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBigDataPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of nodes in the Big Data pool.
func (o LookupBigDataPoolResultOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

// The level of compute power that each node in the Big Data pool has.
func (o LookupBigDataPoolResultOutput) NodeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.NodeSize }).(pulumi.StringPtrOutput)
}

// The kind of nodes that the Big Data pool provides.
func (o LookupBigDataPoolResultOutput) NodeSizeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.NodeSizeFamily }).(pulumi.StringPtrOutput)
}

// The state of the Big Data pool.
func (o LookupBigDataPoolResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Whether session level packages enabled.
func (o LookupBigDataPoolResultOutput) SessionLevelPackagesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *bool { return v.SessionLevelPackagesEnabled }).(pulumi.BoolPtrOutput)
}

// Spark configuration file to specify additional properties
func (o LookupBigDataPoolResultOutput) SparkConfigProperties() SparkConfigPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *SparkConfigPropertiesResponse { return v.SparkConfigProperties }).(SparkConfigPropertiesResponsePtrOutput)
}

// The Spark events folder
func (o LookupBigDataPoolResultOutput) SparkEventsFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.SparkEventsFolder }).(pulumi.StringPtrOutput)
}

// The Apache Spark version.
func (o LookupBigDataPoolResultOutput) SparkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.SparkVersion }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupBigDataPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBigDataPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBigDataPoolResultOutput{})
}

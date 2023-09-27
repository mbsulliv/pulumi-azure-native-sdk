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

// Implements resourcePool GET method.
func LookupResourcePool(ctx *pulumi.Context, args *LookupResourcePoolArgs, opts ...pulumi.InvokeOption) (*LookupResourcePoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupResourcePoolResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20230301preview:getResourcePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourcePoolArgs struct {
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the resourcePool.
	ResourcePoolName string `pulumi:"resourcePoolName"`
}

// Define the resourcePool.
type LookupResourcePoolResult struct {
	// Gets the max CPU usage across all cores on the pool in MHz.
	CpuCapacityMHz float64 `pulumi:"cpuCapacityMHz"`
	// Gets or sets CPULimitMHz which specifies a CPU usage limit in MHz.
	// Utilization will not exceed this limit even if there are available resources.
	CpuLimitMHz float64 `pulumi:"cpuLimitMHz"`
	// Gets the used CPU usage across all cores on the pool in MHz.
	CpuOverallUsageMHz float64 `pulumi:"cpuOverallUsageMHz"`
	// Gets or sets CPUReservationMHz which specifies the CPU size in MHz that is guaranteed
	// to be available.
	CpuReservationMHz float64 `pulumi:"cpuReservationMHz"`
	// Gets or sets CPUSharesLevel which specifies the CPU allocation level for this pool.
	// This property is used in relative allocation between resource consumers.
	CpuSharesLevel string `pulumi:"cpuSharesLevel"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName string `pulumi:"customResourceName"`
	// Gets the datastore ARM ids.
	DatastoreIds []string `pulumi:"datastoreIds"`
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Gets or sets the Id.
	Id string `pulumi:"id"`
	// Gets or sets the inventory Item ID for the resource pool.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Gets the total amount of physical memory on the pool in GB.
	MemCapacityGB float64 `pulumi:"memCapacityGB"`
	// Gets or sets MemLimitMB specifies a memory usage limit in megabytes.
	// Utilization will not exceed the specified limit even if there are available resources.
	MemLimitMB float64 `pulumi:"memLimitMB"`
	// Gets the used physical memory on the pool in GB.
	MemOverallUsageGB float64 `pulumi:"memOverallUsageGB"`
	// Gets or sets MemReservationMB which specifies the guaranteed available memory in
	// megabytes.
	MemReservationMB float64 `pulumi:"memReservationMB"`
	// Gets or sets CPUSharesLevel which specifies the memory allocation level for this pool.
	// This property is used in relative allocation between resource consumers.
	MemSharesLevel string `pulumi:"memSharesLevel"`
	// Gets or sets the vCenter Managed Object name for the resource pool.
	MoName string `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the resource pool.
	MoRefId *string `pulumi:"moRefId"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Gets the network ARM ids.
	NetworkIds []string `pulumi:"networkIds"`
	// Gets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource status information.
	Statuses []ResourceStatusResponse `pulumi:"statuses"`
	// The system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid string `pulumi:"uuid"`
	// Gets or sets the ARM Id of the vCenter resource in which this resource pool resides.
	VCenterId *string `pulumi:"vCenterId"`
}

func LookupResourcePoolOutput(ctx *pulumi.Context, args LookupResourcePoolOutputArgs, opts ...pulumi.InvokeOption) LookupResourcePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourcePoolResult, error) {
			args := v.(LookupResourcePoolArgs)
			r, err := LookupResourcePool(ctx, &args, opts...)
			var s LookupResourcePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourcePoolResultOutput)
}

type LookupResourcePoolOutputArgs struct {
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the resourcePool.
	ResourcePoolName pulumi.StringInput `pulumi:"resourcePoolName"`
}

func (LookupResourcePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePoolArgs)(nil)).Elem()
}

// Define the resourcePool.
type LookupResourcePoolResultOutput struct{ *pulumi.OutputState }

func (LookupResourcePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePoolResult)(nil)).Elem()
}

func (o LookupResourcePoolResultOutput) ToLookupResourcePoolResultOutput() LookupResourcePoolResultOutput {
	return o
}

func (o LookupResourcePoolResultOutput) ToLookupResourcePoolResultOutputWithContext(ctx context.Context) LookupResourcePoolResultOutput {
	return o
}

func (o LookupResourcePoolResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupResourcePoolResult] {
	return pulumix.Output[LookupResourcePoolResult]{
		OutputState: o.OutputState,
	}
}

// Gets the max CPU usage across all cores on the pool in MHz.
func (o LookupResourcePoolResultOutput) CpuCapacityMHz() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.CpuCapacityMHz }).(pulumi.Float64Output)
}

// Gets or sets CPULimitMHz which specifies a CPU usage limit in MHz.
// Utilization will not exceed this limit even if there are available resources.
func (o LookupResourcePoolResultOutput) CpuLimitMHz() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.CpuLimitMHz }).(pulumi.Float64Output)
}

// Gets the used CPU usage across all cores on the pool in MHz.
func (o LookupResourcePoolResultOutput) CpuOverallUsageMHz() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.CpuOverallUsageMHz }).(pulumi.Float64Output)
}

// Gets or sets CPUReservationMHz which specifies the CPU size in MHz that is guaranteed
// to be available.
func (o LookupResourcePoolResultOutput) CpuReservationMHz() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.CpuReservationMHz }).(pulumi.Float64Output)
}

// Gets or sets CPUSharesLevel which specifies the CPU allocation level for this pool.
// This property is used in relative allocation between resource consumers.
func (o LookupResourcePoolResultOutput) CpuSharesLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.CpuSharesLevel }).(pulumi.StringOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o LookupResourcePoolResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets the datastore ARM ids.
func (o LookupResourcePoolResultOutput) DatastoreIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) []string { return v.DatastoreIds }).(pulumi.StringArrayOutput)
}

// Gets or sets the extended location.
func (o LookupResourcePoolResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the Id.
func (o LookupResourcePoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the inventory Item ID for the resource pool.
func (o LookupResourcePoolResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupResourcePoolResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o LookupResourcePoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets the total amount of physical memory on the pool in GB.
func (o LookupResourcePoolResultOutput) MemCapacityGB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.MemCapacityGB }).(pulumi.Float64Output)
}

// Gets or sets MemLimitMB specifies a memory usage limit in megabytes.
// Utilization will not exceed the specified limit even if there are available resources.
func (o LookupResourcePoolResultOutput) MemLimitMB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.MemLimitMB }).(pulumi.Float64Output)
}

// Gets the used physical memory on the pool in GB.
func (o LookupResourcePoolResultOutput) MemOverallUsageGB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.MemOverallUsageGB }).(pulumi.Float64Output)
}

// Gets or sets MemReservationMB which specifies the guaranteed available memory in
// megabytes.
func (o LookupResourcePoolResultOutput) MemReservationMB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupResourcePoolResult) float64 { return v.MemReservationMB }).(pulumi.Float64Output)
}

// Gets or sets CPUSharesLevel which specifies the memory allocation level for this pool.
// This property is used in relative allocation between resource consumers.
func (o LookupResourcePoolResultOutput) MemSharesLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.MemSharesLevel }).(pulumi.StringOutput)
}

// Gets or sets the vCenter Managed Object name for the resource pool.
func (o LookupResourcePoolResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the resource pool.
func (o LookupResourcePoolResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o LookupResourcePoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the network ARM ids.
func (o LookupResourcePoolResultOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) []string { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

// Gets the provisioning state.
func (o LookupResourcePoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o LookupResourcePoolResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o LookupResourcePoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o LookupResourcePoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o LookupResourcePoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o LookupResourcePoolResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this resource pool resides.
func (o LookupResourcePoolResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourcePoolResultOutput{})
}

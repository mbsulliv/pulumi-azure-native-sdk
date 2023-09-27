// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Define the resourcePool.
type ResourcePool struct {
	pulumi.CustomResourceState

	// Gets the max CPU usage across all cores on the pool in MHz.
	CpuCapacityMHz pulumi.Float64Output `pulumi:"cpuCapacityMHz"`
	// Gets or sets CPULimitMHz which specifies a CPU usage limit in MHz.
	// Utilization will not exceed this limit even if there are available resources.
	CpuLimitMHz pulumi.Float64Output `pulumi:"cpuLimitMHz"`
	// Gets the used CPU usage across all cores on the pool in MHz.
	CpuOverallUsageMHz pulumi.Float64Output `pulumi:"cpuOverallUsageMHz"`
	// Gets or sets CPUReservationMHz which specifies the CPU size in MHz that is guaranteed
	// to be available.
	CpuReservationMHz pulumi.Float64Output `pulumi:"cpuReservationMHz"`
	// Gets or sets CPUSharesLevel which specifies the CPU allocation level for this pool.
	// This property is used in relative allocation between resource consumers.
	CpuSharesLevel pulumi.StringOutput `pulumi:"cpuSharesLevel"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName pulumi.StringOutput `pulumi:"customResourceName"`
	// Gets the datastore ARM ids.
	DatastoreIds pulumi.StringArrayOutput `pulumi:"datastoreIds"`
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the resource pool.
	InventoryItemId pulumi.StringPtrOutput `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Gets the total amount of physical memory on the pool in GB.
	MemCapacityGB pulumi.Float64Output `pulumi:"memCapacityGB"`
	// Gets or sets MemLimitMB specifies a memory usage limit in megabytes.
	// Utilization will not exceed the specified limit even if there are available resources.
	MemLimitMB pulumi.Float64Output `pulumi:"memLimitMB"`
	// Gets the used physical memory on the pool in GB.
	MemOverallUsageGB pulumi.Float64Output `pulumi:"memOverallUsageGB"`
	// Gets or sets MemReservationMB which specifies the guaranteed available memory in
	// megabytes.
	MemReservationMB pulumi.Float64Output `pulumi:"memReservationMB"`
	// Gets or sets CPUSharesLevel which specifies the memory allocation level for this pool.
	// This property is used in relative allocation between resource consumers.
	MemSharesLevel pulumi.StringOutput `pulumi:"memSharesLevel"`
	// Gets or sets the vCenter Managed Object name for the resource pool.
	MoName pulumi.StringOutput `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the resource pool.
	MoRefId pulumi.StringPtrOutput `pulumi:"moRefId"`
	// Gets or sets the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the network ARM ids.
	NetworkIds pulumi.StringArrayOutput `pulumi:"networkIds"`
	// Gets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource status information.
	Statuses ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Gets or sets the ARM Id of the vCenter resource in which this resource pool resides.
	VCenterId pulumi.StringPtrOutput `pulumi:"vCenterId"`
}

// NewResourcePool registers a new resource with the given unique name, arguments, and options.
func NewResourcePool(ctx *pulumi.Context,
	name string, args *ResourcePoolArgs, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:ResourcePool"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:ResourcePool"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:ResourcePool"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:ResourcePool"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20230301preview:ResourcePool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ResourcePool
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20231001:ResourcePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePool gets an existing ResourcePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePoolState, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	var resource ResourcePool
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20231001:ResourcePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePool resources.
type resourcePoolState struct {
}

type ResourcePoolState struct {
}

func (ResourcePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolState)(nil)).Elem()
}

type resourcePoolArgs struct {
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the resource pool.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the resource pool.
	MoRefId *string `pulumi:"moRefId"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the resourcePool.
	ResourcePoolName *string `pulumi:"resourcePoolName"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the ARM Id of the vCenter resource in which this resource pool resides.
	VCenterId *string `pulumi:"vCenterId"`
}

// The set of arguments for constructing a ResourcePool resource.
type ResourcePoolArgs struct {
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// Gets or sets the inventory Item ID for the resource pool.
	InventoryItemId pulumi.StringPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the resource pool.
	MoRefId pulumi.StringPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Name of the resourcePool.
	ResourcePoolName pulumi.StringPtrInput
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapInput
	// Gets or sets the ARM Id of the vCenter resource in which this resource pool resides.
	VCenterId pulumi.StringPtrInput
}

func (ResourcePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolArgs)(nil)).Elem()
}

type ResourcePoolInput interface {
	pulumi.Input

	ToResourcePoolOutput() ResourcePoolOutput
	ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput
}

func (*ResourcePool) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (i *ResourcePool) ToResourcePoolOutput() ResourcePoolOutput {
	return i.ToResourcePoolOutputWithContext(context.Background())
}

func (i *ResourcePool) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolOutput)
}

func (i *ResourcePool) ToOutput(ctx context.Context) pulumix.Output[*ResourcePool] {
	return pulumix.Output[*ResourcePool]{
		OutputState: i.ToResourcePoolOutputWithContext(ctx).OutputState,
	}
}

type ResourcePoolOutput struct{ *pulumi.OutputState }

func (ResourcePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (o ResourcePoolOutput) ToResourcePoolOutput() ResourcePoolOutput {
	return o
}

func (o ResourcePoolOutput) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return o
}

func (o ResourcePoolOutput) ToOutput(ctx context.Context) pulumix.Output[*ResourcePool] {
	return pulumix.Output[*ResourcePool]{
		OutputState: o.OutputState,
	}
}

// Gets the max CPU usage across all cores on the pool in MHz.
func (o ResourcePoolOutput) CpuCapacityMHz() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.CpuCapacityMHz }).(pulumi.Float64Output)
}

// Gets or sets CPULimitMHz which specifies a CPU usage limit in MHz.
// Utilization will not exceed this limit even if there are available resources.
func (o ResourcePoolOutput) CpuLimitMHz() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.CpuLimitMHz }).(pulumi.Float64Output)
}

// Gets the used CPU usage across all cores on the pool in MHz.
func (o ResourcePoolOutput) CpuOverallUsageMHz() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.CpuOverallUsageMHz }).(pulumi.Float64Output)
}

// Gets or sets CPUReservationMHz which specifies the CPU size in MHz that is guaranteed
// to be available.
func (o ResourcePoolOutput) CpuReservationMHz() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.CpuReservationMHz }).(pulumi.Float64Output)
}

// Gets or sets CPUSharesLevel which specifies the CPU allocation level for this pool.
// This property is used in relative allocation between resource consumers.
func (o ResourcePoolOutput) CpuSharesLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.CpuSharesLevel }).(pulumi.StringOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o ResourcePoolOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets the datastore ARM ids.
func (o ResourcePoolOutput) DatastoreIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringArrayOutput { return v.DatastoreIds }).(pulumi.StringArrayOutput)
}

// Gets or sets the extended location.
func (o ResourcePoolOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *ResourcePool) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the inventory Item ID for the resource pool.
func (o ResourcePoolOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o ResourcePoolOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o ResourcePoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Gets the total amount of physical memory on the pool in GB.
func (o ResourcePoolOutput) MemCapacityGB() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.MemCapacityGB }).(pulumi.Float64Output)
}

// Gets or sets MemLimitMB specifies a memory usage limit in megabytes.
// Utilization will not exceed the specified limit even if there are available resources.
func (o ResourcePoolOutput) MemLimitMB() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.MemLimitMB }).(pulumi.Float64Output)
}

// Gets the used physical memory on the pool in GB.
func (o ResourcePoolOutput) MemOverallUsageGB() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.MemOverallUsageGB }).(pulumi.Float64Output)
}

// Gets or sets MemReservationMB which specifies the guaranteed available memory in
// megabytes.
func (o ResourcePoolOutput) MemReservationMB() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.MemReservationMB }).(pulumi.Float64Output)
}

// Gets or sets CPUSharesLevel which specifies the memory allocation level for this pool.
// This property is used in relative allocation between resource consumers.
func (o ResourcePoolOutput) MemSharesLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.MemSharesLevel }).(pulumi.StringOutput)
}

// Gets or sets the vCenter Managed Object name for the resource pool.
func (o ResourcePoolOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the resource pool.
func (o ResourcePoolOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o ResourcePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the network ARM ids.
func (o ResourcePoolOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringArrayOutput { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

// Gets the provisioning state.
func (o ResourcePoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o ResourcePoolOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *ResourcePool) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o ResourcePoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ResourcePool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o ResourcePoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o ResourcePoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o ResourcePoolOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this resource pool resides.
func (o ResourcePoolOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourcePoolOutput{})
}

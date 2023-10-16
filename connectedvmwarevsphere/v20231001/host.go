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

// Define the host.
type Host struct {
	pulumi.CustomResourceState

	// Gets the max CPU usage across all cores in MHz.
	CpuMhz pulumi.Float64Output `pulumi:"cpuMhz"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName pulumi.StringOutput `pulumi:"customResourceName"`
	// Gets the datastore ARM ids.
	DatastoreIds pulumi.StringArrayOutput `pulumi:"datastoreIds"`
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the host.
	InventoryItemId pulumi.StringPtrOutput `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Gets the total amount of physical memory on the host in GB.
	MemorySizeGB pulumi.Float64Output `pulumi:"memorySizeGB"`
	// Gets or sets the vCenter Managed Object name for the host.
	MoName pulumi.StringOutput `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the host.
	MoRefId pulumi.StringPtrOutput `pulumi:"moRefId"`
	// Gets or sets the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the network ARM ids.
	NetworkIds pulumi.StringArrayOutput `pulumi:"networkIds"`
	// Gets the used CPU usage across all cores in MHz.
	OverallCpuUsageMHz pulumi.Float64Output `pulumi:"overallCpuUsageMHz"`
	// Gets the used physical memory on the host in GB.
	OverallMemoryUsageGB pulumi.Float64Output `pulumi:"overallMemoryUsageGB"`
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
	// Gets or sets the ARM Id of the vCenter resource in which this host resides.
	VCenterId pulumi.StringPtrOutput `pulumi:"vCenterId"`
}

// NewHost registers a new resource with the given unique name, arguments, and options.
func NewHost(ctx *pulumi.Context,
	name string, args *HostArgs, opts ...pulumi.ResourceOption) (*Host, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:Host"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:Host"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:Host"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:Host"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20230301preview:Host"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Host
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20231001:Host", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHost gets an existing Host resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostState, opts ...pulumi.ResourceOption) (*Host, error) {
	var resource Host
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20231001:Host", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Host resources.
type hostState struct {
}

type HostState struct {
}

func (HostState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostState)(nil)).Elem()
}

type hostArgs struct {
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Name of the host.
	HostName *string `pulumi:"hostName"`
	// Gets or sets the inventory Item ID for the host.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the host.
	MoRefId *string `pulumi:"moRefId"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the ARM Id of the vCenter resource in which this host resides.
	VCenterId *string `pulumi:"vCenterId"`
}

// The set of arguments for constructing a Host resource.
type HostArgs struct {
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// Name of the host.
	HostName pulumi.StringPtrInput
	// Gets or sets the inventory Item ID for the host.
	InventoryItemId pulumi.StringPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the host.
	MoRefId pulumi.StringPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapInput
	// Gets or sets the ARM Id of the vCenter resource in which this host resides.
	VCenterId pulumi.StringPtrInput
}

func (HostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostArgs)(nil)).Elem()
}

type HostInput interface {
	pulumi.Input

	ToHostOutput() HostOutput
	ToHostOutputWithContext(ctx context.Context) HostOutput
}

func (*Host) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (i *Host) ToHostOutput() HostOutput {
	return i.ToHostOutputWithContext(context.Background())
}

func (i *Host) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostOutput)
}

func (i *Host) ToOutput(ctx context.Context) pulumix.Output[*Host] {
	return pulumix.Output[*Host]{
		OutputState: i.ToHostOutputWithContext(ctx).OutputState,
	}
}

type HostOutput struct{ *pulumi.OutputState }

func (HostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (o HostOutput) ToHostOutput() HostOutput {
	return o
}

func (o HostOutput) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return o
}

func (o HostOutput) ToOutput(ctx context.Context) pulumix.Output[*Host] {
	return pulumix.Output[*Host]{
		OutputState: o.OutputState,
	}
}

// Gets the max CPU usage across all cores in MHz.
func (o HostOutput) CpuMhz() pulumi.Float64Output {
	return o.ApplyT(func(v *Host) pulumi.Float64Output { return v.CpuMhz }).(pulumi.Float64Output)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o HostOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets the datastore ARM ids.
func (o HostOutput) DatastoreIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Host) pulumi.StringArrayOutput { return v.DatastoreIds }).(pulumi.StringArrayOutput)
}

// Gets or sets the extended location.
func (o HostOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Host) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the inventory Item ID for the host.
func (o HostOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o HostOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o HostOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Gets the total amount of physical memory on the host in GB.
func (o HostOutput) MemorySizeGB() pulumi.Float64Output {
	return o.ApplyT(func(v *Host) pulumi.Float64Output { return v.MemorySizeGB }).(pulumi.Float64Output)
}

// Gets or sets the vCenter Managed Object name for the host.
func (o HostOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the host.
func (o HostOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o HostOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the network ARM ids.
func (o HostOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Host) pulumi.StringArrayOutput { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

// Gets the used CPU usage across all cores in MHz.
func (o HostOutput) OverallCpuUsageMHz() pulumi.Float64Output {
	return o.ApplyT(func(v *Host) pulumi.Float64Output { return v.OverallCpuUsageMHz }).(pulumi.Float64Output)
}

// Gets the used physical memory on the host in GB.
func (o HostOutput) OverallMemoryUsageGB() pulumi.Float64Output {
	return o.ApplyT(func(v *Host) pulumi.Float64Output { return v.OverallMemoryUsageGB }).(pulumi.Float64Output)
}

// Gets the provisioning state.
func (o HostOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o HostOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *Host) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o HostOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Host) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o HostOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Host) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o HostOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o HostOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this host resides.
func (o HostOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HostOutput{})
}
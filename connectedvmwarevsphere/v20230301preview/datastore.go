// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the datastore.
type Datastore struct {
	pulumi.CustomResourceState

	// Gets or sets Maximum capacity of this datastore in GBs.
	CapacityGB pulumi.Float64Output `pulumi:"capacityGB"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName pulumi.StringOutput `pulumi:"customResourceName"`
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Gets or sets Available space of this datastore in GBs.
	FreeSpaceGB pulumi.Float64Output `pulumi:"freeSpaceGB"`
	// Gets or sets the inventory Item ID for the datastore.
	InventoryItemId pulumi.StringPtrOutput `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Gets or sets the vCenter Managed Object name for the datastore.
	MoName pulumi.StringOutput `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the datastore.
	MoRefId pulumi.StringPtrOutput `pulumi:"moRefId"`
	// Gets or sets the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
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
	// Gets or sets the ARM Id of the vCenter resource in which this datastore resides.
	VCenterId pulumi.StringPtrOutput `pulumi:"vCenterId"`
}

// NewDatastore registers a new resource with the given unique name, arguments, and options.
func NewDatastore(ctx *pulumi.Context,
	name string, args *DatastoreArgs, opts ...pulumi.ResourceOption) (*Datastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:Datastore"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20231001:Datastore"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Datastore
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20230301preview:Datastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatastore gets an existing Datastore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastoreState, opts ...pulumi.ResourceOption) (*Datastore, error) {
	var resource Datastore
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20230301preview:Datastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datastore resources.
type datastoreState struct {
}

type DatastoreState struct {
}

func (DatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreState)(nil)).Elem()
}

type datastoreArgs struct {
	// Name of the datastore.
	DatastoreName *string `pulumi:"datastoreName"`
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the inventory Item ID for the datastore.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the datastore.
	MoRefId *string `pulumi:"moRefId"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the ARM Id of the vCenter resource in which this datastore resides.
	VCenterId *string `pulumi:"vCenterId"`
}

// The set of arguments for constructing a Datastore resource.
type DatastoreArgs struct {
	// Name of the datastore.
	DatastoreName pulumi.StringPtrInput
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// Gets or sets the inventory Item ID for the datastore.
	InventoryItemId pulumi.StringPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the datastore.
	MoRefId pulumi.StringPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapInput
	// Gets or sets the ARM Id of the vCenter resource in which this datastore resides.
	VCenterId pulumi.StringPtrInput
}

func (DatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreArgs)(nil)).Elem()
}

type DatastoreInput interface {
	pulumi.Input

	ToDatastoreOutput() DatastoreOutput
	ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput
}

func (*Datastore) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (i *Datastore) ToDatastoreOutput() DatastoreOutput {
	return i.ToDatastoreOutputWithContext(context.Background())
}

func (i *Datastore) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreOutput)
}

type DatastoreOutput struct{ *pulumi.OutputState }

func (DatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datastore)(nil)).Elem()
}

func (o DatastoreOutput) ToDatastoreOutput() DatastoreOutput {
	return o
}

func (o DatastoreOutput) ToDatastoreOutputWithContext(ctx context.Context) DatastoreOutput {
	return o
}

// Gets or sets Maximum capacity of this datastore in GBs.
func (o DatastoreOutput) CapacityGB() pulumi.Float64Output {
	return o.ApplyT(func(v *Datastore) pulumi.Float64Output { return v.CapacityGB }).(pulumi.Float64Output)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o DatastoreOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets or sets the extended location.
func (o DatastoreOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Datastore) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets Available space of this datastore in GBs.
func (o DatastoreOutput) FreeSpaceGB() pulumi.Float64Output {
	return o.ApplyT(func(v *Datastore) pulumi.Float64Output { return v.FreeSpaceGB }).(pulumi.Float64Output)
}

// Gets or sets the inventory Item ID for the datastore.
func (o DatastoreOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o DatastoreOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o DatastoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets the vCenter Managed Object name for the datastore.
func (o DatastoreOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the datastore.
func (o DatastoreOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o DatastoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o DatastoreOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o DatastoreOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *Datastore) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o DatastoreOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Datastore) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o DatastoreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o DatastoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o DatastoreOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this datastore resides.
func (o DatastoreOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datastore) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatastoreOutput{})
}

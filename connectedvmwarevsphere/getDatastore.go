// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements datastore GET method.
// Azure REST API version: 2022-07-15-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-10-01.
func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatastoreArgs struct {
	// Name of the datastore.
	DatastoreName string `pulumi:"datastoreName"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Define the datastore.
type LookupDatastoreResult struct {
	// Gets or sets Maximum capacity of this datastore in GBs.
	CapacityGB float64 `pulumi:"capacityGB"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName string `pulumi:"customResourceName"`
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Gets or sets Available space of this datastore in GBs.
	FreeSpaceGB float64 `pulumi:"freeSpaceGB"`
	// Gets or sets the Id.
	Id string `pulumi:"id"`
	// Gets or sets the inventory Item ID for the datastore.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Gets or sets the vCenter Managed Object name for the datastore.
	MoName string `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the datastore.
	MoRefId *string `pulumi:"moRefId"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
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
	// Gets or sets the ARM Id of the vCenter resource in which this datastore resides.
	VCenterId *string `pulumi:"vCenterId"`
}

func LookupDatastoreOutput(ctx *pulumi.Context, args LookupDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatastoreResult, error) {
			args := v.(LookupDatastoreArgs)
			r, err := LookupDatastore(ctx, &args, opts...)
			var s LookupDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatastoreResultOutput)
}

type LookupDatastoreOutputArgs struct {
	// Name of the datastore.
	DatastoreName pulumi.StringInput `pulumi:"datastoreName"`
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreArgs)(nil)).Elem()
}

// Define the datastore.
type LookupDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreResult)(nil)).Elem()
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutput() LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutputWithContext(ctx context.Context) LookupDatastoreResultOutput {
	return o
}

// Gets or sets Maximum capacity of this datastore in GBs.
func (o LookupDatastoreResultOutput) CapacityGB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDatastoreResult) float64 { return v.CapacityGB }).(pulumi.Float64Output)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o LookupDatastoreResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets or sets the extended location.
func (o LookupDatastoreResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets Available space of this datastore in GBs.
func (o LookupDatastoreResultOutput) FreeSpaceGB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDatastoreResult) float64 { return v.FreeSpaceGB }).(pulumi.Float64Output)
}

// Gets or sets the Id.
func (o LookupDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the inventory Item ID for the datastore.
func (o LookupDatastoreResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupDatastoreResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o LookupDatastoreResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets the vCenter Managed Object name for the datastore.
func (o LookupDatastoreResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the datastore.
func (o LookupDatastoreResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o LookupDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupDatastoreResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o LookupDatastoreResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupDatastoreResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o LookupDatastoreResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDatastoreResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o LookupDatastoreResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatastoreResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o LookupDatastoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o LookupDatastoreResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this datastore resides.
func (o LookupDatastoreResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatastoreResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220715preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements host GET method.
func LookupHost(ctx *pulumi.Context, args *LookupHostArgs, opts ...pulumi.InvokeOption) (*LookupHostResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHostResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20220715preview:getHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostArgs struct {
	// Name of the host.
	HostName string `pulumi:"hostName"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Define the host.
type LookupHostResult struct {
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName string `pulumi:"customResourceName"`
	// Gets or sets the datastore ARM ids.
	DatastoreIds []string `pulumi:"datastoreIds"`
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Gets or sets the Id.
	Id string `pulumi:"id"`
	// Gets or sets the inventory Item ID for the host.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Gets or sets the vCenter Managed Object name for the host.
	MoName string `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the host.
	MoRefId *string `pulumi:"moRefId"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Gets or sets the network ARM ids.
	NetworkIds []string `pulumi:"networkIds"`
	// Gets or sets the provisioning state.
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
	// Gets or sets the ARM Id of the vCenter resource in which this host resides.
	VCenterId *string `pulumi:"vCenterId"`
}

func LookupHostOutput(ctx *pulumi.Context, args LookupHostOutputArgs, opts ...pulumi.InvokeOption) LookupHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostResult, error) {
			args := v.(LookupHostArgs)
			r, err := LookupHost(ctx, &args, opts...)
			var s LookupHostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostResultOutput)
}

type LookupHostOutputArgs struct {
	// Name of the host.
	HostName pulumi.StringInput `pulumi:"hostName"`
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostArgs)(nil)).Elem()
}

// Define the host.
type LookupHostResultOutput struct{ *pulumi.OutputState }

func (LookupHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostResult)(nil)).Elem()
}

func (o LookupHostResultOutput) ToLookupHostResultOutput() LookupHostResultOutput {
	return o
}

func (o LookupHostResultOutput) ToLookupHostResultOutputWithContext(ctx context.Context) LookupHostResultOutput {
	return o
}

// Gets the name of the corresponding resource in Kubernetes.
func (o LookupHostResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets or sets the datastore ARM ids.
func (o LookupHostResultOutput) DatastoreIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostResult) []string { return v.DatastoreIds }).(pulumi.StringArrayOutput)
}

// Gets or sets the extended location.
func (o LookupHostResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupHostResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the Id.
func (o LookupHostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the inventory Item ID for the host.
func (o LookupHostResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupHostResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o LookupHostResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets the vCenter Managed Object name for the host.
func (o LookupHostResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the host.
func (o LookupHostResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o LookupHostResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the network ARM ids.
func (o LookupHostResultOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostResult) []string { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

// Gets or sets the provisioning state.
func (o LookupHostResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o LookupHostResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupHostResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o LookupHostResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHostResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o LookupHostResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHostResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o LookupHostResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o LookupHostResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this host resides.
func (o LookupHostResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostResultOutput{})
}

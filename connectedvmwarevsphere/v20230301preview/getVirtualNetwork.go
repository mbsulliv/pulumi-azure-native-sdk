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

// Implements virtual network GET method.
func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20230301preview:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkArgs struct {
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the virtual network resource.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// Define the virtualNetwork.
type LookupVirtualNetworkResult struct {
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName string `pulumi:"customResourceName"`
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Gets or sets the Id.
	Id string `pulumi:"id"`
	// Gets or sets the inventory Item ID for the virtual network.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Gets or sets the vCenter Managed Object name for the virtual network.
	MoName string `pulumi:"moName"`
	// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual network.
	MoRefId *string `pulumi:"moRefId"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
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
	// Gets or sets the ARM Id of the vCenter resource in which this template resides.
	VCenterId *string `pulumi:"vCenterId"`
}

func LookupVirtualNetworkOutput(ctx *pulumi.Context, args LookupVirtualNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkResult, error) {
			args := v.(LookupVirtualNetworkArgs)
			r, err := LookupVirtualNetwork(ctx, &args, opts...)
			var s LookupVirtualNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkResultOutput)
}

type LookupVirtualNetworkOutputArgs struct {
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the virtual network resource.
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (LookupVirtualNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkArgs)(nil)).Elem()
}

// Define the virtualNetwork.
type LookupVirtualNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkResult)(nil)).Elem()
}

func (o LookupVirtualNetworkResultOutput) ToLookupVirtualNetworkResultOutput() LookupVirtualNetworkResultOutput {
	return o
}

func (o LookupVirtualNetworkResultOutput) ToLookupVirtualNetworkResultOutputWithContext(ctx context.Context) LookupVirtualNetworkResultOutput {
	return o
}

func (o LookupVirtualNetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualNetworkResult] {
	return pulumix.Output[LookupVirtualNetworkResult]{
		OutputState: o.OutputState,
	}
}

// Gets the name of the corresponding resource in Kubernetes.
func (o LookupVirtualNetworkResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets or sets the extended location.
func (o LookupVirtualNetworkResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the Id.
func (o LookupVirtualNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the inventory Item ID for the virtual network.
func (o LookupVirtualNetworkResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupVirtualNetworkResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o LookupVirtualNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets the vCenter Managed Object name for the virtual network.
func (o LookupVirtualNetworkResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.MoName }).(pulumi.StringOutput)
}

// Gets or sets the vCenter MoRef (Managed Object Reference) ID for the virtual network.
func (o LookupVirtualNetworkResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o LookupVirtualNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state.
func (o LookupVirtualNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o LookupVirtualNetworkResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o LookupVirtualNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o LookupVirtualNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o LookupVirtualNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o LookupVirtualNetworkResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the ARM Id of the vCenter resource in which this template resides.
func (o LookupVirtualNetworkResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkResultOutput{})
}

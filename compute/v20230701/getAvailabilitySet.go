// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about an availability set.
func LookupAvailabilitySet(ctx *pulumi.Context, args *LookupAvailabilitySetArgs, opts ...pulumi.InvokeOption) (*LookupAvailabilitySetResult, error) {
	var rv LookupAvailabilitySetResult
	err := ctx.Invoke("azure-native:compute/v20230701:getAvailabilitySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAvailabilitySetArgs struct {
	// The name of the availability set.
	AvailabilitySetName string `pulumi:"availabilitySetName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to an availability set at creation time. An existing VM cannot be added to an availability set.
type LookupAvailabilitySetResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Fault Domain count.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Update Domain count.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// Specifies information about the proximity placement group that the availability set should be assigned to. Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResourceResponse `pulumi:"proximityPlacementGroup"`
	// Sku of the availability set, only name is required to be set. See AvailabilitySetSkuTypes for possible set of values. Use 'Aligned' for virtual machines with managed disks and 'Classic' for virtual machines with unmanaged disks. Default value is 'Classic'.
	Sku *SkuResponse `pulumi:"sku"`
	// The resource status information.
	Statuses []InstanceViewStatusResponse `pulumi:"statuses"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// A list of references to all virtual machines in the availability set.
	VirtualMachines []SubResourceResponse `pulumi:"virtualMachines"`
}

func LookupAvailabilitySetOutput(ctx *pulumi.Context, args LookupAvailabilitySetOutputArgs, opts ...pulumi.InvokeOption) LookupAvailabilitySetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAvailabilitySetResult, error) {
			args := v.(LookupAvailabilitySetArgs)
			r, err := LookupAvailabilitySet(ctx, &args, opts...)
			var s LookupAvailabilitySetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAvailabilitySetResultOutput)
}

type LookupAvailabilitySetOutputArgs struct {
	// The name of the availability set.
	AvailabilitySetName pulumi.StringInput `pulumi:"availabilitySetName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAvailabilitySetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailabilitySetArgs)(nil)).Elem()
}

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to an availability set at creation time. An existing VM cannot be added to an availability set.
type LookupAvailabilitySetResultOutput struct{ *pulumi.OutputState }

func (LookupAvailabilitySetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailabilitySetResult)(nil)).Elem()
}

func (o LookupAvailabilitySetResultOutput) ToLookupAvailabilitySetResultOutput() LookupAvailabilitySetResultOutput {
	return o
}

func (o LookupAvailabilitySetResultOutput) ToLookupAvailabilitySetResultOutputWithContext(ctx context.Context) LookupAvailabilitySetResultOutput {
	return o
}

// Resource Id
func (o LookupAvailabilitySetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupAvailabilitySetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupAvailabilitySetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Fault Domain count.
func (o LookupAvailabilitySetResultOutput) PlatformFaultDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *int { return v.PlatformFaultDomainCount }).(pulumi.IntPtrOutput)
}

// Update Domain count.
func (o LookupAvailabilitySetResultOutput) PlatformUpdateDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *int { return v.PlatformUpdateDomainCount }).(pulumi.IntPtrOutput)
}

// Specifies information about the proximity placement group that the availability set should be assigned to. Minimum api-version: 2018-04-01.
func (o LookupAvailabilitySetResultOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *SubResourceResponse { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

// Sku of the availability set, only name is required to be set. See AvailabilitySetSkuTypes for possible set of values. Use 'Aligned' for virtual machines with managed disks and 'Classic' for virtual machines with unmanaged disks. Default value is 'Classic'.
func (o LookupAvailabilitySetResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The resource status information.
func (o LookupAvailabilitySetResultOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

// Resource tags
func (o LookupAvailabilitySetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupAvailabilitySetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of references to all virtual machines in the availability set.
func (o LookupAvailabilitySetResultOutput) VirtualMachines() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupAvailabilitySetResult) []SubResourceResponse { return v.VirtualMachines }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAvailabilitySetResultOutput{})
}

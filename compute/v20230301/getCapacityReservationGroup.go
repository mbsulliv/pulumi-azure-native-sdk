// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The operation that retrieves information about a capacity reservation group.
func LookupCapacityReservationGroup(ctx *pulumi.Context, args *LookupCapacityReservationGroupArgs, opts ...pulumi.InvokeOption) (*LookupCapacityReservationGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCapacityReservationGroupResult
	err := ctx.Invoke("azure-native:compute/v20230301:getCapacityReservationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityReservationGroupArgs struct {
	// The name of the capacity reservation group.
	CapacityReservationGroupName string `pulumi:"capacityReservationGroupName"`
	// The expand expression to apply on the operation. 'InstanceView' will retrieve the list of instance views of the capacity reservations under the capacity reservation group which is a snapshot of the runtime properties of a capacity reservation that is managed by the platform and can change outside of control plane operations.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the capacity reservation group that the capacity reservations should be assigned to. Currently, a capacity reservation can only be added to a capacity reservation group at creation time. An existing capacity reservation cannot be added or moved to another capacity reservation group.
type LookupCapacityReservationGroupResult struct {
	// A list of all capacity reservation resource ids that belong to capacity reservation group.
	CapacityReservations []SubResourceReadOnlyResponse `pulumi:"capacityReservations"`
	// Resource Id
	Id string `pulumi:"id"`
	// The capacity reservation group instance view which has the list of instance views for all the capacity reservations that belong to the capacity reservation group.
	InstanceView CapacityReservationGroupInstanceViewResponse `pulumi:"instanceView"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// A list of references to all virtual machines associated to the capacity reservation group.
	VirtualMachinesAssociated []SubResourceReadOnlyResponse `pulumi:"virtualMachinesAssociated"`
	// Availability Zones to use for this capacity reservation group. The zones can be assigned only during creation. If not provided, the group supports only regional resources in the region. If provided, enforces each capacity reservation in the group to be in one of the zones.
	Zones []string `pulumi:"zones"`
}

func LookupCapacityReservationGroupOutput(ctx *pulumi.Context, args LookupCapacityReservationGroupOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityReservationGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityReservationGroupResult, error) {
			args := v.(LookupCapacityReservationGroupArgs)
			r, err := LookupCapacityReservationGroup(ctx, &args, opts...)
			var s LookupCapacityReservationGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityReservationGroupResultOutput)
}

type LookupCapacityReservationGroupOutputArgs struct {
	// The name of the capacity reservation group.
	CapacityReservationGroupName pulumi.StringInput `pulumi:"capacityReservationGroupName"`
	// The expand expression to apply on the operation. 'InstanceView' will retrieve the list of instance views of the capacity reservations under the capacity reservation group which is a snapshot of the runtime properties of a capacity reservation that is managed by the platform and can change outside of control plane operations.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCapacityReservationGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationGroupArgs)(nil)).Elem()
}

// Specifies information about the capacity reservation group that the capacity reservations should be assigned to. Currently, a capacity reservation can only be added to a capacity reservation group at creation time. An existing capacity reservation cannot be added or moved to another capacity reservation group.
type LookupCapacityReservationGroupResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityReservationGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationGroupResult)(nil)).Elem()
}

func (o LookupCapacityReservationGroupResultOutput) ToLookupCapacityReservationGroupResultOutput() LookupCapacityReservationGroupResultOutput {
	return o
}

func (o LookupCapacityReservationGroupResultOutput) ToLookupCapacityReservationGroupResultOutputWithContext(ctx context.Context) LookupCapacityReservationGroupResultOutput {
	return o
}

func (o LookupCapacityReservationGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCapacityReservationGroupResult] {
	return pulumix.Output[LookupCapacityReservationGroupResult]{
		OutputState: o.OutputState,
	}
}

// A list of all capacity reservation resource ids that belong to capacity reservation group.
func (o LookupCapacityReservationGroupResultOutput) CapacityReservations() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) []SubResourceReadOnlyResponse {
		return v.CapacityReservations
	}).(SubResourceReadOnlyResponseArrayOutput)
}

// Resource Id
func (o LookupCapacityReservationGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The capacity reservation group instance view which has the list of instance views for all the capacity reservations that belong to the capacity reservation group.
func (o LookupCapacityReservationGroupResultOutput) InstanceView() CapacityReservationGroupInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) CapacityReservationGroupInstanceViewResponse {
		return v.InstanceView
	}).(CapacityReservationGroupInstanceViewResponseOutput)
}

// Resource location
func (o LookupCapacityReservationGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupCapacityReservationGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource tags
func (o LookupCapacityReservationGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupCapacityReservationGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of references to all virtual machines associated to the capacity reservation group.
func (o LookupCapacityReservationGroupResultOutput) VirtualMachinesAssociated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) []SubResourceReadOnlyResponse {
		return v.VirtualMachinesAssociated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

// Availability Zones to use for this capacity reservation group. The zones can be assigned only during creation. If not provided, the group supports only regional resources in the region. If provided, enforces each capacity reservation in the group to be in one of the zones.
func (o LookupCapacityReservationGroupResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityReservationGroupResultOutput{})
}

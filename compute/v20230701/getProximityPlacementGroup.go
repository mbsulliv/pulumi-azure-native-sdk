// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a proximity placement group .
func LookupProximityPlacementGroup(ctx *pulumi.Context, args *LookupProximityPlacementGroupArgs, opts ...pulumi.InvokeOption) (*LookupProximityPlacementGroupResult, error) {
	var rv LookupProximityPlacementGroupResult
	err := ctx.Invoke("azure-native:compute/v20230701:getProximityPlacementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProximityPlacementGroupArgs struct {
	// includeColocationStatus=true enables fetching the colocation status of all the resources in the proximity placement group.
	IncludeColocationStatus *string `pulumi:"includeColocationStatus"`
	// The name of the proximity placement group.
	ProximityPlacementGroupName string `pulumi:"proximityPlacementGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the proximity placement group.
type LookupProximityPlacementGroupResult struct {
	// A list of references to all availability sets in the proximity placement group.
	AvailabilitySets []SubResourceWithColocationStatusResponse `pulumi:"availabilitySets"`
	// Describes colocation status of the Proximity Placement Group.
	ColocationStatus *InstanceViewStatusResponse `pulumi:"colocationStatus"`
	// Resource Id
	Id string `pulumi:"id"`
	// Specifies the user intent of the proximity placement group.
	Intent *ProximityPlacementGroupPropertiesResponseIntent `pulumi:"intent"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Specifies the type of the proximity placement group. Possible values are: **Standard** : Co-locate resources within an Azure region or Availability Zone. **Ultra** : For future use.
	ProximityPlacementGroupType *string `pulumi:"proximityPlacementGroupType"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// A list of references to all virtual machine scale sets in the proximity placement group.
	VirtualMachineScaleSets []SubResourceWithColocationStatusResponse `pulumi:"virtualMachineScaleSets"`
	// A list of references to all virtual machines in the proximity placement group.
	VirtualMachines []SubResourceWithColocationStatusResponse `pulumi:"virtualMachines"`
	// Specifies the Availability Zone where virtual machine, virtual machine scale set or availability set associated with the  proximity placement group can be created.
	Zones []string `pulumi:"zones"`
}

func LookupProximityPlacementGroupOutput(ctx *pulumi.Context, args LookupProximityPlacementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupProximityPlacementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProximityPlacementGroupResult, error) {
			args := v.(LookupProximityPlacementGroupArgs)
			r, err := LookupProximityPlacementGroup(ctx, &args, opts...)
			var s LookupProximityPlacementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProximityPlacementGroupResultOutput)
}

type LookupProximityPlacementGroupOutputArgs struct {
	// includeColocationStatus=true enables fetching the colocation status of all the resources in the proximity placement group.
	IncludeColocationStatus pulumi.StringPtrInput `pulumi:"includeColocationStatus"`
	// The name of the proximity placement group.
	ProximityPlacementGroupName pulumi.StringInput `pulumi:"proximityPlacementGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProximityPlacementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProximityPlacementGroupArgs)(nil)).Elem()
}

// Specifies information about the proximity placement group.
type LookupProximityPlacementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupProximityPlacementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProximityPlacementGroupResult)(nil)).Elem()
}

func (o LookupProximityPlacementGroupResultOutput) ToLookupProximityPlacementGroupResultOutput() LookupProximityPlacementGroupResultOutput {
	return o
}

func (o LookupProximityPlacementGroupResultOutput) ToLookupProximityPlacementGroupResultOutputWithContext(ctx context.Context) LookupProximityPlacementGroupResultOutput {
	return o
}

// A list of references to all availability sets in the proximity placement group.
func (o LookupProximityPlacementGroupResultOutput) AvailabilitySets() SubResourceWithColocationStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) []SubResourceWithColocationStatusResponse {
		return v.AvailabilitySets
	}).(SubResourceWithColocationStatusResponseArrayOutput)
}

// Describes colocation status of the Proximity Placement Group.
func (o LookupProximityPlacementGroupResultOutput) ColocationStatus() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) *InstanceViewStatusResponse { return v.ColocationStatus }).(InstanceViewStatusResponsePtrOutput)
}

// Resource Id
func (o LookupProximityPlacementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the user intent of the proximity placement group.
func (o LookupProximityPlacementGroupResultOutput) Intent() ProximityPlacementGroupPropertiesResponseIntentPtrOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) *ProximityPlacementGroupPropertiesResponseIntent {
		return v.Intent
	}).(ProximityPlacementGroupPropertiesResponseIntentPtrOutput)
}

// Resource location
func (o LookupProximityPlacementGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupProximityPlacementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the type of the proximity placement group. Possible values are: **Standard** : Co-locate resources within an Azure region or Availability Zone. **Ultra** : For future use.
func (o LookupProximityPlacementGroupResultOutput) ProximityPlacementGroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) *string { return v.ProximityPlacementGroupType }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o LookupProximityPlacementGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupProximityPlacementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of references to all virtual machine scale sets in the proximity placement group.
func (o LookupProximityPlacementGroupResultOutput) VirtualMachineScaleSets() SubResourceWithColocationStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) []SubResourceWithColocationStatusResponse {
		return v.VirtualMachineScaleSets
	}).(SubResourceWithColocationStatusResponseArrayOutput)
}

// A list of references to all virtual machines in the proximity placement group.
func (o LookupProximityPlacementGroupResultOutput) VirtualMachines() SubResourceWithColocationStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) []SubResourceWithColocationStatusResponse {
		return v.VirtualMachines
	}).(SubResourceWithColocationStatusResponseArrayOutput)
}

// Specifies the Availability Zone where virtual machine, virtual machine scale set or availability set associated with the  proximity placement group can be created.
func (o LookupProximityPlacementGroupResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProximityPlacementGroupResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProximityPlacementGroupResultOutput{})
}

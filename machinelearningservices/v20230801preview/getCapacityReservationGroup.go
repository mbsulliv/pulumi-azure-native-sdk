// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupCapacityReservationGroup(ctx *pulumi.Context, args *LookupCapacityReservationGroupArgs, opts ...pulumi.InvokeOption) (*LookupCapacityReservationGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCapacityReservationGroupResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230801preview:getCapacityReservationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityReservationGroupArgs struct {
	GroupId string `pulumi:"groupId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupCapacityReservationGroupResult struct {
	// [Required] Additional attributes of the entity.
	CapacityReservationGroupProperties CapacityReservationGroupResponse `pulumi:"capacityReservationGroupProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Sku details required for ARM contract for Autoscaling.
	Sku *SkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
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
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCapacityReservationGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationGroupArgs)(nil)).Elem()
}

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

// [Required] Additional attributes of the entity.
func (o LookupCapacityReservationGroupResultOutput) CapacityReservationGroupProperties() CapacityReservationGroupResponseOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) CapacityReservationGroupResponse {
		return v.CapacityReservationGroupProperties
	}).(CapacityReservationGroupResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCapacityReservationGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o LookupCapacityReservationGroupResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o LookupCapacityReservationGroupResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupCapacityReservationGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCapacityReservationGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Sku details required for ARM contract for Autoscaling.
func (o LookupCapacityReservationGroupResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCapacityReservationGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupCapacityReservationGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCapacityReservationGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityReservationGroupResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230707preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a scaling plan.
func LookupScalingPlan(ctx *pulumi.Context, args *LookupScalingPlanArgs, opts ...pulumi.InvokeOption) (*LookupScalingPlanResult, error) {
	var rv LookupScalingPlanResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20230707preview:getScalingPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupScalingPlanArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the scaling plan.
	ScalingPlanName string `pulumi:"scalingPlanName"`
}

// Represents a scaling plan definition.
type LookupScalingPlanResult struct {
	// Description of scaling plan.
	Description *string `pulumi:"description"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag string `pulumi:"etag"`
	// Exclusion tag for scaling plan.
	ExclusionTag *string `pulumi:"exclusionTag"`
	// User friendly name of scaling plan.
	FriendlyName *string `pulumi:"friendlyName"`
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences []ScalingHostPoolReferenceResponse `pulumi:"hostPoolReferences"`
	// HostPool type for desktop.
	HostPoolType *string `pulumi:"hostPoolType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id       string                                               `pulumi:"id"`
	Identity *ResourceModelWithAllowedPropertySetResponseIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The name of the resource
	Name string `pulumi:"name"`
	// ObjectId of scaling plan. (internal use)
	ObjectId string                                           `pulumi:"objectId"`
	Plan     *ResourceModelWithAllowedPropertySetResponsePlan `pulumi:"plan"`
	// List of ScalingPlanPooledSchedule definitions.
	Schedules []ScalingScheduleResponse                       `pulumi:"schedules"`
	Sku       *ResourceModelWithAllowedPropertySetResponseSku `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Timezone of the scaling plan.
	TimeZone string `pulumi:"timeZone"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupScalingPlanResult
func (val *LookupScalingPlanResult) Defaults() *LookupScalingPlanResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.HostPoolType == nil {
		hostPoolType_ := "Pooled"
		tmp.HostPoolType = &hostPoolType_
	}
	return &tmp
}

func LookupScalingPlanOutput(ctx *pulumi.Context, args LookupScalingPlanOutputArgs, opts ...pulumi.InvokeOption) LookupScalingPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScalingPlanResult, error) {
			args := v.(LookupScalingPlanArgs)
			r, err := LookupScalingPlan(ctx, &args, opts...)
			var s LookupScalingPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScalingPlanResultOutput)
}

type LookupScalingPlanOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the scaling plan.
	ScalingPlanName pulumi.StringInput `pulumi:"scalingPlanName"`
}

func (LookupScalingPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanArgs)(nil)).Elem()
}

// Represents a scaling plan definition.
type LookupScalingPlanResultOutput struct{ *pulumi.OutputState }

func (LookupScalingPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanResult)(nil)).Elem()
}

func (o LookupScalingPlanResultOutput) ToLookupScalingPlanResultOutput() LookupScalingPlanResultOutput {
	return o
}

func (o LookupScalingPlanResultOutput) ToLookupScalingPlanResultOutputWithContext(ctx context.Context) LookupScalingPlanResultOutput {
	return o
}

// Description of scaling plan.
func (o LookupScalingPlanResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o LookupScalingPlanResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Exclusion tag for scaling plan.
func (o LookupScalingPlanResultOutput) ExclusionTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.ExclusionTag }).(pulumi.StringPtrOutput)
}

// User friendly name of scaling plan.
func (o LookupScalingPlanResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// List of ScalingHostPoolReference definitions.
func (o LookupScalingPlanResultOutput) HostPoolReferences() ScalingHostPoolReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) []ScalingHostPoolReferenceResponse { return v.HostPoolReferences }).(ScalingHostPoolReferenceResponseArrayOutput)
}

// HostPool type for desktop.
func (o LookupScalingPlanResultOutput) HostPoolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.HostPoolType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupScalingPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScalingPlanResultOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *ResourceModelWithAllowedPropertySetResponseIdentity {
		return v.Identity
	}).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupScalingPlanResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupScalingPlanResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o LookupScalingPlanResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupScalingPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

// ObjectId of scaling plan. (internal use)
func (o LookupScalingPlanResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupScalingPlanResultOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *ResourceModelWithAllowedPropertySetResponsePlan { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

// List of ScalingPlanPooledSchedule definitions.
func (o LookupScalingPlanResultOutput) Schedules() ScalingScheduleResponseArrayOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) []ScalingScheduleResponse { return v.Schedules }).(ScalingScheduleResponseArrayOutput)
}

func (o LookupScalingPlanResultOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *ResourceModelWithAllowedPropertySetResponseSku { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupScalingPlanResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupScalingPlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Timezone of the scaling plan.
func (o LookupScalingPlanResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupScalingPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScalingPlanResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220909

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a scaling plan definition.
type ScalingPlan struct {
	pulumi.CustomResourceState

	// Description of scaling plan.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Exclusion tag for scaling plan.
	ExclusionTag pulumi.StringPtrOutput `pulumi:"exclusionTag"`
	// User friendly name of scaling plan.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences ScalingHostPoolReferenceResponseArrayOutput `pulumi:"hostPoolReferences"`
	// HostPool type for desktop.
	HostPoolType pulumi.StringPtrOutput                                       `pulumi:"hostPoolType"`
	Identity     ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ObjectId of scaling plan. (internal use)
	ObjectId pulumi.StringOutput                                      `pulumi:"objectId"`
	Plan     ResourceModelWithAllowedPropertySetResponsePlanPtrOutput `pulumi:"plan"`
	// List of ScalingPlanPooledSchedule definitions.
	Schedules ScalingScheduleResponseArrayOutput                      `pulumi:"schedules"`
	Sku       ResourceModelWithAllowedPropertySetResponseSkuPtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Timezone of the scaling plan.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScalingPlan registers a new resource with the given unique name, arguments, and options.
func NewScalingPlan(ctx *pulumi.Context,
	name string, args *ScalingPlanArgs, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	if args.HostPoolType == nil {
		args.HostPoolType = pulumi.StringPtr("Pooled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230707preview:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230905:ScalingPlan"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20231004preview:ScalingPlan"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ScalingPlan
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20220909:ScalingPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingPlan gets an existing ScalingPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPlanState, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	var resource ScalingPlan
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20220909:ScalingPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingPlan resources.
type scalingPlanState struct {
}

type ScalingPlanState struct {
}

func (ScalingPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanState)(nil)).Elem()
}

type scalingPlanArgs struct {
	// Description of scaling plan.
	Description *string `pulumi:"description"`
	// Exclusion tag for scaling plan.
	ExclusionTag *string `pulumi:"exclusionTag"`
	// User friendly name of scaling plan.
	FriendlyName *string `pulumi:"friendlyName"`
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences []ScalingHostPoolReference `pulumi:"hostPoolReferences"`
	// HostPool type for desktop.
	HostPoolType *string                                      `pulumi:"hostPoolType"`
	Identity     *ResourceModelWithAllowedPropertySetIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy *string                                  `pulumi:"managedBy"`
	Plan      *ResourceModelWithAllowedPropertySetPlan `pulumi:"plan"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the scaling plan.
	ScalingPlanName *string `pulumi:"scalingPlanName"`
	// List of ScalingPlanPooledSchedule definitions.
	Schedules []ScalingSchedule                       `pulumi:"schedules"`
	Sku       *ResourceModelWithAllowedPropertySetSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Timezone of the scaling plan.
	TimeZone string `pulumi:"timeZone"`
}

// The set of arguments for constructing a ScalingPlan resource.
type ScalingPlanArgs struct {
	// Description of scaling plan.
	Description pulumi.StringPtrInput
	// Exclusion tag for scaling plan.
	ExclusionTag pulumi.StringPtrInput
	// User friendly name of scaling plan.
	FriendlyName pulumi.StringPtrInput
	// List of ScalingHostPoolReference definitions.
	HostPoolReferences ScalingHostPoolReferenceArrayInput
	// HostPool type for desktop.
	HostPoolType pulumi.StringPtrInput
	Identity     ResourceModelWithAllowedPropertySetIdentityPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrInput
	Plan      ResourceModelWithAllowedPropertySetPlanPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the scaling plan.
	ScalingPlanName pulumi.StringPtrInput
	// List of ScalingPlanPooledSchedule definitions.
	Schedules ScalingScheduleArrayInput
	Sku       ResourceModelWithAllowedPropertySetSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Timezone of the scaling plan.
	TimeZone pulumi.StringInput
}

func (ScalingPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanArgs)(nil)).Elem()
}

type ScalingPlanInput interface {
	pulumi.Input

	ToScalingPlanOutput() ScalingPlanOutput
	ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput
}

func (*ScalingPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlan)(nil)).Elem()
}

func (i *ScalingPlan) ToScalingPlanOutput() ScalingPlanOutput {
	return i.ToScalingPlanOutputWithContext(context.Background())
}

func (i *ScalingPlan) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanOutput)
}

type ScalingPlanOutput struct{ *pulumi.OutputState }

func (ScalingPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlan)(nil)).Elem()
}

func (o ScalingPlanOutput) ToScalingPlanOutput() ScalingPlanOutput {
	return o
}

func (o ScalingPlanOutput) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return o
}

// Description of scaling plan.
func (o ScalingPlanOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o ScalingPlanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Exclusion tag for scaling plan.
func (o ScalingPlanOutput) ExclusionTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.ExclusionTag }).(pulumi.StringPtrOutput)
}

// User friendly name of scaling plan.
func (o ScalingPlanOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// List of ScalingHostPoolReference definitions.
func (o ScalingPlanOutput) HostPoolReferences() ScalingHostPoolReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ScalingPlan) ScalingHostPoolReferenceResponseArrayOutput { return v.HostPoolReferences }).(ScalingHostPoolReferenceResponseArrayOutput)
}

// HostPool type for desktop.
func (o ScalingPlanOutput) HostPoolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.HostPoolType }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o ScalingPlanOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o ScalingPlanOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o ScalingPlanOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ScalingPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ObjectId of scaling plan. (internal use)
func (o ScalingPlanOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

func (o ScalingPlanOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

// List of ScalingPlanPooledSchedule definitions.
func (o ScalingPlanOutput) Schedules() ScalingScheduleResponseArrayOutput {
	return o.ApplyT(func(v *ScalingPlan) ScalingScheduleResponseArrayOutput { return v.Schedules }).(ScalingScheduleResponseArrayOutput)
}

func (o ScalingPlanOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v *ScalingPlan) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ScalingPlanOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScalingPlan) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ScalingPlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Timezone of the scaling plan.
func (o ScalingPlanOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScalingPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalingPlanOutput{})
}

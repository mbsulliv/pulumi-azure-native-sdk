// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get an application group.
// Azure REST API version: 2022-09-09.
//
// Other available API versions: 2020-11-10-preview, 2022-04-01-preview, 2022-10-14-preview, 2023-07-07-preview, 2023-09-05.
func LookupApplicationGroup(ctx *pulumi.Context, args *LookupApplicationGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationGroupResult
	err := ctx.Invoke("azure-native:desktopvirtualization:getApplicationGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGroupArgs struct {
	// The name of the application group
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a ApplicationGroup definition.
type LookupApplicationGroupResult struct {
	// Resource Type of ApplicationGroup.
	ApplicationGroupType string `pulumi:"applicationGroupType"`
	// Is cloud pc resource.
	CloudPcResource bool `pulumi:"cloudPcResource"`
	// Description of ApplicationGroup.
	Description *string `pulumi:"description"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag string `pulumi:"etag"`
	// Friendly name of ApplicationGroup.
	FriendlyName *string `pulumi:"friendlyName"`
	// HostPool arm path of ApplicationGroup.
	HostPoolArmPath string `pulumi:"hostPoolArmPath"`
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
	// ObjectId of ApplicationGroup. (internal use)
	ObjectId string                                           `pulumi:"objectId"`
	Plan     *ResourceModelWithAllowedPropertySetResponsePlan `pulumi:"plan"`
	Sku      *ResourceModelWithAllowedPropertySetResponseSku  `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Workspace arm path of ApplicationGroup.
	WorkspaceArmPath string `pulumi:"workspaceArmPath"`
}

func LookupApplicationGroupOutput(ctx *pulumi.Context, args LookupApplicationGroupOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationGroupResult, error) {
			args := v.(LookupApplicationGroupArgs)
			r, err := LookupApplicationGroup(ctx, &args, opts...)
			var s LookupApplicationGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationGroupResultOutput)
}

type LookupApplicationGroupOutputArgs struct {
	// The name of the application group
	ApplicationGroupName pulumi.StringInput `pulumi:"applicationGroupName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGroupArgs)(nil)).Elem()
}

// Represents a ApplicationGroup definition.
type LookupApplicationGroupResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGroupResult)(nil)).Elem()
}

func (o LookupApplicationGroupResultOutput) ToLookupApplicationGroupResultOutput() LookupApplicationGroupResultOutput {
	return o
}

func (o LookupApplicationGroupResultOutput) ToLookupApplicationGroupResultOutputWithContext(ctx context.Context) LookupApplicationGroupResultOutput {
	return o
}

func (o LookupApplicationGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApplicationGroupResult] {
	return pulumix.Output[LookupApplicationGroupResult]{
		OutputState: o.OutputState,
	}
}

// Resource Type of ApplicationGroup.
func (o LookupApplicationGroupResultOutput) ApplicationGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.ApplicationGroupType }).(pulumi.StringOutput)
}

// Is cloud pc resource.
func (o LookupApplicationGroupResultOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) bool { return v.CloudPcResource }).(pulumi.BoolOutput)
}

// Description of ApplicationGroup.
func (o LookupApplicationGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o LookupApplicationGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Friendly name of ApplicationGroup.
func (o LookupApplicationGroupResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// HostPool arm path of ApplicationGroup.
func (o LookupApplicationGroupResultOutput) HostPoolArmPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.HostPoolArmPath }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApplicationGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *ResourceModelWithAllowedPropertySetResponseIdentity {
		return v.Identity
	}).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupApplicationGroupResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupApplicationGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o LookupApplicationGroupResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupApplicationGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// ObjectId of ApplicationGroup. (internal use)
func (o LookupApplicationGroupResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupApplicationGroupResultOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *ResourceModelWithAllowedPropertySetResponsePlan { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

func (o LookupApplicationGroupResultOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) *ResourceModelWithAllowedPropertySetResponseSku { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApplicationGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupApplicationGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApplicationGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// Workspace arm path of ApplicationGroup.
func (o LookupApplicationGroupResultOutput) WorkspaceArmPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGroupResult) string { return v.WorkspaceArmPath }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationGroupResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230707preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a workspace.
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20230707preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a Workspace definition.
type LookupWorkspaceResult struct {
	// List of applicationGroup resource Ids.
	ApplicationGroupReferences []string `pulumi:"applicationGroupReferences"`
	// Is cloud pc resource.
	CloudPcResource bool `pulumi:"cloudPcResource"`
	// Description of Workspace.
	Description *string `pulumi:"description"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag string `pulumi:"etag"`
	// Friendly name of Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
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
	// ObjectId of Workspace. (internal use)
	ObjectId string                                           `pulumi:"objectId"`
	Plan     *ResourceModelWithAllowedPropertySetResponsePlan `pulumi:"plan"`
	// List of private endpoint connection associated with the specified resource
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
	PublicNetworkAccess *string                                         `pulumi:"publicNetworkAccess"`
	Sku                 *ResourceModelWithAllowedPropertySetResponseSku `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceOutput(ctx *pulumi.Context, args LookupWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceResult, error) {
			args := v.(LookupWorkspaceArgs)
			r, err := LookupWorkspace(ctx, &args, opts...)
			var s LookupWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceResultOutput)
}

type LookupWorkspaceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

// Represents a Workspace definition.
type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceResult] {
	return pulumix.Output[LookupWorkspaceResult]{
		OutputState: o.OutputState,
	}
}

// List of applicationGroup resource Ids.
func (o LookupWorkspaceResultOutput) ApplicationGroupReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []string { return v.ApplicationGroupReferences }).(pulumi.StringArrayOutput)
}

// Is cloud pc resource.
func (o LookupWorkspaceResultOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) bool { return v.CloudPcResource }).(pulumi.BoolOutput)
}

// Description of Workspace.
func (o LookupWorkspaceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o LookupWorkspaceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Friendly name of Workspace.
func (o LookupWorkspaceResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ResourceModelWithAllowedPropertySetResponseIdentity { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupWorkspaceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o LookupWorkspaceResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// ObjectId of Workspace. (internal use)
func (o LookupWorkspaceResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ResourceModelWithAllowedPropertySetResponsePlan { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

// List of private endpoint connection associated with the specified resource
func (o LookupWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
func (o LookupWorkspaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ResourceModelWithAllowedPropertySetResponseSku { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}

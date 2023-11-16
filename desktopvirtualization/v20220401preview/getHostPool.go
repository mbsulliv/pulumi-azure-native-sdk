// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a host pool.
func LookupHostPool(ctx *pulumi.Context, args *LookupHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupHostPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHostPoolResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20220401preview:getHostPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostPoolArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName string `pulumi:"hostPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a HostPool definition.
type LookupHostPoolResult struct {
	// The session host configuration for updating agent, monitoring agent, and stack component.
	AgentUpdate *AgentUpdatePropertiesResponse `pulumi:"agentUpdate"`
	// List of applicationGroup links.
	ApplicationGroupReferences []string `pulumi:"applicationGroupReferences"`
	// Is cloud pc resource.
	CloudPcResource bool `pulumi:"cloudPcResource"`
	// Custom rdp property of HostPool.
	CustomRdpProperty *string `pulumi:"customRdpProperty"`
	// Description of HostPool.
	Description *string `pulumi:"description"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag string `pulumi:"etag"`
	// Friendly name of HostPool.
	FriendlyName *string `pulumi:"friendlyName"`
	// HostPool type for desktop.
	HostPoolType string `pulumi:"hostPoolType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id       string                                               `pulumi:"id"`
	Identity *ResourceModelWithAllowedPropertySetResponseIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// The type of the load balancer.
	LoadBalancerType string `pulumi:"loadBalancerType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The max session limit of HostPool.
	MaxSessionLimit *int `pulumi:"maxSessionLimit"`
	// The registration info of HostPool.
	MigrationRequest *MigrationRequestPropertiesResponse `pulumi:"migrationRequest"`
	// The name of the resource
	Name string `pulumi:"name"`
	// ObjectId of HostPool. (internal use)
	ObjectId string `pulumi:"objectId"`
	// PersonalDesktopAssignment type for HostPool.
	PersonalDesktopAssignmentType *string                                          `pulumi:"personalDesktopAssignmentType"`
	Plan                          *ResourceModelWithAllowedPropertySetResponsePlan `pulumi:"plan"`
	// The type of preferred application group type, default to Desktop Application Group
	PreferredAppGroupType string `pulumi:"preferredAppGroupType"`
	// List of private endpoint connection associated with the specified resource
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The registration info of HostPool.
	RegistrationInfo *RegistrationInfoResponse `pulumi:"registrationInfo"`
	// The ring number of HostPool.
	Ring *int                                            `pulumi:"ring"`
	Sku  *ResourceModelWithAllowedPropertySetResponseSku `pulumi:"sku"`
	// ClientId for the registered Relying Party used to issue WVD SSO certificates.
	SsoClientId *string `pulumi:"ssoClientId"`
	// Path to Azure KeyVault storing the secret used for communication to ADFS.
	SsoClientSecretKeyVaultPath *string `pulumi:"ssoClientSecretKeyVaultPath"`
	// The type of single sign on Secret Type.
	SsoSecretType *string `pulumi:"ssoSecretType"`
	// URL to customer ADFS server for signing WVD SSO certificates.
	SsoadfsAuthority *string `pulumi:"ssoadfsAuthority"`
	// The flag to turn on/off StartVMOnConnect feature.
	StartVMOnConnect *bool `pulumi:"startVMOnConnect"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Is validation environment.
	ValidationEnvironment *bool `pulumi:"validationEnvironment"`
	// VM template for sessionhosts configuration within hostpool.
	VmTemplate *string `pulumi:"vmTemplate"`
}

func LookupHostPoolOutput(ctx *pulumi.Context, args LookupHostPoolOutputArgs, opts ...pulumi.InvokeOption) LookupHostPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostPoolResult, error) {
			args := v.(LookupHostPoolArgs)
			r, err := LookupHostPool(ctx, &args, opts...)
			var s LookupHostPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostPoolResultOutput)
}

type LookupHostPoolOutputArgs struct {
	// The name of the host pool within the specified resource group
	HostPoolName pulumi.StringInput `pulumi:"hostPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHostPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostPoolArgs)(nil)).Elem()
}

// Represents a HostPool definition.
type LookupHostPoolResultOutput struct{ *pulumi.OutputState }

func (LookupHostPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostPoolResult)(nil)).Elem()
}

func (o LookupHostPoolResultOutput) ToLookupHostPoolResultOutput() LookupHostPoolResultOutput {
	return o
}

func (o LookupHostPoolResultOutput) ToLookupHostPoolResultOutputWithContext(ctx context.Context) LookupHostPoolResultOutput {
	return o
}

// The session host configuration for updating agent, monitoring agent, and stack component.
func (o LookupHostPoolResultOutput) AgentUpdate() AgentUpdatePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *AgentUpdatePropertiesResponse { return v.AgentUpdate }).(AgentUpdatePropertiesResponsePtrOutput)
}

// List of applicationGroup links.
func (o LookupHostPoolResultOutput) ApplicationGroupReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostPoolResult) []string { return v.ApplicationGroupReferences }).(pulumi.StringArrayOutput)
}

// Is cloud pc resource.
func (o LookupHostPoolResultOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostPoolResult) bool { return v.CloudPcResource }).(pulumi.BoolOutput)
}

// Custom rdp property of HostPool.
func (o LookupHostPoolResultOutput) CustomRdpProperty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.CustomRdpProperty }).(pulumi.StringPtrOutput)
}

// Description of HostPool.
func (o LookupHostPoolResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o LookupHostPoolResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Friendly name of HostPool.
func (o LookupHostPoolResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// HostPool type for desktop.
func (o LookupHostPoolResultOutput) HostPoolType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.HostPoolType }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupHostPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *ResourceModelWithAllowedPropertySetResponseIdentity { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupHostPoolResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The type of the load balancer.
func (o LookupHostPoolResultOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.LoadBalancerType }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupHostPoolResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o LookupHostPoolResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The max session limit of HostPool.
func (o LookupHostPoolResultOutput) MaxSessionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *int { return v.MaxSessionLimit }).(pulumi.IntPtrOutput)
}

// The registration info of HostPool.
func (o LookupHostPoolResultOutput) MigrationRequest() MigrationRequestPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *MigrationRequestPropertiesResponse { return v.MigrationRequest }).(MigrationRequestPropertiesResponsePtrOutput)
}

// The name of the resource
func (o LookupHostPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// ObjectId of HostPool. (internal use)
func (o LookupHostPoolResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// PersonalDesktopAssignment type for HostPool.
func (o LookupHostPoolResultOutput) PersonalDesktopAssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.PersonalDesktopAssignmentType }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *ResourceModelWithAllowedPropertySetResponsePlan { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

// The type of preferred application group type, default to Desktop Application Group
func (o LookupHostPoolResultOutput) PreferredAppGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.PreferredAppGroupType }).(pulumi.StringOutput)
}

// List of private endpoint connection associated with the specified resource
func (o LookupHostPoolResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupHostPoolResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
func (o LookupHostPoolResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The registration info of HostPool.
func (o LookupHostPoolResultOutput) RegistrationInfo() RegistrationInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *RegistrationInfoResponse { return v.RegistrationInfo }).(RegistrationInfoResponsePtrOutput)
}

// The ring number of HostPool.
func (o LookupHostPoolResultOutput) Ring() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *int { return v.Ring }).(pulumi.IntPtrOutput)
}

func (o LookupHostPoolResultOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *ResourceModelWithAllowedPropertySetResponseSku { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

// ClientId for the registered Relying Party used to issue WVD SSO certificates.
func (o LookupHostPoolResultOutput) SsoClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoClientId }).(pulumi.StringPtrOutput)
}

// Path to Azure KeyVault storing the secret used for communication to ADFS.
func (o LookupHostPoolResultOutput) SsoClientSecretKeyVaultPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoClientSecretKeyVaultPath }).(pulumi.StringPtrOutput)
}

// The type of single sign on Secret Type.
func (o LookupHostPoolResultOutput) SsoSecretType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoSecretType }).(pulumi.StringPtrOutput)
}

// URL to customer ADFS server for signing WVD SSO certificates.
func (o LookupHostPoolResultOutput) SsoadfsAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoadfsAuthority }).(pulumi.StringPtrOutput)
}

// The flag to turn on/off StartVMOnConnect feature.
func (o LookupHostPoolResultOutput) StartVMOnConnect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *bool { return v.StartVMOnConnect }).(pulumi.BoolPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupHostPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHostPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupHostPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHostPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupHostPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// Is validation environment.
func (o LookupHostPoolResultOutput) ValidationEnvironment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *bool { return v.ValidationEnvironment }).(pulumi.BoolPtrOutput)
}

// VM template for sessionhosts configuration within hostpool.
func (o LookupHostPoolResultOutput) VmTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.VmTemplate }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostPoolResultOutput{})
}

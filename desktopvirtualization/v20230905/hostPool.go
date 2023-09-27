// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230905

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a HostPool definition.
type HostPool struct {
	pulumi.CustomResourceState

	// The session host configuration for updating agent, monitoring agent, and stack component.
	AgentUpdate AgentUpdatePropertiesResponsePtrOutput `pulumi:"agentUpdate"`
	// List of applicationGroup links.
	ApplicationGroupReferences pulumi.StringArrayOutput `pulumi:"applicationGroupReferences"`
	// Is cloud pc resource.
	CloudPcResource pulumi.BoolOutput `pulumi:"cloudPcResource"`
	// Custom rdp property of HostPool.
	CustomRdpProperty pulumi.StringPtrOutput `pulumi:"customRdpProperty"`
	// Description of HostPool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Friendly name of HostPool.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// HostPool type for desktop.
	HostPoolType pulumi.StringOutput                                          `pulumi:"hostPoolType"`
	Identity     ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The type of the load balancer.
	LoadBalancerType pulumi.StringOutput `pulumi:"loadBalancerType"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// The max session limit of HostPool.
	MaxSessionLimit pulumi.IntPtrOutput `pulumi:"maxSessionLimit"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ObjectId of HostPool. (internal use)
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// PersonalDesktopAssignment type for HostPool.
	PersonalDesktopAssignmentType pulumi.StringPtrOutput                                   `pulumi:"personalDesktopAssignmentType"`
	Plan                          ResourceModelWithAllowedPropertySetResponsePlanPtrOutput `pulumi:"plan"`
	// The type of preferred application group type, default to Desktop Application Group
	PreferredAppGroupType pulumi.StringOutput `pulumi:"preferredAppGroupType"`
	// List of private endpoint connection associated with the specified resource
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The registration info of HostPool.
	RegistrationInfo RegistrationInfoResponsePtrOutput `pulumi:"registrationInfo"`
	// The ring number of HostPool.
	Ring pulumi.IntPtrOutput                                     `pulumi:"ring"`
	Sku  ResourceModelWithAllowedPropertySetResponseSkuPtrOutput `pulumi:"sku"`
	// ClientId for the registered Relying Party used to issue WVD SSO certificates.
	SsoClientId pulumi.StringPtrOutput `pulumi:"ssoClientId"`
	// Path to Azure KeyVault storing the secret used for communication to ADFS.
	SsoClientSecretKeyVaultPath pulumi.StringPtrOutput `pulumi:"ssoClientSecretKeyVaultPath"`
	// The type of single sign on Secret Type.
	SsoSecretType pulumi.StringPtrOutput `pulumi:"ssoSecretType"`
	// URL to customer ADFS server for signing WVD SSO certificates.
	SsoadfsAuthority pulumi.StringPtrOutput `pulumi:"ssoadfsAuthority"`
	// The flag to turn on/off StartVMOnConnect feature.
	StartVMOnConnect pulumi.BoolPtrOutput `pulumi:"startVMOnConnect"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Is validation environment.
	ValidationEnvironment pulumi.BoolPtrOutput `pulumi:"validationEnvironment"`
	// VM template for sessionhosts configuration within hostpool.
	VmTemplate pulumi.StringPtrOutput `pulumi:"vmTemplate"`
}

// NewHostPool registers a new resource with the given unique name, arguments, and options.
func NewHostPool(ctx *pulumi.Context,
	name string, args *HostPoolArgs, opts ...pulumi.ResourceOption) (*HostPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostPoolType == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolType'")
	}
	if args.LoadBalancerType == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerType'")
	}
	if args.PreferredAppGroupType == nil {
		return nil, errors.New("invalid value for required argument 'PreferredAppGroupType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220909:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:HostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230707preview:HostPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HostPool
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20230905:HostPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostPool gets an existing HostPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostPoolState, opts ...pulumi.ResourceOption) (*HostPool, error) {
	var resource HostPool
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20230905:HostPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostPool resources.
type hostPoolState struct {
}

type HostPoolState struct {
}

func (HostPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolState)(nil)).Elem()
}

type hostPoolArgs struct {
	// The session host configuration for updating agent, monitoring agent, and stack component.
	AgentUpdate *AgentUpdateProperties `pulumi:"agentUpdate"`
	// Custom rdp property of HostPool.
	CustomRdpProperty *string `pulumi:"customRdpProperty"`
	// Description of HostPool.
	Description *string `pulumi:"description"`
	// Friendly name of HostPool.
	FriendlyName *string `pulumi:"friendlyName"`
	// The name of the host pool within the specified resource group
	HostPoolName *string `pulumi:"hostPoolName"`
	// HostPool type for desktop.
	HostPoolType string                                       `pulumi:"hostPoolType"`
	Identity     *ResourceModelWithAllowedPropertySetIdentity `pulumi:"identity"`
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
	// PersonalDesktopAssignment type for HostPool.
	PersonalDesktopAssignmentType *string                                  `pulumi:"personalDesktopAssignmentType"`
	Plan                          *ResourceModelWithAllowedPropertySetPlan `pulumi:"plan"`
	// The type of preferred application group type, default to Desktop Application Group
	PreferredAppGroupType string `pulumi:"preferredAppGroupType"`
	// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The registration info of HostPool.
	RegistrationInfo *RegistrationInfo `pulumi:"registrationInfo"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ring number of HostPool.
	Ring *int                                    `pulumi:"ring"`
	Sku  *ResourceModelWithAllowedPropertySetSku `pulumi:"sku"`
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
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Is validation environment.
	ValidationEnvironment *bool `pulumi:"validationEnvironment"`
	// VM template for sessionhosts configuration within hostpool.
	VmTemplate *string `pulumi:"vmTemplate"`
}

// The set of arguments for constructing a HostPool resource.
type HostPoolArgs struct {
	// The session host configuration for updating agent, monitoring agent, and stack component.
	AgentUpdate AgentUpdatePropertiesPtrInput
	// Custom rdp property of HostPool.
	CustomRdpProperty pulumi.StringPtrInput
	// Description of HostPool.
	Description pulumi.StringPtrInput
	// Friendly name of HostPool.
	FriendlyName pulumi.StringPtrInput
	// The name of the host pool within the specified resource group
	HostPoolName pulumi.StringPtrInput
	// HostPool type for desktop.
	HostPoolType pulumi.StringInput
	Identity     ResourceModelWithAllowedPropertySetIdentityPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// The type of the load balancer.
	LoadBalancerType pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrInput
	// The max session limit of HostPool.
	MaxSessionLimit pulumi.IntPtrInput
	// PersonalDesktopAssignment type for HostPool.
	PersonalDesktopAssignmentType pulumi.StringPtrInput
	Plan                          ResourceModelWithAllowedPropertySetPlanPtrInput
	// The type of preferred application group type, default to Desktop Application Group
	PreferredAppGroupType pulumi.StringInput
	// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
	PublicNetworkAccess pulumi.StringPtrInput
	// The registration info of HostPool.
	RegistrationInfo RegistrationInfoPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The ring number of HostPool.
	Ring pulumi.IntPtrInput
	Sku  ResourceModelWithAllowedPropertySetSkuPtrInput
	// ClientId for the registered Relying Party used to issue WVD SSO certificates.
	SsoClientId pulumi.StringPtrInput
	// Path to Azure KeyVault storing the secret used for communication to ADFS.
	SsoClientSecretKeyVaultPath pulumi.StringPtrInput
	// The type of single sign on Secret Type.
	SsoSecretType pulumi.StringPtrInput
	// URL to customer ADFS server for signing WVD SSO certificates.
	SsoadfsAuthority pulumi.StringPtrInput
	// The flag to turn on/off StartVMOnConnect feature.
	StartVMOnConnect pulumi.BoolPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Is validation environment.
	ValidationEnvironment pulumi.BoolPtrInput
	// VM template for sessionhosts configuration within hostpool.
	VmTemplate pulumi.StringPtrInput
}

func (HostPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPoolArgs)(nil)).Elem()
}

type HostPoolInput interface {
	pulumi.Input

	ToHostPoolOutput() HostPoolOutput
	ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput
}

func (*HostPool) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPool)(nil)).Elem()
}

func (i *HostPool) ToHostPoolOutput() HostPoolOutput {
	return i.ToHostPoolOutputWithContext(context.Background())
}

func (i *HostPool) ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPoolOutput)
}

func (i *HostPool) ToOutput(ctx context.Context) pulumix.Output[*HostPool] {
	return pulumix.Output[*HostPool]{
		OutputState: i.ToHostPoolOutputWithContext(ctx).OutputState,
	}
}

type HostPoolOutput struct{ *pulumi.OutputState }

func (HostPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPool)(nil)).Elem()
}

func (o HostPoolOutput) ToHostPoolOutput() HostPoolOutput {
	return o
}

func (o HostPoolOutput) ToHostPoolOutputWithContext(ctx context.Context) HostPoolOutput {
	return o
}

func (o HostPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*HostPool] {
	return pulumix.Output[*HostPool]{
		OutputState: o.OutputState,
	}
}

// The session host configuration for updating agent, monitoring agent, and stack component.
func (o HostPoolOutput) AgentUpdate() AgentUpdatePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *HostPool) AgentUpdatePropertiesResponsePtrOutput { return v.AgentUpdate }).(AgentUpdatePropertiesResponsePtrOutput)
}

// List of applicationGroup links.
func (o HostPoolOutput) ApplicationGroupReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringArrayOutput { return v.ApplicationGroupReferences }).(pulumi.StringArrayOutput)
}

// Is cloud pc resource.
func (o HostPoolOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v *HostPool) pulumi.BoolOutput { return v.CloudPcResource }).(pulumi.BoolOutput)
}

// Custom rdp property of HostPool.
func (o HostPoolOutput) CustomRdpProperty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.CustomRdpProperty }).(pulumi.StringPtrOutput)
}

// Description of HostPool.
func (o HostPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o HostPoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Friendly name of HostPool.
func (o HostPoolOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// HostPool type for desktop.
func (o HostPoolOutput) HostPoolType() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.HostPoolType }).(pulumi.StringOutput)
}

func (o HostPoolOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v *HostPool) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o HostPoolOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The type of the load balancer.
func (o HostPoolOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.LoadBalancerType }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o HostPoolOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o HostPoolOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The max session limit of HostPool.
func (o HostPoolOutput) MaxSessionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.IntPtrOutput { return v.MaxSessionLimit }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o HostPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ObjectId of HostPool. (internal use)
func (o HostPoolOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

// PersonalDesktopAssignment type for HostPool.
func (o HostPoolOutput) PersonalDesktopAssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.PersonalDesktopAssignmentType }).(pulumi.StringPtrOutput)
}

func (o HostPoolOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v *HostPool) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

// The type of preferred application group type, default to Desktop Application Group
func (o HostPoolOutput) PreferredAppGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.PreferredAppGroupType }).(pulumi.StringOutput)
}

// List of private endpoint connection associated with the specified resource
func (o HostPoolOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *HostPool) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Enabled allows this resource to be accessed from both public and private networks, Disabled allows this resource to only be accessed via private endpoints
func (o HostPoolOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The registration info of HostPool.
func (o HostPoolOutput) RegistrationInfo() RegistrationInfoResponsePtrOutput {
	return o.ApplyT(func(v *HostPool) RegistrationInfoResponsePtrOutput { return v.RegistrationInfo }).(RegistrationInfoResponsePtrOutput)
}

// The ring number of HostPool.
func (o HostPoolOutput) Ring() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.IntPtrOutput { return v.Ring }).(pulumi.IntPtrOutput)
}

func (o HostPoolOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v *HostPool) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

// ClientId for the registered Relying Party used to issue WVD SSO certificates.
func (o HostPoolOutput) SsoClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.SsoClientId }).(pulumi.StringPtrOutput)
}

// Path to Azure KeyVault storing the secret used for communication to ADFS.
func (o HostPoolOutput) SsoClientSecretKeyVaultPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.SsoClientSecretKeyVaultPath }).(pulumi.StringPtrOutput)
}

// The type of single sign on Secret Type.
func (o HostPoolOutput) SsoSecretType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.SsoSecretType }).(pulumi.StringPtrOutput)
}

// URL to customer ADFS server for signing WVD SSO certificates.
func (o HostPoolOutput) SsoadfsAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.SsoadfsAuthority }).(pulumi.StringPtrOutput)
}

// The flag to turn on/off StartVMOnConnect feature.
func (o HostPoolOutput) StartVMOnConnect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.BoolPtrOutput { return v.StartVMOnConnect }).(pulumi.BoolPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o HostPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HostPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o HostPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HostPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Is validation environment.
func (o HostPoolOutput) ValidationEnvironment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.BoolPtrOutput { return v.ValidationEnvironment }).(pulumi.BoolPtrOutput)
}

// VM template for sessionhosts configuration within hostpool.
func (o HostPoolOutput) VmTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostPool) pulumi.StringPtrOutput { return v.VmTemplate }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HostPoolOutput{})
}

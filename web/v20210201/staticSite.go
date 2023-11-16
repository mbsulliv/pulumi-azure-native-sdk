// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Static Site ARM resource.
type StaticSite struct {
	pulumi.CustomResourceState

	// <code>false</code> if config file is locked for this static web app; otherwise, <code>true</code>.
	AllowConfigFileUpdates pulumi.BoolPtrOutput `pulumi:"allowConfigFileUpdates"`
	// The target branch in the repository.
	Branch pulumi.StringPtrOutput `pulumi:"branch"`
	// Build properties to configure on the repository.
	BuildProperties StaticSiteBuildPropertiesResponsePtrOutput `pulumi:"buildProperties"`
	// The content distribution endpoint for the static site.
	ContentDistributionEndpoint pulumi.StringOutput `pulumi:"contentDistributionEndpoint"`
	// The custom domains associated with this static site.
	CustomDomains pulumi.StringArrayOutput `pulumi:"customDomains"`
	// The default autogenerated hostname for the static site.
	DefaultHostname pulumi.StringOutput `pulumi:"defaultHostname"`
	// Managed service identity.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Identity to use for Key Vault Reference authentication.
	KeyVaultReferenceIdentity pulumi.StringOutput `pulumi:"keyVaultReferenceIdentity"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint connections
	PrivateEndpointConnections ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The provider that submitted the last deployment to the primary environment of the static site.
	Provider pulumi.StringOutput `pulumi:"provider"`
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken pulumi.StringPtrOutput `pulumi:"repositoryToken"`
	// URL for the repository of the static site.
	RepositoryUrl pulumi.StringPtrOutput `pulumi:"repositoryUrl"`
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrOutput `pulumi:"sku"`
	// State indicating whether staging environments are allowed or not allowed for a static web app.
	StagingEnvironmentPolicy pulumi.StringPtrOutput `pulumi:"stagingEnvironmentPolicy"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Template options for generating a new repository.
	TemplateProperties StaticSiteTemplateOptionsResponsePtrOutput `pulumi:"templateProperties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// User provided function apps registered with the static site
	UserProvidedFunctionApps StaticSiteUserProvidedFunctionAppResponseArrayOutput `pulumi:"userProvidedFunctionApps"`
}

// NewStaticSite registers a new resource with the given unique name, arguments, and options.
func NewStaticSite(ctx *pulumi.Context,
	name string, args *StaticSiteArgs, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:StaticSite"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StaticSite
	err := ctx.RegisterResource("azure-native:web/v20210201:StaticSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticSite gets an existing StaticSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteState, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	var resource StaticSite
	err := ctx.ReadResource("azure-native:web/v20210201:StaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticSite resources.
type staticSiteState struct {
}

type StaticSiteState struct {
}

func (StaticSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteState)(nil)).Elem()
}

type staticSiteArgs struct {
	// <code>false</code> if config file is locked for this static web app; otherwise, <code>true</code>.
	AllowConfigFileUpdates *bool `pulumi:"allowConfigFileUpdates"`
	// The target branch in the repository.
	Branch *string `pulumi:"branch"`
	// Build properties to configure on the repository.
	BuildProperties *StaticSiteBuildProperties `pulumi:"buildProperties"`
	// Managed service identity.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Name of the static site to create or update.
	Name *string `pulumi:"name"`
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken *string `pulumi:"repositoryToken"`
	// URL for the repository of the static site.
	RepositoryUrl *string `pulumi:"repositoryUrl"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescription `pulumi:"sku"`
	// State indicating whether staging environments are allowed or not allowed for a static web app.
	StagingEnvironmentPolicy *StagingEnvironmentPolicy `pulumi:"stagingEnvironmentPolicy"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Template options for generating a new repository.
	TemplateProperties *StaticSiteTemplateOptions `pulumi:"templateProperties"`
}

// The set of arguments for constructing a StaticSite resource.
type StaticSiteArgs struct {
	// <code>false</code> if config file is locked for this static web app; otherwise, <code>true</code>.
	AllowConfigFileUpdates pulumi.BoolPtrInput
	// The target branch in the repository.
	Branch pulumi.StringPtrInput
	// Build properties to configure on the repository.
	BuildProperties StaticSiteBuildPropertiesPtrInput
	// Managed service identity.
	Identity ManagedServiceIdentityPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Name of the static site to create or update.
	Name pulumi.StringPtrInput
	// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
	RepositoryToken pulumi.StringPtrInput
	// URL for the repository of the static site.
	RepositoryUrl pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionPtrInput
	// State indicating whether staging environments are allowed or not allowed for a static web app.
	StagingEnvironmentPolicy StagingEnvironmentPolicyPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Template options for generating a new repository.
	TemplateProperties StaticSiteTemplateOptionsPtrInput
}

func (StaticSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteArgs)(nil)).Elem()
}

type StaticSiteInput interface {
	pulumi.Input

	ToStaticSiteOutput() StaticSiteOutput
	ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput
}

func (*StaticSite) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSite)(nil)).Elem()
}

func (i *StaticSite) ToStaticSiteOutput() StaticSiteOutput {
	return i.ToStaticSiteOutputWithContext(context.Background())
}

func (i *StaticSite) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteOutput)
}

type StaticSiteOutput struct{ *pulumi.OutputState }

func (StaticSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSite)(nil)).Elem()
}

func (o StaticSiteOutput) ToStaticSiteOutput() StaticSiteOutput {
	return o
}

func (o StaticSiteOutput) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return o
}

// <code>false</code> if config file is locked for this static web app; otherwise, <code>true</code>.
func (o StaticSiteOutput) AllowConfigFileUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.BoolPtrOutput { return v.AllowConfigFileUpdates }).(pulumi.BoolPtrOutput)
}

// The target branch in the repository.
func (o StaticSiteOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

// Build properties to configure on the repository.
func (o StaticSiteOutput) BuildProperties() StaticSiteBuildPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *StaticSite) StaticSiteBuildPropertiesResponsePtrOutput { return v.BuildProperties }).(StaticSiteBuildPropertiesResponsePtrOutput)
}

// The content distribution endpoint for the static site.
func (o StaticSiteOutput) ContentDistributionEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringOutput { return v.ContentDistributionEndpoint }).(pulumi.StringOutput)
}

// The custom domains associated with this static site.
func (o StaticSiteOutput) CustomDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringArrayOutput { return v.CustomDomains }).(pulumi.StringArrayOutput)
}

// The default autogenerated hostname for the static site.
func (o StaticSiteOutput) DefaultHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringOutput { return v.DefaultHostname }).(pulumi.StringOutput)
}

// Managed service identity.
func (o StaticSiteOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *StaticSite) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Identity to use for Key Vault Reference authentication.
func (o StaticSiteOutput) KeyVaultReferenceIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringOutput { return v.KeyVaultReferenceIdentity }).(pulumi.StringOutput)
}

// Kind of resource.
func (o StaticSiteOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location.
func (o StaticSiteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource Name.
func (o StaticSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint connections
func (o StaticSiteOutput) PrivateEndpointConnections() ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *StaticSite) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput)
}

// The provider that submitted the last deployment to the primary environment of the static site.
func (o StaticSiteOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringOutput { return v.Provider }).(pulumi.StringOutput)
}

// A user's github repository token. This is used to setup the Github Actions workflow file and API secrets.
func (o StaticSiteOutput) RepositoryToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.RepositoryToken }).(pulumi.StringPtrOutput)
}

// URL for the repository of the static site.
func (o StaticSiteOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.RepositoryUrl }).(pulumi.StringPtrOutput)
}

// Description of a SKU for a scalable resource.
func (o StaticSiteOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *StaticSite) SkuDescriptionResponsePtrOutput { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

// State indicating whether staging environments are allowed or not allowed for a static web app.
func (o StaticSiteOutput) StagingEnvironmentPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringPtrOutput { return v.StagingEnvironmentPolicy }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o StaticSiteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Template options for generating a new repository.
func (o StaticSiteOutput) TemplateProperties() StaticSiteTemplateOptionsResponsePtrOutput {
	return o.ApplyT(func(v *StaticSite) StaticSiteTemplateOptionsResponsePtrOutput { return v.TemplateProperties }).(StaticSiteTemplateOptionsResponsePtrOutput)
}

// Resource type.
func (o StaticSiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// User provided function apps registered with the static site
func (o StaticSiteOutput) UserProvidedFunctionApps() StaticSiteUserProvidedFunctionAppResponseArrayOutput {
	return o.ApplyT(func(v *StaticSite) StaticSiteUserProvidedFunctionAppResponseArrayOutput {
		return v.UserProvidedFunctionApps
	}).(StaticSiteUserProvidedFunctionAppResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSiteOutput{})
}

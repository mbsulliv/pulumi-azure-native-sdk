// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Template resource definition.
type ContentTemplate struct {
	pulumi.CustomResourceState

	// The creator of the content item.
	Author MetadataAuthorResponsePtrOutput `pulumi:"author"`
	// Categories for the item
	Categories MetadataCategoriesResponsePtrOutput `pulumi:"categories"`
	// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
	ContentId pulumi.StringOutput `pulumi:"contentId"`
	// The kind of content the template is for.
	ContentKind pulumi.StringOutput `pulumi:"contentKind"`
	// Unique ID for the content. It should be generated based on the contentId of the package, contentId of the template, contentKind of the template and the contentVersion of the template
	ContentProductId pulumi.StringOutput `pulumi:"contentProductId"`
	// Schema version of the content. Can be used to distinguish between different flow based on the schema version
	ContentSchemaVersion pulumi.StringPtrOutput `pulumi:"contentSchemaVersion"`
	// The custom version of the content. A optional free text
	CustomVersion pulumi.StringPtrOutput `pulumi:"customVersion"`
	// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
	Dependencies MetadataDependenciesResponsePtrOutput `pulumi:"dependencies"`
	// The display name of the template
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// first publish date content item
	FirstPublishDate pulumi.StringPtrOutput `pulumi:"firstPublishDate"`
	// the icon identifier. this id can later be fetched from the content metadata
	Icon pulumi.StringPtrOutput `pulumi:"icon"`
	// last publish date for the content item
	LastPublishDate pulumi.StringPtrOutput `pulumi:"lastPublishDate"`
	// The JSON of the ARM template to deploy active content
	MainTemplate pulumi.AnyOutput `pulumi:"mainTemplate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// the package Id contains this template
	PackageId pulumi.StringOutput `pulumi:"packageId"`
	// the packageKind of the package contains this template
	PackageKind pulumi.StringPtrOutput `pulumi:"packageKind"`
	// the name of the package contains this template
	PackageName pulumi.StringPtrOutput `pulumi:"packageName"`
	// Version of the package.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
	PackageVersion pulumi.StringOutput `pulumi:"packageVersion"`
	// preview image file names. These will be taken from the solution artifacts
	PreviewImages pulumi.StringArrayOutput `pulumi:"previewImages"`
	// preview image file names. These will be taken from the solution artifacts. used for dark theme support
	PreviewImagesDark pulumi.StringArrayOutput `pulumi:"previewImagesDark"`
	// Providers for the content item
	Providers pulumi.StringArrayOutput `pulumi:"providers"`
	// Source of the content.  This is where/how it was created.
	Source MetadataSourceResponseOutput `pulumi:"source"`
	// Support information for the template - type, name, contact information
	Support MetadataSupportResponsePtrOutput `pulumi:"support"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// the tactics the resource covers
	ThreatAnalysisTactics pulumi.StringArrayOutput `pulumi:"threatAnalysisTactics"`
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques pulumi.StringArrayOutput `pulumi:"threatAnalysisTechniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewContentTemplate registers a new resource with the given unique name, arguments, and options.
func NewContentTemplate(ctx *pulumi.Context,
	name string, args *ContentTemplateArgs, opts ...pulumi.ResourceOption) (*ContentTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentId == nil {
		return nil, errors.New("invalid value for required argument 'ContentId'")
	}
	if args.ContentKind == nil {
		return nil, errors.New("invalid value for required argument 'ContentKind'")
	}
	if args.ContentProductId == nil {
		return nil, errors.New("invalid value for required argument 'ContentProductId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.PackageId == nil {
		return nil, errors.New("invalid value for required argument 'PackageId'")
	}
	if args.PackageVersion == nil {
		return nil, errors.New("invalid value for required argument 'PackageVersion'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ContentTemplate"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:ContentTemplate"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:ContentTemplate"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:ContentTemplate"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:ContentTemplate"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:ContentTemplate"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:ContentTemplate"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:ContentTemplate"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:ContentTemplate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ContentTemplate
	err := ctx.RegisterResource("azure-native:securityinsights/v20230901preview:ContentTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentTemplate gets an existing ContentTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentTemplateState, opts ...pulumi.ResourceOption) (*ContentTemplate, error) {
	var resource ContentTemplate
	err := ctx.ReadResource("azure-native:securityinsights/v20230901preview:ContentTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContentTemplate resources.
type contentTemplateState struct {
}

type ContentTemplateState struct {
}

func (ContentTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentTemplateState)(nil)).Elem()
}

type contentTemplateArgs struct {
	// The creator of the content item.
	Author *MetadataAuthor `pulumi:"author"`
	// Categories for the item
	Categories *MetadataCategories `pulumi:"categories"`
	// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
	ContentId string `pulumi:"contentId"`
	// The kind of content the template is for.
	ContentKind string `pulumi:"contentKind"`
	// Unique ID for the content. It should be generated based on the contentId of the package, contentId of the template, contentKind of the template and the contentVersion of the template
	ContentProductId string `pulumi:"contentProductId"`
	// Schema version of the content. Can be used to distinguish between different flow based on the schema version
	ContentSchemaVersion *string `pulumi:"contentSchemaVersion"`
	// The custom version of the content. A optional free text
	CustomVersion *string `pulumi:"customVersion"`
	// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
	Dependencies *MetadataDependencies `pulumi:"dependencies"`
	// The display name of the template
	DisplayName string `pulumi:"displayName"`
	// first publish date content item
	FirstPublishDate *string `pulumi:"firstPublishDate"`
	// the icon identifier. this id can later be fetched from the content metadata
	Icon *string `pulumi:"icon"`
	// last publish date for the content item
	LastPublishDate *string `pulumi:"lastPublishDate"`
	// The JSON of the ARM template to deploy active content
	MainTemplate interface{} `pulumi:"mainTemplate"`
	// the package Id contains this template
	PackageId string `pulumi:"packageId"`
	// the packageKind of the package contains this template
	PackageKind *string `pulumi:"packageKind"`
	// the name of the package contains this template
	PackageName *string `pulumi:"packageName"`
	// Version of the package.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
	PackageVersion string `pulumi:"packageVersion"`
	// preview image file names. These will be taken from the solution artifacts
	PreviewImages []string `pulumi:"previewImages"`
	// preview image file names. These will be taken from the solution artifacts. used for dark theme support
	PreviewImagesDark []string `pulumi:"previewImagesDark"`
	// Providers for the content item
	Providers []string `pulumi:"providers"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Source of the content.  This is where/how it was created.
	Source MetadataSource `pulumi:"source"`
	// Support information for the template - type, name, contact information
	Support *MetadataSupport `pulumi:"support"`
	// template Id
	TemplateId *string `pulumi:"templateId"`
	// the tactics the resource covers
	ThreatAnalysisTactics []string `pulumi:"threatAnalysisTactics"`
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques []string `pulumi:"threatAnalysisTechniques"`
	// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
	Version string `pulumi:"version"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ContentTemplate resource.
type ContentTemplateArgs struct {
	// The creator of the content item.
	Author MetadataAuthorPtrInput
	// Categories for the item
	Categories MetadataCategoriesPtrInput
	// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
	ContentId pulumi.StringInput
	// The kind of content the template is for.
	ContentKind pulumi.StringInput
	// Unique ID for the content. It should be generated based on the contentId of the package, contentId of the template, contentKind of the template and the contentVersion of the template
	ContentProductId pulumi.StringInput
	// Schema version of the content. Can be used to distinguish between different flow based on the schema version
	ContentSchemaVersion pulumi.StringPtrInput
	// The custom version of the content. A optional free text
	CustomVersion pulumi.StringPtrInput
	// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
	Dependencies MetadataDependenciesPtrInput
	// The display name of the template
	DisplayName pulumi.StringInput
	// first publish date content item
	FirstPublishDate pulumi.StringPtrInput
	// the icon identifier. this id can later be fetched from the content metadata
	Icon pulumi.StringPtrInput
	// last publish date for the content item
	LastPublishDate pulumi.StringPtrInput
	// The JSON of the ARM template to deploy active content
	MainTemplate pulumi.Input
	// the package Id contains this template
	PackageId pulumi.StringInput
	// the packageKind of the package contains this template
	PackageKind pulumi.StringPtrInput
	// the name of the package contains this template
	PackageName pulumi.StringPtrInput
	// Version of the package.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
	PackageVersion pulumi.StringInput
	// preview image file names. These will be taken from the solution artifacts
	PreviewImages pulumi.StringArrayInput
	// preview image file names. These will be taken from the solution artifacts. used for dark theme support
	PreviewImagesDark pulumi.StringArrayInput
	// Providers for the content item
	Providers pulumi.StringArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Source of the content.  This is where/how it was created.
	Source MetadataSourceInput
	// Support information for the template - type, name, contact information
	Support MetadataSupportPtrInput
	// template Id
	TemplateId pulumi.StringPtrInput
	// the tactics the resource covers
	ThreatAnalysisTactics pulumi.StringArrayInput
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques pulumi.StringArrayInput
	// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
	Version pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ContentTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentTemplateArgs)(nil)).Elem()
}

type ContentTemplateInput interface {
	pulumi.Input

	ToContentTemplateOutput() ContentTemplateOutput
	ToContentTemplateOutputWithContext(ctx context.Context) ContentTemplateOutput
}

func (*ContentTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentTemplate)(nil)).Elem()
}

func (i *ContentTemplate) ToContentTemplateOutput() ContentTemplateOutput {
	return i.ToContentTemplateOutputWithContext(context.Background())
}

func (i *ContentTemplate) ToContentTemplateOutputWithContext(ctx context.Context) ContentTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentTemplateOutput)
}

type ContentTemplateOutput struct{ *pulumi.OutputState }

func (ContentTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentTemplate)(nil)).Elem()
}

func (o ContentTemplateOutput) ToContentTemplateOutput() ContentTemplateOutput {
	return o
}

func (o ContentTemplateOutput) ToContentTemplateOutputWithContext(ctx context.Context) ContentTemplateOutput {
	return o
}

// The creator of the content item.
func (o ContentTemplateOutput) Author() MetadataAuthorResponsePtrOutput {
	return o.ApplyT(func(v *ContentTemplate) MetadataAuthorResponsePtrOutput { return v.Author }).(MetadataAuthorResponsePtrOutput)
}

// Categories for the item
func (o ContentTemplateOutput) Categories() MetadataCategoriesResponsePtrOutput {
	return o.ApplyT(func(v *ContentTemplate) MetadataCategoriesResponsePtrOutput { return v.Categories }).(MetadataCategoriesResponsePtrOutput)
}

// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
func (o ContentTemplateOutput) ContentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.ContentId }).(pulumi.StringOutput)
}

// The kind of content the template is for.
func (o ContentTemplateOutput) ContentKind() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.ContentKind }).(pulumi.StringOutput)
}

// Unique ID for the content. It should be generated based on the contentId of the package, contentId of the template, contentKind of the template and the contentVersion of the template
func (o ContentTemplateOutput) ContentProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.ContentProductId }).(pulumi.StringOutput)
}

// Schema version of the content. Can be used to distinguish between different flow based on the schema version
func (o ContentTemplateOutput) ContentSchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringPtrOutput { return v.ContentSchemaVersion }).(pulumi.StringPtrOutput)
}

// The custom version of the content. A optional free text
func (o ContentTemplateOutput) CustomVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringPtrOutput { return v.CustomVersion }).(pulumi.StringPtrOutput)
}

// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
func (o ContentTemplateOutput) Dependencies() MetadataDependenciesResponsePtrOutput {
	return o.ApplyT(func(v *ContentTemplate) MetadataDependenciesResponsePtrOutput { return v.Dependencies }).(MetadataDependenciesResponsePtrOutput)
}

// The display name of the template
func (o ContentTemplateOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o ContentTemplateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// first publish date content item
func (o ContentTemplateOutput) FirstPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringPtrOutput { return v.FirstPublishDate }).(pulumi.StringPtrOutput)
}

// the icon identifier. this id can later be fetched from the content metadata
func (o ContentTemplateOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringPtrOutput { return v.Icon }).(pulumi.StringPtrOutput)
}

// last publish date for the content item
func (o ContentTemplateOutput) LastPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringPtrOutput { return v.LastPublishDate }).(pulumi.StringPtrOutput)
}

// The JSON of the ARM template to deploy active content
func (o ContentTemplateOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.AnyOutput { return v.MainTemplate }).(pulumi.AnyOutput)
}

// The name of the resource
func (o ContentTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the package Id contains this template
func (o ContentTemplateOutput) PackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.PackageId }).(pulumi.StringOutput)
}

// the packageKind of the package contains this template
func (o ContentTemplateOutput) PackageKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringPtrOutput { return v.PackageKind }).(pulumi.StringPtrOutput)
}

// the name of the package contains this template
func (o ContentTemplateOutput) PackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringPtrOutput { return v.PackageName }).(pulumi.StringPtrOutput)
}

// Version of the package.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
func (o ContentTemplateOutput) PackageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.PackageVersion }).(pulumi.StringOutput)
}

// preview image file names. These will be taken from the solution artifacts
func (o ContentTemplateOutput) PreviewImages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringArrayOutput { return v.PreviewImages }).(pulumi.StringArrayOutput)
}

// preview image file names. These will be taken from the solution artifacts. used for dark theme support
func (o ContentTemplateOutput) PreviewImagesDark() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringArrayOutput { return v.PreviewImagesDark }).(pulumi.StringArrayOutput)
}

// Providers for the content item
func (o ContentTemplateOutput) Providers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringArrayOutput { return v.Providers }).(pulumi.StringArrayOutput)
}

// Source of the content.  This is where/how it was created.
func (o ContentTemplateOutput) Source() MetadataSourceResponseOutput {
	return o.ApplyT(func(v *ContentTemplate) MetadataSourceResponseOutput { return v.Source }).(MetadataSourceResponseOutput)
}

// Support information for the template - type, name, contact information
func (o ContentTemplateOutput) Support() MetadataSupportResponsePtrOutput {
	return o.ApplyT(func(v *ContentTemplate) MetadataSupportResponsePtrOutput { return v.Support }).(MetadataSupportResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ContentTemplateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContentTemplate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// the tactics the resource covers
func (o ContentTemplateOutput) ThreatAnalysisTactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringArrayOutput { return v.ThreatAnalysisTactics }).(pulumi.StringArrayOutput)
}

// the techniques the resource covers, these have to be aligned with the tactics being used
func (o ContentTemplateOutput) ThreatAnalysisTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringArrayOutput { return v.ThreatAnalysisTechniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ContentTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
func (o ContentTemplateOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentTemplate) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentTemplateOutput{})
}

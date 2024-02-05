// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Package in Azure Security Insights.
type ContentPackage struct {
	pulumi.CustomResourceState

	// The author of the package
	Author MetadataAuthorResponsePtrOutput `pulumi:"author"`
	// The categories of the package
	Categories MetadataCategoriesResponsePtrOutput `pulumi:"categories"`
	// The content id of the package
	ContentId pulumi.StringOutput `pulumi:"contentId"`
	// The package kind
	ContentKind pulumi.StringOutput `pulumi:"contentKind"`
	// Unique ID for the content. It should be generated based on the contentId, contentKind and the contentVersion of the package
	ContentProductId pulumi.StringOutput `pulumi:"contentProductId"`
	// The version of the content schema.
	ContentSchemaVersion pulumi.StringPtrOutput `pulumi:"contentSchemaVersion"`
	// The support tier of the package
	Dependencies MetadataDependenciesResponsePtrOutput `pulumi:"dependencies"`
	// The description of the package
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the package
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// first publish date package item
	FirstPublishDate pulumi.StringPtrOutput `pulumi:"firstPublishDate"`
	// the icon identifier. this id can later be fetched from the content metadata
	Icon pulumi.StringPtrOutput `pulumi:"icon"`
	// Flag indicates if this template is deprecated
	IsDeprecated pulumi.StringPtrOutput `pulumi:"isDeprecated"`
	// Flag indicates if this package is among the featured list.
	IsFeatured pulumi.StringPtrOutput `pulumi:"isFeatured"`
	// Flag indicates if this is a newly published package.
	IsNew pulumi.StringPtrOutput `pulumi:"isNew"`
	// Flag indicates if this package is in preview.
	IsPreview pulumi.StringPtrOutput `pulumi:"isPreview"`
	// last publish date for the package item
	LastPublishDate pulumi.StringPtrOutput `pulumi:"lastPublishDate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Providers for the package item
	Providers pulumi.StringArrayOutput `pulumi:"providers"`
	// The publisher display name of the package
	PublisherDisplayName pulumi.StringPtrOutput `pulumi:"publisherDisplayName"`
	// The source of the package
	Source MetadataSourceResponsePtrOutput `pulumi:"source"`
	// The support tier of the package
	Support MetadataSupportResponsePtrOutput `pulumi:"support"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// the tactics the resource covers
	ThreatAnalysisTactics pulumi.StringArrayOutput `pulumi:"threatAnalysisTactics"`
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques pulumi.StringArrayOutput `pulumi:"threatAnalysisTechniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// the latest version number of the package
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewContentPackage registers a new resource with the given unique name, arguments, and options.
func NewContentPackage(ctx *pulumi.Context,
	name string, args *ContentPackageArgs, opts ...pulumi.ResourceOption) (*ContentPackage, error) {
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
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ContentPackage"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:ContentPackage"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:ContentPackage"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:ContentPackage"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:ContentPackage"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:ContentPackage"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:ContentPackage"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:ContentPackage"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:ContentPackage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ContentPackage
	err := ctx.RegisterResource("azure-native:securityinsights/v20231101:ContentPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentPackage gets an existing ContentPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentPackageState, opts ...pulumi.ResourceOption) (*ContentPackage, error) {
	var resource ContentPackage
	err := ctx.ReadResource("azure-native:securityinsights/v20231101:ContentPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContentPackage resources.
type contentPackageState struct {
}

type ContentPackageState struct {
}

func (ContentPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentPackageState)(nil)).Elem()
}

type contentPackageArgs struct {
	// The author of the package
	Author *MetadataAuthor `pulumi:"author"`
	// The categories of the package
	Categories *MetadataCategories `pulumi:"categories"`
	// The content id of the package
	ContentId string `pulumi:"contentId"`
	// The package kind
	ContentKind string `pulumi:"contentKind"`
	// Unique ID for the content. It should be generated based on the contentId, contentKind and the contentVersion of the package
	ContentProductId string `pulumi:"contentProductId"`
	// The version of the content schema.
	ContentSchemaVersion *string `pulumi:"contentSchemaVersion"`
	// The support tier of the package
	Dependencies *MetadataDependencies `pulumi:"dependencies"`
	// The description of the package
	Description *string `pulumi:"description"`
	// The display name of the package
	DisplayName string `pulumi:"displayName"`
	// first publish date package item
	FirstPublishDate *string `pulumi:"firstPublishDate"`
	// the icon identifier. this id can later be fetched from the content metadata
	Icon *string `pulumi:"icon"`
	// Flag indicates if this template is deprecated
	IsDeprecated *string `pulumi:"isDeprecated"`
	// Flag indicates if this package is among the featured list.
	IsFeatured *string `pulumi:"isFeatured"`
	// Flag indicates if this is a newly published package.
	IsNew *string `pulumi:"isNew"`
	// Flag indicates if this package is in preview.
	IsPreview *string `pulumi:"isPreview"`
	// last publish date for the package item
	LastPublishDate *string `pulumi:"lastPublishDate"`
	// package Id
	PackageId *string `pulumi:"packageId"`
	// Providers for the package item
	Providers []string `pulumi:"providers"`
	// The publisher display name of the package
	PublisherDisplayName *string `pulumi:"publisherDisplayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The source of the package
	Source *MetadataSource `pulumi:"source"`
	// The support tier of the package
	Support *MetadataSupport `pulumi:"support"`
	// the tactics the resource covers
	ThreatAnalysisTactics []string `pulumi:"threatAnalysisTactics"`
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques []string `pulumi:"threatAnalysisTechniques"`
	// the latest version number of the package
	Version string `pulumi:"version"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ContentPackage resource.
type ContentPackageArgs struct {
	// The author of the package
	Author MetadataAuthorPtrInput
	// The categories of the package
	Categories MetadataCategoriesPtrInput
	// The content id of the package
	ContentId pulumi.StringInput
	// The package kind
	ContentKind pulumi.StringInput
	// Unique ID for the content. It should be generated based on the contentId, contentKind and the contentVersion of the package
	ContentProductId pulumi.StringInput
	// The version of the content schema.
	ContentSchemaVersion pulumi.StringPtrInput
	// The support tier of the package
	Dependencies MetadataDependenciesPtrInput
	// The description of the package
	Description pulumi.StringPtrInput
	// The display name of the package
	DisplayName pulumi.StringInput
	// first publish date package item
	FirstPublishDate pulumi.StringPtrInput
	// the icon identifier. this id can later be fetched from the content metadata
	Icon pulumi.StringPtrInput
	// Flag indicates if this template is deprecated
	IsDeprecated pulumi.StringPtrInput
	// Flag indicates if this package is among the featured list.
	IsFeatured pulumi.StringPtrInput
	// Flag indicates if this is a newly published package.
	IsNew pulumi.StringPtrInput
	// Flag indicates if this package is in preview.
	IsPreview pulumi.StringPtrInput
	// last publish date for the package item
	LastPublishDate pulumi.StringPtrInput
	// package Id
	PackageId pulumi.StringPtrInput
	// Providers for the package item
	Providers pulumi.StringArrayInput
	// The publisher display name of the package
	PublisherDisplayName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The source of the package
	Source MetadataSourcePtrInput
	// The support tier of the package
	Support MetadataSupportPtrInput
	// the tactics the resource covers
	ThreatAnalysisTactics pulumi.StringArrayInput
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques pulumi.StringArrayInput
	// the latest version number of the package
	Version pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ContentPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentPackageArgs)(nil)).Elem()
}

type ContentPackageInput interface {
	pulumi.Input

	ToContentPackageOutput() ContentPackageOutput
	ToContentPackageOutputWithContext(ctx context.Context) ContentPackageOutput
}

func (*ContentPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentPackage)(nil)).Elem()
}

func (i *ContentPackage) ToContentPackageOutput() ContentPackageOutput {
	return i.ToContentPackageOutputWithContext(context.Background())
}

func (i *ContentPackage) ToContentPackageOutputWithContext(ctx context.Context) ContentPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentPackageOutput)
}

type ContentPackageOutput struct{ *pulumi.OutputState }

func (ContentPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentPackage)(nil)).Elem()
}

func (o ContentPackageOutput) ToContentPackageOutput() ContentPackageOutput {
	return o
}

func (o ContentPackageOutput) ToContentPackageOutputWithContext(ctx context.Context) ContentPackageOutput {
	return o
}

// The author of the package
func (o ContentPackageOutput) Author() MetadataAuthorResponsePtrOutput {
	return o.ApplyT(func(v *ContentPackage) MetadataAuthorResponsePtrOutput { return v.Author }).(MetadataAuthorResponsePtrOutput)
}

// The categories of the package
func (o ContentPackageOutput) Categories() MetadataCategoriesResponsePtrOutput {
	return o.ApplyT(func(v *ContentPackage) MetadataCategoriesResponsePtrOutput { return v.Categories }).(MetadataCategoriesResponsePtrOutput)
}

// The content id of the package
func (o ContentPackageOutput) ContentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringOutput { return v.ContentId }).(pulumi.StringOutput)
}

// The package kind
func (o ContentPackageOutput) ContentKind() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringOutput { return v.ContentKind }).(pulumi.StringOutput)
}

// Unique ID for the content. It should be generated based on the contentId, contentKind and the contentVersion of the package
func (o ContentPackageOutput) ContentProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringOutput { return v.ContentProductId }).(pulumi.StringOutput)
}

// The version of the content schema.
func (o ContentPackageOutput) ContentSchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.ContentSchemaVersion }).(pulumi.StringPtrOutput)
}

// The support tier of the package
func (o ContentPackageOutput) Dependencies() MetadataDependenciesResponsePtrOutput {
	return o.ApplyT(func(v *ContentPackage) MetadataDependenciesResponsePtrOutput { return v.Dependencies }).(MetadataDependenciesResponsePtrOutput)
}

// The description of the package
func (o ContentPackageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the package
func (o ContentPackageOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o ContentPackageOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// first publish date package item
func (o ContentPackageOutput) FirstPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.FirstPublishDate }).(pulumi.StringPtrOutput)
}

// the icon identifier. this id can later be fetched from the content metadata
func (o ContentPackageOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.Icon }).(pulumi.StringPtrOutput)
}

// Flag indicates if this template is deprecated
func (o ContentPackageOutput) IsDeprecated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.IsDeprecated }).(pulumi.StringPtrOutput)
}

// Flag indicates if this package is among the featured list.
func (o ContentPackageOutput) IsFeatured() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.IsFeatured }).(pulumi.StringPtrOutput)
}

// Flag indicates if this is a newly published package.
func (o ContentPackageOutput) IsNew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.IsNew }).(pulumi.StringPtrOutput)
}

// Flag indicates if this package is in preview.
func (o ContentPackageOutput) IsPreview() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.IsPreview }).(pulumi.StringPtrOutput)
}

// last publish date for the package item
func (o ContentPackageOutput) LastPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.LastPublishDate }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ContentPackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Providers for the package item
func (o ContentPackageOutput) Providers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringArrayOutput { return v.Providers }).(pulumi.StringArrayOutput)
}

// The publisher display name of the package
func (o ContentPackageOutput) PublisherDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringPtrOutput { return v.PublisherDisplayName }).(pulumi.StringPtrOutput)
}

// The source of the package
func (o ContentPackageOutput) Source() MetadataSourceResponsePtrOutput {
	return o.ApplyT(func(v *ContentPackage) MetadataSourceResponsePtrOutput { return v.Source }).(MetadataSourceResponsePtrOutput)
}

// The support tier of the package
func (o ContentPackageOutput) Support() MetadataSupportResponsePtrOutput {
	return o.ApplyT(func(v *ContentPackage) MetadataSupportResponsePtrOutput { return v.Support }).(MetadataSupportResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ContentPackageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContentPackage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// the tactics the resource covers
func (o ContentPackageOutput) ThreatAnalysisTactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringArrayOutput { return v.ThreatAnalysisTactics }).(pulumi.StringArrayOutput)
}

// the techniques the resource covers, these have to be aligned with the tactics being used
func (o ContentPackageOutput) ThreatAnalysisTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringArrayOutput { return v.ThreatAnalysisTechniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ContentPackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// the latest version number of the package
func (o ContentPackageOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentPackage) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentPackageOutput{})
}

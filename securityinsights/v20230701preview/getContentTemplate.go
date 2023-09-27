// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a template byt its identifier.
func LookupContentTemplate(ctx *pulumi.Context, args *LookupContentTemplateArgs, opts ...pulumi.InvokeOption) (*LookupContentTemplateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupContentTemplateResult
	err := ctx.Invoke("azure-native:securityinsights/v20230701preview:getContentTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentTemplateArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// template Id
	TemplateId string `pulumi:"templateId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Template resource definition.
type LookupContentTemplateResult struct {
	// The creator of the content item.
	Author *MetadataAuthorResponse `pulumi:"author"`
	// Categories for the item
	Categories *MetadataCategoriesResponse `pulumi:"categories"`
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
	Dependencies *MetadataDependenciesResponse `pulumi:"dependencies"`
	// The display name of the template
	DisplayName string `pulumi:"displayName"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// first publish date content item
	FirstPublishDate *string `pulumi:"firstPublishDate"`
	// the icon identifier. this id can later be fetched from the content metadata
	Icon *string `pulumi:"icon"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// last publish date for the content item
	LastPublishDate *string `pulumi:"lastPublishDate"`
	// The JSON of the ARM template to deploy active content
	MainTemplate interface{} `pulumi:"mainTemplate"`
	// The name of the resource
	Name string `pulumi:"name"`
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
	// Source of the content.  This is where/how it was created.
	Source MetadataSourceResponse `pulumi:"source"`
	// Support information for the template - type, name, contact information
	Support *MetadataSupportResponse `pulumi:"support"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// the tactics the resource covers
	ThreatAnalysisTactics []string `pulumi:"threatAnalysisTactics"`
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques []string `pulumi:"threatAnalysisTechniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
	Version string `pulumi:"version"`
}

func LookupContentTemplateOutput(ctx *pulumi.Context, args LookupContentTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupContentTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentTemplateResult, error) {
			args := v.(LookupContentTemplateArgs)
			r, err := LookupContentTemplate(ctx, &args, opts...)
			var s LookupContentTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentTemplateResultOutput)
}

type LookupContentTemplateOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// template Id
	TemplateId pulumi.StringInput `pulumi:"templateId"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupContentTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentTemplateArgs)(nil)).Elem()
}

// Template resource definition.
type LookupContentTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupContentTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentTemplateResult)(nil)).Elem()
}

func (o LookupContentTemplateResultOutput) ToLookupContentTemplateResultOutput() LookupContentTemplateResultOutput {
	return o
}

func (o LookupContentTemplateResultOutput) ToLookupContentTemplateResultOutputWithContext(ctx context.Context) LookupContentTemplateResultOutput {
	return o
}

func (o LookupContentTemplateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupContentTemplateResult] {
	return pulumix.Output[LookupContentTemplateResult]{
		OutputState: o.OutputState,
	}
}

// The creator of the content item.
func (o LookupContentTemplateResultOutput) Author() MetadataAuthorResponsePtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *MetadataAuthorResponse { return v.Author }).(MetadataAuthorResponsePtrOutput)
}

// Categories for the item
func (o LookupContentTemplateResultOutput) Categories() MetadataCategoriesResponsePtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *MetadataCategoriesResponse { return v.Categories }).(MetadataCategoriesResponsePtrOutput)
}

// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
func (o LookupContentTemplateResultOutput) ContentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.ContentId }).(pulumi.StringOutput)
}

// The kind of content the template is for.
func (o LookupContentTemplateResultOutput) ContentKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.ContentKind }).(pulumi.StringOutput)
}

// Unique ID for the content. It should be generated based on the contentId of the package, contentId of the template, contentKind of the template and the contentVersion of the template
func (o LookupContentTemplateResultOutput) ContentProductId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.ContentProductId }).(pulumi.StringOutput)
}

// Schema version of the content. Can be used to distinguish between different flow based on the schema version
func (o LookupContentTemplateResultOutput) ContentSchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *string { return v.ContentSchemaVersion }).(pulumi.StringPtrOutput)
}

// The custom version of the content. A optional free text
func (o LookupContentTemplateResultOutput) CustomVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *string { return v.CustomVersion }).(pulumi.StringPtrOutput)
}

// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
func (o LookupContentTemplateResultOutput) Dependencies() MetadataDependenciesResponsePtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *MetadataDependenciesResponse { return v.Dependencies }).(MetadataDependenciesResponsePtrOutput)
}

// The display name of the template
func (o LookupContentTemplateResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o LookupContentTemplateResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// first publish date content item
func (o LookupContentTemplateResultOutput) FirstPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *string { return v.FirstPublishDate }).(pulumi.StringPtrOutput)
}

// the icon identifier. this id can later be fetched from the content metadata
func (o LookupContentTemplateResultOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *string { return v.Icon }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupContentTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// last publish date for the content item
func (o LookupContentTemplateResultOutput) LastPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *string { return v.LastPublishDate }).(pulumi.StringPtrOutput)
}

// The JSON of the ARM template to deploy active content
func (o LookupContentTemplateResultOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) interface{} { return v.MainTemplate }).(pulumi.AnyOutput)
}

// The name of the resource
func (o LookupContentTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// the package Id contains this template
func (o LookupContentTemplateResultOutput) PackageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.PackageId }).(pulumi.StringOutput)
}

// the packageKind of the package contains this template
func (o LookupContentTemplateResultOutput) PackageKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *string { return v.PackageKind }).(pulumi.StringPtrOutput)
}

// the name of the package contains this template
func (o LookupContentTemplateResultOutput) PackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *string { return v.PackageName }).(pulumi.StringPtrOutput)
}

// Version of the package.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
func (o LookupContentTemplateResultOutput) PackageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.PackageVersion }).(pulumi.StringOutput)
}

// preview image file names. These will be taken from the solution artifacts
func (o LookupContentTemplateResultOutput) PreviewImages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) []string { return v.PreviewImages }).(pulumi.StringArrayOutput)
}

// preview image file names. These will be taken from the solution artifacts. used for dark theme support
func (o LookupContentTemplateResultOutput) PreviewImagesDark() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) []string { return v.PreviewImagesDark }).(pulumi.StringArrayOutput)
}

// Providers for the content item
func (o LookupContentTemplateResultOutput) Providers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) []string { return v.Providers }).(pulumi.StringArrayOutput)
}

// Source of the content.  This is where/how it was created.
func (o LookupContentTemplateResultOutput) Source() MetadataSourceResponseOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) MetadataSourceResponse { return v.Source }).(MetadataSourceResponseOutput)
}

// Support information for the template - type, name, contact information
func (o LookupContentTemplateResultOutput) Support() MetadataSupportResponsePtrOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) *MetadataSupportResponse { return v.Support }).(MetadataSupportResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupContentTemplateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// the tactics the resource covers
func (o LookupContentTemplateResultOutput) ThreatAnalysisTactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) []string { return v.ThreatAnalysisTactics }).(pulumi.StringArrayOutput)
}

// the techniques the resource covers, these have to be aligned with the tactics being used
func (o LookupContentTemplateResultOutput) ThreatAnalysisTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) []string { return v.ThreatAnalysisTechniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupContentTemplateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM metadata best practices.  Can also be any string, but then we cannot guarantee any version checks
func (o LookupContentTemplateResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTemplateResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentTemplateResultOutput{})
}

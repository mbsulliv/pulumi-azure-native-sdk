// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Metadata.
func LookupMetadata(ctx *pulumi.Context, args *LookupMetadataArgs, opts ...pulumi.InvokeOption) (*LookupMetadataResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMetadataResult
	err := ctx.Invoke("azure-native:securityinsights/v20230801preview:getMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetadataArgs struct {
	// The Metadata name.
	MetadataName string `pulumi:"metadataName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Metadata resource definition.
type LookupMetadataResult struct {
	// The creator of the content item.
	Author *MetadataAuthorResponse `pulumi:"author"`
	// Categories for the solution content item
	Categories *MetadataCategoriesResponse `pulumi:"categories"`
	// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
	ContentId *string `pulumi:"contentId"`
	// Schema version of the content. Can be used to distinguish between different flow based on the schema version
	ContentSchemaVersion *string `pulumi:"contentSchemaVersion"`
	// The custom version of the content. A optional free text
	CustomVersion *string `pulumi:"customVersion"`
	// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
	Dependencies *MetadataDependenciesResponse `pulumi:"dependencies"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// first publish date solution content item
	FirstPublishDate *string `pulumi:"firstPublishDate"`
	// the icon identifier. this id can later be fetched from the solution template
	Icon *string `pulumi:"icon"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of content the metadata is for.
	Kind string `pulumi:"kind"`
	// last publish date for the solution content item
	LastPublishDate *string `pulumi:"lastPublishDate"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Full parent resource ID of the content item the metadata is for.  This is the full resource ID including the scope (subscription and resource group)
	ParentId string `pulumi:"parentId"`
	// preview image file names. These will be taken from the solution artifacts
	PreviewImages []string `pulumi:"previewImages"`
	// preview image file names. These will be taken from the solution artifacts. used for dark theme support
	PreviewImagesDark []string `pulumi:"previewImagesDark"`
	// Providers for the solution content item
	Providers []string `pulumi:"providers"`
	// Source of the content.  This is where/how it was created.
	Source *MetadataSourceResponse `pulumi:"source"`
	// Support information for the metadata - type, name, contact information
	Support *MetadataSupportResponse `pulumi:"support"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// the tactics the resource covers
	ThreatAnalysisTactics []string `pulumi:"threatAnalysisTactics"`
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques []string `pulumi:"threatAnalysisTechniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM template best practices.  Can also be any string, but then we cannot guarantee any version checks
	Version *string `pulumi:"version"`
}

func LookupMetadataOutput(ctx *pulumi.Context, args LookupMetadataOutputArgs, opts ...pulumi.InvokeOption) LookupMetadataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetadataResult, error) {
			args := v.(LookupMetadataArgs)
			r, err := LookupMetadata(ctx, &args, opts...)
			var s LookupMetadataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetadataResultOutput)
}

type LookupMetadataOutputArgs struct {
	// The Metadata name.
	MetadataName pulumi.StringInput `pulumi:"metadataName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMetadataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetadataArgs)(nil)).Elem()
}

// Metadata resource definition.
type LookupMetadataResultOutput struct{ *pulumi.OutputState }

func (LookupMetadataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetadataResult)(nil)).Elem()
}

func (o LookupMetadataResultOutput) ToLookupMetadataResultOutput() LookupMetadataResultOutput {
	return o
}

func (o LookupMetadataResultOutput) ToLookupMetadataResultOutputWithContext(ctx context.Context) LookupMetadataResultOutput {
	return o
}

func (o LookupMetadataResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMetadataResult] {
	return pulumix.Output[LookupMetadataResult]{
		OutputState: o.OutputState,
	}
}

// The creator of the content item.
func (o LookupMetadataResultOutput) Author() MetadataAuthorResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataAuthorResponse { return v.Author }).(MetadataAuthorResponsePtrOutput)
}

// Categories for the solution content item
func (o LookupMetadataResultOutput) Categories() MetadataCategoriesResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataCategoriesResponse { return v.Categories }).(MetadataCategoriesResponsePtrOutput)
}

// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
func (o LookupMetadataResultOutput) ContentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.ContentId }).(pulumi.StringPtrOutput)
}

// Schema version of the content. Can be used to distinguish between different flow based on the schema version
func (o LookupMetadataResultOutput) ContentSchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.ContentSchemaVersion }).(pulumi.StringPtrOutput)
}

// The custom version of the content. A optional free text
func (o LookupMetadataResultOutput) CustomVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.CustomVersion }).(pulumi.StringPtrOutput)
}

// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
func (o LookupMetadataResultOutput) Dependencies() MetadataDependenciesResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataDependenciesResponse { return v.Dependencies }).(MetadataDependenciesResponsePtrOutput)
}

// Etag of the azure resource
func (o LookupMetadataResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// first publish date solution content item
func (o LookupMetadataResultOutput) FirstPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.FirstPublishDate }).(pulumi.StringPtrOutput)
}

// the icon identifier. this id can later be fetched from the solution template
func (o LookupMetadataResultOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.Icon }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupMetadataResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of content the metadata is for.
func (o LookupMetadataResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.Kind }).(pulumi.StringOutput)
}

// last publish date for the solution content item
func (o LookupMetadataResultOutput) LastPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.LastPublishDate }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupMetadataResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.Name }).(pulumi.StringOutput)
}

// Full parent resource ID of the content item the metadata is for.  This is the full resource ID including the scope (subscription and resource group)
func (o LookupMetadataResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.ParentId }).(pulumi.StringOutput)
}

// preview image file names. These will be taken from the solution artifacts
func (o LookupMetadataResultOutput) PreviewImages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetadataResult) []string { return v.PreviewImages }).(pulumi.StringArrayOutput)
}

// preview image file names. These will be taken from the solution artifacts. used for dark theme support
func (o LookupMetadataResultOutput) PreviewImagesDark() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetadataResult) []string { return v.PreviewImagesDark }).(pulumi.StringArrayOutput)
}

// Providers for the solution content item
func (o LookupMetadataResultOutput) Providers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetadataResult) []string { return v.Providers }).(pulumi.StringArrayOutput)
}

// Source of the content.  This is where/how it was created.
func (o LookupMetadataResultOutput) Source() MetadataSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataSourceResponse { return v.Source }).(MetadataSourceResponsePtrOutput)
}

// Support information for the metadata - type, name, contact information
func (o LookupMetadataResultOutput) Support() MetadataSupportResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataSupportResponse { return v.Support }).(MetadataSupportResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMetadataResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMetadataResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// the tactics the resource covers
func (o LookupMetadataResultOutput) ThreatAnalysisTactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetadataResult) []string { return v.ThreatAnalysisTactics }).(pulumi.StringArrayOutput)
}

// the techniques the resource covers, these have to be aligned with the tactics being used
func (o LookupMetadataResultOutput) ThreatAnalysisTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetadataResult) []string { return v.ThreatAnalysisTechniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMetadataResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM template best practices.  Can also be any string, but then we cannot guarantee any version checks
func (o LookupMetadataResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetadataResultOutput{})
}

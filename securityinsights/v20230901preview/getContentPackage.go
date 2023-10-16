// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an installed packages by its id.
func LookupContentPackage(ctx *pulumi.Context, args *LookupContentPackageArgs, opts ...pulumi.InvokeOption) (*LookupContentPackageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupContentPackageResult
	err := ctx.Invoke("azure-native:securityinsights/v20230901preview:getContentPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentPackageArgs struct {
	// package Id
	PackageId string `pulumi:"packageId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a Package in Azure Security Insights.
type LookupContentPackageResult struct {
	// The author of the package
	Author *MetadataAuthorResponse `pulumi:"author"`
	// The categories of the package
	Categories *MetadataCategoriesResponse `pulumi:"categories"`
	// The content id of the package
	ContentId string `pulumi:"contentId"`
	// The package kind
	ContentKind string `pulumi:"contentKind"`
	// Unique ID for the content. It should be generated based on the contentId, contentKind and the contentVersion of the package
	ContentProductId string `pulumi:"contentProductId"`
	// The version of the content schema.
	ContentSchemaVersion *string `pulumi:"contentSchemaVersion"`
	// The support tier of the package
	Dependencies *MetadataDependenciesResponse `pulumi:"dependencies"`
	// The description of the package
	Description *string `pulumi:"description"`
	// The display name of the package
	DisplayName string `pulumi:"displayName"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// first publish date package item
	FirstPublishDate *string `pulumi:"firstPublishDate"`
	// the icon identifier. this id can later be fetched from the content metadata
	Icon *string `pulumi:"icon"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Flag indicates if this package is among the featured list.
	IsFeatured *string `pulumi:"isFeatured"`
	// Flag indicates if this is a newly published package.
	IsNew *string `pulumi:"isNew"`
	// Flag indicates if this package is in preview.
	IsPreview *string `pulumi:"isPreview"`
	// last publish date for the package item
	LastPublishDate *string `pulumi:"lastPublishDate"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Providers for the package item
	Providers []string `pulumi:"providers"`
	// The publisher display name of the package
	PublisherDisplayName *string `pulumi:"publisherDisplayName"`
	// The source of the package
	Source *MetadataSourceResponse `pulumi:"source"`
	// The support tier of the package
	Support *MetadataSupportResponse `pulumi:"support"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// the tactics the resource covers
	ThreatAnalysisTactics []string `pulumi:"threatAnalysisTactics"`
	// the techniques the resource covers, these have to be aligned with the tactics being used
	ThreatAnalysisTechniques []string `pulumi:"threatAnalysisTechniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// the latest version number of the package
	Version string `pulumi:"version"`
}

func LookupContentPackageOutput(ctx *pulumi.Context, args LookupContentPackageOutputArgs, opts ...pulumi.InvokeOption) LookupContentPackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentPackageResult, error) {
			args := v.(LookupContentPackageArgs)
			r, err := LookupContentPackage(ctx, &args, opts...)
			var s LookupContentPackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentPackageResultOutput)
}

type LookupContentPackageOutputArgs struct {
	// package Id
	PackageId pulumi.StringInput `pulumi:"packageId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupContentPackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentPackageArgs)(nil)).Elem()
}

// Represents a Package in Azure Security Insights.
type LookupContentPackageResultOutput struct{ *pulumi.OutputState }

func (LookupContentPackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentPackageResult)(nil)).Elem()
}

func (o LookupContentPackageResultOutput) ToLookupContentPackageResultOutput() LookupContentPackageResultOutput {
	return o
}

func (o LookupContentPackageResultOutput) ToLookupContentPackageResultOutputWithContext(ctx context.Context) LookupContentPackageResultOutput {
	return o
}

func (o LookupContentPackageResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupContentPackageResult] {
	return pulumix.Output[LookupContentPackageResult]{
		OutputState: o.OutputState,
	}
}

// The author of the package
func (o LookupContentPackageResultOutput) Author() MetadataAuthorResponsePtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *MetadataAuthorResponse { return v.Author }).(MetadataAuthorResponsePtrOutput)
}

// The categories of the package
func (o LookupContentPackageResultOutput) Categories() MetadataCategoriesResponsePtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *MetadataCategoriesResponse { return v.Categories }).(MetadataCategoriesResponsePtrOutput)
}

// The content id of the package
func (o LookupContentPackageResultOutput) ContentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentPackageResult) string { return v.ContentId }).(pulumi.StringOutput)
}

// The package kind
func (o LookupContentPackageResultOutput) ContentKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentPackageResult) string { return v.ContentKind }).(pulumi.StringOutput)
}

// Unique ID for the content. It should be generated based on the contentId, contentKind and the contentVersion of the package
func (o LookupContentPackageResultOutput) ContentProductId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentPackageResult) string { return v.ContentProductId }).(pulumi.StringOutput)
}

// The version of the content schema.
func (o LookupContentPackageResultOutput) ContentSchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.ContentSchemaVersion }).(pulumi.StringPtrOutput)
}

// The support tier of the package
func (o LookupContentPackageResultOutput) Dependencies() MetadataDependenciesResponsePtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *MetadataDependenciesResponse { return v.Dependencies }).(MetadataDependenciesResponsePtrOutput)
}

// The description of the package
func (o LookupContentPackageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the package
func (o LookupContentPackageResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentPackageResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o LookupContentPackageResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// first publish date package item
func (o LookupContentPackageResultOutput) FirstPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.FirstPublishDate }).(pulumi.StringPtrOutput)
}

// the icon identifier. this id can later be fetched from the content metadata
func (o LookupContentPackageResultOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.Icon }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupContentPackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentPackageResult) string { return v.Id }).(pulumi.StringOutput)
}

// Flag indicates if this package is among the featured list.
func (o LookupContentPackageResultOutput) IsFeatured() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.IsFeatured }).(pulumi.StringPtrOutput)
}

// Flag indicates if this is a newly published package.
func (o LookupContentPackageResultOutput) IsNew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.IsNew }).(pulumi.StringPtrOutput)
}

// Flag indicates if this package is in preview.
func (o LookupContentPackageResultOutput) IsPreview() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.IsPreview }).(pulumi.StringPtrOutput)
}

// last publish date for the package item
func (o LookupContentPackageResultOutput) LastPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.LastPublishDate }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupContentPackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentPackageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Providers for the package item
func (o LookupContentPackageResultOutput) Providers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContentPackageResult) []string { return v.Providers }).(pulumi.StringArrayOutput)
}

// The publisher display name of the package
func (o LookupContentPackageResultOutput) PublisherDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *string { return v.PublisherDisplayName }).(pulumi.StringPtrOutput)
}

// The source of the package
func (o LookupContentPackageResultOutput) Source() MetadataSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *MetadataSourceResponse { return v.Source }).(MetadataSourceResponsePtrOutput)
}

// The support tier of the package
func (o LookupContentPackageResultOutput) Support() MetadataSupportResponsePtrOutput {
	return o.ApplyT(func(v LookupContentPackageResult) *MetadataSupportResponse { return v.Support }).(MetadataSupportResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupContentPackageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContentPackageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// the tactics the resource covers
func (o LookupContentPackageResultOutput) ThreatAnalysisTactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContentPackageResult) []string { return v.ThreatAnalysisTactics }).(pulumi.StringArrayOutput)
}

// the techniques the resource covers, these have to be aligned with the tactics being used
func (o LookupContentPackageResultOutput) ThreatAnalysisTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContentPackageResult) []string { return v.ThreatAnalysisTechniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupContentPackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentPackageResult) string { return v.Type }).(pulumi.StringOutput)
}

// the latest version number of the package
func (o LookupContentPackageResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentPackageResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentPackageResultOutput{})
}
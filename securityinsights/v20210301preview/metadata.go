// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Metadata resource definition.
type Metadata struct {
	pulumi.CustomResourceState

	// The creator of the content item.
	Author MetadataAuthorResponsePtrOutput `pulumi:"author"`
	// Categories for the solution content item
	Categories MetadataCategoriesResponsePtrOutput `pulumi:"categories"`
	// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
	ContentId pulumi.StringPtrOutput `pulumi:"contentId"`
	// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
	Dependencies MetadataDependenciesResponsePtrOutput `pulumi:"dependencies"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// first publish date solution content item
	FirstPublishDate pulumi.StringPtrOutput `pulumi:"firstPublishDate"`
	// The kind of content the metadata is for.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// last publish date for the solution content item
	LastPublishDate pulumi.StringPtrOutput `pulumi:"lastPublishDate"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Full parent resource ID of the content item the metadata is for.  This is the full resource ID including the scope (subscription and resource group)
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// Providers for the solution content item
	Providers pulumi.StringArrayOutput `pulumi:"providers"`
	// Source of the content.  This is where/how it was created.
	Source MetadataSourceResponsePtrOutput `pulumi:"source"`
	// Support information for the metadata - type, name, contact information
	Support MetadataSupportResponsePtrOutput `pulumi:"support"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM template best practices.  Can also be any string, but then we cannot guarantee any version checks
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewMetadata registers a new resource with the given unique name, arguments, and options.
func NewMetadata(ctx *pulumi.Context,
	name string, args *MetadataArgs, opts ...pulumi.ResourceOption) (*Metadata, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:Metadata"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Metadata
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:Metadata", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetadata gets an existing Metadata resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetadata(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetadataState, opts ...pulumi.ResourceOption) (*Metadata, error) {
	var resource Metadata
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:Metadata", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Metadata resources.
type metadataState struct {
}

type MetadataState struct {
}

func (MetadataState) ElementType() reflect.Type {
	return reflect.TypeOf((*metadataState)(nil)).Elem()
}

type metadataArgs struct {
	// The creator of the content item.
	Author *MetadataAuthor `pulumi:"author"`
	// Categories for the solution content item
	Categories *MetadataCategories `pulumi:"categories"`
	// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
	ContentId *string `pulumi:"contentId"`
	// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
	Dependencies *MetadataDependencies `pulumi:"dependencies"`
	// first publish date solution content item
	FirstPublishDate *string `pulumi:"firstPublishDate"`
	// The kind of content the metadata is for.
	Kind string `pulumi:"kind"`
	// last publish date for the solution content item
	LastPublishDate *string `pulumi:"lastPublishDate"`
	// The Metadata name.
	MetadataName *string `pulumi:"metadataName"`
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	// Full parent resource ID of the content item the metadata is for.  This is the full resource ID including the scope (subscription and resource group)
	ParentId string `pulumi:"parentId"`
	// Providers for the solution content item
	Providers []string `pulumi:"providers"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Source of the content.  This is where/how it was created.
	Source *MetadataSource `pulumi:"source"`
	// Support information for the metadata - type, name, contact information
	Support *MetadataSupport `pulumi:"support"`
	// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM template best practices.  Can also be any string, but then we cannot guarantee any version checks
	Version *string `pulumi:"version"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Metadata resource.
type MetadataArgs struct {
	// The creator of the content item.
	Author MetadataAuthorPtrInput
	// Categories for the solution content item
	Categories MetadataCategoriesPtrInput
	// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
	ContentId pulumi.StringPtrInput
	// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
	Dependencies MetadataDependenciesPtrInput
	// first publish date solution content item
	FirstPublishDate pulumi.StringPtrInput
	// The kind of content the metadata is for.
	Kind pulumi.StringInput
	// last publish date for the solution content item
	LastPublishDate pulumi.StringPtrInput
	// The Metadata name.
	MetadataName pulumi.StringPtrInput
	// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
	OperationalInsightsResourceProvider pulumi.StringInput
	// Full parent resource ID of the content item the metadata is for.  This is the full resource ID including the scope (subscription and resource group)
	ParentId pulumi.StringInput
	// Providers for the solution content item
	Providers pulumi.StringArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Source of the content.  This is where/how it was created.
	Source MetadataSourcePtrInput
	// Support information for the metadata - type, name, contact information
	Support MetadataSupportPtrInput
	// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM template best practices.  Can also be any string, but then we cannot guarantee any version checks
	Version pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metadataArgs)(nil)).Elem()
}

type MetadataInput interface {
	pulumi.Input

	ToMetadataOutput() MetadataOutput
	ToMetadataOutputWithContext(ctx context.Context) MetadataOutput
}

func (*Metadata) ElementType() reflect.Type {
	return reflect.TypeOf((**Metadata)(nil)).Elem()
}

func (i *Metadata) ToMetadataOutput() MetadataOutput {
	return i.ToMetadataOutputWithContext(context.Background())
}

func (i *Metadata) ToMetadataOutputWithContext(ctx context.Context) MetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataOutput)
}

type MetadataOutput struct{ *pulumi.OutputState }

func (MetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Metadata)(nil)).Elem()
}

func (o MetadataOutput) ToMetadataOutput() MetadataOutput {
	return o
}

func (o MetadataOutput) ToMetadataOutputWithContext(ctx context.Context) MetadataOutput {
	return o
}

// The creator of the content item.
func (o MetadataOutput) Author() MetadataAuthorResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataAuthorResponsePtrOutput { return v.Author }).(MetadataAuthorResponsePtrOutput)
}

// Categories for the solution content item
func (o MetadataOutput) Categories() MetadataCategoriesResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataCategoriesResponsePtrOutput { return v.Categories }).(MetadataCategoriesResponsePtrOutput)
}

// Static ID for the content.  Used to identify dependencies and content from solutions or community.  Hard-coded/static for out of the box content and solutions. Dynamic for user-created.  This is the resource name
func (o MetadataOutput) ContentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.ContentId }).(pulumi.StringPtrOutput)
}

// Dependencies for the content item, what other content items it requires to work.  Can describe more complex dependencies using a recursive/nested structure. For a single dependency an id/kind/version can be supplied or operator/criteria for complex formats.
func (o MetadataOutput) Dependencies() MetadataDependenciesResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataDependenciesResponsePtrOutput { return v.Dependencies }).(MetadataDependenciesResponsePtrOutput)
}

// Etag of the azure resource
func (o MetadataOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// first publish date solution content item
func (o MetadataOutput) FirstPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.FirstPublishDate }).(pulumi.StringPtrOutput)
}

// The kind of content the metadata is for.
func (o MetadataOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// last publish date for the solution content item
func (o MetadataOutput) LastPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.LastPublishDate }).(pulumi.StringPtrOutput)
}

// Azure resource name
func (o MetadataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Full parent resource ID of the content item the metadata is for.  This is the full resource ID including the scope (subscription and resource group)
func (o MetadataOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

// Providers for the solution content item
func (o MetadataOutput) Providers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringArrayOutput { return v.Providers }).(pulumi.StringArrayOutput)
}

// Source of the content.  This is where/how it was created.
func (o MetadataOutput) Source() MetadataSourceResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataSourceResponsePtrOutput { return v.Source }).(MetadataSourceResponsePtrOutput)
}

// Support information for the metadata - type, name, contact information
func (o MetadataOutput) Support() MetadataSupportResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataSupportResponsePtrOutput { return v.Support }).(MetadataSupportResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MetadataOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Metadata) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource type
func (o MetadataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the content.  Default and recommended format is numeric (e.g. 1, 1.0, 1.0.0, 1.0.0.0), following ARM template best practices.  Can also be any string, but then we cannot guarantee any version checks
func (o MetadataOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MetadataOutput{})
}

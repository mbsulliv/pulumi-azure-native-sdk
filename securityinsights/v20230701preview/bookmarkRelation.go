// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a relation between two resources
type BookmarkRelation struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the related resource
	RelatedResourceId pulumi.StringOutput `pulumi:"relatedResourceId"`
	// The resource kind of the related resource
	RelatedResourceKind pulumi.StringOutput `pulumi:"relatedResourceKind"`
	// The name of the related resource
	RelatedResourceName pulumi.StringOutput `pulumi:"relatedResourceName"`
	// The resource type of the related resource
	RelatedResourceType pulumi.StringOutput `pulumi:"relatedResourceType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBookmarkRelation registers a new resource with the given unique name, arguments, and options.
func NewBookmarkRelation(ctx *pulumi.Context,
	name string, args *BookmarkRelationArgs, opts ...pulumi.ResourceOption) (*BookmarkRelation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BookmarkId == nil {
		return nil, errors.New("invalid value for required argument 'BookmarkId'")
	}
	if args.RelatedResourceId == nil {
		return nil, errors.New("invalid value for required argument 'RelatedResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:BookmarkRelation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BookmarkRelation
	err := ctx.RegisterResource("azure-native:securityinsights/v20230701preview:BookmarkRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBookmarkRelation gets an existing BookmarkRelation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBookmarkRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BookmarkRelationState, opts ...pulumi.ResourceOption) (*BookmarkRelation, error) {
	var resource BookmarkRelation
	err := ctx.ReadResource("azure-native:securityinsights/v20230701preview:BookmarkRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BookmarkRelation resources.
type bookmarkRelationState struct {
}

type BookmarkRelationState struct {
}

func (BookmarkRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkRelationState)(nil)).Elem()
}

type bookmarkRelationArgs struct {
	// Bookmark ID
	BookmarkId string `pulumi:"bookmarkId"`
	// The resource ID of the related resource
	RelatedResourceId string `pulumi:"relatedResourceId"`
	// Relation Name
	RelationName *string `pulumi:"relationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a BookmarkRelation resource.
type BookmarkRelationArgs struct {
	// Bookmark ID
	BookmarkId pulumi.StringInput
	// The resource ID of the related resource
	RelatedResourceId pulumi.StringInput
	// Relation Name
	RelationName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (BookmarkRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkRelationArgs)(nil)).Elem()
}

type BookmarkRelationInput interface {
	pulumi.Input

	ToBookmarkRelationOutput() BookmarkRelationOutput
	ToBookmarkRelationOutputWithContext(ctx context.Context) BookmarkRelationOutput
}

func (*BookmarkRelation) ElementType() reflect.Type {
	return reflect.TypeOf((**BookmarkRelation)(nil)).Elem()
}

func (i *BookmarkRelation) ToBookmarkRelationOutput() BookmarkRelationOutput {
	return i.ToBookmarkRelationOutputWithContext(context.Background())
}

func (i *BookmarkRelation) ToBookmarkRelationOutputWithContext(ctx context.Context) BookmarkRelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BookmarkRelationOutput)
}

type BookmarkRelationOutput struct{ *pulumi.OutputState }

func (BookmarkRelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BookmarkRelation)(nil)).Elem()
}

func (o BookmarkRelationOutput) ToBookmarkRelationOutput() BookmarkRelationOutput {
	return o
}

func (o BookmarkRelationOutput) ToBookmarkRelationOutputWithContext(ctx context.Context) BookmarkRelationOutput {
	return o
}

// Etag of the azure resource
func (o BookmarkRelationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o BookmarkRelationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the related resource
func (o BookmarkRelationOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.RelatedResourceId }).(pulumi.StringOutput)
}

// The resource kind of the related resource
func (o BookmarkRelationOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

// The name of the related resource
func (o BookmarkRelationOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.RelatedResourceName }).(pulumi.StringOutput)
}

// The resource type of the related resource
func (o BookmarkRelationOutput) RelatedResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.RelatedResourceType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BookmarkRelationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BookmarkRelation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BookmarkRelationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BookmarkRelationOutput{})
}

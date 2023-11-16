// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Hunt Relation in Azure Security Insights.
type HuntRelation struct {
	pulumi.CustomResourceState

	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// List of labels relevant to this hunt
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the related resource
	RelatedResourceId pulumi.StringOutput `pulumi:"relatedResourceId"`
	// The resource that the relation is related to
	RelatedResourceKind pulumi.StringOutput `pulumi:"relatedResourceKind"`
	// The name of the related resource
	RelatedResourceName pulumi.StringOutput `pulumi:"relatedResourceName"`
	// The type of the hunt relation
	RelationType pulumi.StringOutput `pulumi:"relationType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHuntRelation registers a new resource with the given unique name, arguments, and options.
func NewHuntRelation(ctx *pulumi.Context,
	name string, args *HuntRelationArgs, opts ...pulumi.ResourceOption) (*HuntRelation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HuntId == nil {
		return nil, errors.New("invalid value for required argument 'HuntId'")
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
			Type: pulumi.String("azure-native:securityinsights:HuntRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:HuntRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:HuntRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:HuntRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:HuntRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:HuntRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:HuntRelation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HuntRelation
	err := ctx.RegisterResource("azure-native:securityinsights/v20230801preview:HuntRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHuntRelation gets an existing HuntRelation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHuntRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HuntRelationState, opts ...pulumi.ResourceOption) (*HuntRelation, error) {
	var resource HuntRelation
	err := ctx.ReadResource("azure-native:securityinsights/v20230801preview:HuntRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HuntRelation resources.
type huntRelationState struct {
}

type HuntRelationState struct {
}

func (HuntRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*huntRelationState)(nil)).Elem()
}

type huntRelationArgs struct {
	// The hunt id (GUID)
	HuntId string `pulumi:"huntId"`
	// The hunt relation id (GUID)
	HuntRelationId *string `pulumi:"huntRelationId"`
	// List of labels relevant to this hunt
	Labels []string `pulumi:"labels"`
	// The id of the related resource
	RelatedResourceId string `pulumi:"relatedResourceId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a HuntRelation resource.
type HuntRelationArgs struct {
	// The hunt id (GUID)
	HuntId pulumi.StringInput
	// The hunt relation id (GUID)
	HuntRelationId pulumi.StringPtrInput
	// List of labels relevant to this hunt
	Labels pulumi.StringArrayInput
	// The id of the related resource
	RelatedResourceId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (HuntRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*huntRelationArgs)(nil)).Elem()
}

type HuntRelationInput interface {
	pulumi.Input

	ToHuntRelationOutput() HuntRelationOutput
	ToHuntRelationOutputWithContext(ctx context.Context) HuntRelationOutput
}

func (*HuntRelation) ElementType() reflect.Type {
	return reflect.TypeOf((**HuntRelation)(nil)).Elem()
}

func (i *HuntRelation) ToHuntRelationOutput() HuntRelationOutput {
	return i.ToHuntRelationOutputWithContext(context.Background())
}

func (i *HuntRelation) ToHuntRelationOutputWithContext(ctx context.Context) HuntRelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HuntRelationOutput)
}

type HuntRelationOutput struct{ *pulumi.OutputState }

func (HuntRelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HuntRelation)(nil)).Elem()
}

func (o HuntRelationOutput) ToHuntRelationOutput() HuntRelationOutput {
	return o
}

func (o HuntRelationOutput) ToHuntRelationOutputWithContext(ctx context.Context) HuntRelationOutput {
	return o
}

// Etag of the azure resource
func (o HuntRelationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HuntRelation) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// List of labels relevant to this hunt
func (o HuntRelationOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HuntRelation) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o HuntRelationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntRelation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the related resource
func (o HuntRelationOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntRelation) pulumi.StringOutput { return v.RelatedResourceId }).(pulumi.StringOutput)
}

// The resource that the relation is related to
func (o HuntRelationOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntRelation) pulumi.StringOutput { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

// The name of the related resource
func (o HuntRelationOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntRelation) pulumi.StringOutput { return v.RelatedResourceName }).(pulumi.StringOutput)
}

// The type of the hunt relation
func (o HuntRelationOutput) RelationType() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntRelation) pulumi.StringOutput { return v.RelationType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o HuntRelationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HuntRelation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HuntRelationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HuntRelation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HuntRelationOutput{})
}

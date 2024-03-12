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

// Represents a relation between two resources
type IncidentRelation struct {
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

// NewIncidentRelation registers a new resource with the given unique name, arguments, and options.
func NewIncidentRelation(ctx *pulumi.Context,
	name string, args *IncidentRelationArgs, opts ...pulumi.ResourceOption) (*IncidentRelation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncidentId == nil {
		return nil, errors.New("invalid value for required argument 'IncidentId'")
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
			Type: pulumi.String("azure-native:securityinsights:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:IncidentRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:IncidentRelation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IncidentRelation
	err := ctx.RegisterResource("azure-native:securityinsights/v20230801preview:IncidentRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIncidentRelation gets an existing IncidentRelation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIncidentRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentRelationState, opts ...pulumi.ResourceOption) (*IncidentRelation, error) {
	var resource IncidentRelation
	err := ctx.ReadResource("azure-native:securityinsights/v20230801preview:IncidentRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IncidentRelation resources.
type incidentRelationState struct {
}

type IncidentRelationState struct {
}

func (IncidentRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentRelationState)(nil)).Elem()
}

type incidentRelationArgs struct {
	// Incident ID
	IncidentId string `pulumi:"incidentId"`
	// The resource ID of the related resource
	RelatedResourceId string `pulumi:"relatedResourceId"`
	// Relation Name
	RelationName *string `pulumi:"relationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IncidentRelation resource.
type IncidentRelationArgs struct {
	// Incident ID
	IncidentId pulumi.StringInput
	// The resource ID of the related resource
	RelatedResourceId pulumi.StringInput
	// Relation Name
	RelationName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IncidentRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentRelationArgs)(nil)).Elem()
}

type IncidentRelationInput interface {
	pulumi.Input

	ToIncidentRelationOutput() IncidentRelationOutput
	ToIncidentRelationOutputWithContext(ctx context.Context) IncidentRelationOutput
}

func (*IncidentRelation) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentRelation)(nil)).Elem()
}

func (i *IncidentRelation) ToIncidentRelationOutput() IncidentRelationOutput {
	return i.ToIncidentRelationOutputWithContext(context.Background())
}

func (i *IncidentRelation) ToIncidentRelationOutputWithContext(ctx context.Context) IncidentRelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentRelationOutput)
}

type IncidentRelationOutput struct{ *pulumi.OutputState }

func (IncidentRelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentRelation)(nil)).Elem()
}

func (o IncidentRelationOutput) ToIncidentRelationOutput() IncidentRelationOutput {
	return o
}

func (o IncidentRelationOutput) ToIncidentRelationOutputWithContext(ctx context.Context) IncidentRelationOutput {
	return o
}

// Etag of the azure resource
func (o IncidentRelationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o IncidentRelationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the related resource
func (o IncidentRelationOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.RelatedResourceId }).(pulumi.StringOutput)
}

// The resource kind of the related resource
func (o IncidentRelationOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

// The name of the related resource
func (o IncidentRelationOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.RelatedResourceName }).(pulumi.StringOutput)
}

// The resource type of the related resource
func (o IncidentRelationOutput) RelatedResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.RelatedResourceType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IncidentRelationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IncidentRelation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IncidentRelationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IncidentRelation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentRelationOutput{})
}

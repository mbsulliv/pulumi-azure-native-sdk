// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Activity entity query.
// Azure REST API version: 2023-06-01-preview. Prior API version in Azure Native 1.x: 2021-03-01-preview.
type ActivityCustomEntityQuery struct {
	pulumi.CustomResourceState

	// The entity query content to display in timeline
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// The time the activity was created
	CreatedTimeUtc pulumi.StringOutput `pulumi:"createdTimeUtc"`
	// The entity query description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Determines whether this activity is enabled or disabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The query applied only to entities matching to all filters
	EntitiesFilter pulumi.StringArrayMapOutput `pulumi:"entitiesFilter"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The type of the query's source entity
	InputEntityType pulumi.StringPtrOutput `pulumi:"inputEntityType"`
	// The kind of the entity query
	// Expected value is 'Activity'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time the activity was updated
	LastModifiedTimeUtc pulumi.StringOutput `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Activity query definitions
	QueryDefinitions ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput `pulumi:"queryDefinitions"`
	// List of the fields of the source entity that are required to run the query
	RequiredInputFieldsSets pulumi.StringArrayArrayOutput `pulumi:"requiredInputFieldsSets"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The template id this activity was created from
	TemplateName pulumi.StringPtrOutput `pulumi:"templateName"`
	// The entity query title
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewActivityCustomEntityQuery registers a new resource with the given unique name, arguments, and options.
func NewActivityCustomEntityQuery(ctx *pulumi.Context,
	name string, args *ActivityCustomEntityQueryArgs, opts ...pulumi.ResourceOption) (*ActivityCustomEntityQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Activity")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:ActivityCustomEntityQuery"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ActivityCustomEntityQuery
	err := ctx.RegisterResource("azure-native:securityinsights:ActivityCustomEntityQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivityCustomEntityQuery gets an existing ActivityCustomEntityQuery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivityCustomEntityQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivityCustomEntityQueryState, opts ...pulumi.ResourceOption) (*ActivityCustomEntityQuery, error) {
	var resource ActivityCustomEntityQuery
	err := ctx.ReadResource("azure-native:securityinsights:ActivityCustomEntityQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActivityCustomEntityQuery resources.
type activityCustomEntityQueryState struct {
}

type ActivityCustomEntityQueryState struct {
}

func (ActivityCustomEntityQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*activityCustomEntityQueryState)(nil)).Elem()
}

type activityCustomEntityQueryArgs struct {
	// The entity query content to display in timeline
	Content *string `pulumi:"content"`
	// The entity query description
	Description *string `pulumi:"description"`
	// Determines whether this activity is enabled or disabled.
	Enabled *bool `pulumi:"enabled"`
	// The query applied only to entities matching to all filters
	EntitiesFilter map[string][]string `pulumi:"entitiesFilter"`
	// entity query ID
	EntityQueryId *string `pulumi:"entityQueryId"`
	// The type of the query's source entity
	InputEntityType *string `pulumi:"inputEntityType"`
	// The kind of the entity query that supports put request.
	// Expected value is 'Activity'.
	Kind string `pulumi:"kind"`
	// The Activity query definitions
	QueryDefinitions *ActivityEntityQueriesPropertiesQueryDefinitions `pulumi:"queryDefinitions"`
	// List of the fields of the source entity that are required to run the query
	RequiredInputFieldsSets [][]string `pulumi:"requiredInputFieldsSets"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The template id this activity was created from
	TemplateName *string `pulumi:"templateName"`
	// The entity query title
	Title *string `pulumi:"title"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ActivityCustomEntityQuery resource.
type ActivityCustomEntityQueryArgs struct {
	// The entity query content to display in timeline
	Content pulumi.StringPtrInput
	// The entity query description
	Description pulumi.StringPtrInput
	// Determines whether this activity is enabled or disabled.
	Enabled pulumi.BoolPtrInput
	// The query applied only to entities matching to all filters
	EntitiesFilter pulumi.StringArrayMapInput
	// entity query ID
	EntityQueryId pulumi.StringPtrInput
	// The type of the query's source entity
	InputEntityType pulumi.StringPtrInput
	// The kind of the entity query that supports put request.
	// Expected value is 'Activity'.
	Kind pulumi.StringInput
	// The Activity query definitions
	QueryDefinitions ActivityEntityQueriesPropertiesQueryDefinitionsPtrInput
	// List of the fields of the source entity that are required to run the query
	RequiredInputFieldsSets pulumi.StringArrayArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The template id this activity was created from
	TemplateName pulumi.StringPtrInput
	// The entity query title
	Title pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ActivityCustomEntityQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activityCustomEntityQueryArgs)(nil)).Elem()
}

type ActivityCustomEntityQueryInput interface {
	pulumi.Input

	ToActivityCustomEntityQueryOutput() ActivityCustomEntityQueryOutput
	ToActivityCustomEntityQueryOutputWithContext(ctx context.Context) ActivityCustomEntityQueryOutput
}

func (*ActivityCustomEntityQuery) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityCustomEntityQuery)(nil)).Elem()
}

func (i *ActivityCustomEntityQuery) ToActivityCustomEntityQueryOutput() ActivityCustomEntityQueryOutput {
	return i.ToActivityCustomEntityQueryOutputWithContext(context.Background())
}

func (i *ActivityCustomEntityQuery) ToActivityCustomEntityQueryOutputWithContext(ctx context.Context) ActivityCustomEntityQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityCustomEntityQueryOutput)
}

type ActivityCustomEntityQueryOutput struct{ *pulumi.OutputState }

func (ActivityCustomEntityQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityCustomEntityQuery)(nil)).Elem()
}

func (o ActivityCustomEntityQueryOutput) ToActivityCustomEntityQueryOutput() ActivityCustomEntityQueryOutput {
	return o
}

func (o ActivityCustomEntityQueryOutput) ToActivityCustomEntityQueryOutputWithContext(ctx context.Context) ActivityCustomEntityQueryOutput {
	return o
}

// The entity query content to display in timeline
func (o ActivityCustomEntityQueryOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// The time the activity was created
func (o ActivityCustomEntityQueryOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// The entity query description
func (o ActivityCustomEntityQueryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Determines whether this activity is enabled or disabled.
func (o ActivityCustomEntityQueryOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The query applied only to entities matching to all filters
func (o ActivityCustomEntityQueryOutput) EntitiesFilter() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringArrayMapOutput { return v.EntitiesFilter }).(pulumi.StringArrayMapOutput)
}

// Etag of the azure resource
func (o ActivityCustomEntityQueryOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The type of the query's source entity
func (o ActivityCustomEntityQueryOutput) InputEntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringPtrOutput { return v.InputEntityType }).(pulumi.StringPtrOutput)
}

// The kind of the entity query
// Expected value is 'Activity'.
func (o ActivityCustomEntityQueryOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time the activity was updated
func (o ActivityCustomEntityQueryOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o ActivityCustomEntityQueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Activity query definitions
func (o ActivityCustomEntityQueryOutput) QueryDefinitions() ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
		return v.QueryDefinitions
	}).(ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput)
}

// List of the fields of the source entity that are required to run the query
func (o ActivityCustomEntityQueryOutput) RequiredInputFieldsSets() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringArrayArrayOutput { return v.RequiredInputFieldsSets }).(pulumi.StringArrayArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ActivityCustomEntityQueryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The template id this activity was created from
func (o ActivityCustomEntityQueryOutput) TemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringPtrOutput { return v.TemplateName }).(pulumi.StringPtrOutput)
}

// The entity query title
func (o ActivityCustomEntityQueryOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ActivityCustomEntityQueryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ActivityCustomEntityQuery) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ActivityCustomEntityQueryOutput{})
}

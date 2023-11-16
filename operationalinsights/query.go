// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Log Analytics QueryPack-Query definition.
// Azure REST API version: 2019-09-01. Prior API version in Azure Native 1.x: 2019-09-01.
//
// Other available API versions: 2019-09-01-preview.
type Query struct {
	pulumi.CustomResourceState

	// Object Id of user creating the query.
	Author pulumi.StringOutput `pulumi:"author"`
	// Body of the query.
	Body pulumi.StringOutput `pulumi:"body"`
	// Description of the query.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique display name for your query within the Query Pack.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Additional properties that can be set for the query.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The related metadata items for the function.
	Related LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput `pulumi:"related"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags associated with the query.
	Tags pulumi.StringArrayMapOutput `pulumi:"tags"`
	// Creation Date for the Log Analytics Query, in ISO 8601 format.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Last modified date of the Log Analytics Query, in ISO 8601 format.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewQuery registers a new resource with the given unique name, arguments, and options.
func NewQuery(ctx *pulumi.Context,
	name string, args *QueryArgs, opts ...pulumi.ResourceOption) (*Query, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.QueryPackName == nil {
		return nil, errors.New("invalid value for required argument 'QueryPackName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190901:Query"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190901preview:Query"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Query
	err := ctx.RegisterResource("azure-native:operationalinsights:Query", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuery gets an existing Query resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryState, opts ...pulumi.ResourceOption) (*Query, error) {
	var resource Query
	err := ctx.ReadResource("azure-native:operationalinsights:Query", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Query resources.
type queryState struct {
}

type QueryState struct {
}

func (QueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryState)(nil)).Elem()
}

type queryArgs struct {
	// Body of the query.
	Body string `pulumi:"body"`
	// Description of the query.
	Description *string `pulumi:"description"`
	// Unique display name for your query within the Query Pack.
	DisplayName string `pulumi:"displayName"`
	// The id of a specific query defined in the Log Analytics QueryPack
	Id *string `pulumi:"id"`
	// Additional properties that can be set for the query.
	Properties interface{} `pulumi:"properties"`
	// The name of the Log Analytics QueryPack resource.
	QueryPackName string `pulumi:"queryPackName"`
	// The related metadata items for the function.
	Related *LogAnalyticsQueryPackQueryPropertiesRelated `pulumi:"related"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags associated with the query.
	Tags map[string][]string `pulumi:"tags"`
}

// The set of arguments for constructing a Query resource.
type QueryArgs struct {
	// Body of the query.
	Body pulumi.StringInput
	// Description of the query.
	Description pulumi.StringPtrInput
	// Unique display name for your query within the Query Pack.
	DisplayName pulumi.StringInput
	// The id of a specific query defined in the Log Analytics QueryPack
	Id pulumi.StringPtrInput
	// Additional properties that can be set for the query.
	Properties pulumi.Input
	// The name of the Log Analytics QueryPack resource.
	QueryPackName pulumi.StringInput
	// The related metadata items for the function.
	Related LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags associated with the query.
	Tags pulumi.StringArrayMapInput
}

func (QueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryArgs)(nil)).Elem()
}

type QueryInput interface {
	pulumi.Input

	ToQueryOutput() QueryOutput
	ToQueryOutputWithContext(ctx context.Context) QueryOutput
}

func (*Query) ElementType() reflect.Type {
	return reflect.TypeOf((**Query)(nil)).Elem()
}

func (i *Query) ToQueryOutput() QueryOutput {
	return i.ToQueryOutputWithContext(context.Background())
}

func (i *Query) ToQueryOutputWithContext(ctx context.Context) QueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryOutput)
}

type QueryOutput struct{ *pulumi.OutputState }

func (QueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Query)(nil)).Elem()
}

func (o QueryOutput) ToQueryOutput() QueryOutput {
	return o
}

func (o QueryOutput) ToQueryOutputWithContext(ctx context.Context) QueryOutput {
	return o
}

// Object Id of user creating the query.
func (o QueryOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

// Body of the query.
func (o QueryOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Body }).(pulumi.StringOutput)
}

// Description of the query.
func (o QueryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Query) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique display name for your query within the Query Pack.
func (o QueryOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Azure resource name
func (o QueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Additional properties that can be set for the query.
func (o QueryOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Query) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The related metadata items for the function.
func (o QueryOutput) Related() LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o.ApplyT(func(v *Query) LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput { return v.Related }).(LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput)
}

// Read only system data
func (o QueryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Query) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags associated with the query.
func (o QueryOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *Query) pulumi.StringArrayMapOutput { return v.Tags }).(pulumi.StringArrayMapOutput)
}

// Creation Date for the Log Analytics Query, in ISO 8601 format.
func (o QueryOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

// Last modified date of the Log Analytics Query, in ISO 8601 format.
func (o QueryOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

// Azure resource type
func (o QueryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueryOutput{})
}

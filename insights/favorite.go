// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Properties that define a favorite that is associated to an Application Insights component.
// Azure REST API version: 2015-05-01. Prior API version in Azure Native 1.x: 2015-05-01
type Favorite struct {
	pulumi.CustomResourceState

	// Favorite category, as defined by the user at creation time.
	Category pulumi.StringPtrOutput `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config pulumi.StringPtrOutput `pulumi:"config"`
	// Internally assigned unique id of the favorite definition.
	FavoriteId pulumi.StringOutput `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType pulumi.StringPtrOutput `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate pulumi.BoolPtrOutput `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The source of the favorite definition.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this favorite definition.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Unique user id of the specific user that owns this favorite.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewFavorite registers a new resource with the given unique name, arguments, and options.
func NewFavorite(ctx *pulumi.Context,
	name string, args *FavoriteArgs, opts ...pulumi.ResourceOption) (*Favorite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20150501:Favorite"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Favorite
	err := ctx.RegisterResource("azure-native:insights:Favorite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFavorite gets an existing Favorite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFavorite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FavoriteState, opts ...pulumi.ResourceOption) (*Favorite, error) {
	var resource Favorite
	err := ctx.ReadResource("azure-native:insights:Favorite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Favorite resources.
type favoriteState struct {
}

type FavoriteState struct {
}

func (FavoriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteState)(nil)).Elem()
}

type favoriteArgs struct {
	// Favorite category, as defined by the user at creation time.
	Category *string `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config *string `pulumi:"config"`
	// The Id of a specific favorite defined in the Application Insights component
	FavoriteId *string `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType *FavoriteType `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate *bool `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
	// The source of the favorite definition.
	SourceType *string `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags []string `pulumi:"tags"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Favorite resource.
type FavoriteArgs struct {
	// Favorite category, as defined by the user at creation time.
	Category pulumi.StringPtrInput
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config pulumi.StringPtrInput
	// The Id of a specific favorite defined in the Application Insights component
	FavoriteId pulumi.StringPtrInput
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType FavoriteTypePtrInput
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate pulumi.BoolPtrInput
	// The user-defined name of the favorite.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput
	// The source of the favorite definition.
	SourceType pulumi.StringPtrInput
	// A list of 0 or more tags that are associated with this favorite definition
	Tags pulumi.StringArrayInput
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version pulumi.StringPtrInput
}

func (FavoriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteArgs)(nil)).Elem()
}

type FavoriteInput interface {
	pulumi.Input

	ToFavoriteOutput() FavoriteOutput
	ToFavoriteOutputWithContext(ctx context.Context) FavoriteOutput
}

func (*Favorite) ElementType() reflect.Type {
	return reflect.TypeOf((**Favorite)(nil)).Elem()
}

func (i *Favorite) ToFavoriteOutput() FavoriteOutput {
	return i.ToFavoriteOutputWithContext(context.Background())
}

func (i *Favorite) ToFavoriteOutputWithContext(ctx context.Context) FavoriteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FavoriteOutput)
}

func (i *Favorite) ToOutput(ctx context.Context) pulumix.Output[*Favorite] {
	return pulumix.Output[*Favorite]{
		OutputState: i.ToFavoriteOutputWithContext(ctx).OutputState,
	}
}

type FavoriteOutput struct{ *pulumi.OutputState }

func (FavoriteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Favorite)(nil)).Elem()
}

func (o FavoriteOutput) ToFavoriteOutput() FavoriteOutput {
	return o
}

func (o FavoriteOutput) ToFavoriteOutputWithContext(ctx context.Context) FavoriteOutput {
	return o
}

func (o FavoriteOutput) ToOutput(ctx context.Context) pulumix.Output[*Favorite] {
	return pulumix.Output[*Favorite]{
		OutputState: o.OutputState,
	}
}

// Favorite category, as defined by the user at creation time.
func (o FavoriteOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
func (o FavoriteOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.Config }).(pulumi.StringPtrOutput)
}

// Internally assigned unique id of the favorite definition.
func (o FavoriteOutput) FavoriteId() pulumi.StringOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringOutput { return v.FavoriteId }).(pulumi.StringOutput)
}

// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
func (o FavoriteOutput) FavoriteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.FavoriteType }).(pulumi.StringPtrOutput)
}

// Flag denoting wether or not this favorite was generated from a template.
func (o FavoriteOutput) IsGeneratedFromTemplate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.BoolPtrOutput { return v.IsGeneratedFromTemplate }).(pulumi.BoolPtrOutput)
}

// The user-defined name of the favorite.
func (o FavoriteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The source of the favorite definition.
func (o FavoriteOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

// A list of 0 or more tags that are associated with this favorite definition
func (o FavoriteOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Date and time in UTC of the last modification that was made to this favorite definition.
func (o FavoriteOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

// Unique user id of the specific user that owns this favorite.
func (o FavoriteOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
func (o FavoriteOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Favorite) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FavoriteOutput{})
}

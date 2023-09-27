// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a single favorite by its FavoriteId, defined within an Application Insights component.
// Azure REST API version: 2015-05-01.
func LookupFavorite(ctx *pulumi.Context, args *LookupFavoriteArgs, opts ...pulumi.InvokeOption) (*LookupFavoriteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFavoriteResult
	err := ctx.Invoke("azure-native:insights:getFavorite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFavoriteArgs struct {
	// The Id of a specific favorite defined in the Application Insights component
	FavoriteId string `pulumi:"favoriteId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
}

// Properties that define a favorite that is associated to an Application Insights component.
type LookupFavoriteResult struct {
	// Favorite category, as defined by the user at creation time.
	Category *string `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config *string `pulumi:"config"`
	// Internally assigned unique id of the favorite definition.
	FavoriteId string `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType *string `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate *bool `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name *string `pulumi:"name"`
	// The source of the favorite definition.
	SourceType *string `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags []string `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this favorite definition.
	TimeModified string `pulumi:"timeModified"`
	// Unique user id of the specific user that owns this favorite.
	UserId string `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version *string `pulumi:"version"`
}

func LookupFavoriteOutput(ctx *pulumi.Context, args LookupFavoriteOutputArgs, opts ...pulumi.InvokeOption) LookupFavoriteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFavoriteResult, error) {
			args := v.(LookupFavoriteArgs)
			r, err := LookupFavorite(ctx, &args, opts...)
			var s LookupFavoriteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFavoriteResultOutput)
}

type LookupFavoriteOutputArgs struct {
	// The Id of a specific favorite defined in the Application Insights component
	FavoriteId pulumi.StringInput `pulumi:"favoriteId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupFavoriteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFavoriteArgs)(nil)).Elem()
}

// Properties that define a favorite that is associated to an Application Insights component.
type LookupFavoriteResultOutput struct{ *pulumi.OutputState }

func (LookupFavoriteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFavoriteResult)(nil)).Elem()
}

func (o LookupFavoriteResultOutput) ToLookupFavoriteResultOutput() LookupFavoriteResultOutput {
	return o
}

func (o LookupFavoriteResultOutput) ToLookupFavoriteResultOutputWithContext(ctx context.Context) LookupFavoriteResultOutput {
	return o
}

func (o LookupFavoriteResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFavoriteResult] {
	return pulumix.Output[LookupFavoriteResult]{
		OutputState: o.OutputState,
	}
}

// Favorite category, as defined by the user at creation time.
func (o LookupFavoriteResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
func (o LookupFavoriteResultOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.Config }).(pulumi.StringPtrOutput)
}

// Internally assigned unique id of the favorite definition.
func (o LookupFavoriteResultOutput) FavoriteId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteResult) string { return v.FavoriteId }).(pulumi.StringOutput)
}

// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
func (o LookupFavoriteResultOutput) FavoriteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.FavoriteType }).(pulumi.StringPtrOutput)
}

// Flag denoting wether or not this favorite was generated from a template.
func (o LookupFavoriteResultOutput) IsGeneratedFromTemplate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *bool { return v.IsGeneratedFromTemplate }).(pulumi.BoolPtrOutput)
}

// The user-defined name of the favorite.
func (o LookupFavoriteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The source of the favorite definition.
func (o LookupFavoriteResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

// A list of 0 or more tags that are associated with this favorite definition
func (o LookupFavoriteResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFavoriteResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// Date and time in UTC of the last modification that was made to this favorite definition.
func (o LookupFavoriteResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

// Unique user id of the specific user that owns this favorite.
func (o LookupFavoriteResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFavoriteResult) string { return v.UserId }).(pulumi.StringOutput)
}

// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
func (o LookupFavoriteResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFavoriteResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFavoriteResultOutput{})
}

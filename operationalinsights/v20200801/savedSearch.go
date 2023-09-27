// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Value object for saved search results.
type SavedSearch struct {
	pulumi.CustomResourceState

	// The category of the saved search. This helps the user to find a saved search faster.
	Category pulumi.StringOutput `pulumi:"category"`
	// Saved search display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The ETag of the saved search. To override an existing saved search, use "*" or specify the current Etag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The function alias if query serves as a function.
	FunctionAlias pulumi.StringPtrOutput `pulumi:"functionAlias"`
	// The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
	FunctionParameters pulumi.StringPtrOutput `pulumi:"functionParameters"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The query expression for the saved search.
	Query pulumi.StringOutput `pulumi:"query"`
	// The tags attached to the saved search.
	Tags TagResponseArrayOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The version number of the query language. The current version is 2 and is the default.
	Version pulumi.Float64PtrOutput `pulumi:"version"`
}

// NewSavedSearch registers a new resource with the given unique name, arguments, and options.
func NewSavedSearch(ctx *pulumi.Context,
	name string, args *SavedSearchArgs, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20150320:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:SavedSearch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SavedSearch
	err := ctx.RegisterResource("azure-native:operationalinsights/v20200801:SavedSearch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSavedSearch gets an existing SavedSearch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSavedSearch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SavedSearchState, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	var resource SavedSearch
	err := ctx.ReadResource("azure-native:operationalinsights/v20200801:SavedSearch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SavedSearch resources.
type savedSearchState struct {
}

type SavedSearchState struct {
}

func (SavedSearchState) ElementType() reflect.Type {
	return reflect.TypeOf((*savedSearchState)(nil)).Elem()
}

type savedSearchArgs struct {
	// The category of the saved search. This helps the user to find a saved search faster.
	Category string `pulumi:"category"`
	// Saved search display name.
	DisplayName string `pulumi:"displayName"`
	// The function alias if query serves as a function.
	FunctionAlias *string `pulumi:"functionAlias"`
	// The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
	FunctionParameters *string `pulumi:"functionParameters"`
	// The query expression for the saved search.
	Query string `pulumi:"query"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The id of the saved search.
	SavedSearchId *string `pulumi:"savedSearchId"`
	// The tags attached to the saved search.
	Tags []Tag `pulumi:"tags"`
	// The version number of the query language. The current version is 2 and is the default.
	Version *float64 `pulumi:"version"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SavedSearch resource.
type SavedSearchArgs struct {
	// The category of the saved search. This helps the user to find a saved search faster.
	Category pulumi.StringInput
	// Saved search display name.
	DisplayName pulumi.StringInput
	// The function alias if query serves as a function.
	FunctionAlias pulumi.StringPtrInput
	// The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
	FunctionParameters pulumi.StringPtrInput
	// The query expression for the saved search.
	Query pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The id of the saved search.
	SavedSearchId pulumi.StringPtrInput
	// The tags attached to the saved search.
	Tags TagArrayInput
	// The version number of the query language. The current version is 2 and is the default.
	Version pulumi.Float64PtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (SavedSearchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*savedSearchArgs)(nil)).Elem()
}

type SavedSearchInput interface {
	pulumi.Input

	ToSavedSearchOutput() SavedSearchOutput
	ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput
}

func (*SavedSearch) ElementType() reflect.Type {
	return reflect.TypeOf((**SavedSearch)(nil)).Elem()
}

func (i *SavedSearch) ToSavedSearchOutput() SavedSearchOutput {
	return i.ToSavedSearchOutputWithContext(context.Background())
}

func (i *SavedSearch) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavedSearchOutput)
}

func (i *SavedSearch) ToOutput(ctx context.Context) pulumix.Output[*SavedSearch] {
	return pulumix.Output[*SavedSearch]{
		OutputState: i.ToSavedSearchOutputWithContext(ctx).OutputState,
	}
}

type SavedSearchOutput struct{ *pulumi.OutputState }

func (SavedSearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SavedSearch)(nil)).Elem()
}

func (o SavedSearchOutput) ToSavedSearchOutput() SavedSearchOutput {
	return o
}

func (o SavedSearchOutput) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return o
}

func (o SavedSearchOutput) ToOutput(ctx context.Context) pulumix.Output[*SavedSearch] {
	return pulumix.Output[*SavedSearch]{
		OutputState: o.OutputState,
	}
}

// The category of the saved search. This helps the user to find a saved search faster.
func (o SavedSearchOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// Saved search display name.
func (o SavedSearchOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The ETag of the saved search. To override an existing saved search, use "*" or specify the current Etag
func (o SavedSearchOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The function alias if query serves as a function.
func (o SavedSearchOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringPtrOutput { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

// The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
func (o SavedSearchOutput) FunctionParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringPtrOutput { return v.FunctionParameters }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SavedSearchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The query expression for the saved search.
func (o SavedSearchOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// The tags attached to the saved search.
func (o SavedSearchOutput) Tags() TagResponseArrayOutput {
	return o.ApplyT(func(v *SavedSearch) TagResponseArrayOutput { return v.Tags }).(TagResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SavedSearchOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The version number of the query language. The current version is 2 and is the default.
func (o SavedSearchOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.Float64PtrOutput { return v.Version }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SavedSearchOutput{})
}

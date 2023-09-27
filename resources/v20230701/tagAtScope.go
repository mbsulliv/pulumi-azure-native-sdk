// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Wrapper resource for tags API requests and responses.
type TagAtScope struct {
	pulumi.CustomResourceState

	// The name of the tags wrapper resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of tags.
	Properties TagsResponseOutput `pulumi:"properties"`
	// The type of the tags wrapper resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagAtScope registers a new resource with the given unique name, arguments, and options.
func NewTagAtScope(ctx *pulumi.Context,
	name string, args *TagAtScopeArgs, opts ...pulumi.ResourceOption) (*TagAtScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:TagAtScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TagAtScope
	err := ctx.RegisterResource("azure-native:resources/v20230701:TagAtScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagAtScope gets an existing TagAtScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagAtScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagAtScopeState, opts ...pulumi.ResourceOption) (*TagAtScope, error) {
	var resource TagAtScope
	err := ctx.ReadResource("azure-native:resources/v20230701:TagAtScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagAtScope resources.
type tagAtScopeState struct {
}

type TagAtScopeState struct {
}

func (TagAtScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagAtScopeState)(nil)).Elem()
}

type tagAtScopeArgs struct {
	// The set of tags.
	Properties Tags `pulumi:"properties"`
	// The resource scope.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a TagAtScope resource.
type TagAtScopeArgs struct {
	// The set of tags.
	Properties TagsInput
	// The resource scope.
	Scope pulumi.StringInput
}

func (TagAtScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagAtScopeArgs)(nil)).Elem()
}

type TagAtScopeInput interface {
	pulumi.Input

	ToTagAtScopeOutput() TagAtScopeOutput
	ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput
}

func (*TagAtScope) ElementType() reflect.Type {
	return reflect.TypeOf((**TagAtScope)(nil)).Elem()
}

func (i *TagAtScope) ToTagAtScopeOutput() TagAtScopeOutput {
	return i.ToTagAtScopeOutputWithContext(context.Background())
}

func (i *TagAtScope) ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagAtScopeOutput)
}

func (i *TagAtScope) ToOutput(ctx context.Context) pulumix.Output[*TagAtScope] {
	return pulumix.Output[*TagAtScope]{
		OutputState: i.ToTagAtScopeOutputWithContext(ctx).OutputState,
	}
}

type TagAtScopeOutput struct{ *pulumi.OutputState }

func (TagAtScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagAtScope)(nil)).Elem()
}

func (o TagAtScopeOutput) ToTagAtScopeOutput() TagAtScopeOutput {
	return o
}

func (o TagAtScopeOutput) ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput {
	return o
}

func (o TagAtScopeOutput) ToOutput(ctx context.Context) pulumix.Output[*TagAtScope] {
	return pulumix.Output[*TagAtScope]{
		OutputState: o.OutputState,
	}
}

// The name of the tags wrapper resource.
func (o TagAtScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagAtScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The set of tags.
func (o TagAtScopeOutput) Properties() TagsResponseOutput {
	return o.ApplyT(func(v *TagAtScope) TagsResponseOutput { return v.Properties }).(TagsResponseOutput)
}

// The type of the tags wrapper resource.
func (o TagAtScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagAtScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagAtScopeOutput{})
}
